---
title: "pg_dbms_errlog"
linkTitle: "pg_dbms_errlog"
description: "Emulate DBMS_ERRLOG Oracle module to log DML errors in a dedicated table."
weight: 9270
categories: ["SIM"]
width: full
---

[**pg_dbms_errlog**](https://github.com/HexaCluster/pg_dbms_errlog) : Emulate DBMS_ERRLOG Oracle module to log DML errors in a dedicated table.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9270** | {{< badge content="pg_dbms_errlog" link="https://github.com/HexaCluster/pg_dbms_errlog" >}} | {{< ext "pg_dbms_errlog" >}} | `2.4` | {{< category "SIM" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `dbms_errlog` |
|   **Requires**    | {{< ext "pg_statement_rollback" >}} |
|   **See Also**    | {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pg_dbms_job" >}} |

> [!Note] Requires pg_statement_rollback and shared_preload_libraries=pg_dbms_errlog; restart required.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `2.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_dbms_errlog` | `pg_statement_rollback` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.2` | {{< bg "18" "pg_dbms_errlog_18" "green" >}} {{< bg "17" "pg_dbms_errlog_17" "green" >}} {{< bg "16" "pg_dbms_errlog_16" "green" >}} {{< bg "15" "pg_dbms_errlog_15" "green" >}} {{< bg "14" "pg_dbms_errlog_14" "green" >}} | `pg_dbms_errlog_$v` | `pg_statement_rollback_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4` | {{< bg "18" "postgresql-18-pg-dbms-errlog" "green" >}} {{< bg "17" "postgresql-17-pg-dbms-errlog" "green" >}} {{< bg "16" "postgresql-16-pg-dbms-errlog" "green" >}} {{< bg "15" "postgresql-15-pg-dbms-errlog" "green" >}} {{< bg "14" "postgresql-14-pg-dbms-errlog" "green" >}} | `postgresql-$v-pg-dbms-errlog` | `postgresql-$v-pg-statement-rollback` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.4" "postgresql-18-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-17-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-16-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-15-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-14-pg-dbms-errlog : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.4" "postgresql-18-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-17-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-16-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-15-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-14-pg-dbms-errlog : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.4" "postgresql-18-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-17-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-16-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-15-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-14-pg-dbms-errlog : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.4" "postgresql-18-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-17-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-16-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-15-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-14-pg-dbms-errlog : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.4" "postgresql-18-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-17-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-16-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-15-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-14-pg-dbms-errlog : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.4" "postgresql-18-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-17-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-16-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-15-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-14-pg-dbms-errlog : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.4" "postgresql-18-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-17-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-16-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-15-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-14-pg-dbms-errlog : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.4" "postgresql-18-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-17-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-16-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-15-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-14-pg-dbms-errlog : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.4" "postgresql-18-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-17-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-16-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-15-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-14-pg-dbms-errlog : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.4" "postgresql-18-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-17-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-16-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-15-pg-dbms-errlog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4" "postgresql-14-pg-dbms-errlog : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_18` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_errlog_18-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_errlog_18-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.9 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_errlog_18-2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_errlog_18-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_errlog_18-2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_errlog_18-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.3 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_dbms_errlog_18-2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.6 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_dbms_errlog_18-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.9 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_errlog_18-2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_errlog_18-2.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-dbms-errlog` | `2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 62.1 KiB | [postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-errlog/postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-dbms-errlog` | `2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.1 KiB | [postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-errlog/postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-dbms-errlog` | `2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 62.2 KiB | [postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-errlog/postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-dbms-errlog` | `2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.1 KiB | [postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-errlog/postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-dbms-errlog` | `2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 67.3 KiB | [postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-errlog/postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-dbms-errlog` | `2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 66.5 KiB | [postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-errlog/postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-dbms-errlog` | `2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 64.7 KiB | [postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-errlog/postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-dbms-errlog` | `2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.0 KiB | [postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-errlog/postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-dbms-errlog` | `2.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 64.6 KiB | [postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-errlog/postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-dbms-errlog` | `2.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 64.0 KiB | [postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-errlog/postgresql-18-pg-dbms-errlog_2.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_17` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_errlog_17-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_errlog_17-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.9 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_errlog_17-2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_errlog_17-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_errlog_17-2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_errlog_17-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.3 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_dbms_errlog_17-2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.6 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_dbms_errlog_17-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_errlog_17-2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_errlog_17-2.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-dbms-errlog` | `2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 62.0 KiB | [postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-errlog/postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-dbms-errlog` | `2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.1 KiB | [postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-errlog/postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-dbms-errlog` | `2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 62.2 KiB | [postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-errlog/postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-dbms-errlog` | `2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.1 KiB | [postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-errlog/postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-dbms-errlog` | `2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 73.9 KiB | [postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-errlog/postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-dbms-errlog` | `2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 73.1 KiB | [postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-errlog/postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-dbms-errlog` | `2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 64.6 KiB | [postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-errlog/postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-dbms-errlog` | `2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.0 KiB | [postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-errlog/postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-dbms-errlog` | `2.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 64.5 KiB | [postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-errlog/postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-dbms-errlog` | `2.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 64.1 KiB | [postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-errlog/postgresql-17-pg-dbms-errlog_2.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_16` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_errlog_16-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_errlog_16-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_errlog_16-2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_errlog_16-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_errlog_16-2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_errlog_16-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.3 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_dbms_errlog_16-2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.6 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_dbms_errlog_16-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.9 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_errlog_16-2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_errlog_16-2.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-dbms-errlog` | `2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 62.0 KiB | [postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-errlog/postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-dbms-errlog` | `2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 60.9 KiB | [postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-errlog/postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-dbms-errlog` | `2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 62.2 KiB | [postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-errlog/postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-dbms-errlog` | `2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.0 KiB | [postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-errlog/postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-dbms-errlog` | `2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 73.4 KiB | [postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-errlog/postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-dbms-errlog` | `2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 72.7 KiB | [postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-errlog/postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-dbms-errlog` | `2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 64.5 KiB | [postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-errlog/postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-dbms-errlog` | `2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.0 KiB | [postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-errlog/postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-dbms-errlog` | `2.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 64.5 KiB | [postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-errlog/postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-dbms-errlog` | `2.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 63.9 KiB | [postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-errlog/postgresql-16-pg-dbms-errlog_2.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_15` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.6 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_errlog_15-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.8 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dbms_errlog_15-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.0 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_errlog_15-2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.1 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_errlog_15-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.7 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_errlog_15-2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.7 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_errlog_15-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.1 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_dbms_errlog_15-2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.5 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_dbms_errlog_15-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.8 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_errlog_15-2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_errlog_15-2.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-dbms-errlog` | `2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 62.3 KiB | [postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-errlog/postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-dbms-errlog` | `2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.3 KiB | [postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-errlog/postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-dbms-errlog` | `2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 62.5 KiB | [postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-errlog/postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-dbms-errlog` | `2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.3 KiB | [postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-errlog/postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-dbms-errlog` | `2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 74.0 KiB | [postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-errlog/postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-dbms-errlog` | `2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 73.2 KiB | [postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-errlog/postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-dbms-errlog` | `2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 65.1 KiB | [postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-errlog/postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-dbms-errlog` | `2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.6 KiB | [postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-errlog/postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-dbms-errlog` | `2.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 64.5 KiB | [postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-errlog/postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-dbms-errlog` | `2.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 64.2 KiB | [postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-errlog/postgresql-15-pg-dbms-errlog_2.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_14` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.5 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_errlog_14-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.8 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dbms_errlog_14-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.9 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_errlog_14-2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.1 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_errlog_14-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.6 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_errlog_14-2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.6 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_errlog_14-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.1 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_dbms_errlog_14-2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.4 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_dbms_errlog_14-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.9 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_dbms_errlog_14-2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.1 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_dbms_errlog_14-2.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-dbms-errlog` | `2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 62.1 KiB | [postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-errlog/postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-dbms-errlog` | `2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.0 KiB | [postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-errlog/postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-dbms-errlog` | `2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 62.3 KiB | [postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-errlog/postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-dbms-errlog` | `2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.0 KiB | [postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-errlog/postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-dbms-errlog` | `2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 73.7 KiB | [postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-errlog/postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-dbms-errlog` | `2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 72.9 KiB | [postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-errlog/postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-dbms-errlog` | `2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 64.9 KiB | [postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-errlog/postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-dbms-errlog` | `2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.4 KiB | [postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-errlog/postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-dbms-errlog` | `2.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 64.4 KiB | [postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-errlog/postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-dbms-errlog` | `2.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 64.0 KiB | [postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-errlog/postgresql-14-pg-dbms-errlog_2.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HexaCluster/pg_dbms_errlog" title="Repository" icon="github" subtitle="github.com/HexaCluster/pg_dbms_errlog" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_dbms_errlog-2.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_dbms_errlog;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_dbms_errlog;		# install via package name, for the active PG version

pig install pg_dbms_errlog -v 18;   # install for PG 18
pig install pg_dbms_errlog -v 17;   # install for PG 17
pig install pg_dbms_errlog -v 16;   # install for PG 16
pig install pg_dbms_errlog -v 15;   # install for PG 15
pig install pg_dbms_errlog -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_statement_rollback, pg_dbms_errlog';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_dbms_errlog CASCADE; -- requires pg_statement_rollback
```

## Usage

Sources:

- [Official v2.4 README](https://github.com/HexaCluster/pg_dbms_errlog/blob/v2.4/README.md)
- [v2.4 release changelog](https://github.com/HexaCluster/pg_dbms_errlog/blob/v2.4/ChangeLog)
- [v2.4 control file](https://github.com/HexaCluster/pg_dbms_errlog/blob/v2.4/pg_dbms_errlog.control)
- [v2.4 extension SQL](https://github.com/HexaCluster/pg_dbms_errlog/blob/v2.4/sql/pg_dbms_errlog--2.4.sql)

`pg_dbms_errlog` provides Oracle-style DML error logging for PostgreSQL. It queues an error from a failed `INSERT`, `UPDATE`, or `DELETE`, writes it to a registered `ERR$_...` table through background workers, and lets the surrounding script continue after rolling back to a savepoint. It requires either `pg_statement_rollback` or explicit savepoint management by the caller.

### Enable the Extension

Add the library to `shared_preload_libraries`, ensure `max_worker_processes` can accommodate `pg_dbms_errlog.max_workers` plus the fixed worker, and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_dbms_errlog'
```

```sql
CREATE EXTENSION pg_dbms_errlog;
```

Create and register an error table for each DML target:

```sql
CREATE TABLE raises (
    employee_id integer,
    salary integer CHECK (salary > 8000)
);

CALL dbms_errlog.create_error_log('raises');
-- Creates and registers public."ERR$_raises" by default.
```

### Log and Continue after an Error

```sql
LOAD 'pg_statement_rollback';

SET pg_statement_rollback.enabled = on;
SET pg_dbms_errlog.enabled = on;
SET pg_dbms_errlog.query_tag = 'daily_load';
SET pg_dbms_errlog.reject_limit = 10;

BEGIN;
INSERT INTO raises VALUES (145, 15400);
INSERT INTO raises VALUES (161, 7700);  -- logged failure
ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt";
INSERT INTO raises VALUES (175, 9680);
COMMIT;

SELECT * FROM "ERR$_raises";
```

The error table contains `pg_err_number$`, `pg_err_mesg$`, `pg_err_optyp$`, `pg_err_tag$`, `pg_err_query$`, and `pg_err_detail$`.

### API and Configuration Index

- `dbms_errlog.create_error_log(dml_table_name, err_log_table_name, err_log_table_owner, err_log_table_space)`: creates and registers an error table.
- `dbms_errlog.publish_queue(wait_for_completion)`: asks workers to process queued errors; execution is not granted to `PUBLIC` by default.
- `dbms_errlog.queue_size()`: reports queued errors.
- `pg_dbms_errlog.synchronous`: `transaction` by default, `query`, or `off`. Transaction mode guarantees that only errors from committed transactions are logged.
- `pg_dbms_errlog.reject_limit`: transaction-wide error limit; `-1` is unlimited and `0` logs nothing and rolls back.
- `pg_dbms_errlog.no_client_error`: suppresses client error messages while retaining server logging; enabled by default.
- `pg_dbms_errlog.frequency` and `pg_dbms_errlog.max_workers`: asynchronous worker timing and concurrency.

### Caveats

- A caller needs DML privileges on the target and error tables; creating an error table also requires execution and registration-table privileges described upstream.
- `INSERT INTO ... SELECT ...` is one PostgreSQL statement and cannot preserve only successful rows in the Oracle manner.
- Syntax and other parse-time errors are not logged. Stored query text must remain below PostgreSQL's 1 GB value limit.
- Version `2.4` changes no SQL API; it fixes worker shutdown loops and a dynamic-background-worker crash.
