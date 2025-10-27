---
title: "pg_ivm"
linkTitle: "pg_ivm"
description: "incremental view maintenance on PostgreSQL"
weight: 2870
categories: ["FEAT"]
width: full
---

incremental view maintenance on PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2870** | {{< badge content="pg_ivm" link="https://github.com/sraoss/pg_ivm" >}} | {{< ext "pg_ivm" >}} | `1.13` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "timeseries" >}} |
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_ivm" >}} | `1.13` | {{< bg "18" "pg_ivm_18*" "green" >}} {{< bg "17" "pg_ivm_17*" "green" >}} {{< bg "16" "pg_ivm_16*" "green" >}} {{< bg "15" "pg_ivm_15*" "green" >}} {{< bg "14" "pg_ivm_14*" "green" >}} | `pg_ivm_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_ivm" >}} | `1.12` | {{< bg "18" "postgresql-18-pg-ivm" "green" >}} {{< bg "17" "postgresql-17-pg-ivm" "green" >}} {{< bg "16" "postgresql-16-pg-ivm" "green" >}} {{< bg "15" "postgresql-15-pg-ivm" "green" >}} {{< bg "14" "postgresql-14-pg-ivm" "green" >}} | `postgresql-$v-pg-ivm` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.13" "pg_ivm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.9" "pg_ivm_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.8" "pg_ivm_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.8" "pg_ivm_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.8" "pg_ivm_14 : AVAIL 14" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.13" "pg_ivm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.9" "pg_ivm_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.8" "pg_ivm_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.8" "pg_ivm_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.8" "pg_ivm_14 : AVAIL 10" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.13" "pg_ivm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.9" "pg_ivm_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.8" "pg_ivm_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.8" "pg_ivm_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.8" "pg_ivm_14 : AVAIL 13" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.13" "pg_ivm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.9" "pg_ivm_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.8" "pg_ivm_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.8" "pg_ivm_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.8" "pg_ivm_14 : AVAIL 10" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-ivm : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.12" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-ivm : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.12" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-ivm : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.12" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-ivm : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.12" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-ivm : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.12" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-ivm : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.12" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.12" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_18` | 1.13 | `el8.x86_64` | pgdg | 49.5 KiB | [pg_ivm_18-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_ivm_18-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_18` | 1.12 | `el8.x86_64` | pgdg | 43.3 KiB | [pg_ivm_18-1.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_ivm_18-1.12-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_18` | 1.13 | `el8.aarch64` | pgdg | 47.5 KiB | [pg_ivm_18-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_ivm_18-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_18` | 1.12 | `el8.aarch64` | pgdg | 41.2 KiB | [pg_ivm_18-1.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_ivm_18-1.12-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_18` | 1.13 | `el9.x86_64` | pgdg | 49.3 KiB | [pg_ivm_18-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_18` | 1.12 | `el9.x86_64` | pgdg | 43.3 KiB | [pg_ivm_18-1.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.12-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_18` | 1.13 | `el9.aarch64` | pgdg | 48.1 KiB | [pg_ivm_18-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_18` | 1.12 | `el9.aarch64` | pgdg | 42.0 KiB | [pg_ivm_18-1.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.12-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_17` | 1.9 | `el8.x86_64` | pgdg | 40.6 KiB | [pg_ivm_17-1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.9-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | 1.13 | `el8.x86_64` | pgdg | 49.4 KiB | [pg_ivm_17-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | 1.11 | `el8.x86_64` | pgdg | 42.9 KiB | [pg_ivm_17-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | 1.10 | `el8.x86_64` | pgdg | 42.6 KiB | [pg_ivm_17-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | 1.9 | `el8.aarch64` | pgdg | 38.6 KiB | [pg_ivm_17-1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.9-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | 1.13 | `el8.aarch64` | pgdg | 47.3 KiB | [pg_ivm_17-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | 1.11 | `el8.aarch64` | pgdg | 40.9 KiB | [pg_ivm_17-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | 1.10 | `el8.aarch64` | pgdg | 40.5 KiB | [pg_ivm_17-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | 1.9 | `el9.x86_64` | pgdg | 41.4 KiB | [pg_ivm_17-1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.9-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | 1.13 | `el9.x86_64` | pgdg | 49.3 KiB | [pg_ivm_17-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | 1.11 | `el9.x86_64` | pgdg | 43.2 KiB | [pg_ivm_17-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | 1.10 | `el9.x86_64` | pgdg | 42.9 KiB | [pg_ivm_17-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | 1.9 | `el9.aarch64` | pgdg | 39.9 KiB | [pg_ivm_17-1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.9-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | 1.13 | `el9.aarch64` | pgdg | 48.0 KiB | [pg_ivm_17-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | 1.11 | `el9.aarch64` | pgdg | 41.9 KiB | [pg_ivm_17-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | 1.10 | `el9.aarch64` | pgdg | 41.6 KiB | [pg_ivm_17-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pg-ivm` | 1.12 | `d12.x86_64` | pigsty | 102.0 KiB | [postgresql-17-pg-ivm_1.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-ivm` | 1.12 | `d12.aarch64` | pigsty | 98.9 KiB | [postgresql-17-pg-ivm_1.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-ivm` | 1.12 | `u22.x86_64` | pigsty | 131.8 KiB | [postgresql-17-pg-ivm_1.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-ivm` | 1.12 | `u22.aarch64` | pigsty | 129.7 KiB | [postgresql-17-pg-ivm_1.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-ivm` | 1.12 | `u24.x86_64` | pigsty | 106.7 KiB | [postgresql-17-pg-ivm_1.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-ivm` | 1.12 | `u24.aarch64` | pigsty | 105.1 KiB | [postgresql-17-pg-ivm_1.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_16` | 1.8 | `el8.x86_64` | pgdg | 39.9 KiB | [pg_ivm_16-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | 1.7 | `el8.x86_64` | pgdg | 41.5 KiB | [pg_ivm_16-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | 1.13 | `el8.x86_64` | pgdg | 49.5 KiB | [pg_ivm_16-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | 1.11 | `el8.x86_64` | pgdg | 43.0 KiB | [pg_ivm_16-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | 1.10 | `el8.x86_64` | pgdg | 42.7 KiB | [pg_ivm_16-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | 1.8 | `el8.aarch64` | pgdg | 37.9 KiB | [pg_ivm_16-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | 1.7 | `el8.aarch64` | pgdg | 39.7 KiB | [pg_ivm_16-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | 1.13 | `el8.aarch64` | pgdg | 47.4 KiB | [pg_ivm_16-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | 1.11 | `el8.aarch64` | pgdg | 40.9 KiB | [pg_ivm_16-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | 1.10 | `el8.aarch64` | pgdg | 40.6 KiB | [pg_ivm_16-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | 1.8 | `el9.x86_64` | pgdg | 40.7 KiB | [pg_ivm_16-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | 1.7 | `el9.x86_64` | pgdg | 42.6 KiB | [pg_ivm_16-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | 1.13 | `el9.x86_64` | pgdg | 49.3 KiB | [pg_ivm_16-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | 1.11 | `el9.x86_64` | pgdg | 43.2 KiB | [pg_ivm_16-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | 1.10 | `el9.x86_64` | pgdg | 43.0 KiB | [pg_ivm_16-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | 1.8 | `el9.aarch64` | pgdg | 39.3 KiB | [pg_ivm_16-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | 1.7 | `el9.aarch64` | pgdg | 41.4 KiB | [pg_ivm_16-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | 1.13 | `el9.aarch64` | pgdg | 48.2 KiB | [pg_ivm_16-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | 1.11 | `el9.aarch64` | pgdg | 42.0 KiB | [pg_ivm_16-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | 1.10 | `el9.aarch64` | pgdg | 41.7 KiB | [pg_ivm_16-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pg-ivm` | 1.12 | `d12.x86_64` | pigsty | 102.1 KiB | [postgresql-16-pg-ivm_1.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-ivm` | 1.12 | `d12.aarch64` | pigsty | 99.0 KiB | [postgresql-16-pg-ivm_1.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-ivm` | 1.12 | `u22.x86_64` | pigsty | 130.0 KiB | [postgresql-16-pg-ivm_1.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-ivm` | 1.12 | `u22.aarch64` | pigsty | 128.1 KiB | [postgresql-16-pg-ivm_1.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-ivm` | 1.12 | `u24.x86_64` | pigsty | 106.7 KiB | [postgresql-16-pg-ivm_1.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-ivm` | 1.12 | `u24.aarch64` | pigsty | 105.2 KiB | [postgresql-16-pg-ivm_1.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_15` | 1.8 | `el8.x86_64` | pgdg | 40.3 KiB | [pg_ivm_15-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | 1.7 | `el8.x86_64` | pgdg | 41.8 KiB | [pg_ivm_15-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | 1.6 | `el8.x86_64` | pgdg | 41.6 KiB | [pg_ivm_15-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | 1.5.1 | `el8.x86_64` | pgdg | 39.1 KiB | [pg_ivm_15-1.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.5.1-1.rhel8.x86_64.rpm) |
| `pg_ivm_15` | 1.5 | `el8.x86_64` | pgdg | 39.2 KiB | [pg_ivm_15-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.5-1.rhel8.x86_64.rpm) |
| `pg_ivm_15` | 1.4 | `el8.x86_64` | pgdg | 38.3 KiB | [pg_ivm_15-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.4-1.rhel8.x86_64.rpm) |
| `pg_ivm_15` | 1.3 | `el8.x86_64` | pgdg | 37.8 KiB | [pg_ivm_15-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.3-1.rhel8.x86_64.rpm) |
| `pg_ivm_15` | 1.13 | `el8.x86_64` | pgdg | 49.8 KiB | [pg_ivm_15-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | 1.11 | `el8.x86_64` | pgdg | 43.3 KiB | [pg_ivm_15-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | 1.10 | `el8.x86_64` | pgdg | 43.0 KiB | [pg_ivm_15-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | 1.8 | `el8.aarch64` | pgdg | 38.1 KiB | [pg_ivm_15-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | 1.7 | `el8.aarch64` | pgdg | 40.0 KiB | [pg_ivm_15-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | 1.6 | `el8.aarch64` | pgdg | 39.8 KiB | [pg_ivm_15-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | 1.5.1 | `el8.aarch64` | pgdg | 37.5 KiB | [pg_ivm_15-1.5.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.5.1-1.rhel8.aarch64.rpm) |
| `pg_ivm_15` | 1.5 | `el8.aarch64` | pgdg | 37.6 KiB | [pg_ivm_15-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.5-1.rhel8.aarch64.rpm) |
| `pg_ivm_15` | 1.4 | `el8.aarch64` | pgdg | 36.4 KiB | [pg_ivm_15-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.4-1.rhel8.aarch64.rpm) |
| `pg_ivm_15` | 1.3 | `el8.aarch64` | pgdg | 36.0 KiB | [pg_ivm_15-1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.3-1.rhel8.aarch64.rpm) |
| `pg_ivm_15` | 1.13 | `el8.aarch64` | pgdg | 47.6 KiB | [pg_ivm_15-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | 1.11 | `el8.aarch64` | pgdg | 41.3 KiB | [pg_ivm_15-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | 1.10 | `el8.aarch64` | pgdg | 40.9 KiB | [pg_ivm_15-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | 1.8 | `el9.x86_64` | pgdg | 41.6 KiB | [pg_ivm_15-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | 1.7 | `el9.x86_64` | pgdg | 43.6 KiB | [pg_ivm_15-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | 1.6 | `el9.x86_64` | pgdg | 43.4 KiB | [pg_ivm_15-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | 1.5.1 | `el9.x86_64` | pgdg | 41.0 KiB | [pg_ivm_15-1.5.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.5.1-1.rhel9.x86_64.rpm) |
| `pg_ivm_15` | 1.5 | `el9.x86_64` | pgdg | 41.1 KiB | [pg_ivm_15-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.5-1.rhel9.x86_64.rpm) |
| `pg_ivm_15` | 1.4 | `el9.x86_64` | pgdg | 40.3 KiB | [pg_ivm_15-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.4-1.rhel9.x86_64.rpm) |
| `pg_ivm_15` | 1.3 | `el9.x86_64` | pgdg | 39.7 KiB | [pg_ivm_15-1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.3-1.rhel9.x86_64.rpm) |
| `pg_ivm_15` | 1.13 | `el9.x86_64` | pgdg | 50.0 KiB | [pg_ivm_15-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | 1.11 | `el9.x86_64` | pgdg | 44.0 KiB | [pg_ivm_15-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | 1.10 | `el9.x86_64` | pgdg | 43.7 KiB | [pg_ivm_15-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | 1.8 | `el9.aarch64` | pgdg | 39.9 KiB | [pg_ivm_15-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | 1.7 | `el9.aarch64` | pgdg | 42.2 KiB | [pg_ivm_15-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | 1.6 | `el9.aarch64` | pgdg | 42.0 KiB | [pg_ivm_15-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | 1.5.1 | `el9.aarch64` | pgdg | 39.7 KiB | [pg_ivm_15-1.5.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.5.1-1.rhel9.aarch64.rpm) |
| `pg_ivm_15` | 1.5 | `el9.aarch64` | pgdg | 39.7 KiB | [pg_ivm_15-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.5-1.rhel9.aarch64.rpm) |
| `pg_ivm_15` | 1.4 | `el9.aarch64` | pgdg | 38.7 KiB | [pg_ivm_15-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.4-1.rhel9.aarch64.rpm) |
| `pg_ivm_15` | 1.3 | `el9.aarch64` | pgdg | 38.3 KiB | [pg_ivm_15-1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.3-1.rhel9.aarch64.rpm) |
| `pg_ivm_15` | 1.13 | `el9.aarch64` | pgdg | 48.8 KiB | [pg_ivm_15-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | 1.11 | `el9.aarch64` | pgdg | 42.8 KiB | [pg_ivm_15-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | 1.10 | `el9.aarch64` | pgdg | 42.5 KiB | [pg_ivm_15-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-pg-ivm` | 1.12 | `d12.x86_64` | pigsty | 102.7 KiB | [postgresql-15-pg-ivm_1.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-ivm` | 1.12 | `d12.aarch64` | pigsty | 99.3 KiB | [postgresql-15-pg-ivm_1.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-ivm` | 1.12 | `u22.x86_64` | pigsty | 130.9 KiB | [postgresql-15-pg-ivm_1.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-ivm` | 1.12 | `u22.aarch64` | pigsty | 128.6 KiB | [postgresql-15-pg-ivm_1.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-ivm` | 1.12 | `u24.x86_64` | pigsty | 107.4 KiB | [postgresql-15-pg-ivm_1.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-ivm` | 1.12 | `u24.aarch64` | pigsty | 105.9 KiB | [postgresql-15-pg-ivm_1.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_14` | 1.8 | `el8.x86_64` | pgdg | 68.6 KiB | [pg_ivm_14-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.7 | `el8.x86_64` | pgdg | 71.6 KiB | [pg_ivm_14-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.6 | `el8.x86_64` | pgdg | 71.4 KiB | [pg_ivm_14-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.5.1 | `el8.x86_64` | pgdg | 69.0 KiB | [pg_ivm_14-1.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.5.1-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.5 | `el8.x86_64` | pgdg | 69.1 KiB | [pg_ivm_14-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.5-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.4 | `el8.x86_64` | pgdg | 68.2 KiB | [pg_ivm_14-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.4-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.3 | `el8.x86_64` | pgdg | 67.6 KiB | [pg_ivm_14-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.3-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.2 | `el8.x86_64` | pgdg | 66.2 KiB | [pg_ivm_14-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.2-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.13 | `el8.x86_64` | pgdg | 78.0 KiB | [pg_ivm_14-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.11 | `el8.x86_64` | pgdg | 71.8 KiB | [pg_ivm_14-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.10 | `el8.x86_64` | pgdg | 71.5 KiB | [pg_ivm_14-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.1 | `el8.x86_64` | pgdg | 32.4 KiB | [pg_ivm_14-1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.1-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.0 | `el8.x86_64` | pgdg | 62.9 KiB | [pg_ivm_14-1.0-alpha1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.0-alpha1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.0 | `el8.x86_64` | pgdg | 74.8 KiB | [pg_ivm_14-1.0-.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.0-.rhel8.x86_64.rpm) |
| `pg_ivm_14` | 1.8 | `el8.aarch64` | pgdg | 64.0 KiB | [pg_ivm_14-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | 1.7 | `el8.aarch64` | pgdg | 67.0 KiB | [pg_ivm_14-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | 1.6 | `el8.aarch64` | pgdg | 66.8 KiB | [pg_ivm_14-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | 1.5.1 | `el8.aarch64` | pgdg | 64.7 KiB | [pg_ivm_14-1.5.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.5.1-1.rhel8.aarch64.rpm) |
| `pg_ivm_14` | 1.5 | `el8.aarch64` | pgdg | 64.9 KiB | [pg_ivm_14-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.5-1.rhel8.aarch64.rpm) |
| `pg_ivm_14` | 1.4 | `el8.aarch64` | pgdg | 63.8 KiB | [pg_ivm_14-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.4-1.rhel8.aarch64.rpm) |
| `pg_ivm_14` | 1.3 | `el8.aarch64` | pgdg | 63.4 KiB | [pg_ivm_14-1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.3-1.rhel8.aarch64.rpm) |
| `pg_ivm_14` | 1.13 | `el8.aarch64` | pgdg | 73.4 KiB | [pg_ivm_14-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | 1.11 | `el8.aarch64` | pgdg | 67.1 KiB | [pg_ivm_14-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | 1.10 | `el8.aarch64` | pgdg | 66.8 KiB | [pg_ivm_14-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | 1.8 | `el9.x86_64` | pgdg | 71.0 KiB | [pg_ivm_14-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.7 | `el9.x86_64` | pgdg | 74.4 KiB | [pg_ivm_14-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.6 | `el9.x86_64` | pgdg | 74.6 KiB | [pg_ivm_14-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.5.1 | `el9.x86_64` | pgdg | 72.3 KiB | [pg_ivm_14-1.5.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.5.1-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.5 | `el9.x86_64` | pgdg | 72.4 KiB | [pg_ivm_14-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.5-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.4 | `el9.x86_64` | pgdg | 71.5 KiB | [pg_ivm_14-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.4-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.3 | `el9.x86_64` | pgdg | 71.0 KiB | [pg_ivm_14-1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.3-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.2 | `el9.x86_64` | pgdg | 69.4 KiB | [pg_ivm_14-1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.2-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.13 | `el9.x86_64` | pgdg | 79.3 KiB | [pg_ivm_14-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.11 | `el9.x86_64` | pgdg | 73.5 KiB | [pg_ivm_14-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.10 | `el9.x86_64` | pgdg | 73.2 KiB | [pg_ivm_14-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.1 | `el9.x86_64` | pgdg | 34.5 KiB | [pg_ivm_14-1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.1-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.0 | `el9.x86_64` | pgdg | 77.2 KiB | [pg_ivm_14-1.0-.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.0-.rhel9.x86_64.rpm) |
| `pg_ivm_14` | 1.8 | `el9.aarch64` | pgdg | 68.2 KiB | [pg_ivm_14-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | 1.7 | `el9.aarch64` | pgdg | 71.6 KiB | [pg_ivm_14-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | 1.6 | `el9.aarch64` | pgdg | 71.4 KiB | [pg_ivm_14-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | 1.5.1 | `el9.aarch64` | pgdg | 69.2 KiB | [pg_ivm_14-1.5.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.5.1-1.rhel9.aarch64.rpm) |
| `pg_ivm_14` | 1.5 | `el9.aarch64` | pgdg | 69.3 KiB | [pg_ivm_14-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.5-1.rhel9.aarch64.rpm) |
| `pg_ivm_14` | 1.4 | `el9.aarch64` | pgdg | 68.3 KiB | [pg_ivm_14-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.4-1.rhel9.aarch64.rpm) |
| `pg_ivm_14` | 1.3 | `el9.aarch64` | pgdg | 68.0 KiB | [pg_ivm_14-1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.3-1.rhel9.aarch64.rpm) |
| `pg_ivm_14` | 1.13 | `el9.aarch64` | pgdg | 77.0 KiB | [pg_ivm_14-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | 1.11 | `el9.aarch64` | pgdg | 70.9 KiB | [pg_ivm_14-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | 1.10 | `el9.aarch64` | pgdg | 70.6 KiB | [pg_ivm_14-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.10-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-14-pg-ivm` | 1.12 | `d12.x86_64` | pigsty | 193.2 KiB | [postgresql-14-pg-ivm_1.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-ivm` | 1.12 | `d12.aarch64` | pigsty | 186.7 KiB | [postgresql-14-pg-ivm_1.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-ivm` | 1.12 | `u22.x86_64` | pigsty | 235.4 KiB | [postgresql-14-pg-ivm_1.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-ivm` | 1.12 | `u22.aarch64` | pigsty | 231.1 KiB | [postgresql-14-pg-ivm_1.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-ivm` | 1.12 | `u24.x86_64` | pigsty | 202.1 KiB | [postgresql-14-pg-ivm_1.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-ivm` | 1.12 | `u24.aarch64` | pigsty | 198.9 KiB | [postgresql-14-pg-ivm_1.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/sraoss/pg_ivm" title="Repository" icon="github" subtitle="github.com/sraoss/pg_ivm" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_ivm-1.13.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_ivm; # get pg_ivm source code
pig build dep pg_ivm; # install build dependencies
pig build pkg pg_ivm; # build extension rpm or deb
pig build ext pg_ivm; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_ivm; # install by extension name, for the current active PG version
pig ext install pg_ivm; # install via package alias, for the active PG version
pig ext install pg_ivm -v 18;   # install for PG 18
pig ext install pg_ivm -v 17;   # install for PG 17
pig ext install pg_ivm -v 16;   # install for PG 16
pig ext install pg_ivm -v 15;   # install for PG 15
pig ext install pg_ivm -v 14;   # install for PG 14
pig ext install pg_ivm -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_ivm CASCADE SCHEMA pg_catalog;
```

