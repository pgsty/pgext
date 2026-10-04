---
title: "pg_statvfs"
linkTitle: "pg_statvfs"
description: "Server filesystem capacity and inode statistics through statvfs"
weight: 6470
categories: ["STAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_statvfs**](https://github.com/michaelpq/pg_plugins/tree/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs) : Server filesystem capacity and inode statistics through statvfs


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6470** | {{< badge content="pg_statvfs" link="https://github.com/michaelpq/pg_plugins/tree/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs" >}} | {{< ext "pg_statvfs" >}} | `1.0` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "system_stats" >}} {{< ext "pgnodemx" >}} {{< ext "pgexporter_ext" >}} {{< ext "pg_proctab" >}} |

> [!Note] Superuser-only; checks server paths under data/log-directory rules. No preload. PostgreSQL 14-18 verified in PGSTY builds; upstream declares no major-version matrix.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_statvfs` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_statvfs_18" "green" >}} {{< bg "17" "pg_statvfs_17" "green" >}} {{< bg "16" "pg_statvfs_16" "green" >}} {{< bg "15" "pg_statvfs_15" "green" >}} {{< bg "14" "pg_statvfs_14" "green" >}} | `pg_statvfs_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-statvfs" "green" >}} {{< bg "17" "postgresql-17-pg-statvfs" "green" >}} {{< bg "16" "postgresql-16-pg-statvfs" "green" >}} {{< bg "15" "postgresql-15-pg-statvfs" "green" >}} {{< bg "14" "postgresql-14-pg-statvfs" "green" >}} | `postgresql-$v-pg-statvfs` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_statvfs_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-statvfs : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-statvfs : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-statvfs : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-statvfs : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-statvfs : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-statvfs : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-statvfs : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-statvfs : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-statvfs : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-statvfs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-statvfs : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statvfs_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.4 KiB | [pg_statvfs_18-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_statvfs_18-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_statvfs_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.4 KiB | [pg_statvfs_18-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_statvfs_18-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_statvfs_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.6 KiB | [pg_statvfs_18-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_statvfs_18-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_statvfs_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.5 KiB | [pg_statvfs_18-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_statvfs_18-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_statvfs_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.6 KiB | [pg_statvfs_18-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_statvfs_18-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_statvfs_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.8 KiB | [pg_statvfs_18-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_statvfs_18-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-statvfs` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.3 KiB | [postgresql-18-pg-statvfs_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statvfs/postgresql-18-pg-statvfs_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-statvfs` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.2 KiB | [postgresql-18-pg-statvfs_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statvfs/postgresql-18-pg-statvfs_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-statvfs` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.3 KiB | [postgresql-18-pg-statvfs_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statvfs/postgresql-18-pg-statvfs_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-statvfs` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.2 KiB | [postgresql-18-pg-statvfs_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statvfs/postgresql-18-pg-statvfs_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-statvfs` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.7 KiB | [postgresql-18-pg-statvfs_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statvfs/postgresql-18-pg-statvfs_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-statvfs` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.5 KiB | [postgresql-18-pg-statvfs_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statvfs/postgresql-18-pg-statvfs_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-statvfs` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.8 KiB | [postgresql-18-pg-statvfs_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statvfs/postgresql-18-pg-statvfs_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-statvfs` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.6 KiB | [postgresql-18-pg-statvfs_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statvfs/postgresql-18-pg-statvfs_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-statvfs` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.7 KiB | [postgresql-18-pg-statvfs_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statvfs/postgresql-18-pg-statvfs_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-statvfs` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.7 KiB | [postgresql-18-pg-statvfs_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statvfs/postgresql-18-pg-statvfs_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statvfs_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.4 KiB | [pg_statvfs_17-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_statvfs_17-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_statvfs_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.4 KiB | [pg_statvfs_17-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_statvfs_17-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_statvfs_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.5 KiB | [pg_statvfs_17-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_statvfs_17-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_statvfs_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.5 KiB | [pg_statvfs_17-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_statvfs_17-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_statvfs_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.6 KiB | [pg_statvfs_17-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_statvfs_17-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_statvfs_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.8 KiB | [pg_statvfs_17-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_statvfs_17-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-statvfs` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.3 KiB | [postgresql-17-pg-statvfs_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statvfs/postgresql-17-pg-statvfs_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-statvfs` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.2 KiB | [postgresql-17-pg-statvfs_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statvfs/postgresql-17-pg-statvfs_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-statvfs` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.3 KiB | [postgresql-17-pg-statvfs_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statvfs/postgresql-17-pg-statvfs_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-statvfs` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.2 KiB | [postgresql-17-pg-statvfs_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statvfs/postgresql-17-pg-statvfs_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-statvfs` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-17-pg-statvfs_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statvfs/postgresql-17-pg-statvfs_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-statvfs` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.2 KiB | [postgresql-17-pg-statvfs_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statvfs/postgresql-17-pg-statvfs_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-statvfs` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.8 KiB | [postgresql-17-pg-statvfs_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statvfs/postgresql-17-pg-statvfs_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-statvfs` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.6 KiB | [postgresql-17-pg-statvfs_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statvfs/postgresql-17-pg-statvfs_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-statvfs` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.7 KiB | [postgresql-17-pg-statvfs_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statvfs/postgresql-17-pg-statvfs_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-statvfs` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.7 KiB | [postgresql-17-pg-statvfs_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statvfs/postgresql-17-pg-statvfs_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statvfs_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.3 KiB | [pg_statvfs_16-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_statvfs_16-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_statvfs_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.4 KiB | [pg_statvfs_16-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_statvfs_16-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_statvfs_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.5 KiB | [pg_statvfs_16-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_statvfs_16-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_statvfs_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.5 KiB | [pg_statvfs_16-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_statvfs_16-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_statvfs_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.6 KiB | [pg_statvfs_16-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_statvfs_16-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_statvfs_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.8 KiB | [pg_statvfs_16-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_statvfs_16-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-statvfs` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.3 KiB | [postgresql-16-pg-statvfs_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statvfs/postgresql-16-pg-statvfs_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-statvfs` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.2 KiB | [postgresql-16-pg-statvfs_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statvfs/postgresql-16-pg-statvfs_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-statvfs` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.3 KiB | [postgresql-16-pg-statvfs_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statvfs/postgresql-16-pg-statvfs_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-statvfs` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.2 KiB | [postgresql-16-pg-statvfs_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statvfs/postgresql-16-pg-statvfs_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-statvfs` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-16-pg-statvfs_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statvfs/postgresql-16-pg-statvfs_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-statvfs` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.2 KiB | [postgresql-16-pg-statvfs_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statvfs/postgresql-16-pg-statvfs_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-statvfs` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.8 KiB | [postgresql-16-pg-statvfs_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statvfs/postgresql-16-pg-statvfs_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-statvfs` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.6 KiB | [postgresql-16-pg-statvfs_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statvfs/postgresql-16-pg-statvfs_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-statvfs` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.7 KiB | [postgresql-16-pg-statvfs_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statvfs/postgresql-16-pg-statvfs_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-statvfs` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.7 KiB | [postgresql-16-pg-statvfs_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statvfs/postgresql-16-pg-statvfs_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statvfs_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.4 KiB | [pg_statvfs_15-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_statvfs_15-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_statvfs_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.4 KiB | [pg_statvfs_15-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_statvfs_15-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_statvfs_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.5 KiB | [pg_statvfs_15-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_statvfs_15-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_statvfs_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.5 KiB | [pg_statvfs_15-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_statvfs_15-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_statvfs_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.6 KiB | [pg_statvfs_15-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_statvfs_15-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_statvfs_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.8 KiB | [pg_statvfs_15-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_statvfs_15-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-statvfs` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.3 KiB | [postgresql-15-pg-statvfs_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statvfs/postgresql-15-pg-statvfs_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-statvfs` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.2 KiB | [postgresql-15-pg-statvfs_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statvfs/postgresql-15-pg-statvfs_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-statvfs` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.3 KiB | [postgresql-15-pg-statvfs_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statvfs/postgresql-15-pg-statvfs_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-statvfs` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.2 KiB | [postgresql-15-pg-statvfs_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statvfs/postgresql-15-pg-statvfs_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-statvfs` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 KiB | [postgresql-15-pg-statvfs_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statvfs/postgresql-15-pg-statvfs_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-statvfs` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.2 KiB | [postgresql-15-pg-statvfs_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statvfs/postgresql-15-pg-statvfs_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-statvfs` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.8 KiB | [postgresql-15-pg-statvfs_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statvfs/postgresql-15-pg-statvfs_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-statvfs` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.6 KiB | [postgresql-15-pg-statvfs_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statvfs/postgresql-15-pg-statvfs_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-statvfs` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.7 KiB | [postgresql-15-pg-statvfs_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statvfs/postgresql-15-pg-statvfs_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-statvfs` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.7 KiB | [postgresql-15-pg-statvfs_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statvfs/postgresql-15-pg-statvfs_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statvfs_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.3 KiB | [pg_statvfs_14-1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_statvfs_14-1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_statvfs_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.4 KiB | [pg_statvfs_14-1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_statvfs_14-1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_statvfs_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.5 KiB | [pg_statvfs_14-1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_statvfs_14-1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_statvfs_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.5 KiB | [pg_statvfs_14-1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_statvfs_14-1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_statvfs_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.6 KiB | [pg_statvfs_14-1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_statvfs_14-1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_statvfs_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.7 KiB | [pg_statvfs_14-1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_statvfs_14-1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-statvfs` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.9 KiB | [postgresql-14-pg-statvfs_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statvfs/postgresql-14-pg-statvfs_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-statvfs` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.7 KiB | [postgresql-14-pg-statvfs_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statvfs/postgresql-14-pg-statvfs_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-statvfs` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.9 KiB | [postgresql-14-pg-statvfs_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statvfs/postgresql-14-pg-statvfs_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-statvfs` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.8 KiB | [postgresql-14-pg-statvfs_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statvfs/postgresql-14-pg-statvfs_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-statvfs` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.0 KiB | [postgresql-14-pg-statvfs_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statvfs/postgresql-14-pg-statvfs_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-statvfs` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.8 KiB | [postgresql-14-pg-statvfs_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statvfs/postgresql-14-pg-statvfs_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-statvfs` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.5 KiB | [postgresql-14-pg-statvfs_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statvfs/postgresql-14-pg-statvfs_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-statvfs` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.2 KiB | [postgresql-14-pg-statvfs_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statvfs/postgresql-14-pg-statvfs_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-statvfs` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.3 KiB | [postgresql-14-pg-statvfs_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statvfs/postgresql-14-pg-statvfs_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-statvfs` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.3 KiB | [postgresql-14-pg-statvfs_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statvfs/postgresql-14-pg-statvfs_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/michaelpq/pg_plugins/tree/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs" title="Repository" icon="github" subtitle="github.com/michaelpq/pg_plugins/tree/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_plugins-ae57c1f3df69.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_statvfs;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_statvfs;		# install via package name, for the active PG version

