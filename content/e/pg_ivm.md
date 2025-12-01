---
title: "pg_ivm"
linkTitle: "pg_ivm"
description: "incremental view maintenance on PostgreSQL"
weight: 2870
categories: ["FEAT"]
width: full
---

[**pg_ivm**](https://github.com/sraoss/pg_ivm) : incremental view maintenance on PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2870** | {{< badge content="pg_ivm" link="https://github.com/sraoss/pg_ivm" >}} | {{< ext "pg_ivm" >}} | `1.13` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "timeseries" >}} |
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.13` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_ivm` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.13` | {{< bg "18" "pg_ivm_18*" "green" >}} {{< bg "17" "pg_ivm_17*" "green" >}} {{< bg "16" "pg_ivm_16*" "green" >}} {{< bg "15" "pg_ivm_15*" "green" >}} {{< bg "14" "pg_ivm_14*" "green" >}} {{< bg "13" "pg_ivm_13*" "green" >}} | `pg_ivm_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.13` | {{< bg "18" "postgresql-18-pg-ivm" "green" >}} {{< bg "17" "postgresql-17-pg-ivm" "green" >}} {{< bg "16" "postgresql-16-pg-ivm" "green" >}} {{< bg "15" "postgresql-15-pg-ivm" "green" >}} {{< bg "14" "postgresql-14-pg-ivm" "green" >}} {{< bg "13" "postgresql-13-pg-ivm" "green" >}} | `postgresql-$v-pg-ivm` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.13" "pg_ivm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_14 : AVAIL 14" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_13 : AVAIL 12" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.13" "pg_ivm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_13 : AVAIL 10" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.13" "pg_ivm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_14 : AVAIL 13" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_13 : AVAIL 12" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.13" "pg_ivm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_13 : AVAIL 10" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.13" "pg_ivm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.13" "pg_ivm_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.13" "pg_ivm_13 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.13" "postgresql-18-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-13-pg-ivm : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.13" "postgresql-18-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-13-pg-ivm : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.13" "postgresql-18-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-13-pg-ivm : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.13" "postgresql-18-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-13-pg-ivm : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.13" "postgresql-18-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-13-pg-ivm : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.13" "postgresql-18-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-13-pg-ivm : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.13" "postgresql-18-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-13-pg-ivm : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.13" "postgresql-18-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-17-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-16-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-15-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-14-pg-ivm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.13" "postgresql-13-pg-ivm : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_18` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [pg_ivm_18-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_ivm_18-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_18` | `1.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.3 KiB | [pg_ivm_18-1.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_ivm_18-1.12-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_18` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.5 KiB | [pg_ivm_18-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_ivm_18-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_18` | `1.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.2 KiB | [pg_ivm_18-1.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_ivm_18-1.12-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_18` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [pg_ivm_18-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_18` | `1.12` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.3 KiB | [pg_ivm_18-1.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.12-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_18` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.1 KiB | [pg_ivm_18-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_18` | `1.12` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.0 KiB | [pg_ivm_18-1.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.12-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_18` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.6 KiB | [pg_ivm_18-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_ivm_18-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_18` | `1.12` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.1 KiB | [pg_ivm_18-1.12-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_ivm_18-1.12-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_18` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_18-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_ivm_18-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_18` | `1.12` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pg_ivm_18-1.12-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_ivm_18-1.12-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 118.5 KiB | [postgresql-18-pg-ivm_1.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 115.2 KiB | [postgresql-18-pg-ivm_1.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 118.3 KiB | [postgresql-18-pg-ivm_1.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 114.9 KiB | [postgresql-18-pg-ivm_1.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 129.6 KiB | [postgresql-18-pg-ivm_1.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 127.3 KiB | [postgresql-18-pg-ivm_1.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 123.6 KiB | [postgresql-18-pg-ivm_1.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 121.8 KiB | [postgresql-18-pg-ivm_1.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_17` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.4 KiB | [pg_ivm_17-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | `1.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.9 KiB | [pg_ivm_17-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.6 KiB | [pg_ivm_17-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | `1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.6 KiB | [pg_ivm_17-1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.9-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.3 KiB | [pg_ivm_17-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | `1.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.9 KiB | [pg_ivm_17-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.5 KiB | [pg_ivm_17-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | `1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [pg_ivm_17-1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.9-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [pg_ivm_17-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | `1.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.2 KiB | [pg_ivm_17-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.9 KiB | [pg_ivm_17-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | `1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.4 KiB | [pg_ivm_17-1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.9-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.0 KiB | [pg_ivm_17-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | `1.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.9 KiB | [pg_ivm_17-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [pg_ivm_17-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.10-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | `1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.9 KiB | [pg_ivm_17-1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.9-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.7 KiB | [pg_ivm_17-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_17` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.2 KiB | [pg_ivm_17-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_17` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.9 KiB | [pg_ivm_17-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_17` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.6 KiB | [pg_ivm_17-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_17` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.7 KiB | [pg_ivm_17-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_17` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pg_ivm_17-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 118.2 KiB | [postgresql-17-pg-ivm_1.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 114.9 KiB | [postgresql-17-pg-ivm_1.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 118.0 KiB | [postgresql-17-pg-ivm_1.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 114.5 KiB | [postgresql-17-pg-ivm_1.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 149.7 KiB | [postgresql-17-pg-ivm_1.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 147.6 KiB | [postgresql-17-pg-ivm_1.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 123.0 KiB | [postgresql-17-pg-ivm_1.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 121.5 KiB | [postgresql-17-pg-ivm_1.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_16` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [pg_ivm_16-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.0 KiB | [pg_ivm_16-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.7 KiB | [pg_ivm_16-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.9 KiB | [pg_ivm_16-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [pg_ivm_16-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.4 KiB | [pg_ivm_16-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.9 KiB | [pg_ivm_16-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.6 KiB | [pg_ivm_16-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.9 KiB | [pg_ivm_16-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.7 KiB | [pg_ivm_16-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [pg_ivm_16-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.2 KiB | [pg_ivm_16-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.0 KiB | [pg_ivm_16-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.7 KiB | [pg_ivm_16-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.6 KiB | [pg_ivm_16-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.2 KiB | [pg_ivm_16-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.0 KiB | [pg_ivm_16-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.7 KiB | [pg_ivm_16-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.10-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.3 KiB | [pg_ivm_16-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.4 KiB | [pg_ivm_16-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.7 KiB | [pg_ivm_16-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_16` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [pg_ivm_16-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_16` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.9 KiB | [pg_ivm_16-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_16` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_16-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_16` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.7 KiB | [pg_ivm_16-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_16` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pg_ivm_16-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 118.0 KiB | [postgresql-16-pg-ivm_1.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 114.9 KiB | [postgresql-16-pg-ivm_1.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 117.8 KiB | [postgresql-16-pg-ivm_1.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 114.6 KiB | [postgresql-16-pg-ivm_1.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 148.6 KiB | [postgresql-16-pg-ivm_1.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 146.4 KiB | [postgresql-16-pg-ivm_1.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 123.1 KiB | [postgresql-16-pg-ivm_1.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 121.7 KiB | [postgresql-16-pg-ivm_1.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
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
| `pg_ivm_15` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.5 KiB | [pg_ivm_15-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_15` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.0 KiB | [pg_ivm_15-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_15` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.7 KiB | [pg_ivm_15-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_15` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.4 KiB | [pg_ivm_15-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_15` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.5 KiB | [pg_ivm_15-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_15` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.1 KiB | [pg_ivm_15-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 118.2 KiB | [postgresql-15-pg-ivm_1.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 114.8 KiB | [postgresql-15-pg-ivm_1.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 118.0 KiB | [postgresql-15-pg-ivm_1.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 114.8 KiB | [postgresql-15-pg-ivm_1.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 148.7 KiB | [postgresql-15-pg-ivm_1.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 146.3 KiB | [postgresql-15-pg-ivm_1.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 123.2 KiB | [postgresql-15-pg-ivm_1.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 121.8 KiB | [postgresql-15-pg-ivm_1.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
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
| `pg_ivm_14` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 80.8 KiB | [pg_ivm_14-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_14` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 74.9 KiB | [pg_ivm_14-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_14` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 74.6 KiB | [pg_ivm_14-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_14` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 79.1 KiB | [pg_ivm_14-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_14` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 72.4 KiB | [pg_ivm_14-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_14` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 72.1 KiB | [pg_ivm_14-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 208.8 KiB | [postgresql-14-pg-ivm_1.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 201.5 KiB | [postgresql-14-pg-ivm_1.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 208.3 KiB | [postgresql-14-pg-ivm_1.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 201.9 KiB | [postgresql-14-pg-ivm_1.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 252.8 KiB | [postgresql-14-pg-ivm_1.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 248.6 KiB | [postgresql-14-pg-ivm_1.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 217.4 KiB | [postgresql-14-pg-ivm_1.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 214.7 KiB | [postgresql-14-pg-ivm_1.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_13` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.7 KiB | [pg_ivm_13-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_ivm_13-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_13` | `1.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.7 KiB | [pg_ivm_13-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_ivm_13-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_13` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.4 KiB | [pg_ivm_13-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_ivm_13-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_13` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.5 KiB | [pg_ivm_13-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_ivm_13-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_13` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.4 KiB | [pg_ivm_13-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_ivm_13-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_13` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.1 KiB | [pg_ivm_13-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_ivm_13-1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_13` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.9 KiB | [pg_ivm_13-1.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_ivm_13-1.5.1-1.rhel8.x86_64.rpm) |
| `pg_ivm_13` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.2 KiB | [pg_ivm_13-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_ivm_13-1.5-1.rhel8.x86_64.rpm) |
| `pg_ivm_13` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.1 KiB | [pg_ivm_13-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_ivm_13-1.4-1.rhel8.x86_64.rpm) |
| `pg_ivm_13` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 66.5 KiB | [pg_ivm_13-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_ivm_13-1.3-1.rhel8.x86_64.rpm) |
| `pg_ivm_13` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 65.3 KiB | [pg_ivm_13-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_ivm_13-1.2-1.rhel8.x86_64.rpm) |
| `pg_ivm_13` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.6 KiB | [pg_ivm_13-1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_ivm_13-1.1-1.rhel8.x86_64.rpm) |
| `pg_ivm_13` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 72.6 KiB | [pg_ivm_13-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_ivm_13-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_13` | `1.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.7 KiB | [pg_ivm_13-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_ivm_13-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_13` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.4 KiB | [pg_ivm_13-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_ivm_13-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_13` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.5 KiB | [pg_ivm_13-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_ivm_13-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_13` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.5 KiB | [pg_ivm_13-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_ivm_13-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_13` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.4 KiB | [pg_ivm_13-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_ivm_13-1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_13` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.3 KiB | [pg_ivm_13-1.5.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_ivm_13-1.5.1-1.rhel8.aarch64.rpm) |
| `pg_ivm_13` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.4 KiB | [pg_ivm_13-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_ivm_13-1.5-1.rhel8.aarch64.rpm) |
| `pg_ivm_13` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.3 KiB | [pg_ivm_13-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_ivm_13-1.4-1.rhel8.aarch64.rpm) |
| `pg_ivm_13` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 62.8 KiB | [pg_ivm_13-1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_ivm_13-1.3-1.rhel8.aarch64.rpm) |
| `pg_ivm_13` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 79.5 KiB | [pg_ivm_13-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_ivm_13-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_13` | `1.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.5 KiB | [pg_ivm_13-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_ivm_13-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_13` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.3 KiB | [pg_ivm_13-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_ivm_13-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_13` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 70.9 KiB | [pg_ivm_13-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_ivm_13-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_13` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.5 KiB | [pg_ivm_13-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_ivm_13-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_13` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.3 KiB | [pg_ivm_13-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_ivm_13-1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_13` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.0 KiB | [pg_ivm_13-1.5.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_ivm_13-1.5.1-1.rhel9.x86_64.rpm) |
| `pg_ivm_13` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.0 KiB | [pg_ivm_13-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_ivm_13-1.5-1.rhel9.x86_64.rpm) |
| `pg_ivm_13` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 71.1 KiB | [pg_ivm_13-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_ivm_13-1.4-1.rhel9.x86_64.rpm) |
| `pg_ivm_13` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 70.8 KiB | [pg_ivm_13-1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_ivm_13-1.3-1.rhel9.x86_64.rpm) |
| `pg_ivm_13` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.5 KiB | [pg_ivm_13-1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_ivm_13-1.2-1.rhel9.x86_64.rpm) |
| `pg_ivm_13` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.9 KiB | [pg_ivm_13-1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_ivm_13-1.1-1.rhel9.x86_64.rpm) |
| `pg_ivm_13` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 76.1 KiB | [pg_ivm_13-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_ivm_13-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_13` | `1.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 70.2 KiB | [pg_ivm_13-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_ivm_13-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_13` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.9 KiB | [pg_ivm_13-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_ivm_13-1.10-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_13` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.7 KiB | [pg_ivm_13-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_ivm_13-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_13` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.0 KiB | [pg_ivm_13-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_ivm_13-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_13` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 70.9 KiB | [pg_ivm_13-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_ivm_13-1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_13` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.6 KiB | [pg_ivm_13-1.5.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_ivm_13-1.5.1-1.rhel9.aarch64.rpm) |
| `pg_ivm_13` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.8 KiB | [pg_ivm_13-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_ivm_13-1.5-1.rhel9.aarch64.rpm) |
| `pg_ivm_13` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.8 KiB | [pg_ivm_13-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_ivm_13-1.4-1.rhel9.aarch64.rpm) |
| `pg_ivm_13` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.7 KiB | [pg_ivm_13-1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_ivm_13-1.3-1.rhel9.aarch64.rpm) |
| `pg_ivm_13` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 81.5 KiB | [pg_ivm_13-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_ivm_13-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_13` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 74.9 KiB | [pg_ivm_13-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_ivm_13-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_13` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 74.9 KiB | [pg_ivm_13-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_ivm_13-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_13` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 78.7 KiB | [pg_ivm_13-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_ivm_13-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_13` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 72.1 KiB | [pg_ivm_13-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_ivm_13-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_13` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 71.8 KiB | [pg_ivm_13-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_ivm_13-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 207.2 KiB | [postgresql-13-pg-ivm_1.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-13-pg-ivm_1.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 200.0 KiB | [postgresql-13-pg-ivm_1.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-13-pg-ivm_1.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 206.8 KiB | [postgresql-13-pg-ivm_1.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-13-pg-ivm_1.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 200.0 KiB | [postgresql-13-pg-ivm_1.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-13-pg-ivm_1.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 250.8 KiB | [postgresql-13-pg-ivm_1.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-13-pg-ivm_1.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 245.6 KiB | [postgresql-13-pg-ivm_1.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-13-pg-ivm_1.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 215.6 KiB | [postgresql-13-pg-ivm_1.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-13-pg-ivm_1.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 212.3 KiB | [postgresql-13-pg-ivm_1.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-13-pg-ivm_1.13-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/sraoss/pg_ivm" title="Repository" icon="github" subtitle="github.com/sraoss/pg_ivm" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_ivm-1.13.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_ivm;		# build rpm / deb with pig
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
pig install pg_ivm -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_ivm;
```
