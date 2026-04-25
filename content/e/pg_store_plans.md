---
title: "pg_store_plans"
linkTitle: "pg_store_plans"
description: "track plan statistics of all SQL statements executed"
weight: 6250
categories: ["STAT"]
width: full
---

[**pg_store_plans**](https://github.com/ossc-db/pg_store_plans) : track plan statistics of all SQL statements executed


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6250** | {{< badge content="pg_store_plans" link="https://github.com/ossc-db/pg_store_plans" >}} | {{< ext "pg_store_plans" >}} | `1.10` | {{< category "STAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_show_plans" >}} {{< ext "auto_explain" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_hint_plan" >}} {{< ext "pre_prepare" >}} {{< ext "pg_stat_monitor" >}} {{< ext "explain_ui" >}} {{< ext "plprofiler" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.10` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_store_plans` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.10` | {{< bg "18" "pg_store_plans_18" "green" >}} {{< bg "17" "pg_store_plans_17" "green" >}} {{< bg "16" "pg_store_plans_16" "green" >}} {{< bg "15" "pg_store_plans_15" "green" >}} {{< bg "14" "pg_store_plans_14" "green" >}} | `pg_store_plans_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.10` | {{< bg "18" "postgresql-18-pg-store-plan" "green" >}} {{< bg "17" "postgresql-17-pg-store-plan" "green" >}} {{< bg "16" "postgresql-16-pg-store-plan" "green" >}} {{< bg "15" "postgresql-15-pg-store-plan" "green" >}} {{< bg "14" "postgresql-14-pg-store-plan" "green" >}} | `postgresql-$v-pg-store-plan` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.10" "pg_store_plans_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.10" "postgresql-18-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-17-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-16-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-15-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-14-pg-store-plan : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.10" "postgresql-18-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-17-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-16-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-15-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-14-pg-store-plan : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.10" "postgresql-18-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-17-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-16-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-15-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-14-pg-store-plan : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.10" "postgresql-18-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-17-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-16-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-15-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-14-pg-store-plan : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.10" "postgresql-18-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-17-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-16-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-15-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-14-pg-store-plan : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.10" "postgresql-18-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-17-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-16-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-15-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-14-pg-store-plan : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.10" "postgresql-18-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-17-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-16-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-15-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-14-pg-store-plan : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.10" "postgresql-18-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-17-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-16-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-15-pg-store-plan : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.10" "postgresql-14-pg-store-plan : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_store_plans_18` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.6 KiB | [pg_store_plans_18-1.10-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_store_plans_18-1.10-1PIGSTY.el8.x86_64.rpm) |
| `pg_store_plans_18` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.9 KiB | [pg_store_plans_18-1.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_store_plans_18-1.10-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_store_plans_18` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.4 KiB | [pg_store_plans_18-1.10-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_store_plans_18-1.10-1PIGSTY.el8.aarch64.rpm) |
| `pg_store_plans_18` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.6 KiB | [pg_store_plans_18-1.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_store_plans_18-1.10-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_store_plans_18` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.5 KiB | [pg_store_plans_18-1.10-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_store_plans_18-1.10-1PIGSTY.el9.x86_64.rpm) |
| `pg_store_plans_18` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.1 KiB | [pg_store_plans_18-1.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_store_plans_18-1.10-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_store_plans_18` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.1 KiB | [pg_store_plans_18-1.10-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_store_plans_18-1.10-1PIGSTY.el9.aarch64.rpm) |
| `pg_store_plans_18` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.1 KiB | [pg_store_plans_18-1.10-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_store_plans_18-1.10-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_store_plans_18` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pigsty | 44.2 KiB | [pg_store_plans_18-1.10-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_store_plans_18-1.10-1PIGSTY.el10.x86_64.rpm) |
| `pg_store_plans_18` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.9 KiB | [pg_store_plans_18-1.10-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_store_plans_18-1.10-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_store_plans_18` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.8 KiB | [pg_store_plans_18-1.10-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_store_plans_18-1.10-1PIGSTY.el10.aarch64.rpm) |
| `pg_store_plans_18` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.8 KiB | [pg_store_plans_18-1.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_store_plans_18-1.10-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-18-pg-store-plan` | `1.10` | [d12.x86_64](/os/d12.x86_64) | pigsty | 108.2 KiB | [postgresql-18-pg-store-plan_1.10-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-18-pg-store-plan_1.10-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-store-plan` | `1.10` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.2 KiB | [postgresql-18-pg-store-plan_1.10-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-18-pg-store-plan_1.10-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-store-plan` | `1.10` | [d13.x86_64](/os/d13.x86_64) | pigsty | 108.8 KiB | [postgresql-18-pg-store-plan_1.10-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-store-plan/postgresql-18-pg-store-plan_1.10-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-store-plan` | `1.10` | [d13.aarch64](/os/d13.aarch64) | pigsty | 105.4 KiB | [postgresql-18-pg-store-plan_1.10-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-store-plan/postgresql-18-pg-store-plan_1.10-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-store-plan` | `1.10` | [u22.x86_64](/os/u22.x86_64) | pigsty | 118.0 KiB | [postgresql-18-pg-store-plan_1.10-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-18-pg-store-plan_1.10-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-store-plan` | `1.10` | [u22.aarch64](/os/u22.aarch64) | pigsty | 116.5 KiB | [postgresql-18-pg-store-plan_1.10-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-18-pg-store-plan_1.10-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-store-plan` | `1.10` | [u24.x86_64](/os/u24.x86_64) | pigsty | 113.8 KiB | [postgresql-18-pg-store-plan_1.10-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-18-pg-store-plan_1.10-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-store-plan` | `1.10` | [u24.aarch64](/os/u24.aarch64) | pigsty | 112.5 KiB | [postgresql-18-pg-store-plan_1.10-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-18-pg-store-plan_1.10-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_store_plans_17` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.7 KiB | [pg_store_plans_17-1.10-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_store_plans_17-1.10-1PIGSTY.el8.x86_64.rpm) |
| `pg_store_plans_17` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.9 KiB | [pg_store_plans_17-1.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_store_plans_17-1.10-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_store_plans_17` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.4 KiB | [pg_store_plans_17-1.10-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_store_plans_17-1.10-1PIGSTY.el8.aarch64.rpm) |
| `pg_store_plans_17` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.6 KiB | [pg_store_plans_17-1.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_store_plans_17-1.10-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_store_plans_17` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.7 KiB | [pg_store_plans_17-1.10-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_store_plans_17-1.10-1PIGSTY.el9.x86_64.rpm) |
| `pg_store_plans_17` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.1 KiB | [pg_store_plans_17-1.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_store_plans_17-1.10-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_store_plans_17` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.2 KiB | [pg_store_plans_17-1.10-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_store_plans_17-1.10-1PIGSTY.el9.aarch64.rpm) |
| `pg_store_plans_17` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.1 KiB | [pg_store_plans_17-1.10-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_store_plans_17-1.10-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_store_plans_17` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pigsty | 44.9 KiB | [pg_store_plans_17-1.10-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_store_plans_17-1.10-1PIGSTY.el10.x86_64.rpm) |
| `pg_store_plans_17` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.0 KiB | [pg_store_plans_17-1.10-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_store_plans_17-1.10-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_store_plans_17` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.9 KiB | [pg_store_plans_17-1.10-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_store_plans_17-1.10-1PIGSTY.el10.aarch64.rpm) |
| `pg_store_plans_17` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_store_plans_17-1.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_store_plans_17-1.10-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-17-pg-store-plan` | `1.10` | [d12.x86_64](/os/d12.x86_64) | pigsty | 108.0 KiB | [postgresql-17-pg-store-plan_1.10-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.10-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-store-plan` | `1.10` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.2 KiB | [postgresql-17-pg-store-plan_1.10-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.10-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-store-plan` | `1.10` | [d13.x86_64](/os/d13.x86_64) | pigsty | 108.6 KiB | [postgresql-17-pg-store-plan_1.10-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.10-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-store-plan` | `1.10` | [d13.aarch64](/os/d13.aarch64) | pigsty | 105.6 KiB | [postgresql-17-pg-store-plan_1.10-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.10-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-store-plan` | `1.10` | [u22.x86_64](/os/u22.x86_64) | pigsty | 129.5 KiB | [postgresql-17-pg-store-plan_1.10-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.10-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-store-plan` | `1.10` | [u22.aarch64](/os/u22.aarch64) | pigsty | 127.9 KiB | [postgresql-17-pg-store-plan_1.10-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.10-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-store-plan` | `1.10` | [u24.x86_64](/os/u24.x86_64) | pigsty | 113.9 KiB | [postgresql-17-pg-store-plan_1.10-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.10-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-store-plan` | `1.10` | [u24.aarch64](/os/u24.aarch64) | pigsty | 112.7 KiB | [postgresql-17-pg-store-plan_1.10-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.10-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_store_plans_16` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.7 KiB | [pg_store_plans_16-1.10-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_store_plans_16-1.10-1PIGSTY.el8.x86_64.rpm) |
| `pg_store_plans_16` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.0 KiB | [pg_store_plans_16-1.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_store_plans_16-1.10-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_store_plans_16` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.4 KiB | [pg_store_plans_16-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_store_plans_16-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_store_plans_16` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.4 KiB | [pg_store_plans_16-1.10-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_store_plans_16-1.10-1PIGSTY.el8.aarch64.rpm) |
| `pg_store_plans_16` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.6 KiB | [pg_store_plans_16-1.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_store_plans_16-1.10-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_store_plans_16` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.1 KiB | [pg_store_plans_16-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_store_plans_16-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_store_plans_16` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.7 KiB | [pg_store_plans_16-1.10-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_store_plans_16-1.10-1PIGSTY.el9.x86_64.rpm) |
| `pg_store_plans_16` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.1 KiB | [pg_store_plans_16-1.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_store_plans_16-1.10-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_store_plans_16` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.1 KiB | [pg_store_plans_16-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_store_plans_16-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_store_plans_16` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.2 KiB | [pg_store_plans_16-1.10-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_store_plans_16-1.10-1PIGSTY.el9.aarch64.rpm) |
| `pg_store_plans_16` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.1 KiB | [pg_store_plans_16-1.10-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_store_plans_16-1.10-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_store_plans_16` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.1 KiB | [pg_store_plans_16-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_store_plans_16-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_store_plans_16` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pigsty | 44.8 KiB | [pg_store_plans_16-1.10-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_store_plans_16-1.10-1PIGSTY.el10.x86_64.rpm) |
| `pg_store_plans_16` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.9 KiB | [pg_store_plans_16-1.10-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_store_plans_16-1.10-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_store_plans_16` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.1 KiB | [pg_store_plans_16-1.8-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_store_plans_16-1.8-3PGDG.rhel10.x86_64.rpm) |
| `pg_store_plans_16` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.8 KiB | [pg_store_plans_16-1.10-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_store_plans_16-1.10-1PIGSTY.el10.aarch64.rpm) |
| `pg_store_plans_16` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_store_plans_16-1.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_store_plans_16-1.10-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_store_plans_16` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.0 KiB | [pg_store_plans_16-1.8-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_store_plans_16-1.8-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-store-plan` | `1.10` | [d12.x86_64](/os/d12.x86_64) | pigsty | 108.4 KiB | [postgresql-16-pg-store-plan_1.10-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.10-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-store-plan` | `1.10` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.4 KiB | [postgresql-16-pg-store-plan_1.10-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.10-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-store-plan` | `1.10` | [d13.x86_64](/os/d13.x86_64) | pigsty | 109.1 KiB | [postgresql-16-pg-store-plan_1.10-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.10-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-store-plan` | `1.10` | [d13.aarch64](/os/d13.aarch64) | pigsty | 105.7 KiB | [postgresql-16-pg-store-plan_1.10-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.10-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-store-plan` | `1.10` | [u22.x86_64](/os/u22.x86_64) | pigsty | 129.2 KiB | [postgresql-16-pg-store-plan_1.10-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.10-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-store-plan` | `1.10` | [u22.aarch64](/os/u22.aarch64) | pigsty | 127.6 KiB | [postgresql-16-pg-store-plan_1.10-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.10-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-store-plan` | `1.10` | [u24.x86_64](/os/u24.x86_64) | pigsty | 114.0 KiB | [postgresql-16-pg-store-plan_1.10-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.10-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-store-plan` | `1.10` | [u24.aarch64](/os/u24.aarch64) | pigsty | 113.0 KiB | [postgresql-16-pg-store-plan_1.10-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.10-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_store_plans_15` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 46.2 KiB | [pg_store_plans_15-1.10-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_store_plans_15-1.10-1PIGSTY.el8.x86_64.rpm) |
| `pg_store_plans_15` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.8 KiB | [pg_store_plans_15-1.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_store_plans_15-1.10-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_store_plans_15` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.3 KiB | [pg_store_plans_15-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_store_plans_15-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_store_plans_15` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.3 KiB | [pg_store_plans_15-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_store_plans_15-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_store_plans_15` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 44.9 KiB | [pg_store_plans_15-1.10-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_store_plans_15-1.10-1PIGSTY.el8.aarch64.rpm) |
| `pg_store_plans_15` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.4 KiB | [pg_store_plans_15-1.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_store_plans_15-1.10-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_store_plans_15` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.0 KiB | [pg_store_plans_15-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_store_plans_15-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_store_plans_15` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.8 KiB | [pg_store_plans_15-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_store_plans_15-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_store_plans_15` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.6 KiB | [pg_store_plans_15-1.10-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_store_plans_15-1.10-1PIGSTY.el9.x86_64.rpm) |
| `pg_store_plans_15` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 47.1 KiB | [pg_store_plans_15-1.10-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_store_plans_15-1.10-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_store_plans_15` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 47.2 KiB | [pg_store_plans_15-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_store_plans_15-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_store_plans_15` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 47.3 KiB | [pg_store_plans_15-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_store_plans_15-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_store_plans_15` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 45.1 KiB | [pg_store_plans_15-1.10-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_store_plans_15-1.10-1PIGSTY.el9.aarch64.rpm) |
| `pg_store_plans_15` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.2 KiB | [pg_store_plans_15-1.10-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_store_plans_15-1.10-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_store_plans_15` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.1 KiB | [pg_store_plans_15-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_store_plans_15-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_store_plans_15` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.9 KiB | [pg_store_plans_15-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_store_plans_15-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_store_plans_15` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pigsty | 46.5 KiB | [pg_store_plans_15-1.10-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_store_plans_15-1.10-1PIGSTY.el10.x86_64.rpm) |
| `pg_store_plans_15` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 48.1 KiB | [pg_store_plans_15-1.10-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_store_plans_15-1.10-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_store_plans_15` | `1.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 48.3 KiB | [pg_store_plans_15-1.8-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_store_plans_15-1.8-3PGDG.rhel10.x86_64.rpm) |
| `pg_store_plans_15` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pigsty | 45.7 KiB | [pg_store_plans_15-1.10-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_store_plans_15-1.10-1PIGSTY.el10.aarch64.rpm) |
| `pg_store_plans_15` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.9 KiB | [pg_store_plans_15-1.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_store_plans_15-1.10-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_store_plans_15` | `1.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 47.2 KiB | [pg_store_plans_15-1.8-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_store_plans_15-1.8-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-store-plan` | `1.10` | [d12.x86_64](/os/d12.x86_64) | pigsty | 109.8 KiB | [postgresql-15-pg-store-plan_1.10-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.10-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-store-plan` | `1.10` | [d12.aarch64](/os/d12.aarch64) | pigsty | 106.2 KiB | [postgresql-15-pg-store-plan_1.10-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.10-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-store-plan` | `1.10` | [d13.x86_64](/os/d13.x86_64) | pigsty | 110.2 KiB | [postgresql-15-pg-store-plan_1.10-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.10-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-store-plan` | `1.10` | [d13.aarch64](/os/d13.aarch64) | pigsty | 106.8 KiB | [postgresql-15-pg-store-plan_1.10-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.10-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-store-plan` | `1.10` | [u22.x86_64](/os/u22.x86_64) | pigsty | 131.0 KiB | [postgresql-15-pg-store-plan_1.10-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.10-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-store-plan` | `1.10` | [u22.aarch64](/os/u22.aarch64) | pigsty | 129.2 KiB | [postgresql-15-pg-store-plan_1.10-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.10-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-store-plan` | `1.10` | [u24.x86_64](/os/u24.x86_64) | pigsty | 116.3 KiB | [postgresql-15-pg-store-plan_1.10-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.10-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-store-plan` | `1.10` | [u24.aarch64](/os/u24.aarch64) | pigsty | 114.1 KiB | [postgresql-15-pg-store-plan_1.10-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.10-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_store_plans_14` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 46.0 KiB | [pg_store_plans_14-1.10-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_store_plans_14-1.10-1PIGSTY.el8.x86_64.rpm) |
| `pg_store_plans_14` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 44.9 KiB | [pg_store_plans_14-1.10-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_store_plans_14-1.10-1PIGSTY.el8.aarch64.rpm) |
| `pg_store_plans_14` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 45.1 KiB | [pg_store_plans_14-1.10-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_store_plans_14-1.10-1PIGSTY.el9.x86_64.rpm) |
| `pg_store_plans_14` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 45.0 KiB | [pg_store_plans_14-1.10-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_store_plans_14-1.10-1PIGSTY.el9.aarch64.rpm) |
| `pg_store_plans_14` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pigsty | 46.5 KiB | [pg_store_plans_14-1.10-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_store_plans_14-1.10-1PIGSTY.el10.x86_64.rpm) |
| `pg_store_plans_14` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pigsty | 45.7 KiB | [pg_store_plans_14-1.10-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_store_plans_14-1.10-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-store-plan` | `1.10` | [d12.x86_64](/os/d12.x86_64) | pigsty | 109.7 KiB | [postgresql-14-pg-store-plan_1.10-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.10-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-store-plan` | `1.10` | [d12.aarch64](/os/d12.aarch64) | pigsty | 106.1 KiB | [postgresql-14-pg-store-plan_1.10-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.10-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-store-plan` | `1.10` | [d13.x86_64](/os/d13.x86_64) | pigsty | 109.7 KiB | [postgresql-14-pg-store-plan_1.10-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.10-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-store-plan` | `1.10` | [d13.aarch64](/os/d13.aarch64) | pigsty | 106.7 KiB | [postgresql-14-pg-store-plan_1.10-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.10-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-store-plan` | `1.10` | [u22.x86_64](/os/u22.x86_64) | pigsty | 130.6 KiB | [postgresql-14-pg-store-plan_1.10-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.10-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-store-plan` | `1.10` | [u22.aarch64](/os/u22.aarch64) | pigsty | 128.8 KiB | [postgresql-14-pg-store-plan_1.10-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.10-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-store-plan` | `1.10` | [u24.x86_64](/os/u24.x86_64) | pigsty | 116.0 KiB | [postgresql-14-pg-store-plan_1.10-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.10-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-store-plan` | `1.10` | [u24.aarch64](/os/u24.aarch64) | pigsty | 114.3 KiB | [postgresql-14-pg-store-plan_1.10-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.10-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ossc-db/pg_store_plans" title="Repository" icon="github" subtitle="github.com/ossc-db/pg_store_plans" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_store_plans-1.10.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_store_plans;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_store_plans;		# install via package name, for the active PG version

pig install pg_store_plans -v 18;   # install for PG 18
pig install pg_store_plans -v 17;   # install for PG 17
pig install pg_store_plans -v 16;   # install for PG 16
pig install pg_store_plans -v 15;   # install for PG 15
pig install pg_store_plans -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_store_plans';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_store_plans;
```

## Usage

Sources: [official docs](https://ossc-db.github.io/pg_store_plans/), [repo](https://github.com/ossc-db/pg_store_plans), [1.10 release notes](https://github.com/ossc-db/pg_store_plans/releases/tag/1.10)

`pg_store_plans` tracks execution plan statistics for SQL statements, similar to how `pg_stat_statements` tracks statement statistics. The upstream `1.10` release note says this version adds PostgreSQL 18 support; the documented SQL surface is otherwise the same as current docs.

### Required server settings

```ini
shared_preload_libraries = 'pg_store_plans'
compute_query_id = 'on'
```

`pg_store_plans` requires shared memory, so adding or removing it needs a server restart. The docs say it is silently disabled if `compute_query_id` is `no`.

### Inspect stored plans

```sql
SELECT queryid, planid, plan, calls, total_time, rows
FROM pg_store_plans
ORDER BY total_time DESC;

SELECT * FROM pg_store_plans_info;
```

The docs describe `queryid` as the join key for `pg_stat_statements`, and `pg_store_plans_info` as a one-row view that exposes module-level stats such as `dealloc` and `stats_reset`.

### Helper functions

```sql
SELECT pg_store_plans_reset();
SELECT pg_store_hash_query('SELECT 1');
SELECT pg_store_plans_textplan(plan);
SELECT pg_store_plans_jsonplan(plan);
SELECT pg_store_plans_xmlplan(plan);
SELECT pg_store_plans_yamlplan(plan);
```

`pg_store_plans_*plan()` is useful when `pg_store_plans.plan_format = 'raw'`.

### Key GUCs

- `pg_store_plans.max`
- `pg_store_plans.track`
- `pg_store_plans.max_plan_length`
- `pg_store_plans.plan_storage`
- `pg_store_plans.plan_format`
- `pg_store_plans.min_duration`
- `pg_store_plans.log_analyze`
- `pg_store_plans.log_buffers`
- `pg_store_plans.log_timing`
- `pg_store_plans.save`

The docs describe `plan_storage` as `file` or `shmem`, and `plan_format` as `text`, `json`, `xml`, `yaml`, or `raw`.

### Caveats

- Non-superusers cannot see `plan`, `queryid`, or `planid` for statements executed by other users.
- `pg_store_plans` and `pg_stat_statements` maintain entries independently, so low-frequency rows may not always have a matching peer.
