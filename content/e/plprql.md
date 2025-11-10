---
title: "plprql"
linkTitle: "plprql"
description: "Use PRQL in PostgreSQL - Pipelined Relational Query Language"
weight: 3040
categories: ["LANG"]
width: full
---

[**plprql**](https://github.com/kaspermarstal/plprql) : Use PRQL in PostgreSQL - Pipelined Relational Query Language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3040** | {{< badge content="plprql" link="https://github.com/kaspermarstal/plprql" >}} | {{< ext "plprql" >}} | `18.0.0` | {{< category "LANG" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_tle" >}} {{< ext "plpgsql" >}} {{< ext "plv8" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "pllua" >}} {{< ext "hstore_pllua" >}} {{< ext "plluau" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `18.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `plprql` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `18.0.0` | {{< bg "18" "plprql_18" "green" >}} {{< bg "17" "plprql_17" "green" >}} {{< bg "16" "plprql_16" "green" >}} {{< bg "15" "plprql_15" "green" >}} {{< bg "14" "plprql_14" "green" >}} {{< bg "13" "plprql_13" "green" >}} | `plprql_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `18.0.0` | {{< bg "18" "postgresql-18-plprql" "green" >}} {{< bg "17" "postgresql-17-plprql" "green" >}} {{< bg "16" "postgresql-16-plprql" "green" >}} {{< bg "15" "postgresql-15-plprql" "green" >}} {{< bg "14" "postgresql-14-plprql" "green" >}} {{< bg "13" "postgresql-13-plprql" "green" >}} | `postgresql-$v-plprql` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 18.0.0" "plprql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 18.0.0" "plprql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 18.0.0" "plprql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 18.0.0" "plprql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 18.0.0" "plprql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 18.0.0" "plprql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "plprql_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-18-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-17-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-18-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-17-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-18-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-17-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-18-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-17-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-18-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-17-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-18-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-17-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-18-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-17-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-18-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-17-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 18.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprql_18` | `18.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [plprql_18-18.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plprql_18-18.0.0-1PIGSTY.el8.x86_64.rpm) |
| `plprql_18` | `18.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.4 MiB | [plprql_18-18.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plprql_18-18.0.0-1PIGSTY.el8.aarch64.rpm) |
| `plprql_18` | `18.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [plprql_18-18.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plprql_18-18.0.0-1PIGSTY.el9.x86_64.rpm) |
| `plprql_18` | `18.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.5 MiB | [plprql_18-18.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plprql_18-18.0.0-1PIGSTY.el9.aarch64.rpm) |
| `plprql_18` | `18.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.7 MiB | [plprql_18-18.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plprql_18-18.0.0-1PIGSTY.el10.x86_64.rpm) |
| `plprql_18` | `18.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.5 MiB | [plprql_18-18.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plprql_18-18.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-plprql` | `18.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.2 KiB | [postgresql-18-plprql_18.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-18-plprql_18.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-plprql` | `18.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 KiB | [postgresql-18-plprql_18.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-18-plprql_18.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-plprql` | `18.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.2 KiB | [postgresql-18-plprql_18.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plprql/postgresql-18-plprql_18.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-plprql` | `18.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 KiB | [postgresql-18-plprql_18.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plprql/postgresql-18-plprql_18.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-plprql` | `18.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.2 KiB | [postgresql-18-plprql_18.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-18-plprql_18.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-plprql` | `18.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.2 KiB | [postgresql-18-plprql_18.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-18-plprql_18.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-plprql` | `18.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.2 KiB | [postgresql-18-plprql_18.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-18-plprql_18.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-plprql` | `18.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.2 KiB | [postgresql-18-plprql_18.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-18-plprql_18.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprql_17` | `18.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [plprql_17-18.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plprql_17-18.0.0-1PIGSTY.el8.x86_64.rpm) |
| `plprql_17` | `18.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.4 MiB | [plprql_17-18.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plprql_17-18.0.0-1PIGSTY.el8.aarch64.rpm) |
| `plprql_17` | `18.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [plprql_17-18.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plprql_17-18.0.0-1PIGSTY.el9.x86_64.rpm) |
| `plprql_17` | `18.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.5 MiB | [plprql_17-18.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plprql_17-18.0.0-1PIGSTY.el9.aarch64.rpm) |
| `plprql_17` | `18.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.7 MiB | [plprql_17-18.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plprql_17-18.0.0-1PIGSTY.el10.x86_64.rpm) |
| `plprql_17` | `18.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.5 MiB | [plprql_17-18.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plprql_17-18.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-plprql` | `18.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.2 KiB | [postgresql-17-plprql_18.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-17-plprql_18.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-plprql` | `18.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 KiB | [postgresql-17-plprql_18.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-17-plprql_18.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-plprql` | `18.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.2 KiB | [postgresql-17-plprql_18.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plprql/postgresql-17-plprql_18.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-plprql` | `18.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 KiB | [postgresql-17-plprql_18.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plprql/postgresql-17-plprql_18.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-plprql` | `18.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.2 KiB | [postgresql-17-plprql_18.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-17-plprql_18.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-plprql` | `18.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.2 KiB | [postgresql-17-plprql_18.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-17-plprql_18.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-plprql` | `18.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.2 KiB | [postgresql-17-plprql_18.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-17-plprql_18.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-plprql` | `18.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.2 KiB | [postgresql-17-plprql_18.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-17-plprql_18.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprql_16` | `18.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [plprql_16-18.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plprql_16-18.0.0-1PIGSTY.el8.x86_64.rpm) |
| `plprql_16` | `18.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.4 MiB | [plprql_16-18.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plprql_16-18.0.0-1PIGSTY.el8.aarch64.rpm) |
| `plprql_16` | `18.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [plprql_16-18.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plprql_16-18.0.0-1PIGSTY.el9.x86_64.rpm) |
| `plprql_16` | `18.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.5 MiB | [plprql_16-18.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plprql_16-18.0.0-1PIGSTY.el9.aarch64.rpm) |
| `plprql_16` | `18.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.7 MiB | [plprql_16-18.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plprql_16-18.0.0-1PIGSTY.el10.x86_64.rpm) |
| `plprql_16` | `18.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.5 MiB | [plprql_16-18.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plprql_16-18.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-plprql` | `18.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.3 MiB | [postgresql-16-plprql_18.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-16-plprql_18.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-plprql` | `18.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.9 MiB | [postgresql-16-plprql_18.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-16-plprql_18.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-plprql` | `18.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.3 MiB | [postgresql-16-plprql_18.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plprql/postgresql-16-plprql_18.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-plprql` | `18.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.9 MiB | [postgresql-16-plprql_18.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plprql/postgresql-16-plprql_18.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-plprql` | `18.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.5 MiB | [postgresql-16-plprql_18.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-16-plprql_18.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-plprql` | `18.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.3 MiB | [postgresql-16-plprql_18.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-16-plprql_18.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-plprql` | `18.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.5 MiB | [postgresql-16-plprql_18.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-16-plprql_18.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-plprql` | `18.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.3 MiB | [postgresql-16-plprql_18.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-16-plprql_18.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprql_15` | `18.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [plprql_15-18.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plprql_15-18.0.0-1PIGSTY.el8.x86_64.rpm) |
| `plprql_15` | `18.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.4 MiB | [plprql_15-18.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plprql_15-18.0.0-1PIGSTY.el8.aarch64.rpm) |
| `plprql_15` | `18.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [plprql_15-18.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plprql_15-18.0.0-1PIGSTY.el9.x86_64.rpm) |
| `plprql_15` | `18.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.5 MiB | [plprql_15-18.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plprql_15-18.0.0-1PIGSTY.el9.aarch64.rpm) |
| `plprql_15` | `18.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.7 MiB | [plprql_15-18.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plprql_15-18.0.0-1PIGSTY.el10.x86_64.rpm) |
| `plprql_15` | `18.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.5 MiB | [plprql_15-18.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plprql_15-18.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-plprql` | `18.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.3 MiB | [postgresql-15-plprql_18.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-15-plprql_18.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-plprql` | `18.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.9 MiB | [postgresql-15-plprql_18.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-15-plprql_18.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-plprql` | `18.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.3 MiB | [postgresql-15-plprql_18.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plprql/postgresql-15-plprql_18.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-plprql` | `18.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.9 MiB | [postgresql-15-plprql_18.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plprql/postgresql-15-plprql_18.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-plprql` | `18.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.5 MiB | [postgresql-15-plprql_18.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-15-plprql_18.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-plprql` | `18.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.3 MiB | [postgresql-15-plprql_18.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-15-plprql_18.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-plprql` | `18.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.5 MiB | [postgresql-15-plprql_18.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-15-plprql_18.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-plprql` | `18.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.3 MiB | [postgresql-15-plprql_18.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-15-plprql_18.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprql_14` | `18.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [plprql_14-18.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plprql_14-18.0.0-1PIGSTY.el8.x86_64.rpm) |
| `plprql_14` | `18.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.4 MiB | [plprql_14-18.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plprql_14-18.0.0-1PIGSTY.el8.aarch64.rpm) |
| `plprql_14` | `18.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [plprql_14-18.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plprql_14-18.0.0-1PIGSTY.el9.x86_64.rpm) |
| `plprql_14` | `18.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.5 MiB | [plprql_14-18.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plprql_14-18.0.0-1PIGSTY.el9.aarch64.rpm) |
| `plprql_14` | `18.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.7 MiB | [plprql_14-18.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plprql_14-18.0.0-1PIGSTY.el10.x86_64.rpm) |
| `plprql_14` | `18.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.5 MiB | [plprql_14-18.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plprql_14-18.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-plprql` | `18.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.3 MiB | [postgresql-14-plprql_18.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-14-plprql_18.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-plprql` | `18.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.9 MiB | [postgresql-14-plprql_18.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-14-plprql_18.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-plprql` | `18.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.3 MiB | [postgresql-14-plprql_18.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plprql/postgresql-14-plprql_18.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-plprql` | `18.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.9 MiB | [postgresql-14-plprql_18.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plprql/postgresql-14-plprql_18.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-plprql` | `18.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.5 MiB | [postgresql-14-plprql_18.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-14-plprql_18.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-plprql` | `18.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.3 MiB | [postgresql-14-plprql_18.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-14-plprql_18.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-plprql` | `18.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.5 MiB | [postgresql-14-plprql_18.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-14-plprql_18.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-plprql` | `18.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.3 MiB | [postgresql-14-plprql_18.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-14-plprql_18.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprql_13` | `18.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [plprql_13-18.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plprql_13-18.0.0-1PIGSTY.el8.x86_64.rpm) |
| `plprql_13` | `18.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.4 MiB | [plprql_13-18.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plprql_13-18.0.0-1PIGSTY.el8.aarch64.rpm) |
| `plprql_13` | `18.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [plprql_13-18.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plprql_13-18.0.0-1PIGSTY.el9.x86_64.rpm) |
| `plprql_13` | `18.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.5 MiB | [plprql_13-18.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plprql_13-18.0.0-1PIGSTY.el9.aarch64.rpm) |
| `plprql_13` | `18.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.7 MiB | [plprql_13-18.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plprql_13-18.0.0-1PIGSTY.el10.x86_64.rpm) |
| `plprql_13` | `18.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.5 MiB | [plprql_13-18.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plprql_13-18.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-plprql` | `18.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.3 MiB | [postgresql-13-plprql_18.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-13-plprql_18.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-plprql` | `18.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.9 MiB | [postgresql-13-plprql_18.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-13-plprql_18.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-plprql` | `18.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.3 MiB | [postgresql-13-plprql_18.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plprql/postgresql-13-plprql_18.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-plprql` | `18.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.9 MiB | [postgresql-13-plprql_18.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plprql/postgresql-13-plprql_18.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-plprql` | `18.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.5 MiB | [postgresql-13-plprql_18.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-13-plprql_18.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-plprql` | `18.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.3 MiB | [postgresql-13-plprql_18.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-13-plprql_18.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-plprql` | `18.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.5 MiB | [postgresql-13-plprql_18.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-13-plprql_18.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-plprql` | `18.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.3 MiB | [postgresql-13-plprql_18.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-13-plprql_18.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/kaspermarstal/plprql" title="Repository" icon="github" subtitle="github.com/kaspermarstal/plprql" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plprql-18.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg plprql;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install plprql;		# install via package name, for the active PG version

pig install plprql -v 18;   # install for PG 18
pig install plprql -v 17;   # install for PG 17
pig install plprql -v 16;   # install for PG 16
pig install plprql -v 15;   # install for PG 15
pig install plprql -v 14;   # install for PG 14
pig install plprql -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plprql;
```
