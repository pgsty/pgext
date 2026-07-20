---
title: "pg_dbms_lock"
linkTitle: "pg_dbms_lock"
description: "Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL"
weight: 9250
categories: ["SIM"]
width: full
---

[**pg_dbms_lock**](https://github.com/HexaCluster/pg_dbms_lock) : Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9250** | {{< badge content="pg_dbms_lock" link="https://github.com/HexaCluster/pg_dbms_lock" >}} | {{< ext "pg_dbms_lock" >}} | `2.0` | {{< category "SIM" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `dbms_lock` |
|   **Requires**    | {{< ext "pg_background" >}} |
|   **See Also**    | {{< ext "orafce" >}} {{< ext "session_variable" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_job" >}} {{< ext "oracle_fdw" >}} {{< ext "pgtt" >}} {{< ext "pg_statement_rollback" >}} {{< ext "mysql_fdw" >}} |

> [!Note] Requires pg_background >= 2.0; the PIGSTY DEB carries a downstream pg_background v2 API compatibility patch.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_dbms_lock` | `pg_background` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0` | {{< bg "18" "pg_dbms_lock_18" "green" >}} {{< bg "17" "pg_dbms_lock_17" "green" >}} {{< bg "16" "pg_dbms_lock_16" "green" >}} {{< bg "15" "pg_dbms_lock_15" "green" >}} {{< bg "14" "pg_dbms_lock_14" "green" >}} | `pg_dbms_lock_$v` | `pg_background_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0` | {{< bg "18" "postgresql-18-pg-dbms-lock" "green" >}} {{< bg "17" "postgresql-17-pg-dbms-lock" "green" >}} {{< bg "16" "postgresql-16-pg-dbms-lock" "green" >}} {{< bg "15" "postgresql-15-pg-dbms-lock" "green" >}} {{< bg "14" "postgresql-14-pg-dbms-lock" "green" >}} | `postgresql-$v-pg-dbms-lock` | `postgresql-$v-pg-background` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_14 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0" "pg_dbms_lock_14 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg-dbms-lock : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg-dbms-lock : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg-dbms-lock : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg-dbms-lock : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg-dbms-lock : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg-dbms-lock : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg-dbms-lock : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg-dbms-lock : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg-dbms-lock : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg-dbms-lock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg-dbms-lock : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_lock_18` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.1 KiB | [pg_dbms_lock_18-2.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_lock_18-2.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_lock_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.7 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_lock_18-1.0-3PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_18` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.1 KiB | [pg_dbms_lock_18-2.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_lock_18-2.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_lock_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.7 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_lock_18-1.0-3PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_18` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [pg_dbms_lock_18-2.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_lock_18-2.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.6 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_lock_18-1.0-3PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.6 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_lock_18-1.0-3PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_18` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pg_dbms_lock_18-2.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_lock_18-2.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.6 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_lock_18-1.0-3PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.6 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_lock_18-1.0-3PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_18` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pg_dbms_lock_18-2.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_dbms_lock_18-2.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.8 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_dbms_lock_18-1.0-3PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_dbms_lock_18-1.0-3PGDG.rhel10.noarch.rpm) |
| `pg_dbms_lock_18` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [pg_dbms_lock_18-2.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_lock_18-2.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.7 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_lock_18-1.0-3PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.1 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_lock_18-1.0-3PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pg-dbms-lock` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-18-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-lock/postgresql-18-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-dbms-lock` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-18-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-lock/postgresql-18-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-dbms-lock` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-18-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-lock/postgresql-18-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-dbms-lock` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-18-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-lock/postgresql-18-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-dbms-lock` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 8.8 KiB | [postgresql-18-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-lock/postgresql-18-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-dbms-lock` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.8 KiB | [postgresql-18-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-lock/postgresql-18-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-dbms-lock` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.8 KiB | [postgresql-18-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-lock/postgresql-18-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-dbms-lock` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.8 KiB | [postgresql-18-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-lock/postgresql-18-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-dbms-lock` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 8.9 KiB | [postgresql-18-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-lock/postgresql-18-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb) |
| `postgresql-18-pg-dbms-lock` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.9 KiB | [postgresql-18-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-lock/postgresql-18-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_lock_17` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.1 KiB | [pg_dbms_lock_17-2.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_lock_17-2.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_lock_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.5 KiB | [pg_dbms_lock_17-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_lock_17-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_17` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.1 KiB | [pg_dbms_lock_17-2.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_lock_17-2.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_lock_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.5 KiB | [pg_dbms_lock_17-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_lock_17-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_17` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [pg_dbms_lock_17-2.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_lock_17-2.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.6 KiB | [pg_dbms_lock_17-1.0-3PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_lock_17-1.0-3PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.4 KiB | [pg_dbms_lock_17-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_lock_17-1.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_17` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pg_dbms_lock_17-2.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_lock_17-2.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.6 KiB | [pg_dbms_lock_17-1.0-3PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_lock_17-1.0-3PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.4 KiB | [pg_dbms_lock_17-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_lock_17-1.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_17` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pg_dbms_lock_17-2.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_dbms_lock_17-2.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.8 KiB | [pg_dbms_lock_17-1.0-3PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_dbms_lock_17-1.0-3PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.0 KiB | [pg_dbms_lock_17-1.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_dbms_lock_17-1.0-2PGDG.rhel10.noarch.rpm) |
| `pg_dbms_lock_17` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [pg_dbms_lock_17-2.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_lock_17-2.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.7 KiB | [pg_dbms_lock_17-1.0-3PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_lock_17-1.0-3PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [pg_dbms_lock_17-1.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_lock_17-1.0-2PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pg-dbms-lock` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-17-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-lock/postgresql-17-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-dbms-lock` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-17-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-lock/postgresql-17-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-dbms-lock` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-17-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-lock/postgresql-17-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-dbms-lock` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-17-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-lock/postgresql-17-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-dbms-lock` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 8.9 KiB | [postgresql-17-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-lock/postgresql-17-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-dbms-lock` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.9 KiB | [postgresql-17-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-lock/postgresql-17-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-dbms-lock` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.9 KiB | [postgresql-17-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-lock/postgresql-17-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-dbms-lock` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.9 KiB | [postgresql-17-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-lock/postgresql-17-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-dbms-lock` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 8.9 KiB | [postgresql-17-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-lock/postgresql-17-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb) |
| `postgresql-17-pg-dbms-lock` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.9 KiB | [postgresql-17-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-lock/postgresql-17-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_lock_16` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.1 KiB | [pg_dbms_lock_16-2.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_lock_16-2.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_lock_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.5 KiB | [pg_dbms_lock_16-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_lock_16-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_16` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.1 KiB | [pg_dbms_lock_16-2.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_lock_16-2.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_lock_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.5 KiB | [pg_dbms_lock_16-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_lock_16-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_16` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [pg_dbms_lock_16-2.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_lock_16-2.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.6 KiB | [pg_dbms_lock_16-1.0-3PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_lock_16-1.0-3PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.4 KiB | [pg_dbms_lock_16-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_lock_16-1.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_16` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pg_dbms_lock_16-2.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_lock_16-2.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.6 KiB | [pg_dbms_lock_16-1.0-3PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_lock_16-1.0-3PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.2 KiB | [pg_dbms_lock_16-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_lock_16-1.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_16` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pg_dbms_lock_16-2.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_dbms_lock_16-2.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.8 KiB | [pg_dbms_lock_16-1.0-3PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_dbms_lock_16-1.0-3PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.0 KiB | [pg_dbms_lock_16-1.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_dbms_lock_16-1.0-2PGDG.rhel10.noarch.rpm) |
| `pg_dbms_lock_16` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [pg_dbms_lock_16-2.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_lock_16-2.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.7 KiB | [pg_dbms_lock_16-1.0-3PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_lock_16-1.0-3PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [pg_dbms_lock_16-1.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_lock_16-1.0-2PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pg-dbms-lock` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-16-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-lock/postgresql-16-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-dbms-lock` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-16-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-lock/postgresql-16-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-dbms-lock` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-16-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-lock/postgresql-16-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-dbms-lock` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-16-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-lock/postgresql-16-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-dbms-lock` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 8.9 KiB | [postgresql-16-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-lock/postgresql-16-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-dbms-lock` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.9 KiB | [postgresql-16-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-lock/postgresql-16-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-dbms-lock` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.8 KiB | [postgresql-16-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-lock/postgresql-16-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-dbms-lock` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.8 KiB | [postgresql-16-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-lock/postgresql-16-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-dbms-lock` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 8.9 KiB | [postgresql-16-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-lock/postgresql-16-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb) |
| `postgresql-16-pg-dbms-lock` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.9 KiB | [postgresql-16-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-lock/postgresql-16-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_lock_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.1 KiB | [pg_dbms_lock_15-2.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_lock_15-2.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_lock_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.5 KiB | [pg_dbms_lock_15-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_lock_15-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.1 KiB | [pg_dbms_lock_15-2.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dbms_lock_15-2.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_lock_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.5 KiB | [pg_dbms_lock_15-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dbms_lock_15-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_15` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [pg_dbms_lock_15-2.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_lock_15-2.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.6 KiB | [pg_dbms_lock_15-1.0-3PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_lock_15-1.0-3PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.4 KiB | [pg_dbms_lock_15-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_lock_15-1.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_15` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pg_dbms_lock_15-2.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_lock_15-2.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.6 KiB | [pg_dbms_lock_15-1.0-3PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_lock_15-1.0-3PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.2 KiB | [pg_dbms_lock_15-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_lock_15-1.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_15` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pg_dbms_lock_15-2.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_dbms_lock_15-2.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.8 KiB | [pg_dbms_lock_15-1.0-3PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_dbms_lock_15-1.0-3PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.0 KiB | [pg_dbms_lock_15-1.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_dbms_lock_15-1.0-2PGDG.rhel10.noarch.rpm) |
| `pg_dbms_lock_15` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [pg_dbms_lock_15-2.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_lock_15-2.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.7 KiB | [pg_dbms_lock_15-1.0-3PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_lock_15-1.0-3PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [pg_dbms_lock_15-1.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_lock_15-1.0-2PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pg-dbms-lock` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-15-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-lock/postgresql-15-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-dbms-lock` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-15-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-lock/postgresql-15-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-dbms-lock` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-15-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-lock/postgresql-15-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-dbms-lock` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-15-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-lock/postgresql-15-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-dbms-lock` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 8.9 KiB | [postgresql-15-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-lock/postgresql-15-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-dbms-lock` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.9 KiB | [postgresql-15-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-lock/postgresql-15-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-dbms-lock` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.9 KiB | [postgresql-15-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-lock/postgresql-15-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-dbms-lock` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.9 KiB | [postgresql-15-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-lock/postgresql-15-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-dbms-lock` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 8.9 KiB | [postgresql-15-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-lock/postgresql-15-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb) |
| `postgresql-15-pg-dbms-lock` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.9 KiB | [postgresql-15-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-lock/postgresql-15-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_lock_14` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.1 KiB | [pg_dbms_lock_14-2.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_lock_14-2.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_lock_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.5 KiB | [pg_dbms_lock_14-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_lock_14-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_14` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.1 KiB | [pg_dbms_lock_14-2.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dbms_lock_14-2.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_lock_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.5 KiB | [pg_dbms_lock_14-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dbms_lock_14-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_14` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [pg_dbms_lock_14-2.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_lock_14-2.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.6 KiB | [pg_dbms_lock_14-1.0-3PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_lock_14-1.0-3PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.4 KiB | [pg_dbms_lock_14-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_lock_14-1.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_14` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [pg_dbms_lock_14-2.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_lock_14-2.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.6 KiB | [pg_dbms_lock_14-1.0-3PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_lock_14-1.0-3PGDG.rhel9.8.noarch.rpm) |
| `pg_dbms_lock_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.2 KiB | [pg_dbms_lock_14-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_lock_14-1.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_14` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [pg_dbms_lock_14-2.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_dbms_lock_14-2.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.8 KiB | [pg_dbms_lock_14-1.0-3PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_dbms_lock_14-1.0-3PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.0 KiB | [pg_dbms_lock_14-1.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_dbms_lock_14-1.0-2PGDG.rhel10.noarch.rpm) |
| `pg_dbms_lock_14` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [pg_dbms_lock_14-2.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_dbms_lock_14-2.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.7 KiB | [pg_dbms_lock_14-1.0-3PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_dbms_lock_14-1.0-3PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_lock_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [pg_dbms_lock_14-1.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_dbms_lock_14-1.0-2PGDG.rhel10.noarch.rpm) |
| `postgresql-14-pg-dbms-lock` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-14-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-lock/postgresql-14-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pg-dbms-lock` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-14-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-lock/postgresql-14-pg-dbms-lock_2.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pg-dbms-lock` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-14-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-lock/postgresql-14-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pg-dbms-lock` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-14-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-lock/postgresql-14-pg-dbms-lock_2.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pg-dbms-lock` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 8.8 KiB | [postgresql-14-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-lock/postgresql-14-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pg-dbms-lock` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.8 KiB | [postgresql-14-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-lock/postgresql-14-pg-dbms-lock_2.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pg-dbms-lock` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.8 KiB | [postgresql-14-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-lock/postgresql-14-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-pg-dbms-lock` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.8 KiB | [postgresql-14-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-lock/postgresql-14-pg-dbms-lock_2.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-pg-dbms-lock` | `2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 8.9 KiB | [postgresql-14-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-lock/postgresql-14-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb) |
| `postgresql-14-pg-dbms-lock` | `2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.9 KiB | [postgresql-14-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-lock/postgresql-14-pg-dbms-lock_2.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HexaCluster/pg_dbms_lock" title="Repository" icon="github" subtitle="github.com/HexaCluster/pg_dbms_lock" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_dbms_lock-2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_dbms_lock;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_dbms_lock;		# install via package name, for the active PG version

pig install pg_dbms_lock -v 18;   # install for PG 18
pig install pg_dbms_lock -v 17;   # install for PG 17
pig install pg_dbms_lock -v 16;   # install for PG 16
pig install pg_dbms_lock -v 15;   # install for PG 15
pig install pg_dbms_lock -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_dbms_lock CASCADE; -- requires pg_background
```




## Usage

> [pg_dbms_lock: Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL](https://github.com/HexaCluster/pg_dbms_lock)

Uses PostgreSQL advisory locks to emulate Oracle DBMS_LOCK behavior.

### Enabling

```sql
CREATE EXTENSION pg_dbms_lock;
```

### ALLOCATE_UNIQUE

Allocate a unique lock identifier for a named lock:

```sql
DO $$
DECLARE
    lockhandle varchar;
BEGIN
    CALL dbms_lock.allocate_unique(
        lockname => 'printer_lock',
        lockhandle => lockhandle
    );
    RAISE NOTICE 'Handle: %', lockhandle;
END;
$$;
```

### REQUEST

Request a lock with a specific mode (Exclusive=6, Shared=4):

```sql
DO $$
DECLARE
    lock_res int;
BEGIN
    lock_res := dbms_lock.request(
        id => 123,
        lockmode => 6,           -- Exclusive
        timeout => 300,
        release_on_commit => false
    );
    IF lock_res <> 0 THEN
        RAISE EXCEPTION 'Lock request failed: %', lock_res;
    END IF;
END;
$$;
```

Return values: 0=Success, 1=Timeout, 3=Parameter error, 4=Already own lock, 5=Illegal handle.

### RELEASE

Explicitly release a previously acquired lock:

```sql
DO $$
DECLARE
    lock_res int;
BEGIN
    lock_res := dbms_lock.release(id => 123);
    IF lock_res <> 0 THEN
        RAISE EXCEPTION 'Release failed: %', lock_res;
    END IF;
END;
$$;
```

### SLEEP

Suspend the session for a given duration:

```sql
CALL dbms_lock.sleep(0.70);  -- sleep 0.7 seconds
```

### Complete Example

```sql
DO $$
DECLARE
    lock_res int;
    printer_lockhandle varchar;
BEGIN
    CALL dbms_lock.allocate_unique('printer_lock', printer_lockhandle);
    lock_res := dbms_lock.request(lockhandle => printer_lockhandle, lockmode => 6, timeout => 5);
    IF lock_res <> 0 THEN
        RAISE EXCEPTION 'REQUEST failed: %', lock_res;
    END IF;
    -- do exclusive work here
    lock_res := dbms_lock.release(lockhandle => printer_lockhandle);
END;
$$;
```
