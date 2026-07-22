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
| **9130** | {{< badge content="pg_statement_rollback" link="https://github.com/lzlabs/pg_statement_rollback" >}} | {{< ext "pg_statement_rollback" >}} | `1.6` | {{< category "SIM" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pg_dbms_errlog" >}} |
|   **See Also**    | {{< ext "oracle_fdw" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} {{< ext "session_variable" >}} {{< ext "safeupdate" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pg_hint_plan" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_statement_rollback` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6` | {{< bg "18" "pg_statement_rollback_18" "green" >}} {{< bg "17" "pg_statement_rollback_17" "green" >}} {{< bg "16" "pg_statement_rollback_16" "green" >}} {{< bg "15" "pg_statement_rollback_15" "green" >}} {{< bg "14" "pg_statement_rollback_14" "green" >}} | `pg_statement_rollback_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5` | {{< bg "18" "postgresql-18-pg-statement-rollback" "green" >}} {{< bg "17" "postgresql-17-pg-statement-rollback" "green" >}} {{< bg "16" "postgresql-16-pg-statement-rollback" "green" >}} {{< bg "15" "postgresql-15-pg-statement-rollback" "green" >}} {{< bg "14" "postgresql-14-pg-statement-rollback" "green" >}} | `postgresql-$v-pg-statement-rollback` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_14 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_14 : AVAIL 4" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6" "pg_statement_rollback_14 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-statement-rollback : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-statement-rollback : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statement_rollback_18` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.5 KiB | [pg_statement_rollback_18-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_statement_rollback_18-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.1 KiB | [pg_statement_rollback_18-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_statement_rollback_18-1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_statement_rollback_18` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pg_statement_rollback_18-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_statement_rollback_18-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_18-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_statement_rollback_18-1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_statement_rollback_18` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_18-1.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_statement_rollback_18-1.6-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.6 KiB | [pg_statement_rollback_18-1.5-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_statement_rollback_18-1.5-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.5 KiB | [pg_statement_rollback_18-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_statement_rollback_18-1.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_statement_rollback_18` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.6 KiB | [pg_statement_rollback_18-1.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_statement_rollback_18-1.6-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.4 KiB | [pg_statement_rollback_18-1.5-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_statement_rollback_18-1.5-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_18-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_statement_rollback_18-1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_statement_rollback_18` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.9 KiB | [pg_statement_rollback_18-1.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_statement_rollback_18-1.6-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_18-1.5-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_statement_rollback_18-1.5-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_18-1.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_statement_rollback_18-1.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_statement_rollback_18` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.9 KiB | [pg_statement_rollback_18-1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_statement_rollback_18-1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_18-1.5-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_statement_rollback_18-1.5-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_statement_rollback_18` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_18-1.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_statement_rollback_18-1.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.1 KiB | [postgresql-18-pg-statement-rollback_1.5-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.9 KiB | [postgresql-18-pg-statement-rollback_1.5-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-18-pg-statement-rollback_1.5-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.0 KiB | [postgresql-18-pg-statement-rollback_1.5-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 28.5 KiB | [postgresql-18-pg-statement-rollback_1.5-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 28.6 KiB | [postgresql-18-pg-statement-rollback_1.5-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.4 KiB | [postgresql-18-pg-statement-rollback_1.5-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.4 KiB | [postgresql-18-pg-statement-rollback_1.5-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-2PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 28.0 KiB | [postgresql-18-pg-statement-rollback_1.5-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-2PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-statement-rollback` | `1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 28.2 KiB | [postgresql-18-pg-statement-rollback_1.5-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statement-rollback/postgresql-18-pg-statement-rollback_1.5-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statement_rollback_17` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.5 KiB | [pg_statement_rollback_17-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_statement_rollback_17-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_statement_rollback_17` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.6 KiB | [pg_statement_rollback_17-1.4-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_statement_rollback_17-1.4-3PGDG.rhel8.x86_64.rpm) |
| `pg_statement_rollback_17` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pg_statement_rollback_17-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_statement_rollback_17-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_statement_rollback_17` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.5 KiB | [pg_statement_rollback_17-1.4-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_statement_rollback_17-1.4-3PGDG.rhel8.aarch64.rpm) |
| `pg_statement_rollback_17` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.9 KiB | [pg_statement_rollback_17-1.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_statement_rollback_17-1.6-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_statement_rollback_17` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.7 KiB | [pg_statement_rollback_17-1.5-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_statement_rollback_17-1.5-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_statement_rollback_17` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.5 KiB | [pg_statement_rollback_17-1.4-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_statement_rollback_17-1.4-3PGDG.rhel9.x86_64.rpm) |
| `pg_statement_rollback_17` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.7 KiB | [pg_statement_rollback_17-1.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_statement_rollback_17-1.6-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_statement_rollback_17` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.6 KiB | [pg_statement_rollback_17-1.5-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_statement_rollback_17-1.5-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_statement_rollback_17` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.4 KiB | [pg_statement_rollback_17-1.4-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_statement_rollback_17-1.4-3PGDG.rhel9.aarch64.rpm) |
| `pg_statement_rollback_17` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.9 KiB | [pg_statement_rollback_17-1.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_statement_rollback_17-1.6-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_statement_rollback_17` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_17-1.5-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_statement_rollback_17-1.5-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_statement_rollback_17` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_17-1.4-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_statement_rollback_17-1.4-4PGDG.rhel10.x86_64.rpm) |
| `pg_statement_rollback_17` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.0 KiB | [pg_statement_rollback_17-1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_statement_rollback_17-1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_statement_rollback_17` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.9 KiB | [pg_statement_rollback_17-1.5-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_statement_rollback_17-1.5-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_statement_rollback_17` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_17-1.4-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_statement_rollback_17-1.4-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.4 KiB | [postgresql-17-pg-statement-rollback_1.5-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.9 KiB | [postgresql-17-pg-statement-rollback_1.5-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.5 KiB | [postgresql-17-pg-statement-rollback_1.5-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.9 KiB | [postgresql-17-pg-statement-rollback_1.5-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.8 KiB | [postgresql-17-pg-statement-rollback_1.5-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.7 KiB | [postgresql-17-pg-statement-rollback_1.5-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.8 KiB | [postgresql-17-pg-statement-rollback_1.5-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.4 KiB | [postgresql-17-pg-statement-rollback_1.5-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-2PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 28.3 KiB | [postgresql-17-pg-statement-rollback_1.5-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-2PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-statement-rollback` | `1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 28.2 KiB | [postgresql-17-pg-statement-rollback_1.5-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statement-rollback/postgresql-17-pg-statement-rollback_1.5-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statement_rollback_16` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.5 KiB | [pg_statement_rollback_16-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_statement_rollback_16-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_statement_rollback_16` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.5 KiB | [pg_statement_rollback_16-1.4-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_statement_rollback_16-1.4-2PGDG.rhel8.x86_64.rpm) |
| `pg_statement_rollback_16` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pg_statement_rollback_16-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_statement_rollback_16-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_statement_rollback_16` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.4 KiB | [pg_statement_rollback_16-1.4-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_statement_rollback_16-1.4-2PGDG.rhel8.aarch64.rpm) |
| `pg_statement_rollback_16` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.9 KiB | [pg_statement_rollback_16-1.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_statement_rollback_16-1.6-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_statement_rollback_16` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.7 KiB | [pg_statement_rollback_16-1.5-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_statement_rollback_16-1.5-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_statement_rollback_16` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.3 KiB | [pg_statement_rollback_16-1.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_statement_rollback_16-1.4-2PGDG.rhel9.x86_64.rpm) |
| `pg_statement_rollback_16` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_16-1.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_statement_rollback_16-1.6-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_statement_rollback_16` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.6 KiB | [pg_statement_rollback_16-1.5-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_statement_rollback_16-1.5-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_statement_rollback_16` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_16-1.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_statement_rollback_16-1.4-2PGDG.rhel9.aarch64.rpm) |
| `pg_statement_rollback_16` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.0 KiB | [pg_statement_rollback_16-1.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_statement_rollback_16-1.6-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_statement_rollback_16` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_16-1.5-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_statement_rollback_16-1.5-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_statement_rollback_16` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_16-1.4-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_statement_rollback_16-1.4-4PGDG.rhel10.x86_64.rpm) |
| `pg_statement_rollback_16` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.0 KiB | [pg_statement_rollback_16-1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_statement_rollback_16-1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_statement_rollback_16` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.9 KiB | [pg_statement_rollback_16-1.5-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_statement_rollback_16-1.5-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_statement_rollback_16` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.9 KiB | [pg_statement_rollback_16-1.4-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_statement_rollback_16-1.4-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.4 KiB | [postgresql-16-pg-statement-rollback_1.5-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.9 KiB | [postgresql-16-pg-statement-rollback_1.5-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.4 KiB | [postgresql-16-pg-statement-rollback_1.5-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.9 KiB | [postgresql-16-pg-statement-rollback_1.5-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.8 KiB | [postgresql-16-pg-statement-rollback_1.5-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.7 KiB | [postgresql-16-pg-statement-rollback_1.5-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.8 KiB | [postgresql-16-pg-statement-rollback_1.5-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.4 KiB | [postgresql-16-pg-statement-rollback_1.5-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-2PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 28.3 KiB | [postgresql-16-pg-statement-rollback_1.5-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-2PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-statement-rollback` | `1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 28.3 KiB | [postgresql-16-pg-statement-rollback_1.5-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statement-rollback/postgresql-16-pg-statement-rollback_1.5-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statement_rollback_15` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.2 KiB | [pg_statement_rollback_15-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_statement_rollback_15-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.3 KiB | [pg_statement_rollback_15-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_statement_rollback_15-1.4-1.rhel8.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.4 KiB | [pg_statement_rollback_15-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_statement_rollback_15-1.3-1.rhel8.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.0 KiB | [pg_statement_rollback_15-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_statement_rollback_15-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_15-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_statement_rollback_15-1.4-1.rhel8.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.2 KiB | [pg_statement_rollback_15-1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_statement_rollback_15-1.3-1.rhel8.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.7 KiB | [pg_statement_rollback_15-1.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_statement_rollback_15-1.6-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.7 KiB | [pg_statement_rollback_15-1.5-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_statement_rollback_15-1.5-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.3 KiB | [pg_statement_rollback_15-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_statement_rollback_15-1.4-1.rhel9.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.2 KiB | [pg_statement_rollback_15-1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_statement_rollback_15-1.3-1.rhel9.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.5 KiB | [pg_statement_rollback_15-1.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_statement_rollback_15-1.6-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.6 KiB | [pg_statement_rollback_15-1.5-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_statement_rollback_15-1.5-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_15-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_statement_rollback_15-1.4-1.rhel9.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.9 KiB | [pg_statement_rollback_15-1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_statement_rollback_15-1.3-1.rhel9.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_15-1.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_statement_rollback_15-1.6-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_15-1.5-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_statement_rollback_15-1.5-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_15-1.4-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_statement_rollback_15-1.4-4PGDG.rhel10.x86_64.rpm) |
| `pg_statement_rollback_15` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_15-1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_statement_rollback_15-1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_15-1.5-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_statement_rollback_15-1.5-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_statement_rollback_15` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_15-1.4-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_statement_rollback_15-1.4-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.9 KiB | [postgresql-15-pg-statement-rollback_1.5-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.3 KiB | [postgresql-15-pg-statement-rollback_1.5-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.9 KiB | [postgresql-15-pg-statement-rollback_1.5-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.4 KiB | [postgresql-15-pg-statement-rollback_1.5-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.1 KiB | [postgresql-15-pg-statement-rollback_1.5-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.9 KiB | [postgresql-15-pg-statement-rollback_1.5-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.3 KiB | [postgresql-15-pg-statement-rollback_1.5-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.9 KiB | [postgresql-15-pg-statement-rollback_1.5-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-2PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 28.0 KiB | [postgresql-15-pg-statement-rollback_1.5-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-2PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-statement-rollback` | `1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.5 KiB | [postgresql-15-pg-statement-rollback_1.5-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statement-rollback/postgresql-15-pg-statement-rollback_1.5-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statement_rollback_14` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.2 KiB | [pg_statement_rollback_14-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_statement_rollback_14-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.3 KiB | [pg_statement_rollback_14-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_statement_rollback_14-1.4-1.rhel8.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.0 KiB | [pg_statement_rollback_14-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_statement_rollback_14-1.3-1.rhel8.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.0 KiB | [pg_statement_rollback_14-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_statement_rollback_14-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_14-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_statement_rollback_14-1.4-1.rhel8.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.1 KiB | [pg_statement_rollback_14-1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_statement_rollback_14-1.3-1.rhel8.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.7 KiB | [pg_statement_rollback_14-1.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_statement_rollback_14-1.6-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.7 KiB | [pg_statement_rollback_14-1.5-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_statement_rollback_14-1.5-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.3 KiB | [pg_statement_rollback_14-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_statement_rollback_14-1.4-1.rhel9.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.5 KiB | [pg_statement_rollback_14-1.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_statement_rollback_14-1.6-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.5 KiB | [pg_statement_rollback_14-1.5-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_statement_rollback_14-1.5-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.1 KiB | [pg_statement_rollback_14-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_statement_rollback_14-1.4-1.rhel9.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.8 KiB | [pg_statement_rollback_14-1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_statement_rollback_14-1.3-1.rhel9.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_14-1.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_statement_rollback_14-1.6-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_14-1.5-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_statement_rollback_14-1.5-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.8 KiB | [pg_statement_rollback_14-1.4-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_statement_rollback_14-1.4-4PGDG.rhel10.x86_64.rpm) |
| `pg_statement_rollback_14` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_14-1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_statement_rollback_14-1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_14-1.5-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_statement_rollback_14-1.5-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_statement_rollback_14` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.8 KiB | [pg_statement_rollback_14-1.4-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_statement_rollback_14-1.4-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.9 KiB | [postgresql-14-pg-statement-rollback_1.5-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.3 KiB | [postgresql-14-pg-statement-rollback_1.5-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.9 KiB | [postgresql-14-pg-statement-rollback_1.5-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.4 KiB | [postgresql-14-pg-statement-rollback_1.5-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.0 KiB | [postgresql-14-pg-statement-rollback_1.5-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.9 KiB | [postgresql-14-pg-statement-rollback_1.5-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.3 KiB | [postgresql-14-pg-statement-rollback_1.5-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.9 KiB | [postgresql-14-pg-statement-rollback_1.5-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-2PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 28.0 KiB | [postgresql-14-pg-statement-rollback_1.5-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-2PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-statement-rollback` | `1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.6 KiB | [postgresql-14-pg-statement-rollback_1.5-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statement-rollback/postgresql-14-pg-statement-rollback_1.5-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/lzlabs/pg_statement_rollback" title="Repository" icon="github" subtitle="github.com/lzlabs/pg_statement_rollback" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_statement_rollback-1.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_statement_rollback;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_statement_rollback;		# install via package name, for the active PG version

pig install pg_statement_rollback -v 18;   # install for PG 18
pig install pg_statement_rollback -v 17;   # install for PG 17
pig install pg_statement_rollback -v 16;   # install for PG 16
pig install pg_statement_rollback -v 15;   # install for PG 15
pig install pg_statement_rollback -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_statement_rollback';
```


This extension does not need `CREATE EXTENSION` DDL command



## Usage

Sources:

- [pg_statement_rollback v1.6 README](https://github.com/lzlabs/pg_statement_rollback/blob/v1.6/README.md)
- [Changes from v1.5 to v1.6](https://github.com/lzlabs/pg_statement_rollback/compare/v1.5...v1.6)

pg_statement_rollback keeps an explicit transaction usable after a statement error by creating an automatic savepoint before each eligible statement. It emulates the statement-level rollback behavior familiar from some other databases, but the client must still issue ROLLBACK TO SAVEPOINT after an error.

The module is loaded into a backend and does not require CREATE EXTENSION.

### Load the Module

Load it for one session:

    LOAD 'pg_statement_rollback.so';

For a selected role or database, add it to session_preload_libraries and reconnect:

    session_preload_libraries = 'pg_statement_rollback'

Use shared_preload_libraries only if the deployment specifically needs server-wide loading; changing either preload list at server scope requires the corresponding restart or reconnect boundary.

### Recover from a Failed Statement

    BEGIN;
    INSERT INTO accounts(id, balance) VALUES (1, 100);
    INSERT INTO accounts(id, balance) VALUES (1, 200);
    -- duplicate-key error
    ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt";
    UPDATE accounts SET balance = 150 WHERE id = 1;
    COMMIT;

The savepoint name is case-sensitive when quoted. Applications must detect the statement error and send the rollback command before continuing.

### Configuration Index

- pg_statement_rollback.enabled enables automatic savepoints for the current session.
- pg_statement_rollback.savepoint_name changes the automatic savepoint name and is superuser-controlled.
- pg_statement_rollback.enable_writeonly limits savepoint creation to statements that can write.

### Version 1.6 Behavior

Version 1.6 adds PostgreSQL 19 build support and improves detection of read-only transactions. The module no longer creates automatic savepoints in read-only transactions and releases its initial savepoint before SET TRANSACTION ... READ ONLY, which avoids interfering with dump and other read-only sessions.

### Caveats

- This is not transparent retry logic: clients must explicitly roll back to the automatic savepoint.
- Savepoints add overhead to every covered statement. Measure write-heavy workloads before enabling the module broadly.
- The upstream README warns of a crash with assertion-enabled PostgreSQL builds; do not treat development-build behavior as production-safe without testing.
- Transaction-wide errors, connection failures, and errors that invalidate the session cannot be repaired by a savepoint.
