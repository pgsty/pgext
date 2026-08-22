---
title: "pg_statviz"
linkTitle: "pg_statviz"
description: "Capture PostgreSQL statistics snapshots for time-series analysis and visualization"
weight: 6080
categories: ["STAT"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**pg_statviz**](https://github.com/vyruss/pg_statviz) : Capture PostgreSQL statistics snapshots for time-series analysis and visualization


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6080** | {{< badge content="pg_statviz" link="https://github.com/vyruss/pg_statviz" >}} | {{< ext "pg_statviz" >}} | `1.1` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgstatviz` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pgsampler" >}} {{< ext "pgmonitor" >}} {{< ext "pg_mon" >}} {{< ext "timescaledb" >}} {{< ext "town" >}} {{< ext "pg_stl" >}} |

> [!Note] Cataloged but hidden from default package groups. GitHub release and control are 1.1 while PGXN still serves 1.0. PGDG DEB 1.1 covers active PG14-18 except Ubuntu 22.04 and recommends the separate Python utility, so a normal APT install can pull its Python stack. PGDG RPM remains at 0.9, lacks PG17, and provides PG18 only on EL10; its metadata declares no PostgreSQL dependency, labels GPLv2+ although upstream uses the PostgreSQL License, and describes a CLI although the subpackage contains only extension SQL and control files. The extension itself is pure SQL and PL/pgSQL and needs no preload.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_statviz` | `plpgsql` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.9` | {{< bg "18" "pg_statviz_extension_18" "green" >}} {{< bg "17" "pg_statviz_extension_17" "red" >}} {{< bg "16" "pg_statviz_extension_16" "green" >}} {{< bg "15" "pg_statviz_extension_15" "green" >}} {{< bg "14" "pg_statviz_extension_14" "green" >}} | `pg_statviz_extension_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1` | {{< bg "18" "postgresql-18-statviz" "green" >}} {{< bg "17" "postgresql-17-statviz" "green" >}} {{< bg "16" "postgresql-16-statviz" "green" >}} {{< bg "15" "postgresql-15-statviz" "green" >}} {{< bg "14" "postgresql-14-statviz" "green" >}} | `postgresql-$v-statviz` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "MISS" "pg_statviz_extension_18 : MISS 0" "red" >}} | {{< bg "MISS" "pg_statviz_extension_17 : MISS 0" "red" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_14 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "MISS" "pg_statviz_extension_18 : MISS 0" "red" >}} | {{< bg "MISS" "pg_statviz_extension_17 : MISS 0" "red" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "MISS" "pg_statviz_extension_18 : MISS 0" "red" >}} | {{< bg "MISS" "pg_statviz_extension_17 : MISS 0" "red" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_14 : AVAIL 5" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "MISS" "pg_statviz_extension_18 : MISS 0" "red" >}} | {{< bg "MISS" "pg_statviz_extension_17 : MISS 0" "red" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_14 : AVAIL 5" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_18 : AVAIL 3" "blue" >}} | {{< bg "MISS" "pg_statviz_extension_17 : MISS 0" "red" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_14 : AVAIL 4" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_18 : AVAIL 2" "blue" >}} | {{< bg "MISS" "pg_statviz_extension_17 : MISS 0" "red" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.9" "pg_statviz_extension_14 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.1" "postgresql-18-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-statviz : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.1" "postgresql-18-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-statviz : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.1" "postgresql-18-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-statviz : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.1" "postgresql-18-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-statviz : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "MISS" "postgresql-18-statviz : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-17-statviz : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-16-statviz : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-15-statviz : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-14-statviz : MISS 0" "red" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "MISS" "postgresql-18-statviz : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-17-statviz : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-16-statviz : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-15-statviz : MISS 0" "red" >}} | {{< bg "MISS" "postgresql-14-statviz : MISS 0" "red" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.1" "postgresql-18-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-statviz : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.1" "postgresql-18-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-statviz : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.1" "postgresql-18-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-statviz : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.1" "postgresql-18-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-17-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-16-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-15-statviz : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1" "postgresql-14-statviz : AVAIL 3" "blue" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statviz_extension_18` | `0.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.7 KiB | [pg_statviz_extension_18-0.9-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_statviz_extension_18-0.9-1PGDG.rhel10.2.noarch.rpm) |
| `pg_statviz_extension_18` | `0.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.7 KiB | [pg_statviz_extension_18-0.9-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_statviz_extension_18-0.9-1PGDG.rhel10.1.noarch.rpm) |
| `pg_statviz_extension_18` | `0.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.1 KiB | [pg_statviz_extension_18-0.9-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_statviz_extension_18-0.9-1PGDG.rhel10.0.noarch.rpm) |
| `pg_statviz_extension_18` | `0.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.7 KiB | [pg_statviz_extension_18-0.9-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_statviz_extension_18-0.9-1PGDG.rhel10.1.noarch.rpm) |
| `pg_statviz_extension_18` | `0.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.7 KiB | [pg_statviz_extension_18-0.9-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_statviz_extension_18-0.9-1PGDG.rhel10.0.noarch.rpm) |
| `postgresql-18-statviz` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.8 KiB | [postgresql-18-statviz_1.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.1-1.pgdg12+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-18-statviz_1.0-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-2.pgdg12+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-18-statviz_1.0-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-1.pgdg12+1_all.deb) |
| `postgresql-18-statviz` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-18-statviz_1.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.1-1.pgdg12+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.7 KiB | [postgresql-18-statviz_1.0-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-2.pgdg12+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-18-statviz_1.0-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-1.pgdg12+1_all.deb) |
| `postgresql-18-statviz` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-18-statviz_1.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.1-1.pgdg13+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-18-statviz_1.0-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-2.pgdg13+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-18-statviz_1.0-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-1.pgdg13+1_all.deb) |
| `postgresql-18-statviz` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.8 KiB | [postgresql-18-statviz_1.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.1-1.pgdg13+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-18-statviz_1.0-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-2.pgdg13+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-18-statviz_1.0-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-1.pgdg13+1_all.deb) |
| `postgresql-18-statviz` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-18-statviz_1.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.1-1.pgdg24.04+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-18-statviz_1.0-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-2.pgdg24.04+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.6 KiB | [postgresql-18-statviz_1.0-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-1.pgdg24.04+1_all.deb) |
| `postgresql-18-statviz` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.8 KiB | [postgresql-18-statviz_1.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.1-1.pgdg24.04+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-18-statviz_1.0-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-2.pgdg24.04+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-18-statviz_1.0-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-1.pgdg24.04+1_all.deb) |
| `postgresql-18-statviz` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.8 KiB | [postgresql-18-statviz_1.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.1-1.pgdg26.04+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.7 KiB | [postgresql-18-statviz_1.0-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-2.pgdg26.04+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.6 KiB | [postgresql-18-statviz_1.0-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-1.pgdg26.04+1_all.deb) |
| `postgresql-18-statviz` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-18-statviz_1.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.1-1.pgdg26.04+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.7 KiB | [postgresql-18-statviz_1.0-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-2.pgdg26.04+1_all.deb) |
| `postgresql-18-statviz` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.6 KiB | [postgresql-18-statviz_1.0-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-18-statviz_1.0-1.pgdg26.04+1_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-17-statviz` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.8 KiB | [postgresql-17-statviz_1.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.1-1.pgdg12+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-17-statviz_1.0-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-2.pgdg12+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-17-statviz_1.0-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-1.pgdg12+1_all.deb) |
| `postgresql-17-statviz` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-17-statviz_1.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.1-1.pgdg12+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.7 KiB | [postgresql-17-statviz_1.0-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-2.pgdg12+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-17-statviz_1.0-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-1.pgdg12+1_all.deb) |
| `postgresql-17-statviz` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-17-statviz_1.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.1-1.pgdg13+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-17-statviz_1.0-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-2.pgdg13+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-17-statviz_1.0-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-1.pgdg13+1_all.deb) |
| `postgresql-17-statviz` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.8 KiB | [postgresql-17-statviz_1.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.1-1.pgdg13+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-17-statviz_1.0-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-2.pgdg13+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-17-statviz_1.0-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-1.pgdg13+1_all.deb) |
| `postgresql-17-statviz` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-17-statviz_1.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.1-1.pgdg24.04+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-17-statviz_1.0-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-2.pgdg24.04+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.6 KiB | [postgresql-17-statviz_1.0-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-1.pgdg24.04+1_all.deb) |
| `postgresql-17-statviz` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.8 KiB | [postgresql-17-statviz_1.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.1-1.pgdg24.04+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-17-statviz_1.0-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-2.pgdg24.04+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-17-statviz_1.0-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-1.pgdg24.04+1_all.deb) |
| `postgresql-17-statviz` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.8 KiB | [postgresql-17-statviz_1.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.1-1.pgdg26.04+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.7 KiB | [postgresql-17-statviz_1.0-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-2.pgdg26.04+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.6 KiB | [postgresql-17-statviz_1.0-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-1.pgdg26.04+1_all.deb) |
| `postgresql-17-statviz` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-17-statviz_1.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.1-1.pgdg26.04+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.7 KiB | [postgresql-17-statviz_1.0-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-2.pgdg26.04+1_all.deb) |
| `postgresql-17-statviz` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.6 KiB | [postgresql-17-statviz_1.0-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-17-statviz_1.0-1.pgdg26.04+1_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statviz_extension_16` | `0.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.2 KiB | [pg_statviz_extension_16-0.9-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_statviz_extension_16-0.9-1PGDG.rhel8.10.noarch.rpm) |
| `pg_statviz_extension_16` | `0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 11.8 KiB | [pg_statviz_extension_16-0.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_statviz_extension_16-0.6-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_16` | `0.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.2 KiB | [pg_statviz_extension_16-0.9-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_statviz_extension_16-0.9-1PGDG.rhel8.10.noarch.rpm) |
| `pg_statviz_extension_16` | `0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 11.8 KiB | [pg_statviz_extension_16-0.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_statviz_extension_16-0.6-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_16` | `0.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.6 KiB | [pg_statviz_extension_16-0.9-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_statviz_extension_16-0.9-1PGDG.rhel9.7.noarch.rpm) |
| `pg_statviz_extension_16` | `0.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.6 KiB | [pg_statviz_extension_16-0.9-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_statviz_extension_16-0.9-1PGDG.rhel9.6.noarch.rpm) |
| `pg_statviz_extension_16` | `0.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.8 KiB | [pg_statviz_extension_16-0.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_statviz_extension_16-0.6-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_16` | `0.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.5 KiB | [pg_statviz_extension_16-0.9-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_statviz_extension_16-0.9-1PGDG.rhel9.7.noarch.rpm) |
| `pg_statviz_extension_16` | `0.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.5 KiB | [pg_statviz_extension_16-0.9-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_statviz_extension_16-0.9-1PGDG.rhel9.6.noarch.rpm) |
| `pg_statviz_extension_16` | `0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.6 KiB | [pg_statviz_extension_16-0.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_statviz_extension_16-0.6-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_16` | `0.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.7 KiB | [pg_statviz_extension_16-0.9-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_statviz_extension_16-0.9-1PGDG.rhel10.2.noarch.rpm) |
| `pg_statviz_extension_16` | `0.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.7 KiB | [pg_statviz_extension_16-0.9-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_statviz_extension_16-0.9-1PGDG.rhel10.1.noarch.rpm) |
| `pg_statviz_extension_16` | `0.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.1 KiB | [pg_statviz_extension_16-0.9-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_statviz_extension_16-0.9-1PGDG.rhel10.0.noarch.rpm) |
| `pg_statviz_extension_16` | `0.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.3 KiB | [pg_statviz_extension_16-0.6-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_statviz_extension_16-0.6-1PGDG.rhel10.noarch.rpm) |
| `pg_statviz_extension_16` | `0.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.7 KiB | [pg_statviz_extension_16-0.9-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_statviz_extension_16-0.9-1PGDG.rhel10.1.noarch.rpm) |
| `pg_statviz_extension_16` | `0.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.7 KiB | [pg_statviz_extension_16-0.9-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_statviz_extension_16-0.9-1PGDG.rhel10.0.noarch.rpm) |
| `pg_statviz_extension_16` | `0.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.2 KiB | [pg_statviz_extension_16-0.6-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_statviz_extension_16-0.6-1PGDG.rhel10.noarch.rpm) |
| `postgresql-16-statviz` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.8 KiB | [postgresql-16-statviz_1.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.1-1.pgdg12+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-16-statviz_1.0-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-2.pgdg12+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-16-statviz_1.0-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-1.pgdg12+1_all.deb) |
| `postgresql-16-statviz` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-16-statviz_1.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.1-1.pgdg12+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.7 KiB | [postgresql-16-statviz_1.0-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-2.pgdg12+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-16-statviz_1.0-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-1.pgdg12+1_all.deb) |
| `postgresql-16-statviz` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-16-statviz_1.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.1-1.pgdg13+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-16-statviz_1.0-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-2.pgdg13+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-16-statviz_1.0-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-1.pgdg13+1_all.deb) |
| `postgresql-16-statviz` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.8 KiB | [postgresql-16-statviz_1.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.1-1.pgdg13+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-16-statviz_1.0-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-2.pgdg13+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-16-statviz_1.0-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-1.pgdg13+1_all.deb) |
| `postgresql-16-statviz` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-16-statviz_1.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.1-1.pgdg24.04+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-16-statviz_1.0-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-2.pgdg24.04+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.6 KiB | [postgresql-16-statviz_1.0-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-1.pgdg24.04+1_all.deb) |
| `postgresql-16-statviz` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.8 KiB | [postgresql-16-statviz_1.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.1-1.pgdg24.04+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-16-statviz_1.0-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-2.pgdg24.04+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-16-statviz_1.0-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-1.pgdg24.04+1_all.deb) |
| `postgresql-16-statviz` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.8 KiB | [postgresql-16-statviz_1.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.1-1.pgdg26.04+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.7 KiB | [postgresql-16-statviz_1.0-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-2.pgdg26.04+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.6 KiB | [postgresql-16-statviz_1.0-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-1.pgdg26.04+1_all.deb) |
| `postgresql-16-statviz` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-16-statviz_1.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.1-1.pgdg26.04+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.7 KiB | [postgresql-16-statviz_1.0-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-2.pgdg26.04+1_all.deb) |
| `postgresql-16-statviz` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.6 KiB | [postgresql-16-statviz_1.0-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-16-statviz_1.0-1.pgdg26.04+1_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statviz_extension_15` | `0.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.2 KiB | [pg_statviz_extension_15-0.9-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_statviz_extension_15-0.9-1PGDG.rhel8.10.noarch.rpm) |
| `pg_statviz_extension_15` | `0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 11.8 KiB | [pg_statviz_extension_15-0.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_statviz_extension_15-0.6-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_15` | `0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 11.8 KiB | [pg_statviz_extension_15-0.5-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_statviz_extension_15-0.5-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_15` | `0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 11.5 KiB | [pg_statviz_extension_15-0.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_statviz_extension_15-0.4-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_15` | `0.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.2 KiB | [pg_statviz_extension_15-0.9-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_statviz_extension_15-0.9-1PGDG.rhel8.10.noarch.rpm) |
| `pg_statviz_extension_15` | `0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 11.8 KiB | [pg_statviz_extension_15-0.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_statviz_extension_15-0.6-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_15` | `0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 11.7 KiB | [pg_statviz_extension_15-0.5-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_statviz_extension_15-0.5-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_15` | `0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 11.5 KiB | [pg_statviz_extension_15-0.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_statviz_extension_15-0.4-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_15` | `0.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.6 KiB | [pg_statviz_extension_15-0.9-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_statviz_extension_15-0.9-1PGDG.rhel9.7.noarch.rpm) |
| `pg_statviz_extension_15` | `0.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.6 KiB | [pg_statviz_extension_15-0.9-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_statviz_extension_15-0.9-1PGDG.rhel9.6.noarch.rpm) |
| `pg_statviz_extension_15` | `0.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.8 KiB | [pg_statviz_extension_15-0.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_statviz_extension_15-0.6-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_15` | `0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.7 KiB | [pg_statviz_extension_15-0.5-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_statviz_extension_15-0.5-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_15` | `0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.5 KiB | [pg_statviz_extension_15-0.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_statviz_extension_15-0.4-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_15` | `0.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.5 KiB | [pg_statviz_extension_15-0.9-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_statviz_extension_15-0.9-1PGDG.rhel9.7.noarch.rpm) |
| `pg_statviz_extension_15` | `0.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.5 KiB | [pg_statviz_extension_15-0.9-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_statviz_extension_15-0.9-1PGDG.rhel9.6.noarch.rpm) |
| `pg_statviz_extension_15` | `0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.6 KiB | [pg_statviz_extension_15-0.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_statviz_extension_15-0.6-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_15` | `0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.5 KiB | [pg_statviz_extension_15-0.5-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_statviz_extension_15-0.5-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_15` | `0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.3 KiB | [pg_statviz_extension_15-0.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_statviz_extension_15-0.4-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_15` | `0.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.7 KiB | [pg_statviz_extension_15-0.9-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_statviz_extension_15-0.9-1PGDG.rhel10.2.noarch.rpm) |
| `pg_statviz_extension_15` | `0.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.7 KiB | [pg_statviz_extension_15-0.9-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_statviz_extension_15-0.9-1PGDG.rhel10.1.noarch.rpm) |
| `pg_statviz_extension_15` | `0.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.1 KiB | [pg_statviz_extension_15-0.9-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_statviz_extension_15-0.9-1PGDG.rhel10.0.noarch.rpm) |
| `pg_statviz_extension_15` | `0.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.3 KiB | [pg_statviz_extension_15-0.6-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_statviz_extension_15-0.6-1PGDG.rhel10.noarch.rpm) |
| `pg_statviz_extension_15` | `0.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.7 KiB | [pg_statviz_extension_15-0.9-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_statviz_extension_15-0.9-1PGDG.rhel10.1.noarch.rpm) |
| `pg_statviz_extension_15` | `0.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.7 KiB | [pg_statviz_extension_15-0.9-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_statviz_extension_15-0.9-1PGDG.rhel10.0.noarch.rpm) |
| `pg_statviz_extension_15` | `0.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.2 KiB | [pg_statviz_extension_15-0.6-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_statviz_extension_15-0.6-1PGDG.rhel10.noarch.rpm) |
| `postgresql-15-statviz` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.8 KiB | [postgresql-15-statviz_1.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.1-1.pgdg12+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-15-statviz_1.0-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-2.pgdg12+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-15-statviz_1.0-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-1.pgdg12+1_all.deb) |
| `postgresql-15-statviz` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-15-statviz_1.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.1-1.pgdg12+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.7 KiB | [postgresql-15-statviz_1.0-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-2.pgdg12+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-15-statviz_1.0-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-1.pgdg12+1_all.deb) |
| `postgresql-15-statviz` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-15-statviz_1.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.1-1.pgdg13+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-15-statviz_1.0-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-2.pgdg13+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-15-statviz_1.0-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-1.pgdg13+1_all.deb) |
| `postgresql-15-statviz` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.8 KiB | [postgresql-15-statviz_1.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.1-1.pgdg13+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-15-statviz_1.0-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-2.pgdg13+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-15-statviz_1.0-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-1.pgdg13+1_all.deb) |
| `postgresql-15-statviz` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-15-statviz_1.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.1-1.pgdg24.04+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-15-statviz_1.0-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-2.pgdg24.04+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.6 KiB | [postgresql-15-statviz_1.0-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-1.pgdg24.04+1_all.deb) |
| `postgresql-15-statviz` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.8 KiB | [postgresql-15-statviz_1.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.1-1.pgdg24.04+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-15-statviz_1.0-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-2.pgdg24.04+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-15-statviz_1.0-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-1.pgdg24.04+1_all.deb) |
| `postgresql-15-statviz` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.8 KiB | [postgresql-15-statviz_1.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.1-1.pgdg26.04+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.7 KiB | [postgresql-15-statviz_1.0-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-2.pgdg26.04+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.6 KiB | [postgresql-15-statviz_1.0-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-1.pgdg26.04+1_all.deb) |
| `postgresql-15-statviz` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-15-statviz_1.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.1-1.pgdg26.04+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.7 KiB | [postgresql-15-statviz_1.0-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-2.pgdg26.04+1_all.deb) |
| `postgresql-15-statviz` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.6 KiB | [postgresql-15-statviz_1.0-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-15-statviz_1.0-1.pgdg26.04+1_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statviz_extension_14` | `0.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.2 KiB | [pg_statviz_extension_14-0.9-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_statviz_extension_14-0.9-1PGDG.rhel8.10.noarch.rpm) |
| `pg_statviz_extension_14` | `0.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 11.8 KiB | [pg_statviz_extension_14-0.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_statviz_extension_14-0.6-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_14` | `0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 11.8 KiB | [pg_statviz_extension_14-0.5-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_statviz_extension_14-0.5-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_14` | `0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 11.5 KiB | [pg_statviz_extension_14-0.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_statviz_extension_14-0.4-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_14` | `0.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 15.2 KiB | [pg_statviz_extension_14-0.9-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_statviz_extension_14-0.9-1PGDG.rhel8.10.noarch.rpm) |
| `pg_statviz_extension_14` | `0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 11.8 KiB | [pg_statviz_extension_14-0.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_statviz_extension_14-0.6-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_14` | `0.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 11.7 KiB | [pg_statviz_extension_14-0.5-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_statviz_extension_14-0.5-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_14` | `0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 11.5 KiB | [pg_statviz_extension_14-0.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_statviz_extension_14-0.4-1PGDG.rhel8.noarch.rpm) |
| `pg_statviz_extension_14` | `0.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.6 KiB | [pg_statviz_extension_14-0.9-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_statviz_extension_14-0.9-1PGDG.rhel9.7.noarch.rpm) |
| `pg_statviz_extension_14` | `0.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.6 KiB | [pg_statviz_extension_14-0.9-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_statviz_extension_14-0.9-1PGDG.rhel9.6.noarch.rpm) |
| `pg_statviz_extension_14` | `0.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.8 KiB | [pg_statviz_extension_14-0.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_statviz_extension_14-0.6-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_14` | `0.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.7 KiB | [pg_statviz_extension_14-0.5-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_statviz_extension_14-0.5-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_14` | `0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.5 KiB | [pg_statviz_extension_14-0.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_statviz_extension_14-0.4-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_14` | `0.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.5 KiB | [pg_statviz_extension_14-0.9-1PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_statviz_extension_14-0.9-1PGDG.rhel9.7.noarch.rpm) |
| `pg_statviz_extension_14` | `0.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.5 KiB | [pg_statviz_extension_14-0.9-1PGDG.rhel9.6.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_statviz_extension_14-0.9-1PGDG.rhel9.6.noarch.rpm) |
| `pg_statviz_extension_14` | `0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.6 KiB | [pg_statviz_extension_14-0.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_statviz_extension_14-0.6-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_14` | `0.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.5 KiB | [pg_statviz_extension_14-0.5-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_statviz_extension_14-0.5-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_14` | `0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 11.3 KiB | [pg_statviz_extension_14-0.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_statviz_extension_14-0.4-1PGDG.rhel9.noarch.rpm) |
| `pg_statviz_extension_14` | `0.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.7 KiB | [pg_statviz_extension_14-0.9-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_statviz_extension_14-0.9-1PGDG.rhel10.2.noarch.rpm) |
| `pg_statviz_extension_14` | `0.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.7 KiB | [pg_statviz_extension_14-0.9-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_statviz_extension_14-0.9-1PGDG.rhel10.1.noarch.rpm) |
| `pg_statviz_extension_14` | `0.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.1 KiB | [pg_statviz_extension_14-0.9-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_statviz_extension_14-0.9-1PGDG.rhel10.0.noarch.rpm) |
| `pg_statviz_extension_14` | `0.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.3 KiB | [pg_statviz_extension_14-0.6-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_statviz_extension_14-0.6-1PGDG.rhel10.noarch.rpm) |
| `pg_statviz_extension_14` | `0.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.7 KiB | [pg_statviz_extension_14-0.9-1PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_statviz_extension_14-0.9-1PGDG.rhel10.1.noarch.rpm) |
| `pg_statviz_extension_14` | `0.9` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.7 KiB | [pg_statviz_extension_14-0.9-1PGDG.rhel10.0.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_statviz_extension_14-0.9-1PGDG.rhel10.0.noarch.rpm) |
| `pg_statviz_extension_14` | `0.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.2 KiB | [pg_statviz_extension_14-0.6-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_statviz_extension_14-0.6-1PGDG.rhel10.noarch.rpm) |
| `postgresql-14-statviz` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.8 KiB | [postgresql-14-statviz_1.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.1-1.pgdg12+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-14-statviz_1.0-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-2.pgdg12+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-14-statviz_1.0-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-1.pgdg12+1_all.deb) |
| `postgresql-14-statviz` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-14-statviz_1.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.1-1.pgdg12+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.7 KiB | [postgresql-14-statviz_1.0-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-2.pgdg12+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-14-statviz_1.0-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-1.pgdg12+1_all.deb) |
| `postgresql-14-statviz` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-14-statviz_1.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.1-1.pgdg13+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-14-statviz_1.0-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-2.pgdg13+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-14-statviz_1.0-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-1.pgdg13+1_all.deb) |
| `postgresql-14-statviz` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.8 KiB | [postgresql-14-statviz_1.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.1-1.pgdg13+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-14-statviz_1.0-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-2.pgdg13+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-14-statviz_1.0-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-1.pgdg13+1_all.deb) |
| `postgresql-14-statviz` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-14-statviz_1.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.1-1.pgdg24.04+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-14-statviz_1.0-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-2.pgdg24.04+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.6 KiB | [postgresql-14-statviz_1.0-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-1.pgdg24.04+1_all.deb) |
| `postgresql-14-statviz` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.8 KiB | [postgresql-14-statviz_1.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.1-1.pgdg24.04+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-14-statviz_1.0-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-2.pgdg24.04+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-14-statviz_1.0-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-1.pgdg24.04+1_all.deb) |
| `postgresql-14-statviz` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.8 KiB | [postgresql-14-statviz_1.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.1-1.pgdg26.04+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.7 KiB | [postgresql-14-statviz_1.0-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-2.pgdg26.04+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.6 KiB | [postgresql-14-statviz_1.0-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-1.pgdg26.04+1_all.deb) |
| `postgresql-14-statviz` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-14-statviz_1.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.1-1.pgdg26.04+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.7 KiB | [postgresql-14-statviz_1.0-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-2.pgdg26.04+1_all.deb) |
| `postgresql-14-statviz` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.6 KiB | [postgresql-14-statviz_1.0-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-statviz/postgresql-14-statviz_1.0-1.pgdg26.04+1_all.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/vyruss/pg_statviz" title="Repository" icon="github" subtitle="github.com/vyruss/pg_statviz" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_statviz;		# install via package name, for the active PG version

pig install pg_statviz -v 18;   # install for PG 18
pig install pg_statviz -v 17;   # install for PG 17
pig install pg_statviz -v 16;   # install for PG 16
pig install pg_statviz -v 15;   # install for PG 15
pig install pg_statviz -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_statviz CASCADE; -- requires plpgsql
```

## Usage

Sources:

- [pg_statviz v1.1 release](https://github.com/vyruss/pg_statviz/releases/tag/v1.1)
- [pg_statviz v1.1 README](https://github.com/vyruss/pg_statviz/blob/v1.1/README.md)
- [pg_statviz v1.1 installation SQL](https://github.com/vyruss/pg_statviz/blob/v1.1/pg_statviz--1.1.sql)
- [pg_statviz v1.1 control file](https://github.com/vyruss/pg_statviz/blob/v1.1/pg_statviz.control)
- [pg_statviz v1.1 metadata](https://github.com/vyruss/pg_statviz/blob/v1.1/META.json)
- [pg_statviz v1.1 Python package metadata](https://github.com/vyruss/pg_statviz/blob/v1.1/pyproject.toml)
- [pg_statviz v1.1 AI provider implementation](https://github.com/vyruss/pg_statviz/blob/v1.1/src/pg_statviz/libs/ai.py)
- [Official PGXN distribution](https://pgxn.org/dist/pg_statviz/)

`pg_statviz` v1.1 is a pure SQL and PL/pgSQL statistics snapshot extension plus a separately installed Python visualization utility. The extension stores cumulative and dynamic PostgreSQL statistics in the fixed `pgstatviz` schema; the utility reads a selected time range and generates charts or optional AI-assisted HTML reports. It requires PostgreSQL 13 or later, needs no `shared_preload_libraries`, and does not require a restart. The utility requires Python 3.11 or later.

### Capture and Retain Snapshots

Have an administrator install the extension, then let a dedicated collection role inherit `pg_monitor` and schedule `pgstatviz.snapshot()` with cron or another external job runner.

```sql
CREATE EXTENSION pg_statviz;

GRANT pg_monitor TO stats_collector;

SELECT pgstatviz.snapshot();

DELETE FROM pgstatviz.snapshots
WHERE snapshot_tstamp < CURRENT_DATE - 90;
```

Deleting parent rows cascades to the associated samples. `pgstatviz.delete_snapshots()` instead truncates the complete history. Pick an interval and retention window based on the shortest event worth observing and the resulting table growth; raw PostgreSQL counters are cumulative and can reset independently, so analyze timestamped deltas rather than treating stored values as rates.

### Stored Data and Version Boundaries

The main relations are `pgstatviz.snapshots`, `pgstatviz.buf`, `pgstatviz.conf`, `pgstatviz.conn`, `pgstatviz.db`, `pgstatviz.io`, `pgstatviz.lock`, `pgstatviz.repl`, `pgstatviz.slru`, `pgstatviz.wait`, and `pgstatviz.wal`. Samples include configuration values, connection user names and ages, replication application and slot names, waits, locks, I/O, database counters, and WAL counters. Protect the tables, dumps, charts, and reports as operational data.

Configuration is stored only when it changes, so `pgstatviz.conf` need not contain one row for every snapshot. `pg_stat_wal` data is collected on PostgreSQL 14 and later; `pg_stat_io` data is collected on PostgreSQL 16 and later, with PostgreSQL 18's byte-based fields handled separately. On older supported versions those tables remain part of the schema, but the unavailable collectors are skipped.

The extension marks its snapshot tables for extension-aware dumps. This allows history to be moved with `pg_dump`, but retention and backup size still need deliberate limits.

### Visualize a Time Range

Install the utility separately and pass normal libpq connection options. The `analyze` command runs every analysis module; individual modules such as `conn`, `io`, `wait`, and `wal` can be selected when a narrower report is sufficient.

```bash
pip install pg_statviz

pg_statviz analyze \
  -h /var/run/postgresql -d mydb -U stats_reader \
  -D 2026-08-01T00:00 2026-08-02T00:00 \
  -O /srv/pg_statviz/reports
```

Restrict database credentials and report-directory access. A visualization role needs read access to the captured schema but does not need permission to collect or delete snapshots.

### Privilege Boundary

The v1.1 installation SQL grants every member of `pg_monitor` schema usage, function execution, and `SELECT`, `INSERT`, `DELETE`, and `TRUNCATE` on all `pgstatviz` tables. Consequently, membership allows both snapshot collection and complete history removal through `pgstatviz.delete_snapshots()`; it is not a read-only visualization role.

If collection, visualization, and retention administration must be separated, revise the default grants after installation and grant only the required functions and table privileges to dedicated roles. Recheck those grants after an extension update.

### Optional AI and Cloud Data Review

Normal chart generation makes no LLM request. AI mode requires the optional `pg_statviz[ai]` dependencies and an explicit `--ai` flag. Claude is the default cloud provider and reads `ANTHROPIC_API_KEY`; Gemini reads `GOOGLE_API_KEY`; `--ai local` uses a local Ollama service. The current defaults are `claude-sonnet-4-6`, `gemini-2.5-flash`, and `gemma4:e4b`; these are implementation defaults, not a guarantee that a provider account or local runtime will continue to offer them.

```bash
pip install 'pg_statviz[ai]'

pg_statviz analyze \
  -h /var/run/postgresql -d mydb -U stats_reader \
  -D 2026-08-01T00:00 2026-08-02T00:00 \
  -O /srv/pg_statviz/reports \
  --ai gemini
```

For a cloud provider, the request can include chart images and summarized series together with the captured PostgreSQL version, primary/standby role, hostname, relevant configuration values, deterministic findings, user or role names, and replication identifiers. Treat that as an explicit operational-data export: review provider retention and regional policy, minimize the selected time range, secure generated HTML and PNG files, and use an approved outbound path. The prompt's data envelopes reduce prompt-injection risk but do not provide confidentiality, authorization, or a substitute for provider governance.
