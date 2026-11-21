---
title: "pg_dbms_metadata"
linkTitle: "pg_dbms_metadata"
description: "Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL"
weight: 9240
categories: ["SIM"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["MIXED"]
page_width: full
---

[**pg_dbms_metadata**](https://github.com/HexaCluster/pg_dbms_metadata) : Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9240** | {{< badge content="pg_dbms_metadata" link="https://github.com/HexaCluster/pg_dbms_metadata" >}} | {{< ext "pg_dbms_metadata" >}} | `1.1.0` | {{< category "SIM" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `dbms_metadata` |
|   **See Also**    | {{< ext "pg_get_functiondef" >}} {{< ext "ddlx" >}} {{< ext "orafce" >}} {{< ext "pg_readme" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pg_dbms_job" >}} {{< ext "pg_dbms_errlog" >}} {{< ext "pg_utl_smtp" >}} {{< ext "pgdd" >}} {{< ext "meta" >}} |

> [!Note] PGDG RPM lacks PG15 on EL8 aarch64; PIGSTY DEB covers PG14-18 on amd64 and arm64.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_dbms_metadata` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.0` | {{< bg "19" "pg_dbms_metadata_19" "green" >}} {{< bg "18" "pg_dbms_metadata_18" "green" >}} {{< bg "17" "pg_dbms_metadata_17" "green" >}} {{< bg "16" "pg_dbms_metadata_16" "green" >}} {{< bg "15" "pg_dbms_metadata_15" "green" >}} | `pg_dbms_metadata_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "19" "postgresql-19-pg-dbms-metadata" "green" >}} {{< bg "18" "postgresql-18-pg-dbms-metadata" "green" >}} {{< bg "17" "postgresql-17-pg-dbms-metadata" "green" >}} {{< bg "16" "postgresql-16-pg-dbms-metadata" "green" >}} {{< bg "15" "postgresql-15-pg-dbms-metadata" "green" >}} | `postgresql-$v-pg-dbms-metadata` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_dbms_metadata_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_15 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_dbms_metadata_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_15 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_15 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_dbms_metadata_15 : AVAIL 5" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-dbms-metadata : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-dbms-metadata : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-dbms-metadata : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-dbms-metadata : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-dbms-metadata : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-dbms-metadata : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-dbms-metadata : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-dbms-metadata : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-dbms-metadata : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-dbms-metadata : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-dbms-metadata : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-dbms-metadata : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-dbms-metadata : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-dbms-metadata : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-dbms-metadata : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-dbms-metadata : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-dbms-metadata : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-dbms-metadata : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-dbms-metadata : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-dbms-metadata : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-dbms-metadata : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_metadata_19` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.5 KiB | [pg_dbms_metadata_19-1.1.0-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/pg_dbms_metadata_19-1.1.0-1PGDG.rhel9.9.noarch.rpm) |
| `pg_dbms_metadata_19` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.4 KiB | [pg_dbms_metadata_19-1.1.0-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/pg_dbms_metadata_19-1.1.0-1PGDG.rhel9.9.noarch.rpm) |
| `pg_dbms_metadata_19` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.6 KiB | [pg_dbms_metadata_19-1.1.0-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/pg_dbms_metadata_19-1.1.0-1PGDG.rhel10.3.noarch.rpm) |
| `pg_dbms_metadata_19` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.5 KiB | [pg_dbms_metadata_19-1.1.0-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/pg_dbms_metadata_19-1.1.0-1PGDG.rhel10.3.noarch.rpm) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_metadata_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.3 KiB | [pg_dbms_metadata_18-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_metadata_18-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.2 KiB | [pg_dbms_metadata_18-1.0.0-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_metadata_18-1.0.0-5PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.8 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pg_dbms_metadata_18-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_metadata_18-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.1 KiB | [pg_dbms_metadata_18-1.0.0-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_metadata_18-1.0.0-5PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.7 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.4 KiB | [pg_dbms_metadata_18-1.1.0-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_metadata_18-1.1.0-1PGDG.rhel9.9.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.4 KiB | [pg_dbms_metadata_18-1.1.0-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_metadata_18-1.1.0-1PGDG.rhel9.9.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.6 KiB | [pg_dbms_metadata_18-1.1.0-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_dbms_metadata_18-1.1.0-1PGDG.rhel10.3.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.5 KiB | [pg_dbms_metadata_18-1.1.0-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_metadata_18-1.1.0-1PGDG.rhel10.3.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.6 KiB | [pg_dbms_metadata_18-1.1.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_metadata_18-1.1.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.7 KiB | [pg_dbms_metadata_18-1.0.0-5PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_metadata_18-1.0.0-5PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.5 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pg-dbms-metadata` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.2 KiB | [postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-metadata/postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-dbms-metadata` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.2 KiB | [postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-metadata/postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-dbms-metadata` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.2 KiB | [postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-metadata/postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-dbms-metadata` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.2 KiB | [postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-metadata/postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-dbms-metadata` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.5 KiB | [postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-metadata/postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-dbms-metadata` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.5 KiB | [postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-metadata/postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-dbms-metadata` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.5 KiB | [postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-metadata/postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-dbms-metadata` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.5 KiB | [postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-metadata/postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-dbms-metadata` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.5 KiB | [postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-metadata/postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-18-pg-dbms-metadata` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.5 KiB | [postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-metadata/postgresql-18-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_metadata_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.3 KiB | [pg_dbms_metadata_17-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_metadata_17-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.2 KiB | [pg_dbms_metadata_17-1.0.0-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_metadata_17-1.0.0-5PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.7 KiB | [pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pg_dbms_metadata_17-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_metadata_17-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.1 KiB | [pg_dbms_metadata_17-1.0.0-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_metadata_17-1.0.0-5PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.6 KiB | [pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.5 KiB | [pg_dbms_metadata_17-1.1.0-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_metadata_17-1.1.0-1PGDG.rhel9.9.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.4 KiB | [pg_dbms_metadata_17-1.1.0-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_metadata_17-1.1.0-1PGDG.rhel9.9.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.6 KiB | [pg_dbms_metadata_17-1.1.0-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_dbms_metadata_17-1.1.0-1PGDG.rhel10.3.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.5 KiB | [pg_dbms_metadata_17-1.1.0-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_metadata_17-1.1.0-1PGDG.rhel10.3.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.6 KiB | [pg_dbms_metadata_17-1.1.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_metadata_17-1.1.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.7 KiB | [pg_dbms_metadata_17-1.0.0-5PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_metadata_17-1.0.0-5PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.5 KiB | [pg_dbms_metadata_17-1.0.0-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_metadata_17-1.0.0-2PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pg_dbms_metadata_17-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_metadata_17-1.0.0-2PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pg-dbms-metadata` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.2 KiB | [postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-metadata/postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-dbms-metadata` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.2 KiB | [postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-metadata/postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-dbms-metadata` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.2 KiB | [postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-metadata/postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-dbms-metadata` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.2 KiB | [postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-metadata/postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-dbms-metadata` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.5 KiB | [postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-metadata/postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-dbms-metadata` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.5 KiB | [postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-metadata/postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-dbms-metadata` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.5 KiB | [postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-metadata/postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-dbms-metadata` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.5 KiB | [postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-metadata/postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-dbms-metadata` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.5 KiB | [postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-metadata/postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-17-pg-dbms-metadata` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.5 KiB | [postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-metadata/postgresql-17-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_metadata_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.3 KiB | [pg_dbms_metadata_16-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_metadata_16-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.2 KiB | [pg_dbms_metadata_16-1.0.0-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_metadata_16-1.0.0-5PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.7 KiB | [pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pg_dbms_metadata_16-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_metadata_16-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.1 KiB | [pg_dbms_metadata_16-1.0.0-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_metadata_16-1.0.0-5PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.6 KiB | [pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.5 KiB | [pg_dbms_metadata_16-1.1.0-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_metadata_16-1.1.0-1PGDG.rhel9.9.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.4 KiB | [pg_dbms_metadata_16-1.1.0-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_metadata_16-1.1.0-1PGDG.rhel9.9.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.6 KiB | [pg_dbms_metadata_16-1.1.0-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_dbms_metadata_16-1.1.0-1PGDG.rhel10.3.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.5 KiB | [pg_dbms_metadata_16-1.1.0-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_metadata_16-1.1.0-1PGDG.rhel10.3.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.6 KiB | [pg_dbms_metadata_16-1.1.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_metadata_16-1.1.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.7 KiB | [pg_dbms_metadata_16-1.0.0-5PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_metadata_16-1.0.0-5PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.5 KiB | [pg_dbms_metadata_16-1.0.0-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_metadata_16-1.0.0-2PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pg_dbms_metadata_16-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_metadata_16-1.0.0-2PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pg-dbms-metadata` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.2 KiB | [postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-metadata/postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-dbms-metadata` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.2 KiB | [postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-metadata/postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-dbms-metadata` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.2 KiB | [postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-metadata/postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-dbms-metadata` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.2 KiB | [postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-metadata/postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-dbms-metadata` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.5 KiB | [postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-metadata/postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-dbms-metadata` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.5 KiB | [postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-metadata/postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-dbms-metadata` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.5 KiB | [postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-metadata/postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-dbms-metadata` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.5 KiB | [postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-metadata/postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-dbms-metadata` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.5 KiB | [postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-metadata/postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-16-pg-dbms-metadata` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.5 KiB | [postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-metadata/postgresql-16-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_metadata_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.3 KiB | [pg_dbms_metadata_15-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_metadata_15-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.2 KiB | [pg_dbms_metadata_15-1.0.0-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_metadata_15-1.0.0-5PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.7 KiB | [pg_dbms_metadata_15-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_metadata_15-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pg_dbms_metadata_15-1.1.0-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dbms_metadata_15-1.1.0-1PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.1 KiB | [pg_dbms_metadata_15-1.0.0-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dbms_metadata_15-1.0.0-5PGDG.rhel8.10.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.1 KiB | [pg_dbms_metadata_15-1.0.0-2PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_dbms_metadata_15-1.0.0-2PIGSTY.el8.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.5 KiB | [pg_dbms_metadata_15-1.1.0-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_metadata_15-1.1.0-1PGDG.rhel9.9.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.4 KiB | [pg_dbms_metadata_15-1.1.0-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_metadata_15-1.1.0-1PGDG.rhel9.9.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 19.6 KiB | [pg_dbms_metadata_15-1.1.0-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_dbms_metadata_15-1.1.0-1PGDG.rhel10.3.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.5 KiB | [pg_dbms_metadata_15-1.1.0-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_metadata_15-1.1.0-1PGDG.rhel10.3.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 19.6 KiB | [pg_dbms_metadata_15-1.1.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_metadata_15-1.1.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.7 KiB | [pg_dbms_metadata_15-1.0.0-5PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_metadata_15-1.0.0-5PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.5 KiB | [pg_dbms_metadata_15-1.0.0-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_metadata_15-1.0.0-2PGDG.rhel10.2.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pg_dbms_metadata_15-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_metadata_15-1.0.0-2PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pg-dbms-metadata` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.2 KiB | [postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-metadata/postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-dbms-metadata` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.2 KiB | [postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-dbms-metadata/postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-dbms-metadata` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.2 KiB | [postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-metadata/postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-dbms-metadata` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.2 KiB | [postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-dbms-metadata/postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-dbms-metadata` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.5 KiB | [postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-metadata/postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-dbms-metadata` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.5 KiB | [postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-dbms-metadata/postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-dbms-metadata` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.5 KiB | [postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-metadata/postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-dbms-metadata` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.5 KiB | [postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-dbms-metadata/postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-dbms-metadata` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.5 KiB | [postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-metadata/postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-15-pg-dbms-metadata` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.5 KiB | [postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-dbms-metadata/postgresql-15-pg-dbms-metadata_1.0.0-1PIGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HexaCluster/pg_dbms_metadata" title="Repository" icon="github" subtitle="github.com/HexaCluster/pg_dbms_metadata" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_dbms_metadata-1.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_dbms_metadata;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_dbms_metadata;		# install via package name, for the active PG version

pig install pg_dbms_metadata -v 19;   # install for PG 19
pig install pg_dbms_metadata -v 18;   # install for PG 18
pig install pg_dbms_metadata -v 17;   # install for PG 17
pig install pg_dbms_metadata -v 16;   # install for PG 16
pig install pg_dbms_metadata -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_dbms_metadata;
```

## Usage

Sources:

- [Version 1.1.0 README](https://github.com/HexaCluster/pg_dbms_metadata/blob/v1.1.0/README.md)
- [SQL definitions](https://github.com/HexaCluster/pg_dbms_metadata/blob/v1.1.0/sql/pg_dbms_metadata--1.1.0.sql)
- [Upgrade script](https://github.com/HexaCluster/pg_dbms_metadata/blob/v1.1.0/updates/pg_dbms_metadata--1.0.0--1.1.0.sql)
- [Release notes](https://github.com/HexaCluster/pg_dbms_metadata/releases/tag/v1.1.0)

`pg_dbms_metadata` 1.1.0 extracts PostgreSQL object DDL through an Oracle-style API. It is a SQL/PL/pgSQL extension with no preload requirement. It implements the GET functions, not Oracle FETCH or XML interfaces.

### Enabling

```sql
CREATE EXTENSION pg_dbms_metadata;
```

### GET_DDL

Extract DDL of database objects. Supported types: TABLE, VIEW, SEQUENCE, PROCEDURE, FUNCTION, TRIGGER, INDEX, CONSTRAINT, CHECK_CONSTRAINT, REF_CONSTRAINT, TYPE, ENUM.

```sql
SELECT dbms_metadata.get_ddl('TABLE', 'employees', 'public');
SELECT dbms_metadata.get_ddl('VIEW', 'active_users', 'myschema');
SELECT dbms_metadata.get_ddl('FUNCTION', 'calculate_tax', 'public');
SELECT dbms_metadata.get_ddl('INDEX', 'idx_name', 'public');

-- Schema is optional; uses search_path when omitted
SELECT dbms_metadata.get_ddl('SEQUENCE', 'my_seq');
```

### GET_DEPENDENT_DDL

Extract DDL of all dependent objects of a specified type for a base object. Supported: SEQUENCE, TRIGGER, CONSTRAINT, REF_CONSTRAINT, INDEX, ENUM.

```sql
SELECT dbms_metadata.get_dependent_ddl('CONSTRAINT', 'employees', 'public');
SELECT dbms_metadata.get_dependent_ddl('INDEX', 'orders', 'sales');
SELECT dbms_metadata.get_dependent_ddl('TRIGGER', 'accounts');
```

### GET_GRANTED_DDL

Extract SQL statements to recreate granted privileges and roles. Supported: ROLE_GRANT.

```sql
SELECT dbms_metadata.get_granted_ddl('ROLE_GRANT', 'user_test');
```

### SET_TRANSFORM_PARAM

Customize DDL output with session-level transform parameters:

```sql
-- Append SQL terminator (;) to each DDL statement
CALL dbms_metadata.set_transform_param('SQLTERMINATOR', true);

-- Exclude constraints from TABLE DDL
CALL dbms_metadata.set_transform_param('CONSTRAINTS', false);

-- Exclude foreign keys from TABLE DDL
CALL dbms_metadata.set_transform_param('REF_CONSTRAINTS', false);

-- Exclude partitioning clauses
CALL dbms_metadata.set_transform_param('PARTITIONING', false);

-- Exclude storage parameters
CALL dbms_metadata.set_transform_param('STORAGE', false);

-- Reset all params to defaults
CALL dbms_metadata.set_transform_param('DEFAULT', true);
```

### 1.1.0 Behavior and Upgrade

ENUM extraction is now supported and explicitly supplied schemas qualify enum names. For overloaded functions or procedures, an unqualified routine name returns every matching overload; append argument types to select one:

```sql
SELECT dbms_metadata.get_ddl('FUNCTION', 'my_function(char, char)', 'my_schema');
ALTER EXTENSION pg_dbms_metadata UPDATE TO '1.1.0';
```

1.1.0 also fixes quoted table comments and boolean transform defaults. Install the new extension files before the SQL update. Upstream documents PostgreSQL 9.1–18 and tests this release's upgrade path on 12–19; package support is separate.

DDL visibility is not limited to objects owned by the caller. When a schema is omitted, lookup uses `search_path` and requires schema USAGE; TRIGGER and REF_CONSTRAINT calls require an explicit schema. Restrict API access when object definitions are sensitive. Session transform settings affect subsequent extraction calls and can be reset with the DEFAULT transform; this API does not implement a full database backup.
