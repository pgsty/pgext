---
title: "pg_show_plans"
linkTitle: "pg_show_plans"
description: "show query plans of all currently running SQL statements"
weight: 6210
categories: ["STAT"]
width: full
---

show query plans of all currently running SQL statements


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6210** | {{< badge content="pg_show_plans" link="https://github.com/cybertec-postgresql/pg_show_plans" >}} | {{< ext "pg_show_plans" >}} | `2.1.6` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_store_plans" >}} {{< ext "explain_ui" >}} {{< ext "auto_explain" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_hint_plan" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pre_prepare" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_show_plans" >}} | `2.1.6` | {{< bg "18" "pg_show_plans_18*" "green" >}} {{< bg "17" "pg_show_plans_17*" "green" >}} {{< bg "16" "pg_show_plans_16*" "green" >}} {{< bg "15" "pg_show_plans_15*" "green" >}} {{< bg "14" "pg_show_plans_14*" "green" >}} | `pg_show_plans_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_show_plans" >}} | `2.1.6` | {{< bg "18" "postgresql-18-show-plans" "green" >}} {{< bg "17" "postgresql-17-show-plans" "green" >}} {{< bg "16" "postgresql-16-show-plans" "green" >}} {{< bg "15" "postgresql-15-show-plans" "green" >}} {{< bg "14" "postgresql-14-show-plans" "green" >}} | `postgresql-$v-show-plans` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 2.1.6" "pg_show_plans_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_14 : AVAIL 4" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 2.1.6" "pg_show_plans_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_14 : AVAIL 4" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 2.1.6" "pg_show_plans_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_14 : AVAIL 4" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 2.1.6" "pg_show_plans_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.1.6" "pg_show_plans_14 : AVAIL 4" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 2.1.7" "postgresql-18-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-17-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-16-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-15-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-14-show-plans : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 2.1.7" "postgresql-18-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-17-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-16-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-15-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-14-show-plans : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 2.1.7" "postgresql-18-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-17-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-16-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-15-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-14-show-plans : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 2.1.7" "postgresql-18-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-17-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-16-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-15-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-14-show-plans : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 2.1.7" "postgresql-18-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-17-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-16-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-15-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-14-show-plans : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 2.1.7" "postgresql-18-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-17-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-16-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-15-show-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.7" "postgresql-14-show-plans : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_show_plans_18` | 2.1.6 | `el8.x86_64` | pgdg | 19.2 KiB | [pg_show_plans_18-2.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_show_plans_18-2.1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_18` | 2.1.6 | `el8.aarch64` | pgdg | 19.1 KiB | [pg_show_plans_18-2.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_show_plans_18-2.1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_18` | 2.1.6 | `el9.x86_64` | pgdg | 19.7 KiB | [pg_show_plans_18-2.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_show_plans_18-2.1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_18` | 2.1.6 | `el9.aarch64` | pgdg | 19.4 KiB | [pg_show_plans_18-2.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_show_plans_18-2.1.6-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-show-plans` | 2.1.7 | `d12.x86_64` | pgdg | 23.2 KiB | [postgresql-18-show-plans_2.1.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-18-show-plans_2.1.7-1.pgdg12+1_amd64.deb) |
| `postgresql-18-show-plans` | 2.1.7 | `d12.aarch64` | pgdg | 23.0 KiB | [postgresql-18-show-plans_2.1.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-18-show-plans_2.1.7-1.pgdg12+1_arm64.deb) |
| `postgresql-18-show-plans` | 2.1.7 | `u22.x86_64` | pgdg | 23.6 KiB | [postgresql-18-show-plans_2.1.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-18-show-plans_2.1.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-show-plans` | 2.1.7 | `u22.aarch64` | pgdg | 23.0 KiB | [postgresql-18-show-plans_2.1.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-18-show-plans_2.1.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-show-plans` | 2.1.7 | `u24.x86_64` | pgdg | 23.3 KiB | [postgresql-18-show-plans_2.1.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-18-show-plans_2.1.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-show-plans` | 2.1.7 | `u24.aarch64` | pgdg | 23.1 KiB | [postgresql-18-show-plans_2.1.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-18-show-plans_2.1.7-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_show_plans_17` | 2.1.6 | `el8.x86_64` | pgdg | 19.2 KiB | [pg_show_plans_17-2.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_show_plans_17-2.1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_17` | 2.1.3 | `el8.x86_64` | pgdg | 19.0 KiB | [pg_show_plans_17-2.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_show_plans_17-2.1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_17` | 2.1.2 | `el8.x86_64` | pgdg | 18.8 KiB | [pg_show_plans_17-2.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_show_plans_17-2.1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_17` | 2.1.6 | `el8.aarch64` | pgdg | 19.1 KiB | [pg_show_plans_17-2.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_show_plans_17-2.1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_17` | 2.1.3 | `el8.aarch64` | pgdg | 18.9 KiB | [pg_show_plans_17-2.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_show_plans_17-2.1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_17` | 2.1.2 | `el8.aarch64` | pgdg | 18.7 KiB | [pg_show_plans_17-2.1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_show_plans_17-2.1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_17` | 2.1.6 | `el9.x86_64` | pgdg | 19.7 KiB | [pg_show_plans_17-2.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_show_plans_17-2.1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_17` | 2.1.3 | `el9.x86_64` | pgdg | 19.5 KiB | [pg_show_plans_17-2.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_show_plans_17-2.1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_17` | 2.1.2 | `el9.x86_64` | pgdg | 19.3 KiB | [pg_show_plans_17-2.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_show_plans_17-2.1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_17` | 2.1.6 | `el9.aarch64` | pgdg | 19.4 KiB | [pg_show_plans_17-2.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_show_plans_17-2.1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_show_plans_17` | 2.1.3 | `el9.aarch64` | pgdg | 19.4 KiB | [pg_show_plans_17-2.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_show_plans_17-2.1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_show_plans_17` | 2.1.2 | `el9.aarch64` | pgdg | 19.2 KiB | [pg_show_plans_17-2.1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_show_plans_17-2.1.2-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-show-plans` | 2.1.7 | `d12.x86_64` | pgdg | 23.3 KiB | [postgresql-17-show-plans_2.1.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-17-show-plans_2.1.7-1.pgdg12+1_amd64.deb) |
| `postgresql-17-show-plans` | 2.1.7 | `d12.aarch64` | pgdg | 23.0 KiB | [postgresql-17-show-plans_2.1.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-17-show-plans_2.1.7-1.pgdg12+1_arm64.deb) |
| `postgresql-17-show-plans` | 2.1.7 | `u22.x86_64` | pgdg | 27.1 KiB | [postgresql-17-show-plans_2.1.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-17-show-plans_2.1.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-show-plans` | 2.1.7 | `u22.aarch64` | pgdg | 26.3 KiB | [postgresql-17-show-plans_2.1.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-17-show-plans_2.1.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-show-plans` | 2.1.7 | `u24.x86_64` | pgdg | 23.4 KiB | [postgresql-17-show-plans_2.1.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-17-show-plans_2.1.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-show-plans` | 2.1.7 | `u24.aarch64` | pgdg | 23.0 KiB | [postgresql-17-show-plans_2.1.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-17-show-plans_2.1.7-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_show_plans_16` | 2.1.6 | `el8.x86_64` | pgdg | 19.2 KiB | [pg_show_plans_16-2.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_show_plans_16-2.1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_16` | 2.1.3 | `el8.x86_64` | pgdg | 19.0 KiB | [pg_show_plans_16-2.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_show_plans_16-2.1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_16` | 2.1.2 | `el8.x86_64` | pgdg | 18.8 KiB | [pg_show_plans_16-2.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_show_plans_16-2.1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_16` | 2.1.0 | `el8.x86_64` | pgdg | 18.6 KiB | [pg_show_plans_16-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_show_plans_16-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_16` | 2.1.6 | `el8.aarch64` | pgdg | 19.1 KiB | [pg_show_plans_16-2.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_show_plans_16-2.1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_16` | 2.1.3 | `el8.aarch64` | pgdg | 18.9 KiB | [pg_show_plans_16-2.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_show_plans_16-2.1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_16` | 2.1.2 | `el8.aarch64` | pgdg | 18.7 KiB | [pg_show_plans_16-2.1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_show_plans_16-2.1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_16` | 2.1.0 | `el8.aarch64` | pgdg | 18.5 KiB | [pg_show_plans_16-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_show_plans_16-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_16` | 2.1.6 | `el9.x86_64` | pgdg | 19.7 KiB | [pg_show_plans_16-2.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_show_plans_16-2.1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_16` | 2.1.3 | `el9.x86_64` | pgdg | 19.5 KiB | [pg_show_plans_16-2.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_show_plans_16-2.1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_16` | 2.1.2 | `el9.x86_64` | pgdg | 19.3 KiB | [pg_show_plans_16-2.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_show_plans_16-2.1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_16` | 2.1.0 | `el9.x86_64` | pgdg | 19.0 KiB | [pg_show_plans_16-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_show_plans_16-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_16` | 2.1.6 | `el9.aarch64` | pgdg | 19.4 KiB | [pg_show_plans_16-2.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_show_plans_16-2.1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_show_plans_16` | 2.1.3 | `el9.aarch64` | pgdg | 19.4 KiB | [pg_show_plans_16-2.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_show_plans_16-2.1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_show_plans_16` | 2.1.2 | `el9.aarch64` | pgdg | 19.2 KiB | [pg_show_plans_16-2.1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_show_plans_16-2.1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_show_plans_16` | 2.1.0 | `el9.aarch64` | pgdg | 19.0 KiB | [pg_show_plans_16-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_show_plans_16-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-show-plans` | 2.1.7 | `d12.x86_64` | pgdg | 23.3 KiB | [postgresql-16-show-plans_2.1.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-16-show-plans_2.1.7-1.pgdg12+1_amd64.deb) |
| `postgresql-16-show-plans` | 2.1.7 | `d12.aarch64` | pgdg | 23.0 KiB | [postgresql-16-show-plans_2.1.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-16-show-plans_2.1.7-1.pgdg12+1_arm64.deb) |
| `postgresql-16-show-plans` | 2.1.7 | `u22.x86_64` | pgdg | 27.1 KiB | [postgresql-16-show-plans_2.1.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-16-show-plans_2.1.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-show-plans` | 2.1.7 | `u22.aarch64` | pgdg | 26.2 KiB | [postgresql-16-show-plans_2.1.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-16-show-plans_2.1.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-show-plans` | 2.1.7 | `u24.x86_64` | pgdg | 23.4 KiB | [postgresql-16-show-plans_2.1.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-16-show-plans_2.1.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-show-plans` | 2.1.7 | `u24.aarch64` | pgdg | 23.0 KiB | [postgresql-16-show-plans_2.1.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-16-show-plans_2.1.7-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_show_plans_15` | 2.1.6 | `el8.x86_64` | pgdg | 19.2 KiB | [pg_show_plans_15-2.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_show_plans_15-2.1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_15` | 2.1.3 | `el8.x86_64` | pgdg | 19.0 KiB | [pg_show_plans_15-2.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_show_plans_15-2.1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_15` | 2.1.2 | `el8.x86_64` | pgdg | 18.8 KiB | [pg_show_plans_15-2.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_show_plans_15-2.1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_15` | 2.1.0 | `el8.x86_64` | pgdg | 18.6 KiB | [pg_show_plans_15-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_show_plans_15-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_15` | 2.1.6 | `el8.aarch64` | pgdg | 19.1 KiB | [pg_show_plans_15-2.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_show_plans_15-2.1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_15` | 2.1.3 | `el8.aarch64` | pgdg | 18.9 KiB | [pg_show_plans_15-2.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_show_plans_15-2.1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_15` | 2.1.2 | `el8.aarch64` | pgdg | 18.7 KiB | [pg_show_plans_15-2.1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_show_plans_15-2.1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_15` | 2.1.0 | `el8.aarch64` | pgdg | 18.5 KiB | [pg_show_plans_15-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_show_plans_15-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_15` | 2.1.6 | `el9.x86_64` | pgdg | 19.7 KiB | [pg_show_plans_15-2.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_show_plans_15-2.1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_15` | 2.1.3 | `el9.x86_64` | pgdg | 19.5 KiB | [pg_show_plans_15-2.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_show_plans_15-2.1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_15` | 2.1.2 | `el9.x86_64` | pgdg | 19.3 KiB | [pg_show_plans_15-2.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_show_plans_15-2.1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_15` | 2.1.0 | `el9.x86_64` | pgdg | 19.0 KiB | [pg_show_plans_15-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_show_plans_15-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_15` | 2.1.6 | `el9.aarch64` | pgdg | 19.4 KiB | [pg_show_plans_15-2.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_show_plans_15-2.1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_show_plans_15` | 2.1.3 | `el9.aarch64` | pgdg | 19.4 KiB | [pg_show_plans_15-2.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_show_plans_15-2.1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_show_plans_15` | 2.1.2 | `el9.aarch64` | pgdg | 19.2 KiB | [pg_show_plans_15-2.1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_show_plans_15-2.1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_show_plans_15` | 2.1.0 | `el9.aarch64` | pgdg | 19.0 KiB | [pg_show_plans_15-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_show_plans_15-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-show-plans` | 2.1.7 | `d12.x86_64` | pgdg | 23.3 KiB | [postgresql-15-show-plans_2.1.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-15-show-plans_2.1.7-1.pgdg12+1_amd64.deb) |
| `postgresql-15-show-plans` | 2.1.7 | `d12.aarch64` | pgdg | 23.0 KiB | [postgresql-15-show-plans_2.1.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-15-show-plans_2.1.7-1.pgdg12+1_arm64.deb) |
| `postgresql-15-show-plans` | 2.1.7 | `u22.x86_64` | pgdg | 27.0 KiB | [postgresql-15-show-plans_2.1.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-15-show-plans_2.1.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-show-plans` | 2.1.7 | `u22.aarch64` | pgdg | 26.3 KiB | [postgresql-15-show-plans_2.1.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-15-show-plans_2.1.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-show-plans` | 2.1.7 | `u24.x86_64` | pgdg | 23.4 KiB | [postgresql-15-show-plans_2.1.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-15-show-plans_2.1.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-show-plans` | 2.1.7 | `u24.aarch64` | pgdg | 23.0 KiB | [postgresql-15-show-plans_2.1.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-15-show-plans_2.1.7-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_show_plans_14` | 2.1.6 | `el8.x86_64` | pgdg | 19.1 KiB | [pg_show_plans_14-2.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_show_plans_14-2.1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_14` | 2.1.3 | `el8.x86_64` | pgdg | 19.0 KiB | [pg_show_plans_14-2.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_show_plans_14-2.1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_14` | 2.1.2 | `el8.x86_64` | pgdg | 18.7 KiB | [pg_show_plans_14-2.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_show_plans_14-2.1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_14` | 2.1.0 | `el8.x86_64` | pgdg | 18.5 KiB | [pg_show_plans_14-2.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_show_plans_14-2.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_show_plans_14` | 2.1.6 | `el8.aarch64` | pgdg | 19.0 KiB | [pg_show_plans_14-2.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_show_plans_14-2.1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_14` | 2.1.3 | `el8.aarch64` | pgdg | 18.8 KiB | [pg_show_plans_14-2.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_show_plans_14-2.1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_14` | 2.1.2 | `el8.aarch64` | pgdg | 18.6 KiB | [pg_show_plans_14-2.1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_show_plans_14-2.1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_14` | 2.1.0 | `el8.aarch64` | pgdg | 18.4 KiB | [pg_show_plans_14-2.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_show_plans_14-2.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_show_plans_14` | 2.1.6 | `el9.x86_64` | pgdg | 19.6 KiB | [pg_show_plans_14-2.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_show_plans_14-2.1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_14` | 2.1.3 | `el9.x86_64` | pgdg | 19.4 KiB | [pg_show_plans_14-2.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_show_plans_14-2.1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_14` | 2.1.2 | `el9.x86_64` | pgdg | 19.2 KiB | [pg_show_plans_14-2.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_show_plans_14-2.1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_14` | 2.1.0 | `el9.x86_64` | pgdg | 19.0 KiB | [pg_show_plans_14-2.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_show_plans_14-2.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_show_plans_14` | 2.1.6 | `el9.aarch64` | pgdg | 19.4 KiB | [pg_show_plans_14-2.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_show_plans_14-2.1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_show_plans_14` | 2.1.3 | `el9.aarch64` | pgdg | 19.3 KiB | [pg_show_plans_14-2.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_show_plans_14-2.1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_show_plans_14` | 2.1.2 | `el9.aarch64` | pgdg | 19.1 KiB | [pg_show_plans_14-2.1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_show_plans_14-2.1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_show_plans_14` | 2.1.0 | `el9.aarch64` | pgdg | 18.9 KiB | [pg_show_plans_14-2.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_show_plans_14-2.1.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-14-show-plans` | 2.1.7 | `d12.x86_64` | pgdg | 23.1 KiB | [postgresql-14-show-plans_2.1.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-14-show-plans_2.1.7-1.pgdg12+1_amd64.deb) |
| `postgresql-14-show-plans` | 2.1.7 | `d12.aarch64` | pgdg | 22.8 KiB | [postgresql-14-show-plans_2.1.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-14-show-plans_2.1.7-1.pgdg12+1_arm64.deb) |
| `postgresql-14-show-plans` | 2.1.7 | `u22.x86_64` | pgdg | 26.8 KiB | [postgresql-14-show-plans_2.1.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-14-show-plans_2.1.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-show-plans` | 2.1.7 | `u22.aarch64` | pgdg | 26.0 KiB | [postgresql-14-show-plans_2.1.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-14-show-plans_2.1.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-show-plans` | 2.1.7 | `u24.x86_64` | pgdg | 23.2 KiB | [postgresql-14-show-plans_2.1.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-14-show-plans_2.1.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-show-plans` | 2.1.7 | `u24.aarch64` | pgdg | 22.9 KiB | [postgresql-14-show-plans_2.1.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-show-plans/postgresql-14-show-plans_2.1.7-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/pg_show_plans" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/pg_show_plans" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_show_plans-2.1.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_show_plans; # get pg_show_plans source code
pig build dep pg_show_plans; # install build dependencies
pig build pkg pg_show_plans; # build extension rpm or deb
pig build ext pg_show_plans; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_show_plans; # install by extension name, for the current active PG version
pig ext install pg_show_plans; # install via package alias, for the active PG version
pig ext install pg_show_plans -v 18;   # install for PG 18
pig ext install pg_show_plans -v 17;   # install for PG 17
pig ext install pg_show_plans -v 16;   # install for PG 16
pig ext install pg_show_plans -v 15;   # install for PG 15
pig ext install pg_show_plans -v 14;   # install for PG 14
pig ext install pg_show_plans -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_show_plans;
```

