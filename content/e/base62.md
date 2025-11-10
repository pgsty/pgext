---
title: "base62"
linkTitle: "base62"
description: "Base62 extension for PostgreSQL"
weight: 4810
categories: ["FUNC"]
width: full
---

[**pg_base62**](https://github.com/adjust/pg-base62)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4810** | {{< badge content="base62" link="https://github.com/adjust/pg-base62" >}} | {{< ext "base62" "pg_base62" >}} | `0.0.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "base36" >}} {{< ext "pg_base58" >}} {{< ext "pg_polyline" >}} {{< ext "uri" >}} {{< ext "pg_curl" >}} {{< ext "url_encode" >}} {{< ext "pg_rewrite" >}} {{< ext "sepgsql" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/base62" >}} | `0.0.1` | {{< bg "18" "pg_base62_18*" "green" >}} {{< bg "17" "pg_base62_17*" "green" >}} {{< bg "16" "pg_base62_16*" "green" >}} {{< bg "15" "pg_base62_15*" "green" >}} {{< bg "14" "pg_base62_14*" "green" >}} {{< bg "13" "pg_base62_13*" "green" >}} | `pg_base62_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/base62" >}} | `0.0.1` | {{< bg "18" "postgresql-18-base62" "green" >}} {{< bg "17" "postgresql-17-base62" "green" >}} {{< bg "16" "postgresql-16-base62" "green" >}} {{< bg "15" "postgresql-15-base62" "green" >}} {{< bg "14" "postgresql-14-base62" "green" >}} {{< bg "13" "postgresql-13-base62" "green" >}} | `postgresql-$v-base62` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base62_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-base62 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-base62 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-base62 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-base62 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-base62 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-base62 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-base62 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-base62 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-base62 : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
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

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base62_13` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [pg_base62_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base62_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_base62_13` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.8 KiB | [pg_base62_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base62_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_base62_13` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.7 KiB | [pg_base62_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base62_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_base62_13` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [pg_base62_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base62_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_base62_13` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [pg_base62_13-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base62_13-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_base62_13` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.2 KiB | [pg_base62_13-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base62_13-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-base62` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.8 KiB | [postgresql-13-base62_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base62/postgresql-13-base62_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-base62` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 25.3 KiB | [postgresql-13-base62_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/base62/postgresql-13-base62_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-base62` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.2 KiB | [postgresql-13-base62_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/base62/postgresql-13-base62_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-base62` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.5 KiB | [postgresql-13-base62_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/base62/postgresql-13-base62_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-base62` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.2 KiB | [postgresql-13-base62_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base62/postgresql-13-base62_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-base62` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.9 KiB | [postgresql-13-base62_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/base62/postgresql-13-base62_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-base62` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.4 KiB | [postgresql-13-base62_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base62/postgresql-13-base62_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-base62` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.8 KiB | [postgresql-13-base62_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/base62/postgresql-13-base62_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adjust/pg-base62" title="Repository" icon="github" subtitle="github.com/adjust/pg-base62" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-base62-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get base62; # get base62 source code
pig build dep base62; # install build dependencies
pig build pkg base62; # build extension rpm or deb
pig build ext base62; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install base62; # install by extension name, for the current active PG version
pig ext install pg_base62; # install via package alias, for the active PG version
pig ext install base62 -v 18;   # install for PG 18
pig ext install base62 -v 17;   # install for PG 17
pig ext install base62 -v 16;   # install for PG 16
pig ext install base62 -v 15;   # install for PG 15
pig ext install base62 -v 14;   # install for PG 14
pig ext install base62 -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION base62;
```

