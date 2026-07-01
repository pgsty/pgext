---
title: "pg_stat_plans"
linkTitle: "pg_stat_plans"
description: "Track per-plan call counts, execution times, and example EXPLAIN texts."
weight: 6050
categories: ["STAT"]
width: full
---

[**pg_stat_plans**](https://github.com/pganalyze/pg_stat_plans) : Track per-plan call counts, execution times, and example EXPLAIN texts.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6050** | {{< badge content="pg_stat_plans" link="https://github.com/pganalyze/pg_stat_plans" >}} | {{< ext "pg_stat_plans" >}} | `2.1.0` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_stat_statements" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_stat_plans` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.1.0` | {{< bg "18" "pg_stat_plans_18" "green" >}} {{< bg "17" "pg_stat_plans_17" "green" >}} {{< bg "16" "pg_stat_plans_16" "green" >}} {{< bg "15" "pg_stat_plans_15" "red" >}} {{< bg "14" "pg_stat_plans_14" "red" >}} | `pg_stat_plans_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.1.0` | {{< bg "18" "postgresql-18-pg-stat-plans" "green" >}} {{< bg "17" "postgresql-17-pg-stat-plans" "green" >}} {{< bg "16" "postgresql-16-pg-stat-plans" "green" >}} {{< bg "15" "postgresql-15-pg-stat-plans" "red" >}} {{< bg "14" "postgresql-14-pg-stat-plans" "red" >}} | `postgresql-$v-pg-stat-plans` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_16 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_stat_plans_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_plans_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_16 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_stat_plans_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_plans_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_16 : AVAIL 3" "blue" >}} |      {{< bg "MISS" "pg_stat_plans_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_plans_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_16 : AVAIL 3" "blue" >}} |      {{< bg "MISS" "pg_stat_plans_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_plans_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_16 : AVAIL 3" "blue" >}} |      {{< bg "MISS" "pg_stat_plans_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_plans_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.0" "pg_stat_plans_16 : AVAIL 3" "blue" >}} |      {{< bg "MISS" "pg_stat_plans_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stat_plans_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-stat-plans : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-stat-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-stat-plans : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-plans : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-plans : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-stat-plans : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-stat-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-stat-plans : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-plans : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-plans : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-stat-plans : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-stat-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-stat-plans : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-plans : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-plans : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-stat-plans : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-stat-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-stat-plans : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-plans : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-plans : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-stat-plans : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-stat-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-stat-plans : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-plans : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-plans : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-stat-plans : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-stat-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-stat-plans : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-plans : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-plans : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-stat-plans : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-stat-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-stat-plans : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-plans : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-plans : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-stat-plans : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-stat-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-stat-plans : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-plans : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-plans : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-stat-plans : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-stat-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-stat-plans : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-plans : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-plans : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.1.0" "postgresql-18-pg-stat-plans : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-17-pg-stat-plans : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.0" "postgresql-16-pg-stat-plans : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-15-pg-stat-plans : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stat-plans : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_plans_18` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_stat_plans_18-2.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_stat_plans_18-2.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.3 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_plans_18-2.1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.3 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_plans_18-2.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.4 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_stat_plans_18-2.1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.5 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_plans_18-2.1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.5 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_plans_18-2.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_stat_plans_18-2.1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.8 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_plans_18-2.1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.9 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_plans_18-2.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.2 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_stat_plans_18-2.1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.0 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_plans_18-2.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.0 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_plans_18-2.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_stat_plans_18` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.0 KiB | [pg_stat_plans_18-2.1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_stat_plans_18-2.1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-pg-stat-plans` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 85.8 KiB | [postgresql-18-pg-stat-plans_2.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-stat-plans` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 81.6 KiB | [postgresql-18-pg-stat-plans_2.0.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.0.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-stat-plans` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 83.4 KiB | [postgresql-18-pg-stat-plans_2.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-stat-plans` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.2 KiB | [postgresql-18-pg-stat-plans_2.0.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.0.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-stat-plans` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 85.8 KiB | [postgresql-18-pg-stat-plans_2.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-stat-plans` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 81.5 KiB | [postgresql-18-pg-stat-plans_2.0.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.0.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-stat-plans` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 83.3 KiB | [postgresql-18-pg-stat-plans_2.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-stat-plans` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.2 KiB | [postgresql-18-pg-stat-plans_2.0.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.0.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-stat-plans` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 89.2 KiB | [postgresql-18-pg-stat-plans_2.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-stat-plans` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 84.8 KiB | [postgresql-18-pg-stat-plans_2.0.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.0.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-stat-plans` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 86.7 KiB | [postgresql-18-pg-stat-plans_2.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-stat-plans` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 82.9 KiB | [postgresql-18-pg-stat-plans_2.0.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.0.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-stat-plans` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 83.5 KiB | [postgresql-18-pg-stat-plans_2.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-stat-plans` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 80.3 KiB | [postgresql-18-pg-stat-plans_2.0.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.0.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-stat-plans` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 81.2 KiB | [postgresql-18-pg-stat-plans_2.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.1.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-stat-plans` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 78.0 KiB | [postgresql-18-pg-stat-plans_2.0.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.0.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-stat-plans` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 82.8 KiB | [postgresql-18-pg-stat-plans_2.1.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.1.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-stat-plans` | `2.0.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 80.3 KiB | [postgresql-18-pg-stat-plans_2.0.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.0.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-stat-plans` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 80.6 KiB | [postgresql-18-pg-stat-plans_2.1.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.1.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pg-stat-plans` | `2.0.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 77.3 KiB | [postgresql-18-pg-stat-plans_2.0.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-18-pg-stat-plans_2.0.0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_plans_17` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 47.2 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_stat_plans_17-2.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.0 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_stat_plans_17-2.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.3 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_plans_17-2.1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.3 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_plans_17-2.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.4 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_stat_plans_17-2.1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.7 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_plans_17-2.1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.7 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_plans_17-2.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.8 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_stat_plans_17-2.1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.5 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_stat_plans_17-2.1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.5 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_stat_plans_17-2.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.9 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_stat_plans_17-2.1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.4 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_plans_17-2.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.4 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_plans_17-2.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_stat_plans_17` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.4 KiB | [pg_stat_plans_17-2.1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_stat_plans_17-2.1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-17-pg-stat-plans` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 101.2 KiB | [postgresql-17-pg-stat-plans_2.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-17-pg-stat-plans_2.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-stat-plans` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 97.9 KiB | [postgresql-17-pg-stat-plans_2.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-17-pg-stat-plans_2.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-stat-plans` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 101.5 KiB | [postgresql-17-pg-stat-plans_2.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-17-pg-stat-plans_2.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-stat-plans` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 98.3 KiB | [postgresql-17-pg-stat-plans_2.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-17-pg-stat-plans_2.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-stat-plans` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 109.6 KiB | [postgresql-17-pg-stat-plans_2.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-17-pg-stat-plans_2.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-stat-plans` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 106.4 KiB | [postgresql-17-pg-stat-plans_2.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-17-pg-stat-plans_2.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-stat-plans` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.1 KiB | [postgresql-17-pg-stat-plans_2.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-17-pg-stat-plans_2.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-stat-plans` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 95.3 KiB | [postgresql-17-pg-stat-plans_2.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-17-pg-stat-plans_2.1.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-stat-plans` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 98.5 KiB | [postgresql-17-pg-stat-plans_2.1.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-17-pg-stat-plans_2.1.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-stat-plans` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 95.8 KiB | [postgresql-17-pg-stat-plans_2.1.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-17-pg-stat-plans_2.1.0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stat_plans_16` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.5 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_stat_plans_16-2.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_stat_plans_16-2.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.6 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_plans_16-2.1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.6 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_plans_16-2.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.8 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_stat_plans_16-2.1.0-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.5 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_plans_16-2.1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.5 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_plans_16-2.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.6 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_stat_plans_16-2.1.0-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.1 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_stat_plans_16-2.1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.0 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_stat_plans_16-2.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.4 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_stat_plans_16-2.1.0-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.7 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_plans_16-2.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.7 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_plans_16-2.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_stat_plans_16` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.7 KiB | [pg_stat_plans_16-2.1.0-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_stat_plans_16-2.1.0-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-16-pg-stat-plans` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 100.3 KiB | [postgresql-16-pg-stat-plans_2.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-16-pg-stat-plans_2.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-stat-plans` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 97.1 KiB | [postgresql-16-pg-stat-plans_2.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-16-pg-stat-plans_2.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-stat-plans` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 100.5 KiB | [postgresql-16-pg-stat-plans_2.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-16-pg-stat-plans_2.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-stat-plans` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 97.4 KiB | [postgresql-16-pg-stat-plans_2.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-16-pg-stat-plans_2.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-stat-plans` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.9 KiB | [postgresql-16-pg-stat-plans_2.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-16-pg-stat-plans_2.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-stat-plans` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.9 KiB | [postgresql-16-pg-stat-plans_2.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-16-pg-stat-plans_2.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-stat-plans` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.1 KiB | [postgresql-16-pg-stat-plans_2.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-16-pg-stat-plans_2.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-stat-plans` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 94.9 KiB | [postgresql-16-pg-stat-plans_2.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-16-pg-stat-plans_2.1.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-stat-plans` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 97.6 KiB | [postgresql-16-pg-stat-plans_2.1.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-16-pg-stat-plans_2.1.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-stat-plans` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 94.2 KiB | [postgresql-16-pg-stat-plans_2.1.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-stat-plans/postgresql-16-pg-stat-plans_2.1.0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pganalyze/pg_stat_plans" title="Repository" icon="github" subtitle="github.com/pganalyze/pg_stat_plans" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_stat_plans;		# install via package name, for the active PG version

pig install pg_stat_plans -v 18;   # install for PG 18
pig install pg_stat_plans -v 17;   # install for PG 17
pig install pg_stat_plans -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = '$libdir/pg_stat_plans';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_stat_plans;
```
