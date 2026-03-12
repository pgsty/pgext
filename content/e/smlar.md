---
title: "smlar"
linkTitle: "smlar"
description: "Effective similarity search"
weight: 1850
categories: ["RAG"]
width: full
---

[**smlar**](https://github.com/jirutka/smlar) : Effective similarity search


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1850** | {{< badge content="smlar" link="https://github.com/jirutka/smlar" >}} | {{< ext "smlar" >}} | `1.0` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_similarity" >}} {{< ext "fuzzystrmatch" >}} {{< ext "pg_trgm" >}} {{< ext "intarray" >}} {{< ext "vector" >}} {{< ext "pg_bigm" >}} {{< ext "unaccent" >}} {{< ext "vchord" >}} |

> [!Note] fix pg18 break issue by https://github.com/Vonng/smlar


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `smlar` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "smlar_18" "green" >}} {{< bg "17" "smlar_17" "green" >}} {{< bg "16" "smlar_16" "green" >}} {{< bg "15" "smlar_15" "green" >}} {{< bg "14" "smlar_14" "green" >}} | `smlar_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-smlar" "green" >}} {{< bg "17" "postgresql-17-smlar" "green" >}} {{< bg "16" "postgresql-16-smlar" "green" >}} {{< bg "15" "postgresql-15-smlar" "green" >}} {{< bg "14" "postgresql-14-smlar" "green" >}} | `postgresql-$v-smlar` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "smlar_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "smlar_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "smlar_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "smlar_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "smlar_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "smlar_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "smlar_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-smlar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-smlar : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `smlar_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.8 KiB | [smlar_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/smlar_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `smlar_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [smlar_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/smlar_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `smlar_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.6 KiB | [smlar_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/smlar_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `smlar_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.6 KiB | [smlar_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/smlar_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `smlar_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.9 KiB | [smlar_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/smlar_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `smlar_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.3 KiB | [smlar_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/smlar_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-smlar` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 71.7 KiB | [postgresql-18-smlar_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-18-smlar_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-smlar` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 70.1 KiB | [postgresql-18-smlar_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-18-smlar_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-smlar` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 71.7 KiB | [postgresql-18-smlar_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/smlar/postgresql-18-smlar_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-smlar` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 70.5 KiB | [postgresql-18-smlar_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/smlar/postgresql-18-smlar_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-smlar` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 77.4 KiB | [postgresql-18-smlar_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-18-smlar_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-smlar` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 76.1 KiB | [postgresql-18-smlar_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-18-smlar_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-smlar` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 75.4 KiB | [postgresql-18-smlar_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-18-smlar_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-smlar` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 74.9 KiB | [postgresql-18-smlar_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-18-smlar_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `smlar_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.7 KiB | [smlar_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/smlar_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `smlar_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [smlar_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/smlar_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `smlar_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.6 KiB | [smlar_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/smlar_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `smlar_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.6 KiB | [smlar_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/smlar_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `smlar_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.9 KiB | [smlar_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/smlar_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `smlar_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.3 KiB | [smlar_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/smlar_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-smlar` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 71.7 KiB | [postgresql-17-smlar_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-smlar` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 70.0 KiB | [postgresql-17-smlar_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-smlar` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 71.6 KiB | [postgresql-17-smlar_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-smlar` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 70.3 KiB | [postgresql-17-smlar_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-smlar` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 85.8 KiB | [postgresql-17-smlar_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-smlar` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 84.4 KiB | [postgresql-17-smlar_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-smlar` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 75.4 KiB | [postgresql-17-smlar_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-smlar` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 74.9 KiB | [postgresql-17-smlar_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-17-smlar_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `smlar_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.7 KiB | [smlar_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/smlar_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `smlar_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [smlar_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/smlar_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `smlar_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.6 KiB | [smlar_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/smlar_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `smlar_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.6 KiB | [smlar_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/smlar_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `smlar_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.9 KiB | [smlar_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/smlar_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `smlar_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.3 KiB | [smlar_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/smlar_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-smlar` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 71.6 KiB | [postgresql-16-smlar_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-smlar` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 70.0 KiB | [postgresql-16-smlar_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-smlar` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 71.6 KiB | [postgresql-16-smlar_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-smlar` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 70.3 KiB | [postgresql-16-smlar_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-smlar` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 85.6 KiB | [postgresql-16-smlar_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-smlar` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 84.2 KiB | [postgresql-16-smlar_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-smlar` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 75.3 KiB | [postgresql-16-smlar_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-smlar` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 74.9 KiB | [postgresql-16-smlar_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-16-smlar_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `smlar_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.2 KiB | [smlar_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/smlar_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `smlar_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.3 KiB | [smlar_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/smlar_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `smlar_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.1 KiB | [smlar_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/smlar_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `smlar_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.8 KiB | [smlar_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/smlar_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `smlar_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.9 KiB | [smlar_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/smlar_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `smlar_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.2 KiB | [smlar_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/smlar_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-smlar` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 72.1 KiB | [postgresql-15-smlar_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-smlar` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 70.3 KiB | [postgresql-15-smlar_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-smlar` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 71.9 KiB | [postgresql-15-smlar_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-smlar` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 70.5 KiB | [postgresql-15-smlar_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-smlar` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 86.2 KiB | [postgresql-15-smlar_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-smlar` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 84.6 KiB | [postgresql-15-smlar_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-smlar` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 75.0 KiB | [postgresql-15-smlar_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-smlar` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 74.4 KiB | [postgresql-15-smlar_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-15-smlar_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `smlar_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.1 KiB | [smlar_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/smlar_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `smlar_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.3 KiB | [smlar_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/smlar_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `smlar_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.1 KiB | [smlar_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/smlar_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `smlar_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 32.8 KiB | [smlar_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/smlar_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `smlar_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.9 KiB | [smlar_14-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/smlar_14-1.0-1PIGSTY.el10.x86_64.rpm) |
| `smlar_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.2 KiB | [smlar_14-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/smlar_14-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-smlar` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 72.1 KiB | [postgresql-14-smlar_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-smlar` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 70.4 KiB | [postgresql-14-smlar_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-smlar` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 71.9 KiB | [postgresql-14-smlar_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-smlar` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 70.5 KiB | [postgresql-14-smlar_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-smlar` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 86.1 KiB | [postgresql-14-smlar_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-smlar` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 84.5 KiB | [postgresql-14-smlar_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-smlar` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 75.0 KiB | [postgresql-14-smlar_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-smlar` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 74.4 KiB | [postgresql-14-smlar_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/smlar/postgresql-14-smlar_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jirutka/smlar" title="Repository" icon="github" subtitle="github.com/jirutka/smlar" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="smlar-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg smlar;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install smlar;		# install via package name, for the active PG version

pig install smlar -v 18;   # install for PG 18
pig install smlar -v 17;   # install for PG 17
pig install smlar -v 16;   # install for PG 16
pig install smlar -v 15;   # install for PG 15
pig install smlar -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION smlar;
```



## Usage

> [smlar](https://github.com/jirutka/smlar): Effective similarity search for PostgreSQL arrays.
> Source: [README](https://github.com/jirutka/smlar/blob/master/README)

The `smlar` extension provides effective similarity search on PostgreSQL arrays using configurable similarity formulas, GiST and GIN index support, and TF/IDF weighting.


--------

## Functions

```
float4 smlar(anyarray, anyarray)
```
Computes similarity of two arrays. Arrays should be the same type.

```
float4 smlar(anyarray, anyarray, bool useIntersect)
```
Computes similarity of two arrays of composite types. Composite type looks like:

```sql
CREATE TYPE type_name AS (element_name anytype, weight_name FLOAT4);
```

The `useIntersect` option points to use only intersected elements in the denominator.

```
float4 smlar(anyarray a, anyarray b, text formula)
```
Computes similarity of two arrays by a given formula. Predefined variables in formula:

- `N.i` -- number of common elements in both arrays (intersection)
- `N.a` -- number of unique elements in first array
- `N.b` -- number of unique elements in second array

Example:

```sql
SELECT smlar('{1,4,6}'::int[], '{5,4,6}');
SELECT smlar('{1,4,6}'::int[], '{5,4,6}', 'N.i / sqrt(N.a * N.b)');
-- These two calls are equivalent.
```

```
anyarray % anyarray
```
Returns true if similarity of the arrays is greater than the threshold limit.

```
text[] tsvector2textarray(tsvector)
```
Transforms tsvector type to text array.

```
anyarray array_unique(anyarray)
```
Sort and unique array.

```
float4 inarray(anyarray, anyelement)
```
Returns zero if second argument does not present in the first one and 1.0 in opposite case.

```
float4 inarray(anyarray, anyelement, float4, float4)
```
Returns fourth argument if second argument does not present in the first one and third argument in opposite case.


--------

## GUC Configuration Variables

```
smlar.threshold  FLOAT
```
Arrays with similarity lower than threshold are not similar by `%` operation.

```
smlar.persistent_cache  BOOL
```
Cache of global stat is stored in transaction-independent memory.

```
smlar.type  STRING
```
Type of similarity formula: `cosine` (default), `tfidf`, `overlap`.

```
smlar.stattable  STRING
```
Name of table storing set-wide statistic. Table should be defined as:

```sql
CREATE TABLE table_name (
    value   data_type UNIQUE,
    ndoc    int4 (or bigint)  NOT NULL CHECK (ndoc > 0)
);
```

A row with null value means total number of documents. Used only for `smlar.type = 'tfidf'`.

```
smlar.tf_method  STRING
```
Calculation method for term frequency. Values:
- `"n"` -- simple counting of entries (default)
- `"log"` -- 1 + log(n)
- `"const"` -- TF is equal to 1

Used only for `smlar.type = 'tfidf'`.

```
smlar.idf_plus_one  BOOL
```
If false (default), calculate idf as `log(d/df)`. If true, as `log(1+d/df)`. Used only for `smlar.type = 'tfidf'`.

It is highly recommended to add to `postgresql.conf`:

```
smlar.threshold = 0.6  # or any other value > 0 and < 1
```


--------

## GiST/GIN Index Support

The `%` and `&&` operations are supported with GiST and GIN indexes for many array types:

| Array Type | GIN operator class | GiST operator class |
|---|---|---|
| `bit[]` | `_bit_sml_ops` | |
| `bytea[]` | `_bytea_sml_ops` | `_bytea_sml_ops` |
| `char[]` | `_char_sml_ops` | `_char_sml_ops` |
| `cidr[]` | `_cidr_sml_ops` | `_cidr_sml_ops` |
| `date[]` | `_date_sml_ops` | `_date_sml_ops` |
| `float4[]` | `_float4_sml_ops` | `_float4_sml_ops` |
| `float8[]` | `_float8_sml_ops` | `_float8_sml_ops` |
| `inet[]` | `_inet_sml_ops` | `_inet_sml_ops` |
| `int2[]` | `_int2_sml_ops` | `_int2_sml_ops` |
| `int4[]` | `_int4_sml_ops` | `_int4_sml_ops` |
| `int8[]` | `_int8_sml_ops` | `_int8_sml_ops` |
| `interval[]` | `_interval_sml_ops` | `_interval_sml_ops` |
| `macaddr[]` | `_macaddr_sml_ops` | `_macaddr_sml_ops` |
| `money[]` | `_money_sml_ops` | |
| `numeric[]` | `_numeric_sml_ops` | `_numeric_sml_ops` |
| `oid[]` | `_oid_sml_ops` | `_oid_sml_ops` |
| `text[]` | `_text_sml_ops` | `_text_sml_ops` |
| `time[]` | `_time_sml_ops` | `_time_sml_ops` |
| `timestamp[]` | `_timestamp_sml_ops` | `_timestamp_sml_ops` |
| `timestamptz[]` | `_timestamptz_sml_ops` | `_timestamptz_sml_ops` |
| `timetz[]` | `_timetz_sml_ops` | `_timetz_sml_ops` |
| `varbit[]` | `_varbit_sml_ops` | |
| `varchar[]` | `_varchar_sml_ops` | `_varchar_sml_ops` |
