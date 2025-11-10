---
title: "random"
linkTitle: "random"
description: "random data generator"
weight: 4790
categories: ["FUNC"]
width: full
---

[**pg_random**](https://github.com/tvondra/random)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4790** | {{< badge content="random" link="https://github.com/tvondra/random" >}} | {{< ext "random" "pg_random" >}} | `2.0.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "permuteseq" >}} {{< ext "tsm_system_rows" >}} {{< ext "tsm_system_time" >}} {{< ext "pg_idkit" >}} {{< ext "sequential_uuids" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/random" >}} | `2.0.0` | {{< bg "18" "pg_random_18*" "green" >}} {{< bg "17" "pg_random_17*" "green" >}} {{< bg "16" "pg_random_16*" "green" >}} {{< bg "15" "pg_random_15*" "green" >}} {{< bg "14" "pg_random_14*" "green" >}} {{< bg "13" "pg_random_13*" "green" >}} | `pg_random_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/random" >}} | `2.0.0` | {{< bg "18" "postgresql-18-random" "green" >}} {{< bg "17" "postgresql-17-random" "green" >}} {{< bg "16" "postgresql-16-random" "green" >}} {{< bg "15" "postgresql-15-random" "green" >}} {{< bg "14" "postgresql-14-random" "green" >}} {{< bg "13" "postgresql-13-random" "green" >}} | `postgresql-$v-random` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_random_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-13-random : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-13-random : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-13-random : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-13-random : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-13-random : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-13-random : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-13-random : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-random : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-13-random : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

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
{{< tab >}}

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
{{< tab >}}

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
{{< tab >}}

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
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_random_13` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.0 KiB | [pg_random_13-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_random_13-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_random_13` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.0 KiB | [pg_random_13-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_random_13-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_random_13` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.0 KiB | [pg_random_13-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_random_13-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_random_13` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.8 KiB | [pg_random_13-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_random_13-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_random_13` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.0 KiB | [pg_random_13-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_random_13-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_random_13` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.0 KiB | [pg_random_13-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_random_13-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-random` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.3 KiB | [postgresql-13-random_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-random` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.9 KiB | [postgresql-13-random_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-random` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.4 KiB | [postgresql-13-random_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-random` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.9 KiB | [postgresql-13-random_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-random` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.7 KiB | [postgresql-13-random_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-random` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.4 KiB | [postgresql-13-random_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-random` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.3 KiB | [postgresql-13-random_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-random` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.8 KiB | [postgresql-13-random_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/random/postgresql-13-random_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/random" title="Repository" icon="github" subtitle="github.com/tvondra/random" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="random-2.0.0-dev.tar.gz" >}}
{{< /cards >}}


```bash
pig build get random; # get random source code
pig build dep random; # install build dependencies
pig build pkg random; # build extension rpm or deb
pig build ext random; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install random; # install by extension name, for the current active PG version
pig ext install pg_random; # install via package alias, for the active PG version
pig ext install random -v 18;   # install for PG 18
pig ext install random -v 17;   # install for PG 17
pig ext install random -v 16;   # install for PG 16
pig ext install random -v 15;   # install for PG 15
pig ext install random -v 14;   # install for PG 14
pig ext install random -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION random;
```

