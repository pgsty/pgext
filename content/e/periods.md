---
title: "periods"
linkTitle: "periods"
description: "Provide Standard SQL functionality for PERIODs and SYSTEM VERSIONING"
weight: 1030
categories: ["TIME"]
width: full
---

[**periods**](https://github.com/xocolatl/periods) : Provide Standard SQL functionality for PERIODs and SYSTEM VERSIONING


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1030** | {{< badge content="periods" link="https://github.com/xocolatl/periods" >}} | {{< ext "periods" >}} | `1.2.3` | {{< category "TIME" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "btree_gist" >}} |
|   **See Also**    | {{< ext "timescaledb_toolkit" >}} {{< ext "timescaledb" >}} {{< ext "timeseries" >}} {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "table_version" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `periods` | `btree_gist` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2.3` | {{< bg "18" "periods_18" "green" >}} {{< bg "17" "periods_17" "green" >}} {{< bg "16" "periods_16" "green" >}} {{< bg "15" "periods_15" "green" >}} {{< bg "14" "periods_14" "green" >}} | `periods_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2.3` | {{< bg "18" "postgresql-18-periods" "green" >}} {{< bg "17" "postgresql-17-periods" "green" >}} {{< bg "16" "postgresql-16-periods" "green" >}} {{< bg "15" "postgresql-15-periods" "green" >}} {{< bg "14" "postgresql-14-periods" "green" >}} | `postgresql-$v-periods` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "periods_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_14 : AVAIL 3" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "periods_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_14 : AVAIL 3" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "periods_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "periods_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_14 : AVAIL 3" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "periods_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_14 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "periods_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.3" "periods_14 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.2.3" "postgresql-18-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-17-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-16-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-15-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-14-periods : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.2.3" "postgresql-18-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-17-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-16-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-15-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-14-periods : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.2.3" "postgresql-18-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-17-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-16-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-15-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-14-periods : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.2.3" "postgresql-18-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-17-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-16-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-15-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-14-periods : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.2.3" "postgresql-18-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-17-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-16-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-15-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-14-periods : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.2.3" "postgresql-18-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-17-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-16-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-15-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-14-periods : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.2.3" "postgresql-18-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-17-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-16-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-15-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-14-periods : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.2.3" "postgresql-18-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-17-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-16-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-15-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-14-periods : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.2.3" "postgresql-18-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-17-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-16-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-15-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-14-periods : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.2.3" "postgresql-18-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-17-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-16-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-15-periods : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.3" "postgresql-14-periods : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `periods_18` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.7 KiB | [periods_18-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/periods_18-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `periods_18` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.6 KiB | [periods_18-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/periods_18-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `periods_18` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.0 KiB | [periods_18-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/periods_18-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `periods_18` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 42.0 KiB | [periods_18-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/periods_18-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `periods_18` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.1 KiB | [periods_18-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/periods_18-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `periods_18` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 42.0 KiB | [periods_18-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/periods_18-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-periods` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 47.0 KiB | [postgresql-18-periods_1.2.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg12+1_amd64.deb) |
| `postgresql-18-periods` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 46.4 KiB | [postgresql-18-periods_1.2.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg12+1_arm64.deb) |
| `postgresql-18-periods` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 47.2 KiB | [postgresql-18-periods_1.2.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg13+1_amd64.deb) |
| `postgresql-18-periods` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 46.4 KiB | [postgresql-18-periods_1.2.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg13+1_arm64.deb) |
| `postgresql-18-periods` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 46.0 KiB | [postgresql-18-periods_1.2.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-periods` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 45.5 KiB | [postgresql-18-periods_1.2.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-periods` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 46.0 KiB | [postgresql-18-periods_1.2.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-periods` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.4 KiB | [postgresql-18-periods_1.2.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-periods` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 46.0 KiB | [postgresql-18-periods_1.2.3-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-periods` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 45.4 KiB | [postgresql-18-periods_1.2.3-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-18-periods_1.2.3-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `periods_17` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.7 KiB | [periods_17-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/periods_17-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `periods_17` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.4 KiB | [periods_17-1.2.2-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/periods_17-1.2.2-3PGDG.rhel8.x86_64.rpm) |
| `periods_17` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.5 KiB | [periods_17-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/periods_17-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `periods_17` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.1 KiB | [periods_17-1.2.2-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/periods_17-1.2.2-3PGDG.rhel8.aarch64.rpm) |
| `periods_17` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.0 KiB | [periods_17-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/periods_17-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `periods_17` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.4 KiB | [periods_17-1.2.2-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/periods_17-1.2.2-3PGDG.rhel9.x86_64.rpm) |
| `periods_17` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.8 KiB | [periods_17-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/periods_17-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `periods_17` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.1 KiB | [periods_17-1.2.2-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/periods_17-1.2.2-3PGDG.rhel9.aarch64.rpm) |
| `periods_17` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.0 KiB | [periods_17-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/periods_17-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `periods_17` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.7 KiB | [periods_17-1.2.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/periods_17-1.2.2-4PGDG.rhel10.x86_64.rpm) |
| `periods_17` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 42.0 KiB | [periods_17-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/periods_17-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `periods_17` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.5 KiB | [periods_17-1.2.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/periods_17-1.2.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-periods` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 47.0 KiB | [postgresql-17-periods_1.2.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg12+1_amd64.deb) |
| `postgresql-17-periods` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 46.3 KiB | [postgresql-17-periods_1.2.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg12+1_arm64.deb) |
| `postgresql-17-periods` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 47.2 KiB | [postgresql-17-periods_1.2.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg13+1_amd64.deb) |
| `postgresql-17-periods` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 46.5 KiB | [postgresql-17-periods_1.2.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg13+1_arm64.deb) |
| `postgresql-17-periods` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.1 KiB | [postgresql-17-periods_1.2.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-periods` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 49.6 KiB | [postgresql-17-periods_1.2.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-periods` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 46.1 KiB | [postgresql-17-periods_1.2.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-periods` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.3 KiB | [postgresql-17-periods_1.2.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-periods` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 46.0 KiB | [postgresql-17-periods_1.2.3-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-periods` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 45.4 KiB | [postgresql-17-periods_1.2.3-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-17-periods_1.2.3-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `periods_16` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.7 KiB | [periods_16-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/periods_16-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `periods_16` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.2 KiB | [periods_16-1.2.2-1.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/periods_16-1.2.2-1.rhel8.1.x86_64.rpm) |
| `periods_16` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.5 KiB | [periods_16-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/periods_16-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `periods_16` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.9 KiB | [periods_16-1.2.2-1.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/periods_16-1.2.2-1.rhel8.1.aarch64.rpm) |
| `periods_16` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.9 KiB | [periods_16-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/periods_16-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `periods_16` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.1 KiB | [periods_16-1.2.2-1.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/periods_16-1.2.2-1.rhel9.1.x86_64.rpm) |
| `periods_16` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.9 KiB | [periods_16-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/periods_16-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `periods_16` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [periods_16-1.2.2-1.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/periods_16-1.2.2-1.rhel9.1.aarch64.rpm) |
| `periods_16` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.0 KiB | [periods_16-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/periods_16-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `periods_16` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.7 KiB | [periods_16-1.2.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/periods_16-1.2.2-4PGDG.rhel10.x86_64.rpm) |
| `periods_16` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 42.0 KiB | [periods_16-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/periods_16-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `periods_16` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.5 KiB | [periods_16-1.2.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/periods_16-1.2.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-periods` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 47.0 KiB | [postgresql-16-periods_1.2.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg12+1_amd64.deb) |
| `postgresql-16-periods` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 46.3 KiB | [postgresql-16-periods_1.2.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg12+1_arm64.deb) |
| `postgresql-16-periods` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 47.2 KiB | [postgresql-16-periods_1.2.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg13+1_amd64.deb) |
| `postgresql-16-periods` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 46.5 KiB | [postgresql-16-periods_1.2.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg13+1_arm64.deb) |
| `postgresql-16-periods` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 49.7 KiB | [postgresql-16-periods_1.2.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-periods` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 49.2 KiB | [postgresql-16-periods_1.2.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-periods` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 46.0 KiB | [postgresql-16-periods_1.2.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-periods` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.4 KiB | [postgresql-16-periods_1.2.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-periods` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 46.0 KiB | [postgresql-16-periods_1.2.3-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-periods` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 45.4 KiB | [postgresql-16-periods_1.2.3-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-16-periods_1.2.3-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `periods_15` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.7 KiB | [periods_15-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/periods_15-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `periods_15` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.2 KiB | [periods_15-1.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/periods_15-1.2.2-1.rhel8.x86_64.rpm) |
| `periods_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.8 KiB | [periods_15-1.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/periods_15-1.2-2.rhel8.x86_64.rpm) |
| `periods_15` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.6 KiB | [periods_15-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/periods_15-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `periods_15` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.9 KiB | [periods_15-1.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/periods_15-1.2.2-1.rhel8.aarch64.rpm) |
| `periods_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 60.4 KiB | [periods_15-1.2-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/periods_15-1.2-2.rhel8.aarch64.rpm) |
| `periods_15` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.0 KiB | [periods_15-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/periods_15-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `periods_15` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.1 KiB | [periods_15-1.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/periods_15-1.2.2-1.rhel9.x86_64.rpm) |
| `periods_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 59.8 KiB | [periods_15-1.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/periods_15-1.2-2.rhel9.x86_64.rpm) |
| `periods_15` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.9 KiB | [periods_15-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/periods_15-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `periods_15` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [periods_15-1.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/periods_15-1.2.2-1.rhel9.aarch64.rpm) |
| `periods_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 59.2 KiB | [periods_15-1.2-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/periods_15-1.2-2.rhel9.aarch64.rpm) |
| `periods_15` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.0 KiB | [periods_15-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/periods_15-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `periods_15` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.7 KiB | [periods_15-1.2.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/periods_15-1.2.2-4PGDG.rhel10.x86_64.rpm) |
| `periods_15` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 42.0 KiB | [periods_15-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/periods_15-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `periods_15` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.5 KiB | [periods_15-1.2.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/periods_15-1.2.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-periods` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 46.9 KiB | [postgresql-15-periods_1.2.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg12+1_amd64.deb) |
| `postgresql-15-periods` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 46.3 KiB | [postgresql-15-periods_1.2.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg12+1_arm64.deb) |
| `postgresql-15-periods` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 47.2 KiB | [postgresql-15-periods_1.2.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg13+1_amd64.deb) |
| `postgresql-15-periods` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 46.5 KiB | [postgresql-15-periods_1.2.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg13+1_arm64.deb) |
| `postgresql-15-periods` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 49.7 KiB | [postgresql-15-periods_1.2.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-periods` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 49.3 KiB | [postgresql-15-periods_1.2.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-periods` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 46.0 KiB | [postgresql-15-periods_1.2.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-periods` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.3 KiB | [postgresql-15-periods_1.2.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-periods` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 46.0 KiB | [postgresql-15-periods_1.2.3-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-periods` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 45.4 KiB | [postgresql-15-periods_1.2.3-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-15-periods_1.2.3-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `periods_14` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.7 KiB | [periods_14-1.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/periods_14-1.2.3-1PIGSTY.el8.x86_64.rpm) |
| `periods_14` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.2 KiB | [periods_14-1.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/periods_14-1.2.2-1.rhel8.x86_64.rpm) |
| `periods_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 61.2 KiB | [periods_14-1.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/periods_14-1.2-2.rhel8.x86_64.rpm) |
| `periods_14` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.6 KiB | [periods_14-1.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/periods_14-1.2.3-1PIGSTY.el8.aarch64.rpm) |
| `periods_14` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.9 KiB | [periods_14-1.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/periods_14-1.2.2-1.rhel8.aarch64.rpm) |
| `periods_14` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 60.2 KiB | [periods_14-1.2-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/periods_14-1.2-2.rhel8.aarch64.rpm) |
| `periods_14` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.0 KiB | [periods_14-1.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/periods_14-1.2.3-1PIGSTY.el9.x86_64.rpm) |
| `periods_14` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.1 KiB | [periods_14-1.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/periods_14-1.2.2-1.rhel9.x86_64.rpm) |
| `periods_14` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.9 KiB | [periods_14-1.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/periods_14-1.2.3-1PIGSTY.el9.aarch64.rpm) |
| `periods_14` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [periods_14-1.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/periods_14-1.2.2-1.rhel9.aarch64.rpm) |
| `periods_14` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 59.0 KiB | [periods_14-1.2-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/periods_14-1.2-2.rhel9.aarch64.rpm) |
| `periods_14` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.0 KiB | [periods_14-1.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/periods_14-1.2.3-1PIGSTY.el10.x86_64.rpm) |
| `periods_14` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.7 KiB | [periods_14-1.2.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/periods_14-1.2.2-4PGDG.rhel10.x86_64.rpm) |
| `periods_14` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 42.0 KiB | [periods_14-1.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/periods_14-1.2.3-1PIGSTY.el10.aarch64.rpm) |
| `periods_14` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.5 KiB | [periods_14-1.2.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/periods_14-1.2.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-periods` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 46.9 KiB | [postgresql-14-periods_1.2.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg12+1_amd64.deb) |
| `postgresql-14-periods` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 46.2 KiB | [postgresql-14-periods_1.2.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg12+1_arm64.deb) |
| `postgresql-14-periods` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 47.1 KiB | [postgresql-14-periods_1.2.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg13+1_amd64.deb) |
| `postgresql-14-periods` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 46.5 KiB | [postgresql-14-periods_1.2.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg13+1_arm64.deb) |
| `postgresql-14-periods` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 49.6 KiB | [postgresql-14-periods_1.2.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-periods` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 49.2 KiB | [postgresql-14-periods_1.2.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-periods` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 46.0 KiB | [postgresql-14-periods_1.2.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-periods` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 45.3 KiB | [postgresql-14-periods_1.2.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-periods` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.9 KiB | [postgresql-14-periods_1.2.3-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg26.04+1_amd64.deb) |
| `postgresql-14-periods` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 45.4 KiB | [postgresql-14-periods_1.2.3-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-periods/postgresql-14-periods_1.2.3-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/xocolatl/periods" title="Repository" icon="github" subtitle="github.com/xocolatl/periods" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="periods-1.2.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg periods;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install periods;		# install via package name, for the active PG version

pig install periods -v 18;   # install for PG 18
pig install periods -v 17;   # install for PG 17
pig install periods -v 16;   # install for PG 16
pig install periods -v 15;   # install for PG 15
pig install periods -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION periods CASCADE; -- requires btree_gist
```



## Usage

> [periods: Periods and SYSTEM VERSIONING for PostgreSQL](https://github.com/xocolatl/periods)

This extension recreates the behavior defined in [SQL:2016](https://www.iso.org/standard/63556.html) (originally in SQL:2011) around periods and tables with `SYSTEM VERSIONING`. The idea is to figure out all the rules that PostgreSQL would like to adopt and to allow earlier versions of PostgreSQL to simulate the behavior once the feature is finally integrated.

### What is a period?

A period is a definition on a table which specifies a name and two columns. The period's name cannot be the same as any column name of the table.

```sql
-- Standard SQL
CREATE TABLE example (
    id bigint,
    start_date date,
    end_date date,
    PERIOD FOR validity (start_date, end_date)
);
```

Since extensions cannot modify PostgreSQL's grammar, we use functions, views, and triggers to get as close to the same thing as possible.

```sql
CREATE TABLE example (
    id bigint,
    start_date date,
    end_date date
);
SELECT periods.add_period('example', 'validity', 'start_date', 'end_date');
```

Defining a period constrains the two columns such that the start column's value must be strictly inferior to the end column's value, and that both columns be non-null. The period's value includes the start value but excludes the end value.

## Unique Constraints

Periods may be part of `PRIMARY KEY`s and `UNIQUE` constraints.

```sql
CREATE TABLE example (
    id bigint,
    start_date date,
    end_date date
);
SELECT periods.add_period('example', 'validity', 'start_date', 'end_date');
SELECT periods.add_unique_key('example', ARRAY['id'], 'validity');
```

The extension will create a unique constraint over all of the columns specified and the two columns of the period given. It will also create an exclusion constraint using gist to implement the `WITHOUT OVERLAPS` part of the constraint.

## Foreign Keys

If you can have unique keys with periods, you can also have foreign keys pointing at them.

```sql
SELECT periods.add_foreign_key('example2', 'ARRAY[ex_id]', 'validity', 'example_id_validity');
```

## Portions

The SQL standard allows syntax for updating or deleting just a portion of a period. Rows are inserted as needed for the portions not being updated or deleted.

```sql
-- Standard SQL
UPDATE example
FOR PORTION OF validity FROM '...' TO '...'
SET ...
WHERE ...;
```

This extension uses a view with an `INSTEAD OF` trigger to figure out what portion of the period you would like to modify:

```sql
UPDATE example__for_portion_of_validity
SET ...,
    start_date = ...,
    end_date = ...
WHERE ...;
```

In order to use this feature, the table must have a primary key.

## Predicates

The SQL standard provides for several predicates on periods, implemented as inlined functions:

```sql
-- "t" and "u" are tables with respective periods "p" and "q".
-- Both periods have underlying columns "s" and "e".

WHERE periods.contains(t.s, t.e, 42)            -- t.p CONTAINS 42
WHERE periods.contains(t.s, t.e, u.s, u.e)      -- t.p CONTAINS u.q
WHERE periods.equals(t.s, t.e, u.s, u.e)        -- t.p EQUALS u.q
WHERE periods.overlaps(t.s, t.e, u.s, u.e)      -- t.p OVERLAPS u.q
WHERE periods.precedes(t.s, t.e, u.s, u.e)      -- t.p PRECEDES u.q
WHERE periods.succeeds(t.s, t.e, u.s, u.e)      -- t.p SUCCEEDS u.q
WHERE periods.immediately_precedes(t.s, t.e, u.s, u.e)  -- t.p IMMEDIATELY PRECEDES u.q
WHERE periods.immediately_succeeds(t.s, t.e, u.s, u.e)  -- t.p IMMEDIATELY SUCCEEDS u.q
```


## System-Versioned Tables

### SYSTEM_TIME

If the period is named `SYSTEM_TIME`, then special rules apply. The type of the columns must be `date`, `timestamp without time zone`, or `timestamp with time zone`; and they are not modifiable by the user. This extension uses triggers to set the start column to `transaction_timestamp()` and the end column is always `'infinity'`.

***Note:*** It is generally unwise to use anything but `timestamp with time zone` because changes in the `TimeZone` configuration parameter or Daylight Savings Time changes can distort the history.

```sql
CREATE TABLE example (
    id bigint PRIMARY KEY,
    value text
);
SELECT periods.add_system_time_period('example', 'row_start', 'row_end');
```

The columns need not exist — they will be created by the extension.

### Excluding Columns

It might be desirable to prevent some columns from updating the `SYSTEM_TIME` values:

```sql
SELECT periods.add_system_time_period(
            'example',
            excluded_column_names => ARRAY['foo', 'bar']);
```

Excluded columns can be defined after the fact as well:

```sql
SELECT periods.set_system_time_period_excluded_columns(
            'example',
            ARRAY['foo', 'bar']);
```

### WITH SYSTEM VERSIONING

This special `SYSTEM_TIME` period can be used to keep track of changes in the table.

```sql
CREATE TABLE example (
    id bigint PRIMARY KEY,
    value text
);
SELECT periods.add_system_time_period('example', 'row_start', 'row_end');
SELECT periods.add_system_versioning('example');
```

This instructs the system to keep a record of all changes in the table. A separate history table is used. You can create the history table yourself and instruct the extension to use it if you want to do things like add partitioning.

### Temporal Querying

The SQL standard extends the `FROM` and `JOIN` clauses to allow specifying a point in time, or a range of time. This extension implements them through inlined functions:

```sql
SELECT * FROM t__as_of('...');                       -- FOR system_time AS OF '...'
SELECT * FROM t__from_to('...', '...');              -- FOR system_time FROM '...' TO '...'
SELECT * FROM t__between('...', '...');              -- FOR system_time BETWEEN '...' AND '...'
SELECT * FROM t__between_symmetric('...', '...');    -- FOR system_time BETWEEN SYMMETRIC '...' AND '...'
```

### Access Control

The history table as well as the helper functions all follow the ownership and access privileges of the base table. The history data is read-only. In order to trim old data, `SYSTEM VERSIONING` must be suspended:

```sql
BEGIN;
SELECT periods.drop_system_versioning('t');
GRANT DELETE ON TABLE t TO CURRENT_USER;
DELETE FROM t_history WHERE system_time_end < now() - interval '1 year';
SELECT periods.add_system_versioning('t');
COMMIT;
```

### Altering a Table with System Versioning

This extension prevents you from dropping objects while system versioning is active. The suggested way to make changes is:

```sql
BEGIN;
SELECT periods.drop_system_versioning('t');
ALTER TABLE t ...;
ALTER TABLE t_history ...;
SELECT periods.add_system_versioning('t');
COMMIT;
```

It is up to you to make sure you alter the history table in a way that is compatible with the main table. Re-activating system versioning will verify this.
