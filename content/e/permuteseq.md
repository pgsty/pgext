---
title: "permuteseq"
linkTitle: "permuteseq"
description: "Pseudo-randomly permute sequences with a format-preserving encryption on elements"
weight: 4530
categories: ["FUNC"]
width: full
---

[**permuteseq**](https://github.com/dverite/permuteseq) : Pseudo-randomly permute sequences with a format-preserving encryption on elements


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4530** | {{< badge content="permuteseq" link="https://github.com/dverite/permuteseq" >}} | {{< ext "permuteseq" >}} | `1.2.2` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "random" >}} {{< ext "sequential_uuids" >}} {{< ext "pg_hashids" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "uuid-ossp" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `permuteseq` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "permuteseq_18" "green" >}} {{< bg "17" "permuteseq_17" "green" >}} {{< bg "16" "permuteseq_16" "green" >}} {{< bg "15" "permuteseq_15" "green" >}} {{< bg "14" "permuteseq_14" "green" >}} | `permuteseq_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "postgresql-18-permuteseq" "green" >}} {{< bg "17" "postgresql-17-permuteseq" "green" >}} {{< bg "16" "postgresql-16-permuteseq" "green" >}} {{< bg "15" "postgresql-15-permuteseq" "green" >}} {{< bg "14" "postgresql-14-permuteseq" "green" >}} | `postgresql-$v-permuteseq` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "permuteseq_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-permuteseq : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-permuteseq : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-permuteseq : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-permuteseq : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-permuteseq : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-permuteseq : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-permuteseq : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-permuteseq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-permuteseq : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-permuteseq : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-permuteseq : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-permuteseq : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-permuteseq : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-permuteseq : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-permuteseq : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-permuteseq : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-permuteseq : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-permuteseq : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-permuteseq : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `permuteseq_18` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.2 KiB | [permuteseq_18-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_18-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `permuteseq_18` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.4 KiB | [permuteseq_18-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_18-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `permuteseq_18` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.1 KiB | [permuteseq_18-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_18-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `permuteseq_18` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.1 KiB | [permuteseq_18-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_18-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `permuteseq_18` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.1 KiB | [permuteseq_18-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/permuteseq_18-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `permuteseq_18` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.3 KiB | [permuteseq_18-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/permuteseq_18-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-permuteseq` | `1.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.4 KiB | [postgresql-18-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-18-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-permuteseq` | `1.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.2 KiB | [postgresql-18-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-18-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-permuteseq` | `1.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.4 KiB | [postgresql-18-permuteseq_1.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/permuteseq/postgresql-18-permuteseq_1.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-permuteseq` | `1.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.2 KiB | [postgresql-18-permuteseq_1.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/permuteseq/postgresql-18-permuteseq_1.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-permuteseq` | `1.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 15.8 KiB | [postgresql-18-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-18-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-permuteseq` | `1.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.6 KiB | [postgresql-18-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-18-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-permuteseq` | `1.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.9 KiB | [postgresql-18-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-18-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-permuteseq` | `1.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.6 KiB | [postgresql-18-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-18-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `permuteseq_17` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.2 KiB | [permuteseq_17-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_17-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `permuteseq_17` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.4 KiB | [permuteseq_17-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_17-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `permuteseq_17` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.1 KiB | [permuteseq_17-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_17-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `permuteseq_17` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.1 KiB | [permuteseq_17-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_17-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `permuteseq_17` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.2 KiB | [permuteseq_17-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/permuteseq_17-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `permuteseq_17` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.3 KiB | [permuteseq_17-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/permuteseq_17-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-permuteseq` | `1.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.4 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-permuteseq` | `1.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.2 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-permuteseq` | `1.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.4 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-permuteseq` | `1.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.2 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-permuteseq` | `1.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.3 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-permuteseq` | `1.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 16.1 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-permuteseq` | `1.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.9 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-permuteseq` | `1.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.6 KiB | [postgresql-17-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-17-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `permuteseq_16` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.2 KiB | [permuteseq_16-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_16-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `permuteseq_16` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.4 KiB | [permuteseq_16-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_16-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `permuteseq_16` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.1 KiB | [permuteseq_16-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_16-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `permuteseq_16` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.1 KiB | [permuteseq_16-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_16-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `permuteseq_16` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.2 KiB | [permuteseq_16-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/permuteseq_16-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `permuteseq_16` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.3 KiB | [permuteseq_16-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/permuteseq_16-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-permuteseq` | `1.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.4 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-permuteseq` | `1.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.2 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-permuteseq` | `1.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.4 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-permuteseq` | `1.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.2 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-permuteseq` | `1.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.3 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-permuteseq` | `1.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 16.1 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-permuteseq` | `1.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.9 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-permuteseq` | `1.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.6 KiB | [postgresql-16-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-16-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `permuteseq_15` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.0 KiB | [permuteseq_15-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_15-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `permuteseq_15` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.2 KiB | [permuteseq_15-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_15-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `permuteseq_15` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.8 KiB | [permuteseq_15-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_15-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `permuteseq_15` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.8 KiB | [permuteseq_15-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_15-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `permuteseq_15` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.8 KiB | [permuteseq_15-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/permuteseq_15-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `permuteseq_15` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.0 KiB | [permuteseq_15-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/permuteseq_15-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-permuteseq` | `1.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 14.5 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-permuteseq` | `1.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 14.4 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-permuteseq` | `1.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.5 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-permuteseq` | `1.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 14.5 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-permuteseq` | `1.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 15.2 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-permuteseq` | `1.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.2 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-permuteseq` | `1.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.0 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-permuteseq` | `1.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.8 KiB | [postgresql-15-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-15-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `permuteseq_14` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.0 KiB | [permuteseq_14-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/permuteseq_14-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `permuteseq_14` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.2 KiB | [permuteseq_14-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/permuteseq_14-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `permuteseq_14` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.8 KiB | [permuteseq_14-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/permuteseq_14-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `permuteseq_14` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.8 KiB | [permuteseq_14-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/permuteseq_14-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `permuteseq_14` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.8 KiB | [permuteseq_14-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/permuteseq_14-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `permuteseq_14` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.0 KiB | [permuteseq_14-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/permuteseq_14-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-permuteseq` | `1.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 14.5 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-permuteseq` | `1.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 14.4 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-permuteseq` | `1.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.5 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-permuteseq` | `1.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 14.4 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-permuteseq` | `1.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 15.1 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-permuteseq` | `1.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.1 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-permuteseq` | `1.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.0 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-permuteseq` | `1.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.8 KiB | [postgresql-14-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/permuteseq/postgresql-14-permuteseq_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dverite/permuteseq" title="Repository" icon="github" subtitle="github.com/dverite/permuteseq" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="permuteseq-1.2.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg permuteseq;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install permuteseq;		# install via package name, for the active PG version

pig install permuteseq -v 18;   # install for PG 18
pig install permuteseq -v 17;   # install for PG 17
pig install permuteseq -v 16;   # install for PG 16
pig install permuteseq -v 15;   # install for PG 15
pig install permuteseq -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION permuteseq;
```



