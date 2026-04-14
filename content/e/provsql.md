---
title: "provsql"
linkTitle: "provsql"
description: "Semiring provenance and uncertainty management for PostgreSQL"
weight: 2900
categories: ["FEAT"]
width: full
---

[**provsql**](https://github.com/PierreSenellart/provsql) : Semiring provenance and uncertainty management for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2900** | {{< badge content="provsql" link="https://github.com/PierreSenellart/provsql" >}} | {{< ext "provsql" >}} | `1.2.3` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLdt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "uuid-ossp" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `provsql` | `uuid-ossp` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.3` | {{< bg "18" "provsql_18" "green" >}} {{< bg "17" "provsql_17" "green" >}} {{< bg "16" "provsql_16" "green" >}} {{< bg "15" "provsql_15" "green" >}} {{< bg "14" "provsql_14" "green" >}} | `provsql_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.3` | {{< bg "18" "postgresql-18-provsql" "green" >}} {{< bg "17" "postgresql-17-provsql" "green" >}} {{< bg "16" "postgresql-16-provsql" "green" >}} {{< bg "15" "postgresql-15-provsql" "green" >}} {{< bg "14" "postgresql-14-provsql" "green" >}} | `postgresql-$v-provsql` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "provsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "provsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "provsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "provsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "provsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "provsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "provsql_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-provsql : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `provsql_18` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 313.0 KiB | [provsql_18-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/provsql_18-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `provsql_18` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 289.8 KiB | [provsql_18-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/provsql_18-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `provsql_18` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 321.4 KiB | [provsql_18-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/provsql_18-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `provsql_18` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 307.1 KiB | [provsql_18-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/provsql_18-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `provsql_18` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 323.8 KiB | [provsql_18-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/provsql_18-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `provsql_18` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 296.5 KiB | [provsql_18-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/provsql_18-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-provsql` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 271.0 KiB | [postgresql-18-provsql_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-18-provsql_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-provsql` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 239.3 KiB | [postgresql-18-provsql_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-18-provsql_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-provsql` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 295.5 KiB | [postgresql-18-provsql_1.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-18-provsql_1.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-provsql` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 257.7 KiB | [postgresql-18-provsql_1.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-18-provsql_1.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-provsql` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 285.5 KiB | [postgresql-18-provsql_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-18-provsql_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-provsql` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 268.6 KiB | [postgresql-18-provsql_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-18-provsql_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-provsql` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 297.9 KiB | [postgresql-18-provsql_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-18-provsql_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-provsql` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 279.5 KiB | [postgresql-18-provsql_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-18-provsql_1.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `provsql_17` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 312.8 KiB | [provsql_17-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/provsql_17-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `provsql_17` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 289.7 KiB | [provsql_17-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/provsql_17-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `provsql_17` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 321.1 KiB | [provsql_17-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/provsql_17-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `provsql_17` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 306.9 KiB | [provsql_17-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/provsql_17-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `provsql_17` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 323.5 KiB | [provsql_17-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/provsql_17-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `provsql_17` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 296.3 KiB | [provsql_17-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/provsql_17-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-provsql` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 270.5 KiB | [postgresql-17-provsql_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-17-provsql_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-provsql` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 239.2 KiB | [postgresql-17-provsql_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-17-provsql_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-provsql` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 294.9 KiB | [postgresql-17-provsql_1.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-17-provsql_1.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-provsql` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 257.4 KiB | [postgresql-17-provsql_1.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-17-provsql_1.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-provsql` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 285.4 KiB | [postgresql-17-provsql_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-17-provsql_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-provsql` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 268.4 KiB | [postgresql-17-provsql_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-17-provsql_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-provsql` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 298.5 KiB | [postgresql-17-provsql_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-17-provsql_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-provsql` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 279.4 KiB | [postgresql-17-provsql_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-17-provsql_1.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `provsql_16` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 312.9 KiB | [provsql_16-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/provsql_16-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `provsql_16` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 290.0 KiB | [provsql_16-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/provsql_16-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `provsql_16` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 321.6 KiB | [provsql_16-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/provsql_16-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `provsql_16` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 307.4 KiB | [provsql_16-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/provsql_16-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `provsql_16` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 323.5 KiB | [provsql_16-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/provsql_16-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `provsql_16` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 296.4 KiB | [provsql_16-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/provsql_16-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-provsql` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 270.7 KiB | [postgresql-16-provsql_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-16-provsql_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-provsql` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 239.2 KiB | [postgresql-16-provsql_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-16-provsql_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-provsql` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 295.1 KiB | [postgresql-16-provsql_1.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-16-provsql_1.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-provsql` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 257.4 KiB | [postgresql-16-provsql_1.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-16-provsql_1.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-provsql` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 286.4 KiB | [postgresql-16-provsql_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-16-provsql_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-provsql` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 268.7 KiB | [postgresql-16-provsql_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-16-provsql_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-provsql` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 300.0 KiB | [postgresql-16-provsql_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-16-provsql_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-provsql` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 280.1 KiB | [postgresql-16-provsql_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-16-provsql_1.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `provsql_15` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 340.2 KiB | [provsql_15-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/provsql_15-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `provsql_15` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 317.9 KiB | [provsql_15-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/provsql_15-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `provsql_15` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 347.3 KiB | [provsql_15-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/provsql_15-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `provsql_15` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 333.8 KiB | [provsql_15-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/provsql_15-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `provsql_15` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 347.5 KiB | [provsql_15-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/provsql_15-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `provsql_15` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 320.4 KiB | [provsql_15-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/provsql_15-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-provsql` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 302.5 KiB | [postgresql-15-provsql_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-15-provsql_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-provsql` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 271.1 KiB | [postgresql-15-provsql_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-15-provsql_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-provsql` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 325.1 KiB | [postgresql-15-provsql_1.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-15-provsql_1.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-provsql` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 288.1 KiB | [postgresql-15-provsql_1.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-15-provsql_1.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-provsql` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 319.1 KiB | [postgresql-15-provsql_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-15-provsql_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-provsql` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 302.6 KiB | [postgresql-15-provsql_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-15-provsql_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-provsql` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 331.6 KiB | [postgresql-15-provsql_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-15-provsql_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-provsql` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 314.0 KiB | [postgresql-15-provsql_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-15-provsql_1.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `provsql_14` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 339.9 KiB | [provsql_14-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/provsql_14-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `provsql_14` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 317.7 KiB | [provsql_14-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/provsql_14-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `provsql_14` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 347.1 KiB | [provsql_14-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/provsql_14-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `provsql_14` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 333.4 KiB | [provsql_14-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/provsql_14-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `provsql_14` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 347.3 KiB | [provsql_14-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/provsql_14-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `provsql_14` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 319.9 KiB | [provsql_14-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/provsql_14-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-provsql` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 302.3 KiB | [postgresql-14-provsql_1.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-14-provsql_1.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-provsql` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 270.9 KiB | [postgresql-14-provsql_1.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-14-provsql_1.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-provsql` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 324.4 KiB | [postgresql-14-provsql_1.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-14-provsql_1.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-provsql` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 287.1 KiB | [postgresql-14-provsql_1.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-14-provsql_1.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-provsql` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 318.9 KiB | [postgresql-14-provsql_1.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-14-provsql_1.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-provsql` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 303.2 KiB | [postgresql-14-provsql_1.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-14-provsql_1.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-provsql` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 330.3 KiB | [postgresql-14-provsql_1.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-14-provsql_1.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-provsql` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 312.5 KiB | [postgresql-14-provsql_1.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-14-provsql_1.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/PierreSenellart/provsql" title="Repository" icon="github" subtitle="github.com/PierreSenellart/provsql" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="provsql-1.2.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg provsql;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install provsql;		# install via package name, for the active PG version

pig install provsql -v 18;   # install for PG 18
pig install provsql -v 17;   # install for PG 17
pig install provsql -v 16;   # install for PG 16
pig install provsql -v 15;   # install for PG 15
pig install provsql -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'provsql';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION provsql CASCADE; -- requires uuid-ossp
```

## Usage

- Source: [GitHub repo](https://github.com/PierreSenellart/provsql), [project docs](https://provsql.org/docs/), [getting started guide](https://provsql.org/docs/user/getting-provsql.html)
- ProvSQL adds m-semiring provenance and uncertainty management to PostgreSQL, with support for probabilities, Shapley values, and semiring evaluation.

```sql
shared_preload_libraries = 'provsql'
CREATE EXTENSION provsql CASCADE;
```

The upstream quick install also calls out these prerequisites: PostgreSQL 10 or later, a C++17 compiler, PostgreSQL headers, `uuid-ossp`, and Boost libraries.

## Core Workflow

ProvSQL is loaded through `shared_preload_libraries`, then installed with `CREATE EXTENSION provsql CASCADE;`.

Typical use cases include:

- evaluating provenance over different semirings
- computing probabilities and expected values
- computing game-theoretic contributions such as Shapley values
- using the built-in compiled semirings for common cases

## Notes

The project homepage and documentation live at [provsql.org](https://provsql.org/). The README links to the user guide for the full installation and testing workflow.
