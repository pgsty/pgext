---
title: "pg_profile"
linkTitle: "pg_profile"
description: "PostgreSQL load profile repository and report builder"
weight: 6000
categories: ["STAT"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_profile**](https://github.com/zubkov-andrei/pg_profile) : PostgreSQL load profile repository and report builder


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6000** | {{< badge content="pg_profile" link="https://github.com/zubkov-andrei/pg_profile" >}} | {{< ext "pg_profile" >}} | `4.16` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "dblink" >}} {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pg_stat_monitor" >}} {{< ext "powa" >}} {{< ext "pg_wait_sampling" >}} {{< ext "pgsentinel" >}} {{< ext "pg_datasentinel" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_stat_plans" >}} {{< ext "pg_track_settings" >}} {{< ext "pg_track_optimizer" >}} |

> [!Note] SQL-only; requires dblink and plpgsql.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.16` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_profile` | `dblink`, `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.16` | {{< bg "18" "pg_profile_18" "green" >}} {{< bg "17" "pg_profile_17" "green" >}} {{< bg "16" "pg_profile_16" "green" >}} {{< bg "15" "pg_profile_15" "green" >}} {{< bg "14" "pg_profile_14" "green" >}} | `pg_profile_$v` | `postgresql$v-contrib` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.16` | {{< bg "18" "postgresql-18-pg-profile" "green" >}} {{< bg "17" "postgresql-17-pg-profile" "green" >}} {{< bg "16" "postgresql-16-pg-profile" "green" >}} {{< bg "15" "postgresql-15-pg-profile" "green" >}} {{< bg "14" "postgresql-14-pg-profile" "green" >}} | `postgresql-$v-pg-profile` | `postgresql-contrib-$v` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 4.16" "pg_profile_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_16 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_15 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_14 : AVAIL 8" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 4.16" "pg_profile_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_16 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_15 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_14 : AVAIL 8" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 4.16" "pg_profile_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_16 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_15 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_14 : AVAIL 9" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 4.16" "pg_profile_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_16 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_15 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_14 : AVAIL 9" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 4.16" "pg_profile_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_15 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_14 : AVAIL 6" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 4.16" "pg_profile_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_15 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 4.16" "pg_profile_14 : AVAIL 6" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.16" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-14-pg-profile : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.16" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-14-pg-profile : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.16" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-14-pg-profile : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.16" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-14-pg-profile : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.16" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-14-pg-profile : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.16" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-14-pg-profile : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.16" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-14-pg-profile : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.16" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-14-pg-profile : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 4.16" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-14-pg-profile : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 4.16" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.16" "postgresql-14-pg-profile : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_profile_18` | `4.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 244.7 KiB | [pg_profile_18-4.16-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_profile_18-4.16-1PGSTY.el8.noarch.rpm) |
| `pg_profile_18` | `4.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 221.0 KiB | [pg_profile_18-4.15-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_profile_18-4.15-1PGDG.rhel8.10.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.4 KiB | [pg_profile_18-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_profile_18-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_18` | `4.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.2 KiB | [pg_profile_18-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_profile_18-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_18` | `4.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 244.7 KiB | [pg_profile_18-4.16-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_profile_18-4.16-1PGSTY.el8.noarch.rpm) |
| `pg_profile_18` | `4.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 220.9 KiB | [pg_profile_18-4.15-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_profile_18-4.15-1PGDG.rhel8.10.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.3 KiB | [pg_profile_18-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_profile_18-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_18` | `4.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.2 KiB | [pg_profile_18-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_profile_18-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_18` | `4.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 223.6 KiB | [pg_profile_18-4.16-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_profile_18-4.16-1PGSTY.el9.noarch.rpm) |
| `pg_profile_18` | `4.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 201.7 KiB | [pg_profile_18-4.15-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_profile_18-4.15-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.8 KiB | [pg_profile_18-4.11-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_profile_18-4.11-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.0 KiB | [pg_profile_18-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_profile_18-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_18` | `4.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 196.9 KiB | [pg_profile_18-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_profile_18-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_18` | `4.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 223.5 KiB | [pg_profile_18-4.16-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_profile_18-4.16-1PGSTY.el9.noarch.rpm) |
| `pg_profile_18` | `4.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 201.6 KiB | [pg_profile_18-4.15-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_profile_18-4.15-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.7 KiB | [pg_profile_18-4.11-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_profile_18-4.11-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.9 KiB | [pg_profile_18-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_profile_18-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_18` | `4.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.9 KiB | [pg_profile_18-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_profile_18-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_18` | `4.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 223.7 KiB | [pg_profile_18-4.16-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_profile_18-4.16-1PGSTY.el10.noarch.rpm) |
| `pg_profile_18` | `4.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 201.9 KiB | [pg_profile_18-4.15-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_profile_18-4.15-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 198.9 KiB | [pg_profile_18-4.11-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_profile_18-4.11-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.5 KiB | [pg_profile_18-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_profile_18-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_18` | `4.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.4 KiB | [pg_profile_18-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_profile_18-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_18` | `4.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 223.6 KiB | [pg_profile_18-4.16-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_profile_18-4.16-1PGSTY.el10.noarch.rpm) |
| `pg_profile_18` | `4.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 201.8 KiB | [pg_profile_18-4.15-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_profile_18-4.15-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 198.9 KiB | [pg_profile_18-4.11-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_profile_18-4.11-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.5 KiB | [pg_profile_18-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_profile_18-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_18` | `4.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.4 KiB | [pg_profile_18-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_profile_18-4.10-1PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pg-profile` | `4.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 198.9 KiB | [postgresql-18-pg-profile_4.16-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-18-pg-profile_4.16-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-profile` | `4.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 198.9 KiB | [postgresql-18-pg-profile_4.16-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-18-pg-profile_4.16-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-profile` | `4.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 198.9 KiB | [postgresql-18-pg-profile_4.16-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-18-pg-profile_4.16-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-profile` | `4.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 198.9 KiB | [postgresql-18-pg-profile_4.16-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-18-pg-profile_4.16-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-profile` | `4.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 198.6 KiB | [postgresql-18-pg-profile_4.16-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-18-pg-profile_4.16-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-profile` | `4.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 198.6 KiB | [postgresql-18-pg-profile_4.16-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-18-pg-profile_4.16-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-profile` | `4.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 197.2 KiB | [postgresql-18-pg-profile_4.16-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-18-pg-profile_4.16-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-profile` | `4.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 197.2 KiB | [postgresql-18-pg-profile_4.16-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-18-pg-profile_4.16-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-profile` | `4.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 197.3 KiB | [postgresql-18-pg-profile_4.16-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-profile/postgresql-18-pg-profile_4.16-1PGSTY~resolute_all.deb) |
| `postgresql-18-pg-profile` | `4.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 197.3 KiB | [postgresql-18-pg-profile_4.16-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-profile/postgresql-18-pg-profile_4.16-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_profile_17` | `4.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 244.7 KiB | [pg_profile_17-4.16-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_profile_17-4.16-1PGSTY.el8.noarch.rpm) |
| `pg_profile_17` | `4.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 221.0 KiB | [pg_profile_17-4.15-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_profile_17-4.15-1PGDG.rhel8.10.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.4 KiB | [pg_profile_17-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_profile_17-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.2 KiB | [pg_profile_17-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_profile_17-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.9 KiB | [pg_profile_17-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_profile_17-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.4 KiB | [pg_profile_17-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_profile_17-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 244.7 KiB | [pg_profile_17-4.16-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_profile_17-4.16-1PGSTY.el8.noarch.rpm) |
| `pg_profile_17` | `4.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 220.9 KiB | [pg_profile_17-4.15-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_profile_17-4.15-1PGDG.rhel8.10.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.3 KiB | [pg_profile_17-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_profile_17-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.2 KiB | [pg_profile_17-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_profile_17-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.9 KiB | [pg_profile_17-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_profile_17-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.3 KiB | [pg_profile_17-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_profile_17-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 223.6 KiB | [pg_profile_17-4.16-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_profile_17-4.16-1PGSTY.el9.noarch.rpm) |
| `pg_profile_17` | `4.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 201.7 KiB | [pg_profile_17-4.15-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_profile_17-4.15-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.8 KiB | [pg_profile_17-4.11-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_profile_17-4.11-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.0 KiB | [pg_profile_17-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_profile_17-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 196.9 KiB | [pg_profile_17-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_profile_17-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 117.0 KiB | [pg_profile_17-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_profile_17-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.8 KiB | [pg_profile_17-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_profile_17-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 223.5 KiB | [pg_profile_17-4.16-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_profile_17-4.16-1PGSTY.el9.noarch.rpm) |
| `pg_profile_17` | `4.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 201.6 KiB | [pg_profile_17-4.15-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_profile_17-4.15-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.8 KiB | [pg_profile_17-4.11-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_profile_17-4.11-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.9 KiB | [pg_profile_17-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_profile_17-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.8 KiB | [pg_profile_17-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_profile_17-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 117.0 KiB | [pg_profile_17-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_profile_17-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 115.7 KiB | [pg_profile_17-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_profile_17-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 223.7 KiB | [pg_profile_17-4.16-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_profile_17-4.16-1PGSTY.el10.noarch.rpm) |
| `pg_profile_17` | `4.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 201.9 KiB | [pg_profile_17-4.15-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_profile_17-4.15-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 198.9 KiB | [pg_profile_17-4.11-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_profile_17-4.11-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.5 KiB | [pg_profile_17-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_profile_17-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_17` | `4.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.4 KiB | [pg_profile_17-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_profile_17-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_17` | `4.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.5 KiB | [pg_profile_17-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_profile_17-4.8-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_17` | `4.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 223.6 KiB | [pg_profile_17-4.16-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_profile_17-4.16-1PGSTY.el10.noarch.rpm) |
| `pg_profile_17` | `4.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 201.8 KiB | [pg_profile_17-4.15-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_profile_17-4.15-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 198.9 KiB | [pg_profile_17-4.11-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_profile_17-4.11-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.5 KiB | [pg_profile_17-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_profile_17-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_17` | `4.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.4 KiB | [pg_profile_17-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_profile_17-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_17` | `4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [pg_profile_17-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_profile_17-4.8-1PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pg-profile` | `4.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 198.9 KiB | [postgresql-17-pg-profile_4.16-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-17-pg-profile_4.16-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-profile` | `4.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 198.9 KiB | [postgresql-17-pg-profile_4.16-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-17-pg-profile_4.16-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-profile` | `4.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 198.9 KiB | [postgresql-17-pg-profile_4.16-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-17-pg-profile_4.16-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-profile` | `4.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 198.9 KiB | [postgresql-17-pg-profile_4.16-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-17-pg-profile_4.16-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-profile` | `4.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 198.6 KiB | [postgresql-17-pg-profile_4.16-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-17-pg-profile_4.16-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-profile` | `4.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 198.6 KiB | [postgresql-17-pg-profile_4.16-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-17-pg-profile_4.16-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-profile` | `4.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 197.2 KiB | [postgresql-17-pg-profile_4.16-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-17-pg-profile_4.16-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-profile` | `4.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 197.2 KiB | [postgresql-17-pg-profile_4.16-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-17-pg-profile_4.16-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-profile` | `4.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 197.4 KiB | [postgresql-17-pg-profile_4.16-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-profile/postgresql-17-pg-profile_4.16-1PGSTY~resolute_all.deb) |
| `postgresql-17-pg-profile` | `4.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 197.4 KiB | [postgresql-17-pg-profile_4.16-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-profile/postgresql-17-pg-profile_4.16-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_profile_16` | `4.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 244.7 KiB | [pg_profile_16-4.16-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_profile_16-4.16-1PGSTY.el8.noarch.rpm) |
| `pg_profile_16` | `4.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 221.0 KiB | [pg_profile_16-4.15-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.15-1PGDG.rhel8.10.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.4 KiB | [pg_profile_16-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.2 KiB | [pg_profile_16-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.9 KiB | [pg_profile_16-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.4 KiB | [pg_profile_16-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 119.8 KiB | [pg_profile_16-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.7 KiB | [pg_profile_16-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 244.7 KiB | [pg_profile_16-4.16-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_profile_16-4.16-1PGSTY.el8.noarch.rpm) |
| `pg_profile_16` | `4.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 220.9 KiB | [pg_profile_16-4.15-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.15-1PGDG.rhel8.10.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.3 KiB | [pg_profile_16-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.2 KiB | [pg_profile_16-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.9 KiB | [pg_profile_16-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.3 KiB | [pg_profile_16-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 119.8 KiB | [pg_profile_16-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 109.7 KiB | [pg_profile_16-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 223.6 KiB | [pg_profile_16-4.16-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_profile_16-4.16-1PGSTY.el9.noarch.rpm) |
| `pg_profile_16` | `4.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 201.7 KiB | [pg_profile_16-4.15-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.15-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.8 KiB | [pg_profile_16-4.11-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.11-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.0 KiB | [pg_profile_16-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 196.9 KiB | [pg_profile_16-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 117.0 KiB | [pg_profile_16-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.8 KiB | [pg_profile_16-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.8 KiB | [pg_profile_16-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 99.1 KiB | [pg_profile_16-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 223.5 KiB | [pg_profile_16-4.16-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_profile_16-4.16-1PGSTY.el9.noarch.rpm) |
| `pg_profile_16` | `4.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 201.7 KiB | [pg_profile_16-4.15-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.15-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.8 KiB | [pg_profile_16-4.11-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.11-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.9 KiB | [pg_profile_16-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.8 KiB | [pg_profile_16-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 117.0 KiB | [pg_profile_16-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 115.7 KiB | [pg_profile_16-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.7 KiB | [pg_profile_16-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.0 KiB | [pg_profile_16-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 223.7 KiB | [pg_profile_16-4.16-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_profile_16-4.16-1PGSTY.el10.noarch.rpm) |
| `pg_profile_16` | `4.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 201.9 KiB | [pg_profile_16-4.15-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_profile_16-4.15-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 198.9 KiB | [pg_profile_16-4.11-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_profile_16-4.11-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.5 KiB | [pg_profile_16-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_profile_16-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_16` | `4.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.4 KiB | [pg_profile_16-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_profile_16-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_16` | `4.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.5 KiB | [pg_profile_16-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_profile_16-4.8-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_16` | `4.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 223.6 KiB | [pg_profile_16-4.16-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_profile_16-4.16-1PGSTY.el10.noarch.rpm) |
| `pg_profile_16` | `4.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 201.8 KiB | [pg_profile_16-4.15-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_profile_16-4.15-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 198.9 KiB | [pg_profile_16-4.11-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_profile_16-4.11-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.5 KiB | [pg_profile_16-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_profile_16-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_16` | `4.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.4 KiB | [pg_profile_16-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_profile_16-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_16` | `4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [pg_profile_16-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_profile_16-4.8-1PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pg-profile` | `4.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 198.9 KiB | [postgresql-16-pg-profile_4.16-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-16-pg-profile_4.16-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-profile` | `4.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 198.9 KiB | [postgresql-16-pg-profile_4.16-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-16-pg-profile_4.16-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-profile` | `4.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 198.9 KiB | [postgresql-16-pg-profile_4.16-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-16-pg-profile_4.16-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-profile` | `4.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 198.9 KiB | [postgresql-16-pg-profile_4.16-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-16-pg-profile_4.16-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-profile` | `4.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 198.7 KiB | [postgresql-16-pg-profile_4.16-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-16-pg-profile_4.16-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-profile` | `4.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 198.7 KiB | [postgresql-16-pg-profile_4.16-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-16-pg-profile_4.16-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-profile` | `4.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 197.3 KiB | [postgresql-16-pg-profile_4.16-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-16-pg-profile_4.16-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-profile` | `4.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 197.3 KiB | [postgresql-16-pg-profile_4.16-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-16-pg-profile_4.16-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-profile` | `4.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 197.4 KiB | [postgresql-16-pg-profile_4.16-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-profile/postgresql-16-pg-profile_4.16-1PGSTY~resolute_all.deb) |
| `postgresql-16-pg-profile` | `4.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 197.4 KiB | [postgresql-16-pg-profile_4.16-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-profile/postgresql-16-pg-profile_4.16-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_profile_15` | `4.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 244.7 KiB | [pg_profile_15-4.16-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_profile_15-4.16-1PGSTY.el8.noarch.rpm) |
| `pg_profile_15` | `4.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 221.0 KiB | [pg_profile_15-4.15-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.15-1PGDG.rhel8.10.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.4 KiB | [pg_profile_15-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.2 KiB | [pg_profile_15-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.9 KiB | [pg_profile_15-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.4 KiB | [pg_profile_15-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 119.8 KiB | [pg_profile_15-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.7 KiB | [pg_profile_15-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 244.7 KiB | [pg_profile_15-4.16-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_profile_15-4.16-1PGSTY.el8.noarch.rpm) |
| `pg_profile_15` | `4.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 220.9 KiB | [pg_profile_15-4.15-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.15-1PGDG.rhel8.10.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.3 KiB | [pg_profile_15-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.2 KiB | [pg_profile_15-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.9 KiB | [pg_profile_15-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.3 KiB | [pg_profile_15-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 119.8 KiB | [pg_profile_15-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 109.7 KiB | [pg_profile_15-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 223.6 KiB | [pg_profile_15-4.16-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_profile_15-4.16-1PGSTY.el9.noarch.rpm) |
| `pg_profile_15` | `4.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 201.7 KiB | [pg_profile_15-4.15-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.15-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.8 KiB | [pg_profile_15-4.11-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.11-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.0 KiB | [pg_profile_15-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 196.9 KiB | [pg_profile_15-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 117.0 KiB | [pg_profile_15-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.8 KiB | [pg_profile_15-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.8 KiB | [pg_profile_15-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 99.1 KiB | [pg_profile_15-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 223.5 KiB | [pg_profile_15-4.16-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_profile_15-4.16-1PGSTY.el9.noarch.rpm) |
| `pg_profile_15` | `4.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 201.6 KiB | [pg_profile_15-4.15-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.15-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.7 KiB | [pg_profile_15-4.11-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.11-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.9 KiB | [pg_profile_15-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.8 KiB | [pg_profile_15-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 117.0 KiB | [pg_profile_15-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 115.7 KiB | [pg_profile_15-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.7 KiB | [pg_profile_15-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.0 KiB | [pg_profile_15-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 223.7 KiB | [pg_profile_15-4.16-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_profile_15-4.16-1PGSTY.el10.noarch.rpm) |
| `pg_profile_15` | `4.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 201.9 KiB | [pg_profile_15-4.15-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_profile_15-4.15-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 198.9 KiB | [pg_profile_15-4.11-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_profile_15-4.11-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.5 KiB | [pg_profile_15-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_profile_15-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_15` | `4.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.4 KiB | [pg_profile_15-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_profile_15-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_15` | `4.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.5 KiB | [pg_profile_15-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_profile_15-4.8-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_15` | `4.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 223.6 KiB | [pg_profile_15-4.16-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_profile_15-4.16-1PGSTY.el10.noarch.rpm) |
| `pg_profile_15` | `4.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 201.8 KiB | [pg_profile_15-4.15-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_profile_15-4.15-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 198.9 KiB | [pg_profile_15-4.11-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_profile_15-4.11-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.5 KiB | [pg_profile_15-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_profile_15-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_15` | `4.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.4 KiB | [pg_profile_15-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_profile_15-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_15` | `4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [pg_profile_15-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_profile_15-4.8-1PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pg-profile` | `4.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 198.9 KiB | [postgresql-15-pg-profile_4.16-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-15-pg-profile_4.16-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pg-profile` | `4.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 198.9 KiB | [postgresql-15-pg-profile_4.16-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-15-pg-profile_4.16-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pg-profile` | `4.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 198.8 KiB | [postgresql-15-pg-profile_4.16-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-15-pg-profile_4.16-1PGSTY~trixie_all.deb) |
| `postgresql-15-pg-profile` | `4.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 198.8 KiB | [postgresql-15-pg-profile_4.16-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-15-pg-profile_4.16-1PGSTY~trixie_all.deb) |
| `postgresql-15-pg-profile` | `4.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 198.7 KiB | [postgresql-15-pg-profile_4.16-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-15-pg-profile_4.16-1PGSTY~jammy_all.deb) |
| `postgresql-15-pg-profile` | `4.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 198.7 KiB | [postgresql-15-pg-profile_4.16-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-15-pg-profile_4.16-1PGSTY~jammy_all.deb) |
| `postgresql-15-pg-profile` | `4.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 197.2 KiB | [postgresql-15-pg-profile_4.16-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-15-pg-profile_4.16-1PGSTY~noble_all.deb) |
| `postgresql-15-pg-profile` | `4.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 197.2 KiB | [postgresql-15-pg-profile_4.16-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-15-pg-profile_4.16-1PGSTY~noble_all.deb) |
| `postgresql-15-pg-profile` | `4.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 197.4 KiB | [postgresql-15-pg-profile_4.16-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-profile/postgresql-15-pg-profile_4.16-1PGSTY~resolute_all.deb) |
| `postgresql-15-pg-profile` | `4.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 197.4 KiB | [postgresql-15-pg-profile_4.16-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-profile/postgresql-15-pg-profile_4.16-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_profile_14` | `4.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 244.7 KiB | [pg_profile_14-4.16-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_profile_14-4.16-1PGSTY.el8.noarch.rpm) |
| `pg_profile_14` | `4.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 221.0 KiB | [pg_profile_14-4.15-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.15-1PGDG.rhel8.10.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.4 KiB | [pg_profile_14-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.2 KiB | [pg_profile_14-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.9 KiB | [pg_profile_14-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.4 KiB | [pg_profile_14-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 119.8 KiB | [pg_profile_14-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.7 KiB | [pg_profile_14-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 244.7 KiB | [pg_profile_14-4.16-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_profile_14-4.16-1PGSTY.el8.noarch.rpm) |
| `pg_profile_14` | `4.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 220.9 KiB | [pg_profile_14-4.15-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.15-1PGDG.rhel8.10.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.3 KiB | [pg_profile_14-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.2 KiB | [pg_profile_14-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.9 KiB | [pg_profile_14-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.3 KiB | [pg_profile_14-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 119.8 KiB | [pg_profile_14-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 109.7 KiB | [pg_profile_14-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 223.6 KiB | [pg_profile_14-4.16-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_profile_14-4.16-1PGSTY.el9.noarch.rpm) |
| `pg_profile_14` | `4.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 201.7 KiB | [pg_profile_14-4.15-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.15-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 198.8 KiB | [pg_profile_14-4.11-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.11-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.0 KiB | [pg_profile_14-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 196.9 KiB | [pg_profile_14-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 117.1 KiB | [pg_profile_14-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.8 KiB | [pg_profile_14-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.8 KiB | [pg_profile_14-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 99.1 KiB | [pg_profile_14-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 223.5 KiB | [pg_profile_14-4.16-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_profile_14-4.16-1PGSTY.el9.noarch.rpm) |
| `pg_profile_14` | `4.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 201.6 KiB | [pg_profile_14-4.15-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.15-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 198.8 KiB | [pg_profile_14-4.11-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.11-1PGDG.rhel9.8.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 197.0 KiB | [pg_profile_14-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.8 KiB | [pg_profile_14-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 117.0 KiB | [pg_profile_14-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 115.7 KiB | [pg_profile_14-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.7 KiB | [pg_profile_14-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.0 KiB | [pg_profile_14-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 223.7 KiB | [pg_profile_14-4.16-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_profile_14-4.16-1PGSTY.el10.noarch.rpm) |
| `pg_profile_14` | `4.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 201.9 KiB | [pg_profile_14-4.15-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_profile_14-4.15-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 198.9 KiB | [pg_profile_14-4.11-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_profile_14-4.11-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.5 KiB | [pg_profile_14-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_profile_14-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_14` | `4.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.4 KiB | [pg_profile_14-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_profile_14-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_14` | `4.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.5 KiB | [pg_profile_14-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_profile_14-4.8-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_14` | `4.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 223.6 KiB | [pg_profile_14-4.16-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_profile_14-4.16-1PGSTY.el10.noarch.rpm) |
| `pg_profile_14` | `4.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 201.8 KiB | [pg_profile_14-4.15-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_profile_14-4.15-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 198.9 KiB | [pg_profile_14-4.11-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_profile_14-4.11-1PGDG.rhel10.2.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.5 KiB | [pg_profile_14-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_profile_14-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_14` | `4.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.4 KiB | [pg_profile_14-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_profile_14-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_14` | `4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [pg_profile_14-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_profile_14-4.8-1PGDG.rhel10.noarch.rpm) |
| `postgresql-14-pg-profile` | `4.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 198.9 KiB | [postgresql-14-pg-profile_4.16-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-14-pg-profile_4.16-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pg-profile` | `4.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 198.9 KiB | [postgresql-14-pg-profile_4.16-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-14-pg-profile_4.16-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pg-profile` | `4.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 198.9 KiB | [postgresql-14-pg-profile_4.16-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-14-pg-profile_4.16-1PGSTY~trixie_all.deb) |
| `postgresql-14-pg-profile` | `4.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 198.9 KiB | [postgresql-14-pg-profile_4.16-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-14-pg-profile_4.16-1PGSTY~trixie_all.deb) |
| `postgresql-14-pg-profile` | `4.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 198.7 KiB | [postgresql-14-pg-profile_4.16-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-14-pg-profile_4.16-1PGSTY~jammy_all.deb) |
| `postgresql-14-pg-profile` | `4.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 198.7 KiB | [postgresql-14-pg-profile_4.16-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-14-pg-profile_4.16-1PGSTY~jammy_all.deb) |
| `postgresql-14-pg-profile` | `4.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 197.3 KiB | [postgresql-14-pg-profile_4.16-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-14-pg-profile_4.16-1PGSTY~noble_all.deb) |
| `postgresql-14-pg-profile` | `4.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 197.3 KiB | [postgresql-14-pg-profile_4.16-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-14-pg-profile_4.16-1PGSTY~noble_all.deb) |
| `postgresql-14-pg-profile` | `4.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 197.3 KiB | [postgresql-14-pg-profile_4.16-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-profile/postgresql-14-pg-profile_4.16-1PGSTY~resolute_all.deb) |
| `postgresql-14-pg-profile` | `4.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 197.3 KiB | [postgresql-14-pg-profile_4.16-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-profile/postgresql-14-pg-profile_4.16-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/zubkov-andrei/pg_profile" title="Repository" icon="github" subtitle="github.com/zubkov-andrei/pg_profile" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_profile-4.16.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_profile;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_profile;		# install via package name, for the active PG version

pig install pg_profile -v 18;   # install for PG 18
pig install pg_profile -v 17;   # install for PG 17
pig install pg_profile -v 16;   # install for PG 16
pig install pg_profile -v 15;   # install for PG 15
pig install pg_profile -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_profile CASCADE; -- requires dblink, plpgsql
```




## Usage

> [pg_profile: historical performance profiling tool for PostgreSQL](https://github.com/zubkov-andrei/pg_profile)

pg_profile collects periodic samples of PostgreSQL statistics and generates detailed historical performance reports. It depends on `pg_stat_statements` and optionally uses `pg_stat_kcache` and `pg_wait_sampling` for additional metrics.

### Taking Samples

Samples must be taken periodically (e.g., via cron). Each sample captures the current state of statistics:

```sql
SELECT profile.take_sample();
```

### Generating Reports

Build a report between two sample IDs to analyze performance during that interval:

```sql
-- Regular report between samples 1 and 2
SELECT profile.get_report(1, 2);

-- Differential report comparing two intervals
SELECT profile.get_diffreport(1, 2, 3, 4);
```

### Managing Servers

pg_profile can collect statistics from remote clusters:

```sql
-- Define a remote server
SELECT profile.create_server('remote', 'host=remote_host dbname=postgres');

-- List defined servers
SELECT * FROM profile.show_servers();

-- Enable/disable a server
SELECT profile.enable_server('remote');
SELECT profile.disable_server('remote');
```

### Baselines

Baselines protect sample ranges from automatic cleanup:

```sql
-- Create a baseline preserving samples 10 through 20
SELECT profile.create_baseline('incident_2024', 10, 20);

-- List baselines
SELECT * FROM profile.show_baselines();

-- Drop a baseline
SELECT profile.drop_baseline('incident_2024');
```

### Retention

Control how long samples are kept:

```sql
-- Set retention to 7 days for the local server
SELECT profile.set_server_max_sample_age('local', 7);
```

### Sample Information

```sql
-- Show available samples
SELECT * FROM profile.show_samples();

-- Show time spent taking samples (requires pg_profile.track_sample_timings = on)
SELECT * FROM v_sample_timings;
```

### Recommended Settings

```
track_activities = on
track_counts = on
track_io_timing = on
track_wal_io_timing = on      # PG 14+
track_functions = all
```