## Usage

> [permuteseq: scalable pseudo-random permutations of sequences](https://github.com/dverite/permuteseq)

Generate unique, non-sequential, random-looking series of numbers without storing previous values. Uses a Feistel cipher with cycle-walking for format-preserving encryption.

```sql
CREATE EXTENSION permuteseq;
```

### Functions

| Function | Description |
|---|---|
| `permute_nextval(seq_oid, crypt_key bigint)` | Advance sequence and return encrypted value within sequence bounds |
| `reverse_permute(seq_oid, value bigint, crypt_key bigint)` | Compute original value from its permuted element |
| `range_encrypt_element(clear_val bigint, min_val bigint, max_val bigint, crypt_key bigint)` | Encrypt a bigint in a given range |
| `range_decrypt_element(crypt_val bigint, min_val bigint, max_val bigint, crypt_key bigint)` | Decrypt a previously encrypted value |

### Examples

```sql
CREATE SEQUENCE s MINVALUE -10000 MAXVALUE 15000;

-- Generate random-looking unique values from a sequence
SELECT permute_nextval('s'::regclass, 123456789012345)
  FROM generate_series(1, 10);

-- Reverse a permuted value back to the original
SELECT reverse_permute('s'::regclass, -545, 123456789012345);
-- -10000

-- Encrypt/decrypt within an arbitrary range
SELECT range_encrypt_element(91919191919, 1e10::bigint, 1e11::bigint, 123456789012345);
-- 83028080992

SELECT range_decrypt_element(83028080992, 1e10::bigint, 1e11::bigint, 123456789012345);
-- 91919191919
```
