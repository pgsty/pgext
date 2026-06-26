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
| **2900** | {{< badge content="provsql" link="https://github.com/PierreSenellart/provsql" >}} | {{< ext "provsql" >}} | `1.10.0` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLdt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "uuid-ossp" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.10.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `provsql` | `uuid-ossp` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.10.0` | {{< bg "18" "provsql_18" "green" >}} {{< bg "17" "provsql_17" "green" >}} {{< bg "16" "provsql_16" "green" >}} {{< bg "15" "provsql_15" "green" >}} {{< bg "14" "provsql_14" "green" >}} | `provsql_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.10.0` | {{< bg "18" "postgresql-18-provsql" "green" >}} {{< bg "17" "postgresql-17-provsql" "green" >}} {{< bg "16" "postgresql-16-provsql" "green" >}} {{< bg "15" "postgresql-15-provsql" "green" >}} {{< bg "14" "postgresql-14-provsql" "green" >}} | `postgresql-$v-provsql` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.10.0" "provsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.10.0" "provsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.10.0" "provsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.10.0" "provsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.10.0" "provsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.10.0" "provsql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "provsql_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-14-provsql : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-18-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-17-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-16-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-15-provsql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.0" "postgresql-14-provsql : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `provsql_18` | `1.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [provsql_18-1.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/provsql_18-1.10.0-1PIGSTY.el8.x86_64.rpm) |
| `provsql_18` | `1.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1012.2 KiB | [provsql_18-1.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/provsql_18-1.10.0-1PIGSTY.el8.aarch64.rpm) |
| `provsql_18` | `1.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [provsql_18-1.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/provsql_18-1.10.0-1PIGSTY.el9.x86_64.rpm) |
| `provsql_18` | `1.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [provsql_18-1.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/provsql_18-1.10.0-1PIGSTY.el9.aarch64.rpm) |
| `provsql_18` | `1.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.2 MiB | [provsql_18-1.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/provsql_18-1.10.0-1PIGSTY.el10.x86_64.rpm) |
| `provsql_18` | `1.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.1 MiB | [provsql_18-1.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/provsql_18-1.10.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-provsql` | `1.10.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.0 MiB | [postgresql-18-provsql_1.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-18-provsql_1.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-provsql` | `1.10.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 907.0 KiB | [postgresql-18-provsql_1.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-18-provsql_1.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-provsql` | `1.10.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.1 MiB | [postgresql-18-provsql_1.10.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-18-provsql_1.10.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-provsql` | `1.10.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 985.7 KiB | [postgresql-18-provsql_1.10.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-18-provsql_1.10.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-provsql` | `1.10.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.0 MiB | [postgresql-18-provsql_1.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-18-provsql_1.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-provsql` | `1.10.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 997.4 KiB | [postgresql-18-provsql_1.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-18-provsql_1.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-provsql` | `1.10.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.1 MiB | [postgresql-18-provsql_1.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-18-provsql_1.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-provsql` | `1.10.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.0 MiB | [postgresql-18-provsql_1.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-18-provsql_1.10.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-provsql` | `1.10.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.1 MiB | [postgresql-18-provsql_1.10.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/provsql/postgresql-18-provsql_1.10.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-provsql` | `1.10.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.0 MiB | [postgresql-18-provsql_1.10.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/provsql/postgresql-18-provsql_1.10.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `provsql_17` | `1.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [provsql_17-1.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/provsql_17-1.10.0-1PIGSTY.el8.x86_64.rpm) |
| `provsql_17` | `1.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1011.7 KiB | [provsql_17-1.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/provsql_17-1.10.0-1PIGSTY.el8.aarch64.rpm) |
| `provsql_17` | `1.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [provsql_17-1.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/provsql_17-1.10.0-1PIGSTY.el9.x86_64.rpm) |
| `provsql_17` | `1.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [provsql_17-1.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/provsql_17-1.10.0-1PIGSTY.el9.aarch64.rpm) |
| `provsql_17` | `1.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.2 MiB | [provsql_17-1.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/provsql_17-1.10.0-1PIGSTY.el10.x86_64.rpm) |
| `provsql_17` | `1.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.1 MiB | [provsql_17-1.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/provsql_17-1.10.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-provsql` | `1.10.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.0 MiB | [postgresql-17-provsql_1.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-17-provsql_1.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-provsql` | `1.10.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 906.7 KiB | [postgresql-17-provsql_1.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-17-provsql_1.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-provsql` | `1.10.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.1 MiB | [postgresql-17-provsql_1.10.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-17-provsql_1.10.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-provsql` | `1.10.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 985.1 KiB | [postgresql-17-provsql_1.10.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-17-provsql_1.10.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-provsql` | `1.10.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.0 MiB | [postgresql-17-provsql_1.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-17-provsql_1.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-provsql` | `1.10.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 998.0 KiB | [postgresql-17-provsql_1.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-17-provsql_1.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-provsql` | `1.10.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.1 MiB | [postgresql-17-provsql_1.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-17-provsql_1.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-provsql` | `1.10.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.0 MiB | [postgresql-17-provsql_1.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-17-provsql_1.10.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-provsql` | `1.10.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.1 MiB | [postgresql-17-provsql_1.10.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/provsql/postgresql-17-provsql_1.10.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-provsql` | `1.10.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.0 MiB | [postgresql-17-provsql_1.10.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/provsql/postgresql-17-provsql_1.10.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `provsql_16` | `1.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [provsql_16-1.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/provsql_16-1.10.0-1PIGSTY.el8.x86_64.rpm) |
| `provsql_16` | `1.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1012.4 KiB | [provsql_16-1.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/provsql_16-1.10.0-1PIGSTY.el8.aarch64.rpm) |
| `provsql_16` | `1.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [provsql_16-1.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/provsql_16-1.10.0-1PIGSTY.el9.x86_64.rpm) |
| `provsql_16` | `1.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [provsql_16-1.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/provsql_16-1.10.0-1PIGSTY.el9.aarch64.rpm) |
| `provsql_16` | `1.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.2 MiB | [provsql_16-1.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/provsql_16-1.10.0-1PIGSTY.el10.x86_64.rpm) |
| `provsql_16` | `1.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.1 MiB | [provsql_16-1.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/provsql_16-1.10.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-provsql` | `1.10.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.0 MiB | [postgresql-16-provsql_1.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-16-provsql_1.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-provsql` | `1.10.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 908.0 KiB | [postgresql-16-provsql_1.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-16-provsql_1.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-provsql` | `1.10.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.1 MiB | [postgresql-16-provsql_1.10.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-16-provsql_1.10.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-provsql` | `1.10.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 985.9 KiB | [postgresql-16-provsql_1.10.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-16-provsql_1.10.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-provsql` | `1.10.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.0 MiB | [postgresql-16-provsql_1.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-16-provsql_1.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-provsql` | `1.10.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 992.0 KiB | [postgresql-16-provsql_1.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-16-provsql_1.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-provsql` | `1.10.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.1 MiB | [postgresql-16-provsql_1.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-16-provsql_1.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-provsql` | `1.10.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.0 MiB | [postgresql-16-provsql_1.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-16-provsql_1.10.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-provsql` | `1.10.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.1 MiB | [postgresql-16-provsql_1.10.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/provsql/postgresql-16-provsql_1.10.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-provsql` | `1.10.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.0 MiB | [postgresql-16-provsql_1.10.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/provsql/postgresql-16-provsql_1.10.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `provsql_15` | `1.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.2 MiB | [provsql_15-1.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/provsql_15-1.10.0-1PIGSTY.el8.x86_64.rpm) |
| `provsql_15` | `1.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.0 MiB | [provsql_15-1.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/provsql_15-1.10.0-1PIGSTY.el8.aarch64.rpm) |
| `provsql_15` | `1.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [provsql_15-1.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/provsql_15-1.10.0-1PIGSTY.el9.x86_64.rpm) |
| `provsql_15` | `1.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [provsql_15-1.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/provsql_15-1.10.0-1PIGSTY.el9.aarch64.rpm) |
| `provsql_15` | `1.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.2 MiB | [provsql_15-1.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/provsql_15-1.10.0-1PIGSTY.el10.x86_64.rpm) |
| `provsql_15` | `1.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.1 MiB | [provsql_15-1.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/provsql_15-1.10.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-provsql` | `1.10.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-15-provsql_1.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-15-provsql_1.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-provsql` | `1.10.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 958.2 KiB | [postgresql-15-provsql_1.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-15-provsql_1.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-provsql` | `1.10.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-15-provsql_1.10.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-15-provsql_1.10.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-provsql` | `1.10.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.0 MiB | [postgresql-15-provsql_1.10.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-15-provsql_1.10.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-provsql` | `1.10.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.1 MiB | [postgresql-15-provsql_1.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-15-provsql_1.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-provsql` | `1.10.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.0 MiB | [postgresql-15-provsql_1.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-15-provsql_1.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-provsql` | `1.10.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.1 MiB | [postgresql-15-provsql_1.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-15-provsql_1.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-provsql` | `1.10.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-15-provsql_1.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-15-provsql_1.10.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-provsql` | `1.10.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.2 MiB | [postgresql-15-provsql_1.10.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/provsql/postgresql-15-provsql_1.10.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-provsql` | `1.10.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.1 MiB | [postgresql-15-provsql_1.10.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/provsql/postgresql-15-provsql_1.10.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `provsql_14` | `1.10.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.2 MiB | [provsql_14-1.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/provsql_14-1.10.0-1PIGSTY.el8.x86_64.rpm) |
| `provsql_14` | `1.10.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.0 MiB | [provsql_14-1.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/provsql_14-1.10.0-1PIGSTY.el8.aarch64.rpm) |
| `provsql_14` | `1.10.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [provsql_14-1.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/provsql_14-1.10.0-1PIGSTY.el9.x86_64.rpm) |
| `provsql_14` | `1.10.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [provsql_14-1.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/provsql_14-1.10.0-1PIGSTY.el9.aarch64.rpm) |
| `provsql_14` | `1.10.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.2 MiB | [provsql_14-1.10.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/provsql_14-1.10.0-1PIGSTY.el10.x86_64.rpm) |
| `provsql_14` | `1.10.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.1 MiB | [provsql_14-1.10.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/provsql_14-1.10.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-provsql` | `1.10.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-14-provsql_1.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-14-provsql_1.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-provsql` | `1.10.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 952.0 KiB | [postgresql-14-provsql_1.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/provsql/postgresql-14-provsql_1.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-provsql` | `1.10.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-14-provsql_1.10.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-14-provsql_1.10.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-provsql` | `1.10.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.0 MiB | [postgresql-14-provsql_1.10.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/provsql/postgresql-14-provsql_1.10.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-provsql` | `1.10.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.1 MiB | [postgresql-14-provsql_1.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-14-provsql_1.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-provsql` | `1.10.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.0 MiB | [postgresql-14-provsql_1.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/provsql/postgresql-14-provsql_1.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-provsql` | `1.10.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.1 MiB | [postgresql-14-provsql_1.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-14-provsql_1.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-provsql` | `1.10.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-14-provsql_1.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/provsql/postgresql-14-provsql_1.10.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-provsql` | `1.10.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.2 MiB | [postgresql-14-provsql_1.10.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/provsql/postgresql-14-provsql_1.10.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-provsql` | `1.10.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.1 MiB | [postgresql-14-provsql_1.10.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/provsql/postgresql-14-provsql_1.10.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/PierreSenellart/provsql" title="Repository" icon="github" subtitle="github.com/PierreSenellart/provsql" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="provsql-1.10.0.tar.gz" >}}
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

Sources: [README](https://github.com/PierreSenellart/provsql/blob/v1.9.0/doc/provsql.md), [v1.9.0 release](https://github.com/PierreSenellart/provsql/releases/tag/v1.9.0), [v1.9.0 control](https://github.com/PierreSenellart/provsql/blob/v1.9.0/provsql.common.control), [getting started](https://provsql.org/docs/user/getting-provsql.html), [configuration](https://provsql.org/docs/user/configuration.html), [semirings](https://provsql.org/docs/user/semirings.html)

`provsql` adds semiring provenance and uncertainty management to PostgreSQL. Upstream documents provenance tracking, semiring evaluation, probabilities, Shapley and Banzhaf values, where-provenance, update provenance, and temporal features.

### Load and Track Provenance

```ini
shared_preload_libraries = 'provsql'
```

```sql
CREATE EXTENSION provsql CASCADE;
```

The `CASCADE` form installs `uuid-ossp` automatically if needed. The getting-started guide says the preload step is mandatory because ProvSQL installs a planner hook.

```sql
SELECT provsql.add_provenance('mytable');

