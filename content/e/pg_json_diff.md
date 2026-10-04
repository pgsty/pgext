---
title: "pg_json_diff"
linkTitle: "pg_json_diff"
description: "JSONB diff, JSON Patch, and JSON Merge Patch functions"
weight: 4160
categories: ["UTIL"]
languages: ["C++"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_json_diff**](https://github.com/KhaledSMQ/pg-jsondiff) : JSONB diff, JSON Patch, and JSON Merge Patch functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4160** | {{< badge content="pg_json_diff" link="https://github.com/KhaledSMQ/pg-jsondiff" >}} | {{< ext "pg_json_diff" >}} | `1.0` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgjq" >}} {{< ext "pgbson" >}} {{< ext "jsquery" >}} |

> [!Note] Source requires PostgreSQL 12+ and C++17. Packaging removes the unsupported control-file author field; no extension dependency.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_json_diff` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_json_diff_18" "green" >}} {{< bg "17" "pg_json_diff_17" "green" >}} {{< bg "16" "pg_json_diff_16" "green" >}} {{< bg "15" "pg_json_diff_15" "green" >}} {{< bg "14" "pg_json_diff_14" "green" >}} | `pg_json_diff_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-json-diff" "green" >}} {{< bg "17" "postgresql-17-pg-json-diff" "green" >}} {{< bg "16" "postgresql-16-pg-json-diff" "green" >}} {{< bg "15" "postgresql-15-pg-json-diff" "green" >}} {{< bg "14" "postgresql-14-pg-json-diff" "green" >}} | `postgresql-$v-pg-json-diff` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_json_diff_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-json-diff : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-json-diff : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-json-diff : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-json-diff : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-json-diff : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-json-diff : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-json-diff : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-json-diff : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-json-diff : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-json-diff : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-json-diff : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_json_diff_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 305.9 KiB | [pg_json_diff_18-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_json_diff_18-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_json_diff_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 298.2 KiB | [pg_json_diff_18-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_json_diff_18-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_json_diff_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 313.1 KiB | [pg_json_diff_18-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_json_diff_18-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_json_diff_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 307.2 KiB | [pg_json_diff_18-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_json_diff_18-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_json_diff_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 318.7 KiB | [pg_json_diff_18-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_json_diff_18-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_json_diff_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 310.1 KiB | [pg_json_diff_18-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_json_diff_18-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-json-diff` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 276.4 KiB | [postgresql-18-pg-json-diff_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-json-diff/postgresql-18-pg-json-diff_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-json-diff` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 269.8 KiB | [postgresql-18-pg-json-diff_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-json-diff/postgresql-18-pg-json-diff_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-json-diff` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 285.2 KiB | [postgresql-18-pg-json-diff_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-json-diff/postgresql-18-pg-json-diff_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-json-diff` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 277.2 KiB | [postgresql-18-pg-json-diff_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-json-diff/postgresql-18-pg-json-diff_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-json-diff` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 291.6 KiB | [postgresql-18-pg-json-diff_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-json-diff/postgresql-18-pg-json-diff_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-json-diff` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 288.5 KiB | [postgresql-18-pg-json-diff_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-json-diff/postgresql-18-pg-json-diff_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-json-diff` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 295.5 KiB | [postgresql-18-pg-json-diff_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-json-diff/postgresql-18-pg-json-diff_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-json-diff` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 291.2 KiB | [postgresql-18-pg-json-diff_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-json-diff/postgresql-18-pg-json-diff_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-json-diff` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 296.4 KiB | [postgresql-18-pg-json-diff_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-json-diff/postgresql-18-pg-json-diff_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-json-diff` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 289.1 KiB | [postgresql-18-pg-json-diff_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-json-diff/postgresql-18-pg-json-diff_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_json_diff_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 305.8 KiB | [pg_json_diff_17-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_json_diff_17-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_json_diff_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 297.8 KiB | [pg_json_diff_17-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_json_diff_17-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_json_diff_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 312.8 KiB | [pg_json_diff_17-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_json_diff_17-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_json_diff_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 306.7 KiB | [pg_json_diff_17-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_json_diff_17-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_json_diff_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 318.4 KiB | [pg_json_diff_17-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_json_diff_17-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_json_diff_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 309.6 KiB | [pg_json_diff_17-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_json_diff_17-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-json-diff` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 276.9 KiB | [postgresql-17-pg-json-diff_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-json-diff/postgresql-17-pg-json-diff_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-json-diff` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 270.5 KiB | [postgresql-17-pg-json-diff_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-json-diff/postgresql-17-pg-json-diff_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-json-diff` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 285.1 KiB | [postgresql-17-pg-json-diff_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-json-diff/postgresql-17-pg-json-diff_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-json-diff` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 277.1 KiB | [postgresql-17-pg-json-diff_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-json-diff/postgresql-17-pg-json-diff_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-json-diff` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 323.8 KiB | [postgresql-17-pg-json-diff_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-json-diff/postgresql-17-pg-json-diff_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-json-diff` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 321.0 KiB | [postgresql-17-pg-json-diff_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-json-diff/postgresql-17-pg-json-diff_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-json-diff` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 295.5 KiB | [postgresql-17-pg-json-diff_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-json-diff/postgresql-17-pg-json-diff_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-json-diff` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 291.5 KiB | [postgresql-17-pg-json-diff_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-json-diff/postgresql-17-pg-json-diff_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-json-diff` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 296.1 KiB | [postgresql-17-pg-json-diff_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-json-diff/postgresql-17-pg-json-diff_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-json-diff` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 288.8 KiB | [postgresql-17-pg-json-diff_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-json-diff/postgresql-17-pg-json-diff_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_json_diff_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 305.8 KiB | [pg_json_diff_16-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_json_diff_16-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_json_diff_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 297.8 KiB | [pg_json_diff_16-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_json_diff_16-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_json_diff_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 312.8 KiB | [pg_json_diff_16-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_json_diff_16-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_json_diff_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 306.7 KiB | [pg_json_diff_16-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_json_diff_16-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_json_diff_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 318.3 KiB | [pg_json_diff_16-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_json_diff_16-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_json_diff_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 309.6 KiB | [pg_json_diff_16-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_json_diff_16-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-json-diff` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 277.0 KiB | [postgresql-16-pg-json-diff_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-json-diff/postgresql-16-pg-json-diff_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-json-diff` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 270.4 KiB | [postgresql-16-pg-json-diff_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-json-diff/postgresql-16-pg-json-diff_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-json-diff` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 285.1 KiB | [postgresql-16-pg-json-diff_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-json-diff/postgresql-16-pg-json-diff_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-json-diff` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 277.1 KiB | [postgresql-16-pg-json-diff_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-json-diff/postgresql-16-pg-json-diff_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-json-diff` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 323.7 KiB | [postgresql-16-pg-json-diff_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-json-diff/postgresql-16-pg-json-diff_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-json-diff` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 321.0 KiB | [postgresql-16-pg-json-diff_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-json-diff/postgresql-16-pg-json-diff_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-json-diff` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 295.4 KiB | [postgresql-16-pg-json-diff_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-json-diff/postgresql-16-pg-json-diff_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-json-diff` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 291.5 KiB | [postgresql-16-pg-json-diff_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-json-diff/postgresql-16-pg-json-diff_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-json-diff` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 296.1 KiB | [postgresql-16-pg-json-diff_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-json-diff/postgresql-16-pg-json-diff_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-json-diff` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 288.8 KiB | [postgresql-16-pg-json-diff_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-json-diff/postgresql-16-pg-json-diff_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_json_diff_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 307.1 KiB | [pg_json_diff_15-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_json_diff_15-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_json_diff_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 298.7 KiB | [pg_json_diff_15-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_json_diff_15-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_json_diff_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 313.3 KiB | [pg_json_diff_15-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_json_diff_15-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_json_diff_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 306.9 KiB | [pg_json_diff_15-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_json_diff_15-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_json_diff_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 318.9 KiB | [pg_json_diff_15-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_json_diff_15-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_json_diff_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 310.0 KiB | [pg_json_diff_15-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_json_diff_15-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-json-diff` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 281.4 KiB | [postgresql-15-pg-json-diff_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-json-diff/postgresql-15-pg-json-diff_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-json-diff` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 274.8 KiB | [postgresql-15-pg-json-diff_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-json-diff/postgresql-15-pg-json-diff_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-json-diff` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 289.2 KiB | [postgresql-15-pg-json-diff_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-json-diff/postgresql-15-pg-json-diff_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-json-diff` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 280.6 KiB | [postgresql-15-pg-json-diff_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-json-diff/postgresql-15-pg-json-diff_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-json-diff` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 324.2 KiB | [postgresql-15-pg-json-diff_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-json-diff/postgresql-15-pg-json-diff_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-json-diff` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 321.1 KiB | [postgresql-15-pg-json-diff_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-json-diff/postgresql-15-pg-json-diff_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-json-diff` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 295.7 KiB | [postgresql-15-pg-json-diff_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-json-diff/postgresql-15-pg-json-diff_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-json-diff` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 291.7 KiB | [postgresql-15-pg-json-diff_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-json-diff/postgresql-15-pg-json-diff_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-json-diff` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 296.4 KiB | [postgresql-15-pg-json-diff_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-json-diff/postgresql-15-pg-json-diff_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-json-diff` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 289.2 KiB | [postgresql-15-pg-json-diff_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-json-diff/postgresql-15-pg-json-diff_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_json_diff_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 307.1 KiB | [pg_json_diff_14-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_json_diff_14-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_json_diff_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 298.9 KiB | [pg_json_diff_14-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_json_diff_14-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_json_diff_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 313.4 KiB | [pg_json_diff_14-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_json_diff_14-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_json_diff_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 307.0 KiB | [pg_json_diff_14-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_json_diff_14-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_json_diff_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 319.0 KiB | [pg_json_diff_14-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_json_diff_14-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_json_diff_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 310.1 KiB | [pg_json_diff_14-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_json_diff_14-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-json-diff` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 284.8 KiB | [postgresql-14-pg-json-diff_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-json-diff/postgresql-14-pg-json-diff_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-json-diff` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 278.4 KiB | [postgresql-14-pg-json-diff_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-json-diff/postgresql-14-pg-json-diff_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-json-diff` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 293.1 KiB | [postgresql-14-pg-json-diff_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-json-diff/postgresql-14-pg-json-diff_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-json-diff` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 284.3 KiB | [postgresql-14-pg-json-diff_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-json-diff/postgresql-14-pg-json-diff_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-json-diff` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 327.7 KiB | [postgresql-14-pg-json-diff_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-json-diff/postgresql-14-pg-json-diff_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-json-diff` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 324.5 KiB | [postgresql-14-pg-json-diff_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-json-diff/postgresql-14-pg-json-diff_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-json-diff` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 299.3 KiB | [postgresql-14-pg-json-diff_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-json-diff/postgresql-14-pg-json-diff_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-json-diff` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 295.4 KiB | [postgresql-14-pg-json-diff_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-json-diff/postgresql-14-pg-json-diff_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-json-diff` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 300.2 KiB | [postgresql-14-pg-json-diff_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-json-diff/postgresql-14-pg-json-diff_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-json-diff` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 292.7 KiB | [postgresql-14-pg-json-diff_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-json-diff/postgresql-14-pg-json-diff_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/KhaledSMQ/pg-jsondiff" title="Repository" icon="github" subtitle="github.com/KhaledSMQ/pg-jsondiff" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_json_diff-1.0-95cd668636ce.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_json_diff;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_json_diff;		# install via package name, for the active PG version

pig install pg_json_diff -v 18;   # install for PG 18
pig install pg_json_diff -v 17;   # install for PG 17
pig install pg_json_diff -v 16;   # install for PG 16
pig install pg_json_diff -v 15;   # install for PG 15
pig install pg_json_diff -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_json_diff;
```

## Usage

Sources:

- [Official README](https://github.com/KhaledSMQ/pg-jsondiff/blob/95cd668636ce9c88d2f6ac57d92ac5301a81ed3c/README.md)
- [Extension SQL](https://github.com/KhaledSMQ/pg-jsondiff/blob/95cd668636ce9c88d2f6ac57d92ac5301a81ed3c/pg_json_diff--1.0.sql)
- [C++ implementation](https://github.com/KhaledSMQ/pg-jsondiff/blob/95cd668636ce9c88d2f6ac57d92ac5301a81ed3c/src/json_diff_impl.cpp)

`pg_json_diff` 1.0 provides three immutable JSONB transformation functions backed by nlohmann/json: generate an RFC 6902 JSON Patch, apply such a patch, or apply an RFC 7386 merge patch. Use it for explicit document transformations, not as automatic conflict resolution or a stable audit-diff format.

### Core Workflow

Generate a patch, store or inspect it, and verify that applying it produces the expected target:

```sql
CREATE EXTENSION pg_json_diff;

WITH docs AS (
  SELECT '{"name":"Ada","age":30}'::jsonb AS old_doc,
         '{"name":"Ada","age":31,"city":"London"}'::jsonb AS new_doc
), patch AS (
  SELECT old_doc, new_doc, generate_json_diff(old_doc, new_doc) AS operations
  FROM docs
)
SELECT operations, apply_patch(old_doc, operations) = new_doc AS verified
FROM patch;
```

Validate externally supplied patches before applying them and run the transformation in a transaction with application-level authorization.

### Function Index

- `generate_json_diff(source, target)` returns an RFC 6902 operation array.
- `apply_patch(document, patch)` supports add, remove, replace, move, copy, and test operations.
- `merge_patch(document, patch)` recursively merges objects; a null member removes a key and arrays are replaced rather than merged.

All three functions accept and return `jsonb`.

### Operational Notes

Each call serializes PostgreSQL JSONB to text, parses it into a C++ JSON value, performs the operation, serializes again, and reparses the result as JSONB. Large or deeply nested documents therefore incur copying and parser cost. nlohmann/json and PostgreSQL JSONB have different numeric models, so test very large integers and high-precision decimals for fidelity. Patch errors are surfaced as general internal errors, and generated operation order or paths should not be treated as a canonical human audit record.
