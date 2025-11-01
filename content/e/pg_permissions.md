---
title: "pg_permissions"
linkTitle: "pg_permissions"
description: "view object permissions and compare them with the desired state"
weight: 5140
categories: ["ADMIN"]
width: full
---

view object permissions and compare them with the desired state


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5140** | {{< badge content="pg_permissions" link="https://github.com/cybertec-postgresql/pg_permissions" >}} | {{< ext "pg_permissions" >}} | `1.4` | {{< category "ADMIN" >}} | {{< license "BSD 2-Clause" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_readonly" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "pg_upless" >}} {{< ext "safeupdate" >}} {{< ext "pgauditlogtofile" >}} {{< ext "credcheck" >}} {{< ext "login_hook" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_permissions" >}} | `1.4` | {{< bg "18" "pg_permissions_18" "green" >}} {{< bg "17" "pg_permissions_17" "green" >}} {{< bg "16" "pg_permissions_16" "green" >}} {{< bg "15" "pg_permissions_15" "green" >}} {{< bg "14" "pg_permissions_14" "green" >}} {{< bg "13" "pg_permissions_13" "green" >}} | `pg_permissions_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_permissions" >}} | `1.4` | {{< bg "18" "postgresql-18-pg-permissions" "green" >}} {{< bg "17" "postgresql-17-pg-permissions" "green" >}} {{< bg "16" "postgresql-16-pg-permissions" "green" >}} {{< bg "15" "postgresql-15-pg-permissions" "green" >}} {{< bg "14" "postgresql-14-pg-permissions" "green" >}} {{< bg "13" "postgresql-13-pg-permissions" "green" >}} | `postgresql-$v-pg-permissions` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.4" "pg_permissions_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_13 : AVAIL 4" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.4" "pg_permissions_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_13 : AVAIL 4" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.4" "pg_permissions_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_13 : AVAIL 4" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.4" "pg_permissions_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_13 : AVAIL 4" "blue" >}} |
|    `el10.x86_64`    | {{< bg "PGDG 1.4" "pg_permissions_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_13 : AVAIL 2" "blue" >}} |
|    `el10.aarch64`    | {{< bg "PGDG 1.4" "pg_permissions_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_permissions_13 : AVAIL 2" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 1.4" "postgresql-18-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-pg-permissions : AVAIL 2" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 1.4" "postgresql-18-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-pg-permissions : AVAIL 2" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 1.4" "postgresql-18-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-pg-permissions : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 1.4" "postgresql-18-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-pg-permissions : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 1.4" "postgresql-18-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-pg-permissions : AVAIL 2" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 1.4" "postgresql-18-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-pg-permissions : AVAIL 2" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 1.4" "postgresql-18-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-pg-permissions : AVAIL 2" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 1.4" "postgresql-18-pg-permissions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-14-pg-permissions : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-13-pg-permissions : AVAIL 2" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_permissions_18` | 1.4 | `el8.x86_64` | pgdg | 13.8 KiB | [pg_permissions_18-1.4-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_permissions_18-1.4-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_18` | 1.4 | `el8.aarch64` | pgdg | 13.8 KiB | [pg_permissions_18-1.4-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_permissions_18-1.4-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_18` | 1.4 | `el9.x86_64` | pgdg | 13.4 KiB | [pg_permissions_18-1.4-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_permissions_18-1.4-2PGDG.rhel9.noarch.rpm) |
| `pg_permissions_18` | 1.4 | `el9.aarch64` | pgdg | 13.3 KiB | [pg_permissions_18-1.4-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_permissions_18-1.4-2PGDG.rhel9.noarch.rpm) |
| `pg_permissions_18` | 1.4 | `el10.x86_64` | pgdg | 13.9 KiB | [pg_permissions_18-1.4-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_permissions_18-1.4-2PGDG.rhel10.noarch.rpm) |
| `pg_permissions_18` | 1.4 | `el10.aarch64` | pgdg | 13.8 KiB | [pg_permissions_18-1.4-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_permissions_18-1.4-2PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pg-permissions` | 1.4 | `d12.x86_64` | pgdg | 8.5 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-18-pg-permissions` | 1.4 | `d12.aarch64` | pgdg | 8.5 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-18-pg-permissions` | 1.4 | `d13.x86_64` | pgdg | 8.5 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-18-pg-permissions` | 1.4 | `d13.aarch64` | pgdg | 8.5 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-18-pg-permissions` | 1.4 | `u22.x86_64` | pgdg | 8.1 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-18-pg-permissions` | 1.4 | `u22.aarch64` | pgdg | 8.1 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-18-pg-permissions` | 1.4 | `u24.x86_64` | pgdg | 8.1 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-18-pg-permissions` | 1.4 | `u24.aarch64` | pgdg | 8.1 KiB | [postgresql-18-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-18-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_permissions_17` | 1.4 | `el8.x86_64` | pgdg | 13.8 KiB | [pg_permissions_17-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_permissions_17-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_17` | 1.3 | `el8.x86_64` | pgdg | 13.6 KiB | [pg_permissions_17-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_permissions_17-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_17` | 1.4 | `el8.aarch64` | pgdg | 13.7 KiB | [pg_permissions_17-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_permissions_17-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_17` | 1.3 | `el8.aarch64` | pgdg | 13.5 KiB | [pg_permissions_17-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_permissions_17-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_17` | 1.4 | `el9.x86_64` | pgdg | 13.3 KiB | [pg_permissions_17-1.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_permissions_17-1.4-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_17` | 1.3 | `el9.x86_64` | pgdg | 13.2 KiB | [pg_permissions_17-1.3-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_permissions_17-1.3-2PGDG.rhel9.noarch.rpm) |
| `pg_permissions_17` | 1.4 | `el9.aarch64` | pgdg | 13.2 KiB | [pg_permissions_17-1.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_permissions_17-1.4-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_17` | 1.3 | `el9.aarch64` | pgdg | 13.2 KiB | [pg_permissions_17-1.3-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_permissions_17-1.3-2PGDG.rhel9.noarch.rpm) |
| `pg_permissions_17` | 1.4 | `el10.x86_64` | pgdg | 13.8 KiB | [pg_permissions_17-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_permissions_17-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_17` | 1.3 | `el10.x86_64` | pgdg | 13.7 KiB | [pg_permissions_17-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_permissions_17-1.3-2PGDG.rhel10.noarch.rpm) |
| `pg_permissions_17` | 1.4 | `el10.aarch64` | pgdg | 13.7 KiB | [pg_permissions_17-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_permissions_17-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_17` | 1.3 | `el10.aarch64` | pgdg | 13.7 KiB | [pg_permissions_17-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_permissions_17-1.3-2PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pg-permissions` | 1.4 | `d12.x86_64` | pgdg | 8.5 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-17-pg-permissions` | 1.3 | `d12.x86_64` | pigsty | 7.9 KiB | [postgresql-17-pg-permissions_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-permissions` | 1.4 | `d12.aarch64` | pgdg | 8.5 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-17-pg-permissions` | 1.3 | `d12.aarch64` | pigsty | 7.9 KiB | [postgresql-17-pg-permissions_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-permissions` | 1.4 | `d13.x86_64` | pgdg | 8.5 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-17-pg-permissions` | 1.4 | `d13.aarch64` | pgdg | 8.5 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-17-pg-permissions` | 1.4 | `u22.x86_64` | pgdg | 8.1 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-17-pg-permissions` | 1.3 | `u22.x86_64` | pigsty | 7.6 KiB | [postgresql-17-pg-permissions_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-permissions` | 1.4 | `u22.aarch64` | pgdg | 8.1 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-17-pg-permissions` | 1.3 | `u22.aarch64` | pigsty | 7.6 KiB | [postgresql-17-pg-permissions_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-permissions` | 1.4 | `u24.x86_64` | pgdg | 8.1 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-17-pg-permissions` | 1.3 | `u24.x86_64` | pigsty | 7.6 KiB | [postgresql-17-pg-permissions_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-permissions` | 1.4 | `u24.aarch64` | pgdg | 8.1 KiB | [postgresql-17-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-17-pg-permissions` | 1.3 | `u24.aarch64` | pigsty | 7.6 KiB | [postgresql-17-pg-permissions_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-permissions/postgresql-17-pg-permissions_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_permissions_16` | 1.4 | `el8.x86_64` | pgdg | 13.8 KiB | [pg_permissions_16-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_permissions_16-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_16` | 1.3 | `el8.x86_64` | pgdg | 13.1 KiB | [pg_permissions_16-1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_permissions_16-1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_16` | 1.3 | `el8.x86_64` | pgdg | 13.6 KiB | [pg_permissions_16-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_permissions_16-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_16` | 1.1 | `el8.x86_64` | pgdg | 12.4 KiB | [pg_permissions_16-1.1-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_permissions_16-1.1-3.rhel8.noarch.rpm) |
| `pg_permissions_16` | 1.4 | `el8.aarch64` | pgdg | 13.7 KiB | [pg_permissions_16-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_permissions_16-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_16` | 1.3 | `el8.aarch64` | pgdg | 13.0 KiB | [pg_permissions_16-1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_permissions_16-1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_16` | 1.3 | `el8.aarch64` | pgdg | 13.5 KiB | [pg_permissions_16-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_permissions_16-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_16` | 1.1 | `el8.aarch64` | pgdg | 12.3 KiB | [pg_permissions_16-1.1-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_permissions_16-1.1-3.rhel8.noarch.rpm) |
| `pg_permissions_16` | 1.4 | `el9.x86_64` | pgdg | 13.3 KiB | [pg_permissions_16-1.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_permissions_16-1.4-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_16` | 1.3 | `el9.x86_64` | pgdg | 12.7 KiB | [pg_permissions_16-1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_permissions_16-1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_16` | 1.3 | `el9.x86_64` | pgdg | 13.2 KiB | [pg_permissions_16-1.3-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_permissions_16-1.3-2PGDG.rhel9.noarch.rpm) |
| `pg_permissions_16` | 1.1 | `el9.x86_64` | pgdg | 12.2 KiB | [pg_permissions_16-1.1-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_permissions_16-1.1-3.rhel9.noarch.rpm) |
| `pg_permissions_16` | 1.4 | `el9.aarch64` | pgdg | 13.2 KiB | [pg_permissions_16-1.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_permissions_16-1.4-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_16` | 1.3 | `el9.aarch64` | pgdg | 12.7 KiB | [pg_permissions_16-1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_permissions_16-1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_16` | 1.3 | `el9.aarch64` | pgdg | 13.2 KiB | [pg_permissions_16-1.3-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_permissions_16-1.3-2PGDG.rhel9.noarch.rpm) |
| `pg_permissions_16` | 1.1 | `el9.aarch64` | pgdg | 12.1 KiB | [pg_permissions_16-1.1-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_permissions_16-1.1-3.rhel9.noarch.rpm) |
| `pg_permissions_16` | 1.4 | `el10.x86_64` | pgdg | 13.8 KiB | [pg_permissions_16-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_permissions_16-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_16` | 1.3 | `el10.x86_64` | pgdg | 13.7 KiB | [pg_permissions_16-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_permissions_16-1.3-2PGDG.rhel10.noarch.rpm) |
| `pg_permissions_16` | 1.4 | `el10.aarch64` | pgdg | 13.7 KiB | [pg_permissions_16-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_permissions_16-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_16` | 1.3 | `el10.aarch64` | pgdg | 13.7 KiB | [pg_permissions_16-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_permissions_16-1.3-2PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pg-permissions` | 1.4 | `d12.x86_64` | pgdg | 8.5 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-16-pg-permissions` | 1.3 | `d12.x86_64` | pigsty | 7.9 KiB | [postgresql-16-pg-permissions_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-permissions` | 1.4 | `d12.aarch64` | pgdg | 8.5 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-16-pg-permissions` | 1.3 | `d12.aarch64` | pigsty | 7.9 KiB | [postgresql-16-pg-permissions_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-permissions` | 1.4 | `d13.x86_64` | pgdg | 8.5 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-16-pg-permissions` | 1.4 | `d13.aarch64` | pgdg | 8.5 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-16-pg-permissions` | 1.4 | `u22.x86_64` | pgdg | 8.1 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-16-pg-permissions` | 1.3 | `u22.x86_64` | pigsty | 7.6 KiB | [postgresql-16-pg-permissions_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-permissions` | 1.4 | `u22.aarch64` | pgdg | 8.1 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-16-pg-permissions` | 1.3 | `u22.aarch64` | pigsty | 7.6 KiB | [postgresql-16-pg-permissions_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-permissions` | 1.4 | `u24.x86_64` | pgdg | 8.1 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-16-pg-permissions` | 1.3 | `u24.x86_64` | pigsty | 7.6 KiB | [postgresql-16-pg-permissions_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-permissions` | 1.4 | `u24.aarch64` | pgdg | 8.1 KiB | [postgresql-16-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-16-pg-permissions` | 1.3 | `u24.aarch64` | pigsty | 7.6 KiB | [postgresql-16-pg-permissions_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-permissions/postgresql-16-pg-permissions_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_permissions_15` | 1.4 | `el8.x86_64` | pgdg | 13.8 KiB | [pg_permissions_15-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_permissions_15-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_15` | 1.3 | `el8.x86_64` | pgdg | 13.6 KiB | [pg_permissions_15-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_permissions_15-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_15` | 1.3 | `el8.x86_64` | pgdg | 13.1 KiB | [pg_permissions_15-1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_permissions_15-1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_15` | 1.1 | `el8.x86_64` | pgdg | 12.3 KiB | [pg_permissions_15-1.1-2.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_permissions_15-1.1-2.rhel8.noarch.rpm) |
| `pg_permissions_15` | 1.4 | `el8.aarch64` | pgdg | 13.7 KiB | [pg_permissions_15-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_permissions_15-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_15` | 1.3 | `el8.aarch64` | pgdg | 13.5 KiB | [pg_permissions_15-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_permissions_15-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_15` | 1.3 | `el8.aarch64` | pgdg | 13.0 KiB | [pg_permissions_15-1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_permissions_15-1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_15` | 1.1 | `el8.aarch64` | pgdg | 12.2 KiB | [pg_permissions_15-1.1-2.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_permissions_15-1.1-2.rhel8.noarch.rpm) |
| `pg_permissions_15` | 1.4 | `el9.x86_64` | pgdg | 13.3 KiB | [pg_permissions_15-1.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_permissions_15-1.4-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_15` | 1.3 | `el9.x86_64` | pgdg | 12.7 KiB | [pg_permissions_15-1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_permissions_15-1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_15` | 1.3 | `el9.x86_64` | pgdg | 13.2 KiB | [pg_permissions_15-1.3-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_permissions_15-1.3-2PGDG.rhel9.noarch.rpm) |
| `pg_permissions_15` | 1.1 | `el9.x86_64` | pgdg | 12.2 KiB | [pg_permissions_15-1.1-2.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_permissions_15-1.1-2.rhel9.noarch.rpm) |
| `pg_permissions_15` | 1.4 | `el9.aarch64` | pgdg | 13.2 KiB | [pg_permissions_15-1.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_permissions_15-1.4-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_15` | 1.3 | `el9.aarch64` | pgdg | 13.2 KiB | [pg_permissions_15-1.3-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_permissions_15-1.3-2PGDG.rhel9.noarch.rpm) |
| `pg_permissions_15` | 1.3 | `el9.aarch64` | pgdg | 12.7 KiB | [pg_permissions_15-1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_permissions_15-1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_15` | 1.1 | `el9.aarch64` | pgdg | 12.0 KiB | [pg_permissions_15-1.1-2.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_permissions_15-1.1-2.rhel9.noarch.rpm) |
| `pg_permissions_15` | 1.4 | `el10.x86_64` | pgdg | 13.8 KiB | [pg_permissions_15-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_permissions_15-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_15` | 1.3 | `el10.x86_64` | pgdg | 13.7 KiB | [pg_permissions_15-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_permissions_15-1.3-2PGDG.rhel10.noarch.rpm) |
| `pg_permissions_15` | 1.4 | `el10.aarch64` | pgdg | 13.7 KiB | [pg_permissions_15-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_permissions_15-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_15` | 1.3 | `el10.aarch64` | pgdg | 13.7 KiB | [pg_permissions_15-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_permissions_15-1.3-2PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pg-permissions` | 1.4 | `d12.x86_64` | pgdg | 8.5 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-15-pg-permissions` | 1.3 | `d12.x86_64` | pigsty | 7.9 KiB | [postgresql-15-pg-permissions_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-permissions` | 1.4 | `d12.aarch64` | pgdg | 8.5 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-15-pg-permissions` | 1.3 | `d12.aarch64` | pigsty | 7.9 KiB | [postgresql-15-pg-permissions_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-permissions` | 1.4 | `d13.x86_64` | pgdg | 8.5 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-15-pg-permissions` | 1.4 | `d13.aarch64` | pgdg | 8.5 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-15-pg-permissions` | 1.4 | `u22.x86_64` | pgdg | 8.1 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-15-pg-permissions` | 1.3 | `u22.x86_64` | pigsty | 7.6 KiB | [postgresql-15-pg-permissions_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-permissions` | 1.4 | `u22.aarch64` | pgdg | 8.1 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-15-pg-permissions` | 1.3 | `u22.aarch64` | pigsty | 7.6 KiB | [postgresql-15-pg-permissions_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-permissions` | 1.4 | `u24.x86_64` | pgdg | 8.1 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-15-pg-permissions` | 1.3 | `u24.x86_64` | pigsty | 7.6 KiB | [postgresql-15-pg-permissions_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-permissions` | 1.4 | `u24.aarch64` | pgdg | 8.1 KiB | [postgresql-15-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-15-pg-permissions` | 1.3 | `u24.aarch64` | pigsty | 7.6 KiB | [postgresql-15-pg-permissions_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-permissions/postgresql-15-pg-permissions_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_permissions_14` | 1.4 | `el8.x86_64` | pgdg | 13.8 KiB | [pg_permissions_14-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_permissions_14-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_14` | 1.3 | `el8.x86_64` | pgdg | 13.1 KiB | [pg_permissions_14-1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_permissions_14-1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_14` | 1.3 | `el8.x86_64` | pgdg | 13.6 KiB | [pg_permissions_14-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_permissions_14-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_14` | 1.1 | `el8.x86_64` | pgdg | 12.3 KiB | [pg_permissions_14-1.1-2.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_permissions_14-1.1-2.rhel8.noarch.rpm) |
| `pg_permissions_14` | 1.4 | `el8.aarch64` | pgdg | 13.7 KiB | [pg_permissions_14-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_permissions_14-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_14` | 1.3 | `el8.aarch64` | pgdg | 13.5 KiB | [pg_permissions_14-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_permissions_14-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_14` | 1.3 | `el8.aarch64` | pgdg | 13.0 KiB | [pg_permissions_14-1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_permissions_14-1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_14` | 1.1 | `el8.aarch64` | pgdg | 12.2 KiB | [pg_permissions_14-1.1-2.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_permissions_14-1.1-2.rhel8.noarch.rpm) |
| `pg_permissions_14` | 1.4 | `el9.x86_64` | pgdg | 13.3 KiB | [pg_permissions_14-1.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_permissions_14-1.4-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_14` | 1.3 | `el9.x86_64` | pgdg | 12.7 KiB | [pg_permissions_14-1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_permissions_14-1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_14` | 1.3 | `el9.x86_64` | pgdg | 13.2 KiB | [pg_permissions_14-1.3-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_permissions_14-1.3-2PGDG.rhel9.noarch.rpm) |
| `pg_permissions_14` | 1.1 | `el9.x86_64` | pgdg | 12.2 KiB | [pg_permissions_14-1.1-2.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_permissions_14-1.1-2.rhel9.noarch.rpm) |
| `pg_permissions_14` | 1.4 | `el9.aarch64` | pgdg | 13.2 KiB | [pg_permissions_14-1.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_permissions_14-1.4-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_14` | 1.3 | `el9.aarch64` | pgdg | 13.2 KiB | [pg_permissions_14-1.3-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_permissions_14-1.3-2PGDG.rhel9.noarch.rpm) |
| `pg_permissions_14` | 1.3 | `el9.aarch64` | pgdg | 12.7 KiB | [pg_permissions_14-1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_permissions_14-1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_14` | 1.1 | `el9.aarch64` | pgdg | 12.0 KiB | [pg_permissions_14-1.1-2.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_permissions_14-1.1-2.rhel9.noarch.rpm) |
| `pg_permissions_14` | 1.4 | `el10.x86_64` | pgdg | 13.8 KiB | [pg_permissions_14-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_permissions_14-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_14` | 1.3 | `el10.x86_64` | pgdg | 13.7 KiB | [pg_permissions_14-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_permissions_14-1.3-2PGDG.rhel10.noarch.rpm) |
| `pg_permissions_14` | 1.4 | `el10.aarch64` | pgdg | 13.7 KiB | [pg_permissions_14-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_permissions_14-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_14` | 1.3 | `el10.aarch64` | pgdg | 13.7 KiB | [pg_permissions_14-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_permissions_14-1.3-2PGDG.rhel10.noarch.rpm) |
| `postgresql-14-pg-permissions` | 1.4 | `d12.x86_64` | pgdg | 8.5 KiB | [postgresql-14-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-14-pg-permissions` | 1.3 | `d12.x86_64` | pigsty | 7.9 KiB | [postgresql-14-pg-permissions_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-permissions` | 1.4 | `d12.aarch64` | pgdg | 8.5 KiB | [postgresql-14-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-14-pg-permissions` | 1.3 | `d12.aarch64` | pigsty | 7.9 KiB | [postgresql-14-pg-permissions_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-permissions` | 1.4 | `d13.x86_64` | pgdg | 8.5 KiB | [postgresql-14-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-14-pg-permissions` | 1.4 | `d13.aarch64` | pgdg | 8.5 KiB | [postgresql-14-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-14-pg-permissions` | 1.4 | `u22.x86_64` | pgdg | 8.1 KiB | [postgresql-14-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-14-pg-permissions` | 1.3 | `u22.x86_64` | pigsty | 7.6 KiB | [postgresql-14-pg-permissions_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-permissions` | 1.4 | `u22.aarch64` | pgdg | 8.1 KiB | [postgresql-14-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-14-pg-permissions` | 1.3 | `u22.aarch64` | pigsty | 7.6 KiB | [postgresql-14-pg-permissions_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-permissions` | 1.4 | `u24.x86_64` | pgdg | 8.1 KiB | [postgresql-14-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-14-pg-permissions` | 1.3 | `u24.x86_64` | pigsty | 7.6 KiB | [postgresql-14-pg-permissions_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-permissions` | 1.4 | `u24.aarch64` | pgdg | 8.1 KiB | [postgresql-14-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-14-pg-permissions` | 1.3 | `u24.aarch64` | pigsty | 7.6 KiB | [postgresql-14-pg-permissions_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-permissions/postgresql-14-pg-permissions_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_permissions_13` | 1.4 | `el8.x86_64` | pgdg | 13.8 KiB | [pg_permissions_13-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_permissions_13-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_13` | 1.3 | `el8.x86_64` | pgdg | 13.6 KiB | [pg_permissions_13-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_permissions_13-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_13` | 1.3 | `el8.x86_64` | pgdg | 13.1 KiB | [pg_permissions_13-1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_permissions_13-1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_13` | 1.1 | `el8.x86_64` | pgdg | 12.3 KiB | [pg_permissions_13-1.1-2.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_permissions_13-1.1-2.rhel8.noarch.rpm) |
| `pg_permissions_13` | 1.4 | `el8.aarch64` | pgdg | 13.7 KiB | [pg_permissions_13-1.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_permissions_13-1.4-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_13` | 1.3 | `el8.aarch64` | pgdg | 13.5 KiB | [pg_permissions_13-1.3-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_permissions_13-1.3-2PGDG.rhel8.noarch.rpm) |
| `pg_permissions_13` | 1.3 | `el8.aarch64` | pgdg | 13.0 KiB | [pg_permissions_13-1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_permissions_13-1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_permissions_13` | 1.1 | `el8.aarch64` | pgdg | 12.2 KiB | [pg_permissions_13-1.1-2.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_permissions_13-1.1-2.rhel8.noarch.rpm) |
| `pg_permissions_13` | 1.4 | `el9.x86_64` | pgdg | 13.3 KiB | [pg_permissions_13-1.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_permissions_13-1.4-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_13` | 1.3 | `el9.x86_64` | pgdg | 13.2 KiB | [pg_permissions_13-1.3-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_permissions_13-1.3-2PGDG.rhel9.noarch.rpm) |
| `pg_permissions_13` | 1.3 | `el9.x86_64` | pgdg | 12.7 KiB | [pg_permissions_13-1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_permissions_13-1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_13` | 1.1 | `el9.x86_64` | pgdg | 12.2 KiB | [pg_permissions_13-1.1-2.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_permissions_13-1.1-2.rhel9.noarch.rpm) |
| `pg_permissions_13` | 1.4 | `el9.aarch64` | pgdg | 13.2 KiB | [pg_permissions_13-1.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_permissions_13-1.4-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_13` | 1.3 | `el9.aarch64` | pgdg | 12.7 KiB | [pg_permissions_13-1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_permissions_13-1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_permissions_13` | 1.3 | `el9.aarch64` | pgdg | 13.2 KiB | [pg_permissions_13-1.3-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_permissions_13-1.3-2PGDG.rhel9.noarch.rpm) |
| `pg_permissions_13` | 1.1 | `el9.aarch64` | pgdg | 12.0 KiB | [pg_permissions_13-1.1-2.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_permissions_13-1.1-2.rhel9.noarch.rpm) |
| `pg_permissions_13` | 1.4 | `el10.x86_64` | pgdg | 13.8 KiB | [pg_permissions_13-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_permissions_13-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_13` | 1.3 | `el10.x86_64` | pgdg | 13.7 KiB | [pg_permissions_13-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_permissions_13-1.3-2PGDG.rhel10.noarch.rpm) |
| `pg_permissions_13` | 1.4 | `el10.aarch64` | pgdg | 13.7 KiB | [pg_permissions_13-1.4-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_permissions_13-1.4-1PGDG.rhel10.noarch.rpm) |
| `pg_permissions_13` | 1.3 | `el10.aarch64` | pgdg | 13.7 KiB | [pg_permissions_13-1.3-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_permissions_13-1.3-2PGDG.rhel10.noarch.rpm) |
| `postgresql-13-pg-permissions` | 1.4 | `d12.x86_64` | pgdg | 8.5 KiB | [postgresql-13-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-13-pg-permissions` | 1.3 | `d12.x86_64` | pigsty | 7.9 KiB | [postgresql-13-pg-permissions_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-permissions` | 1.4 | `d12.aarch64` | pgdg | 8.5 KiB | [postgresql-13-pg-permissions_1.4-2.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.4-2.pgdg12+1_all.deb) |
| `postgresql-13-pg-permissions` | 1.3 | `d12.aarch64` | pigsty | 7.9 KiB | [postgresql-13-pg-permissions_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-permissions` | 1.4 | `d13.x86_64` | pgdg | 8.4 KiB | [postgresql-13-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-13-pg-permissions` | 1.4 | `d13.aarch64` | pgdg | 8.4 KiB | [postgresql-13-pg-permissions_1.4-2.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.4-2.pgdg13+1_all.deb) |
| `postgresql-13-pg-permissions` | 1.4 | `u22.x86_64` | pgdg | 8.1 KiB | [postgresql-13-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-13-pg-permissions` | 1.3 | `u22.x86_64` | pigsty | 7.6 KiB | [postgresql-13-pg-permissions_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-permissions` | 1.4 | `u22.aarch64` | pgdg | 8.1 KiB | [postgresql-13-pg-permissions_1.4-2.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.4-2.pgdg22.04+1_all.deb) |
| `postgresql-13-pg-permissions` | 1.3 | `u22.aarch64` | pigsty | 7.6 KiB | [postgresql-13-pg-permissions_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-permissions` | 1.4 | `u24.x86_64` | pgdg | 8.1 KiB | [postgresql-13-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-13-pg-permissions` | 1.3 | `u24.x86_64` | pigsty | 7.6 KiB | [postgresql-13-pg-permissions_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-permissions` | 1.4 | `u24.aarch64` | pgdg | 8.1 KiB | [postgresql-13-pg-permissions_1.4-2.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.4-2.pgdg24.04+1_all.deb) |
| `postgresql-13-pg-permissions` | 1.3 | `u24.aarch64` | pigsty | 7.6 KiB | [postgresql-13-pg-permissions_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-permissions/postgresql-13-pg-permissions_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/pg_permissions" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/pg_permissions" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_permissions-REL_1_3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_permissions; # get pg_permissions source code
pig build dep pg_permissions; # install build dependencies
pig build pkg pg_permissions; # build extension rpm or deb
pig build ext pg_permissions; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_permissions; # install by extension name, for the current active PG version
pig ext install pg_permissions; # install via package alias, for the active PG version
pig ext install pg_permissions -v 18;   # install for PG 18
pig ext install pg_permissions -v 17;   # install for PG 17
pig ext install pg_permissions -v 16;   # install for PG 16
pig ext install pg_permissions -v 15;   # install for PG 15
pig ext install pg_permissions -v 14;   # install for PG 14
pig ext install pg_permissions -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_permissions;
```

