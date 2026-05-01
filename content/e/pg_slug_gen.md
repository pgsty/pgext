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
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-slug-gen : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-slug-gen : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-slug-gen : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-slug-gen : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-slug-gen : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_slug_gen_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.7 KiB | [pg_slug_gen_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_slug_gen_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_slug_gen_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_slug_gen_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_slug_gen_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_slug_gen_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_slug_gen_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_slug_gen_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_slug_gen_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_slug_gen_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.5 KiB | [pg_slug_gen_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_slug_gen_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_slug_gen_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_slug_gen_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.0 KiB | [postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.0 KiB | [postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.0 KiB | [postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.0 KiB | [postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.0 KiB | [postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.1 KiB | [postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.1 KiB | [postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.4 KiB | [postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-slug-gen` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.4 KiB | [postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-slug-gen/postgresql-18-pg-slug-gen_1.0.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_slug_gen_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.7 KiB | [pg_slug_gen_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_slug_gen_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_slug_gen_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_slug_gen_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_slug_gen_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_slug_gen_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_slug_gen_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_slug_gen_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_slug_gen_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_slug_gen_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_slug_gen_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_slug_gen_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_slug_gen_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.0 KiB | [postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.0 KiB | [postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.0 KiB | [postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.0 KiB | [postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.2 KiB | [postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.1 KiB | [postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.3 KiB | [postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-slug-gen` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.4 KiB | [postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-slug-gen/postgresql-17-pg-slug-gen_1.0.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_slug_gen_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.7 KiB | [pg_slug_gen_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_slug_gen_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_slug_gen_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_slug_gen_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_slug_gen_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_slug_gen_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_slug_gen_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_slug_gen_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_slug_gen_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_slug_gen_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_slug_gen_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_slug_gen_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_slug_gen_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.0 KiB | [postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.0 KiB | [postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.0 KiB | [postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.0 KiB | [postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.2 KiB | [postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.1 KiB | [postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.3 KiB | [postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-slug-gen` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.4 KiB | [postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-slug-gen/postgresql-16-pg-slug-gen_1.0.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_slug_gen_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.7 KiB | [pg_slug_gen_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_slug_gen_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_slug_gen_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_slug_gen_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_slug_gen_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_slug_gen_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_slug_gen_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_slug_gen_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_slug_gen_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_slug_gen_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.6 KiB | [pg_slug_gen_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_slug_gen_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_slug_gen_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.8 KiB | [pg_slug_gen_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_slug_gen_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.0 KiB | [postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.0 KiB | [postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.0 KiB | [postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.0 KiB | [postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.1 KiB | [postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.3 KiB | [postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-slug-gen` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.4 KiB | [postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-slug-gen/postgresql-15-pg-slug-gen_1.0.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/fernandoolle/pg_slug_gen" title="Repository" icon="github" subtitle="github.com/fernandoolle/pg_slug_gen" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_slug_gen-1.0.0.tar.gz" >}}
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

Sources: [official PGXN release page](https://pgxn.org/dist/pg_slug_gen/), [official release README](https://api.pgxn.org/src/pg_slug_gen/pg_slug_gen-1.0.0/README.md), [official release SQL](https://api.pgxn.org/src/pg_slug_gen/pg_slug_gen-1.0.0/sql/pg_slug_gen--1.0.sql), [official release metadata](https://api.pgxn.org/src/pg_slug_gen/pg_slug_gen-1.0.0/META.json)

`pg_slug_gen` generates timestamp-based slugs with cryptographic randomness. The official 1.0.0 release describes it as a secure, URL-friendly short ID generator where the requested length selects the timestamp precision.

```sql
CREATE EXTENSION pg_slug_gen;

SELECT gen_random_slug();
SELECT gen_random_slug(13);
```

### Function

- `gen_random_slug(slug_length int DEFAULT 16) returns text`

The release SQL comment and README document these supported values:

- `10`: seconds
- `13`: milliseconds
- `16`: microseconds, also the default
- `19`: nanoseconds

### Precision And Format

Each precision maps to a timestamp width and a fixed slug shape:

- `10` digits: `5-5` format, 11 characters total
- `13` digits: `6-7` format, 14 characters total
- `16` digits: `8-8` format, 17 characters total
- `19` digits: `9-10` format, 20 characters total

The README states the collision-free window is bounded by timestamp precision: at most 1 insert per second, millisecond, microsecond, or nanosecond respectively.

### Examples

```sql
SELECT gen_random_slug();
SELECT gen_random_slug(10);
SELECT gen_random_slug(16);

CREATE TABLE products (
  id serial PRIMARY KEY,
  name text NOT NULL,
  slug text DEFAULT gen_random_slug() UNIQUE
);
```

### How It Works

The official README describes this algorithm:

- take the current timestamp at the chosen precision
- map each digit to a QWERTY-based character bucket
- choose one random character from that bucket with `pg_strong_random()`
- insert a hyphen at the midpoint

### Caveats

- This is a secure short-ID generator, not a text transliteration or title-to-URL slugifier.
- Same-timestamp collisions are still possible; upstream only claims uniqueness when inserts do not exceed one per chosen time unit.
- The official release metadata still points to `https://github.com/fernandoolle/pg_slug_gen`, but that repo URL currently returns 404.
