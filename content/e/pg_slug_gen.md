---
title: "pg_slug_gen"
linkTitle: "pg_slug_gen"
description: "Generate cryptographically secure timestamp-based slugs"
weight: 4550
categories: ["FUNC"]
width: full
---

[**pg_slug_gen**](https://github.com/fernandoolle/pg_slug_gen) : Generate cryptographically secure timestamp-based slugs


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4550** | {{< badge content="pg_slug_gen" link="https://github.com/fernandoolle/pg_slug_gen" >}} | {{< ext "pg_slug_gen" >}} | `1.0.0` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "uuid-ossp" >}} {{< ext "pg_uuidv7" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_slug_gen` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_slug_gen_18" "green" >}} {{< bg "17" "pg_slug_gen_17" "green" >}} {{< bg "16" "pg_slug_gen_16" "green" >}} {{< bg "15" "pg_slug_gen_15" "green" >}} {{< bg "14" "pg_slug_gen_14" "red" >}} | `pg_slug_gen_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pg-slug-gen" "green" >}} {{< bg "17" "postgresql-17-pg-slug-gen" "green" >}} {{< bg "16" "postgresql-16-pg-slug-gen" "green" >}} {{< bg "15" "postgresql-15-pg-slug-gen" "green" >}} {{< bg "14" "postgresql-14-pg-slug-gen" "red" >}} | `postgresql-$v-pg-slug-gen` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_slug_gen_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_slug_gen_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_slug_gen_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_slug_gen_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_slug_gen_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_slug_gen_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_slug_gen_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-slug-gen : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-slug-gen : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-slug-gen : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-slug-gen : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-slug-gen : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-slug-gen : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-slug-gen : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-slug-gen : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-slug-gen : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-slug-gen : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-slug-gen : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-slug-gen : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-slug-gen : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-slug-gen : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-slug-gen : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-slug-gen : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_slug_gen_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.7 KiB | [pg_slug_gen_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_slug_gen_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_slug_gen_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_slug_gen_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_slug_gen_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_slug_gen_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_slug_gen_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_slug_gen_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_slug_gen_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_slug_gen_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.5 KiB | [pg_slug_gen_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_slug_gen_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_slug_gen_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_slug_gen_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.1 KiB | [postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.0 KiB | [postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.0 KiB | [postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.0 KiB | [postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.0 KiB | [postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.0 KiB | [postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.1 KiB | [postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_slug_gen_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.7 KiB | [pg_slug_gen_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_slug_gen_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_slug_gen_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_slug_gen_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_slug_gen_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_slug_gen_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_slug_gen_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_slug_gen_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_slug_gen_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_slug_gen_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_slug_gen_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_slug_gen_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_slug_gen_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.0 KiB | [postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.0 KiB | [postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.0 KiB | [postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.0 KiB | [postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.1 KiB | [postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_slug_gen_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.7 KiB | [pg_slug_gen_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_slug_gen_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_slug_gen_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_slug_gen_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_slug_gen_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_slug_gen_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_slug_gen_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_slug_gen_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_slug_gen_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_slug_gen_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_slug_gen_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_slug_gen_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_slug_gen_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.0 KiB | [postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.0 KiB | [postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.0 KiB | [postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.0 KiB | [postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.1 KiB | [postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_slug_gen_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.7 KiB | [pg_slug_gen_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_slug_gen_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_slug_gen_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_slug_gen_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_slug_gen_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_slug_gen_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_slug_gen_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_slug_gen_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_slug_gen_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_slug_gen_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_slug_gen_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_slug_gen_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_slug_gen_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.0 KiB | [postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.0 KiB | [postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.0 KiB | [postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.9 KiB | [postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.1 KiB | [postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/fernandoolle/pg_slug_gen" title="Repository" icon="github" subtitle="github.com/fernandoolle/pg_slug_gen" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_slug_gen-1.0.0.zip" >}}
{{< /cards >}}


```bash
pig build pkg pg_slug_gen;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_slug_gen;		# install via package name, for the active PG version

pig install pg_slug_gen -v 18;   # install for PG 18
pig install pg_slug_gen -v 17;   # install for PG 17
pig install pg_slug_gen -v 16;   # install for PG 16
pig install pg_slug_gen -v 15;   # install for PG 15

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_slug_gen;
```


## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_slug_gen;
> SELECT gen_random_slug();
> SELECT gen_random_slug(13);
> ```
>
> Source: [PGXN release README](https://pgxn.org/dist/pg_slug_gen/1.0.0/)

`pg_slug_gen` generates timestamp-based slugs using cryptographic randomness. The PGXN release README describes it as a PostgreSQL extension that maps timestamp digits into letter buckets and inserts a hyphen in the middle, producing URL-friendly slugs.

## Function

The documented SQL function is:

```sql
gen_random_slug(slug_length int DEFAULT 16) RETURNS text
```

The README shows these interfaces:

```sql
gen_random_slug()      -- default: 16 (microseconds)
gen_random_slug(10)    -- seconds
gen_random_slug(13)    -- milliseconds
gen_random_slug(16)    -- microseconds
gen_random_slug(19)    -- nanoseconds
```

## Precision and Length

The release README maps precision to timestamp digits and maximum collision-free throughput:

- `10` digits for seconds, up to 1 insert per second
- `13` digits for milliseconds, up to 1,000 inserts per second
- `16` digits for microseconds, up to 1,000,000 inserts per second
- `19` digits for nanoseconds, up to 1 billion inserts per second

The slug includes a midpoint hyphen:

- seconds: `5-5` pattern, 11 characters total
- milliseconds: `6-7` pattern, 14 characters
- microseconds: `8-8` pattern, 17 characters
- nanoseconds: `9-10` pattern, 20 characters

## Examples

```sql
SELECT gen_random_slug();
SELECT gen_random_slug(10);
SELECT gen_random_slug(13);
SELECT gen_random_slug(16);
SELECT gen_random_slug(19);
```

As a table default:

```sql
CREATE TABLE products (
    id serial PRIMARY KEY,
    name text NOT NULL,
    slug text DEFAULT gen_random_slug() UNIQUE
);
```

## How It Works

The release README describes the algorithm as:

1. take the current timestamp at the chosen precision
2. map each digit to a QWERTY-based character bucket
3. choose one random character from that bucket using `pg_strong_random()`
4. insert a hyphen at the midpoint

The README also notes that same-timestamp collisions remain possible, but with microsecond precision the probability is stated as roughly 1 in 10 million.
