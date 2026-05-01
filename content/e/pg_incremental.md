---
title: "pg_incremental"
linkTitle: "pg_incremental"
description: "Incremental Processing by Crunchy Data"
weight: 2850
categories: ["FEAT"]
width: full
---

[**pg_incremental**](https://github.com/CrunchyData/pg_incremental) : Incremental Processing by Crunchy Data


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2850** | {{< badge content="pg_incremental" link="https://github.com/CrunchyData/pg_incremental" >}} | {{< ext "pg_incremental" >}} | `1.5.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} |

> [!Note] pg_cron is optional since v1.3 and only required for scheduled pipelines.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_incremental` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.0` | {{< bg "18" "pg_incremental_18" "green" >}} {{< bg "17" "pg_incremental_17" "green" >}} {{< bg "16" "pg_incremental_16" "green" >}} {{< bg "15" "pg_incremental_15" "red" >}} {{< bg "14" "pg_incremental_14" "red" >}} | `pg_incremental_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.0` | {{< bg "18" "postgresql-18-pg-incremental" "green" >}} {{< bg "17" "postgresql-17-pg-incremental" "green" >}} {{< bg "16" "postgresql-16-pg-incremental" "green" >}} {{< bg "15" "postgresql-15-pg-incremental" "red" >}} {{< bg "14" "postgresql-14-pg-incremental" "red" >}} | `postgresql-$v-pg-incremental` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_16 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_incremental_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_incremental_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_16 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_incremental_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_incremental_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_16 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_incremental_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_incremental_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_16 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_incremental_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_incremental_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_16 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_incremental_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_incremental_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.0" "pg_incremental_16 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_incremental_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_incremental_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-incremental : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-incremental : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-incremental : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-incremental : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-incremental : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-incremental : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-incremental : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-incremental : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-incremental : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-incremental : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-incremental : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-incremental : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-incremental : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-incremental : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-incremental : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-incremental : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-incremental : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-incremental : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.0" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-incremental : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-incremental : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_incremental_18` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.5 KiB | [pg_incremental_18-1.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_incremental_18-1.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.6 KiB | [pg_incremental_18-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_incremental_18-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [pg_incremental_18-1.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_incremental_18-1.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.0 KiB | [pg_incremental_18-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_incremental_18-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [pg_incremental_18-1.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_incremental_18-1.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.1 KiB | [pg_incremental_18-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_incremental_18-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.6 KiB | [pg_incremental_18-1.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_incremental_18-1.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.3 KiB | [pg_incremental_18-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_incremental_18-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.4 KiB | [pg_incremental_18-1.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_incremental_18-1.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.8 KiB | [pg_incremental_18-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_incremental_18-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.9 KiB | [pg_incremental_18-1.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_incremental_18-1.5.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.1 KiB | [pg_incremental_18-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_incremental_18-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-incremental` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 55.9 KiB | [postgresql-18-pg-incremental_1.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-incremental` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 54.8 KiB | [postgresql-18-pg-incremental_1.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-incremental` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 55.9 KiB | [postgresql-18-pg-incremental_1.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-incremental` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 54.8 KiB | [postgresql-18-pg-incremental_1.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-incremental` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 57.4 KiB | [postgresql-18-pg-incremental_1.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-incremental` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 56.6 KiB | [postgresql-18-pg-incremental_1.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-incremental` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 56.3 KiB | [postgresql-18-pg-incremental_1.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-incremental` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 55.6 KiB | [postgresql-18-pg-incremental_1.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-incremental` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 56.0 KiB | [postgresql-18-pg-incremental_1.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-incremental` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 55.5 KiB | [postgresql-18-pg-incremental_1.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_incremental_17` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.5 KiB | [pg_incremental_17-1.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_incremental_17-1.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.6 KiB | [pg_incremental_17-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_incremental_17-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [pg_incremental_17-1.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_incremental_17-1.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.0 KiB | [pg_incremental_17-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_incremental_17-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [pg_incremental_17-1.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_incremental_17-1.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.1 KiB | [pg_incremental_17-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_incremental_17-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.6 KiB | [pg_incremental_17-1.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_incremental_17-1.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.5 KiB | [pg_incremental_17-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_incremental_17-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.4 KiB | [pg_incremental_17-1.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_incremental_17-1.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.8 KiB | [pg_incremental_17-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_incremental_17-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.8 KiB | [pg_incremental_17-1.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_incremental_17-1.5.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.1 KiB | [pg_incremental_17-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_incremental_17-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-incremental` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 55.9 KiB | [postgresql-17-pg-incremental_1.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-incremental` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 54.8 KiB | [postgresql-17-pg-incremental_1.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-incremental` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 55.9 KiB | [postgresql-17-pg-incremental_1.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-incremental` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 54.9 KiB | [postgresql-17-pg-incremental_1.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-incremental` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 62.4 KiB | [postgresql-17-pg-incremental_1.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-incremental` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 61.6 KiB | [postgresql-17-pg-incremental_1.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-incremental` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 56.3 KiB | [postgresql-17-pg-incremental_1.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-incremental` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 55.6 KiB | [postgresql-17-pg-incremental_1.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-incremental` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 56.0 KiB | [postgresql-17-pg-incremental_1.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-incremental` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 55.6 KiB | [postgresql-17-pg-incremental_1.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_incremental_16` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.5 KiB | [pg_incremental_16-1.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_incremental_16-1.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.7 KiB | [pg_incremental_16-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_incremental_16-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.1 KiB | [pg_incremental_16-1.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_incremental_16-1.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.0 KiB | [pg_incremental_16-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_incremental_16-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [pg_incremental_16-1.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_incremental_16-1.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.1 KiB | [pg_incremental_16-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_incremental_16-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.6 KiB | [pg_incremental_16-1.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_incremental_16-1.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.5 KiB | [pg_incremental_16-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_incremental_16-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.4 KiB | [pg_incremental_16-1.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_incremental_16-1.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.8 KiB | [pg_incremental_16-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_incremental_16-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.8 KiB | [pg_incremental_16-1.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_incremental_16-1.5.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.1 KiB | [pg_incremental_16-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_incremental_16-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-incremental` | `1.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 55.9 KiB | [postgresql-16-pg-incremental_1.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-incremental` | `1.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 54.7 KiB | [postgresql-16-pg-incremental_1.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-incremental` | `1.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 55.9 KiB | [postgresql-16-pg-incremental_1.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-incremental` | `1.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 54.8 KiB | [postgresql-16-pg-incremental_1.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-incremental` | `1.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 62.4 KiB | [postgresql-16-pg-incremental_1.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-incremental` | `1.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 61.6 KiB | [postgresql-16-pg-incremental_1.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-incremental` | `1.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 56.4 KiB | [postgresql-16-pg-incremental_1.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-incremental` | `1.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 55.7 KiB | [postgresql-16-pg-incremental_1.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-incremental` | `1.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 56.0 KiB | [postgresql-16-pg-incremental_1.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-incremental` | `1.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 55.6 KiB | [postgresql-16-pg-incremental_1.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrunchyData/pg_incremental" title="Repository" icon="github" subtitle="github.com/CrunchyData/pg_incremental" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_incremental-1.5.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_incremental;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_incremental;		# install via package name, for the active PG version

pig install pg_incremental -v 18;   # install for PG 18
pig install pg_incremental -v 17;   # install for PG 17
pig install pg_incremental -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_incremental;
```

## Usage

- Sources: [README](https://github.com/CrunchyData/pg_incremental/blob/main/README.md), [v1.5.0 release](https://github.com/CrunchyData/pg_incremental/releases/tag/v1.5.0)

`pg_incremental` defines exactly-once incremental pipelines for append-only tables and file feeds. Upstream documents three pipeline types: sequence, time-interval, and file-list.

### Install And Scheduling Model

The upstream README still documents `pg_cron`-backed scheduling and installs with:

```sql
CREATE EXTENSION pg_incremental CASCADE;
```

Pipelines run immediately when created unless `execute_immediately := false`, then continue on a `pg_cron` schedule. The README notes that each scheduled execution appears in `cron.job_run_details` even when no new data is available.

### Sequence Pipelines

Use sequence pipelines to process safe ranges of sequence values:

```sql
SELECT incremental.create_sequence_pipeline('event-aggregation', 'events', $$
  INSERT INTO events_agg
  SELECT date_trunc('day', event_time), count(*)
  FROM events
  WHERE event_id BETWEEN $1 AND $2
  GROUP BY 1
  ON CONFLICT (day) DO UPDATE
  SET event_count = events_agg.event_count + excluded.event_count
$$);
```

The README documents `max_batch_size` for limiting how many sequence IDs are processed per run.

### Time-Interval Pipelines

Use time windows when the command should receive `$1` and `$2` as a passed interval:

```sql
SELECT incremental.create_time_interval_pipeline('event-aggregation', '1 day', $$
  INSERT INTO events_agg
  SELECT event_time::date, count(DISTINCT event_id)
  FROM events
  WHERE event_time >= $1 AND event_time < $2
  GROUP BY 1
$$);
```

For export-style jobs, the README documents `batched := false` so each interval runs separately.

### File-List Pipelines

Use file-list pipelines to process newly discovered files:

```sql
SELECT incremental.create_file_list_pipeline('event-import', 's3://mybucket/events/*.csv', $$
  SELECT import_events($1)
$$);
```

The v1.5.0 release adds `max_batches_per_run` to file-list pipelines. The README documents `incremental.skip_file()` for permanently marking a bad file as processed.

### Operations And Monitoring

The README documents:

- `CALL incremental.execute_pipeline(name)`: run once if new work exists.
- `SELECT incremental.reset_pipeline(name)`: reset progress.
- `SELECT incremental.drop_pipeline(name)`: remove a pipeline.
- Views and tables such as `incremental.sequence_pipelines`, `incremental.time_interval_pipelines`, `incremental.file_list_pipelines`, and `incremental.processed_files`.

The v1.5.0 release note also calls out a `DROP EXTENSION` fix for environments where `pg_cron` is not present.
