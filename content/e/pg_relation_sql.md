---
title: "pg_relation_sql"
linkTitle: "pg_relation_sql"
description: "Generate inlinable SQL functions for navigating PostgreSQL foreign-key relations"
weight: 4210
categories: ["UTIL"]
languages: ["PLpgSQL"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_relation_sql**](https://github.com/asmgit/pg_relation_sql) : Generate inlinable SQL functions for navigating PostgreSQL foreign-key relations


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4210** | {{< badge content="pg_relation_sql" link="https://github.com/asmgit/pg_relation_sql" >}} | {{< ext "pg_relation_sql" >}} | `0.2.2` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "PLpgSQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-------" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_readme" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_get_functiondef" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_query_rewrite" >}} {{< ext "ddl_historization" >}} {{< ext "data_historization" >}} |

> [!Note] Upstream intentionally ships no control file or CREATE EXTENSION path; execute the packaged relation_sql.sql in each database; relation_sql('install') requires superuser only for its optional event trigger.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_relation_sql` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.2` | {{< bg "18" "pg_relation_sql_18" "green" >}} {{< bg "17" "pg_relation_sql_17" "green" >}} {{< bg "16" "pg_relation_sql_16" "green" >}} {{< bg "15" "pg_relation_sql_15" "green" >}} {{< bg "14" "pg_relation_sql_14" "green" >}} | `pg_relation_sql_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.2` | {{< bg "18" "postgresql-18-pg-relation-sql" "green" >}} {{< bg "17" "postgresql-17-pg-relation-sql" "green" >}} {{< bg "16" "postgresql-16-pg-relation-sql" "green" >}} {{< bg "15" "postgresql-15-pg-relation-sql" "green" >}} {{< bg "14" "postgresql-14-pg-relation-sql" "green" >}} | `postgresql-$v-pg-relation-sql` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pg_relation_sql_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-relation-sql : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-relation-sql : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-relation-sql : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-relation-sql : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-relation-sql : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-relation-sql : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-relation-sql : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-relation-sql : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-relation-sql : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pg-relation-sql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pg-relation-sql : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relation_sql_18` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.0 KiB | [pg_relation_sql_18-0.2.2-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relation_sql_18-0.2.2-1PGSTY.el8.noarch.rpm) |
| `pg_relation_sql_18` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [pg_relation_sql_18-0.2.2-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relation_sql_18-0.2.2-1PGSTY.el8.noarch.rpm) |
| `pg_relation_sql_18` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.8 KiB | [pg_relation_sql_18-0.2.2-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relation_sql_18-0.2.2-1PGSTY.el9.noarch.rpm) |
| `pg_relation_sql_18` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_relation_sql_18-0.2.2-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relation_sql_18-0.2.2-1PGSTY.el9.noarch.rpm) |
| `pg_relation_sql_18` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.0 KiB | [pg_relation_sql_18-0.2.2-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relation_sql_18-0.2.2-1PGSTY.el10.noarch.rpm) |
| `pg_relation_sql_18` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [pg_relation_sql_18-0.2.2-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relation_sql_18-0.2.2-1PGSTY.el10.noarch.rpm) |
| `postgresql-18-pg-relation-sql` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 14.1 KiB | [postgresql-18-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relation-sql/postgresql-18-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-relation-sql` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 14.1 KiB | [postgresql-18-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relation-sql/postgresql-18-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-relation-sql` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.1 KiB | [postgresql-18-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relation-sql/postgresql-18-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-relation-sql` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 14.1 KiB | [postgresql-18-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relation-sql/postgresql-18-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-relation-sql` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.3 KiB | [postgresql-18-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relation-sql/postgresql-18-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-relation-sql` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.3 KiB | [postgresql-18-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relation-sql/postgresql-18-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-relation-sql` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.3 KiB | [postgresql-18-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relation-sql/postgresql-18-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-relation-sql` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.3 KiB | [postgresql-18-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relation-sql/postgresql-18-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-relation-sql` | `0.2.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 14.3 KiB | [postgresql-18-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relation-sql/postgresql-18-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb) |
| `postgresql-18-pg-relation-sql` | `0.2.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 14.3 KiB | [postgresql-18-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relation-sql/postgresql-18-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relation_sql_17` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.0 KiB | [pg_relation_sql_17-0.2.2-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relation_sql_17-0.2.2-1PGSTY.el8.noarch.rpm) |
| `pg_relation_sql_17` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [pg_relation_sql_17-0.2.2-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relation_sql_17-0.2.2-1PGSTY.el8.noarch.rpm) |
| `pg_relation_sql_17` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.8 KiB | [pg_relation_sql_17-0.2.2-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relation_sql_17-0.2.2-1PGSTY.el9.noarch.rpm) |
| `pg_relation_sql_17` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_relation_sql_17-0.2.2-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relation_sql_17-0.2.2-1PGSTY.el9.noarch.rpm) |
| `pg_relation_sql_17` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.0 KiB | [pg_relation_sql_17-0.2.2-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relation_sql_17-0.2.2-1PGSTY.el10.noarch.rpm) |
| `pg_relation_sql_17` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [pg_relation_sql_17-0.2.2-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relation_sql_17-0.2.2-1PGSTY.el10.noarch.rpm) |
| `postgresql-17-pg-relation-sql` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 14.1 KiB | [postgresql-17-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relation-sql/postgresql-17-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-relation-sql` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 14.1 KiB | [postgresql-17-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relation-sql/postgresql-17-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-relation-sql` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.1 KiB | [postgresql-17-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relation-sql/postgresql-17-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-relation-sql` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 14.1 KiB | [postgresql-17-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relation-sql/postgresql-17-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-relation-sql` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.3 KiB | [postgresql-17-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relation-sql/postgresql-17-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-relation-sql` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.3 KiB | [postgresql-17-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relation-sql/postgresql-17-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-relation-sql` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.3 KiB | [postgresql-17-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relation-sql/postgresql-17-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-relation-sql` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.3 KiB | [postgresql-17-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relation-sql/postgresql-17-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-relation-sql` | `0.2.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 14.3 KiB | [postgresql-17-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relation-sql/postgresql-17-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb) |
| `postgresql-17-pg-relation-sql` | `0.2.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 14.3 KiB | [postgresql-17-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relation-sql/postgresql-17-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relation_sql_16` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.0 KiB | [pg_relation_sql_16-0.2.2-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relation_sql_16-0.2.2-1PGSTY.el8.noarch.rpm) |
| `pg_relation_sql_16` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [pg_relation_sql_16-0.2.2-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relation_sql_16-0.2.2-1PGSTY.el8.noarch.rpm) |
| `pg_relation_sql_16` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.8 KiB | [pg_relation_sql_16-0.2.2-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relation_sql_16-0.2.2-1PGSTY.el9.noarch.rpm) |
| `pg_relation_sql_16` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_relation_sql_16-0.2.2-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relation_sql_16-0.2.2-1PGSTY.el9.noarch.rpm) |
| `pg_relation_sql_16` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.0 KiB | [pg_relation_sql_16-0.2.2-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relation_sql_16-0.2.2-1PGSTY.el10.noarch.rpm) |
| `pg_relation_sql_16` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [pg_relation_sql_16-0.2.2-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relation_sql_16-0.2.2-1PGSTY.el10.noarch.rpm) |
| `postgresql-16-pg-relation-sql` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 14.1 KiB | [postgresql-16-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relation-sql/postgresql-16-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-relation-sql` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 14.1 KiB | [postgresql-16-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relation-sql/postgresql-16-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-relation-sql` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.1 KiB | [postgresql-16-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relation-sql/postgresql-16-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-relation-sql` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 14.1 KiB | [postgresql-16-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relation-sql/postgresql-16-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-relation-sql` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.3 KiB | [postgresql-16-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relation-sql/postgresql-16-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-relation-sql` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.3 KiB | [postgresql-16-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relation-sql/postgresql-16-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-relation-sql` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.3 KiB | [postgresql-16-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relation-sql/postgresql-16-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-relation-sql` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.3 KiB | [postgresql-16-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relation-sql/postgresql-16-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-relation-sql` | `0.2.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 14.3 KiB | [postgresql-16-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relation-sql/postgresql-16-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb) |
| `postgresql-16-pg-relation-sql` | `0.2.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 14.3 KiB | [postgresql-16-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relation-sql/postgresql-16-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relation_sql_15` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.0 KiB | [pg_relation_sql_15-0.2.2-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relation_sql_15-0.2.2-1PGSTY.el8.noarch.rpm) |
| `pg_relation_sql_15` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [pg_relation_sql_15-0.2.2-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relation_sql_15-0.2.2-1PGSTY.el8.noarch.rpm) |
| `pg_relation_sql_15` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.8 KiB | [pg_relation_sql_15-0.2.2-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relation_sql_15-0.2.2-1PGSTY.el9.noarch.rpm) |
| `pg_relation_sql_15` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_relation_sql_15-0.2.2-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relation_sql_15-0.2.2-1PGSTY.el9.noarch.rpm) |
| `pg_relation_sql_15` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.0 KiB | [pg_relation_sql_15-0.2.2-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relation_sql_15-0.2.2-1PGSTY.el10.noarch.rpm) |
| `pg_relation_sql_15` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [pg_relation_sql_15-0.2.2-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relation_sql_15-0.2.2-1PGSTY.el10.noarch.rpm) |
| `postgresql-15-pg-relation-sql` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 14.1 KiB | [postgresql-15-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relation-sql/postgresql-15-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pg-relation-sql` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 14.1 KiB | [postgresql-15-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relation-sql/postgresql-15-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pg-relation-sql` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.1 KiB | [postgresql-15-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relation-sql/postgresql-15-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb) |
| `postgresql-15-pg-relation-sql` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 14.1 KiB | [postgresql-15-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relation-sql/postgresql-15-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb) |
| `postgresql-15-pg-relation-sql` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.3 KiB | [postgresql-15-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relation-sql/postgresql-15-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb) |
| `postgresql-15-pg-relation-sql` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.3 KiB | [postgresql-15-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relation-sql/postgresql-15-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb) |
| `postgresql-15-pg-relation-sql` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.3 KiB | [postgresql-15-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relation-sql/postgresql-15-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb) |
| `postgresql-15-pg-relation-sql` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.3 KiB | [postgresql-15-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relation-sql/postgresql-15-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb) |
| `postgresql-15-pg-relation-sql` | `0.2.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 14.3 KiB | [postgresql-15-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relation-sql/postgresql-15-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb) |
| `postgresql-15-pg-relation-sql` | `0.2.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 14.3 KiB | [postgresql-15-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relation-sql/postgresql-15-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relation_sql_14` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.0 KiB | [pg_relation_sql_14-0.2.2-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relation_sql_14-0.2.2-1PGSTY.el8.noarch.rpm) |
| `pg_relation_sql_14` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [pg_relation_sql_14-0.2.2-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relation_sql_14-0.2.2-1PGSTY.el8.noarch.rpm) |
| `pg_relation_sql_14` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.8 KiB | [pg_relation_sql_14-0.2.2-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relation_sql_14-0.2.2-1PGSTY.el9.noarch.rpm) |
| `pg_relation_sql_14` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.8 KiB | [pg_relation_sql_14-0.2.2-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relation_sql_14-0.2.2-1PGSTY.el9.noarch.rpm) |
| `pg_relation_sql_14` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.0 KiB | [pg_relation_sql_14-0.2.2-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relation_sql_14-0.2.2-1PGSTY.el10.noarch.rpm) |
| `pg_relation_sql_14` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [pg_relation_sql_14-0.2.2-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relation_sql_14-0.2.2-1PGSTY.el10.noarch.rpm) |
| `postgresql-14-pg-relation-sql` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 14.1 KiB | [postgresql-14-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relation-sql/postgresql-14-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pg-relation-sql` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 14.1 KiB | [postgresql-14-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relation-sql/postgresql-14-pg-relation-sql_0.2.2-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pg-relation-sql` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.1 KiB | [postgresql-14-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relation-sql/postgresql-14-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb) |
| `postgresql-14-pg-relation-sql` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 14.1 KiB | [postgresql-14-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relation-sql/postgresql-14-pg-relation-sql_0.2.2-1PGSTY~trixie_all.deb) |
| `postgresql-14-pg-relation-sql` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.3 KiB | [postgresql-14-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relation-sql/postgresql-14-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb) |
| `postgresql-14-pg-relation-sql` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.3 KiB | [postgresql-14-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relation-sql/postgresql-14-pg-relation-sql_0.2.2-1PGSTY~jammy_all.deb) |
| `postgresql-14-pg-relation-sql` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.3 KiB | [postgresql-14-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relation-sql/postgresql-14-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb) |
| `postgresql-14-pg-relation-sql` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.3 KiB | [postgresql-14-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relation-sql/postgresql-14-pg-relation-sql_0.2.2-1PGSTY~noble_all.deb) |
| `postgresql-14-pg-relation-sql` | `0.2.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 14.3 KiB | [postgresql-14-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relation-sql/postgresql-14-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb) |
| `postgresql-14-pg-relation-sql` | `0.2.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 14.3 KiB | [postgresql-14-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relation-sql/postgresql-14-pg-relation-sql_0.2.2-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/asmgit/pg_relation_sql" title="Repository" icon="github" subtitle="github.com/asmgit/pg_relation_sql" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_relation_sql-0.2.2.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_relation_sql;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_relation_sql;		# install via package name, for the active PG version

pig install pg_relation_sql -v 18;   # install for PG 18
pig install pg_relation_sql -v 17;   # install for PG 17
pig install pg_relation_sql -v 16;   # install for PG 16
pig install pg_relation_sql -v 15;   # install for PG 15
pig install pg_relation_sql -v 14;   # install for PG 14

```


This extension does not need `CREATE EXTENSION` DDL command



## Usage

Sources:

- [pg_relation_sql 0.2.2 on PGXN](https://pgxn.org/dist/pg_relation_sql/0.2.2/)
- [pg_relation_sql 0.2.2 README](https://api.pgxn.org/src/pg_relation_sql/pg_relation_sql-0.2.2/README.md)
- [pg_relation_sql 0.2.2 SQL script](https://api.pgxn.org/src/pg_relation_sql/pg_relation_sql-0.2.2/relation_sql.sql)
- [pg_relation_sql 0.2.2 plan comparison](https://api.pgxn.org/src/pg_relation_sql/pg_relation_sql-0.2.2/EXPLAIN.md)

`pg_relation_sql` 0.2.2 generates pairs of SQL functions from PostgreSQL foreign keys: a lookup follows a reference, while a list function returns rows that point back. The generated `LANGUAGE sql` functions are designed to be inlined by the planner, allowing queries to navigate declared relations without repeating join conditions.

Upstream deliberately ships one standalone `relation_sql.sql` file rather than a control file. There is no `CREATE EXTENSION pg_relation_sql`; execute the packaged script in every database where the functions are needed.

```bash
psql app -f /usr/pgsql-17/share/pg_relation_sql/relation_sql.sql
psql app -f /usr/share/postgresql/17/pg_relation_sql/relation_sql.sql
```

The script creates `relation_sql(text)` in the current schema and finishes by requesting `relation_sql('install')`.

### Generate and Use Relations

```sql
CREATE TABLE profile (
  id bigint PRIMARY KEY,
  name text
);

CREATE TABLE address (
  id bigint PRIMARY KEY,
  profile_id bigint REFERENCES profile(id),
  city text
);

SELECT status, command FROM relation_sql('sync');

SELECT a.city, p.name
FROM address AS a, profile(a) AS p;

SELECT p.name, a.city
FROM profile AS p, address_list(p) AS a;
```

For each foreign key, the lookup function follows the referenced row and the reverse function uses a `_list` suffix unless the foreign key is one-to-one. Composite and cross-schema foreign keys are supported, and several foreign keys to the same target receive role-specific names.

### Generator Modes

- `relation_sql()` returns a status dashboard.
- `relation_sql('show')` reports the computed functions and ready-to-run synchronization commands without changing objects.
- `relation_sql('sync')` creates, replaces, or removes marked relation functions to match current foreign keys.
- `relation_sql('install')` adds a `ddl_command_end` event trigger and synchronizes immediately.
- `relation_sql('uninstall')` removes the event trigger; `relation_sql('drop')` removes generated functions.

### Operational Boundaries

- Creating the event trigger requires superuser privileges. Without them, installation emits a warning and the one-time synchronization still runs with the caller's object privileges.
- Install the generator in a trusted schema with a controlled `search_path`: automatic mode creates a `SECURITY DEFINER` event-trigger helper that preserves the installation-time path.
- Generated functions depend on table row types. Dropping a table whose row type is used by them can require `CASCADE`; inspect dependencies before destructive DDL.
- The generated bodies use `SELECT *`, so column-level `SELECT` grants do not combine cleanly with them. Row-level security continues to apply.
- Put relation functions in `FROM` for plan-sensitive queries. Attribute notation in a select list becomes a `ProjectSet`, and `NOT EXISTS (SELECT FROM relation_function(row))` can remain a correlated probe instead of becoming the equivalent anti-join.
- Queries depend on generated functions just as they depend on views. Run `relation_sql('sync')` in the migration path when not using the event trigger.
- Upstream requires PostgreSQL 11 or later; Pigsty packages cover PostgreSQL 14–18.