pig install pg_statvfs -v 18;   # install for PG 18
pig install pg_statvfs -v 17;   # install for PG 17
pig install pg_statvfs -v 16;   # install for PG 16
pig install pg_statvfs -v 15;   # install for PG 15
pig install pg_statvfs -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_statvfs;
```

## Usage

Sources:

- [Official documentation](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs/README)
- [Control file](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs/pg_statvfs.control)
- [Version 1.0 SQL](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs/pg_statvfs--1.0.sql)
- [Privilege and path checks](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs/pg_statvfs.c)

`pg_statvfs` 1.0 exposes filesystem capacity, inode availability, and mount flags through the server's statvfs call. It inspects the filesystem containing a server-side path, not a client directory or the size of an individual relation.

### Core Workflow

After installing the extension files, use a superuser session:

```sql
CREATE EXTENSION pg_statvfs;
SELECT * FROM pg_statvfs(current_setting('data_directory'));
SELECT f_frsize * f_blocks AS total_bytes,
       f_frsize * f_bavail AS available_bytes
FROM pg_statvfs(current_setting('data_directory'));
```

### Result and Access

`pg_statvfs(path text)` returns one record. `f_bsize` is the preferred block size; `f_frsize` is the unit for block counts. `f_blocks`, `f_bfree`, and `f_bavail` report total, free, and unprivileged-available blocks. `f_files`, `f_ffree`, and `f_favail` describe inode counts; `f_fsid`, `f_namemax`, and `flags` expose filesystem identity, name limits, and supported mount flags.

The C function explicitly requires superuser privileges and validates the path before asking the operating system. Missing or inaccessible paths produce an error. The extension needs no preload or restart of its own. Its control file permits relocation, but does not mark it trusted. The upstream subdirectory does not declare a complete PostgreSQL-major compatibility matrix; operating-system support and mount-flag coverage also vary.
