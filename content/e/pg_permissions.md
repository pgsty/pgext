---
title: "pg_permissions"
linkTitle: "pg_permissions"
description: "view object permissions and compare them with the desired state"
weight: 5140
categories: ["ADMIN"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**pg_permissions**](https://github.com/cybertec-postgresql/pg_permissions) : view object permissions and compare them with the desired state


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5140** | {{< badge content="pg_permissions" link="https://github.com/cybertec-postgresql/pg_permissions" >}} | {{< ext "pg_permissions" >}} | `1.6` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "set_user" >}} {{< ext "pg_roast" >}} {{< ext "pg_readonly" >}} {{< ext "pgaudit" >}} {{< ext "pg_command_fw" >}} {{< ext "acl" >}} {{< ext "pg_auditor" >}} {{< ext "pg_snakeoil" >}} |

> [!Note] PGDG RPM SQL 1.3.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_permissions` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6` | {{< bg "19" "pg_permissions_19" "green" >}} {{< bg "18" "pg_permissions_18" "green" >}} {{< bg "17" "pg_permissions_17" "green" >}} {{< bg "16" "pg_permissions_16" "green" >}} {{< bg "15" "pg_permissions_15" "green" >}} | `pg_permissions_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6` | {{< bg "19" "postgresql-19-pg-permissions" "green" >}} {{< bg "18" "postgresql-18-pg-permissions" "green" >}} {{< bg "17" "postgresql-17-pg-permissions" "green" >}} {{< bg "16" "postgresql-16-pg-permissions" "green" >}} {{< bg "15" "postgresql-15-pg-permissions" "green" >}} | `postgresql-$v-pg-permissions` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_permissions_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_15 : AVAIL 7" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_permissions_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_15 : AVAIL 7" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_15 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.4.1" "pg_permissions_15 : AVAIL 7" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.4.1" "postgresql-19-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-18-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-17-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pg-permissions : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.4.1" "postgresql-19-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-18-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-17-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pg-permissions : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.4.1" "postgresql-19-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-18-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-17-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pg-permissions : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.4.1" "postgresql-19-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-18-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-17-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pg-permissions : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.4.1" "postgresql-19-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-18-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-17-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pg-permissions : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.4.1" "postgresql-19-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-18-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-17-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pg-permissions : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.4.1" "postgresql-19-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-18-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-17-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pg-permissions : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.4.1" "postgresql-19-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-18-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-17-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pg-permissions : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.4.1" "postgresql-19-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-18-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-17-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pg-permissions : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.4.1" "postgresql-19-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-18-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-17-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-16-pg-permissions : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "postgresql-15-pg-permissions : AVAIL 3" "blue" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_permissions_19` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.6 KiB | [pg_permissions_19-1.4.1-4PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/pg_permissions_19-1.4.1-4PGDG.rhel9.9.noarch.rpm) |
| `pg_permissions_19` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.5 KiB | [pg_permissions_19-1.4.1-4PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/pg_permissions_19-1.4.1-4PGDG.rhel9.9.noarch.rpm) |
| `pg_permissions_19` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.7 KiB | [pg_permissions_19-1.4.1-4PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/pg_permissions_19-1.4.1-4PGDG.rhel10.3.noarch.rpm) |
| `pg_permissions_19` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_19-1.4.1-4PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/pg_permissions_19-1.4.1-4PGDG.rhel10.3.noarch.rpm) |
| `postgresql-19-pg-permissions` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg12+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg12+2_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg12+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.5 KiB | [postgresql-19-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg12+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg12+2_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg12+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.5 KiB | [postgresql-19-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg13+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg13+2_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg13+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.4 KiB | [postgresql-19-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg13+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg13+2_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg13+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.4 KiB | [postgresql-19-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 10.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 10.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.1 KiB | [postgresql-19-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 10.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 10.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.1 KiB | [postgresql-19-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 10.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 10.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.1 KiB | [postgresql-19-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 10.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 10.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.1 KiB | [postgresql-19-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 10.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 10.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.1 KiB | [postgresql-19-pg-permissions_1.4-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4-2.pgdg26.04+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 10.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb) |
| `postgresql-19-pg-permissions` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 10.9 KiB | [postgresql-19-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb) |
| `postgresql-19-pg-permissions` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.1 KiB | [postgresql-19-pg-permissions_1.4-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-19-pg-permissions_1.4-2.pgdg26.04+1_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_permissions_18` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.4 KiB | [pg_permissions_18-1.4.1-4PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_permissions_18-1.4.1-4PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.2 KiB | [pg_permissions_18-1.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_permissions_18-1.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.0 KiB | [pg_permissions_18-1.4.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_permissions_18-1.4.1-1PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_18` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.8 KiB | [pg_permissions_18-1.4-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_permissions_18-1.4-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.3 KiB | [pg_permissions_18-1.4.1-4PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_permissions_18-1.4.1-4PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.2 KiB | [pg_permissions_18-1.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_permissions_18-1.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.0 KiB | [pg_permissions_18-1.4.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_permissions_18-1.4.1-1PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_18` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.8 KiB | [pg_permissions_18-1.4-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_permissions_18-1.4-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.6 KiB | [pg_permissions_18-1.4.1-4PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_permissions_18-1.4.1-4PGDG.rhel9.9.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.9 KiB | [pg_permissions_18-1.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_permissions_18-1.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.5 KiB | [pg_permissions_18-1.4.1-4PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_permissions_18-1.4.1-4PGDG.rhel9.9.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.9 KiB | [pg_permissions_18-1.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_permissions_18-1.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.7 KiB | [pg_permissions_18-1.4.1-4PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_permissions_18-1.4.1-4PGDG.rhel10.3.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.1 KiB | [pg_permissions_18-1.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_permissions_18-1.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_18-1.4.1-4PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_permissions_18-1.4.1-4PGDG.rhel10.3.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_18-1.4.1-4PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_permissions_18-1.4.1-4PGDG.rhel10.2.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.1 KiB | [pg_permissions_18-1.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_permissions_18-1.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_permissions_18` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.4 KiB | [pg_permissions_18-1.4.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_permissions_18-1.4.1-1PGDG.rhel10.2.noarch.rpm) |
| `pg_permissions_18` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.4 KiB | [pg_permissions_18-1.4-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_permissions_18-1.4-2PGDG.rhel10.2.noarch.rpm) |
| `pg_permissions_18` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.8 KiB | [pg_permissions_18-1.4-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_permissions_18-1.4-2PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pg-permissions` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg12+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg12+2_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg12+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.5 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg12+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg12+2_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg12+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.5 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg13+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg13+2_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg13+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.5 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg13+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg13+2_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg13+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.5 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 10.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 10.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.1 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 10.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 10.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.1 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 10.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 10.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.1 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 10.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 10.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.1 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 10.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 10.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.1 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg26.04+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 10.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb) |
| `postgresql-18-pg-permissions` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 10.9 KiB | [postgresql-18-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb) |
| `postgresql-18-pg-permissions` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.1 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg26.04+1_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_permissions_17` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.4 KiB | [pg_permissions_17-1.4.1-4PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_permissions_17-1.4.1-4PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.2 KiB | [pg_permissions_17-1.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_permissions_17-1.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.0 KiB | [pg_permissions_17-1.4.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_permissions_17-1.4.1-1PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_17` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.8 KiB | [pg_permissions_17-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_permissions_17-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_17` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.6 KiB | [pg_permissions_17-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_permissions_17-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.3 KiB | [pg_permissions_17-1.4.1-4PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_permissions_17-1.4.1-4PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.2 KiB | [pg_permissions_17-1.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_permissions_17-1.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.0 KiB | [pg_permissions_17-1.4.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_permissions_17-1.4.1-1PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_17` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.7 KiB | [pg_permissions_17-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_permissions_17-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_17` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.5 KiB | [pg_permissions_17-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_permissions_17-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.6 KiB | [pg_permissions_17-1.4.1-4PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_permissions_17-1.4.1-4PGDG.rhel9.9.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.9 KiB | [pg_permissions_17-1.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_permissions_17-1.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.5 KiB | [pg_permissions_17-1.4.1-4PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_permissions_17-1.4.1-4PGDG.rhel9.9.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.9 KiB | [pg_permissions_17-1.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_permissions_17-1.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.7 KiB | [pg_permissions_17-1.4.1-4PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_permissions_17-1.4.1-4PGDG.rhel10.3.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.1 KiB | [pg_permissions_17-1.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_permissions_17-1.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_17-1.4.1-4PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_permissions_17-1.4.1-4PGDG.rhel10.3.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_17-1.4.1-4PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_permissions_17-1.4.1-4PGDG.rhel10.2.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.1 KiB | [pg_permissions_17-1.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_permissions_17-1.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_permissions_17` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.4 KiB | [pg_permissions_17-1.4.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_permissions_17-1.4.1-1PGDG.rhel10.2.noarch.rpm) |
| `pg_permissions_17` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.4 KiB | [pg_permissions_17-1.4-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_permissions_17-1.4-2PGDG.rhel10.2.noarch.rpm) |
| `pg_permissions_17` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_17-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_permissions_17-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_17-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_permissions_17-1.3-2PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pg-permissions` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg12+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg12+2_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg12+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.5 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg12+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg12+2_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg12+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.5 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg13+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg13+2_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg13+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.5 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg13+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg13+2_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg13+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.5 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 10.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 10.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.1 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 10.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 10.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.1 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 10.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 10.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.1 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 10.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 10.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.1 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 10.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 10.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.1 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg26.04+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 10.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb) |
| `postgresql-17-pg-permissions` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 10.9 KiB | [postgresql-17-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb) |
| `postgresql-17-pg-permissions` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.1 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg26.04+1_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_permissions_16` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.4 KiB | [pg_permissions_16-1.4.1-4PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_permissions_16-1.4.1-4PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.2 KiB | [pg_permissions_16-1.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_permissions_16-1.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.0 KiB | [pg_permissions_16-1.4.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_permissions_16-1.4.1-1PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_16` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.8 KiB | [pg_permissions_16-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_permissions_16-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.6 KiB | [pg_permissions_16-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_permissions_16-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.1 KiB | [pg_permissions_16-1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_permissions_16-1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_16` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.4 KiB | [pg_permissions_16-1.1-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_permissions_16-1.1-3.rhel8.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.3 KiB | [pg_permissions_16-1.4.1-4PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_permissions_16-1.4.1-4PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.2 KiB | [pg_permissions_16-1.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_permissions_16-1.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.0 KiB | [pg_permissions_16-1.4.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_permissions_16-1.4.1-1PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_16` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.7 KiB | [pg_permissions_16-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_permissions_16-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.5 KiB | [pg_permissions_16-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_permissions_16-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.0 KiB | [pg_permissions_16-1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_permissions_16-1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_16` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.3 KiB | [pg_permissions_16-1.1-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_permissions_16-1.1-3.rhel8.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.6 KiB | [pg_permissions_16-1.4.1-4PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_permissions_16-1.4.1-4PGDG.rhel9.9.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.9 KiB | [pg_permissions_16-1.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_permissions_16-1.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.5 KiB | [pg_permissions_16-1.4.1-4PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_permissions_16-1.4.1-4PGDG.rhel9.9.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.9 KiB | [pg_permissions_16-1.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_permissions_16-1.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.7 KiB | [pg_permissions_16-1.4.1-4PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_permissions_16-1.4.1-4PGDG.rhel10.3.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.1 KiB | [pg_permissions_16-1.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_permissions_16-1.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_16-1.4.1-4PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_permissions_16-1.4.1-4PGDG.rhel10.3.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_16-1.4.1-4PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_permissions_16-1.4.1-4PGDG.rhel10.2.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.1 KiB | [pg_permissions_16-1.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_permissions_16-1.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_permissions_16` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.4 KiB | [pg_permissions_16-1.4.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_permissions_16-1.4.1-1PGDG.rhel10.2.noarch.rpm) |
| `pg_permissions_16` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.4 KiB | [pg_permissions_16-1.4-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_permissions_16-1.4-2PGDG.rhel10.2.noarch.rpm) |
| `pg_permissions_16` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_16-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_permissions_16-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_16-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_permissions_16-1.3-2PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pg-permissions` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg12+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg12+2_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg12+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.5 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg12+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg12+2_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg12+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.5 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg13+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg13+2_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg13+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.5 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg13+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg13+2_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg13+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.5 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 10.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 10.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.1 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 10.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 10.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.1 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 10.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 10.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.1 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 10.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 10.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.1 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 10.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 10.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.1 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg26.04+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 10.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb) |
| `postgresql-16-pg-permissions` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 10.9 KiB | [postgresql-16-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb) |
| `postgresql-16-pg-permissions` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.1 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg26.04+1_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_permissions_15` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.4 KiB | [pg_permissions_15-1.4.1-4PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_permissions_15-1.4.1-4PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.2 KiB | [pg_permissions_15-1.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_permissions_15-1.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.0 KiB | [pg_permissions_15-1.4.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_permissions_15-1.4.1-1PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_15` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.8 KiB | [pg_permissions_15-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_permissions_15-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.6 KiB | [pg_permissions_15-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_permissions_15-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.1 KiB | [pg_permissions_15-1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_permissions_15-1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_15` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.3 KiB | [pg_permissions_15-1.1-2.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_permissions_15-1.1-2.rhel8.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.3 KiB | [pg_permissions_15-1.4.1-4PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_permissions_15-1.4.1-4PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.2 KiB | [pg_permissions_15-1.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_permissions_15-1.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.0 KiB | [pg_permissions_15-1.4.1-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_permissions_15-1.4.1-1PGDG.rhel8.10.noarch.rpm) |
| `pg_permissions_15` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.7 KiB | [pg_permissions_15-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_permissions_15-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.5 KiB | [pg_permissions_15-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_permissions_15-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.0 KiB | [pg_permissions_15-1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_permissions_15-1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_15` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.2 KiB | [pg_permissions_15-1.1-2.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_permissions_15-1.1-2.rhel8.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.6 KiB | [pg_permissions_15-1.4.1-4PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_permissions_15-1.4.1-4PGDG.rhel9.9.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.9 KiB | [pg_permissions_15-1.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_permissions_15-1.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.5 KiB | [pg_permissions_15-1.4.1-4PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_permissions_15-1.4.1-4PGDG.rhel9.9.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.9 KiB | [pg_permissions_15-1.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_permissions_15-1.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.7 KiB | [pg_permissions_15-1.4.1-4PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_permissions_15-1.4.1-4PGDG.rhel10.3.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.1 KiB | [pg_permissions_15-1.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_permissions_15-1.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_15-1.4.1-4PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_permissions_15-1.4.1-4PGDG.rhel10.3.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_15-1.4.1-4PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_permissions_15-1.4.1-4PGDG.rhel10.2.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.1 KiB | [pg_permissions_15-1.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_permissions_15-1.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_permissions_15` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.4 KiB | [pg_permissions_15-1.4.1-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_permissions_15-1.4.1-1PGDG.rhel10.2.noarch.rpm) |
| `pg_permissions_15` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.4 KiB | [pg_permissions_15-1.4-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_permissions_15-1.4-2PGDG.rhel10.2.noarch.rpm) |
| `pg_permissions_15` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_15-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_permissions_15-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.7 KiB | [pg_permissions_15-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_permissions_15-1.3-2PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pg-permissions` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg12+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg12+2_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg12+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.5 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg12+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg12+2_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg12+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.5 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg13+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg13+2_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg13+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.5 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg13+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg13+2_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg13+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.5 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 10.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 10.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.1 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 10.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg22.04+2_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 10.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg22.04+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.1 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 10.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 10.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.1 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 10.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg24.04+2_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 10.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg24.04+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.1 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 10.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 10.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.1 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg26.04+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 10.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg26.04+2_all.deb) |
| `postgresql-15-pg-permissions` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 10.9 KiB | [postgresql-15-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4.1-1.pgdg26.04+1_all.deb) |
| `postgresql-15-pg-permissions` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.1 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg26.04+1_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/pg_permissions" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/pg_permissions" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_permissions-1.6.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_permissions;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_permissions;		# install via package name, for the active PG version

pig install pg_permissions -v 19;   # install for PG 19
pig install pg_permissions -v 18;   # install for PG 18
pig install pg_permissions -v 17;   # install for PG 17
pig install pg_permissions -v 16;   # install for PG 16
pig install pg_permissions -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_permissions;
```

## Usage

Sources:

- [README.md](https://github.com/cybertec-postgresql/pg_permissions/blob/REL_1_6/README.md)
- [pg_permissions.control](https://github.com/cybertec-postgresql/pg_permissions/blob/REL_1_6/pg_permissions.control)
- [pg_permissions--1.5--1.6.sql](https://github.com/cybertec-postgresql/pg_permissions/blob/REL_1_6/pg_permissions--1.5--1.6.sql)

`pg_permissions` compares actual privileges with a desired baseline and can generate or apply corrective grants. Version 1.6 requires PostgreSQL 15 or later. It is pure SQL, requires no preload, and is not relocatable; choose the installation schema initially. The installer needs CREATE privileges there.

### Define and Inspect a Baseline

The application role and schema in this example must already exist:

```sql
CREATE EXTENSION pg_permissions SCHEMA public;
INSERT INTO public.permission_target
  (role_name, permissions, object_type, schema_name)
VALUES ('appuser', '{SELECT,INSERT,UPDATE,DELETE}', 'TABLE', 'appschema'),
       ('appuser', '{USAGE}', 'SCHEMA', 'appschema');
SELECT * FROM public.permission_diffs();
SELECT * FROM public.permission_diff_summary ORDER BY objects DESC;
SELECT public.permission_apply();
```

A missing flag of true means a required privilege is absent; false means an extra privilege. A NULL object name covers all objects of the specified type in that schema. Audit scope is limited to targeted objects, managed roles and roles inheriting from them. Use `all_permissions` for a broader review. Superusers are omitted because they implicitly hold all privileges.

### Review Before Applying

`permission_fix_sql()` returns corrective SQL. `permission_apply()` defaults to a dry run; `permission_apply(false)` executes it using the caller's privileges. The writable `granted` column on permission views also executes GRANT or REVOKE. Restrict both baseline writes and permission-changing access.

The object views cover databases, schemas, tables, views, columns, functions and sequences. Their `source` column distinguishes ownership, direct grants, PUBLIC and inherited-role privileges. `PUBLIC` can be an explicit target grantee. `permission_target_snapshot(role, schema)` adopts current grants as the baseline; this changes the desired state and should be deliberate.

`default_acl_target`, `default_acl_permissions` and `default_acl_diffs()` separately handle ALTER DEFAULT PRIVILEGES. A default privilege affects future objects; it does not repair existing grants. Baseline tables are readable by PUBLIC but writable only by the extension owner and explicitly authorized roles. `role_dependencies()` inspects dependencies; `role_drop()` can reassign ownership and drop roles and warrants a separate review.

### Upgrade Boundary

Follow the version-by-version update scripts. For an installed 1.4 database:

```sql
ALTER EXTENSION pg_permissions UPDATE TO '1.5';
ALTER EXTENSION pg_permissions UPDATE TO '1.6';
```

Version 1.6 restores the missing 1.2-to-1.3 path. Its 1.5-to-1.6 migration rewrites the baseline: empty permission lists are removed, COLUMN entries without a column become table-level entries, and duplicates are merged. Back up and review the desired state before upgrading. MAINTAIN privileges are rejected before PostgreSQL 17.
