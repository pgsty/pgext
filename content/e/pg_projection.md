---
title: "pg_projection"
linkTitle: "pg_projection"
description: "MongoDB-like read projections for JSONB in PostgreSQL"
weight: 9090
categories: ["SIM"]
width: full
---

[**pg_projection**](https://github.com/suissa/pg_projection) : MongoDB-like read projections for JSONB in PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9090** | {{< badge content="pg_projection" link="https://github.com/suissa/pg_projection" >}} | {{< ext "pg_projection" >}} | `1.0.0` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pgjq" >}} |

> [!Note] SQL-only extension.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_projection` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_projection_18" "green" >}} {{< bg "17" "pg_projection_17" "green" >}} {{< bg "16" "pg_projection_16" "green" >}} {{< bg "15" "pg_projection_15" "green" >}} {{< bg "14" "pg_projection_14" "green" >}} | `pg_projection_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pg-projection" "green" >}} {{< bg "17" "postgresql-17-pg-projection" "green" >}} {{< bg "16" "postgresql-16-pg-projection" "green" >}} {{< bg "15" "postgresql-15-pg-projection" "green" >}} {{< bg "14" "postgresql-14-pg-projection" "green" >}} | `postgresql-$v-pg-projection` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_projection_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-projection : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-projection : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-projection : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-projection : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-projection : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-projection : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-projection : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-projection : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-projection : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-projection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-projection : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_projection_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.8 KiB | [pg_projection_18-1.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_projection_18-1.0.0-1PIGSTY.el8.noarch.rpm) |
| `pg_projection_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.8 KiB | [pg_projection_18-1.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_projection_18-1.0.0-1PIGSTY.el8.noarch.rpm) |
| `pg_projection_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.9 KiB | [pg_projection_18-1.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_projection_18-1.0.0-1PIGSTY.el9.noarch.rpm) |
| `pg_projection_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.9 KiB | [pg_projection_18-1.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_projection_18-1.0.0-1PIGSTY.el9.noarch.rpm) |
| `pg_projection_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.0 KiB | [pg_projection_18-1.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_projection_18-1.0.0-1PIGSTY.el10.noarch.rpm) |
| `pg_projection_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.0 KiB | [pg_projection_18-1.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_projection_18-1.0.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-pg-projection` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.9 KiB | [postgresql-18-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-projection/postgresql-18-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-projection` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.9 KiB | [postgresql-18-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-projection/postgresql-18-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-projection` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.9 KiB | [postgresql-18-pg-projection_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-projection/postgresql-18-pg-projection_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-projection` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.9 KiB | [postgresql-18-pg-projection_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-projection/postgresql-18-pg-projection_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-projection` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.8 KiB | [postgresql-18-pg-projection_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-projection/postgresql-18-pg-projection_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-projection` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.8 KiB | [postgresql-18-pg-projection_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-projection/postgresql-18-pg-projection_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-projection` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.8 KiB | [postgresql-18-pg-projection_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-projection/postgresql-18-pg-projection_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-projection` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.8 KiB | [postgresql-18-pg-projection_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-projection/postgresql-18-pg-projection_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-projection` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.8 KiB | [postgresql-18-pg-projection_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-projection/postgresql-18-pg-projection_1.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-18-pg-projection` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.8 KiB | [postgresql-18-pg-projection_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-projection/postgresql-18-pg-projection_1.0.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_projection_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.8 KiB | [pg_projection_17-1.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_projection_17-1.0.0-1PIGSTY.el8.noarch.rpm) |
| `pg_projection_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.8 KiB | [pg_projection_17-1.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_projection_17-1.0.0-1PIGSTY.el8.noarch.rpm) |
| `pg_projection_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.9 KiB | [pg_projection_17-1.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_projection_17-1.0.0-1PIGSTY.el9.noarch.rpm) |
| `pg_projection_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.9 KiB | [pg_projection_17-1.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_projection_17-1.0.0-1PIGSTY.el9.noarch.rpm) |
| `pg_projection_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.0 KiB | [pg_projection_17-1.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_projection_17-1.0.0-1PIGSTY.el10.noarch.rpm) |
| `pg_projection_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.0 KiB | [pg_projection_17-1.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_projection_17-1.0.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-pg-projection` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.9 KiB | [postgresql-17-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-projection/postgresql-17-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-projection` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.9 KiB | [postgresql-17-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-projection/postgresql-17-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-projection` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.9 KiB | [postgresql-17-pg-projection_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-projection/postgresql-17-pg-projection_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-projection` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.9 KiB | [postgresql-17-pg-projection_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-projection/postgresql-17-pg-projection_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-projection` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.8 KiB | [postgresql-17-pg-projection_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-projection/postgresql-17-pg-projection_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-projection` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.8 KiB | [postgresql-17-pg-projection_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-projection/postgresql-17-pg-projection_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-projection` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.8 KiB | [postgresql-17-pg-projection_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-projection/postgresql-17-pg-projection_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-projection` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.8 KiB | [postgresql-17-pg-projection_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-projection/postgresql-17-pg-projection_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-projection` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.8 KiB | [postgresql-17-pg-projection_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-projection/postgresql-17-pg-projection_1.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-17-pg-projection` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.8 KiB | [postgresql-17-pg-projection_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-projection/postgresql-17-pg-projection_1.0.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_projection_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.8 KiB | [pg_projection_16-1.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_projection_16-1.0.0-1PIGSTY.el8.noarch.rpm) |
| `pg_projection_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.8 KiB | [pg_projection_16-1.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_projection_16-1.0.0-1PIGSTY.el8.noarch.rpm) |
| `pg_projection_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.9 KiB | [pg_projection_16-1.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_projection_16-1.0.0-1PIGSTY.el9.noarch.rpm) |
| `pg_projection_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.9 KiB | [pg_projection_16-1.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_projection_16-1.0.0-1PIGSTY.el9.noarch.rpm) |
| `pg_projection_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.0 KiB | [pg_projection_16-1.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_projection_16-1.0.0-1PIGSTY.el10.noarch.rpm) |
| `pg_projection_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.0 KiB | [pg_projection_16-1.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_projection_16-1.0.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-pg-projection` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.9 KiB | [postgresql-16-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-projection/postgresql-16-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-projection` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.9 KiB | [postgresql-16-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-projection/postgresql-16-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-projection` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.9 KiB | [postgresql-16-pg-projection_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-projection/postgresql-16-pg-projection_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-projection` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.9 KiB | [postgresql-16-pg-projection_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-projection/postgresql-16-pg-projection_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-projection` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.8 KiB | [postgresql-16-pg-projection_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-projection/postgresql-16-pg-projection_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-projection` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.8 KiB | [postgresql-16-pg-projection_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-projection/postgresql-16-pg-projection_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-projection` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.8 KiB | [postgresql-16-pg-projection_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-projection/postgresql-16-pg-projection_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-projection` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.8 KiB | [postgresql-16-pg-projection_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-projection/postgresql-16-pg-projection_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-projection` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.8 KiB | [postgresql-16-pg-projection_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-projection/postgresql-16-pg-projection_1.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-16-pg-projection` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.8 KiB | [postgresql-16-pg-projection_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-projection/postgresql-16-pg-projection_1.0.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_projection_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.8 KiB | [pg_projection_15-1.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_projection_15-1.0.0-1PIGSTY.el8.noarch.rpm) |
| `pg_projection_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.8 KiB | [pg_projection_15-1.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_projection_15-1.0.0-1PIGSTY.el8.noarch.rpm) |
| `pg_projection_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.9 KiB | [pg_projection_15-1.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_projection_15-1.0.0-1PIGSTY.el9.noarch.rpm) |
| `pg_projection_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.8 KiB | [pg_projection_15-1.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_projection_15-1.0.0-1PIGSTY.el9.noarch.rpm) |
| `pg_projection_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.0 KiB | [pg_projection_15-1.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_projection_15-1.0.0-1PIGSTY.el10.noarch.rpm) |
| `pg_projection_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.0 KiB | [pg_projection_15-1.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_projection_15-1.0.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-pg-projection` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.9 KiB | [postgresql-15-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-projection/postgresql-15-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-projection` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.9 KiB | [postgresql-15-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-projection/postgresql-15-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-projection` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.9 KiB | [postgresql-15-pg-projection_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-projection/postgresql-15-pg-projection_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-projection` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.9 KiB | [postgresql-15-pg-projection_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-projection/postgresql-15-pg-projection_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-projection` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.8 KiB | [postgresql-15-pg-projection_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-projection/postgresql-15-pg-projection_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-projection` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.8 KiB | [postgresql-15-pg-projection_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-projection/postgresql-15-pg-projection_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-projection` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.8 KiB | [postgresql-15-pg-projection_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-projection/postgresql-15-pg-projection_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-projection` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.8 KiB | [postgresql-15-pg-projection_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-projection/postgresql-15-pg-projection_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-projection` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.8 KiB | [postgresql-15-pg-projection_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-projection/postgresql-15-pg-projection_1.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-15-pg-projection` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.8 KiB | [postgresql-15-pg-projection_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-projection/postgresql-15-pg-projection_1.0.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_projection_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.8 KiB | [pg_projection_14-1.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_projection_14-1.0.0-1PIGSTY.el8.noarch.rpm) |
| `pg_projection_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.8 KiB | [pg_projection_14-1.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_projection_14-1.0.0-1PIGSTY.el8.noarch.rpm) |
| `pg_projection_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.9 KiB | [pg_projection_14-1.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_projection_14-1.0.0-1PIGSTY.el9.noarch.rpm) |
| `pg_projection_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.9 KiB | [pg_projection_14-1.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_projection_14-1.0.0-1PIGSTY.el9.noarch.rpm) |
| `pg_projection_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.0 KiB | [pg_projection_14-1.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_projection_14-1.0.0-1PIGSTY.el10.noarch.rpm) |
| `pg_projection_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.0 KiB | [pg_projection_14-1.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_projection_14-1.0.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-pg-projection` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.9 KiB | [postgresql-14-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-projection/postgresql-14-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pg-projection` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.9 KiB | [postgresql-14-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-projection/postgresql-14-pg-projection_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pg-projection` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.9 KiB | [postgresql-14-pg-projection_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-projection/postgresql-14-pg-projection_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pg-projection` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.9 KiB | [postgresql-14-pg-projection_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-projection/postgresql-14-pg-projection_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pg-projection` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.8 KiB | [postgresql-14-pg-projection_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-projection/postgresql-14-pg-projection_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pg-projection` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.8 KiB | [postgresql-14-pg-projection_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-projection/postgresql-14-pg-projection_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pg-projection` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.8 KiB | [postgresql-14-pg-projection_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-projection/postgresql-14-pg-projection_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-pg-projection` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.8 KiB | [postgresql-14-pg-projection_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-projection/postgresql-14-pg-projection_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-pg-projection` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.8 KiB | [postgresql-14-pg-projection_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-projection/postgresql-14-pg-projection_1.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-14-pg-projection` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.8 KiB | [postgresql-14-pg-projection_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-projection/postgresql-14-pg-projection_1.0.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/suissa/pg_projection" title="Repository" icon="github" subtitle="github.com/suissa/pg_projection" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_projection-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_projection;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_projection;		# install via package name, for the active PG version

pig install pg_projection -v 18;   # install for PG 18
pig install pg_projection -v 17;   # install for PG 17
pig install pg_projection -v 16;   # install for PG 16
pig install pg_projection -v 15;   # install for PG 15
pig install pg_projection -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_projection;
```




## Usage

Sources: [pg_projection README](https://github.com/suissa/pg_projection), [SQL definitions](https://github.com/suissa/pg_projection/blob/main/pg_projection--1.0.sql), [control file](https://github.com/suissa/pg_projection/blob/main/pg_projection.control).

`pg_projection` provides MongoDB-style read projections for PostgreSQL `jsonb`. The 1.0 SQL file defines two functions: `pg_project(jsonb, jsonb)` for one JSON document and `pg_project_set(text, jsonb)` for a query result converted to a JSON array.

### Project One JSONB Value

Projection values are numeric flags: `1` includes a field and `0` excludes a field.

```sql
CREATE EXTENSION pg_projection;

SELECT pg_project(
  '{"_id": 7, "name": "Ada", "email": "ada@example.test", "secret": "x"}'::jsonb,
  '{"name": 1, "email": 1}'::jsonb
);
-- {"_id": 7, "name": "Ada", "email": "ada@example.test"}
```

In inclusion mode, `_id` is included by default when present. Exclude it explicitly when the caller wants only the selected fields:

```sql
SELECT pg_project(
  '{"_id": 7, "name": "Ada", "email": "ada@example.test"}'::jsonb,
  '{"_id": 0, "name": 1}'::jsonb
);
-- {"name": "Ada"}
```

### Exclude Fields

When the projection uses `0`, the function starts from the original document and removes matching top-level keys:

```sql
SELECT pg_project(
  '{"name": "Ada", "internal_id": "a-1", "secret_key": "k"}'::jsonb,
  '{"internal_id": 0, "secret_key": 0}'::jsonb
);
-- {"name": "Ada"}
```

### Project A Query Result

`pg_project_set(query_text, projection_json)` executes the supplied SQL text, converts each row with `to_jsonb(t)`, applies `pg_project`, and returns a JSON array:

```sql
SELECT pg_project_set(
  'SELECT id, username, password_hash FROM users WHERE active',
  '{"password_hash": 0}'::jsonb
);
```

Because `query_text` is dynamic SQL, pass only trusted query strings assembled by application or migration code you control. Do not concatenate untrusted user input into this argument.

### Caveats

- The SQL implementation projects top-level keys; it does not implement nested MongoDB path projection.
- Projection values are cast to integers internally, so use numeric `0` and `1` flags.
- `pg_project(jsonb, jsonb)` is declared `IMMUTABLE STRICT`; `pg_project_set(text, jsonb)` is declared `STABLE`.
