---
title: "pg_rusage"
linkTitle: "pg_rusage"
description: "Measure backend CPU usage between explicit reset and print calls"
weight: 6230
categories: ["STAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_rusage**](https://github.com/michaelpq/pg_plugins/tree/ae57c1f3df697fb942f4576f873a655187193ace/pg_rusage) : Measure backend CPU usage between explicit reset and print calls


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6230** | {{< badge content="pg_rusage" link="https://github.com/michaelpq/pg_plugins/tree/ae57c1f3df697fb942f4576f873a655187193ace/pg_rusage" >}} | {{< ext "pg_rusage" >}} | `1.0` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_proctab" >}} {{< ext "system_stats" >}} |

> [!Note] pg_rusage_print() emits the current backend measurement as a WARNING and returns void. No preload; PostgreSQL 14-18 verified in PGSTY builds; upstream declares no major-version matrix.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_rusage` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_rusage_18" "green" >}} {{< bg "17" "pg_rusage_17" "green" >}} {{< bg "16" "pg_rusage_16" "green" >}} {{< bg "15" "pg_rusage_15" "green" >}} {{< bg "14" "pg_rusage_14" "green" >}} | `pg_rusage_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-rusage" "green" >}} {{< bg "17" "postgresql-17-pg-rusage" "green" >}} {{< bg "16" "postgresql-16-pg-rusage" "green" >}} {{< bg "15" "postgresql-15-pg-rusage" "green" >}} {{< bg "14" "postgresql-14-pg-rusage" "green" >}} | `postgresql-$v-pg-rusage` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_rusage_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-rusage : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-rusage : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-rusage : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-rusage : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-rusage : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-rusage : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-rusage : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-rusage : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-rusage : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-rusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-rusage : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rusage_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pg_rusage_18-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rusage_18-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_rusage_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.9 KiB | [pg_rusage_18-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rusage_18-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_rusage_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.7 KiB | [pg_rusage_18-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rusage_18-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_rusage_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.6 KiB | [pg_rusage_18-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rusage_18-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_rusage_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.7 KiB | [pg_rusage_18-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rusage_18-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_rusage_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.8 KiB | [pg_rusage_18-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rusage_18-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-rusage` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.4 KiB | [postgresql-18-pg-rusage_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rusage/postgresql-18-pg-rusage_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-rusage` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.5 KiB | [postgresql-18-pg-rusage_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rusage/postgresql-18-pg-rusage_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-rusage` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.4 KiB | [postgresql-18-pg-rusage_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rusage/postgresql-18-pg-rusage_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-rusage` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.6 KiB | [postgresql-18-pg-rusage_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rusage/postgresql-18-pg-rusage_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-rusage` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-18-pg-rusage_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rusage/postgresql-18-pg-rusage_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-rusage` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-18-pg-rusage_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rusage/postgresql-18-pg-rusage_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-rusage` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.7 KiB | [postgresql-18-pg-rusage_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rusage/postgresql-18-pg-rusage_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-rusage` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.7 KiB | [postgresql-18-pg-rusage_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rusage/postgresql-18-pg-rusage_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-rusage` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.8 KiB | [postgresql-18-pg-rusage_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rusage/postgresql-18-pg-rusage_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-rusage` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.9 KiB | [postgresql-18-pg-rusage_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rusage/postgresql-18-pg-rusage_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rusage_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pg_rusage_17-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rusage_17-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_rusage_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.9 KiB | [pg_rusage_17-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rusage_17-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_rusage_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.6 KiB | [pg_rusage_17-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rusage_17-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_rusage_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.5 KiB | [pg_rusage_17-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rusage_17-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_rusage_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.7 KiB | [pg_rusage_17-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rusage_17-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_rusage_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.8 KiB | [pg_rusage_17-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rusage_17-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-rusage` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.4 KiB | [postgresql-17-pg-rusage_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rusage/postgresql-17-pg-rusage_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-rusage` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.5 KiB | [postgresql-17-pg-rusage_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rusage/postgresql-17-pg-rusage_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-rusage` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.4 KiB | [postgresql-17-pg-rusage_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rusage/postgresql-17-pg-rusage_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-rusage` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.5 KiB | [postgresql-17-pg-rusage_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rusage/postgresql-17-pg-rusage_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-rusage` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.5 KiB | [postgresql-17-pg-rusage_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rusage/postgresql-17-pg-rusage_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-rusage` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.5 KiB | [postgresql-17-pg-rusage_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rusage/postgresql-17-pg-rusage_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-rusage` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.7 KiB | [postgresql-17-pg-rusage_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rusage/postgresql-17-pg-rusage_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-rusage` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.6 KiB | [postgresql-17-pg-rusage_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rusage/postgresql-17-pg-rusage_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-rusage` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.8 KiB | [postgresql-17-pg-rusage_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rusage/postgresql-17-pg-rusage_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-rusage` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.9 KiB | [postgresql-17-pg-rusage_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rusage/postgresql-17-pg-rusage_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rusage_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pg_rusage_16-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rusage_16-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_rusage_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.9 KiB | [pg_rusage_16-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rusage_16-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_rusage_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.6 KiB | [pg_rusage_16-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rusage_16-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_rusage_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.5 KiB | [pg_rusage_16-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rusage_16-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_rusage_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.7 KiB | [pg_rusage_16-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rusage_16-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_rusage_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.8 KiB | [pg_rusage_16-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rusage_16-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-rusage` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.4 KiB | [postgresql-16-pg-rusage_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rusage/postgresql-16-pg-rusage_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-rusage` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.5 KiB | [postgresql-16-pg-rusage_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rusage/postgresql-16-pg-rusage_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-rusage` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.4 KiB | [postgresql-16-pg-rusage_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rusage/postgresql-16-pg-rusage_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-rusage` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.5 KiB | [postgresql-16-pg-rusage_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rusage/postgresql-16-pg-rusage_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-rusage` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.5 KiB | [postgresql-16-pg-rusage_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rusage/postgresql-16-pg-rusage_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-rusage` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.5 KiB | [postgresql-16-pg-rusage_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rusage/postgresql-16-pg-rusage_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-rusage` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.6 KiB | [postgresql-16-pg-rusage_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rusage/postgresql-16-pg-rusage_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-rusage` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.6 KiB | [postgresql-16-pg-rusage_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rusage/postgresql-16-pg-rusage_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-rusage` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.8 KiB | [postgresql-16-pg-rusage_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rusage/postgresql-16-pg-rusage_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-rusage` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.9 KiB | [postgresql-16-pg-rusage_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rusage/postgresql-16-pg-rusage_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rusage_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pg_rusage_15-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rusage_15-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_rusage_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.9 KiB | [pg_rusage_15-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rusage_15-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_rusage_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.6 KiB | [pg_rusage_15-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rusage_15-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_rusage_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.5 KiB | [pg_rusage_15-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rusage_15-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_rusage_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.7 KiB | [pg_rusage_15-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rusage_15-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_rusage_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.8 KiB | [pg_rusage_15-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rusage_15-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-rusage` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.4 KiB | [postgresql-15-pg-rusage_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rusage/postgresql-15-pg-rusage_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-rusage` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.5 KiB | [postgresql-15-pg-rusage_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rusage/postgresql-15-pg-rusage_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-rusage` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.4 KiB | [postgresql-15-pg-rusage_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rusage/postgresql-15-pg-rusage_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-rusage` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.5 KiB | [postgresql-15-pg-rusage_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rusage/postgresql-15-pg-rusage_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-rusage` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.5 KiB | [postgresql-15-pg-rusage_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rusage/postgresql-15-pg-rusage_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-rusage` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.5 KiB | [postgresql-15-pg-rusage_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rusage/postgresql-15-pg-rusage_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-rusage` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.7 KiB | [postgresql-15-pg-rusage_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rusage/postgresql-15-pg-rusage_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-rusage` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.6 KiB | [postgresql-15-pg-rusage_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rusage/postgresql-15-pg-rusage_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-rusage` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.8 KiB | [postgresql-15-pg-rusage_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rusage/postgresql-15-pg-rusage_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-rusage` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.9 KiB | [postgresql-15-pg-rusage_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rusage/postgresql-15-pg-rusage_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_rusage_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pg_rusage_14-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rusage_14-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_rusage_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.8 KiB | [pg_rusage_14-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rusage_14-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_rusage_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.6 KiB | [pg_rusage_14-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rusage_14-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_rusage_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.5 KiB | [pg_rusage_14-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rusage_14-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_rusage_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.7 KiB | [pg_rusage_14-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_rusage_14-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_rusage_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.8 KiB | [pg_rusage_14-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_rusage_14-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-rusage` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.0 KiB | [postgresql-14-pg-rusage_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rusage/postgresql-14-pg-rusage_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-rusage` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.1 KiB | [postgresql-14-pg-rusage_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-rusage/postgresql-14-pg-rusage_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-rusage` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.0 KiB | [postgresql-14-pg-rusage_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rusage/postgresql-14-pg-rusage_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-rusage` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.1 KiB | [postgresql-14-pg-rusage_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-rusage/postgresql-14-pg-rusage_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-rusage` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 8.1 KiB | [postgresql-14-pg-rusage_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rusage/postgresql-14-pg-rusage_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-rusage` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.2 KiB | [postgresql-14-pg-rusage_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-rusage/postgresql-14-pg-rusage_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-rusage` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.3 KiB | [postgresql-14-pg-rusage_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rusage/postgresql-14-pg-rusage_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-rusage` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.2 KiB | [postgresql-14-pg-rusage_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-rusage/postgresql-14-pg-rusage_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-rusage` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 8.5 KiB | [postgresql-14-pg-rusage_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rusage/postgresql-14-pg-rusage_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-rusage` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.5 KiB | [postgresql-14-pg-rusage_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-rusage/postgresql-14-pg-rusage_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/michaelpq/pg_plugins/tree/ae57c1f3df697fb942f4576f873a655187193ace/pg_rusage" title="Repository" icon="github" subtitle="github.com/michaelpq/pg_plugins/tree/ae57c1f3df697fb942f4576f873a655187193ace/pg_rusage" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_plugins-ae57c1f3df69.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_rusage;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_rusage;		# install via package name, for the active PG version

pig install pg_rusage -v 18;   # install for PG 18
pig install pg_rusage -v 17;   # install for PG 17
pig install pg_rusage -v 16;   # install for PG 16
pig install pg_rusage -v 15;   # install for PG 15
pig install pg_rusage -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_rusage;
```

## Usage

Sources:

- [Official upstream README](https://github.com/michaelpq/pg_plugins/blob/626fb56b0a0b833d4f23ca55359ce56d38162864/pg_rusage/README)
- [Official extension control file (pg_rusage.control)](https://github.com/michaelpq/pg_plugins/blob/626fb56b0a0b833d4f23ca55359ce56d38162864/pg_rusage/pg_rusage.control)
- [Official extension SQL (pg_rusage--1.0.sql)](https://github.com/michaelpq/pg_plugins/blob/626fb56b0a0b833d4f23ca55359ce56d38162864/pg_rusage/pg_rusage--1.0.sql)

`pg_rusage` — This module is a PostgreSQL extension that can enable CPU measurements, with one SQL function to enable the measurement and one to disable it. When disabling, the existing accumulated results will show up. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_rusage;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_rusage_print()` is an extension function and returns `void`.
- `pg_rusage_reset()` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
