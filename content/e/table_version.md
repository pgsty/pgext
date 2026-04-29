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
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.11.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `table_version` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.11.1` | {{< bg "18" "table_version_18" "green" >}} {{< bg "17" "table_version_17" "green" >}} {{< bg "16" "table_version_16" "green" >}} {{< bg "15" "table_version_15" "green" >}} {{< bg "14" "table_version_14" "green" >}} | `table_version_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.11.1` | {{< bg "18" "postgresql-18-table-version" "green" >}} {{< bg "17" "postgresql-17-table-version" "green" >}} {{< bg "16" "postgresql-16-table-version" "green" >}} {{< bg "15" "postgresql-15-table-version" "green" >}} {{< bg "14" "postgresql-14-table-version" "green" >}} | `postgresql-$v-table-version` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.11.1" "table_version_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.11.1" "table_version_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.11.1" "table_version_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "table_version_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "table_version_14 : AVAIL 3" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.11.1" "table_version_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.11.1" "table_version_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.11.1" "table_version_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.11.1" "table_version_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.11.1" "table_version_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.11.1" "table_version_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.11.1" "table_version_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "table_version_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.11.1" "table_version_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.11.1" "table_version_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.11.1" "table_version_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.11.1" "table_version_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.11.1" "table_version_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.11.1" "table_version_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.11.1" "table_version_14 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.11.1" "table_version_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.11.1" "table_version_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.11.1" "table_version_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-18-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-17-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-16-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-15-table-version : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.11.1" "postgresql-14-table-version : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-table-version : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-table-version : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-table-version : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-table-version : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-table-version : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-table-version : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-table-version : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-table-version : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-table-version : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-table-version : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_version_18` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [table_version_18-1.11.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/table_version_18-1.11.1-2PGDG.rhel8.noarch.rpm) |
| `table_version_18` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.1 KiB | [table_version_18-1.11.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_version_18-1.11.1-1PIGSTY.el8.x86_64.rpm) |
| `table_version_18` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.9 KiB | [table_version_18-1.11.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/table_version_18-1.11.1-2PGDG.rhel8.noarch.rpm) |
| `table_version_18` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.1 KiB | [table_version_18-1.11.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_version_18-1.11.1-1PIGSTY.el8.aarch64.rpm) |
| `table_version_18` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.2 KiB | [table_version_18-1.11.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/table_version_18-1.11.1-2PGDG.rhel9.noarch.rpm) |
| `table_version_18` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.8 KiB | [table_version_18-1.11.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_version_18-1.11.1-1PIGSTY.el9.x86_64.rpm) |
| `table_version_18` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.2 KiB | [table_version_18-1.11.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/table_version_18-1.11.1-2PGDG.rhel9.noarch.rpm) |
| `table_version_18` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.8 KiB | [table_version_18-1.11.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_version_18-1.11.1-1PIGSTY.el9.aarch64.rpm) |
| `table_version_18` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [table_version_18-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/table_version_18-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_18` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.0 KiB | [table_version_18-1.11.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/table_version_18-1.11.1-1PIGSTY.el10.x86_64.rpm) |
| `table_version_18` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [table_version_18-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/table_version_18-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_18` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.9 KiB | [table_version_18-1.11.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/table_version_18-1.11.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-table-version` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.1 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.1 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.1 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.9 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.9 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-table-version` | `1.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-18-table-version_1.11.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-18-table-version_1.11.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_version_17` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.1 KiB | [table_version_17-1.11.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_version_17-1.11.1-1PIGSTY.el8.x86_64.rpm) |
| `table_version_17` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [table_version_17-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/table_version_17-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_17` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.1 KiB | [table_version_17-1.11.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_version_17-1.11.1-1PIGSTY.el8.aarch64.rpm) |
| `table_version_17` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.8 KiB | [table_version_17-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/table_version_17-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_17` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.8 KiB | [table_version_17-1.11.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_version_17-1.11.1-1PIGSTY.el9.x86_64.rpm) |
| `table_version_17` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [table_version_17-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/table_version_17-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_17` | `1.10.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [table_version_17-1.10.3-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/table_version_17-1.10.3-3PGDG.rhel9.noarch.rpm) |
| `table_version_17` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.8 KiB | [table_version_17-1.11.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_version_17-1.11.1-1PIGSTY.el9.aarch64.rpm) |
| `table_version_17` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.2 KiB | [table_version_17-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/table_version_17-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_17` | `1.10.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [table_version_17-1.10.3-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/table_version_17-1.10.3-3PGDG.rhel9.noarch.rpm) |
| `table_version_17` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [table_version_17-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/table_version_17-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_17` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.0 KiB | [table_version_17-1.11.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/table_version_17-1.11.1-1PIGSTY.el10.x86_64.rpm) |
| `table_version_17` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [table_version_17-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/table_version_17-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_17` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.9 KiB | [table_version_17-1.11.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/table_version_17-1.11.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-table-version` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.0 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.0 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.1 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.9 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.9 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-table-version` | `1.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-17-table-version_1.11.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-17-table-version_1.11.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_version_16` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.1 KiB | [table_version_16-1.11.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_version_16-1.11.1-1PIGSTY.el8.x86_64.rpm) |
| `table_version_16` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [table_version_16-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/table_version_16-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_16` | `1.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.3 KiB | [table_version_16-1.10.3-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/table_version_16-1.10.3-3PGDG.rhel8.x86_64.rpm) |
| `table_version_16` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.1 KiB | [table_version_16-1.11.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_version_16-1.11.1-1PIGSTY.el8.aarch64.rpm) |
| `table_version_16` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.8 KiB | [table_version_16-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/table_version_16-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_16` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.8 KiB | [table_version_16-1.11.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_version_16-1.11.1-1PIGSTY.el9.x86_64.rpm) |
| `table_version_16` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [table_version_16-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/table_version_16-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_16` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.8 KiB | [table_version_16-1.11.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_version_16-1.11.1-1PIGSTY.el9.aarch64.rpm) |
| `table_version_16` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.2 KiB | [table_version_16-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/table_version_16-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_16` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [table_version_16-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/table_version_16-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_16` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.0 KiB | [table_version_16-1.11.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/table_version_16-1.11.1-1PIGSTY.el10.x86_64.rpm) |
| `table_version_16` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [table_version_16-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/table_version_16-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_16` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.9 KiB | [table_version_16-1.11.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/table_version_16-1.11.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-table-version` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.1 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.1 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.1 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.9 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.9 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-table-version` | `1.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-16-table-version_1.11.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-16-table-version_1.11.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_version_15` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.1 KiB | [table_version_15-1.11.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_version_15-1.11.1-1PIGSTY.el8.x86_64.rpm) |
| `table_version_15` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [table_version_15-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/table_version_15-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_15` | `1.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.1 KiB | [table_version_15-1.10.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/table_version_15-1.10.3-1.rhel8.x86_64.rpm) |
| `table_version_15` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.1 KiB | [table_version_15-1.11.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_version_15-1.11.1-1PIGSTY.el8.aarch64.rpm) |
| `table_version_15` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.8 KiB | [table_version_15-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/table_version_15-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_15` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.8 KiB | [table_version_15-1.11.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_version_15-1.11.1-1PIGSTY.el9.x86_64.rpm) |
| `table_version_15` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [table_version_15-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/table_version_15-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_15` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.8 KiB | [table_version_15-1.11.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_version_15-1.11.1-1PIGSTY.el9.aarch64.rpm) |
| `table_version_15` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.2 KiB | [table_version_15-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/table_version_15-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_15` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [table_version_15-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/table_version_15-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_15` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.0 KiB | [table_version_15-1.11.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/table_version_15-1.11.1-1PIGSTY.el10.x86_64.rpm) |
| `table_version_15` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [table_version_15-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/table_version_15-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_15` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.9 KiB | [table_version_15-1.11.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/table_version_15-1.11.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-table-version` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.1 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.1 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.1 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.9 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.9 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-table-version` | `1.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-15-table-version_1.11.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-15-table-version_1.11.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_version_14` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.1 KiB | [table_version_14-1.11.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_version_14-1.11.1-1PIGSTY.el8.x86_64.rpm) |
| `table_version_14` | `1.11.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [table_version_14-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/table_version_14-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_14` | `1.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [table_version_14-1.9.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/table_version_14-1.9.0-1.rhel8.x86_64.rpm) |
| `table_version_14` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.1 KiB | [table_version_14-1.11.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_version_14-1.11.1-1PIGSTY.el8.aarch64.rpm) |
| `table_version_14` | `1.11.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.8 KiB | [table_version_14-1.11.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/table_version_14-1.11.1-1PGDG.rhel8.noarch.rpm) |
| `table_version_14` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.8 KiB | [table_version_14-1.11.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_version_14-1.11.1-1PIGSTY.el9.x86_64.rpm) |
| `table_version_14` | `1.11.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [table_version_14-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/table_version_14-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_14` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.8 KiB | [table_version_14-1.11.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_version_14-1.11.1-1PIGSTY.el9.aarch64.rpm) |
| `table_version_14` | `1.11.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [table_version_14-1.11.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/table_version_14-1.11.1-1PGDG.rhel9.noarch.rpm) |
| `table_version_14` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [table_version_14-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/table_version_14-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_14` | `1.11.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.0 KiB | [table_version_14-1.11.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/table_version_14-1.11.1-1PIGSTY.el10.x86_64.rpm) |
| `table_version_14` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.7 KiB | [table_version_14-1.11.1-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/table_version_14-1.11.1-2PGDG.rhel10.noarch.rpm) |
| `table_version_14` | `1.11.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.9 KiB | [table_version_14-1.11.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/table_version_14-1.11.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-table-version` | `1.11.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.0 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.0 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.1 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.1 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.9 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.9 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.5 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-table-version` | `1.11.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-14-table-version_1.11.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/table-version/postgresql-14-table-version_1.11.1-1PIGSTY~noble_arm64.deb) |

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

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION table_version CASCADE; -- requires plpgsql
```



## Usage

> [table_version: PostgreSQL table versioning extension](https://github.com/linz/postgresql-tableversion)

PostgreSQL table versioning extension, recording row modifications and its history. The extension provides APIs for accessing snapshots of a table at certain revisions and the difference generated between any two given revisions. It uses a PL/PgSQL trigger based system to record and provide access to row revisions.

### Quick Start

```sql
CREATE EXTENSION table_version;

CREATE SCHEMA foo;
SET search_path TO foo, public;

CREATE TABLE foo.bar (
    id INTEGER NOT NULL PRIMARY KEY,
    baz TEXT
);

-- Enable versioning
SELECT table_version.ver_enable_versioning('foo', 'bar');

-- Create a revision and insert data
SELECT table_version.ver_create_revision('Insert data');
INSERT INTO foo.bar (id, baz) VALUES
  (1, 'foo bar 1'),
  (2, 'foo bar 2'),
  (3, 'foo bar 3');
SELECT table_version.ver_complete_revision();

-- Show differences between revisions
SELECT * FROM table_version.ver_get_foo_bar_diff(1001, 1002);
```


## How It Works

When a table is versioned, the original table data is left untouched and a new revision table is created with all the same fields plus `_revision_created` and `_revision_expired` fields. A row-level trigger is set up on the original table to record every insert, update and delete in the revision data table. A statement-level trigger is set up to forbid TRUNCATE.

### Table Prerequisites

- The table must have a unique non-composite integer, bigint, text or varchar column
- The table must not be temporary


## Auto Revisions

If you don't want to call `ver_create_revision` and `ver_complete_revision` explicitly, auto-revision mode groups edits by transactions:

```sql
SELECT table_version.ver_enable_versioning('foo', 'bar');

BEGIN;
INSERT INTO foo.bar (id, baz) VALUES (1, 'foo bar 1');
INSERT INTO foo.bar (id, baz) VALUES (2, 'foo bar 2');
COMMIT;

BEGIN;
UPDATE foo.bar SET baz = 'foo bar 1 edit' WHERE id = 1;
COMMIT;

SELECT * FROM table_version.foo_bar_revision;
```

The revision message will be automatically created based on the transaction ID.


## Replicate Data Using Table Differences

To maintain a copy of table data on a remote system:

```sql
-- 1. Determine which tables are versioned
SELECT * FROM table_version.ver_get_versioned_tables();

-- 2. Get the base revision
SELECT table_version.ver_get_table_base_revision('foo', 'bar');

-- 3. Create a base snapshot
CREATE TABLE foo_bar_copy AS
SELECT * FROM table_version.ver_get_foo_bar_revision(
    table_version.ver_get_table_base_revision('foo', 'bar')
);

-- 4. Get differences to apply incremental updates
SELECT * FROM table_version.ver_get_foo_bar_diff(
    my_last_revision,
    table_version.ver_get_table_last_revision('foo', 'bar')
);
```


## Security Model

- Anyone can create revisions
- Revisions can only be completed by their creators
- Only those who have ownership privileges on a table can enable/disable versioning
- Only empty revisions can be deleted
- Only the creator of a revision can delete it

Note: Disabling versioning on a table results in all history for that table being deleted.


## Key Functions

| Function | Description |
|----------|-------------|
| `ver_enable_versioning(schema, table)` | Enable versioning on a table |
| `ver_disable_versioning(schema, table)` | Disable versioning and remove history |
| `ver_create_revision(comment)` | Create a new revision |
| `ver_complete_revision()` | Mark current revision as complete |
| `ver_get_<schema>_<table>_diff(rev1, rev2)` | Get differences between two revisions |
| `ver_get_<schema>_<table>_revision(rev)` | Get snapshot at a specific revision |
| `ver_get_versioned_tables()` | List all versioned tables |
| `ver_get_last_revision()` | Get the last revision number |
