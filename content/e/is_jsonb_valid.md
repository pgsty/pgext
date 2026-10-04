---
title: "is_jsonb_valid"
linkTitle: "is_jsonb_valid"
description: "Native JSONB validation against JSON Schema drafts 4 and 7"
weight: 2765
categories: ["FEAT"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**is_jsonb_valid**](https://github.com/furstenheim/is_jsonb_valid) : Native JSONB validation against JSON Schema drafts 4 and 7


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2765** | {{< badge content="is_jsonb_valid" link="https://github.com/furstenheim/is_jsonb_valid" >}} | {{< ext "is_jsonb_valid" >}} | `0.1.4` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_jsonschema" >}} {{< ext "jsonschema" >}} {{< ext "postgres-json-schema" >}} {{< ext "jsquery" >}} |

> [!Note] Remote $ref and format validation are unsupported; local $ref resolution is limited. Ships two libraries under one extension identity.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `is_jsonb_valid` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.4` | {{< bg "18" "is_jsonb_valid_18" "green" >}} {{< bg "17" "is_jsonb_valid_17" "green" >}} {{< bg "16" "is_jsonb_valid_16" "green" >}} {{< bg "15" "is_jsonb_valid_15" "green" >}} {{< bg "14" "is_jsonb_valid_14" "green" >}} | `is_jsonb_valid_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.4` | {{< bg "18" "postgresql-18-is-jsonb-valid" "green" >}} {{< bg "17" "postgresql-17-is-jsonb-valid" "green" >}} {{< bg "16" "postgresql-16-is-jsonb-valid" "green" >}} {{< bg "15" "postgresql-15-is-jsonb-valid" "green" >}} {{< bg "14" "postgresql-14-is-jsonb-valid" "green" >}} | `postgresql-$v-is-jsonb-valid` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "is_jsonb_valid_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-18-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-17-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-16-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-15-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-14-is-jsonb-valid : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-18-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-17-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-16-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-15-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-14-is-jsonb-valid : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-18-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-17-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-16-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-15-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-14-is-jsonb-valid : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-18-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-17-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-16-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-15-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-14-is-jsonb-valid : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-18-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-17-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-16-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-15-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-14-is-jsonb-valid : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-18-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-17-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-16-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-15-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-14-is-jsonb-valid : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-18-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-17-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-16-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-15-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-14-is-jsonb-valid : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-18-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-17-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-16-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-15-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-14-is-jsonb-valid : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-18-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-17-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-16-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-15-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-14-is-jsonb-valid : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-18-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-17-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-16-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-15-is-jsonb-valid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "postgresql-14-is-jsonb-valid : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `is_jsonb_valid_18` | `0.1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.0 KiB | [is_jsonb_valid_18-0.1.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/is_jsonb_valid_18-0.1.4-1PGSTY.el8.x86_64.rpm) |
| `is_jsonb_valid_18` | `0.1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.4 KiB | [is_jsonb_valid_18-0.1.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/is_jsonb_valid_18-0.1.4-1PGSTY.el8.aarch64.rpm) |
| `is_jsonb_valid_18` | `0.1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 64.8 KiB | [is_jsonb_valid_18-0.1.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/is_jsonb_valid_18-0.1.4-1PGSTY.el9.x86_64.rpm) |
| `is_jsonb_valid_18` | `0.1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 64.2 KiB | [is_jsonb_valid_18-0.1.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/is_jsonb_valid_18-0.1.4-1PGSTY.el9.aarch64.rpm) |
| `is_jsonb_valid_18` | `0.1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 64.9 KiB | [is_jsonb_valid_18-0.1.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/is_jsonb_valid_18-0.1.4-1PGSTY.el10.x86_64.rpm) |
| `is_jsonb_valid_18` | `0.1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 64.2 KiB | [is_jsonb_valid_18-0.1.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/is_jsonb_valid_18-0.1.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-is-jsonb-valid` | `0.1.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.2 KiB | [postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/is-jsonb-valid/postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-is-jsonb-valid` | `0.1.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 49.6 KiB | [postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/is-jsonb-valid/postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-is-jsonb-valid` | `0.1.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.0 KiB | [postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/is-jsonb-valid/postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-is-jsonb-valid` | `0.1.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 49.8 KiB | [postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/is-jsonb-valid/postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-is-jsonb-valid` | `0.1.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 53.8 KiB | [postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/is-jsonb-valid/postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-is-jsonb-valid` | `0.1.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 53.2 KiB | [postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/is-jsonb-valid/postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-is-jsonb-valid` | `0.1.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 53.2 KiB | [postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/is-jsonb-valid/postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~noble_amd64.deb) |
| `postgresql-18-is-jsonb-valid` | `0.1.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 52.7 KiB | [postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/is-jsonb-valid/postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~noble_arm64.deb) |
| `postgresql-18-is-jsonb-valid` | `0.1.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 52.5 KiB | [postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/is-jsonb-valid/postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-is-jsonb-valid` | `0.1.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 51.9 KiB | [postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/is-jsonb-valid/postgresql-18-is-jsonb-valid_0.1.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `is_jsonb_valid_17` | `0.1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 62.9 KiB | [is_jsonb_valid_17-0.1.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/is_jsonb_valid_17-0.1.4-1PGSTY.el8.x86_64.rpm) |
| `is_jsonb_valid_17` | `0.1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.3 KiB | [is_jsonb_valid_17-0.1.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/is_jsonb_valid_17-0.1.4-1PGSTY.el8.aarch64.rpm) |
| `is_jsonb_valid_17` | `0.1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 64.6 KiB | [is_jsonb_valid_17-0.1.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/is_jsonb_valid_17-0.1.4-1PGSTY.el9.x86_64.rpm) |
| `is_jsonb_valid_17` | `0.1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 64.1 KiB | [is_jsonb_valid_17-0.1.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/is_jsonb_valid_17-0.1.4-1PGSTY.el9.aarch64.rpm) |
| `is_jsonb_valid_17` | `0.1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 64.9 KiB | [is_jsonb_valid_17-0.1.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/is_jsonb_valid_17-0.1.4-1PGSTY.el10.x86_64.rpm) |
| `is_jsonb_valid_17` | `0.1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 64.2 KiB | [is_jsonb_valid_17-0.1.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/is_jsonb_valid_17-0.1.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-is-jsonb-valid` | `0.1.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.2 KiB | [postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/is-jsonb-valid/postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-is-jsonb-valid` | `0.1.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 49.5 KiB | [postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/is-jsonb-valid/postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-is-jsonb-valid` | `0.1.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.1 KiB | [postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/is-jsonb-valid/postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-is-jsonb-valid` | `0.1.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 49.6 KiB | [postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/is-jsonb-valid/postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-is-jsonb-valid` | `0.1.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 59.9 KiB | [postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/is-jsonb-valid/postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-is-jsonb-valid` | `0.1.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 59.3 KiB | [postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/is-jsonb-valid/postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-is-jsonb-valid` | `0.1.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 53.3 KiB | [postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/is-jsonb-valid/postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~noble_amd64.deb) |
| `postgresql-17-is-jsonb-valid` | `0.1.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 52.6 KiB | [postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/is-jsonb-valid/postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~noble_arm64.deb) |
| `postgresql-17-is-jsonb-valid` | `0.1.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 52.4 KiB | [postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/is-jsonb-valid/postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-is-jsonb-valid` | `0.1.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 51.8 KiB | [postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/is-jsonb-valid/postgresql-17-is-jsonb-valid_0.1.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `is_jsonb_valid_16` | `0.1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 62.9 KiB | [is_jsonb_valid_16-0.1.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/is_jsonb_valid_16-0.1.4-1PGSTY.el8.x86_64.rpm) |
| `is_jsonb_valid_16` | `0.1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.3 KiB | [is_jsonb_valid_16-0.1.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/is_jsonb_valid_16-0.1.4-1PGSTY.el8.aarch64.rpm) |
| `is_jsonb_valid_16` | `0.1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 64.6 KiB | [is_jsonb_valid_16-0.1.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/is_jsonb_valid_16-0.1.4-1PGSTY.el9.x86_64.rpm) |
| `is_jsonb_valid_16` | `0.1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 64.1 KiB | [is_jsonb_valid_16-0.1.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/is_jsonb_valid_16-0.1.4-1PGSTY.el9.aarch64.rpm) |
| `is_jsonb_valid_16` | `0.1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 64.8 KiB | [is_jsonb_valid_16-0.1.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/is_jsonb_valid_16-0.1.4-1PGSTY.el10.x86_64.rpm) |
| `is_jsonb_valid_16` | `0.1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 64.2 KiB | [is_jsonb_valid_16-0.1.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/is_jsonb_valid_16-0.1.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-is-jsonb-valid` | `0.1.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.1 KiB | [postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/is-jsonb-valid/postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-is-jsonb-valid` | `0.1.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 49.6 KiB | [postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/is-jsonb-valid/postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-is-jsonb-valid` | `0.1.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.1 KiB | [postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/is-jsonb-valid/postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-is-jsonb-valid` | `0.1.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 49.7 KiB | [postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/is-jsonb-valid/postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-is-jsonb-valid` | `0.1.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 59.9 KiB | [postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/is-jsonb-valid/postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-is-jsonb-valid` | `0.1.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 59.3 KiB | [postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/is-jsonb-valid/postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-is-jsonb-valid` | `0.1.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 53.3 KiB | [postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/is-jsonb-valid/postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~noble_amd64.deb) |
| `postgresql-16-is-jsonb-valid` | `0.1.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 52.6 KiB | [postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/is-jsonb-valid/postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~noble_arm64.deb) |
| `postgresql-16-is-jsonb-valid` | `0.1.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 52.4 KiB | [postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/is-jsonb-valid/postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-is-jsonb-valid` | `0.1.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 51.8 KiB | [postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/is-jsonb-valid/postgresql-16-is-jsonb-valid_0.1.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `is_jsonb_valid_15` | `0.1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 62.9 KiB | [is_jsonb_valid_15-0.1.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/is_jsonb_valid_15-0.1.4-1PGSTY.el8.x86_64.rpm) |
| `is_jsonb_valid_15` | `0.1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.2 KiB | [is_jsonb_valid_15-0.1.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/is_jsonb_valid_15-0.1.4-1PGSTY.el8.aarch64.rpm) |
| `is_jsonb_valid_15` | `0.1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 64.7 KiB | [is_jsonb_valid_15-0.1.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/is_jsonb_valid_15-0.1.4-1PGSTY.el9.x86_64.rpm) |
| `is_jsonb_valid_15` | `0.1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 63.9 KiB | [is_jsonb_valid_15-0.1.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/is_jsonb_valid_15-0.1.4-1PGSTY.el9.aarch64.rpm) |
| `is_jsonb_valid_15` | `0.1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 64.8 KiB | [is_jsonb_valid_15-0.1.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/is_jsonb_valid_15-0.1.4-1PGSTY.el10.x86_64.rpm) |
| `is_jsonb_valid_15` | `0.1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 64.0 KiB | [is_jsonb_valid_15-0.1.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/is_jsonb_valid_15-0.1.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-is-jsonb-valid` | `0.1.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.0 KiB | [postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/is-jsonb-valid/postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-is-jsonb-valid` | `0.1.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 49.6 KiB | [postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/is-jsonb-valid/postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-is-jsonb-valid` | `0.1.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 50.9 KiB | [postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/is-jsonb-valid/postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-is-jsonb-valid` | `0.1.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 49.6 KiB | [postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/is-jsonb-valid/postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-is-jsonb-valid` | `0.1.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 60.0 KiB | [postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/is-jsonb-valid/postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-is-jsonb-valid` | `0.1.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 59.3 KiB | [postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/is-jsonb-valid/postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-is-jsonb-valid` | `0.1.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 53.1 KiB | [postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/is-jsonb-valid/postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~noble_amd64.deb) |
| `postgresql-15-is-jsonb-valid` | `0.1.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 52.5 KiB | [postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/is-jsonb-valid/postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~noble_arm64.deb) |
| `postgresql-15-is-jsonb-valid` | `0.1.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 52.5 KiB | [postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/is-jsonb-valid/postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-is-jsonb-valid` | `0.1.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 51.8 KiB | [postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/is-jsonb-valid/postgresql-15-is-jsonb-valid_0.1.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `is_jsonb_valid_14` | `0.1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.0 KiB | [is_jsonb_valid_14-0.1.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/is_jsonb_valid_14-0.1.4-1PGSTY.el8.x86_64.rpm) |
| `is_jsonb_valid_14` | `0.1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.4 KiB | [is_jsonb_valid_14-0.1.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/is_jsonb_valid_14-0.1.4-1PGSTY.el8.aarch64.rpm) |
| `is_jsonb_valid_14` | `0.1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 64.7 KiB | [is_jsonb_valid_14-0.1.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/is_jsonb_valid_14-0.1.4-1PGSTY.el9.x86_64.rpm) |
| `is_jsonb_valid_14` | `0.1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 64.1 KiB | [is_jsonb_valid_14-0.1.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/is_jsonb_valid_14-0.1.4-1PGSTY.el9.aarch64.rpm) |
| `is_jsonb_valid_14` | `0.1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 64.9 KiB | [is_jsonb_valid_14-0.1.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/is_jsonb_valid_14-0.1.4-1PGSTY.el10.x86_64.rpm) |
| `is_jsonb_valid_14` | `0.1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 64.2 KiB | [is_jsonb_valid_14-0.1.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/is_jsonb_valid_14-0.1.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-is-jsonb-valid` | `0.1.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.8 KiB | [postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/is-jsonb-valid/postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-is-jsonb-valid` | `0.1.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.4 KiB | [postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/is-jsonb-valid/postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-is-jsonb-valid` | `0.1.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.8 KiB | [postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/is-jsonb-valid/postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-is-jsonb-valid` | `0.1.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.4 KiB | [postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/is-jsonb-valid/postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-is-jsonb-valid` | `0.1.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 61.3 KiB | [postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/is-jsonb-valid/postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-is-jsonb-valid` | `0.1.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 60.5 KiB | [postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/is-jsonb-valid/postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-is-jsonb-valid` | `0.1.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 55.0 KiB | [postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/is-jsonb-valid/postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~noble_amd64.deb) |
| `postgresql-14-is-jsonb-valid` | `0.1.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 54.6 KiB | [postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/is-jsonb-valid/postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~noble_arm64.deb) |
| `postgresql-14-is-jsonb-valid` | `0.1.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 54.4 KiB | [postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/is-jsonb-valid/postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-is-jsonb-valid` | `0.1.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.8 KiB | [postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/is-jsonb-valid/postgresql-14-is-jsonb-valid_0.1.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/furstenheim/is_jsonb_valid" title="Repository" icon="github" subtitle="github.com/furstenheim/is_jsonb_valid" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="is_jsonb_valid-0.1.4-1f536758ae06.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg is_jsonb_valid;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install is_jsonb_valid;		# install via package name, for the active PG version

pig install is_jsonb_valid -v 18;   # install for PG 18
pig install is_jsonb_valid -v 17;   # install for PG 17
pig install is_jsonb_valid -v 16;   # install for PG 16
pig install is_jsonb_valid -v 15;   # install for PG 15
pig install is_jsonb_valid -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION is_jsonb_valid;
```

## Usage

Sources:

- [Official README at the catalog source revision](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/README.md)
- [Extension control file at the catalog source revision](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/is_jsonb_valid.control)
- [Version 0.1.4 extension SQL](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/is_jsonb_valid--0.1.4.sql)
- [PGXN distribution page](https://pgxn.org/dist/is_jsonb_valid/)

is_jsonb_valid provides native boolean validation of JSONB values against JSON Schema Draft 4 or Draft 7. It is suitable for checks and predicates that need only pass/fail output; it does not return a path or diagnostic explaining a failed keyword.

### Core Workflow

Choose the function that matches the schema draft and pass the schema first, followed by the value being validated.

```sql
CREATE EXTENSION is_jsonb_valid;

SELECT is_jsonb_valid(
    '{"type":"object","required":["id"],"properties":{"id":{"type":"integer"}}}'::jsonb,
    '{"id":42}'::jsonb
);

SELECT is_jsonb_valid_draft_v7(
    '{"if":{"exclusiveMaximum":0},"else":{"multipleOf":2}}'::jsonb,
    '4'::jsonb
);
```

The two functions are strict and immutable in version 0.1.4. This makes them usable in constraints and indexes when the schema is stable.

```sql
CREATE TABLE events (
    payload jsonb NOT NULL,
    CONSTRAINT events_payload_valid CHECK (
        is_jsonb_valid(
            '{"type":"object","required":["id"]}'::jsonb,
            payload
        )
    )
);
```

### Draft Boundary and Limitations

The base validator follows Draft 4; the explicitly named alternative follows Draft 7. Select deliberately, because keyword meaning differs between drafts and the extension does not infer the draft from a schema declaration.

At the pinned revision, remote references are unsupported. Local references are limited to definitions reachable from the schema root, and identifier changes along a reference chain are not resolved. Format validation is also unsupported. These gaps mean a true result is only a claim about the implemented subset, not every behavior in the complete JSON Schema specifications.

No preload or restart is declared. Keep schema literals and application expectations under version control, and add representative valid and invalid cases whenever a constraint depends on this extension.
