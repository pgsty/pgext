---
title: "pg_liquid"
linkTitle: "pg_liquid"
description: "Liquid-inspired Datalog graph query extension for PostgreSQL"
weight: 2705
categories: ["FEAT"]
width: full
---

[**pg_liquid**](https://github.com/michael-golfi/pg_liquid) : Liquid-inspired Datalog graph query extension for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2705** | {{< badge content="pg_liquid" link="https://github.com/michael-golfi/pg_liquid" >}} | {{< ext "pg_liquid" >}} | `0.1.7` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `liquid` |
|   **See Also**    | {{< ext "age" >}} {{< ext "jsquery" >}} {{< ext "pg_jsonschema" >}} {{< ext "pg_search" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_liquid` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.7` | {{< bg "18" "pg_liquid_18" "green" >}} {{< bg "17" "pg_liquid_17" "green" >}} {{< bg "16" "pg_liquid_16" "green" >}} {{< bg "15" "pg_liquid_15" "green" >}} {{< bg "14" "pg_liquid_14" "green" >}} | `pg_liquid_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.7` | {{< bg "18" "postgresql-18-pg-liquid" "green" >}} {{< bg "17" "postgresql-17-pg-liquid" "green" >}} {{< bg "16" "postgresql-16-pg-liquid" "green" >}} {{< bg "15" "postgresql-15-pg-liquid" "green" >}} {{< bg "14" "postgresql-14-pg-liquid" "green" >}} | `postgresql-$v-pg-liquid` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "pg_liquid_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-liquid : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-liquid : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-liquid : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-liquid : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-liquid : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-liquid : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-liquid : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-18-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-17-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-16-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-15-pg-liquid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.7" "postgresql-14-pg-liquid : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_liquid_18` | `0.1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 64.7 KiB | [pg_liquid_18-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_liquid_18-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_liquid_18` | `0.1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.0 KiB | [pg_liquid_18-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_liquid_18-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_liquid_18` | `0.1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.7 KiB | [pg_liquid_18-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_liquid_18-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_liquid_18` | `0.1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.8 KiB | [pg_liquid_18-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_liquid_18-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_liquid_18` | `0.1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 65.0 KiB | [pg_liquid_18-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_liquid_18-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_liquid_18` | `0.1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 62.4 KiB | [pg_liquid_18-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_liquid_18-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-liquid` | `0.1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 383.1 KiB | [postgresql-18-pg-liquid_0.1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-liquid/postgresql-18-pg-liquid_0.1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-liquid` | `0.1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 378.0 KiB | [postgresql-18-pg-liquid_0.1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-liquid/postgresql-18-pg-liquid_0.1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-liquid` | `0.1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 383.5 KiB | [postgresql-18-pg-liquid_0.1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-liquid/postgresql-18-pg-liquid_0.1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-liquid` | `0.1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 378.6 KiB | [postgresql-18-pg-liquid_0.1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-liquid/postgresql-18-pg-liquid_0.1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-liquid` | `0.1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 429.6 KiB | [postgresql-18-pg-liquid_0.1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-liquid/postgresql-18-pg-liquid_0.1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-liquid` | `0.1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 427.6 KiB | [postgresql-18-pg-liquid_0.1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-liquid/postgresql-18-pg-liquid_0.1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-liquid` | `0.1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 407.1 KiB | [postgresql-18-pg-liquid_0.1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-liquid/postgresql-18-pg-liquid_0.1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-liquid` | `0.1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 402.7 KiB | [postgresql-18-pg-liquid_0.1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-liquid/postgresql-18-pg-liquid_0.1.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_liquid_17` | `0.1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 64.7 KiB | [pg_liquid_17-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_liquid_17-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_liquid_17` | `0.1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.0 KiB | [pg_liquid_17-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_liquid_17-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_liquid_17` | `0.1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.7 KiB | [pg_liquid_17-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_liquid_17-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_liquid_17` | `0.1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.8 KiB | [pg_liquid_17-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_liquid_17-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_liquid_17` | `0.1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 65.0 KiB | [pg_liquid_17-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_liquid_17-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_liquid_17` | `0.1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 62.4 KiB | [pg_liquid_17-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_liquid_17-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-liquid` | `0.1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 379.0 KiB | [postgresql-17-pg-liquid_0.1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-liquid/postgresql-17-pg-liquid_0.1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-liquid` | `0.1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 376.7 KiB | [postgresql-17-pg-liquid_0.1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-liquid/postgresql-17-pg-liquid_0.1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-liquid` | `0.1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 379.7 KiB | [postgresql-17-pg-liquid_0.1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-liquid/postgresql-17-pg-liquid_0.1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-liquid` | `0.1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 377.4 KiB | [postgresql-17-pg-liquid_0.1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-liquid/postgresql-17-pg-liquid_0.1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-liquid` | `0.1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 450.1 KiB | [postgresql-17-pg-liquid_0.1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-liquid/postgresql-17-pg-liquid_0.1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-liquid` | `0.1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 450.4 KiB | [postgresql-17-pg-liquid_0.1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-liquid/postgresql-17-pg-liquid_0.1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-liquid` | `0.1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 400.9 KiB | [postgresql-17-pg-liquid_0.1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-liquid/postgresql-17-pg-liquid_0.1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-liquid` | `0.1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 401.5 KiB | [postgresql-17-pg-liquid_0.1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-liquid/postgresql-17-pg-liquid_0.1.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_liquid_16` | `0.1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 64.7 KiB | [pg_liquid_16-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_liquid_16-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_liquid_16` | `0.1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.0 KiB | [pg_liquid_16-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_liquid_16-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_liquid_16` | `0.1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.7 KiB | [pg_liquid_16-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_liquid_16-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_liquid_16` | `0.1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.8 KiB | [pg_liquid_16-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_liquid_16-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_liquid_16` | `0.1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 64.9 KiB | [pg_liquid_16-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_liquid_16-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_liquid_16` | `0.1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 62.4 KiB | [pg_liquid_16-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_liquid_16-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-liquid` | `0.1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 372.0 KiB | [postgresql-16-pg-liquid_0.1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-liquid/postgresql-16-pg-liquid_0.1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-liquid` | `0.1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 368.1 KiB | [postgresql-16-pg-liquid_0.1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-liquid/postgresql-16-pg-liquid_0.1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-liquid` | `0.1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 372.4 KiB | [postgresql-16-pg-liquid_0.1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-liquid/postgresql-16-pg-liquid_0.1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-liquid` | `0.1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 368.5 KiB | [postgresql-16-pg-liquid_0.1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-liquid/postgresql-16-pg-liquid_0.1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-liquid` | `0.1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 437.0 KiB | [postgresql-16-pg-liquid_0.1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-liquid/postgresql-16-pg-liquid_0.1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-liquid` | `0.1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 436.2 KiB | [postgresql-16-pg-liquid_0.1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-liquid/postgresql-16-pg-liquid_0.1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-liquid` | `0.1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 393.9 KiB | [postgresql-16-pg-liquid_0.1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-liquid/postgresql-16-pg-liquid_0.1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-liquid` | `0.1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 391.3 KiB | [postgresql-16-pg-liquid_0.1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-liquid/postgresql-16-pg-liquid_0.1.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_liquid_15` | `0.1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.4 KiB | [pg_liquid_15-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_liquid_15-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_liquid_15` | `0.1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.5 KiB | [pg_liquid_15-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_liquid_15-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_liquid_15` | `0.1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 64.0 KiB | [pg_liquid_15-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_liquid_15-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_liquid_15` | `0.1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 62.8 KiB | [pg_liquid_15-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_liquid_15-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_liquid_15` | `0.1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 66.2 KiB | [pg_liquid_15-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_liquid_15-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_liquid_15` | `0.1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 63.6 KiB | [pg_liquid_15-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_liquid_15-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-liquid` | `0.1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 370.4 KiB | [postgresql-15-pg-liquid_0.1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-liquid/postgresql-15-pg-liquid_0.1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-liquid` | `0.1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 367.5 KiB | [postgresql-15-pg-liquid_0.1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-liquid/postgresql-15-pg-liquid_0.1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-liquid` | `0.1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 370.7 KiB | [postgresql-15-pg-liquid_0.1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-liquid/postgresql-15-pg-liquid_0.1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-liquid` | `0.1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 368.0 KiB | [postgresql-15-pg-liquid_0.1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-liquid/postgresql-15-pg-liquid_0.1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-liquid` | `0.1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 434.7 KiB | [postgresql-15-pg-liquid_0.1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-liquid/postgresql-15-pg-liquid_0.1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-liquid` | `0.1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 435.0 KiB | [postgresql-15-pg-liquid_0.1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-liquid/postgresql-15-pg-liquid_0.1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-liquid` | `0.1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 391.1 KiB | [postgresql-15-pg-liquid_0.1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-liquid/postgresql-15-pg-liquid_0.1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-liquid` | `0.1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 391.8 KiB | [postgresql-15-pg-liquid_0.1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-liquid/postgresql-15-pg-liquid_0.1.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_liquid_14` | `0.1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.4 KiB | [pg_liquid_14-0.1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_liquid_14-0.1.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_liquid_14` | `0.1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.6 KiB | [pg_liquid_14-0.1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_liquid_14-0.1.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_liquid_14` | `0.1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 64.0 KiB | [pg_liquid_14-0.1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_liquid_14-0.1.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_liquid_14` | `0.1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 62.8 KiB | [pg_liquid_14-0.1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_liquid_14-0.1.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_liquid_14` | `0.1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 66.1 KiB | [pg_liquid_14-0.1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_liquid_14-0.1.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_liquid_14` | `0.1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 63.6 KiB | [pg_liquid_14-0.1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_liquid_14-0.1.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-liquid` | `0.1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 366.4 KiB | [postgresql-14-pg-liquid_0.1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-liquid/postgresql-14-pg-liquid_0.1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-liquid` | `0.1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 363.5 KiB | [postgresql-14-pg-liquid_0.1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-liquid/postgresql-14-pg-liquid_0.1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-liquid` | `0.1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 367.0 KiB | [postgresql-14-pg-liquid_0.1.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-liquid/postgresql-14-pg-liquid_0.1.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-liquid` | `0.1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 364.5 KiB | [postgresql-14-pg-liquid_0.1.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-liquid/postgresql-14-pg-liquid_0.1.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-liquid` | `0.1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 427.7 KiB | [postgresql-14-pg-liquid_0.1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-liquid/postgresql-14-pg-liquid_0.1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-liquid` | `0.1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 429.9 KiB | [postgresql-14-pg-liquid_0.1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-liquid/postgresql-14-pg-liquid_0.1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-liquid` | `0.1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 387.5 KiB | [postgresql-14-pg-liquid_0.1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-liquid/postgresql-14-pg-liquid_0.1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-liquid` | `0.1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 388.5 KiB | [postgresql-14-pg-liquid_0.1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-liquid/postgresql-14-pg-liquid_0.1.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/michael-golfi/pg_liquid" title="Repository" icon="github" subtitle="github.com/michael-golfi/pg_liquid" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_liquid-0.1.7.zip" >}}
{{< /cards >}}


```bash
pig build pkg pg_liquid;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_liquid;		# install via package name, for the active PG version

pig install pg_liquid -v 18;   # install for PG 18
pig install pg_liquid -v 17;   # install for PG 17
pig install pg_liquid -v 16;   # install for PG 16
pig install pg_liquid -v 15;   # install for PG 15
pig install pg_liquid -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_liquid;
```


## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_liquid;
> SELECT liquid.query('Edge("a","b"). Edge("b","c"). Path(X,Y) :- Edge(X,Y). Path(X,Y) :- Edge(X,Z), Path(Z,Y). Path("a",Y)?');
> ```
>
> Sources: [README](https://github.com/michael-golfi/pg_liquid), [Docs site](https://michael-golfi.github.io/pg_liquid/)

`pg_liquid` maps the Liquid blog language and data model onto native PostgreSQL storage and execution. The extension exposes SQL entry points for running Liquid-style programs, querying as a principal, and managing row normalizers that project relational rows into Liquid compounds.

## Core Functions

The upstream README lists these main functions:

- `liquid.query(program text)`
- `liquid.query_as(principal text, program text)`
- `liquid.read_as(principal text, program text)`

These support plain execution, principal-aware querying, and CLS-aware reads.

## Language Features

The current README says supported program features include:

- `%` comments
- assertions and rule definitions terminated with `.`
- one terminal `?` query
- `Edge(...)`
- named compounds such as `Type@(cid=..., role=...)`
- query-local recursive rules

## Example Shape

Programs are passed as text and can define facts, rules, and a final query:

```sql
SELECT liquid.query($$
  Edge("a","b").
  Edge("b","c").
  Path(X,Y) :- Edge(X,Y).
  Path(X,Y) :- Edge(X,Z), Path(Z,Y).
  Path("a",Y)?
$$);
```

## Notes

The project README points to the VitePress documentation site as the main documentation surface and notes that operational rollout details are also documented there. The extension is currently published as PGXN package version `0.1.1` and validated against PostgreSQL 14 through 18.
