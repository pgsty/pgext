---
title: "pg_statement_rollback"
linkTitle: "pg_statement_rollback"
description: "Server side rollback at statement level for PostgreSQL like Oracle or DB2"
weight: 9130
categories: ["SIM"]
width: full
---

[**pg_statement_rollback**](https://github.com/lzlabs/pg_statement_rollback) : Server side rollback at statement level for PostgreSQL like Oracle or DB2


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9130** | {{< badge content="pg_statement_rollback" link="https://github.com/lzlabs/pg_statement_rollback" >}} | {{< ext "pg_statement_rollback" >}} | `1.5` | {{< category "SIM" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "oracle_fdw" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} {{< ext "session_variable" >}} {{< ext "safeupdate" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pg_hint_plan" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_statement_rollback` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5` | {{< bg "18" "pg_statement_rollback_18*" "green" >}} {{< bg "17" "pg_statement_rollback_17*" "green" >}} {{< bg "16" "pg_statement_rollback_16*" "green" >}} {{< bg "15" "pg_statement_rollback_15*" "green" >}} {{< bg "14" "pg_statement_rollback_14*" "green" >}} {{< bg "13" "pg_statement_rollback_13*" "green" >}} | `pg_statement_rollback_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5` | {{< bg "18" "postgresql-18-pg-statement-rollback" "green" >}} {{< bg "17" "postgresql-17-pg-statement-rollback" "green" >}} {{< bg "16" "postgresql-16-pg-statement-rollback" "green" >}} {{< bg "15" "postgresql-15-pg-statement-rollback" "green" >}} {{< bg "14" "postgresql-14-pg-statement-rollback" "green" >}} {{< bg "13" "postgresql-13-pg-statement-rollback" "green" >}} | `postgresql-$v-pg-statement-rollback` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.5" "pg_statement_rollback_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_13 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.5" "pg_statement_rollback_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.5" "pg_statement_rollback_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.5" "pg_statement_rollback_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.5" "pg_statement_rollback_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.5" "pg_statement_rollback_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_statement_rollback_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-13-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-13-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-13-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-13-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-13-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-13-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-13-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-13-pg-statement-rollback : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statement_rollback_18` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.1 KiB | [pg_statement_rollback_18-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_statement_rollback_18-1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_18-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_statement_rollback_18-1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.5 KiB | [pg_statement_rollback_18-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_statement_rollback_18-1.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_18-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_statement_rollback_18-1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_18-1.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_statement_rollback_18-1.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_18-1.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10.0-aarch64/pg_statement_rollback_18-1.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.1 KiB | [postgresql-18-pg-statement-rollback_1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.9 KiB | [postgresql-18-pg-statement-rollback_1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-18-pg-statement-rollback_1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.0 KiB | [postgresql-18-pg-statement-rollback_1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 28.5 KiB | [postgresql-18-pg-statement-rollback_1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 28.6 KiB | [postgresql-18-pg-statement-rollback_1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.4 KiB | [postgresql-18-pg-statement-rollback_1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.4 KiB | [postgresql-18-pg-statement-rollback_1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statement_rollback_17` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.6 KiB | [pg_statement_rollback_17-1.4-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_statement_rollback_17-1.4-3PGDG.rhel8.x86_64.rpm) |
| `pg_statement_rollback_17` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.5 KiB | [pg_statement_rollback_17-1.4-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_statement_rollback_17-1.4-3PGDG.rhel8.aarch64.rpm) |
| `pg_statement_rollback_17` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.5 KiB | [pg_statement_rollback_17-1.4-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_statement_rollback_17-1.4-3PGDG.rhel9.x86_64.rpm) |
| `pg_statement_rollback_17` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.4 KiB | [pg_statement_rollback_17-1.4-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_statement_rollback_17-1.4-3PGDG.rhel9.aarch64.rpm) |
| `pg_statement_rollback_17` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_17-1.4-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_statement_rollback_17-1.4-4PGDG.rhel10.x86_64.rpm) |
| `pg_statement_rollback_17` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_17-1.4-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10.0-aarch64/pg_statement_rollback_17-1.4-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.4 KiB | [postgresql-17-pg-statement-rollback_1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.9 KiB | [postgresql-17-pg-statement-rollback_1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.5 KiB | [postgresql-17-pg-statement-rollback_1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.0 KiB | [postgresql-17-pg-statement-rollback_1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.9 KiB | [postgresql-17-pg-statement-rollback_1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.7 KiB | [postgresql-17-pg-statement-rollback_1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.8 KiB | [postgresql-17-pg-statement-rollback_1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.5 KiB | [postgresql-17-pg-statement-rollback_1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statement_rollback_16` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.5 KiB | [pg_statement_rollback_16-1.4-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_statement_rollback_16-1.4-2PGDG.rhel8.x86_64.rpm) |
| `pg_statement_rollback_16` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.4 KiB | [pg_statement_rollback_16-1.4-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_statement_rollback_16-1.4-2PGDG.rhel8.aarch64.rpm) |
| `pg_statement_rollback_16` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.3 KiB | [pg_statement_rollback_16-1.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_statement_rollback_16-1.4-2PGDG.rhel9.x86_64.rpm) |
| `pg_statement_rollback_16` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_16-1.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_statement_rollback_16-1.4-2PGDG.rhel9.aarch64.rpm) |
| `pg_statement_rollback_16` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_16-1.4-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_statement_rollback_16-1.4-4PGDG.rhel10.x86_64.rpm) |
| `pg_statement_rollback_16` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.9 KiB | [pg_statement_rollback_16-1.4-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10.0-aarch64/pg_statement_rollback_16-1.4-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.4 KiB | [postgresql-16-pg-statement-rollback_1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.9 KiB | [postgresql-16-pg-statement-rollback_1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.4 KiB | [postgresql-16-pg-statement-rollback_1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.0 KiB | [postgresql-16-pg-statement-rollback_1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.8 KiB | [postgresql-16-pg-statement-rollback_1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.7 KiB | [postgresql-16-pg-statement-rollback_1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.8 KiB | [postgresql-16-pg-statement-rollback_1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.5 KiB | [postgresql-16-pg-statement-rollback_1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statement_rollback_15` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.3 KiB | [pg_statement_rollback_15-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_statement_rollback_15-1.4-1.rhel8.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.4 KiB | [pg_statement_rollback_15-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_statement_rollback_15-1.3-1.rhel8.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_15-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_statement_rollback_15-1.4-1.rhel8.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.2 KiB | [pg_statement_rollback_15-1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_statement_rollback_15-1.3-1.rhel8.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.3 KiB | [pg_statement_rollback_15-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_statement_rollback_15-1.4-1.rhel9.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.2 KiB | [pg_statement_rollback_15-1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_statement_rollback_15-1.3-1.rhel9.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_15-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_statement_rollback_15-1.4-1.rhel9.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.9 KiB | [pg_statement_rollback_15-1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_statement_rollback_15-1.3-1.rhel9.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_15-1.4-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_statement_rollback_15-1.4-4PGDG.rhel10.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_15-1.4-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10.0-aarch64/pg_statement_rollback_15-1.4-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.9 KiB | [postgresql-15-pg-statement-rollback_1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.3 KiB | [postgresql-15-pg-statement-rollback_1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.9 KiB | [postgresql-15-pg-statement-rollback_1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.3 KiB | [postgresql-15-pg-statement-rollback_1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.1 KiB | [postgresql-15-pg-statement-rollback_1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.0 KiB | [postgresql-15-pg-statement-rollback_1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.3 KiB | [postgresql-15-pg-statement-rollback_1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.0 KiB | [postgresql-15-pg-statement-rollback_1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statement_rollback_14` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.3 KiB | [pg_statement_rollback_14-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_statement_rollback_14-1.4-1.rhel8.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.0 KiB | [pg_statement_rollback_14-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_statement_rollback_14-1.3-1.rhel8.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_14-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_statement_rollback_14-1.4-1.rhel8.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.1 KiB | [pg_statement_rollback_14-1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_statement_rollback_14-1.3-1.rhel8.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.3 KiB | [pg_statement_rollback_14-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_statement_rollback_14-1.4-1.rhel9.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_14-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_statement_rollback_14-1.4-1.rhel9.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.8 KiB | [pg_statement_rollback_14-1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_statement_rollback_14-1.3-1.rhel9.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_14-1.4-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_statement_rollback_14-1.4-4PGDG.rhel10.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_14-1.4-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10.0-aarch64/pg_statement_rollback_14-1.4-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.9 KiB | [postgresql-14-pg-statement-rollback_1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.3 KiB | [postgresql-14-pg-statement-rollback_1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.9 KiB | [postgresql-14-pg-statement-rollback_1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.4 KiB | [postgresql-14-pg-statement-rollback_1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.1 KiB | [postgresql-14-pg-statement-rollback_1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.9 KiB | [postgresql-14-pg-statement-rollback_1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.3 KiB | [postgresql-14-pg-statement-rollback_1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.0 KiB | [postgresql-14-pg-statement-rollback_1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statement_rollback_13` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.1 KiB | [pg_statement_rollback_13-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_statement_rollback_13-1.4-1.rhel8.x86_64.rpm) |
| `pg_statement_rollback_13` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.8 KiB | [pg_statement_rollback_13-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_statement_rollback_13-1.3-1.rhel8.x86_64.rpm) |
| `pg_statement_rollback_13` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.8 KiB | [pg_statement_rollback_13-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_statement_rollback_13-1.2-1.rhel8.x86_64.rpm) |
| `pg_statement_rollback_13` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_13-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_statement_rollback_13-1.4-1.rhel8.aarch64.rpm) |
| `pg_statement_rollback_13` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.0 KiB | [pg_statement_rollback_13-1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_statement_rollback_13-1.3-1.rhel8.aarch64.rpm) |
| `pg_statement_rollback_13` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.3 KiB | [pg_statement_rollback_13-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_statement_rollback_13-1.4-1.rhel9.x86_64.rpm) |
| `pg_statement_rollback_13` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.0 KiB | [pg_statement_rollback_13-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_statement_rollback_13-1.4-1.rhel9.aarch64.rpm) |
| `pg_statement_rollback_13` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.7 KiB | [pg_statement_rollback_13-1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_statement_rollback_13-1.3-1.rhel9.aarch64.rpm) |
| `pg_statement_rollback_13` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.7 KiB | [pg_statement_rollback_13-1.4-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_statement_rollback_13-1.4-4PGDG.rhel10.x86_64.rpm) |
| `pg_statement_rollback_13` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_13-1.4-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10.0-aarch64/pg_statement_rollback_13-1.4-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-statement-rollback` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.6 KiB | [postgresql-13-pg-statement-rollback_1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-13-pg-statement-rollback_1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-statement-rollback` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.2 KiB | [postgresql-13-pg-statement-rollback_1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-13-pg-statement-rollback_1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-statement-rollback` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.6 KiB | [postgresql-13-pg-statement-rollback_1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-13-pg-statement-rollback_1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-statement-rollback` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.3 KiB | [postgresql-13-pg-statement-rollback_1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-13-pg-statement-rollback_1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-statement-rollback` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 30.8 KiB | [postgresql-13-pg-statement-rollback_1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-13-pg-statement-rollback_1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-statement-rollback` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.5 KiB | [postgresql-13-pg-statement-rollback_1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-13-pg-statement-rollback_1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-statement-rollback` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.0 KiB | [postgresql-13-pg-statement-rollback_1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-13-pg-statement-rollback_1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-statement-rollback` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.9 KiB | [postgresql-13-pg-statement-rollback_1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-13-pg-statement-rollback_1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/lzlabs/pg_statement_rollback" title="Repository" icon="github" subtitle="github.com/lzlabs/pg_statement_rollback" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_statement_rollback-1.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_statement_rollback;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_statement_rollback;		# install via package name, for the active PG version

pig install pg_statement_rollback -v 18;   # install for PG 18
pig install pg_statement_rollback -v 17;   # install for PG 17
pig install pg_statement_rollback -v 16;   # install for PG 16
pig install pg_statement_rollback -v 15;   # install for PG 15
pig install pg_statement_rollback -v 14;   # install for PG 14
pig install pg_statement_rollback -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_statement_rollback';
```


This extension does not need `CREATE EXTENSION` DDL command


