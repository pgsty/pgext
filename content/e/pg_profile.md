---
title: "pg_profile"
linkTitle: "pg_profile"
description: "PostgreSQL load profile repository and report builder"
weight: 6000
categories: ["STAT"]
width: full
---

[**pg_profile**](https://github.com/zubkov-andrei/pg_profile) : PostgreSQL load profile repository and report builder


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6000** | {{< badge content="pg_profile" link="https://github.com/zubkov-andrei/pg_profile" >}} | {{< ext "pg_profile" >}} | `4.11` | {{< category "STAT" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "dblink" >}} {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "plprofiler" >}} {{< ext "pg_stat_kcache" >}} {{< ext "powa" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `4.11` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_profile` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.11` | {{< bg "18" "pg_profile_18" "green" >}} {{< bg "17" "pg_profile_17" "green" >}} {{< bg "16" "pg_profile_16" "green" >}} {{< bg "15" "pg_profile_15" "green" >}} {{< bg "14" "pg_profile_14" "green" >}} {{< bg "13" "pg_profile_13" "green" >}} | `pg_profile_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.11` | {{< bg "18" "postgresql-18-pg-profile" "green" >}} {{< bg "17" "postgresql-17-pg-profile" "green" >}} {{< bg "16" "postgresql-16-pg-profile" "green" >}} {{< bg "15" "postgresql-15-pg-profile" "green" >}} {{< bg "14" "postgresql-14-pg-profile" "green" >}} {{< bg "13" "postgresql-13-pg-profile" "green" >}} | `postgresql-$v-pg-profile` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.11" "pg_profile_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_13 : AVAIL 6" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.11" "pg_profile_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_13 : AVAIL 6" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.11" "pg_profile_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_13 : AVAIL 6" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.11" "pg_profile_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_13 : AVAIL 6" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.11" "pg_profile_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.11" "pg_profile_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.11" "pg_profile_13 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.11" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-14-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-13-pg-profile : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.11" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-14-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-13-pg-profile : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.11" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-14-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-13-pg-profile : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.11" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-14-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-13-pg-profile : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.11" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-14-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-13-pg-profile : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.11" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-14-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-13-pg-profile : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.11" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-14-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-13-pg-profile : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.11" "postgresql-18-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-17-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-16-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-15-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-14-pg-profile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.11" "postgresql-13-pg-profile : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_profile_18` | `4.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.4 KiB | [pg_profile_18-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_profile_18-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_18` | `4.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.2 KiB | [pg_profile_18-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_profile_18-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.3 KiB | [pg_profile_18-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_profile_18-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_18` | `4.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.2 KiB | [pg_profile_18-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_profile_18-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.0 KiB | [pg_profile_18-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_profile_18-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_18` | `4.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 196.9 KiB | [pg_profile_18-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_profile_18-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.9 KiB | [pg_profile_18-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_profile_18-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_18` | `4.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.9 KiB | [pg_profile_18-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_profile_18-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.5 KiB | [pg_profile_18-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_profile_18-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_18` | `4.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.4 KiB | [pg_profile_18-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_profile_18-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_18` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.5 KiB | [pg_profile_18-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_profile_18-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_18` | `4.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.4 KiB | [pg_profile_18-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_profile_18-4.10-1PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pg-profile` | `4.11` | [d12.x86_64](/os/d12.x86_64) | pigsty | 192.3 KiB | [postgresql-18-pg-profile_4.11-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-18-pg-profile_4.11-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-profile` | `4.11` | [d12.aarch64](/os/d12.aarch64) | pigsty | 192.3 KiB | [postgresql-18-pg-profile_4.11-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-18-pg-profile_4.11-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-profile` | `4.11` | [d13.x86_64](/os/d13.x86_64) | pigsty | 192.3 KiB | [postgresql-18-pg-profile_4.11-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-18-pg-profile_4.11-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-profile` | `4.11` | [d13.aarch64](/os/d13.aarch64) | pigsty | 192.3 KiB | [postgresql-18-pg-profile_4.11-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-18-pg-profile_4.11-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-profile` | `4.11` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.2 KiB | [postgresql-18-pg-profile_4.11-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-18-pg-profile_4.11-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-profile` | `4.11` | [u22.aarch64](/os/u22.aarch64) | pigsty | 193.2 KiB | [postgresql-18-pg-profile_4.11-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-18-pg-profile_4.11-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-profile` | `4.11` | [u24.x86_64](/os/u24.x86_64) | pigsty | 191.6 KiB | [postgresql-18-pg-profile_4.11-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-18-pg-profile_4.11-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-profile` | `4.11` | [u24.aarch64](/os/u24.aarch64) | pigsty | 191.6 KiB | [postgresql-18-pg-profile_4.11-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-18-pg-profile_4.11-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_profile_17` | `4.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.4 KiB | [pg_profile_17-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_profile_17-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.2 KiB | [pg_profile_17-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_profile_17-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.9 KiB | [pg_profile_17-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_profile_17-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.4 KiB | [pg_profile_17-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_profile_17-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.3 KiB | [pg_profile_17-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_profile_17-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.2 KiB | [pg_profile_17-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_profile_17-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.9 KiB | [pg_profile_17-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_profile_17-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.3 KiB | [pg_profile_17-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_profile_17-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.0 KiB | [pg_profile_17-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_profile_17-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 196.9 KiB | [pg_profile_17-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_profile_17-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 117.0 KiB | [pg_profile_17-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_profile_17-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.8 KiB | [pg_profile_17-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_profile_17-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.9 KiB | [pg_profile_17-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_profile_17-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.8 KiB | [pg_profile_17-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_profile_17-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 117.0 KiB | [pg_profile_17-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_profile_17-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 115.7 KiB | [pg_profile_17-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_profile_17-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.5 KiB | [pg_profile_17-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_profile_17-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_17` | `4.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.4 KiB | [pg_profile_17-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_profile_17-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_17` | `4.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.5 KiB | [pg_profile_17-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_profile_17-4.8-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_17` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.5 KiB | [pg_profile_17-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_profile_17-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_17` | `4.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.4 KiB | [pg_profile_17-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_profile_17-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_17` | `4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [pg_profile_17-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_profile_17-4.8-1PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pg-profile` | `4.11` | [d12.x86_64](/os/d12.x86_64) | pigsty | 192.3 KiB | [postgresql-17-pg-profile_4.11-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-17-pg-profile_4.11-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-profile` | `4.11` | [d12.aarch64](/os/d12.aarch64) | pigsty | 192.3 KiB | [postgresql-17-pg-profile_4.11-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-17-pg-profile_4.11-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-profile` | `4.11` | [d13.x86_64](/os/d13.x86_64) | pigsty | 192.3 KiB | [postgresql-17-pg-profile_4.11-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-17-pg-profile_4.11-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-profile` | `4.11` | [d13.aarch64](/os/d13.aarch64) | pigsty | 192.3 KiB | [postgresql-17-pg-profile_4.11-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-17-pg-profile_4.11-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-profile` | `4.11` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.2 KiB | [postgresql-17-pg-profile_4.11-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-17-pg-profile_4.11-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-profile` | `4.11` | [u22.aarch64](/os/u22.aarch64) | pigsty | 193.2 KiB | [postgresql-17-pg-profile_4.11-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-17-pg-profile_4.11-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-profile` | `4.11` | [u24.x86_64](/os/u24.x86_64) | pigsty | 191.6 KiB | [postgresql-17-pg-profile_4.11-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-17-pg-profile_4.11-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-profile` | `4.11` | [u24.aarch64](/os/u24.aarch64) | pigsty | 191.6 KiB | [postgresql-17-pg-profile_4.11-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-17-pg-profile_4.11-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_profile_16` | `4.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.4 KiB | [pg_profile_16-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.2 KiB | [pg_profile_16-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.9 KiB | [pg_profile_16-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.4 KiB | [pg_profile_16-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 119.8 KiB | [pg_profile_16-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.7 KiB | [pg_profile_16-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_profile_16-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.3 KiB | [pg_profile_16-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.2 KiB | [pg_profile_16-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.9 KiB | [pg_profile_16-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.3 KiB | [pg_profile_16-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 119.8 KiB | [pg_profile_16-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 109.7 KiB | [pg_profile_16-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_profile_16-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.0 KiB | [pg_profile_16-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 196.9 KiB | [pg_profile_16-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 117.0 KiB | [pg_profile_16-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.8 KiB | [pg_profile_16-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.8 KiB | [pg_profile_16-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 99.1 KiB | [pg_profile_16-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_profile_16-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.9 KiB | [pg_profile_16-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.8 KiB | [pg_profile_16-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 117.0 KiB | [pg_profile_16-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 115.7 KiB | [pg_profile_16-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.7 KiB | [pg_profile_16-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.0 KiB | [pg_profile_16-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_profile_16-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.5 KiB | [pg_profile_16-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_profile_16-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_16` | `4.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.4 KiB | [pg_profile_16-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_profile_16-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_16` | `4.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.5 KiB | [pg_profile_16-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_profile_16-4.8-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_16` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.5 KiB | [pg_profile_16-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_profile_16-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_16` | `4.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.4 KiB | [pg_profile_16-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_profile_16-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_16` | `4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [pg_profile_16-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_profile_16-4.8-1PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pg-profile` | `4.11` | [d12.x86_64](/os/d12.x86_64) | pigsty | 192.3 KiB | [postgresql-16-pg-profile_4.11-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-16-pg-profile_4.11-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-profile` | `4.11` | [d12.aarch64](/os/d12.aarch64) | pigsty | 192.3 KiB | [postgresql-16-pg-profile_4.11-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-16-pg-profile_4.11-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-profile` | `4.11` | [d13.x86_64](/os/d13.x86_64) | pigsty | 192.3 KiB | [postgresql-16-pg-profile_4.11-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-16-pg-profile_4.11-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-profile` | `4.11` | [d13.aarch64](/os/d13.aarch64) | pigsty | 192.3 KiB | [postgresql-16-pg-profile_4.11-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-16-pg-profile_4.11-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-profile` | `4.11` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.3 KiB | [postgresql-16-pg-profile_4.11-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-16-pg-profile_4.11-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-profile` | `4.11` | [u22.aarch64](/os/u22.aarch64) | pigsty | 193.3 KiB | [postgresql-16-pg-profile_4.11-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-16-pg-profile_4.11-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-profile` | `4.11` | [u24.x86_64](/os/u24.x86_64) | pigsty | 191.6 KiB | [postgresql-16-pg-profile_4.11-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-16-pg-profile_4.11-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-profile` | `4.11` | [u24.aarch64](/os/u24.aarch64) | pigsty | 191.6 KiB | [postgresql-16-pg-profile_4.11-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-16-pg-profile_4.11-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_profile_15` | `4.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.4 KiB | [pg_profile_15-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.2 KiB | [pg_profile_15-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.9 KiB | [pg_profile_15-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.4 KiB | [pg_profile_15-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 119.8 KiB | [pg_profile_15-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.7 KiB | [pg_profile_15-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_profile_15-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.3 KiB | [pg_profile_15-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.2 KiB | [pg_profile_15-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.9 KiB | [pg_profile_15-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.3 KiB | [pg_profile_15-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 119.8 KiB | [pg_profile_15-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 109.7 KiB | [pg_profile_15-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_profile_15-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.0 KiB | [pg_profile_15-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 196.9 KiB | [pg_profile_15-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 117.0 KiB | [pg_profile_15-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.8 KiB | [pg_profile_15-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.8 KiB | [pg_profile_15-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 99.1 KiB | [pg_profile_15-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_profile_15-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.9 KiB | [pg_profile_15-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.8 KiB | [pg_profile_15-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 117.0 KiB | [pg_profile_15-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 115.7 KiB | [pg_profile_15-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.7 KiB | [pg_profile_15-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.0 KiB | [pg_profile_15-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_profile_15-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.5 KiB | [pg_profile_15-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_profile_15-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_15` | `4.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.4 KiB | [pg_profile_15-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_profile_15-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_15` | `4.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.5 KiB | [pg_profile_15-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_profile_15-4.8-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_15` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.5 KiB | [pg_profile_15-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_profile_15-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_15` | `4.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.4 KiB | [pg_profile_15-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_profile_15-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_15` | `4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [pg_profile_15-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_profile_15-4.8-1PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pg-profile` | `4.11` | [d12.x86_64](/os/d12.x86_64) | pigsty | 192.3 KiB | [postgresql-15-pg-profile_4.11-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-15-pg-profile_4.11-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-profile` | `4.11` | [d12.aarch64](/os/d12.aarch64) | pigsty | 192.3 KiB | [postgresql-15-pg-profile_4.11-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-15-pg-profile_4.11-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-profile` | `4.11` | [d13.x86_64](/os/d13.x86_64) | pigsty | 192.3 KiB | [postgresql-15-pg-profile_4.11-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-15-pg-profile_4.11-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-profile` | `4.11` | [d13.aarch64](/os/d13.aarch64) | pigsty | 192.3 KiB | [postgresql-15-pg-profile_4.11-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-15-pg-profile_4.11-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-profile` | `4.11` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.2 KiB | [postgresql-15-pg-profile_4.11-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-15-pg-profile_4.11-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-profile` | `4.11` | [u22.aarch64](/os/u22.aarch64) | pigsty | 193.2 KiB | [postgresql-15-pg-profile_4.11-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-15-pg-profile_4.11-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-profile` | `4.11` | [u24.x86_64](/os/u24.x86_64) | pigsty | 191.6 KiB | [postgresql-15-pg-profile_4.11-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-15-pg-profile_4.11-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-profile` | `4.11` | [u24.aarch64](/os/u24.aarch64) | pigsty | 191.6 KiB | [postgresql-15-pg-profile_4.11-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-15-pg-profile_4.11-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_profile_14` | `4.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.4 KiB | [pg_profile_14-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.2 KiB | [pg_profile_14-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.9 KiB | [pg_profile_14-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.4 KiB | [pg_profile_14-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 119.8 KiB | [pg_profile_14-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.7 KiB | [pg_profile_14-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_profile_14-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.3 KiB | [pg_profile_14-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.2 KiB | [pg_profile_14-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.9 KiB | [pg_profile_14-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.3 KiB | [pg_profile_14-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 119.8 KiB | [pg_profile_14-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 109.7 KiB | [pg_profile_14-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_profile_14-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.0 KiB | [pg_profile_14-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 196.9 KiB | [pg_profile_14-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 117.1 KiB | [pg_profile_14-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.8 KiB | [pg_profile_14-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.8 KiB | [pg_profile_14-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 99.1 KiB | [pg_profile_14-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_profile_14-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 197.0 KiB | [pg_profile_14-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.8 KiB | [pg_profile_14-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 117.0 KiB | [pg_profile_14-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 115.7 KiB | [pg_profile_14-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.7 KiB | [pg_profile_14-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.0 KiB | [pg_profile_14-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_profile_14-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.5 KiB | [pg_profile_14-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_profile_14-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_14` | `4.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.4 KiB | [pg_profile_14-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_profile_14-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_14` | `4.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.5 KiB | [pg_profile_14-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_profile_14-4.8-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_14` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.5 KiB | [pg_profile_14-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_profile_14-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_14` | `4.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.4 KiB | [pg_profile_14-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_profile_14-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_14` | `4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [pg_profile_14-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_profile_14-4.8-1PGDG.rhel10.noarch.rpm) |
| `postgresql-14-pg-profile` | `4.11` | [d12.x86_64](/os/d12.x86_64) | pigsty | 192.3 KiB | [postgresql-14-pg-profile_4.11-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-14-pg-profile_4.11-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-profile` | `4.11` | [d12.aarch64](/os/d12.aarch64) | pigsty | 192.3 KiB | [postgresql-14-pg-profile_4.11-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-14-pg-profile_4.11-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-profile` | `4.11` | [d13.x86_64](/os/d13.x86_64) | pigsty | 192.3 KiB | [postgresql-14-pg-profile_4.11-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-14-pg-profile_4.11-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-profile` | `4.11` | [d13.aarch64](/os/d13.aarch64) | pigsty | 192.3 KiB | [postgresql-14-pg-profile_4.11-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-14-pg-profile_4.11-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-profile` | `4.11` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.2 KiB | [postgresql-14-pg-profile_4.11-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-14-pg-profile_4.11-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-profile` | `4.11` | [u22.aarch64](/os/u22.aarch64) | pigsty | 193.2 KiB | [postgresql-14-pg-profile_4.11-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-14-pg-profile_4.11-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-profile` | `4.11` | [u24.x86_64](/os/u24.x86_64) | pigsty | 191.6 KiB | [postgresql-14-pg-profile_4.11-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-14-pg-profile_4.11-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-profile` | `4.11` | [u24.aarch64](/os/u24.aarch64) | pigsty | 191.6 KiB | [postgresql-14-pg-profile_4.11-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-14-pg-profile_4.11-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_profile_13` | `4.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.4 KiB | [pg_profile_13-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_profile_13-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_13` | `4.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 214.2 KiB | [pg_profile_13-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_profile_13-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_13` | `4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.9 KiB | [pg_profile_13-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_profile_13-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_13` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.4 KiB | [pg_profile_13-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_profile_13-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_13` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 119.8 KiB | [pg_profile_13-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_profile_13-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_13` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.7 KiB | [pg_profile_13-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_profile_13-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_13` | `4.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.3 KiB | [pg_profile_13-4.11-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_profile_13-4.11-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_13` | `4.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 214.2 KiB | [pg_profile_13-4.10-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_profile_13-4.10-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_13` | `4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.9 KiB | [pg_profile_13-4.8-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_profile_13-4.8-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_13` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.3 KiB | [pg_profile_13-4.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_profile_13-4.7-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_13` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 119.8 KiB | [pg_profile_13-4.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_profile_13-4.6-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_13` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 109.7 KiB | [pg_profile_13-4.4-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_profile_13-4.4-1PGDG.rhel8.noarch.rpm) |
| `pg_profile_13` | `4.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 197.0 KiB | [pg_profile_13-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_profile_13-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_13` | `4.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 196.9 KiB | [pg_profile_13-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_profile_13-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_13` | `4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 117.1 KiB | [pg_profile_13-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_profile_13-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_13` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.8 KiB | [pg_profile_13-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_profile_13-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_13` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.7 KiB | [pg_profile_13-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_profile_13-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_13` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 99.1 KiB | [pg_profile_13-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_profile_13-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_13` | `4.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 197.0 KiB | [pg_profile_13-4.11-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_profile_13-4.11-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_13` | `4.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.8 KiB | [pg_profile_13-4.10-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_profile_13-4.10-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_13` | `4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 117.0 KiB | [pg_profile_13-4.8-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_profile_13-4.8-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_13` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 115.7 KiB | [pg_profile_13-4.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_profile_13-4.7-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_13` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.7 KiB | [pg_profile_13-4.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_profile_13-4.6-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_13` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.0 KiB | [pg_profile_13-4.4-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_profile_13-4.4-1PGDG.rhel9.noarch.rpm) |
| `pg_profile_13` | `4.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.5 KiB | [pg_profile_13-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_profile_13-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_13` | `4.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 197.4 KiB | [pg_profile_13-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_profile_13-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_13` | `4.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.5 KiB | [pg_profile_13-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_profile_13-4.8-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_13` | `4.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.5 KiB | [pg_profile_13-4.11-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_profile_13-4.11-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_13` | `4.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 197.4 KiB | [pg_profile_13-4.10-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_profile_13-4.10-1PGDG.rhel10.noarch.rpm) |
| `pg_profile_13` | `4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [pg_profile_13-4.8-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_profile_13-4.8-1PGDG.rhel10.noarch.rpm) |
| `postgresql-13-pg-profile` | `4.11` | [d12.x86_64](/os/d12.x86_64) | pigsty | 192.3 KiB | [postgresql-13-pg-profile_4.11-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-13-pg-profile_4.11-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-profile` | `4.11` | [d12.aarch64](/os/d12.aarch64) | pigsty | 192.3 KiB | [postgresql-13-pg-profile_4.11-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-profile/postgresql-13-pg-profile_4.11-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-profile` | `4.11` | [d13.x86_64](/os/d13.x86_64) | pigsty | 192.3 KiB | [postgresql-13-pg-profile_4.11-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-13-pg-profile_4.11-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-profile` | `4.11` | [d13.aarch64](/os/d13.aarch64) | pigsty | 192.3 KiB | [postgresql-13-pg-profile_4.11-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-profile/postgresql-13-pg-profile_4.11-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-profile` | `4.11` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.2 KiB | [postgresql-13-pg-profile_4.11-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-13-pg-profile_4.11-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-profile` | `4.11` | [u22.aarch64](/os/u22.aarch64) | pigsty | 193.2 KiB | [postgresql-13-pg-profile_4.11-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-profile/postgresql-13-pg-profile_4.11-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-profile` | `4.11` | [u24.x86_64](/os/u24.x86_64) | pigsty | 191.6 KiB | [postgresql-13-pg-profile_4.11-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-13-pg-profile_4.11-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-profile` | `4.11` | [u24.aarch64](/os/u24.aarch64) | pigsty | 191.6 KiB | [postgresql-13-pg-profile_4.11-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-profile/postgresql-13-pg-profile_4.11-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/zubkov-andrei/pg_profile" title="Repository" icon="github" subtitle="github.com/zubkov-andrei/pg_profile" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_profile-4.11.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_profile;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_profile;		# install via package name, for the active PG version

pig install pg_profile -v 18;   # install for PG 18
pig install pg_profile -v 17;   # install for PG 17
pig install pg_profile -v 16;   # install for PG 16
pig install pg_profile -v 15;   # install for PG 15
pig install pg_profile -v 14;   # install for PG 14
pig install pg_profile -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_profile CASCADE; -- requires dblink, plpgsql
```
