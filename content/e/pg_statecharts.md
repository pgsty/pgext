---
title: "pg_statecharts"
linkTitle: "pg_statecharts"
description: "Hierarchical state machines and transactional event handling"
weight: 2695
categories: ["FEAT"]
languages: ["SQL"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_statecharts**](https://github.com/kronor-io/statecharts) : Hierarchical state machines and transactional event handling


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2695** | {{< badge content="pg_statecharts" link="https://github.com/kronor-io/statecharts" >}} | {{< ext "pg_statecharts" >}} | `0.1.0` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `fsm` |
|   **Requires**    | {{< ext "ltree" >}} |
|    **Need By**    | {{< ext "pg_statecharts_dev" >}} |
|    **Siblings**   | {{< ext "pg_statecharts_dev" >}} |

> [!Note] ltree in public


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_statecharts` | `ltree` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_statecharts_18" "green" >}} {{< bg "17" "pg_statecharts_17" "green" >}} {{< bg "16" "pg_statecharts_16" "green" >}} {{< bg "15" "pg_statecharts_15" "red" >}} {{< bg "14" "pg_statecharts_14" "red" >}} | `pg_statecharts_$v` | `postgresql$v-contrib` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pg-statecharts" "green" >}} {{< bg "17" "postgresql-17-pg-statecharts" "green" >}} {{< bg "16" "postgresql-16-pg-statecharts" "green" >}} {{< bg "15" "postgresql-15-pg-statecharts" "red" >}} {{< bg "14" "postgresql-14-pg-statecharts" "red" >}} | `postgresql-$v-pg-statecharts` | `postgresql-$v` |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_statecharts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_statecharts_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_statecharts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_statecharts_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_statecharts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_statecharts_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_statecharts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_statecharts_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_statecharts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_statecharts_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_statecharts_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_statecharts_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_statecharts_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-statecharts : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-statecharts : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-statecharts : N/A 0" "gray" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statecharts_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.8 KiB | [pg_statecharts_18-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_statecharts_18-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_statecharts_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.8 KiB | [pg_statecharts_18-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_statecharts_18-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_statecharts_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.8 KiB | [pg_statecharts_18-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_statecharts_18-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_statecharts_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.8 KiB | [pg_statecharts_18-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_statecharts_18-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_statecharts_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.9 KiB | [pg_statecharts_18-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_statecharts_18-0.1.0-1PGSTY.el10.noarch.rpm) |
| `pg_statecharts_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.9 KiB | [pg_statecharts_18-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_statecharts_18-0.1.0-1PGSTY.el10.noarch.rpm) |
| `postgresql-18-pg-statecharts` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.1 KiB | [postgresql-18-pg-statecharts_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statecharts/postgresql-18-pg-statecharts_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-statecharts` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 25.1 KiB | [postgresql-18-pg-statecharts_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statecharts/postgresql-18-pg-statecharts_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-statecharts` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.1 KiB | [postgresql-18-pg-statecharts_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statecharts/postgresql-18-pg-statecharts_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-statecharts` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.1 KiB | [postgresql-18-pg-statecharts_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statecharts/postgresql-18-pg-statecharts_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-statecharts` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.7 KiB | [postgresql-18-pg-statecharts_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statecharts/postgresql-18-pg-statecharts_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-statecharts` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.7 KiB | [postgresql-18-pg-statecharts_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statecharts/postgresql-18-pg-statecharts_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-statecharts` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-18-pg-statecharts_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statecharts/postgresql-18-pg-statecharts_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-statecharts` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-18-pg-statecharts_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statecharts/postgresql-18-pg-statecharts_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-statecharts` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 25.5 KiB | [postgresql-18-pg-statecharts_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statecharts/postgresql-18-pg-statecharts_0.1.0-1PGSTY~resolute_all.deb) |
| `postgresql-18-pg-statecharts` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 25.5 KiB | [postgresql-18-pg-statecharts_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statecharts/postgresql-18-pg-statecharts_0.1.0-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statecharts_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.8 KiB | [pg_statecharts_17-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_statecharts_17-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_statecharts_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.8 KiB | [pg_statecharts_17-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_statecharts_17-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_statecharts_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.8 KiB | [pg_statecharts_17-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_statecharts_17-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_statecharts_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.8 KiB | [pg_statecharts_17-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_statecharts_17-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_statecharts_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.9 KiB | [pg_statecharts_17-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_statecharts_17-0.1.0-1PGSTY.el10.noarch.rpm) |
| `pg_statecharts_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.9 KiB | [pg_statecharts_17-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_statecharts_17-0.1.0-1PGSTY.el10.noarch.rpm) |
| `postgresql-17-pg-statecharts` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.1 KiB | [postgresql-17-pg-statecharts_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statecharts/postgresql-17-pg-statecharts_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-statecharts` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 25.1 KiB | [postgresql-17-pg-statecharts_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statecharts/postgresql-17-pg-statecharts_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-statecharts` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.1 KiB | [postgresql-17-pg-statecharts_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statecharts/postgresql-17-pg-statecharts_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-statecharts` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.1 KiB | [postgresql-17-pg-statecharts_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statecharts/postgresql-17-pg-statecharts_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-statecharts` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.7 KiB | [postgresql-17-pg-statecharts_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statecharts/postgresql-17-pg-statecharts_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-statecharts` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.7 KiB | [postgresql-17-pg-statecharts_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statecharts/postgresql-17-pg-statecharts_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-statecharts` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-17-pg-statecharts_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statecharts/postgresql-17-pg-statecharts_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-statecharts` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-17-pg-statecharts_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statecharts/postgresql-17-pg-statecharts_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-statecharts` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 25.5 KiB | [postgresql-17-pg-statecharts_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statecharts/postgresql-17-pg-statecharts_0.1.0-1PGSTY~resolute_all.deb) |
| `postgresql-17-pg-statecharts` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 25.5 KiB | [postgresql-17-pg-statecharts_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statecharts/postgresql-17-pg-statecharts_0.1.0-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statecharts_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.8 KiB | [pg_statecharts_16-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_statecharts_16-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_statecharts_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.8 KiB | [pg_statecharts_16-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_statecharts_16-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_statecharts_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.8 KiB | [pg_statecharts_16-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_statecharts_16-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_statecharts_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.8 KiB | [pg_statecharts_16-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_statecharts_16-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_statecharts_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.9 KiB | [pg_statecharts_16-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_statecharts_16-0.1.0-1PGSTY.el10.noarch.rpm) |
| `pg_statecharts_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.9 KiB | [pg_statecharts_16-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_statecharts_16-0.1.0-1PGSTY.el10.noarch.rpm) |
| `postgresql-16-pg-statecharts` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.1 KiB | [postgresql-16-pg-statecharts_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statecharts/postgresql-16-pg-statecharts_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-statecharts` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 25.1 KiB | [postgresql-16-pg-statecharts_0.1.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statecharts/postgresql-16-pg-statecharts_0.1.0-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-statecharts` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.1 KiB | [postgresql-16-pg-statecharts_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statecharts/postgresql-16-pg-statecharts_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-statecharts` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.1 KiB | [postgresql-16-pg-statecharts_0.1.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statecharts/postgresql-16-pg-statecharts_0.1.0-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-statecharts` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.7 KiB | [postgresql-16-pg-statecharts_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statecharts/postgresql-16-pg-statecharts_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-statecharts` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.7 KiB | [postgresql-16-pg-statecharts_0.1.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statecharts/postgresql-16-pg-statecharts_0.1.0-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-statecharts` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-16-pg-statecharts_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statecharts/postgresql-16-pg-statecharts_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-statecharts` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-16-pg-statecharts_0.1.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statecharts/postgresql-16-pg-statecharts_0.1.0-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-statecharts` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 25.5 KiB | [postgresql-16-pg-statecharts_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statecharts/postgresql-16-pg-statecharts_0.1.0-1PGSTY~resolute_all.deb) |
| `postgresql-16-pg-statecharts` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 25.5 KiB | [postgresql-16-pg-statecharts_0.1.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statecharts/postgresql-16-pg-statecharts_0.1.0-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/kronor-io/statecharts" title="Repository" icon="github" subtitle="github.com/kronor-io/statecharts" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_statecharts-0.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_statecharts;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_statecharts;		# install via package name, for the active PG version

pig install pg_statecharts -v 18;   # install for PG 18
pig install pg_statecharts -v 17;   # install for PG 17
pig install pg_statecharts -v 16;   # install for PG 16

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_statecharts CASCADE; -- requires ltree
```

## Usage

Sources:

- [PGXN 0.1.0 distribution](https://pgxn.org/dist/pg_statecharts/0.1.0/)
- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/README.md)
- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/pg_statecharts.control)
- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/sql/function/start_machine_with_latest_statechart.sql)
- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/sql/table/state_machine_event.sql)

`pg_statecharts` 0.1.0 runs hierarchical state machines and their events transactionally inside PostgreSQL. This release is pure SQL, tested on PostgreSQL 16–18, and needs no preload or restart.

### Core Workflow

Install `ltree` in `public`, then install the extension with a role that has the required object-creation privileges:

```sql
CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;
CREATE EXTENSION pg_statecharts;
```

Deploy a statechart definition through SQL migrations or the development companion before starting a machine. For an already deployed chart named order, this starts its latest version and returns the machine row:

```sql
SELECT * FROM fsm.start_machine_with_latest_statechart(1, 'order', '{}');
```

Events are inserted into `fsm.state_machine_event` with the shard, machine identifier, event name and JSON data. Its triggers process the event in the inserting transaction, including entry and exit callbacks. Callers must have the appropriate table and function privileges; chart callbacks are application code that needs its own review.

### Catalog and Versions

Definitions and running machines live in `fsm`. `fsm.to_semver` parses numeric chart versions and `fsm.semver_text` renders them; stored versions are integer arrays so their native ordering remains numeric. Do not sort a textual rendering when version precedence matters.

### Backups and Legacy Migration

The reviewed 0.1.0 source registers state tables and sequences for database dumps. The event inbox is intentionally omitted to avoid replaying callbacks during restoration. Backups made with 0.0.0 omitted machine state; take a fresh, restore-tested backup after upgrading.

An earlier 0.1.0 installation may predate dump registration, so check the installed catalog instead of trusting its version string. Do not drop a populated extension merely to refresh registration. Legacy Rust or Sqitch migrations require superuser work and the old semver dependency during conversion; follow the migration guide and rehearse on a restored copy. The old 0.0.0 path is no longer independently exercised upstream.