SELECT name, provenance()
FROM mytable;

SELECT provsql.remove_provenance('mytable');
```

The user docs also describe provenance mappings:

```sql
SELECT create_provenance_mapping('my_mapping', 'mytable', 'column_name');
SELECT create_provenance_mapping_view('my_mapping_view', 'mytable', 'column_name');
```

### Probability and Influence

Assign probabilities to tuple tokens:

```sql
SELECT set_prob(provenance(), 0.8)
FROM mytable
WHERE id = 1;

SELECT name, probability_evaluate(provenance()) AS prob
FROM mytable;
```

Compute influence scores:

```sql
SELECT shapley(provenance(), m.token)
FROM mytable, my_mapping AS m;

SELECT banzhaf(provenance(), m.token)
FROM mytable, my_mapping AS m;
```

The docs also describe `shapley_all_vars` and `banzhaf_all_vars` for computing scores for all input variables at once.

### Built-in Semirings

Built-in semiring functions use a provenance token and a provenance mapping table:

```sql
SELECT name, sr_boolean(provenance(), 'my_mapping')
FROM mytable;

SELECT name, sr_formula(provenance(), 'my_mapping')
FROM mytable;

SELECT name, sr_how(provenance(), 'my_mapping')
FROM mytable;
```

Current docs include compiled wrappers for `sr_how`, `sr_which`, `sr_tropical`, `sr_viterbi`, `sr_lukasiewicz`, `sr_minmax`, and `sr_maxmin`. For PostgreSQL 14 and later they also include `sr_temporal`, `sr_interval_num`, and `sr_interval_int` over multirange values.

```sql
SELECT city,
       sr_minmax(provenance(), 'personnel_level',
                 'unclassified'::classification_level) AS clearance
