---
title: "table_version"
linkTitle: "table_version"
description: "PostgreSQL table versioning extension"
weight: 1060
categories: ["TIME"]
width: full
---

[**table_version**](https://github.com/linz/postgresql-tableversion) : PostgreSQL table versioning extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1060** | {{< badge content="table_version" link="https://github.com/linz/postgresql-tableversion" >}} | {{< ext "table_version" >}} | `1.11.1` | {{< category "TIME" >}} | {{< license "BSD 3-Clause" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `table_version` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "pg_cron" >}} {{< ext "timescaledb" >}} {{< ext "timescaledb_toolkit" >}} {{< ext "timeseries" >}} {{< ext "pg_task" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.11.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `table_version` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.11.1` | {{< bg "18" "table_version_18" "green" >}} {{< bg "17" "table_version_17" "green" >}} {{< bg "16" "table_version_16" "green" >}} {{< bg "15" "table_version_15" "green" >}} {{< bg "14" "table_version_14" "green" >}} {{< bg "13" "table_version_13" "green" >}} | `table_version_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.11.1` | {{< bg "18" "postgresql-18-table-version" "green" >}} {{< bg "17" "postgresql-17-table-version" "green" >}} {{< bg "16" "postgresql-16-table-version" "green" >}} {{< bg "15" "postgresql-15-table-version" "green" >}} {{< bg "14" "postgresql-14-table-version" "green" >}} {{< bg "13" "postgresql-13-table-version" "green" >}} | `postgresql-$v-table-version` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.11.1" "table_version_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_13 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.11.1" "table_version_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.11.1" "table_version_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.11.1" "table_version_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.11.1" "table_version_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.11.1" "table_version_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-13-table-version : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-13-table-version : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-13-table-version : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-13-table-version : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-13-table-version : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-13-table-version : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-13-table-version : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-13-table-version : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_version_18` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [table_version_18-1.11.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/table_version_18-1.11.1-2PGDG.rhel8.noarch.rpm) |
| `table_version_18` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.9 KiB | [table_version_18-1.11.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/table_version_18-1.11.1-2PGDG.rhel8.noarch.rpm) |
| `table_version_18` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.2 KiB | [table_version_18-1.11.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/table_version_18-1.11.1-2PGDG.rhel9.noarch.rpm) |
| `table_version_18` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.2 KiB | [table_version_18-1.11.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/table_version_18-1.11.1-2PGDG.rhel9.noarch.rpm) |
| `table_version_18` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [table_version_18-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/table_version_18-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_18` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [table_version_18-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/table_version_18-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `postgresql-18-table-version` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.1 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.0 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.1 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.9 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.9 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_version_17` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [table_version_17-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/table_version_17-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_17` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.0 KiB | [table_version_17-1.11.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_version_17-1.11.0-1PIGSTY.el8.x86_64.rpm) |
| `table_version_17` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.8 KiB | [table_version_17-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/table_version_17-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_17` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.9 KiB | [table_version_17-1.11.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_version_17-1.11.0-1PIGSTY.el8.aarch64.rpm) |
| `table_version_17` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [table_version_17-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/table_version_17-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_17` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.7 KiB | [table_version_17-1.11.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_version_17-1.11.0-1PIGSTY.el9.x86_64.rpm) |
| `table_version_17` | `1.10.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [table_version_17-1.10.3-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/table_version_17-1.10.3-3PGDG.rhel9.noarch.rpm) |
| `table_version_17` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.2 KiB | [table_version_17-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/table_version_17-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_17` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.6 KiB | [table_version_17-1.11.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_version_17-1.11.0-1PIGSTY.el9.aarch64.rpm) |
| `table_version_17` | `1.10.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [table_version_17-1.10.3-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/table_version_17-1.10.3-3PGDG.rhel9.noarch.rpm) |
| `table_version_17` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [table_version_17-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/table_version_17-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_17` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [table_version_17-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/table_version_17-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `postgresql-17-table-version` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.1 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.0 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.1 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.9 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.9 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_version_16` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [table_version_16-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/table_version_16-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_16` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.0 KiB | [table_version_16-1.11.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_version_16-1.11.0-1PIGSTY.el8.x86_64.rpm) |
| `table_version_16` | `1.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.3 KiB | [table_version_16-1.10.3-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/table_version_16-1.10.3-3PGDG.rhel8.x86_64.rpm) |
| `table_version_16` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.8 KiB | [table_version_16-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/table_version_16-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_16` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.9 KiB | [table_version_16-1.11.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_version_16-1.11.0-1PIGSTY.el8.aarch64.rpm) |
| `table_version_16` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [table_version_16-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/table_version_16-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_16` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.7 KiB | [table_version_16-1.11.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_version_16-1.11.0-1PIGSTY.el9.x86_64.rpm) |
| `table_version_16` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.2 KiB | [table_version_16-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/table_version_16-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_16` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.6 KiB | [table_version_16-1.11.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_version_16-1.11.0-1PIGSTY.el9.aarch64.rpm) |
| `table_version_16` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [table_version_16-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/table_version_16-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_16` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [table_version_16-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/table_version_16-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `postgresql-16-table-version` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.1 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.0 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.1 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.9 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.9 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_version_15` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [table_version_15-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/table_version_15-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_15` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.0 KiB | [table_version_15-1.11.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_version_15-1.11.0-1PIGSTY.el8.x86_64.rpm) |
| `table_version_15` | `1.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.1 KiB | [table_version_15-1.10.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/table_version_15-1.10.3-1.rhel8.x86_64.rpm) |
| `table_version_15` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.8 KiB | [table_version_15-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/table_version_15-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_15` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.9 KiB | [table_version_15-1.11.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_version_15-1.11.0-1PIGSTY.el8.aarch64.rpm) |
| `table_version_15` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [table_version_15-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/table_version_15-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_15` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.7 KiB | [table_version_15-1.11.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_version_15-1.11.0-1PIGSTY.el9.x86_64.rpm) |
| `table_version_15` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.2 KiB | [table_version_15-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/table_version_15-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_15` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.6 KiB | [table_version_15-1.11.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_version_15-1.11.0-1PIGSTY.el9.aarch64.rpm) |
| `table_version_15` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [table_version_15-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/table_version_15-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_15` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [table_version_15-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/table_version_15-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `postgresql-15-table-version` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.1 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.0 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.1 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.9 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.9 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_version_14` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [table_version_14-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/table_version_14-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_14` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.0 KiB | [table_version_14-1.11.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_version_14-1.11.0-1PIGSTY.el8.x86_64.rpm) |
| `table_version_14` | `1.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [table_version_14-1.9.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/table_version_14-1.9.0-1.rhel8.x86_64.rpm) |
| `table_version_14` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.8 KiB | [table_version_14-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/table_version_14-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_14` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.9 KiB | [table_version_14-1.11.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_version_14-1.11.0-1PIGSTY.el8.aarch64.rpm) |
| `table_version_14` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [table_version_14-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/table_version_14-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_14` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.7 KiB | [table_version_14-1.11.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_version_14-1.11.0-1PIGSTY.el9.x86_64.rpm) |
| `table_version_14` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [table_version_14-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/table_version_14-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_14` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.6 KiB | [table_version_14-1.11.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_version_14-1.11.0-1PIGSTY.el9.aarch64.rpm) |
| `table_version_14` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [table_version_14-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/table_version_14-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_14` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [table_version_14-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/table_version_14-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `postgresql-14-table-version` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.1 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.0 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.1 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.9 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.9 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_version_13` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [table_version_13-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/table_version_13-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_13` | `1.11.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.0 KiB | [table_version_13-1.11.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_version_13-1.11.0-1PIGSTY.el8.x86_64.rpm) |
| `table_version_13` | `1.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [table_version_13-1.9.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/table_version_13-1.9.0-1.rhel8.x86_64.rpm) |
| `table_version_13` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.8 KiB | [table_version_13-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/table_version_13-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_13` | `1.11.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.9 KiB | [table_version_13-1.11.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_version_13-1.11.0-1PIGSTY.el8.aarch64.rpm) |
| `table_version_13` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [table_version_13-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/table_version_13-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_13` | `1.11.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.7 KiB | [table_version_13-1.11.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_version_13-1.11.0-1PIGSTY.el9.x86_64.rpm) |
| `table_version_13` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [table_version_13-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/table_version_13-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_13` | `1.11.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.6 KiB | [table_version_13-1.11.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_version_13-1.11.0-1PIGSTY.el9.aarch64.rpm) |
| `table_version_13` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [table_version_13-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/table_version_13-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_13` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [table_version_13-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/table_version_13-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `postgresql-13-table-version` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.1 KiB | [postgresql-13-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-13-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-table-version` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.1 KiB | [postgresql-13-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-13-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-table-version` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-13-table-version_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-13-table-version_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-table-version` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.1 KiB | [postgresql-13-table-version_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-13-table-version_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-table-version` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.9 KiB | [postgresql-13-table-version_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-13-table-version_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-table-version` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.9 KiB | [postgresql-13-table-version_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-13-table-version_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-table-version` | `1.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-13-table-version_1.11.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-13-table-version_1.11.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-table-version` | `1.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-13-table-version_1.11.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-13-table-version_1.11.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/linz/postgresql-tableversion" title="Repository" icon="github" subtitle="github.com/linz/postgresql-tableversion" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresql-tableversion-1.11.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg table_version;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install table_version;		# install via package name, for the active PG version

pig install table_version -v 18;   # install for PG 18
pig install table_version -v 17;   # install for PG 17
pig install table_version -v 16;   # install for PG 16
pig install table_version -v 15;   # install for PG 15
pig install table_version -v 14;   # install for PG 14
pig install table_version -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION table_version CASCADE; -- requires plpgsql
```
