---
title: "random"
linkTitle: "random"
description: "random data generator"
weight: 4790
categories: ["FUNC"]
width: full
---

[**pg_random**](https://github.com/tvondra/random) : random data generator


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4790** | {{< badge content="random" link="https://github.com/tvondra/random" >}} | {{< ext "random" "pg_random" >}} | `2.0.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "permuteseq" >}} {{< ext "tsm_system_rows" >}} {{< ext "tsm_system_time" >}} {{< ext "pg_idkit" >}} {{< ext "sequential_uuids" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_random` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "pg_random_18" "green" >}} {{< bg "17" "pg_random_17" "green" >}} {{< bg "16" "pg_random_16" "green" >}} {{< bg "15" "pg_random_15" "green" >}} {{< bg "14" "pg_random_14" "green" >}} | `pg_random_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "postgresql-18-random" "green" >}} {{< bg "17" "postgresql-17-random" "green" >}} {{< bg "16" "postgresql-16-random" "green" >}} {{< bg "15" "postgresql-15-random" "green" >}} {{< bg "14" "postgresql-14-random" "green" >}} | `postgresql-$v-random` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-random : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-random : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-random : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-random : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-random : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-random : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-random : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-random : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-random : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-random : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_random_18` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [pg_random_18-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_18-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_random_18` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.9 KiB | [pg_random_18-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_18-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_random_18` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.1 KiB | [pg_random_18-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_18-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_random_18` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.7 KiB | [pg_random_18-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_18-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_random_18` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_random_18-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_random_18-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_random_18` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.9 KiB | [pg_random_18-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_random_18-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-random` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.2 KiB | [postgresql-18-random_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-18-random_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-random` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.0 KiB | [postgresql-18-random_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-18-random_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-random` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.3 KiB | [postgresql-18-random_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/random/postgresql-18-random_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-random` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.1 KiB | [postgresql-18-random_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/random/postgresql-18-random_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-random` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.6 KiB | [postgresql-18-random_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-18-random_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-random` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.1 KiB | [postgresql-18-random_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-18-random_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-random` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.4 KiB | [postgresql-18-random_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-18-random_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-random` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.1 KiB | [postgresql-18-random_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-18-random_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_random_17` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [pg_random_17-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_17-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_random_17` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.9 KiB | [pg_random_17-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_17-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_random_17` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.1 KiB | [pg_random_17-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_17-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_random_17` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.7 KiB | [pg_random_17-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_17-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_random_17` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_random_17-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_random_17-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_random_17` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.9 KiB | [pg_random_17-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_random_17-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-random` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.3 KiB | [postgresql-17-random_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-random` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.0 KiB | [postgresql-17-random_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-random` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.3 KiB | [postgresql-17-random_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-random` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.1 KiB | [postgresql-17-random_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-random` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.9 KiB | [postgresql-17-random_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-random` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.4 KiB | [postgresql-17-random_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-random` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.4 KiB | [postgresql-17-random_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-random` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.0 KiB | [postgresql-17-random_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-17-random_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_random_16` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [pg_random_16-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_16-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_random_16` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.9 KiB | [pg_random_16-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_16-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_random_16` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.1 KiB | [pg_random_16-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_16-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_random_16` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.7 KiB | [pg_random_16-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_16-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_random_16` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_random_16-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_random_16-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_random_16` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.9 KiB | [pg_random_16-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_random_16-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-random` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.2 KiB | [postgresql-16-random_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-random` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.0 KiB | [postgresql-16-random_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-random` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.3 KiB | [postgresql-16-random_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-random` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.1 KiB | [postgresql-16-random_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-random` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.9 KiB | [postgresql-16-random_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-random` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.4 KiB | [postgresql-16-random_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-random` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.4 KiB | [postgresql-16-random_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-random` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.0 KiB | [postgresql-16-random_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-16-random_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_random_15` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.2 KiB | [pg_random_15-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_15-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_random_15` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.0 KiB | [pg_random_15-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_15-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_random_15` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.1 KiB | [pg_random_15-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_15-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_random_15` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.8 KiB | [pg_random_15-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_15-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_random_15` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.9 KiB | [pg_random_15-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_random_15-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_random_15` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.9 KiB | [pg_random_15-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_random_15-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-random` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.3 KiB | [postgresql-15-random_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-random` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.1 KiB | [postgresql-15-random_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-random` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.3 KiB | [postgresql-15-random_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-random` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.1 KiB | [postgresql-15-random_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-random` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.7 KiB | [postgresql-15-random_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-random` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.5 KiB | [postgresql-15-random_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-random` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.2 KiB | [postgresql-15-random_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-random` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.1 KiB | [postgresql-15-random_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-15-random_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_random_14` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [pg_random_14-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_14-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_random_14` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.0 KiB | [pg_random_14-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_14-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_random_14` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.0 KiB | [pg_random_14-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_14-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_random_14` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.8 KiB | [pg_random_14-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_14-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_random_14` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_random_14-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_random_14-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_random_14` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.9 KiB | [pg_random_14-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_random_14-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-random` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.3 KiB | [postgresql-14-random_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-random` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.1 KiB | [postgresql-14-random_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-random` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.3 KiB | [postgresql-14-random_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-random` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.1 KiB | [postgresql-14-random_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-random` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.7 KiB | [postgresql-14-random_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-random` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.4 KiB | [postgresql-14-random_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-random` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.2 KiB | [postgresql-14-random_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-random` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.0 KiB | [postgresql-14-random_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-14-random_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/random" title="Repository" icon="github" subtitle="github.com/tvondra/random" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="random-2.0.0-dev.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_random;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_random;		# install via package name, for the active PG version
pig install random;		# install by extension name, for the current active PG version

pig install random -v 18;   # install for PG 18
pig install random -v 17;   # install for PG 17
pig install random -v 16;   # install for PG 16
pig install random -v 15;   # install for PG 15
pig install random -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION random;
```



## Usage

> [random: reproducible random data generators for PostgreSQL](https://github.com/tvondra/random)

Provides functions to generate random values for various data types with reproducible output controlled by a seed.

```sql
CREATE EXTENSION random;
```

### Functions

All functions accept `seed` (for reproducibility) and `nvalues` (number of distinct values).

| Function | Description |
|---|---|
| `random_string(seed, nvalues, min_length, max_length)` | Random ASCII string |
| `random_bytea(seed, nvalues, min_length, max_length)` | Random bytea |
| `random_int(seed, nvalues, min_value, max_value)` | Random 32-bit integer |
| `random_bigint(seed, nvalues, min_value, max_value)` | Random 64-bit integer |
| `random_real(seed, nvalues, min_value, max_value)` | Random 32-bit float |
| `random_double_precision(seed, nvalues, min_value, max_value)` | Random 64-bit float |
| `random_inet(seed, nvalues)` | Random INET address (/32 mask) |
| `random_cnet(seed, nvalues)` | Random CIDR with masks 8/16/24/32 |
| `random_cnet2(seed, nvalues)` | Random CIDR with equal fraction per mask length |
| `random_macaddr(seed, nvalues)` | Random 6-byte MAC address |
| `random_macaddr8(seed, nvalues)` | Random 8-byte MAC address |

### Examples

```sql
-- Generate reproducible random integers
SELECT random_int(42, 100, 1, 1000) FROM generate_series(1, 10);

-- Random strings of length 5-10
SELECT random_string(42, 1000, 5, 10) FROM generate_series(1, 5);

-- Random IP addresses
SELECT random_inet(42, 256) FROM generate_series(1, 5);
```
