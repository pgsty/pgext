---
title: "pg_mentat"
linkTitle: "pg_mentat"
description: "Datomic-compatible data model and Datalog query engine inside PostgreSQL"
weight: 2980
categories: ["FEAT"]
languages: ["Rust"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_mentat**](https://codeberg.org/gregburd/mentat) : Datomic-compatible data model and Datalog query engine inside PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2980** | {{< badge content="pg_mentat" link="https://codeberg.org/gregburd/mentat" >}} | {{< ext "pg_mentat" >}} | `1.10.3` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `mentat` |
|   **See Also**    | {{< ext "pg_fts" >}} {{< ext "pg_tre" >}} {{< ext "pg_infer" >}} {{< ext "rum" >}} {{< ext "pg_trgm" >}} {{< ext "fuzzystrmatch" >}} {{< ext "vector" >}} {{< ext "postgis" >}} |

> [!Note] No mentatd binary; integrations are optional. pgrx 0.19.3.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.10.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_mentat` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.10.3` | {{< bg "18" "pg_mentat_18" "green" >}} {{< bg "17" "pg_mentat_17" "green" >}} {{< bg "16" "pg_mentat_16" "green" >}} {{< bg "15" "pg_mentat_15" "green" >}} {{< bg "14" "pg_mentat_14" "green" >}} | `pg_mentat_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.10.3` | {{< bg "18" "postgresql-18-pg-mentat" "green" >}} {{< bg "17" "postgresql-17-pg-mentat" "green" >}} {{< bg "16" "postgresql-16-pg-mentat" "green" >}} {{< bg "15" "postgresql-15-pg-mentat" "green" >}} {{< bg "14" "postgresql-14-pg-mentat" "green" >}} | `postgresql-$v-pg-mentat` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "pg_mentat_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10.3" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mentat_18` | `1.10.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.8 MiB | [pg_mentat_18-1.10.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mentat_18-1.10.3-1PGSTY.el8.x86_64.rpm) |
| `pg_mentat_18` | `1.10.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [pg_mentat_18-1.10.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mentat_18-1.10.3-1PGSTY.el8.aarch64.rpm) |
| `pg_mentat_18` | `1.10.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [pg_mentat_18-1.10.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mentat_18-1.10.3-1PGSTY.el9.x86_64.rpm) |
| `pg_mentat_18` | `1.10.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [pg_mentat_18-1.10.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mentat_18-1.10.3-1PGSTY.el9.aarch64.rpm) |
| `pg_mentat_18` | `1.10.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [pg_mentat_18-1.10.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mentat_18-1.10.3-1PGSTY.el10.x86_64.rpm) |
| `pg_mentat_18` | `1.10.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [pg_mentat_18-1.10.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mentat_18-1.10.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-mentat` | `1.10.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.5 MiB | [postgresql-18-pg-mentat_1.10.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.10.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-mentat` | `1.10.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.3 MiB | [postgresql-18-pg-mentat_1.10.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.10.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-mentat` | `1.10.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.5 MiB | [postgresql-18-pg-mentat_1.10.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.10.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-mentat` | `1.10.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.3 MiB | [postgresql-18-pg-mentat_1.10.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.10.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-mentat` | `1.10.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-18-pg-mentat_1.10.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.10.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-mentat` | `1.10.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.5 MiB | [postgresql-18-pg-mentat_1.10.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.10.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-mentat` | `1.10.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.6 MiB | [postgresql-18-pg-mentat_1.10.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.10.3-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-mentat` | `1.10.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.5 MiB | [postgresql-18-pg-mentat_1.10.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.10.3-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-mentat` | `1.10.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.6 MiB | [postgresql-18-pg-mentat_1.10.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.10.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-mentat` | `1.10.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.5 MiB | [postgresql-18-pg-mentat_1.10.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.10.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mentat_17` | `1.10.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pg_mentat_17-1.10.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mentat_17-1.10.3-1PGSTY.el8.x86_64.rpm) |
| `pg_mentat_17` | `1.10.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [pg_mentat_17-1.10.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mentat_17-1.10.3-1PGSTY.el8.aarch64.rpm) |
| `pg_mentat_17` | `1.10.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [pg_mentat_17-1.10.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mentat_17-1.10.3-1PGSTY.el9.x86_64.rpm) |
| `pg_mentat_17` | `1.10.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [pg_mentat_17-1.10.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mentat_17-1.10.3-1PGSTY.el9.aarch64.rpm) |
| `pg_mentat_17` | `1.10.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [pg_mentat_17-1.10.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mentat_17-1.10.3-1PGSTY.el10.x86_64.rpm) |
| `pg_mentat_17` | `1.10.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [pg_mentat_17-1.10.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mentat_17-1.10.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-mentat` | `1.10.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.5 MiB | [postgresql-17-pg-mentat_1.10.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.10.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-mentat` | `1.10.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.3 MiB | [postgresql-17-pg-mentat_1.10.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.10.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-mentat` | `1.10.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.5 MiB | [postgresql-17-pg-mentat_1.10.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.10.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-mentat` | `1.10.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.3 MiB | [postgresql-17-pg-mentat_1.10.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.10.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-mentat` | `1.10.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-17-pg-mentat_1.10.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.10.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-mentat` | `1.10.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.5 MiB | [postgresql-17-pg-mentat_1.10.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.10.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-mentat` | `1.10.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.6 MiB | [postgresql-17-pg-mentat_1.10.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.10.3-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-mentat` | `1.10.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.5 MiB | [postgresql-17-pg-mentat_1.10.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.10.3-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-mentat` | `1.10.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.6 MiB | [postgresql-17-pg-mentat_1.10.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.10.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-mentat` | `1.10.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.5 MiB | [postgresql-17-pg-mentat_1.10.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.10.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mentat_16` | `1.10.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pg_mentat_16-1.10.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mentat_16-1.10.3-1PGSTY.el8.x86_64.rpm) |
| `pg_mentat_16` | `1.10.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [pg_mentat_16-1.10.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mentat_16-1.10.3-1PGSTY.el8.aarch64.rpm) |
| `pg_mentat_16` | `1.10.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [pg_mentat_16-1.10.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mentat_16-1.10.3-1PGSTY.el9.x86_64.rpm) |
| `pg_mentat_16` | `1.10.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [pg_mentat_16-1.10.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mentat_16-1.10.3-1PGSTY.el9.aarch64.rpm) |
| `pg_mentat_16` | `1.10.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [pg_mentat_16-1.10.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mentat_16-1.10.3-1PGSTY.el10.x86_64.rpm) |
| `pg_mentat_16` | `1.10.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [pg_mentat_16-1.10.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mentat_16-1.10.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-mentat` | `1.10.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.5 MiB | [postgresql-16-pg-mentat_1.10.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.10.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-mentat` | `1.10.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.3 MiB | [postgresql-16-pg-mentat_1.10.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.10.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-mentat` | `1.10.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.5 MiB | [postgresql-16-pg-mentat_1.10.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.10.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-mentat` | `1.10.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.3 MiB | [postgresql-16-pg-mentat_1.10.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.10.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-mentat` | `1.10.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-16-pg-mentat_1.10.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.10.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-mentat` | `1.10.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.5 MiB | [postgresql-16-pg-mentat_1.10.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.10.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-mentat` | `1.10.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.6 MiB | [postgresql-16-pg-mentat_1.10.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.10.3-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-mentat` | `1.10.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.5 MiB | [postgresql-16-pg-mentat_1.10.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.10.3-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-mentat` | `1.10.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.6 MiB | [postgresql-16-pg-mentat_1.10.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.10.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-mentat` | `1.10.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.5 MiB | [postgresql-16-pg-mentat_1.10.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.10.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mentat_15` | `1.10.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pg_mentat_15-1.10.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mentat_15-1.10.3-1PGSTY.el8.x86_64.rpm) |
| `pg_mentat_15` | `1.10.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [pg_mentat_15-1.10.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mentat_15-1.10.3-1PGSTY.el8.aarch64.rpm) |
| `pg_mentat_15` | `1.10.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [pg_mentat_15-1.10.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mentat_15-1.10.3-1PGSTY.el9.x86_64.rpm) |
| `pg_mentat_15` | `1.10.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [pg_mentat_15-1.10.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mentat_15-1.10.3-1PGSTY.el9.aarch64.rpm) |
| `pg_mentat_15` | `1.10.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [pg_mentat_15-1.10.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mentat_15-1.10.3-1PGSTY.el10.x86_64.rpm) |
| `pg_mentat_15` | `1.10.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [pg_mentat_15-1.10.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mentat_15-1.10.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-mentat` | `1.10.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.5 MiB | [postgresql-15-pg-mentat_1.10.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.10.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-mentat` | `1.10.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.3 MiB | [postgresql-15-pg-mentat_1.10.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.10.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-mentat` | `1.10.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.5 MiB | [postgresql-15-pg-mentat_1.10.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.10.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-mentat` | `1.10.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.3 MiB | [postgresql-15-pg-mentat_1.10.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.10.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-mentat` | `1.10.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-15-pg-mentat_1.10.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.10.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-mentat` | `1.10.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.5 MiB | [postgresql-15-pg-mentat_1.10.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.10.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-mentat` | `1.10.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.6 MiB | [postgresql-15-pg-mentat_1.10.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.10.3-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-mentat` | `1.10.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.5 MiB | [postgresql-15-pg-mentat_1.10.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.10.3-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-mentat` | `1.10.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.6 MiB | [postgresql-15-pg-mentat_1.10.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.10.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-mentat` | `1.10.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.5 MiB | [postgresql-15-pg-mentat_1.10.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.10.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mentat_14` | `1.10.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pg_mentat_14-1.10.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mentat_14-1.10.3-1PGSTY.el8.x86_64.rpm) |
| `pg_mentat_14` | `1.10.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [pg_mentat_14-1.10.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mentat_14-1.10.3-1PGSTY.el8.aarch64.rpm) |
| `pg_mentat_14` | `1.10.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [pg_mentat_14-1.10.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mentat_14-1.10.3-1PGSTY.el9.x86_64.rpm) |
| `pg_mentat_14` | `1.10.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [pg_mentat_14-1.10.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mentat_14-1.10.3-1PGSTY.el9.aarch64.rpm) |
| `pg_mentat_14` | `1.10.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [pg_mentat_14-1.10.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mentat_14-1.10.3-1PGSTY.el10.x86_64.rpm) |
| `pg_mentat_14` | `1.10.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [pg_mentat_14-1.10.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mentat_14-1.10.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-mentat` | `1.10.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-14-pg-mentat_1.10.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.10.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-mentat` | `1.10.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.3 MiB | [postgresql-14-pg-mentat_1.10.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.10.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-mentat` | `1.10.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-14-pg-mentat_1.10.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.10.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-mentat` | `1.10.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.3 MiB | [postgresql-14-pg-mentat_1.10.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.10.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-mentat` | `1.10.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-14-pg-mentat_1.10.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.10.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-mentat` | `1.10.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.5 MiB | [postgresql-14-pg-mentat_1.10.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.10.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-mentat` | `1.10.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.6 MiB | [postgresql-14-pg-mentat_1.10.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.10.3-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-mentat` | `1.10.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.5 MiB | [postgresql-14-pg-mentat_1.10.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.10.3-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-mentat` | `1.10.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.6 MiB | [postgresql-14-pg-mentat_1.10.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.10.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-mentat` | `1.10.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.5 MiB | [postgresql-14-pg-mentat_1.10.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.10.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/gregburd/mentat" title="Repository" icon="link" subtitle="codeberg.org/gregburd/mentat" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_mentat-1.10.3.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_mentat;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_mentat;		# install via package name, for the active PG version

pig install pg_mentat -v 18;   # install for PG 18
pig install pg_mentat -v 17;   # install for PG 17
pig install pg_mentat -v 16;   # install for PG 16
pig install pg_mentat -v 15;   # install for PG 15
pig install pg_mentat -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_mentat;
```

## Usage

Sources:

- [v1.10.1 README](https://codeberg.org/gregburd/mentat/src/tag/v1.10.1/README.md)
- [v1.10.1 extension control](https://codeberg.org/gregburd/mentat/src/tag/v1.10.1/crates/pg/pg_mentat/pg_mentat.control)
- [v1.10.1 changelog](https://codeberg.org/gregburd/mentat/src/tag/v1.10.1/CHANGELOG.md)
- [v1.9.1 dump repair](https://codeberg.org/gregburd/mentat/src/tag/v1.10.1/crates/pg/pg_mentat/sql/pg_mentat--1.9.0--1.9.1.sql)
- [v1.10.1 SQL aliases](https://codeberg.org/gregburd/mentat/src/tag/v1.10.1/crates/pg/pg_mentat/sql/07_function_aliases.sql)
- [v1.10.1 history functions](https://codeberg.org/gregburd/mentat/src/tag/v1.10.1/crates/pg/pg_mentat/src/functions/time_travel.rs)
- [v1.10.1 excision function](https://codeberg.org/gregburd/mentat/src/tag/v1.10.1/crates/pg/pg_mentat/src/functions/excision.rs)
- [v1.10.1 transaction functions](https://codeberg.org/gregburd/mentat/src/tag/v1.10.1/crates/pg/pg_mentat/src/functions/transact.rs)
- [v1.10.1 subscriptions](https://codeberg.org/gregburd/mentat/src/tag/v1.10.1/crates/pg/pg_mentat/src/functions/subscriptions.rs)

`pg_mentat` implements a Datomic-compatible data model and Datalog query engine inside PostgreSQL. It stores immutable facts as typed datoms and exposes schema transactions, Datalog queries, pull expressions, time travel, transaction history, and permanent excision through SQL functions. Use it for applications that need this model; it is not a transparent replacement for relational tables or SQL.

### Install and Define a Schema

```sql
CREATE EXTENSION pg_mentat;

SELECT mentat.t('[
  {:db/ident       :person/name
   :db/valueType   :db.type/string
   :db/cardinality :db.cardinality/one}
  {:db/ident       :person/age
   :db/valueType   :db.type/long
   :db/cardinality :db.cardinality/one}
]');
```

The recommended convenience aliases live in schema `mentat`. Schema must be transacted before facts use the new attributes.

### Transact and Query Data

```sql
SELECT mentat.t('[
  {:person/name "Alice" :person/age 30}
  {:person/name "Bob"   :person/age 25}
]');

SELECT mentat.q('
  [:find ?name ?age
   :where [?e :person/name ?name]
          [?e :person/age ?age]
          [(> ?age 28)]]
');
```

`mentat.t(edn)` applies an ACID transaction and returns its transaction report. `mentat.q(query, inputs)` compiles a Datalog query to PostgreSQL execution. Use EDN parameters and input bindings rather than interpolating application strings into a query.

### Pull, History, and What-If Transactions

```sql
SELECT mentat.pull('[*]', 10001);
SELECT public.log('default', 1000001, 1000010);
SELECT public.diff(
  'default', 1000003, 1000007,
  '[:find ?name :where [?e :person/name ?name]]',
  '{}'::jsonb
);

SELECT public.mentat_with('[
  {:person/name "Alice" :person/age 31}
]');
```

`mentat.pull` returns entity-shaped JSON. `public.log` returns history for a transaction interval, while `public.diff` compares the results of a supplied Datalog query at two transaction points and requires all five arguments. `public.mentat_with` evaluates a transaction without persisting it. These secondary functions are exported in the public schema in v1.10.1; they have no mentat-schema aliases. Use entity and transaction IDs from your own store. Queries can also be evaluated as of or since a transaction by using the documented database arguments.

Permanent excision is intentionally separate from normal immutable history:

```sql
SELECT public.mentat_excise(ARRAY[10042]::bigint[], 'default', NULL);
```

Review the target entities and backups before excision; it permanently removes datoms and is intended for requirements such as privacy erasure. The first argument is an array of entity IDs. The containing partition must allow excision; schema entities are protected, and references from entities outside the same excision batch can block the operation.

### Important Objects

- `mentat.t(edn)`: transact schema or data.
- `mentat.q(query, inputs)`: execute Datalog.
- `mentat.pull(pattern, eid)` and `mentat.pull_many(pattern, eids)`: entity-shaped reads.
- `mentat.entity(eid)` and `mentat.schema()`: inspect an entity or current schema.
- `public.log(...)` and `public.diff(...)`: inspect transaction history and query-result changes.
- `mentat.stats()`, `mentat.storage()`, and `mentat.cache_stats()`: operational inspection.
- `public.subscribe(...)`: reactive query notifications through PostgreSQL `LISTEN`/`NOTIFY`.

The extension stores typed datoms in narrow tables under schema `mentat`, including reference, integer, string, boolean, floating-point, instant, keyword, UUID, and byte values.

### Earlier Security Fixes

Version 1.6.2 fixes deeply nested EDN causing stack exhaustion, non-superuser query failures when setting restricted limits, and incorrect UTC decoding. The optional `script` build feature adds Clojure-style evaluation; it is disabled by default. Upstream reports that all earlier releases, including 1.5.7, are affected by the EDN nesting flaw: a backend crash can disconnect all clients and trigger crash recovery. Verify the installed extension version when upgrading an existing deployment.

### Requirements and Caveats

- Upstream v1.10.1 supports PostgreSQL 13-18. Current Pigsty packages target PostgreSQL 14-18 and are rebuilt with pgrx 0.19.2; upstream's tagged source declares pgrx 0.17. Treat the packaged binary as the compatibility boundary.
- The extension is not relocatable and does not require `shared_preload_libraries`.
- The optional `mentatd` HTTP/Datomic-wire daemon is an upstream companion program and is not included in the Pigsty `pg_mentat` package. SQL use of the extension does not require it.
- Datalog compilation, pull recursion, full-text attributes, subscriptions, and history can have very different cost profiles. Inspect generated SQL with the documented explain helper and benchmark representative data.
- Excision bypasses the normal immutable-history model. Restrict privileges and audit its use.

### 1.10.1 APIs and Indexes

The project moved to gregburd/mentat. The new core names are `edn_t`, `edn_q`, `edn_pull`, and optional scripting entry point `edn_eval`; convenience aliases `mentat.t` and `mentat.q` remain available, while older `mentat_*` names are deprecated. `edn_q_rows` returns one JSONB array per row but still materializes the complete result within a call.

`mentat.auto_index` accepts `off`, `schema` (default), or `adaptive`. Current-value tables receive AVET indexes; adaptive mode manages partial history-attribute indexes recorded in `mentat.managed_indexes`. `mentat_tune_indexes` reports a plan by default and applies it only when dry-run is explicitly disabled. It drops only its managed indexes, with a default idle window of seven days.

Aggregates now follow Datalog set semantics: summing 1, 1, 1, 3 yields 4 instead of 6. Add `:with ?e` when entity multiplicity must be preserved, and review queries that rely on the old result before upgrading.

### Upgrade and Logical Backups

1.9.1 fixes an important backup defect: earlier versions did not register extension-member table data, so logical `pg_dump` backups omitted mentat data and restored empty stores. Take and verify a new logical backup immediately after upgrading; older backups are not repaired retroactively. Physical backups are unaffected by this defect.

After installing the new files, update through the available upgrade chain to 1.10.1. The upgrade builds indexes and can block writes; schedule a maintenance window:

```sql
ALTER EXTENSION pg_mentat UPDATE TO '1.10.1';
```

Deployed databases still need the SQL extension update and a verified new logical backup.
