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

[**pg_mentat**](https://codeberg.org/gregburd/pg_mentat) : Datomic-compatible data model and Datalog query engine inside PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2980** | {{< badge content="pg_mentat" link="https://codeberg.org/gregburd/pg_mentat" >}} | {{< ext "pg_mentat" >}} | `1.5.7` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `mentat` |
|   **See Also**    | {{< ext "pg_fts" >}} {{< ext "pg_tre" >}} {{< ext "pg_infer" >}} {{< ext "rum" >}} {{< ext "pg_trgm" >}} {{< ext "fuzzystrmatch" >}} {{< ext "vector" >}} {{< ext "postgis" >}} |

> [!Note] The PIGSTY package omits optional mentatd and installs no user-facing binary; listed integrations are soft dependencies. Effective build uses pgrx 0.19.1, migrated from upstream 0.17.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_mentat` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.7` | {{< bg "18" "pg_mentat_18" "green" >}} {{< bg "17" "pg_mentat_17" "green" >}} {{< bg "16" "pg_mentat_16" "green" >}} {{< bg "15" "pg_mentat_15" "green" >}} {{< bg "14" "pg_mentat_14" "green" >}} | `pg_mentat_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.7` | {{< bg "18" "postgresql-18-pg-mentat" "green" >}} {{< bg "17" "postgresql-17-pg-mentat" "green" >}} {{< bg "16" "postgresql-16-pg-mentat" "green" >}} {{< bg "15" "postgresql-15-pg-mentat" "green" >}} {{< bg "14" "postgresql-14-pg-mentat" "green" >}} | `postgresql-$v-pg-mentat` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "pg_mentat_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-18-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-17-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-16-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-15-pg-mentat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.7" "postgresql-14-pg-mentat : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mentat_18` | `1.5.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.6 MiB | [pg_mentat_18-1.5.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mentat_18-1.5.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_mentat_18` | `1.5.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [pg_mentat_18-1.5.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mentat_18-1.5.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_mentat_18` | `1.5.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.6 MiB | [pg_mentat_18-1.5.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mentat_18-1.5.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_mentat_18` | `1.5.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pg_mentat_18-1.5.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mentat_18-1.5.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_mentat_18` | `1.5.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.6 MiB | [pg_mentat_18-1.5.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mentat_18-1.5.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_mentat_18` | `1.5.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [pg_mentat_18-1.5.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mentat_18-1.5.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-mentat` | `1.5.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.1 MiB | [postgresql-18-pg-mentat_1.5.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.5.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-mentat` | `1.5.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-18-pg-mentat_1.5.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.5.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-mentat` | `1.5.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.1 MiB | [postgresql-18-pg-mentat_1.5.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.5.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-mentat` | `1.5.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-18-pg-mentat_1.5.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.5.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-mentat` | `1.5.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.4 MiB | [postgresql-18-pg-mentat_1.5.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.5.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-mentat` | `1.5.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 MiB | [postgresql-18-pg-mentat_1.5.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.5.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-mentat` | `1.5.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.3 MiB | [postgresql-18-pg-mentat_1.5.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.5.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-mentat` | `1.5.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-18-pg-mentat_1.5.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.5.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-mentat` | `1.5.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.3 MiB | [postgresql-18-pg-mentat_1.5.7-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.5.7-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-mentat` | `1.5.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.3 MiB | [postgresql-18-pg-mentat_1.5.7-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-18-pg-mentat_1.5.7-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mentat_17` | `1.5.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.6 MiB | [pg_mentat_17-1.5.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mentat_17-1.5.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_mentat_17` | `1.5.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [pg_mentat_17-1.5.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mentat_17-1.5.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_mentat_17` | `1.5.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.6 MiB | [pg_mentat_17-1.5.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mentat_17-1.5.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_mentat_17` | `1.5.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pg_mentat_17-1.5.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mentat_17-1.5.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_mentat_17` | `1.5.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.6 MiB | [pg_mentat_17-1.5.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mentat_17-1.5.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_mentat_17` | `1.5.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [pg_mentat_17-1.5.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mentat_17-1.5.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-mentat` | `1.5.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.1 MiB | [postgresql-17-pg-mentat_1.5.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.5.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-mentat` | `1.5.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-17-pg-mentat_1.5.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.5.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-mentat` | `1.5.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.1 MiB | [postgresql-17-pg-mentat_1.5.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.5.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-mentat` | `1.5.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-17-pg-mentat_1.5.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.5.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-mentat` | `1.5.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.4 MiB | [postgresql-17-pg-mentat_1.5.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.5.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-mentat` | `1.5.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-17-pg-mentat_1.5.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.5.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-mentat` | `1.5.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.3 MiB | [postgresql-17-pg-mentat_1.5.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.5.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-mentat` | `1.5.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-17-pg-mentat_1.5.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.5.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-mentat` | `1.5.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.3 MiB | [postgresql-17-pg-mentat_1.5.7-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.5.7-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-mentat` | `1.5.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.3 MiB | [postgresql-17-pg-mentat_1.5.7-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-17-pg-mentat_1.5.7-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mentat_16` | `1.5.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.6 MiB | [pg_mentat_16-1.5.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mentat_16-1.5.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_mentat_16` | `1.5.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [pg_mentat_16-1.5.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mentat_16-1.5.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_mentat_16` | `1.5.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.6 MiB | [pg_mentat_16-1.5.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mentat_16-1.5.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_mentat_16` | `1.5.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pg_mentat_16-1.5.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mentat_16-1.5.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_mentat_16` | `1.5.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.6 MiB | [pg_mentat_16-1.5.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mentat_16-1.5.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_mentat_16` | `1.5.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [pg_mentat_16-1.5.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mentat_16-1.5.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-mentat` | `1.5.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.1 MiB | [postgresql-16-pg-mentat_1.5.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.5.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-mentat` | `1.5.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-16-pg-mentat_1.5.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.5.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-mentat` | `1.5.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.1 MiB | [postgresql-16-pg-mentat_1.5.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.5.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-mentat` | `1.5.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-16-pg-mentat_1.5.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.5.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-mentat` | `1.5.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.4 MiB | [postgresql-16-pg-mentat_1.5.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.5.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-mentat` | `1.5.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-16-pg-mentat_1.5.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.5.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-mentat` | `1.5.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.3 MiB | [postgresql-16-pg-mentat_1.5.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.5.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-mentat` | `1.5.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-16-pg-mentat_1.5.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.5.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-mentat` | `1.5.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.3 MiB | [postgresql-16-pg-mentat_1.5.7-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.5.7-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-mentat` | `1.5.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.3 MiB | [postgresql-16-pg-mentat_1.5.7-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-16-pg-mentat_1.5.7-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mentat_15` | `1.5.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.6 MiB | [pg_mentat_15-1.5.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mentat_15-1.5.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_mentat_15` | `1.5.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [pg_mentat_15-1.5.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mentat_15-1.5.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_mentat_15` | `1.5.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.6 MiB | [pg_mentat_15-1.5.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mentat_15-1.5.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_mentat_15` | `1.5.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pg_mentat_15-1.5.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mentat_15-1.5.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_mentat_15` | `1.5.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.6 MiB | [pg_mentat_15-1.5.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mentat_15-1.5.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_mentat_15` | `1.5.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [pg_mentat_15-1.5.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mentat_15-1.5.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-mentat` | `1.5.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.1 MiB | [postgresql-15-pg-mentat_1.5.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.5.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-mentat` | `1.5.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.1 MiB | [postgresql-15-pg-mentat_1.5.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.5.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-mentat` | `1.5.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.1 MiB | [postgresql-15-pg-mentat_1.5.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.5.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-mentat` | `1.5.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.1 MiB | [postgresql-15-pg-mentat_1.5.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.5.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-mentat` | `1.5.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.3 MiB | [postgresql-15-pg-mentat_1.5.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.5.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-mentat` | `1.5.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-15-pg-mentat_1.5.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.5.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-mentat` | `1.5.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.3 MiB | [postgresql-15-pg-mentat_1.5.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.5.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-mentat` | `1.5.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-15-pg-mentat_1.5.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.5.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-mentat` | `1.5.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.3 MiB | [postgresql-15-pg-mentat_1.5.7-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.5.7-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-mentat` | `1.5.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.3 MiB | [postgresql-15-pg-mentat_1.5.7-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-15-pg-mentat_1.5.7-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mentat_14` | `1.5.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.6 MiB | [pg_mentat_14-1.5.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mentat_14-1.5.7-1PIGSTY.el8.x86_64.rpm) |
| `pg_mentat_14` | `1.5.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [pg_mentat_14-1.5.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mentat_14-1.5.7-1PIGSTY.el8.aarch64.rpm) |
| `pg_mentat_14` | `1.5.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.6 MiB | [pg_mentat_14-1.5.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mentat_14-1.5.7-1PIGSTY.el9.x86_64.rpm) |
| `pg_mentat_14` | `1.5.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pg_mentat_14-1.5.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mentat_14-1.5.7-1PIGSTY.el9.aarch64.rpm) |
| `pg_mentat_14` | `1.5.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.6 MiB | [pg_mentat_14-1.5.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mentat_14-1.5.7-1PIGSTY.el10.x86_64.rpm) |
| `pg_mentat_14` | `1.5.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [pg_mentat_14-1.5.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mentat_14-1.5.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-mentat` | `1.5.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.1 MiB | [postgresql-14-pg-mentat_1.5.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.5.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-mentat` | `1.5.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.1 MiB | [postgresql-14-pg-mentat_1.5.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.5.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-mentat` | `1.5.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.1 MiB | [postgresql-14-pg-mentat_1.5.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.5.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-mentat` | `1.5.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.1 MiB | [postgresql-14-pg-mentat_1.5.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.5.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-mentat` | `1.5.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.3 MiB | [postgresql-14-pg-mentat_1.5.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.5.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-mentat` | `1.5.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-14-pg-mentat_1.5.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.5.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-mentat` | `1.5.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.3 MiB | [postgresql-14-pg-mentat_1.5.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.5.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-mentat` | `1.5.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-14-pg-mentat_1.5.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.5.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-mentat` | `1.5.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.3 MiB | [postgresql-14-pg-mentat_1.5.7-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.5.7-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-mentat` | `1.5.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.3 MiB | [postgresql-14-pg-mentat_1.5.7-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mentat/postgresql-14-pg-mentat_1.5.7-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/gregburd/pg_mentat" title="Repository" icon="link" subtitle="codeberg.org/gregburd/pg_mentat" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_mentat-1.5.7.tar.gz" />}}
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

- [pg_mentat v1.5.7 README](https://codeberg.org/gregburd/pg_mentat/src/tag/v1.5.7/README.md)
- [pg_mentat v1.5.7 control file](https://codeberg.org/gregburd/pg_mentat/src/tag/v1.5.7/pg_mentat/pg_mentat.control)
- [pg_mentat v1.5.6 to v1.5.7 upgrade SQL](https://codeberg.org/gregburd/pg_mentat/src/tag/v1.5.7/pg_mentat/sql/pg_mentat--1.5.6--1.5.7.sql)
- [Pigsty package matrix](https://pgext.cloud/ext/pg_mentat)

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
SELECT mentat.log('default', 1000001, 1000010);
SELECT mentat.diff('default', 1000003, 1000007);

SELECT mentat.mentat_with('[
  {:person/name "Alice" :person/age 31}
]');
```

`mentat.pull` returns entity-shaped JSON. `mentat.log` and `mentat.diff` expose transaction history, and `mentat.mentat_with` evaluates a transaction without persisting it. Queries can also be evaluated as of or since a transaction by using the documented database arguments.

Permanent excision is intentionally separate from normal immutable history:

```sql
SELECT mentat.mentat_excise('default', 10042, NULL);
```

Review the target entity and backups before excision; it permanently removes datoms and is intended for requirements such as privacy erasure.

### Important Objects

- `mentat.t(edn)`: transact schema or data.
- `mentat.q(query, inputs)`: execute Datalog.
- `mentat.pull(pattern, eid)` and `mentat.pull_many(pattern, eids)`: entity-shaped reads.
- `mentat.entity(eid)` and `mentat.schema()`: inspect an entity or current schema.
- `mentat.log(...)` and `mentat.diff(...)`: inspect transaction history.
- `mentat.stats()`, `mentat.storage()`, and `mentat.cache_stats()`: operational inspection.
- `mentat.subscribe(...)`: reactive query notifications through PostgreSQL `LISTEN`/`NOTIFY`.

The extension stores typed datoms in narrow tables under schema `mentat`, including reference, integer, string, boolean, floating-point, instant, keyword, UUID, and byte values.

### Requirements and Caveats

- Upstream v1.5.7 supports PostgreSQL 13-18. Current Pigsty packages target PostgreSQL 14-18 and are rebuilt with pgrx 0.19.1; upstream's tagged source declares pgrx 0.17. Treat the packaged binary as the compatibility boundary.
- The extension is not relocatable and does not require `shared_preload_libraries`.
- The optional `mentatd` HTTP/Datomic-wire daemon is an upstream companion program and is not included in the Pigsty `pg_mentat` package. SQL use of the extension does not require it.
- Datalog compilation, pull recursion, full-text attributes, subscriptions, and history can have very different cost profiles. Inspect generated SQL with the documented explain helper and benchmark representative data.
- Excision bypasses the normal immutable-history model. Restrict privileges and audit its use.
