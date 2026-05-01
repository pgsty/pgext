---
title: "plan_filter"
linkTitle: "plan_filter"
description: "filter statements by their execution plans."
weight: 2810
categories: ["FEAT"]
width: full
---

[**pg_plan_filter**](https://github.com/pgexperts/pg_plan_filter) : filter statements by their execution plans.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2810** | {{< badge content="plan_filter" link="https://github.com/pgexperts/pg_plan_filter" >}} | {{< ext "plan_filter" "pg_plan_filter" >}} | `0.0.1` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_plan_filter` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "pg_plan_filter_18" "green" >}} {{< bg "17" "pg_plan_filter_17" "green" >}} {{< bg "16" "pg_plan_filter_16" "green" >}} {{< bg "15" "pg_plan_filter_15" "green" >}} {{< bg "14" "pg_plan_filter_14" "green" >}} | `pg_plan_filter_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-plan-filter" "green" >}} {{< bg "17" "postgresql-17-pg-plan-filter" "green" >}} {{< bg "16" "postgresql-16-pg-plan-filter" "green" >}} {{< bg "15" "postgresql-15-pg-plan-filter" "green" >}} {{< bg "14" "postgresql-14-pg-plan-filter" "green" >}} | `postgresql-$v-pg-plan-filter` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_plan_filter_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-plan-filter : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-plan-filter : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-plan-filter : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-plan-filter : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-plan-filter : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-plan-filter : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-plan-filter : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-plan-filter : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-plan-filter : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-plan-filter : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-plan-filter : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_plan_filter_18` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.8 KiB | [pg_plan_filter_18-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_18-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_plan_filter_18` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.0 KiB | [pg_plan_filter_18-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_18-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_plan_filter_18` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.5 KiB | [pg_plan_filter_18-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_18-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_plan_filter_18` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.5 KiB | [pg_plan_filter_18-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_18-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_plan_filter_18` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.4 KiB | [pg_plan_filter_18-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_plan_filter_18-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_plan_filter_18` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.6 KiB | [pg_plan_filter_18-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_plan_filter_18-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-plan-filter` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.9 KiB | [postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-plan-filter` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.1 KiB | [postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-plan-filter` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.9 KiB | [postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-plan-filter/postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-plan-filter` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.1 KiB | [postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-plan-filter/postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-plan-filter` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.1 KiB | [postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-plan-filter` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.1 KiB | [postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-plan-filter` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.1 KiB | [postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-plan-filter` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.2 KiB | [postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-plan-filter` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 8.8 KiB | [postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-plan-filter/postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-plan-filter` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.9 KiB | [postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-plan-filter/postgresql-18-pg-plan-filter_0.0.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_plan_filter_17` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.8 KiB | [pg_plan_filter_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_plan_filter_17` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.0 KiB | [pg_plan_filter_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_plan_filter_17` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.5 KiB | [pg_plan_filter_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_plan_filter_17` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.5 KiB | [pg_plan_filter_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_plan_filter_17` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pg_plan_filter_17-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_plan_filter_17-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_plan_filter_17` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.7 KiB | [pg_plan_filter_17-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_plan_filter_17-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-plan-filter` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.9 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-plan-filter` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.1 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-plan-filter` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.9 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-plan-filter` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.1 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-plan-filter` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.6 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-plan-filter` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.6 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-plan-filter` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.2 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-plan-filter` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-plan-filter` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 8.8 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-plan-filter` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.9 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_plan_filter_16` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.8 KiB | [pg_plan_filter_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_plan_filter_16` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.0 KiB | [pg_plan_filter_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_plan_filter_16` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.5 KiB | [pg_plan_filter_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_plan_filter_16` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.5 KiB | [pg_plan_filter_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_plan_filter_16` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pg_plan_filter_16-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_plan_filter_16-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_plan_filter_16` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.7 KiB | [pg_plan_filter_16-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_plan_filter_16-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-plan-filter` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.9 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-plan-filter` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.0 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-plan-filter` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.9 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-plan-filter` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.1 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-plan-filter` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.6 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-plan-filter` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.6 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-plan-filter` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.2 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-plan-filter` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-plan-filter` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 8.8 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-plan-filter` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.9 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_plan_filter_15` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.8 KiB | [pg_plan_filter_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_plan_filter_15` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.0 KiB | [pg_plan_filter_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_plan_filter_15` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.5 KiB | [pg_plan_filter_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_plan_filter_15` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.5 KiB | [pg_plan_filter_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_plan_filter_15` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pg_plan_filter_15-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_plan_filter_15-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_plan_filter_15` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.7 KiB | [pg_plan_filter_15-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_plan_filter_15-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-plan-filter` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.9 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-plan-filter` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.1 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-plan-filter` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.9 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-plan-filter` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.1 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-plan-filter` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.5 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-plan-filter` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.6 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-plan-filter` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.2 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-plan-filter` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.3 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-plan-filter` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 8.8 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-plan-filter` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.9 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_plan_filter_14` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.8 KiB | [pg_plan_filter_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_plan_filter_14` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.0 KiB | [pg_plan_filter_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_plan_filter_14` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.5 KiB | [pg_plan_filter_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_plan_filter_14` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.5 KiB | [pg_plan_filter_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_plan_filter_14` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pg_plan_filter_14-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_plan_filter_14-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_plan_filter_14` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.7 KiB | [pg_plan_filter_14-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_plan_filter_14-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-plan-filter` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.9 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-plan-filter` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.0 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-plan-filter` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.8 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-plan-filter` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.1 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-plan-filter` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.5 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-plan-filter` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.6 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-plan-filter` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.1 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-plan-filter` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.2 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-plan-filter` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 8.8 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-plan-filter` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.9 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgexperts/pg_plan_filter" title="Repository" icon="github" subtitle="github.com/pgexperts/pg_plan_filter" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_plan_filter.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_plan_filter;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_plan_filter;		# install via package name, for the active PG version
pig install plan_filter;		# install by extension name, for the current active PG version

pig install plan_filter -v 18;   # install for PG 18
pig install plan_filter -v 17;   # install for PG 17
pig install plan_filter -v 16;   # install for PG 16
pig install plan_filter -v 15;   # install for PG 15
pig install plan_filter -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'plan_filter';
```


This extension does not need `CREATE EXTENSION` DDL command






## Usage

> [plan_filter: filter statements by their execution plans](https://github.com/pgexperts/pg_plan_filter)

The `pg_plan_filter` module tests statements against specific configured criteria before execution, raising an error if the criteria are violated. This allows administrators to prevent execution of certain queries on production databases.

The only criterion currently supported is the maximum allowed estimated cost of the statement plan.

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'plan_filter'
plan_filter.statement_cost_limit = 100000.0
```

The `statement_cost_limit` must be set to a non-zero value to enable filtering. The default is `0` (no filtering).

### GUC Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `plan_filter.statement_cost_limit` | float | `0` | Maximum allowed estimated plan cost. `0` disables filtering |
| `plan_filter.limit_select_only` | bool | `false` | When true, only filter `SELECT` statements |

### Examples

Prevent expensive queries globally:

```ini
plan_filter.statement_cost_limit = 100000.0
```

Limit filtering to SELECT statements only (note: SELECT != READONLY, since SELECT can also modify data):

```ini
plan_filter.limit_select_only = true
```

When the module is running with a non-zero `statement_cost_limit`, it will also prevent `EXPLAIN` on expensive queries. Temporarily bypass the filter:

```sql
BEGIN;
SET LOCAL plan_filter.statement_cost_limit = 0;
EXPLAIN SELECT ...;
COMMIT;
```

Override the limit per user:

```sql
ALTER USER can_run_expensive SET plan_filter.statement_cost_limit = 0;
ALTER USER only_cheap_queries SET plan_filter.statement_cost_limit = 10000;
```

### Caveats

The `statement_cost_limit` cancels plans based on their **estimated cost**. The PostgreSQL planner can return cost estimates unrelated to actual query execution time. Be prepared for false positive cancellations and set the limit generously.
