---
title: "qos"
linkTitle: "qos"
description: "QoS resource governor extension for PostgreSQL sessions and queries"
weight: 5240
categories: ["ADMIN"]
width: full
---

[**pg_qos**](https://github.com/appstonia/pg_qos) : QoS resource governor extension for PostgreSQL sessions and queries


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5240** | {{< badge content="qos" link="https://github.com/appstonia/pg_qos" >}} | {{< ext "qos" "pg_qos" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prioritize" >}} {{< ext "pg_permissions" >}} {{< ext "pg_readonly" >}} {{< ext "pg_crash" >}} {{< ext "pg_cooldown" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_repack" >}} {{< ext "pgfincore" >}} |

> [!Note] requires shared_preload_libraries = 'qos'; official support PG15+


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_qos` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_qos_18" "green" >}} {{< bg "17" "pg_qos_17" "green" >}} {{< bg "16" "pg_qos_16" "green" >}} {{< bg "15" "pg_qos_15" "green" >}} {{< bg "14" "pg_qos_14" "red" >}} | `pg_qos_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-qos" "green" >}} {{< bg "17" "postgresql-17-qos" "green" >}} {{< bg "16" "postgresql-16-qos" "green" >}} {{< bg "15" "postgresql-15-qos" "green" >}} {{< bg "14" "postgresql-14-qos" "red" >}} | `postgresql-$v-qos` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_qos_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_qos_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_qos_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_qos_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_qos_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_qos_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_qos_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-qos : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-qos : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-qos : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-qos : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-qos : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-qos : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-qos : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-qos : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-qos : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-qos : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qos_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.2 KiB | [pg_qos_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_qos_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_qos_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.0 KiB | [pg_qos_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_qos_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_qos_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.3 KiB | [pg_qos_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_qos_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_qos_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.3 KiB | [pg_qos_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_qos_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_qos_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.7 KiB | [pg_qos_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_qos_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_qos_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.6 KiB | [pg_qos_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_qos_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-qos` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.3 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-qos` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 68.5 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-qos` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.6 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-qos` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 68.6 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-qos` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 73.7 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-qos` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 73.1 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-qos` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 71.7 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-qos` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.4 KiB | [postgresql-18-qos_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-18-qos_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qos_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.2 KiB | [pg_qos_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_qos_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_qos_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.0 KiB | [pg_qos_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_qos_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_qos_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.5 KiB | [pg_qos_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_qos_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_qos_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.5 KiB | [pg_qos_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_qos_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_qos_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.9 KiB | [pg_qos_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_qos_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_qos_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.8 KiB | [pg_qos_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_qos_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-qos` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.3 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-qos` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 68.6 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-qos` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.6 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-qos` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 68.7 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-qos` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 81.3 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-qos` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 80.9 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-qos` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 71.9 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-qos` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.5 KiB | [postgresql-17-qos_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-17-qos_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qos_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.2 KiB | [pg_qos_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_qos_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_qos_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.9 KiB | [pg_qos_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_qos_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_qos_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.4 KiB | [pg_qos_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_qos_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_qos_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.4 KiB | [pg_qos_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_qos_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_qos_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.8 KiB | [pg_qos_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_qos_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_qos_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.7 KiB | [pg_qos_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_qos_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-qos` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.2 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-qos` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 68.3 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-qos` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.5 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-qos` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 68.4 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-qos` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 79.9 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-qos` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 79.5 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-qos` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 71.8 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-qos` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.3 KiB | [postgresql-16-qos_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-16-qos_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_qos_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.5 KiB | [pg_qos_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_qos_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_qos_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.3 KiB | [pg_qos_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_qos_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_qos_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.2 KiB | [pg_qos_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_qos_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_qos_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.3 KiB | [pg_qos_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_qos_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_qos_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.6 KiB | [pg_qos_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_qos_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_qos_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.5 KiB | [pg_qos_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_qos_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-qos` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 69.4 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-qos` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 68.4 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-qos` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 69.4 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-qos` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 68.5 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-qos` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 80.0 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-qos` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 80.0 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-qos` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 72.0 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-qos` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.9 KiB | [postgresql-15-qos_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qos/postgresql-15-qos_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/appstonia/pg_qos" title="Repository" icon="github" subtitle="github.com/appstonia/pg_qos" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_qos-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_qos;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_qos;		# install via package name, for the active PG version
pig install qos;		# install by extension name, for the current active PG version

pig install qos -v 18;   # install for PG 18
pig install qos -v 17;   # install for PG 17
pig install qos -v 16;   # install for PG 16
pig install qos -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'qos';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION qos;
```




## Usage

> [qos: QoS resource governor extension for PostgreSQL sessions and queries](https://github.com/appstonia/pg_qos)

The `qos` extension provides Quality of Service resource governance for PostgreSQL, allowing administrators to set per-role and per-database limits on memory usage, CPU cores, and concurrent transactions/statements.

### Configuration Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `qos.work_mem_limit` | bytes | Maximum effective `work_mem` per session |
| `qos.cpu_core_limit` | integer | Maximum CPU cores available to a session |
| `qos.max_concurrent_tx` | integer | Maximum concurrent transactions |
| `qos.max_concurrent_select` | integer | Maximum concurrent SELECT statements |
| `qos.max_concurrent_update` | integer | Maximum concurrent UPDATE statements |
| `qos.max_concurrent_delete` | integer | Maximum concurrent DELETE statements |
| `qos.max_concurrent_insert` | integer | Maximum concurrent INSERT statements |

### Per-Role Limits

```sql
ALTER ROLE app_user SET qos.work_mem_limit = '32MB';
ALTER ROLE app_user SET qos.cpu_core_limit = '2';
ALTER ROLE app_user SET qos.max_concurrent_select = '100';
```

### Per-Database Limits

```sql
ALTER DATABASE appdb SET qos.max_concurrent_tx = '200';
```

### Combined Role + Database Limits

```sql
ALTER ROLE app_user IN DATABASE appdb SET qos.work_mem_limit = '4MB';
ALTER ROLE app_user IN DATABASE appdb SET qos.max_concurrent_update = '10';
```

### Enforcement Behavior

- **Work memory**: Intercepts `SET work_mem` and rejects values exceeding configured limits
- **CPU limiting** (Linux only): Binds backend to N CPU cores via CPU affinity; on non-Linux platforms, limits parallel workers instead
- **Concurrency**: Executor hooks track active transactions/statements by type; violations block execution

### Observability

```sql
SET client_min_messages = 'debug1';  -- enable debug output for QoS events
```

The most restrictive combination of role-level and database-level settings takes effect. Requires `shared_preload_libraries = 'qos'` and PostgreSQL 15+.