FROM (SELECT DISTINCT city FROM personnel) AS t;

SELECT entity_id, sr_temporal(provenance(), 'validity_mapping')
FROM mytable;
```

Advanced users can still define custom semirings and evaluate them with `provenance_evaluate` or `aggregation_evaluate`; upstream recommends the compiled semirings when one matches the needed algebra.

### Extra Modes and Helpers

Session GUCs documented upstream include:

```sql
SET provsql.active = on;
SET provsql.where_provenance = on;
SET provsql.update_provenance = on;
SET provsql.last_eval_method = on;
SET provsql.tool_search_path = '/opt/d4:/home/postgres/bin';
SET provsql.aggtoken_text_as_uuid = on;
```

`provsql.tool_search_path` is used for external probability and visualization tools such as `d4`, `c2d`, `dsharp`, `minic2d`, `weightmc`, and `graph-easy`. `provsql.last_eval_method` stores the last chosen probability-evaluation method. `provsql.aggtoken_text_as_uuid` makes aggregate-token cells render as their provenance UUIDs; `agg_token_value_text(token)` can recover the display text for those aggregate tokens.

The user guide separately documents where-provenance helpers, update provenance, temporal helpers such as `get_valid_time`, `timetravel`, `timeslice`, `history`, and `undo`, circuit-inspection helpers `circuit_subgraph(root, max_depth)` and `resolve_input(uuid)`, and `setup_search_path()` for preparing the helper search path.

### v1.9.0 Query and Probability Notes

Release `1.9.0` materially expands SQL coverage for provenance-aware queries:

- subqueries outside `FROM`, including `EXISTS`, `NOT EXISTS`, `IN`, `NOT IN`, `ANY`, `ALL`, row-valued `IN`, scalar subqueries, and `ARRAY(SELECT ...)`;
- `LEFT`, `RIGHT`, and `FULL` outer joins, plus corrected `EXCEPT` and `EXCEPT ALL` provenance;
- SQL-faithful `NULL` handling for aggregates and exact `HAVING` aggregate probabilities for `COUNT`, `SUM`, `MIN`, `MAX`, and `AVG`;
- probability-method selection through the method catalog and cost chooser, with `karp-luby`, `stopping-rule`, `sieve`, `d-tree`, and `probability_bounds`;
- idempotent `add_provenance` and `create_provenance_mapping` calls.

The release removes the old `probability_benchmark` helper. `agg_token` now has native arithmetic, unary minus, and comparison support for aggregate-token expressions.

### Notes

- The package row in `db/extension.csv` lists version `1.9.0`, package `provsql`, dependency `uuid-ossp`, and PostgreSQL support for 14 through 18.
- The v1.9.0 control file sets `default_version = '1.9.0'`, requires `uuid-ossp`, marks the extension trusted, and is not relocatable.
- Upstream docs say ProvSQL has been tested on PostgreSQL 10 through 18; the Pigsty package matrix is PostgreSQL 14-18.
- `provsql.update_provenance` and the multirange semirings require PostgreSQL 14 or later.
