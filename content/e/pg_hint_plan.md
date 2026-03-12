---
title: "pg_hint_plan"
linkTitle: "pg_hint_plan"
description: "Give PostgreSQL ability to manually force some decisions in execution plans."
weight: 2780
categories: ["FEAT"]
width: full
---

[**pg_hint_plan**](https://github.com/ossc-db/pg_hint_plan) : Give PostgreSQL ability to manually force some decisions in execution plans.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2780** | {{< badge content="pg_hint_plan" link="https://github.com/ossc-db/pg_hint_plan" >}} | {{< ext "pg_hint_plan" >}} | `1.8.0` | {{< category "FEAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `hint_plan` |
|   **See Also**    | {{< ext "pg_show_plans" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_stat_statements" >}} {{< ext "hypopg" >}} {{< ext "pg_qualstats" >}} {{< ext "auto_explain" >}} {{< ext "index_advisor" >}} {{< ext "pg_profile" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.8.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_hint_plan` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.8.0` | {{< bg "18" "pg_hint_plan_18" "green" >}} {{< bg "17" "pg_hint_plan_17" "green" >}} {{< bg "16" "pg_hint_plan_16" "green" >}} {{< bg "15" "pg_hint_plan_15" "green" >}} {{< bg "14" "pg_hint_plan_14" "green" >}} | `pg_hint_plan_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.8.0` | {{< bg "18" "postgresql-18-pg-hint-plan" "green" >}} {{< bg "17" "postgresql-17-pg-hint-plan" "green" >}} {{< bg "16" "postgresql-16-pg-hint-plan" "green" >}} {{< bg "15" "postgresql-15-pg-hint-plan" "green" >}} {{< bg "14" "postgresql-14-pg-hint-plan" "green" >}} | `postgresql-$v-pg-hint-plan` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.8.0" "pg_hint_plan_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "pg_hint_plan_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.2" "pg_hint_plan_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_hint_plan_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.4" "pg_hint_plan_14 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.8.0" "pg_hint_plan_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "pg_hint_plan_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.2" "pg_hint_plan_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_hint_plan_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.4" "pg_hint_plan_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.8.0" "pg_hint_plan_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "pg_hint_plan_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.2" "pg_hint_plan_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_hint_plan_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.4" "pg_hint_plan_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.8.0" "pg_hint_plan_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "pg_hint_plan_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.2" "pg_hint_plan_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_hint_plan_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.4" "pg_hint_plan_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.8.0" "pg_hint_plan_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "pg_hint_plan_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.2" "pg_hint_plan_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_hint_plan_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.4" "pg_hint_plan_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.8.0" "pg_hint_plan_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "pg_hint_plan_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.2" "pg_hint_plan_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_hint_plan_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.4" "pg_hint_plan_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.8.0" "postgresql-18-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.2" "postgresql-16-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.4" "postgresql-14-pg-hint-plan : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.8.0" "postgresql-18-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.2" "postgresql-16-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.4" "postgresql-14-pg-hint-plan : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.8.0" "postgresql-18-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.2" "postgresql-16-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.4" "postgresql-14-pg-hint-plan : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.8.0" "postgresql-18-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.2" "postgresql-16-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.4" "postgresql-14-pg-hint-plan : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.8.0" "postgresql-18-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.2" "postgresql-16-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.4" "postgresql-14-pg-hint-plan : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.8.0" "postgresql-18-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.2" "postgresql-16-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.4" "postgresql-14-pg-hint-plan : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.8.0" "postgresql-18-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.2" "postgresql-16-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.4" "postgresql-14-pg-hint-plan : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.8.0" "postgresql-18-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.2" "postgresql-16-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-pg-hint-plan : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.4" "postgresql-14-pg-hint-plan : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hint_plan_18` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 55.3 KiB | [pg_hint_plan_18-1.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_hint_plan_18-1.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_hint_plan_18` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 51.7 KiB | [pg_hint_plan_18-1.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_hint_plan_18-1.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_hint_plan_18` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 54.4 KiB | [pg_hint_plan_18-1.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_hint_plan_18-1.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_hint_plan_18` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.5 KiB | [pg_hint_plan_18-1.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_hint_plan_18-1.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_hint_plan_18` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 55.5 KiB | [pg_hint_plan_18-1.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_hint_plan_18-1.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_hint_plan_18` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.7 KiB | [pg_hint_plan_18-1.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_hint_plan_18-1.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-hint-plan` | `1.8.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 132.8 KiB | [postgresql-18-pg-hint-plan_1.8.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-18/postgresql-18-pg-hint-plan_1.8.0-3.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-hint-plan` | `1.8.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 128.8 KiB | [postgresql-18-pg-hint-plan_1.8.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-18/postgresql-18-pg-hint-plan_1.8.0-3.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-hint-plan` | `1.8.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 133.2 KiB | [postgresql-18-pg-hint-plan_1.8.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-18/postgresql-18-pg-hint-plan_1.8.0-3.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-hint-plan` | `1.8.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 128.8 KiB | [postgresql-18-pg-hint-plan_1.8.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-18/postgresql-18-pg-hint-plan_1.8.0-3.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-hint-plan` | `1.8.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 133.8 KiB | [postgresql-18-pg-hint-plan_1.8.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-18/postgresql-18-pg-hint-plan_1.8.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-hint-plan` | `1.8.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 129.8 KiB | [postgresql-18-pg-hint-plan_1.8.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-18/postgresql-18-pg-hint-plan_1.8.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-hint-plan` | `1.8.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 131.7 KiB | [postgresql-18-pg-hint-plan_1.8.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-18/postgresql-18-pg-hint-plan_1.8.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-hint-plan` | `1.8.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 128.0 KiB | [postgresql-18-pg-hint-plan_1.8.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-18/postgresql-18-pg-hint-plan_1.8.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hint_plan_17` | `1.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 53.2 KiB | [pg_hint_plan_17-1.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_hint_plan_17-1.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_hint_plan_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.4 KiB | [pg_hint_plan_17-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_hint_plan_17-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_hint_plan_17` | `1.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.0 KiB | [pg_hint_plan_17-1.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_hint_plan_17-1.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_hint_plan_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.2 KiB | [pg_hint_plan_17-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_hint_plan_17-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_hint_plan_17` | `1.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.6 KiB | [pg_hint_plan_17-1.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_hint_plan_17-1.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_hint_plan_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.7 KiB | [pg_hint_plan_17-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_hint_plan_17-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_hint_plan_17` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.6 KiB | [pg_hint_plan_17-1.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_hint_plan_17-1.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_hint_plan_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.0 KiB | [pg_hint_plan_17-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_hint_plan_17-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_hint_plan_17` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.6 KiB | [pg_hint_plan_17-1.7.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_hint_plan_17-1.7.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_hint_plan_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.0 KiB | [pg_hint_plan_17-1.7.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_hint_plan_17-1.7.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_hint_plan_17` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.1 KiB | [pg_hint_plan_17-1.7.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_hint_plan_17-1.7.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_hint_plan_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.4 KiB | [pg_hint_plan_17-1.7.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_hint_plan_17-1.7.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-hint-plan` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 128.0 KiB | [postgresql-17-pg-hint-plan_1.7.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-17/postgresql-17-pg-hint-plan_1.7.1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-hint-plan` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 123.9 KiB | [postgresql-17-pg-hint-plan_1.7.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-17/postgresql-17-pg-hint-plan_1.7.1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-hint-plan` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 128.4 KiB | [postgresql-17-pg-hint-plan_1.7.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-17/postgresql-17-pg-hint-plan_1.7.1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-hint-plan` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 124.3 KiB | [postgresql-17-pg-hint-plan_1.7.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-17/postgresql-17-pg-hint-plan_1.7.1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-hint-plan` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 140.9 KiB | [postgresql-17-pg-hint-plan_1.7.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-17/postgresql-17-pg-hint-plan_1.7.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-hint-plan` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 137.2 KiB | [postgresql-17-pg-hint-plan_1.7.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-17/postgresql-17-pg-hint-plan_1.7.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-hint-plan` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 126.7 KiB | [postgresql-17-pg-hint-plan_1.7.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-17/postgresql-17-pg-hint-plan_1.7.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-hint-plan` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.4 KiB | [postgresql-17-pg-hint-plan_1.7.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-17/postgresql-17-pg-hint-plan_1.7.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hint_plan_16` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.8 KiB | [pg_hint_plan_16-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_hint_plan_16-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_hint_plan_16` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.2 KiB | [pg_hint_plan_16-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_hint_plan_16-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_hint_plan_16` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.6 KiB | [pg_hint_plan_16-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_hint_plan_16-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_hint_plan_16` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.4 KiB | [pg_hint_plan_16-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_hint_plan_16-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_hint_plan_16` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.8 KiB | [pg_hint_plan_16-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_hint_plan_16-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_hint_plan_16` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.2 KiB | [pg_hint_plan_16-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_hint_plan_16-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_hint_plan_16` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.5 KiB | [pg_hint_plan_16-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_hint_plan_16-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_hint_plan_16` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.9 KiB | [pg_hint_plan_16-1.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_hint_plan_16-1.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_hint_plan_16` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.2 KiB | [pg_hint_plan_16-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_hint_plan_16-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_hint_plan_16` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.0 KiB | [pg_hint_plan_16-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_hint_plan_16-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_hint_plan_16` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.5 KiB | [pg_hint_plan_16-1.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_hint_plan_16-1.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_hint_plan_16` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [pg_hint_plan_16-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_hint_plan_16-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_hint_plan_16` | `1.6.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.7 KiB | [pg_hint_plan_16-1.6.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_hint_plan_16-1.6.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_hint_plan_16` | `1.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.2 KiB | [pg_hint_plan_16-1.6.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_hint_plan_16-1.6.1-3PGDG.rhel10.x86_64.rpm) |
| `pg_hint_plan_16` | `1.6.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.0 KiB | [pg_hint_plan_16-1.6.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_hint_plan_16-1.6.2-1PGDG.rhel10.aarch64.rpm) |
| `pg_hint_plan_16` | `1.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.5 KiB | [pg_hint_plan_16-1.6.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_hint_plan_16-1.6.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-hint-plan` | `1.6.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 104.6 KiB | [postgresql-16-pg-hint-plan_1.6.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-16/postgresql-16-pg-hint-plan_1.6.2-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-hint-plan` | `1.6.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 101.6 KiB | [postgresql-16-pg-hint-plan_1.6.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-16/postgresql-16-pg-hint-plan_1.6.2-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-hint-plan` | `1.6.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 104.8 KiB | [postgresql-16-pg-hint-plan_1.6.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-16/postgresql-16-pg-hint-plan_1.6.2-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-hint-plan` | `1.6.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 101.9 KiB | [postgresql-16-pg-hint-plan_1.6.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-16/postgresql-16-pg-hint-plan_1.6.2-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-hint-plan` | `1.6.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 117.2 KiB | [postgresql-16-pg-hint-plan_1.6.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-16/postgresql-16-pg-hint-plan_1.6.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-hint-plan` | `1.6.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 114.4 KiB | [postgresql-16-pg-hint-plan_1.6.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-16/postgresql-16-pg-hint-plan_1.6.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-hint-plan` | `1.6.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 104.4 KiB | [postgresql-16-pg-hint-plan_1.6.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-16/postgresql-16-pg-hint-plan_1.6.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-hint-plan` | `1.6.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 101.7 KiB | [postgresql-16-pg-hint-plan_1.6.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-16/postgresql-16-pg-hint-plan_1.6.2-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hint_plan_15` | `1.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.3 KiB | [pg_hint_plan_15-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_hint_plan_15-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_hint_plan_15` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.7 KiB | [pg_hint_plan_15-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_hint_plan_15-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_hint_plan_15` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 48.9 KiB | [pg_hint_plan_15-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_hint_plan_15-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_hint_plan_15` | `1.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.0 KiB | [pg_hint_plan_15-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_hint_plan_15-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_hint_plan_15` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.5 KiB | [pg_hint_plan_15-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_hint_plan_15-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_hint_plan_15` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.6 KiB | [pg_hint_plan_15-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_hint_plan_15-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_hint_plan_15` | `1.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.0 KiB | [pg_hint_plan_15-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_hint_plan_15-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_hint_plan_15` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.7 KiB | [pg_hint_plan_15-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_hint_plan_15-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_hint_plan_15` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.4 KiB | [pg_hint_plan_15-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_hint_plan_15-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_hint_plan_15` | `1.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.5 KiB | [pg_hint_plan_15-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_hint_plan_15-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_hint_plan_15` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.1 KiB | [pg_hint_plan_15-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_hint_plan_15-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_hint_plan_15` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 47.8 KiB | [pg_hint_plan_15-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_hint_plan_15-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_hint_plan_15` | `1.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.4 KiB | [pg_hint_plan_15-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_hint_plan_15-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_hint_plan_15` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.2 KiB | [pg_hint_plan_15-1.5.2-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_hint_plan_15-1.5.2-3PGDG.rhel10.x86_64.rpm) |
| `pg_hint_plan_15` | `1.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.5 KiB | [pg_hint_plan_15-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_hint_plan_15-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_hint_plan_15` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.2 KiB | [pg_hint_plan_15-1.5.2-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_hint_plan_15-1.5.2-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-hint-plan` | `1.5.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 104.4 KiB | [postgresql-15-pg-hint-plan_1.5.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-15/postgresql-15-pg-hint-plan_1.5.3-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-hint-plan` | `1.5.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 101.2 KiB | [postgresql-15-pg-hint-plan_1.5.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-15/postgresql-15-pg-hint-plan_1.5.3-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-hint-plan` | `1.5.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 104.8 KiB | [postgresql-15-pg-hint-plan_1.5.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-15/postgresql-15-pg-hint-plan_1.5.3-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-hint-plan` | `1.5.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 101.7 KiB | [postgresql-15-pg-hint-plan_1.5.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-15/postgresql-15-pg-hint-plan_1.5.3-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-hint-plan` | `1.5.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 117.1 KiB | [postgresql-15-pg-hint-plan_1.5.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-15/postgresql-15-pg-hint-plan_1.5.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-hint-plan` | `1.5.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 114.0 KiB | [postgresql-15-pg-hint-plan_1.5.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-15/postgresql-15-pg-hint-plan_1.5.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-hint-plan` | `1.5.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 104.4 KiB | [postgresql-15-pg-hint-plan_1.5.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-15/postgresql-15-pg-hint-plan_1.5.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-hint-plan` | `1.5.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 101.4 KiB | [postgresql-15-pg-hint-plan_1.5.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-15/postgresql-15-pg-hint-plan_1.5.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hint_plan_14` | `1.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.6 KiB | [pg_hint_plan_14-1.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_hint_plan_14-1.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_hint_plan_14` | `1.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.3 KiB | [pg_hint_plan_14-1.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_hint_plan_14-1.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_hint_plan_14` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.0 KiB | [pg_hint_plan_14-1.4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_hint_plan_14-1.4.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_hint_plan_14` | `1.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.3 KiB | [pg_hint_plan_14-1.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_hint_plan_14-1.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_hint_plan_14` | `1.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.0 KiB | [pg_hint_plan_14-1.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_hint_plan_14-1.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_hint_plan_14` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.7 KiB | [pg_hint_plan_14-1.4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_hint_plan_14-1.4.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_hint_plan_14` | `1.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.4 KiB | [pg_hint_plan_14-1.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_hint_plan_14-1.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_hint_plan_14` | `1.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.2 KiB | [pg_hint_plan_14-1.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_hint_plan_14-1.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_hint_plan_14` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.8 KiB | [pg_hint_plan_14-1.4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_hint_plan_14-1.4.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_hint_plan_14` | `1.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.9 KiB | [pg_hint_plan_14-1.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_hint_plan_14-1.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_hint_plan_14` | `1.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.7 KiB | [pg_hint_plan_14-1.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_hint_plan_14-1.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_hint_plan_14` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.2 KiB | [pg_hint_plan_14-1.4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_hint_plan_14-1.4.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_hint_plan_14` | `1.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.8 KiB | [pg_hint_plan_14-1.4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_hint_plan_14-1.4.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_hint_plan_14` | `1.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.6 KiB | [pg_hint_plan_14-1.4.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_hint_plan_14-1.4.3-3PGDG.rhel10.x86_64.rpm) |
| `pg_hint_plan_14` | `1.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.9 KiB | [pg_hint_plan_14-1.4.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_hint_plan_14-1.4.4-1PGDG.rhel10.aarch64.rpm) |
| `pg_hint_plan_14` | `1.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.7 KiB | [pg_hint_plan_14-1.4.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_hint_plan_14-1.4.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-hint-plan` | `1.4.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 104.1 KiB | [postgresql-14-pg-hint-plan_1.4.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-14/postgresql-14-pg-hint-plan_1.4.4-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-hint-plan` | `1.4.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 100.8 KiB | [postgresql-14-pg-hint-plan_1.4.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-14/postgresql-14-pg-hint-plan_1.4.4-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-hint-plan` | `1.4.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 104.5 KiB | [postgresql-14-pg-hint-plan_1.4.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-14/postgresql-14-pg-hint-plan_1.4.4-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-hint-plan` | `1.4.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 101.1 KiB | [postgresql-14-pg-hint-plan_1.4.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-14/postgresql-14-pg-hint-plan_1.4.4-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-hint-plan` | `1.4.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 116.6 KiB | [postgresql-14-pg-hint-plan_1.4.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-14/postgresql-14-pg-hint-plan_1.4.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-hint-plan` | `1.4.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 113.5 KiB | [postgresql-14-pg-hint-plan_1.4.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-14/postgresql-14-pg-hint-plan_1.4.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-hint-plan` | `1.4.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 104.4 KiB | [postgresql-14-pg-hint-plan_1.4.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-14/postgresql-14-pg-hint-plan_1.4.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-hint-plan` | `1.4.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 101.1 KiB | [postgresql-14-pg-hint-plan_1.4.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-hint-plan-14/postgresql-14-pg-hint-plan_1.4.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ossc-db/pg_hint_plan" title="Repository" icon="github" subtitle="github.com/ossc-db/pg_hint_plan" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_hint_plan-REL18_1_8_0.tar.gz" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_hint_plan;		# install via package name, for the active PG version

pig install pg_hint_plan -v 18;   # install for PG 18
pig install pg_hint_plan -v 17;   # install for PG 17
pig install pg_hint_plan -v 16;   # install for PG 16
pig install pg_hint_plan -v 15;   # install for PG 15
pig install pg_hint_plan -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_hint_plan;
```




## Usage

> [pg_hint_plan: Give PostgreSQL ability to manually force some decisions in execution plans](https://github.com/ossc-db/pg_hint_plan)

`pg_hint_plan` overrides the PostgreSQL query planner's decisions using SQL comment hints, allowing you to force specific scan methods, join strategies, and join orders.

### Hint Syntax

Hints are embedded in SQL comments prefixed with `/*+` and closed with `*/`:

```sql
/*+
  HashJoin(a b)
  SeqScan(a)
*/
SELECT *
FROM pgbench_branches b
JOIN pgbench_accounts a ON b.bid = a.bid
ORDER BY a.aid;
```

### Scan Method Hints

| Hint | Description |
|------|-------------|
| `SeqScan(table)` | Force sequential scan |
| `IndexScan(table [index...])` | Force index scan (optionally restrict to specific indexes) |
| `IndexOnlyScan(table [index...])` | Force index-only scan |
| `BitmapScan(table [index...])` | Force bitmap scan |
| `TidScan(table)` | Force TID scan |
| `NoSeqScan(table)` | Prevent sequential scan |
| `NoIndexScan(table)` | Prevent index scan and index-only scan |
| `NoIndexOnlyScan(table)` | Prevent index-only scan |
| `NoBitmapScan(table)` | Prevent bitmap scan |
| `NoTidScan(table)` | Prevent TID scan |
| `IndexScanRegexp(table [regexp...])` | Force index scan with regex-matched index names |
| `DisableIndex(table index...)` | Disable specific indexes during planning |

### Join Method Hints

| Hint | Description |
|------|-------------|
| `NestLoop(t1 t2 [t3...])` | Force nested loop join |
| `HashJoin(t1 t2 [t3...])` | Force hash join |
| `MergeJoin(t1 t2 [t3...])` | Force merge join |
| `NoNestLoop(t1 t2 [t3...])` | Prevent nested loop join |
| `NoHashJoin(t1 t2 [t3...])` | Prevent hash join |
| `NoMergeJoin(t1 t2 [t3...])` | Prevent merge join |
| `Memoize(t1 t2 [t3...])` | Allow memoization of inner result |
| `NoMemoize(t1 t2 [t3...])` | Prevent memoization |

### Join Order Hints

```sql
-- Simple left-deep join order
/*+ Leading(a b c) */

-- Explicit join tree with nesting
/*+ Leading((a (b c))) */
```

### Row Number Correction

```sql
/*+ Rows(a b #100) */    -- Set to absolute number
/*+ Rows(a b +100) */    -- Add to estimate
/*+ Rows(a b -100) */    -- Subtract from estimate
/*+ Rows(a b *2.0) */    -- Multiply estimate
```

### Parallel Query Control

```sql
/*+ Parallel(table 4 hard) */   -- Force 4 parallel workers
/*+ Parallel(table 0 hard) */   -- Disable parallelism
```

### GUC Parameter Override

```sql
/*+ Set(random_page_cost 1.0) Set(seq_page_cost 1.0) */
SELECT * FROM my_table WHERE id = 42;
```

### GUC Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `pg_hint_plan.enable_hint` | Enable or disable hints globally | on |
| `pg_hint_plan.enable_hint_table` | Enable hint table for query-based hints | off |
| `pg_hint_plan.debug_print` | Print debug info for applied hints | off |
| `pg_hint_plan.parse_messages` | Log level for hint parsing messages | INFO |
| `pg_hint_plan.message_level` | Log level for debug messages | LOG |
