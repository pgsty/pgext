---
title: "pg_ivm"
linkTitle: "pg_ivm"
description: "incremental view maintenance on PostgreSQL"
weight: 2840
categories: ["FEAT"]
width: full
---

[**pg_ivm**](https://github.com/sraoss/pg_ivm) : incremental view maintenance on PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2840** | {{< badge content="pg_ivm" link="https://github.com/sraoss/pg_ivm" >}} | {{< ext "pg_ivm" >}} | `1.14` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} |

> [!Note] deb takeover by pgdg since 2026-01


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.14` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_ivm` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.14` | {{< bg "18" "pg_ivm_18" "green" >}} {{< bg "17" "pg_ivm_17" "green" >}} {{< bg "16" "pg_ivm_16" "green" >}} {{< bg "15" "pg_ivm_15" "green" >}} {{< bg "14" "pg_ivm_14" "green" >}} | `pg_ivm_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.14` | {{< bg "18" "postgresql-18-pg-ivm" "green" >}} {{< bg "17" "postgresql-17-pg-ivm" "green" >}} {{< bg "16" "postgresql-16-pg-ivm" "green" >}} {{< bg "15" "postgresql-15-pg-ivm" "green" >}} {{< bg "14" "postgresql-14-pg-ivm" "green" >}} | `postgresql-$v-pg-ivm` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_15 : AVAIL 12" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_14 : AVAIL 16" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_15 : AVAIL 12" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_14 : AVAIL 12" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_15 : AVAIL 12" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_14 : AVAIL 15" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_15 : AVAIL 12" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_14 : AVAIL 12" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_14 : AVAIL 5" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.14" "pg_ivm_14 : AVAIL 5" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.14" "postgresql-18-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-17-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-16-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-15-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-14-pg-ivm : AVAIL 2" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.14" "postgresql-18-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-17-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-16-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-15-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-14-pg-ivm : AVAIL 2" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.14" "postgresql-18-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-17-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-16-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-15-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-14-pg-ivm : AVAIL 2" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.14" "postgresql-18-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-17-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-16-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-15-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-14-pg-ivm : AVAIL 2" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.14" "postgresql-18-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-17-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-16-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-15-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-14-pg-ivm : AVAIL 2" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.14" "postgresql-18-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-17-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-16-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-15-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-14-pg-ivm : AVAIL 2" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.14" "postgresql-18-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-17-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-16-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-15-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-14-pg-ivm : AVAIL 2" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.14" "postgresql-18-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-17-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-16-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-15-pg-ivm : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.14" "postgresql-14-pg-ivm : AVAIL 2" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.13" "postgresql-18-pg-ivm : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.13" "postgresql-17-pg-ivm : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.13" "postgresql-16-pg-ivm : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.13" "postgresql-15-pg-ivm : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.13" "postgresql-14-pg-ivm : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.13" "postgresql-18-pg-ivm : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.13" "postgresql-17-pg-ivm : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.13" "postgresql-16-pg-ivm : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.13" "postgresql-15-pg-ivm : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.13" "postgresql-14-pg-ivm : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_18` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pigsty | 57.9 KiB | [pg_ivm_18-1.14-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ivm_18-1.14-1PIGSTY.el8.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.3 KiB | [pg_ivm_18-1.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_ivm_18-1.14-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_18` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [pg_ivm_18-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_ivm_18-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_18` | `1.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.3 KiB | [pg_ivm_18-1.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_ivm_18-1.12-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.9 KiB | [pg_ivm_18-1.14-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ivm_18-1.14-1PIGSTY.el8.aarch64.rpm) |
| `pg_ivm_18` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.1 KiB | [pg_ivm_18-1.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_ivm_18-1.14-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_18` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.5 KiB | [pg_ivm_18-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_ivm_18-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_18` | `1.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.2 KiB | [pg_ivm_18-1.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_ivm_18-1.12-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_18` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pigsty | 57.5 KiB | [pg_ivm_18-1.14-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ivm_18-1.14-1PIGSTY.el9.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.7 KiB | [pg_ivm_18-1.14-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.14-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_ivm_18` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [pg_ivm_18-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_18` | `1.12` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.3 KiB | [pg_ivm_18-1.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.12-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pigsty | 56.3 KiB | [pg_ivm_18-1.14-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ivm_18-1.14-1PIGSTY.el9.aarch64.rpm) |
| `pg_ivm_18` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [pg_ivm_18-1.14-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.14-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_ivm_18` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.1 KiB | [pg_ivm_18-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_18` | `1.12` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.0 KiB | [pg_ivm_18-1.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.12-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_18` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pigsty | 58.5 KiB | [pg_ivm_18-1.14-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ivm_18-1.14-1PIGSTY.el10.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.8 KiB | [pg_ivm_18-1.14-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_ivm_18-1.14-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_ivm_18` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.6 KiB | [pg_ivm_18-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_ivm_18-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_18` | `1.12` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.1 KiB | [pg_ivm_18-1.12-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_ivm_18-1.12-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pigsty | 57.4 KiB | [pg_ivm_18-1.14-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ivm_18-1.14-1PIGSTY.el10.aarch64.rpm) |
| `pg_ivm_18` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.5 KiB | [pg_ivm_18-1.14-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_ivm_18-1.14-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_ivm_18` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_18-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_ivm_18-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_18` | `1.12` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pg_ivm_18-1.12-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_ivm_18-1.12-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-ivm` | `1.14` | [d12.x86_64](/os/d12.x86_64) | pigsty | 119.7 KiB | [postgresql-18-pg-ivm_1.14-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.14-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pgdg | 118.7 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.14` | [d12.aarch64](/os/d12.aarch64) | pigsty | 116.2 KiB | [postgresql-18-pg-ivm_1.14-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.14-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pgdg | 115.4 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.14` | [d13.x86_64](/os/d13.x86_64) | pigsty | 119.7 KiB | [postgresql-18-pg-ivm_1.14-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.14-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pgdg | 118.8 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.14` | [d13.aarch64](/os/d13.aarch64) | pigsty | 116.0 KiB | [postgresql-18-pg-ivm_1.14-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.14-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pgdg | 114.9 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.14` | [u22.x86_64](/os/u22.x86_64) | pigsty | 130.9 KiB | [postgresql-18-pg-ivm_1.14-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.14-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pgdg | 121.6 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.14` | [u22.aarch64](/os/u22.aarch64) | pigsty | 128.2 KiB | [postgresql-18-pg-ivm_1.14-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.14-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pgdg | 117.9 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.14` | [u24.x86_64](/os/u24.x86_64) | pigsty | 125.2 KiB | [postgresql-18-pg-ivm_1.14-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.14-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pgdg | 118.7 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.14` | [u24.aarch64](/os/u24.aarch64) | pigsty | 123.0 KiB | [postgresql-18-pg-ivm_1.14-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.14-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pgdg | 114.9 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u26.x86_64](/os/u26.x86_64) | pgdg | 117.1 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u26.aarch64](/os/u26.aarch64) | pgdg | 113.6 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_17` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pigsty | 57.8 KiB | [pg_ivm_17-1.14-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ivm_17-1.14-1PIGSTY.el8.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.2 KiB | [pg_ivm_17-1.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.14-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_17` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.4 KiB | [pg_ivm_17-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | `1.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.9 KiB | [pg_ivm_17-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.6 KiB | [pg_ivm_17-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | `1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.6 KiB | [pg_ivm_17-1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.9-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.8 KiB | [pg_ivm_17-1.14-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ivm_17-1.14-1PIGSTY.el8.aarch64.rpm) |
| `pg_ivm_17` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.0 KiB | [pg_ivm_17-1.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.14-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_17` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.3 KiB | [pg_ivm_17-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | `1.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.9 KiB | [pg_ivm_17-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.5 KiB | [pg_ivm_17-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | `1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [pg_ivm_17-1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.9-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pigsty | 57.6 KiB | [pg_ivm_17-1.14-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ivm_17-1.14-1PIGSTY.el9.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.7 KiB | [pg_ivm_17-1.14-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.14-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_ivm_17` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [pg_ivm_17-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | `1.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.2 KiB | [pg_ivm_17-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.9 KiB | [pg_ivm_17-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | `1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.4 KiB | [pg_ivm_17-1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.9-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pigsty | 56.2 KiB | [pg_ivm_17-1.14-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ivm_17-1.14-1PIGSTY.el9.aarch64.rpm) |
| `pg_ivm_17` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [pg_ivm_17-1.14-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.14-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_ivm_17` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.0 KiB | [pg_ivm_17-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | `1.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.9 KiB | [pg_ivm_17-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [pg_ivm_17-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.10-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | `1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.9 KiB | [pg_ivm_17-1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.9-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pigsty | 58.7 KiB | [pg_ivm_17-1.14-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ivm_17-1.14-1PIGSTY.el10.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.9 KiB | [pg_ivm_17-1.14-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.14-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_ivm_17` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.7 KiB | [pg_ivm_17-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_17` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.2 KiB | [pg_ivm_17-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_17` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.9 KiB | [pg_ivm_17-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pigsty | 57.6 KiB | [pg_ivm_17-1.14-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ivm_17-1.14-1PIGSTY.el10.aarch64.rpm) |
| `pg_ivm_17` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_17-1.14-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.14-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_ivm_17` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.6 KiB | [pg_ivm_17-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_17` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.7 KiB | [pg_ivm_17-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_17` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pg_ivm_17-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-ivm` | `1.14` | [d12.x86_64](/os/d12.x86_64) | pigsty | 119.6 KiB | [postgresql-17-pg-ivm_1.14-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.14-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pgdg | 118.3 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.14` | [d12.aarch64](/os/d12.aarch64) | pigsty | 116.1 KiB | [postgresql-17-pg-ivm_1.14-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.14-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pgdg | 115.2 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.14` | [d13.x86_64](/os/d13.x86_64) | pigsty | 119.3 KiB | [postgresql-17-pg-ivm_1.14-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.14-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pgdg | 118.2 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.14` | [d13.aarch64](/os/d13.aarch64) | pigsty | 115.7 KiB | [postgresql-17-pg-ivm_1.14-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.14-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pgdg | 114.8 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.14` | [u22.x86_64](/os/u22.x86_64) | pigsty | 151.3 KiB | [postgresql-17-pg-ivm_1.14-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.14-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pgdg | 141.1 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.14` | [u22.aarch64](/os/u22.aarch64) | pigsty | 148.7 KiB | [postgresql-17-pg-ivm_1.14-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.14-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pgdg | 137.8 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.14` | [u24.x86_64](/os/u24.x86_64) | pigsty | 124.9 KiB | [postgresql-17-pg-ivm_1.14-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.14-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pgdg | 118.3 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.14` | [u24.aarch64](/os/u24.aarch64) | pigsty | 122.8 KiB | [postgresql-17-pg-ivm_1.14-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.14-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pgdg | 114.7 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u26.x86_64](/os/u26.x86_64) | pgdg | 116.8 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u26.aarch64](/os/u26.aarch64) | pgdg | 113.6 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_16` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pigsty | 57.9 KiB | [pg_ivm_16-1.14-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ivm_16-1.14-1PIGSTY.el8.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.4 KiB | [pg_ivm_16-1.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.14-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_16` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [pg_ivm_16-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.0 KiB | [pg_ivm_16-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.7 KiB | [pg_ivm_16-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.9 KiB | [pg_ivm_16-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [pg_ivm_16-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.9 KiB | [pg_ivm_16-1.14-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ivm_16-1.14-1PIGSTY.el8.aarch64.rpm) |
| `pg_ivm_16` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.1 KiB | [pg_ivm_16-1.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.14-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_16` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.4 KiB | [pg_ivm_16-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.9 KiB | [pg_ivm_16-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.6 KiB | [pg_ivm_16-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.9 KiB | [pg_ivm_16-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.7 KiB | [pg_ivm_16-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pigsty | 57.7 KiB | [pg_ivm_16-1.14-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ivm_16-1.14-1PIGSTY.el9.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.7 KiB | [pg_ivm_16-1.14-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.14-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_ivm_16` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [pg_ivm_16-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.2 KiB | [pg_ivm_16-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.0 KiB | [pg_ivm_16-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.7 KiB | [pg_ivm_16-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.6 KiB | [pg_ivm_16-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pigsty | 56.4 KiB | [pg_ivm_16-1.14-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ivm_16-1.14-1PIGSTY.el9.aarch64.rpm) |
| `pg_ivm_16` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.4 KiB | [pg_ivm_16-1.14-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.14-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_ivm_16` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.2 KiB | [pg_ivm_16-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.0 KiB | [pg_ivm_16-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.7 KiB | [pg_ivm_16-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.10-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.3 KiB | [pg_ivm_16-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.4 KiB | [pg_ivm_16-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pigsty | 58.6 KiB | [pg_ivm_16-1.14-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ivm_16-1.14-1PIGSTY.el10.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.9 KiB | [pg_ivm_16-1.14-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.14-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_ivm_16` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.7 KiB | [pg_ivm_16-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_16` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [pg_ivm_16-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_16` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.9 KiB | [pg_ivm_16-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pigsty | 57.6 KiB | [pg_ivm_16-1.14-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ivm_16-1.14-1PIGSTY.el10.aarch64.rpm) |
| `pg_ivm_16` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_16-1.14-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.14-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_ivm_16` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_16-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_16` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.7 KiB | [pg_ivm_16-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_16` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pg_ivm_16-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-ivm` | `1.14` | [d12.x86_64](/os/d12.x86_64) | pigsty | 119.3 KiB | [postgresql-16-pg-ivm_1.14-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.14-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pgdg | 118.1 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.14` | [d12.aarch64](/os/d12.aarch64) | pigsty | 116.0 KiB | [postgresql-16-pg-ivm_1.14-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.14-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pgdg | 115.2 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.14` | [d13.x86_64](/os/d13.x86_64) | pigsty | 119.4 KiB | [postgresql-16-pg-ivm_1.14-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.14-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pgdg | 118.1 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.14` | [d13.aarch64](/os/d13.aarch64) | pigsty | 115.6 KiB | [postgresql-16-pg-ivm_1.14-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.14-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pgdg | 114.8 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.14` | [u22.x86_64](/os/u22.x86_64) | pigsty | 150.2 KiB | [postgresql-16-pg-ivm_1.14-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.14-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pgdg | 140.2 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.14` | [u22.aarch64](/os/u22.aarch64) | pigsty | 147.5 KiB | [postgresql-16-pg-ivm_1.14-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.14-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pgdg | 136.5 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.14` | [u24.x86_64](/os/u24.x86_64) | pigsty | 124.9 KiB | [postgresql-16-pg-ivm_1.14-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.14-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pgdg | 118.1 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.14` | [u24.aarch64](/os/u24.aarch64) | pigsty | 122.7 KiB | [postgresql-16-pg-ivm_1.14-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.14-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pgdg | 114.7 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u26.x86_64](/os/u26.x86_64) | pgdg | 117.0 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u26.aarch64](/os/u26.aarch64) | pgdg | 113.5 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_15` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pigsty | 58.1 KiB | [pg_ivm_15-1.14-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ivm_15-1.14-1PIGSTY.el8.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.5 KiB | [pg_ivm_15-1.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.14-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_15` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.8 KiB | [pg_ivm_15-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.3 KiB | [pg_ivm_15-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.0 KiB | [pg_ivm_15-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.3 KiB | [pg_ivm_15-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_ivm_15-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.6 KiB | [pg_ivm_15-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.1 KiB | [pg_ivm_15-1.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.5.1-1.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.2 KiB | [pg_ivm_15-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.5-1.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.3 KiB | [pg_ivm_15-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.4-1.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.8 KiB | [pg_ivm_15-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.3-1.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pigsty | 56.2 KiB | [pg_ivm_15-1.14-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ivm_15-1.14-1PIGSTY.el8.aarch64.rpm) |
| `pg_ivm_15` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.3 KiB | [pg_ivm_15-1.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.14-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_15` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.6 KiB | [pg_ivm_15-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.3 KiB | [pg_ivm_15-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.9 KiB | [pg_ivm_15-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.1 KiB | [pg_ivm_15-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.0 KiB | [pg_ivm_15-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.8 KiB | [pg_ivm_15-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.5 KiB | [pg_ivm_15-1.5.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.5.1-1.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.6 KiB | [pg_ivm_15-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.5-1.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.4 KiB | [pg_ivm_15-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.4-1.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.0 KiB | [pg_ivm_15-1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.3-1.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pigsty | 58.1 KiB | [pg_ivm_15-1.14-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ivm_15-1.14-1PIGSTY.el9.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.2 KiB | [pg_ivm_15-1.14-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.14-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_ivm_15` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.0 KiB | [pg_ivm_15-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.0 KiB | [pg_ivm_15-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.7 KiB | [pg_ivm_15-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.6 KiB | [pg_ivm_15-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.6 KiB | [pg_ivm_15-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.4 KiB | [pg_ivm_15-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.0 KiB | [pg_ivm_15-1.5.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.5.1-1.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.1 KiB | [pg_ivm_15-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.5-1.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.3 KiB | [pg_ivm_15-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.4-1.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.7 KiB | [pg_ivm_15-1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.3-1.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pigsty | 57.0 KiB | [pg_ivm_15-1.14-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ivm_15-1.14-1PIGSTY.el9.aarch64.rpm) |
| `pg_ivm_15` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.4 KiB | [pg_ivm_15-1.14-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.14-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_ivm_15` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.8 KiB | [pg_ivm_15-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.8 KiB | [pg_ivm_15-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.5 KiB | [pg_ivm_15-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.10-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.9 KiB | [pg_ivm_15-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.2 KiB | [pg_ivm_15-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.0 KiB | [pg_ivm_15-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [pg_ivm_15-1.5.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.5.1-1.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [pg_ivm_15-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.5-1.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.7 KiB | [pg_ivm_15-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.4-1.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.3 KiB | [pg_ivm_15-1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.3-1.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pigsty | 59.1 KiB | [pg_ivm_15-1.14-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ivm_15-1.14-1PIGSTY.el10.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.5 KiB | [pg_ivm_15-1.14-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.14-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_ivm_15` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.5 KiB | [pg_ivm_15-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_15` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.0 KiB | [pg_ivm_15-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_15` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.7 KiB | [pg_ivm_15-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pigsty | 57.9 KiB | [pg_ivm_15-1.14-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ivm_15-1.14-1PIGSTY.el10.aarch64.rpm) |
| `pg_ivm_15` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.3 KiB | [pg_ivm_15-1.14-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.14-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_ivm_15` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.4 KiB | [pg_ivm_15-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_15` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.5 KiB | [pg_ivm_15-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_15` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.1 KiB | [pg_ivm_15-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-ivm` | `1.14` | [d12.x86_64](/os/d12.x86_64) | pigsty | 119.3 KiB | [postgresql-15-pg-ivm_1.14-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.14-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pgdg | 118.7 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.14` | [d12.aarch64](/os/d12.aarch64) | pigsty | 115.9 KiB | [postgresql-15-pg-ivm_1.14-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.14-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pgdg | 115.2 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.14` | [d13.x86_64](/os/d13.x86_64) | pigsty | 119.4 KiB | [postgresql-15-pg-ivm_1.14-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.14-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pgdg | 118.4 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.14` | [d13.aarch64](/os/d13.aarch64) | pigsty | 115.7 KiB | [postgresql-15-pg-ivm_1.14-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.14-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pgdg | 114.9 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.14` | [u22.x86_64](/os/u22.x86_64) | pigsty | 149.9 KiB | [postgresql-15-pg-ivm_1.14-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.14-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pgdg | 140.3 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.14` | [u22.aarch64](/os/u22.aarch64) | pigsty | 147.8 KiB | [postgresql-15-pg-ivm_1.14-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.14-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pgdg | 136.3 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.14` | [u24.x86_64](/os/u24.x86_64) | pigsty | 124.8 KiB | [postgresql-15-pg-ivm_1.14-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.14-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pgdg | 118.3 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.14` | [u24.aarch64](/os/u24.aarch64) | pigsty | 123.0 KiB | [postgresql-15-pg-ivm_1.14-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.14-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pgdg | 115.2 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u26.x86_64](/os/u26.x86_64) | pgdg | 117.1 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u26.aarch64](/os/u26.aarch64) | pgdg | 113.6 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_14` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pigsty | 87.8 KiB | [pg_ivm_14-1.14-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ivm_14-1.14-1PIGSTY.el8.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.8 KiB | [pg_ivm_14-1.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.14-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_14` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.0 KiB | [pg_ivm_14-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 71.8 KiB | [pg_ivm_14-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 71.5 KiB | [pg_ivm_14-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.6 KiB | [pg_ivm_14-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 71.6 KiB | [pg_ivm_14-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 71.4 KiB | [pg_ivm_14-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.0 KiB | [pg_ivm_14-1.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.5.1-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.1 KiB | [pg_ivm_14-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.5-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.2 KiB | [pg_ivm_14-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.4-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.6 KiB | [pg_ivm_14-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.3-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 66.2 KiB | [pg_ivm_14-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.2-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.4 KiB | [pg_ivm_14-1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.1-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.8 KiB | [pg_ivm_14-1.0-.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.0-.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.9 KiB | [pg_ivm_14-1.0-alpha1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.0-alpha1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pigsty | 83.2 KiB | [pg_ivm_14-1.14-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ivm_14-1.14-1PIGSTY.el8.aarch64.rpm) |
| `pg_ivm_14` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.1 KiB | [pg_ivm_14-1.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.14-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_14` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.4 KiB | [pg_ivm_14-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 67.1 KiB | [pg_ivm_14-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.8 KiB | [pg_ivm_14-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.0 KiB | [pg_ivm_14-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 67.0 KiB | [pg_ivm_14-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.8 KiB | [pg_ivm_14-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.7 KiB | [pg_ivm_14-1.5.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.5.1-1.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.9 KiB | [pg_ivm_14-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.5-1.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.8 KiB | [pg_ivm_14-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.4-1.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.4 KiB | [pg_ivm_14-1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.3-1.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pigsty | 87.7 KiB | [pg_ivm_14-1.14-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ivm_14-1.14-1PIGSTY.el9.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 79.8 KiB | [pg_ivm_14-1.14-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.14-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_ivm_14` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 79.3 KiB | [pg_ivm_14-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.5 KiB | [pg_ivm_14-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.2 KiB | [pg_ivm_14-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 71.0 KiB | [pg_ivm_14-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.4 KiB | [pg_ivm_14-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.6 KiB | [pg_ivm_14-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.3 KiB | [pg_ivm_14-1.5.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.5.1-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.4 KiB | [pg_ivm_14-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.5-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 71.5 KiB | [pg_ivm_14-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.4-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 71.0 KiB | [pg_ivm_14-1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.3-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.4 KiB | [pg_ivm_14-1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.2-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.5 KiB | [pg_ivm_14-1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.1-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.2 KiB | [pg_ivm_14-1.0-.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.0-.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pigsty | 85.5 KiB | [pg_ivm_14-1.14-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ivm_14-1.14-1PIGSTY.el9.aarch64.rpm) |
| `pg_ivm_14` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 77.4 KiB | [pg_ivm_14-1.14-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.14-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_ivm_14` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 77.0 KiB | [pg_ivm_14-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 70.9 KiB | [pg_ivm_14-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 70.6 KiB | [pg_ivm_14-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.10-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.2 KiB | [pg_ivm_14-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.6 KiB | [pg_ivm_14-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.4 KiB | [pg_ivm_14-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.2 KiB | [pg_ivm_14-1.5.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.5.1-1.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.3 KiB | [pg_ivm_14-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.5-1.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.3 KiB | [pg_ivm_14-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.4-1.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.0 KiB | [pg_ivm_14-1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.3-1.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pigsty | 89.1 KiB | [pg_ivm_14-1.14-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ivm_14-1.14-1PIGSTY.el10.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 81.1 KiB | [pg_ivm_14-1.14-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.14-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_ivm_14` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 80.8 KiB | [pg_ivm_14-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_14` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 74.9 KiB | [pg_ivm_14-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_14` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 74.6 KiB | [pg_ivm_14-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pigsty | 86.9 KiB | [pg_ivm_14-1.14-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ivm_14-1.14-1PIGSTY.el10.aarch64.rpm) |
| `pg_ivm_14` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 78.9 KiB | [pg_ivm_14-1.14-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.14-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_ivm_14` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 79.1 KiB | [pg_ivm_14-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_14` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 72.4 KiB | [pg_ivm_14-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_14` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 72.1 KiB | [pg_ivm_14-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-ivm` | `1.14` | [d12.x86_64](/os/d12.x86_64) | pigsty | 209.8 KiB | [postgresql-14-pg-ivm_1.14-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.14-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pgdg | 209.0 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.14` | [d12.aarch64](/os/d12.aarch64) | pigsty | 202.6 KiB | [postgresql-14-pg-ivm_1.14-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.14-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pgdg | 201.9 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.14` | [d13.x86_64](/os/d13.x86_64) | pigsty | 209.6 KiB | [postgresql-14-pg-ivm_1.14-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.14-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pgdg | 208.6 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.14` | [d13.aarch64](/os/d13.aarch64) | pigsty | 202.8 KiB | [postgresql-14-pg-ivm_1.14-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.14-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pgdg | 201.9 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.14` | [u22.x86_64](/os/u22.x86_64) | pigsty | 253.7 KiB | [postgresql-14-pg-ivm_1.14-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.14-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pgdg | 238.7 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.14` | [u22.aarch64](/os/u22.aarch64) | pigsty | 250.0 KiB | [postgresql-14-pg-ivm_1.14-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.14-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pgdg | 230.9 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.14` | [u24.x86_64](/os/u24.x86_64) | pigsty | 218.9 KiB | [postgresql-14-pg-ivm_1.14-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.14-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pgdg | 208.9 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.14` | [u24.aarch64](/os/u24.aarch64) | pigsty | 216.0 KiB | [postgresql-14-pg-ivm_1.14-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.14-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pgdg | 202.0 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u26.x86_64](/os/u26.x86_64) | pgdg | 206.0 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u26.aarch64](/os/u26.aarch64) | pgdg | 198.6 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/sraoss/pg_ivm" title="Repository" icon="github" subtitle="github.com/sraoss/pg_ivm" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_ivm-1.13.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_ivm;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_ivm;		# install via package name, for the active PG version

pig install pg_ivm -v 18;   # install for PG 18
pig install pg_ivm -v 17;   # install for PG 17
pig install pg_ivm -v 16;   # install for PG 16
pig install pg_ivm -v 15;   # install for PG 15
pig install pg_ivm -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_ivm';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_ivm;
```

## Usage

Sources: [README](https://github.com/sraoss/pg_ivm/blob/master/README.md), [release 1.14](https://github.com/sraoss/pg_ivm/releases/tag/v1.14)

`pg_ivm` provides immediate Incremental View Maintenance for PostgreSQL materialized views. Instead of recomputing the whole view, it applies deltas in `AFTER` triggers and stores metadata in the `pgivm` schema.

```sql
CREATE EXTENSION pg_ivm;
```

### Required Setup

Upstream says `pg_ivm` should be preloaded so IMMVs are maintained correctly:

```conf
shared_preload_libraries = 'pg_ivm'
session_preload_libraries = 'pg_ivm'
```

The current README says the extension is compatible with PostgreSQL 13 through 18, and the latest GitHub release is `1.14` dated March 31, 2026.

### Main Functions

- `pgivm.create_immv(name, query)` creates an incrementally maintainable materialized view (IMMV), its maintenance triggers, and a unique index when possible.
- `pgivm.refresh_immv(name, with_data)` fully refreshes the IMMV and can disable or re-enable maintenance.
- `pgivm.get_immv_def(regclass)` reconstructs the stored `SELECT` definition.
- `pgivm.pg_ivm_immv` stores IMMV metadata including `immvrelid`, `viewdef`, `ispopulated`, and `lastivmupdate`.

### Common Patterns

Create an IMMV:

```sql
SELECT pgivm.create_immv(
  'immv_agg',
  'SELECT bid, count(*), sum(abalance), avg(abalance)
   FROM pgbench_accounts JOIN pgbench_branches USING(bid)
   GROUP BY bid'
);
```

Query the maintained result after base-table changes:

```sql
UPDATE pgbench_accounts SET abalance = abalance + 1000 WHERE aid = 4112345;
SELECT * FROM immv_agg WHERE bid = 42;
```

Inspect or refresh IMMVs:

```sql
SELECT immvrelid AS immv, pgivm.get_immv_def(immvrelid)
FROM pgivm.pg_ivm_immv;

SELECT pgivm.refresh_immv('immv_agg', true);
```

Pause maintenance for bulk work, then rebuild:

```sql
SELECT pgivm.refresh_immv('myview', false);
-- bulk changes
SELECT pgivm.refresh_immv('myview', true);
```

### Caveats

- Upstream only supports a restricted subset of view definitions: joins, `DISTINCT`, simple subqueries/CTEs, and built-in aggregates `count`, `sum`, `avg`, `min`, and `max`.
- Unsupported constructs include `HAVING`, window functions, `ORDER BY`, `LIMIT/OFFSET`, `UNION`/`INTERSECT`/`EXCEPT`, `DISTINCT ON`, and user-defined aggregates.
- Efficient maintenance depends on having a suitable unique index; `create_immv` creates one automatically only when the definition allows it.
