---
title: "datasketches"
linkTitle: "datasketches"
description: "Approximate analytics sketches and aggregates for PostgreSQL"
weight: 4690
categories: ["FUNC"]
width: full
---

[**datasketches**](https://github.com/apache/datasketches-postgresql) : Approximate analytics sketches and aggregates for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4690** | {{< badge content="datasketches" link="https://github.com/apache/datasketches-postgresql" >}} | {{< ext "datasketches" >}} | `1.7.0` | {{< category "FUNC" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] Built against Apache DataSketches C++ core 5.0.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `datasketches` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "18" "datasketches_18" "green" >}} {{< bg "17" "datasketches_17" "green" >}} {{< bg "16" "datasketches_16" "green" >}} {{< bg "15" "datasketches_15" "green" >}} {{< bg "14" "datasketches_14" "green" >}} | `datasketches_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "18" "postgresql-18-datasketches" "green" >}} {{< bg "17" "postgresql-17-datasketches" "green" >}} {{< bg "16" "postgresql-16-datasketches" "green" >}} {{< bg "15" "postgresql-15-datasketches" "green" >}} {{< bg "14" "postgresql-14-datasketches" "green" >}} | `postgresql-$v-datasketches` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-datasketches : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-datasketches : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-datasketches : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-datasketches : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-datasketches : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-datasketches : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-datasketches : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-datasketches : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `datasketches_18` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 324.4 KiB | [datasketches_18-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/datasketches_18-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `datasketches_18` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 314.1 KiB | [datasketches_18-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/datasketches_18-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `datasketches_18` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 309.4 KiB | [datasketches_18-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/datasketches_18-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `datasketches_18` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 315.1 KiB | [datasketches_18-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/datasketches_18-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `datasketches_18` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 319.1 KiB | [datasketches_18-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/datasketches_18-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `datasketches_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 319.4 KiB | [datasketches_18-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/datasketches_18-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-datasketches` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 918.1 KiB | [postgresql-18-datasketches_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 920.0 KiB | [postgresql-18-datasketches_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 943.3 KiB | [postgresql-18-datasketches_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 944.0 KiB | [postgresql-18-datasketches_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1017.0 KiB | [postgresql-18-datasketches_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1020.8 KiB | [postgresql-18-datasketches_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 977.8 KiB | [postgresql-18-datasketches_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 991.3 KiB | [postgresql-18-datasketches_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `datasketches_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 324.4 KiB | [datasketches_17-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/datasketches_17-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `datasketches_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 314.1 KiB | [datasketches_17-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/datasketches_17-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `datasketches_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 309.4 KiB | [datasketches_17-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/datasketches_17-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `datasketches_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 315.0 KiB | [datasketches_17-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/datasketches_17-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `datasketches_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 319.1 KiB | [datasketches_17-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/datasketches_17-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `datasketches_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 319.4 KiB | [datasketches_17-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/datasketches_17-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-datasketches` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 918.3 KiB | [postgresql-17-datasketches_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 919.2 KiB | [postgresql-17-datasketches_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 942.9 KiB | [postgresql-17-datasketches_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 943.8 KiB | [postgresql-17-datasketches_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.1 MiB | [postgresql-17-datasketches_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-17-datasketches_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 977.8 KiB | [postgresql-17-datasketches_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 991.2 KiB | [postgresql-17-datasketches_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `datasketches_16` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 324.4 KiB | [datasketches_16-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/datasketches_16-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `datasketches_16` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 314.1 KiB | [datasketches_16-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/datasketches_16-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `datasketches_16` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 309.4 KiB | [datasketches_16-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/datasketches_16-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `datasketches_16` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 315.0 KiB | [datasketches_16-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/datasketches_16-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `datasketches_16` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 319.1 KiB | [datasketches_16-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/datasketches_16-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `datasketches_16` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 319.3 KiB | [datasketches_16-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/datasketches_16-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-datasketches` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 918.1 KiB | [postgresql-16-datasketches_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 919.5 KiB | [postgresql-16-datasketches_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 943.1 KiB | [postgresql-16-datasketches_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 943.8 KiB | [postgresql-16-datasketches_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.1 MiB | [postgresql-16-datasketches_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-16-datasketches_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 977.8 KiB | [postgresql-16-datasketches_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 991.2 KiB | [postgresql-16-datasketches_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `datasketches_15` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 342.1 KiB | [datasketches_15-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/datasketches_15-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `datasketches_15` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 332.3 KiB | [datasketches_15-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/datasketches_15-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `datasketches_15` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 323.5 KiB | [datasketches_15-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/datasketches_15-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `datasketches_15` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 329.1 KiB | [datasketches_15-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/datasketches_15-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `datasketches_15` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 325.9 KiB | [datasketches_15-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/datasketches_15-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `datasketches_15` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 325.2 KiB | [datasketches_15-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/datasketches_15-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-datasketches` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 932.6 KiB | [postgresql-15-datasketches_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 933.7 KiB | [postgresql-15-datasketches_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 957.8 KiB | [postgresql-15-datasketches_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 957.9 KiB | [postgresql-15-datasketches_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.1 MiB | [postgresql-15-datasketches_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-15-datasketches_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 984.6 KiB | [postgresql-15-datasketches_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 998.8 KiB | [postgresql-15-datasketches_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `datasketches_14` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 342.1 KiB | [datasketches_14-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/datasketches_14-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `datasketches_14` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 332.3 KiB | [datasketches_14-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/datasketches_14-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `datasketches_14` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 323.9 KiB | [datasketches_14-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/datasketches_14-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `datasketches_14` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 328.8 KiB | [datasketches_14-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/datasketches_14-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `datasketches_14` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 325.9 KiB | [datasketches_14-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/datasketches_14-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `datasketches_14` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 325.2 KiB | [datasketches_14-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/datasketches_14-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-datasketches` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 932.6 KiB | [postgresql-14-datasketches_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-14-datasketches_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-datasketches` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 933.4 KiB | [postgresql-14-datasketches_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-14-datasketches_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-datasketches` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 957.3 KiB | [postgresql-14-datasketches_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-14-datasketches_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-datasketches` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 957.5 KiB | [postgresql-14-datasketches_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-14-datasketches_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-datasketches` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.1 MiB | [postgresql-14-datasketches_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-14-datasketches_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-datasketches` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-14-datasketches_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-14-datasketches_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-datasketches` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 984.5 KiB | [postgresql-14-datasketches_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-14-datasketches_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-datasketches` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 998.7 KiB | [postgresql-14-datasketches_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-14-datasketches_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/apache/datasketches-postgresql" title="Repository" icon="github" subtitle="github.com/apache/datasketches-postgresql" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="apache-datasketches-postgresql-1.7.0-src.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg datasketches;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install datasketches;		# install via package name, for the active PG version

pig install datasketches -v 18;   # install for PG 18
pig install datasketches -v 17;   # install for PG 17
pig install datasketches -v 16;   # install for PG 16
pig install datasketches -v 15;   # install for PG 15
pig install datasketches -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION datasketches;
```

## Usage

Sources: [README](https://github.com/apache/datasketches-postgresql/blob/master/README.md), [latest release 1.7.0](https://github.com/apache/datasketches-postgresql/releases/tag/1.7.0), [Apache DataSketches](https://datasketches.apache.org/)

`datasketches` adds approximate analytics sketch types and aggregates to PostgreSQL. The upstream README lists CPC, HLL, Theta, Array Of Doubles, KLL, Quantiles, and Frequent Strings sketches; the 1.7.0 release is the latest published GitHub release, while the default branch has already moved on to `1.8.0-SNAPSHOT`.

```sql
CREATE EXTENSION datasketches;
```

### Core Sketch Families

- `cpc_sketch` and `hll_sketch` for approximate distinct counting.
- `theta_sketch` for distinct counting plus set operations such as union, intersection, and A-not-B.
- `aod_sketch` for tuple-style metrics keyed by identifiers with arrays of doubles.
- `kll_*_sketch` and `quantiles_*_sketch` for quantiles, ranks, PMF, and CDF.
- `frequent_strings_sketch` for heavy-hitter detection.

### Common Patterns

Build a sketch from raw values:

```sql
SELECT cpc_sketch_build(1);
SELECT kll_float_sketch_build(value) FROM normal;
```

Use one-shot approximate aggregates:

```sql
SELECT cpc_sketch_distinct(id) FROM random_ints_100m;
```

Merge sketches across groups or cube dimensions:

```sql
SELECT cpc_sketch_get_estimate(cpc_sketch_union(sketch)) FROM cpc_sketch_test;
SELECT hll_sketch_get_estimate(hll_sketch_union(sketch)) FROM hll_sketch_test;
SELECT kll_float_sketch_get_quantile(kll_float_sketch_merge(sketch), 0.5)
FROM kll_float_sketch_test;
```

Run set operations on Theta sketches:

```sql
SELECT theta_sketch_get_estimate(theta_sketch_intersection(sketch1, sketch2))
FROM theta_set_op_test;
```

Find frequent items above a threshold:

```sql
SELECT frequent_strings_sketch_result_no_false_negatives(
  frequent_strings_sketch_build(9, value),
  1000000
)
FROM zipf_1p1_8k_100m;
```

### Caveats

- Upstream documents PostgreSQL 9.6+ plus Boost 1.75.0 and DataSketches C++ core 5.0.0 or later as build dependencies.
- These are approximate structures meant to be mergeable across dimensions; they are not exact replacements for `COUNT(DISTINCT ...)` or exact histograms.
