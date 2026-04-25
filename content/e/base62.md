---
title: "base62"
linkTitle: "base62"
description: "Base62 extension for PostgreSQL"
weight: 4810
categories: ["FUNC"]
width: full
---

[**pg_base62**](https://github.com/adjust/pg-base62) : Base62 extension for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4810** | {{< badge content="base62" link="https://github.com/adjust/pg-base62" >}} | {{< ext "base62" "pg_base62" >}} | `0.0.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "base36" >}} {{< ext "pg_base58" >}} {{< ext "pg_polyline" >}} {{< ext "uri" >}} {{< ext "pg_curl" >}} {{< ext "url_encode" >}} {{< ext "pg_rewrite" >}} {{< ext "sepgsql" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_base62` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "pg_base62_18" "green" >}} {{< bg "17" "pg_base62_17" "green" >}} {{< bg "16" "pg_base62_16" "green" >}} {{< bg "15" "pg_base62_15" "green" >}} {{< bg "14" "pg_base62_14" "green" >}} | `pg_base62_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "postgresql-18-base62" "green" >}} {{< bg "17" "postgresql-17-base62" "green" >}} {{< bg "16" "postgresql-16-base62" "green" >}} {{< bg "15" "postgresql-15-base62" "green" >}} {{< bg "14" "postgresql-14-base62" "green" >}} | `postgresql-$v-base62` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-base62 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-base62 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-base62 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-base62 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-base62 : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-base62 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-base62 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-base62 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-base62 : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-base62 : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base62_18` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.9 KiB | [pg_base62_18-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base62_18-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base62_18` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.8 KiB | [pg_base62_18-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base62_18-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base62_18` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_base62_18-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base62_18-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base62_18` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.9 KiB | [pg_base62_18-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base62_18-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_base62_18` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [pg_base62_18-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base62_18-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_base62_18` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.3 KiB | [pg_base62_18-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base62_18-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-base62` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.4 KiB | [postgresql-18-base62_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base62/postgresql-18-base62_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-base62` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 25.2 KiB | [postgresql-18-base62_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base62/postgresql-18-base62_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-base62` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.8 KiB | [postgresql-18-base62_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/base62/postgresql-18-base62_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-base62` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.6 KiB | [postgresql-18-base62_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/base62/postgresql-18-base62_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-base62` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.7 KiB | [postgresql-18-base62_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base62/postgresql-18-base62_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-base62` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.7 KiB | [postgresql-18-base62_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base62/postgresql-18-base62_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-base62` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.9 KiB | [postgresql-18-base62_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base62/postgresql-18-base62_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-base62` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.8 KiB | [postgresql-18-base62_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base62/postgresql-18-base62_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base62_17` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.9 KiB | [pg_base62_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base62_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base62_17` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.8 KiB | [pg_base62_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base62_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base62_17` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_base62_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base62_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base62_17` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.9 KiB | [pg_base62_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base62_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_base62_17` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [pg_base62_17-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base62_17-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_base62_17` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.3 KiB | [pg_base62_17-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base62_17-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-base62` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.4 KiB | [postgresql-17-base62_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base62/postgresql-17-base62_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-base62` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 25.2 KiB | [postgresql-17-base62_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base62/postgresql-17-base62_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-base62` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.8 KiB | [postgresql-17-base62_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/base62/postgresql-17-base62_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-base62` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.5 KiB | [postgresql-17-base62_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/base62/postgresql-17-base62_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-base62` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.1 KiB | [postgresql-17-base62_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base62/postgresql-17-base62_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-base62` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.1 KiB | [postgresql-17-base62_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base62/postgresql-17-base62_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-base62` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.0 KiB | [postgresql-17-base62_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base62/postgresql-17-base62_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-base62` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.8 KiB | [postgresql-17-base62_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base62/postgresql-17-base62_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base62_16` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.0 KiB | [pg_base62_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base62_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base62_16` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.8 KiB | [pg_base62_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base62_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base62_16` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_base62_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base62_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base62_16` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.9 KiB | [pg_base62_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base62_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_base62_16` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [pg_base62_16-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base62_16-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_base62_16` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.3 KiB | [pg_base62_16-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base62_16-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-base62` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.4 KiB | [postgresql-16-base62_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base62/postgresql-16-base62_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-base62` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 25.3 KiB | [postgresql-16-base62_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base62/postgresql-16-base62_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-base62` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.8 KiB | [postgresql-16-base62_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/base62/postgresql-16-base62_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-base62` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.6 KiB | [postgresql-16-base62_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/base62/postgresql-16-base62_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-base62` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.1 KiB | [postgresql-16-base62_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base62/postgresql-16-base62_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-base62` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.1 KiB | [postgresql-16-base62_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base62/postgresql-16-base62_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-base62` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.0 KiB | [postgresql-16-base62_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base62/postgresql-16-base62_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-base62` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.8 KiB | [postgresql-16-base62_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base62/postgresql-16-base62_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base62_15` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.0 KiB | [pg_base62_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base62_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base62_15` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.8 KiB | [pg_base62_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base62_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base62_15` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_base62_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base62_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base62_15` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.9 KiB | [pg_base62_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base62_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_base62_15` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [pg_base62_15-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base62_15-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_base62_15` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.3 KiB | [pg_base62_15-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base62_15-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-base62` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.5 KiB | [postgresql-15-base62_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base62/postgresql-15-base62_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-base62` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 25.4 KiB | [postgresql-15-base62_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base62/postgresql-15-base62_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-base62` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.9 KiB | [postgresql-15-base62_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/base62/postgresql-15-base62_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-base62` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.7 KiB | [postgresql-15-base62_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/base62/postgresql-15-base62_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-base62` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.1 KiB | [postgresql-15-base62_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base62/postgresql-15-base62_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-base62` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.1 KiB | [postgresql-15-base62_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base62/postgresql-15-base62_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-base62` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.1 KiB | [postgresql-15-base62_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base62/postgresql-15-base62_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-base62` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.0 KiB | [postgresql-15-base62_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base62/postgresql-15-base62_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base62_14` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.9 KiB | [pg_base62_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base62_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base62_14` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.8 KiB | [pg_base62_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base62_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base62_14` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_base62_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base62_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base62_14` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.9 KiB | [pg_base62_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base62_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_base62_14` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [pg_base62_14-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base62_14-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_base62_14` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.3 KiB | [pg_base62_14-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base62_14-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-base62` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.2 KiB | [postgresql-14-base62_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base62/postgresql-14-base62_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-base62` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 25.3 KiB | [postgresql-14-base62_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base62/postgresql-14-base62_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-base62` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.6 KiB | [postgresql-14-base62_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/base62/postgresql-14-base62_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-base62` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.6 KiB | [postgresql-14-base62_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/base62/postgresql-14-base62_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-base62` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.1 KiB | [postgresql-14-base62_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base62/postgresql-14-base62_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-base62` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.1 KiB | [postgresql-14-base62_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base62/postgresql-14-base62_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-base62` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.8 KiB | [postgresql-14-base62_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base62/postgresql-14-base62_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-base62` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.8 KiB | [postgresql-14-base62_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base62/postgresql-14-base62_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adjust/pg-base62" title="Repository" icon="github" subtitle="github.com/adjust/pg-base62" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-base62-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_base62;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_base62;		# install via package name, for the active PG version
pig install base62;		# install by extension name, for the current active PG version

pig install base62 -v 18;   # install for PG 18
pig install base62 -v 17;   # install for PG 17
pig install base62 -v 16;   # install for PG 16
pig install base62 -v 15;   # install for PG 15
pig install base62 -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION base62;
```



## Usage

> [base62: base62 encoding/decoding data types for PostgreSQL](https://github.com/adjust/pg-base62)

Provides data types for encoding and decoding values using the base62 scheme (0-9, A-Z, a-z).

```sql
CREATE EXTENSION base62;
```

### Types

| Type | Storage | Max String Length | Max Numeric Value |
|---|---|---|---|
| `base62` | 4 bytes (int) | 6 characters | 2,147,483,647 |
| `bigbase62` | 8 bytes (bigint) | 11 characters | 9,223,372,036,854,775,807 |
| `hugebase62` | 16 bytes | 20 characters | (bytea conversion) |

### Examples

```sql
-- Encode/decode base62
SELECT 2147483647::base62;          -- '2LKcb1'
SELECT '2LKcb1'::base62::int;      -- 2147483647

-- Bigbase62 for larger values
SELECT 9223372036854775807::bigbase62;           -- 'AzL8n0Y58m7'
SELECT 'AzL8n0Y58m7'::bigbase62::bigint;        -- 9223372036854775807

-- Hugebase62 with bytea conversion
SELECT 'AzL8n0Y58m7AzL8n0Y58'::hugebase62;
SELECT 'AzL8n0Y58m7AzL8n0Y58'::hugebase62::bytea;
SELECT '\x960c06065a6ed8ffff1e7149f40b1800'::bytea::hugebase62;

-- Note: base62 is case-sensitive
SELECT '2lkcb'::base62::int;   -- 40933305
SELECT '2LKCB'::base62::int;   -- 34635195
```
