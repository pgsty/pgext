---
title: "pg_mentat"
linkTitle: "pg_mentat"
description: "Datomic-compatible data model and Datalog query engine inside PostgreSQL"
weight: 2980
categories: ["FEAT"]
width: full
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


{{< tabs >}}
{{< tab name="PG18" >}}

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

{{< /tab >}}
{{< tab name="PG17" >}}

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

{{< /tab >}}
{{< tab name="PG16" >}}

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

{{< /tab >}}
{{< tab name="PG15" >}}

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

{{< /tab >}}
{{< tab name="PG14" >}}

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

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/gregburd/pg_mentat" title="Repository" icon="link" subtitle="codeberg.org/gregburd/pg_mentat" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_mentat-1.5.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_mentat;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

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

- [Official upstream README](https://github.com/gburd/pg_mentat/blob/134015ebd3121e1a74eeff2de6a800143b33cb4e/pg_mentat/README.md)
- [Official extension control file (pg_mentat.control)](https://github.com/gburd/pg_mentat/blob/134015ebd3121e1a74eeff2de6a800143b33cb4e/pg_mentat/pg_mentat.control)
- [Official extension SQL (pg_mentat--1.0.0.sql)](https://github.com/gburd/pg_mentat/blob/134015ebd3121e1a74eeff2de6a800143b33cb4e/pg_mentat/sql/pg_mentat--1.0.0.sql)

`pg_mentat` — PostgreSQL extension providing a Datomic-compatible Datalog query engine with a native EDN data type. Use it when porting or emulating the corresponding database API. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_mentat;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `mentat.allocate_entid(partition_name TEXT)` is an extension function and returns `BIGINT`.
- `mentat.fulltext_update_trigger()` is an extension function and returns `trigger`.
- `mentat.resolve_ident(keyword TEXT)` is an extension function and returns `BIGINT`.
- `mentat.cardinality_type` is an extension-defined type.
- `mentat.EdnValue` is an extension-defined type.
- `mentat.unique_type` is an extension-defined type.
- `mentat.value_type` is an extension-defined type.
- `mentat.datoms` is a table installed or managed by the extension.
- `mentat.datoms_bool` is a table installed or managed by the extension.
- `mentat.datoms_bytes` is a table installed or managed by the extension.
- `mentat.datoms_default` is a table installed or managed by the extension.
- `mentat.datoms_double` is a table installed or managed by the extension.
- `mentat.datoms_instant` is a table installed or managed by the extension.
- `mentat.datoms_keyword` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.5.7`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
