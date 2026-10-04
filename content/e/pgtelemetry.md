---
title: "pgtelemetry"
linkTitle: "pgtelemetry"
description: "SQL monitoring views for database activity, WAL, replication, and storage"
weight: 6080
categories: ["STAT"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pgtelemetry**](https://github.com/adjust/pg-telemetry) : SQL monitoring views for database activity, WAL, replication, and storage


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6080** | {{< badge content="pgtelemetry" link="https://github.com/adjust/pg-telemetry" >}} | {{< ext "pgtelemetry" >}} | `1.7` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgtelemetry` |
|   **Requires**    | {{< ext "pg_stat_statements" >}} {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pgmonitor" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_profile" >}} {{< ext "powa" >}} {{< ext "pgexporter_ext" >}} |

> [!Note] Pure SQL in the fixed pgtelemetry schema; requires pg_stat_statements and plpgsql. Preload pg_stat_statements, not pgtelemetry. Control/SQL are 1.7 while META.json still says 1.6.0; upstream CI stops at PG17, and PG14-18 are verified in PGSTY builds.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgtelemetry` | `pg_stat_statements`, `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7` | {{< bg "18" "pgtelemetry_18" "green" >}} {{< bg "17" "pgtelemetry_17" "green" >}} {{< bg "16" "pgtelemetry_16" "green" >}} {{< bg "15" "pgtelemetry_15" "green" >}} {{< bg "14" "pgtelemetry_14" "green" >}} | `pgtelemetry_$v` | `postgresql$v-contrib` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7` | {{< bg "18" "postgresql-18-pgtelemetry" "green" >}} {{< bg "17" "postgresql-17-pgtelemetry" "green" >}} {{< bg "16" "postgresql-16-pgtelemetry" "green" >}} {{< bg "15" "postgresql-15-pgtelemetry" "green" >}} {{< bg "14" "postgresql-14-pgtelemetry" "green" >}} | `postgresql-$v-pgtelemetry` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "pgtelemetry_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-pgtelemetry : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-pgtelemetry : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-pgtelemetry : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-pgtelemetry : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-pgtelemetry : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-pgtelemetry : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-pgtelemetry : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-pgtelemetry : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-pgtelemetry : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.7" "postgresql-18-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-17-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-16-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-15-pgtelemetry : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7" "postgresql-14-pgtelemetry : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtelemetry_18` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.6 KiB | [pgtelemetry_18-1.7-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgtelemetry_18-1.7-1PGSTY.el8.noarch.rpm) |
| `pgtelemetry_18` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.6 KiB | [pgtelemetry_18-1.7-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgtelemetry_18-1.7-1PGSTY.el8.noarch.rpm) |
| `pgtelemetry_18` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.4 KiB | [pgtelemetry_18-1.7-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgtelemetry_18-1.7-1PGSTY.el9.noarch.rpm) |
| `pgtelemetry_18` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.4 KiB | [pgtelemetry_18-1.7-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgtelemetry_18-1.7-1PGSTY.el9.noarch.rpm) |
| `pgtelemetry_18` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.5 KiB | [pgtelemetry_18-1.7-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgtelemetry_18-1.7-1PGSTY.el10.noarch.rpm) |
| `pgtelemetry_18` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [pgtelemetry_18-1.7-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgtelemetry_18-1.7-1PGSTY.el10.noarch.rpm) |
| `postgresql-18-pgtelemetry` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.1 KiB | [postgresql-18-pgtelemetry_1.7-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtelemetry/postgresql-18-pgtelemetry_1.7-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pgtelemetry` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.1 KiB | [postgresql-18-pgtelemetry_1.7-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtelemetry/postgresql-18-pgtelemetry_1.7-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pgtelemetry` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.1 KiB | [postgresql-18-pgtelemetry_1.7-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgtelemetry/postgresql-18-pgtelemetry_1.7-1PGSTY~trixie_all.deb) |
| `postgresql-18-pgtelemetry` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.1 KiB | [postgresql-18-pgtelemetry_1.7-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgtelemetry/postgresql-18-pgtelemetry_1.7-1PGSTY~trixie_all.deb) |
| `postgresql-18-pgtelemetry` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.6 KiB | [postgresql-18-pgtelemetry_1.7-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtelemetry/postgresql-18-pgtelemetry_1.7-1PGSTY~jammy_all.deb) |
| `postgresql-18-pgtelemetry` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.6 KiB | [postgresql-18-pgtelemetry_1.7-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtelemetry/postgresql-18-pgtelemetry_1.7-1PGSTY~jammy_all.deb) |
| `postgresql-18-pgtelemetry` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.6 KiB | [postgresql-18-pgtelemetry_1.7-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtelemetry/postgresql-18-pgtelemetry_1.7-1PGSTY~noble_all.deb) |
| `postgresql-18-pgtelemetry` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.6 KiB | [postgresql-18-pgtelemetry_1.7-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtelemetry/postgresql-18-pgtelemetry_1.7-1PGSTY~noble_all.deb) |
| `postgresql-18-pgtelemetry` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.6 KiB | [postgresql-18-pgtelemetry_1.7-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgtelemetry/postgresql-18-pgtelemetry_1.7-1PGSTY~resolute_all.deb) |
| `postgresql-18-pgtelemetry` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 9.6 KiB | [postgresql-18-pgtelemetry_1.7-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgtelemetry/postgresql-18-pgtelemetry_1.7-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtelemetry_17` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.6 KiB | [pgtelemetry_17-1.7-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgtelemetry_17-1.7-1PGSTY.el8.noarch.rpm) |
| `pgtelemetry_17` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.6 KiB | [pgtelemetry_17-1.7-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgtelemetry_17-1.7-1PGSTY.el8.noarch.rpm) |
| `pgtelemetry_17` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.4 KiB | [pgtelemetry_17-1.7-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgtelemetry_17-1.7-1PGSTY.el9.noarch.rpm) |
| `pgtelemetry_17` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.4 KiB | [pgtelemetry_17-1.7-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgtelemetry_17-1.7-1PGSTY.el9.noarch.rpm) |
| `pgtelemetry_17` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.5 KiB | [pgtelemetry_17-1.7-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgtelemetry_17-1.7-1PGSTY.el10.noarch.rpm) |
| `pgtelemetry_17` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [pgtelemetry_17-1.7-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgtelemetry_17-1.7-1PGSTY.el10.noarch.rpm) |
| `postgresql-17-pgtelemetry` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.1 KiB | [postgresql-17-pgtelemetry_1.7-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtelemetry/postgresql-17-pgtelemetry_1.7-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pgtelemetry` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.1 KiB | [postgresql-17-pgtelemetry_1.7-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtelemetry/postgresql-17-pgtelemetry_1.7-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pgtelemetry` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.1 KiB | [postgresql-17-pgtelemetry_1.7-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgtelemetry/postgresql-17-pgtelemetry_1.7-1PGSTY~trixie_all.deb) |
| `postgresql-17-pgtelemetry` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.1 KiB | [postgresql-17-pgtelemetry_1.7-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgtelemetry/postgresql-17-pgtelemetry_1.7-1PGSTY~trixie_all.deb) |
| `postgresql-17-pgtelemetry` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.7 KiB | [postgresql-17-pgtelemetry_1.7-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtelemetry/postgresql-17-pgtelemetry_1.7-1PGSTY~jammy_all.deb) |
| `postgresql-17-pgtelemetry` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.7 KiB | [postgresql-17-pgtelemetry_1.7-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtelemetry/postgresql-17-pgtelemetry_1.7-1PGSTY~jammy_all.deb) |
| `postgresql-17-pgtelemetry` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.6 KiB | [postgresql-17-pgtelemetry_1.7-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtelemetry/postgresql-17-pgtelemetry_1.7-1PGSTY~noble_all.deb) |
| `postgresql-17-pgtelemetry` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.6 KiB | [postgresql-17-pgtelemetry_1.7-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtelemetry/postgresql-17-pgtelemetry_1.7-1PGSTY~noble_all.deb) |
| `postgresql-17-pgtelemetry` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.7 KiB | [postgresql-17-pgtelemetry_1.7-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgtelemetry/postgresql-17-pgtelemetry_1.7-1PGSTY~resolute_all.deb) |
| `postgresql-17-pgtelemetry` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 9.7 KiB | [postgresql-17-pgtelemetry_1.7-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgtelemetry/postgresql-17-pgtelemetry_1.7-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtelemetry_16` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.6 KiB | [pgtelemetry_16-1.7-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgtelemetry_16-1.7-1PGSTY.el8.noarch.rpm) |
| `pgtelemetry_16` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.6 KiB | [pgtelemetry_16-1.7-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgtelemetry_16-1.7-1PGSTY.el8.noarch.rpm) |
| `pgtelemetry_16` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.4 KiB | [pgtelemetry_16-1.7-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgtelemetry_16-1.7-1PGSTY.el9.noarch.rpm) |
| `pgtelemetry_16` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.4 KiB | [pgtelemetry_16-1.7-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgtelemetry_16-1.7-1PGSTY.el9.noarch.rpm) |
| `pgtelemetry_16` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.5 KiB | [pgtelemetry_16-1.7-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgtelemetry_16-1.7-1PGSTY.el10.noarch.rpm) |
| `pgtelemetry_16` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [pgtelemetry_16-1.7-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgtelemetry_16-1.7-1PGSTY.el10.noarch.rpm) |
| `postgresql-16-pgtelemetry` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.1 KiB | [postgresql-16-pgtelemetry_1.7-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtelemetry/postgresql-16-pgtelemetry_1.7-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pgtelemetry` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.1 KiB | [postgresql-16-pgtelemetry_1.7-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtelemetry/postgresql-16-pgtelemetry_1.7-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pgtelemetry` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.1 KiB | [postgresql-16-pgtelemetry_1.7-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgtelemetry/postgresql-16-pgtelemetry_1.7-1PGSTY~trixie_all.deb) |
| `postgresql-16-pgtelemetry` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.1 KiB | [postgresql-16-pgtelemetry_1.7-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgtelemetry/postgresql-16-pgtelemetry_1.7-1PGSTY~trixie_all.deb) |
| `postgresql-16-pgtelemetry` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.7 KiB | [postgresql-16-pgtelemetry_1.7-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtelemetry/postgresql-16-pgtelemetry_1.7-1PGSTY~jammy_all.deb) |
| `postgresql-16-pgtelemetry` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.7 KiB | [postgresql-16-pgtelemetry_1.7-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtelemetry/postgresql-16-pgtelemetry_1.7-1PGSTY~jammy_all.deb) |
| `postgresql-16-pgtelemetry` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.6 KiB | [postgresql-16-pgtelemetry_1.7-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtelemetry/postgresql-16-pgtelemetry_1.7-1PGSTY~noble_all.deb) |
| `postgresql-16-pgtelemetry` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.6 KiB | [postgresql-16-pgtelemetry_1.7-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtelemetry/postgresql-16-pgtelemetry_1.7-1PGSTY~noble_all.deb) |
| `postgresql-16-pgtelemetry` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.7 KiB | [postgresql-16-pgtelemetry_1.7-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgtelemetry/postgresql-16-pgtelemetry_1.7-1PGSTY~resolute_all.deb) |
| `postgresql-16-pgtelemetry` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 9.7 KiB | [postgresql-16-pgtelemetry_1.7-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgtelemetry/postgresql-16-pgtelemetry_1.7-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtelemetry_15` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.6 KiB | [pgtelemetry_15-1.7-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgtelemetry_15-1.7-1PGSTY.el8.noarch.rpm) |
| `pgtelemetry_15` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.6 KiB | [pgtelemetry_15-1.7-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgtelemetry_15-1.7-1PGSTY.el8.noarch.rpm) |
| `pgtelemetry_15` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.4 KiB | [pgtelemetry_15-1.7-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgtelemetry_15-1.7-1PGSTY.el9.noarch.rpm) |
| `pgtelemetry_15` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.4 KiB | [pgtelemetry_15-1.7-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgtelemetry_15-1.7-1PGSTY.el9.noarch.rpm) |
| `pgtelemetry_15` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.5 KiB | [pgtelemetry_15-1.7-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgtelemetry_15-1.7-1PGSTY.el10.noarch.rpm) |
| `pgtelemetry_15` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [pgtelemetry_15-1.7-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgtelemetry_15-1.7-1PGSTY.el10.noarch.rpm) |
| `postgresql-15-pgtelemetry` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.1 KiB | [postgresql-15-pgtelemetry_1.7-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtelemetry/postgresql-15-pgtelemetry_1.7-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pgtelemetry` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.1 KiB | [postgresql-15-pgtelemetry_1.7-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtelemetry/postgresql-15-pgtelemetry_1.7-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pgtelemetry` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.1 KiB | [postgresql-15-pgtelemetry_1.7-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgtelemetry/postgresql-15-pgtelemetry_1.7-1PGSTY~trixie_all.deb) |
| `postgresql-15-pgtelemetry` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.1 KiB | [postgresql-15-pgtelemetry_1.7-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgtelemetry/postgresql-15-pgtelemetry_1.7-1PGSTY~trixie_all.deb) |
| `postgresql-15-pgtelemetry` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.7 KiB | [postgresql-15-pgtelemetry_1.7-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtelemetry/postgresql-15-pgtelemetry_1.7-1PGSTY~jammy_all.deb) |
| `postgresql-15-pgtelemetry` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.7 KiB | [postgresql-15-pgtelemetry_1.7-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtelemetry/postgresql-15-pgtelemetry_1.7-1PGSTY~jammy_all.deb) |
| `postgresql-15-pgtelemetry` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.6 KiB | [postgresql-15-pgtelemetry_1.7-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtelemetry/postgresql-15-pgtelemetry_1.7-1PGSTY~noble_all.deb) |
| `postgresql-15-pgtelemetry` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.6 KiB | [postgresql-15-pgtelemetry_1.7-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtelemetry/postgresql-15-pgtelemetry_1.7-1PGSTY~noble_all.deb) |
| `postgresql-15-pgtelemetry` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.7 KiB | [postgresql-15-pgtelemetry_1.7-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgtelemetry/postgresql-15-pgtelemetry_1.7-1PGSTY~resolute_all.deb) |
| `postgresql-15-pgtelemetry` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 9.7 KiB | [postgresql-15-pgtelemetry_1.7-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgtelemetry/postgresql-15-pgtelemetry_1.7-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgtelemetry_14` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.6 KiB | [pgtelemetry_14-1.7-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgtelemetry_14-1.7-1PGSTY.el8.noarch.rpm) |
| `pgtelemetry_14` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.6 KiB | [pgtelemetry_14-1.7-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgtelemetry_14-1.7-1PGSTY.el8.noarch.rpm) |
| `pgtelemetry_14` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.4 KiB | [pgtelemetry_14-1.7-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgtelemetry_14-1.7-1PGSTY.el9.noarch.rpm) |
| `pgtelemetry_14` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.4 KiB | [pgtelemetry_14-1.7-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgtelemetry_14-1.7-1PGSTY.el9.noarch.rpm) |
| `pgtelemetry_14` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.5 KiB | [pgtelemetry_14-1.7-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgtelemetry_14-1.7-1PGSTY.el10.noarch.rpm) |
| `pgtelemetry_14` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [pgtelemetry_14-1.7-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgtelemetry_14-1.7-1PGSTY.el10.noarch.rpm) |
| `postgresql-14-pgtelemetry` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.3 KiB | [postgresql-14-pgtelemetry_1.7-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtelemetry/postgresql-14-pgtelemetry_1.7-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pgtelemetry` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.3 KiB | [postgresql-14-pgtelemetry_1.7-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgtelemetry/postgresql-14-pgtelemetry_1.7-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pgtelemetry` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.3 KiB | [postgresql-14-pgtelemetry_1.7-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgtelemetry/postgresql-14-pgtelemetry_1.7-1PGSTY~trixie_all.deb) |
| `postgresql-14-pgtelemetry` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.3 KiB | [postgresql-14-pgtelemetry_1.7-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgtelemetry/postgresql-14-pgtelemetry_1.7-1PGSTY~trixie_all.deb) |
| `postgresql-14-pgtelemetry` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.8 KiB | [postgresql-14-pgtelemetry_1.7-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtelemetry/postgresql-14-pgtelemetry_1.7-1PGSTY~jammy_all.deb) |
| `postgresql-14-pgtelemetry` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.8 KiB | [postgresql-14-pgtelemetry_1.7-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgtelemetry/postgresql-14-pgtelemetry_1.7-1PGSTY~jammy_all.deb) |
| `postgresql-14-pgtelemetry` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.8 KiB | [postgresql-14-pgtelemetry_1.7-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtelemetry/postgresql-14-pgtelemetry_1.7-1PGSTY~noble_all.deb) |
| `postgresql-14-pgtelemetry` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.8 KiB | [postgresql-14-pgtelemetry_1.7-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgtelemetry/postgresql-14-pgtelemetry_1.7-1PGSTY~noble_all.deb) |
| `postgresql-14-pgtelemetry` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 10.9 KiB | [postgresql-14-pgtelemetry_1.7-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgtelemetry/postgresql-14-pgtelemetry_1.7-1PGSTY~resolute_all.deb) |
| `postgresql-14-pgtelemetry` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.9 KiB | [postgresql-14-pgtelemetry_1.7-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgtelemetry/postgresql-14-pgtelemetry_1.7-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adjust/pg-telemetry" title="Repository" icon="github" subtitle="github.com/adjust/pg-telemetry" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgtelemetry-1.7-cac3d192a119.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgtelemetry;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgtelemetry;		# install via package name, for the active PG version

pig install pgtelemetry -v 18;   # install for PG 18
pig install pgtelemetry -v 17;   # install for PG 17
pig install pgtelemetry -v 16;   # install for PG 16
pig install pgtelemetry -v 15;   # install for PG 15
pig install pgtelemetry -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgtelemetry CASCADE; -- requires pg_stat_statements, plpgsql
```

## Usage

Sources:

- [Upstream README](https://github.com/adjust/pg-telemetry/blob/cac3d192a119f25dc1964b13624e49b5a5a6a27c/README.md)
- [Extension control file](https://github.com/adjust/pg-telemetry/blob/cac3d192a119f25dc1964b13624e49b5a5a6a27c/pgtelemetry.control)
- [Version 1.7 SQL](https://github.com/adjust/pg-telemetry/blob/cac3d192a119f25dc1964b13624e49b5a5a6a27c/extension/pgtelemetry--1.7.sql)
- [Generated object documentation](https://github.com/adjust/pg-telemetry/blob/cac3d192a119f25dc1964b13624e49b5a5a6a27c/doc/pgtelemetry.html)

`pgtelemetry` version `1.7` is a pure-SQL monitoring bundle. Views and functions summarize relation/database/tablespace size, connections and waits, locks, table access and autovacuum, `pg_stat_statements` timing and buffers, WAL rate, replication slots, and standby lag.

### Example

Preload and install `pg_stat_statements` first, then install this extension in its fixed `pgtelemetry` schema:

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pgtelemetry;
SELECT * FROM pgtelemetry.database_size;
SELECT * FROM pgtelemetry.connections_by_state;
SELECT * FROM pgtelemetry.longest_running_active_queries LIMIT 10;
```

Installation is superuser-only. Several objects expose cluster-wide query text, client addresses, sizes, locks, replication state, and other operational metadata; `wal_telemetry()` also maintains a history table. Grant only the specific views/functions a monitoring role needs, protect exported metrics, set retention for WAL samples, and account for polling cost and PostgreSQL-version changes in statistics catalogs.
