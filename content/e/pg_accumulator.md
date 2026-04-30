---
title: "pg_accumulator"
linkTitle: "pg_accumulator"
description: "Accumulation registers for balance and turnover tracking in PostgreSQL"
weight: 4845
categories: ["FUNC"]
width: full
---

[**pg_accumulator**](https://github.com/Treedo/pg_accumulator) : Accumulation registers for balance and turnover tracking in PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4845** | {{< badge content="pg_accumulator" link="https://github.com/Treedo/pg_accumulator" >}} | {{< ext "pg_accumulator" >}} | `1.1.3` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `accum` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "financial" >}} {{< ext "topn" >}} {{< ext "quantile" >}} {{< ext "first_last_agg" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_accumulator` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.3` | {{< bg "18" "pg_accumulator_18" "green" >}} {{< bg "17" "pg_accumulator_17" "green" >}} {{< bg "16" "pg_accumulator_16" "green" >}} {{< bg "15" "pg_accumulator_15" "green" >}} {{< bg "14" "pg_accumulator_14" "green" >}} | `pg_accumulator_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.3` | {{< bg "18" "postgresql-18-pg-accumulator" "green" >}} {{< bg "17" "postgresql-17-pg-accumulator" "green" >}} {{< bg "16" "postgresql-16-pg-accumulator" "green" >}} {{< bg "15" "postgresql-15-pg-accumulator" "green" >}} {{< bg "14" "postgresql-14-pg-accumulator" "green" >}} | `postgresql-$v-pg-accumulator` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "pg_accumulator_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-18-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-17-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-16-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-15-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-14-pg-accumulator : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-18-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-17-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-16-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-15-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-14-pg-accumulator : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-18-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-17-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-16-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-15-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-14-pg-accumulator : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-18-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-17-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-16-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-15-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-14-pg-accumulator : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-18-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-17-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-16-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-15-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-14-pg-accumulator : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-18-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-17-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-16-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-15-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-14-pg-accumulator : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-accumulator : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-accumulator : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-18-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-17-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-16-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-15-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-14-pg-accumulator : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-18-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-17-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-16-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-15-pg-accumulator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.3" "postgresql-14-pg-accumulator : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_accumulator_18` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.6 KiB | [pg_accumulator_18-1.1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_accumulator_18-1.1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_accumulator_18` | `1.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 56.8 KiB | [pg_accumulator_18-1.1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_accumulator_18-1.1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_accumulator_18` | `1.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.3 KiB | [pg_accumulator_18-1.1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_accumulator_18-1.1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_accumulator_18` | `1.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.2 KiB | [pg_accumulator_18-1.1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_accumulator_18-1.1.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_accumulator_18` | `1.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.3 KiB | [pg_accumulator_18-1.1.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_accumulator_18-1.1.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_accumulator_18` | `1.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.5 KiB | [pg_accumulator_18-1.1.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_accumulator_18-1.1.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-accumulator` | `1.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.1 KiB | [postgresql-18-pg-accumulator_1.1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-accumulator/postgresql-18-pg-accumulator_1.1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-accumulator` | `1.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.9 KiB | [postgresql-18-pg-accumulator_1.1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-accumulator/postgresql-18-pg-accumulator_1.1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-accumulator` | `1.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.2 KiB | [postgresql-18-pg-accumulator_1.1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-accumulator/postgresql-18-pg-accumulator_1.1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-accumulator` | `1.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 50.9 KiB | [postgresql-18-pg-accumulator_1.1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-accumulator/postgresql-18-pg-accumulator_1.1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-accumulator` | `1.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 52.8 KiB | [postgresql-18-pg-accumulator_1.1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-accumulator/postgresql-18-pg-accumulator_1.1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-accumulator` | `1.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 52.5 KiB | [postgresql-18-pg-accumulator_1.1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-accumulator/postgresql-18-pg-accumulator_1.1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-accumulator` | `1.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 52.3 KiB | [postgresql-18-pg-accumulator_1.1.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-accumulator/postgresql-18-pg-accumulator_1.1.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-accumulator` | `1.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 52.1 KiB | [postgresql-18-pg-accumulator_1.1.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-accumulator/postgresql-18-pg-accumulator_1.1.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_accumulator_17` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.6 KiB | [pg_accumulator_17-1.1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_accumulator_17-1.1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_accumulator_17` | `1.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 56.8 KiB | [pg_accumulator_17-1.1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_accumulator_17-1.1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_accumulator_17` | `1.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.3 KiB | [pg_accumulator_17-1.1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_accumulator_17-1.1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_accumulator_17` | `1.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.2 KiB | [pg_accumulator_17-1.1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_accumulator_17-1.1.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_accumulator_17` | `1.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.2 KiB | [pg_accumulator_17-1.1.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_accumulator_17-1.1.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_accumulator_17` | `1.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.5 KiB | [pg_accumulator_17-1.1.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_accumulator_17-1.1.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-accumulator` | `1.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.1 KiB | [postgresql-17-pg-accumulator_1.1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-accumulator/postgresql-17-pg-accumulator_1.1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-accumulator` | `1.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.9 KiB | [postgresql-17-pg-accumulator_1.1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-accumulator/postgresql-17-pg-accumulator_1.1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-accumulator` | `1.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.1 KiB | [postgresql-17-pg-accumulator_1.1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-accumulator/postgresql-17-pg-accumulator_1.1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-accumulator` | `1.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 50.9 KiB | [postgresql-17-pg-accumulator_1.1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-accumulator/postgresql-17-pg-accumulator_1.1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-accumulator` | `1.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 54.1 KiB | [postgresql-17-pg-accumulator_1.1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-accumulator/postgresql-17-pg-accumulator_1.1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-accumulator` | `1.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 53.8 KiB | [postgresql-17-pg-accumulator_1.1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-accumulator/postgresql-17-pg-accumulator_1.1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-accumulator` | `1.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 52.3 KiB | [postgresql-17-pg-accumulator_1.1.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-accumulator/postgresql-17-pg-accumulator_1.1.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-accumulator` | `1.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 52.1 KiB | [postgresql-17-pg-accumulator_1.1.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-accumulator/postgresql-17-pg-accumulator_1.1.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_accumulator_16` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.6 KiB | [pg_accumulator_16-1.1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_accumulator_16-1.1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_accumulator_16` | `1.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 56.8 KiB | [pg_accumulator_16-1.1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_accumulator_16-1.1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_accumulator_16` | `1.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.3 KiB | [pg_accumulator_16-1.1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_accumulator_16-1.1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_accumulator_16` | `1.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.2 KiB | [pg_accumulator_16-1.1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_accumulator_16-1.1.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_accumulator_16` | `1.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.3 KiB | [pg_accumulator_16-1.1.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_accumulator_16-1.1.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_accumulator_16` | `1.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.5 KiB | [pg_accumulator_16-1.1.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_accumulator_16-1.1.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-accumulator` | `1.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.1 KiB | [postgresql-16-pg-accumulator_1.1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-accumulator/postgresql-16-pg-accumulator_1.1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-accumulator` | `1.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.9 KiB | [postgresql-16-pg-accumulator_1.1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-accumulator/postgresql-16-pg-accumulator_1.1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-accumulator` | `1.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.1 KiB | [postgresql-16-pg-accumulator_1.1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-accumulator/postgresql-16-pg-accumulator_1.1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-accumulator` | `1.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.0 KiB | [postgresql-16-pg-accumulator_1.1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-accumulator/postgresql-16-pg-accumulator_1.1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-accumulator` | `1.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 54.1 KiB | [postgresql-16-pg-accumulator_1.1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-accumulator/postgresql-16-pg-accumulator_1.1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-accumulator` | `1.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 53.8 KiB | [postgresql-16-pg-accumulator_1.1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-accumulator/postgresql-16-pg-accumulator_1.1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-accumulator` | `1.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 52.3 KiB | [postgresql-16-pg-accumulator_1.1.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-accumulator/postgresql-16-pg-accumulator_1.1.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-accumulator` | `1.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 52.1 KiB | [postgresql-16-pg-accumulator_1.1.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-accumulator/postgresql-16-pg-accumulator_1.1.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_accumulator_15` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.7 KiB | [pg_accumulator_15-1.1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_accumulator_15-1.1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_accumulator_15` | `1.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 56.9 KiB | [pg_accumulator_15-1.1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_accumulator_15-1.1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_accumulator_15` | `1.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.3 KiB | [pg_accumulator_15-1.1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_accumulator_15-1.1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_accumulator_15` | `1.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.4 KiB | [pg_accumulator_15-1.1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_accumulator_15-1.1.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_accumulator_15` | `1.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.4 KiB | [pg_accumulator_15-1.1.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_accumulator_15-1.1.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_accumulator_15` | `1.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.7 KiB | [pg_accumulator_15-1.1.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_accumulator_15-1.1.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-accumulator` | `1.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.2 KiB | [postgresql-15-pg-accumulator_1.1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-accumulator/postgresql-15-pg-accumulator_1.1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-accumulator` | `1.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.9 KiB | [postgresql-15-pg-accumulator_1.1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-accumulator/postgresql-15-pg-accumulator_1.1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-accumulator` | `1.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.2 KiB | [postgresql-15-pg-accumulator_1.1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-accumulator/postgresql-15-pg-accumulator_1.1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-accumulator` | `1.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.0 KiB | [postgresql-15-pg-accumulator_1.1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-accumulator/postgresql-15-pg-accumulator_1.1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-accumulator` | `1.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 54.2 KiB | [postgresql-15-pg-accumulator_1.1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-accumulator/postgresql-15-pg-accumulator_1.1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-accumulator` | `1.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 53.9 KiB | [postgresql-15-pg-accumulator_1.1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-accumulator/postgresql-15-pg-accumulator_1.1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-accumulator` | `1.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 52.4 KiB | [postgresql-15-pg-accumulator_1.1.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-accumulator/postgresql-15-pg-accumulator_1.1.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-accumulator` | `1.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 52.3 KiB | [postgresql-15-pg-accumulator_1.1.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-accumulator/postgresql-15-pg-accumulator_1.1.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_accumulator_14` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.7 KiB | [pg_accumulator_14-1.1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_accumulator_14-1.1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_accumulator_14` | `1.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 56.9 KiB | [pg_accumulator_14-1.1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_accumulator_14-1.1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_accumulator_14` | `1.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.3 KiB | [pg_accumulator_14-1.1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_accumulator_14-1.1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_accumulator_14` | `1.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.4 KiB | [pg_accumulator_14-1.1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_accumulator_14-1.1.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_accumulator_14` | `1.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.4 KiB | [pg_accumulator_14-1.1.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_accumulator_14-1.1.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_accumulator_14` | `1.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.7 KiB | [pg_accumulator_14-1.1.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_accumulator_14-1.1.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-accumulator` | `1.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 51.1 KiB | [postgresql-14-pg-accumulator_1.1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-accumulator/postgresql-14-pg-accumulator_1.1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-accumulator` | `1.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.9 KiB | [postgresql-14-pg-accumulator_1.1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-accumulator/postgresql-14-pg-accumulator_1.1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-accumulator` | `1.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 51.2 KiB | [postgresql-14-pg-accumulator_1.1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-accumulator/postgresql-14-pg-accumulator_1.1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-accumulator` | `1.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.0 KiB | [postgresql-14-pg-accumulator_1.1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-accumulator/postgresql-14-pg-accumulator_1.1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-accumulator` | `1.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 54.2 KiB | [postgresql-14-pg-accumulator_1.1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-accumulator/postgresql-14-pg-accumulator_1.1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-accumulator` | `1.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 53.9 KiB | [postgresql-14-pg-accumulator_1.1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-accumulator/postgresql-14-pg-accumulator_1.1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-accumulator` | `1.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 52.3 KiB | [postgresql-14-pg-accumulator_1.1.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-accumulator/postgresql-14-pg-accumulator_1.1.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-accumulator` | `1.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 52.2 KiB | [postgresql-14-pg-accumulator_1.1.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-accumulator/postgresql-14-pg-accumulator_1.1.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Treedo/pg_accumulator" title="Repository" icon="github" subtitle="github.com/Treedo/pg_accumulator" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_accumulator-1.1.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_accumulator;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_accumulator;		# install via package name, for the active PG version

pig install pg_accumulator -v 18;   # install for PG 18
pig install pg_accumulator -v 17;   # install for PG 17
pig install pg_accumulator -v 16;   # install for PG 16
pig install pg_accumulator -v 15;   # install for PG 15
pig install pg_accumulator -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_accumulator CASCADE; -- requires plpgsql
```
