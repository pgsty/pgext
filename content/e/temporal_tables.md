---
title: "temporal_tables"
linkTitle: "temporal_tables"
description: "temporal tables"
weight: 1040
categories: ["TIME"]
width: full
---

temporal tables


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1040** | {{< badge content="temporal_tables" link="https://pgxn.org/dist/temporal_tables/" >}} | {{< ext "temporal_tables" >}} | `1.2.2` | {{< category "TIME" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "timescaledb_toolkit" >}} {{< ext "timescaledb" >}} {{< ext "timeseries" >}} {{< ext "periods" >}} {{< ext "emaj" >}} {{< ext "table_version" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} |

> [!Note] no pg17 on el8/9 pgdg repo


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/temporal_tables" >}} | `1.2.2` | {{< bg "18" "temporal_tables_18*" "green" >}} {{< bg "17" "temporal_tables_17*" "green" >}} {{< bg "16" "temporal_tables_16*" "green" >}} {{< bg "15" "temporal_tables_15*" "green" >}} {{< bg "14" "temporal_tables_14*" "green" >}} {{< bg "13" "temporal_tables_13*" "green" >}} | `temporal_tables_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/temporal_tables" >}} | `1.2.2` | {{< bg "18" "postgresql-18-temporal-tables" "green" >}} {{< bg "17" "postgresql-17-temporal-tables" "green" >}} {{< bg "16" "postgresql-16-temporal-tables" "green" >}} {{< bg "15" "postgresql-15-temporal-tables" "green" >}} {{< bg "14" "postgresql-14-temporal-tables" "green" >}} {{< bg "13" "postgresql-13-temporal-tables" "green" >}} | `postgresql-$v-temporal-tables` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_17 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_13 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_17 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-temporal-tables : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-13-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-temporal-tables : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-13-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-temporal-tables : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-temporal-tables : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-temporal-tables : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-13-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-temporal-tables : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-13-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-temporal-tables : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-13-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-temporal-tables : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-13-temporal-tables : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `temporal_tables_18` | 1.2.2 | `el8.x86_64` | pigsty | 18.5 KiB | [temporal_tables_18-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/temporal_tables_18-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `temporal_tables_18` | 1.2.2 | `el8.aarch64` | pigsty | 18.2 KiB | [temporal_tables_18-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/temporal_tables_18-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `temporal_tables_18` | 1.2.2 | `el9.x86_64` | pigsty | 18.3 KiB | [temporal_tables_18-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/temporal_tables_18-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `temporal_tables_18` | 1.2.2 | `el9.aarch64` | pigsty | 18.1 KiB | [temporal_tables_18-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/temporal_tables_18-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `temporal_tables_18` | 1.2.2 | `el10.x86_64` | pigsty | 18.5 KiB | [temporal_tables_18-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/temporal_tables_18-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `temporal_tables_18` | 1.2.2 | `el10.aarch64` | pigsty | 18.3 KiB | [temporal_tables_18-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/temporal_tables_18-1.2.2-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `temporal_tables_17` | 1.2.2 | `el8.x86_64` | pigsty | 18.5 KiB | [temporal_tables_17-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/temporal_tables_17-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `temporal_tables_17` | 1.2.2 | `el8.aarch64` | pigsty | 18.3 KiB | [temporal_tables_17-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/temporal_tables_17-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `temporal_tables_17` | 1.2.2 | `el9.x86_64` | pigsty | 18.5 KiB | [temporal_tables_17-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/temporal_tables_17-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `temporal_tables_17` | 1.2.2 | `el9.x86_64` | pgdg | 24.0 KiB | [temporal_tables_17-1.2.2-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/temporal_tables_17-1.2.2-4PGDG.rhel9.x86_64.rpm) |
| `temporal_tables_17` | 1.2.2 | `el9.aarch64` | pigsty | 18.1 KiB | [temporal_tables_17-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/temporal_tables_17-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `temporal_tables_17` | 1.2.2 | `el9.aarch64` | pgdg | 23.3 KiB | [temporal_tables_17-1.2.2-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/temporal_tables_17-1.2.2-4PGDG.rhel9.aarch64.rpm) |
| `temporal_tables_17` | 1.2.2 | `el10.x86_64` | pigsty | 18.5 KiB | [temporal_tables_17-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/temporal_tables_17-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `temporal_tables_17` | 1.2.2 | `el10.x86_64` | pgdg | 24.5 KiB | [temporal_tables_17-1.2.2-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/temporal_tables_17-1.2.2-5PGDG.rhel10.x86_64.rpm) |
| `temporal_tables_17` | 1.2.2 | `el10.aarch64` | pigsty | 18.3 KiB | [temporal_tables_17-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/temporal_tables_17-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `temporal_tables_17` | 1.2.2 | `el10.aarch64` | pgdg | 24.0 KiB | [temporal_tables_17-1.2.2-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/temporal_tables_17-1.2.2-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-temporal-tables` | 1.2.2 | `d12.x86_64` | pigsty | 30.1 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-temporal-tables` | 1.2.2 | `d12.aarch64` | pigsty | 29.4 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-temporal-tables` | 1.2.2 | `u22.x86_64` | pigsty | 31.4 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-temporal-tables` | 1.2.2 | `u22.aarch64` | pigsty | 30.8 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-temporal-tables` | 1.2.2 | `u24.x86_64` | pigsty | 26.6 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-temporal-tables` | 1.2.2 | `u24.aarch64` | pigsty | 26.0 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `temporal_tables_16` | 1.2.2 | `el8.x86_64` | pigsty | 18.5 KiB | [temporal_tables_16-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/temporal_tables_16-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `temporal_tables_16` | 1.2.2 | `el8.x86_64` | pgdg | 23.6 KiB | [temporal_tables_16-1.2.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/temporal_tables_16-1.2.2-2PGDG.rhel8.x86_64.rpm) |
| `temporal_tables_16` | 1.2.2 | `el8.aarch64` | pigsty | 18.3 KiB | [temporal_tables_16-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/temporal_tables_16-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `temporal_tables_16` | 1.2.2 | `el8.aarch64` | pgdg | 23.2 KiB | [temporal_tables_16-1.2.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/temporal_tables_16-1.2.2-2PGDG.rhel8.aarch64.rpm) |
| `temporal_tables_16` | 1.2.2 | `el9.x86_64` | pigsty | 18.5 KiB | [temporal_tables_16-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/temporal_tables_16-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `temporal_tables_16` | 1.2.2 | `el9.x86_64` | pgdg | 23.3 KiB | [temporal_tables_16-1.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/temporal_tables_16-1.2.2-1PGDG.rhel9.x86_64.rpm) |
| `temporal_tables_16` | 1.2.2 | `el9.aarch64` | pigsty | 18.1 KiB | [temporal_tables_16-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/temporal_tables_16-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `temporal_tables_16` | 1.2.2 | `el9.aarch64` | pgdg | 22.4 KiB | [temporal_tables_16-1.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/temporal_tables_16-1.2.2-1PGDG.rhel9.aarch64.rpm) |
| `temporal_tables_16` | 1.2.2 | `el10.x86_64` | pigsty | 18.5 KiB | [temporal_tables_16-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/temporal_tables_16-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `temporal_tables_16` | 1.2.2 | `el10.x86_64` | pgdg | 24.5 KiB | [temporal_tables_16-1.2.2-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/temporal_tables_16-1.2.2-5PGDG.rhel10.x86_64.rpm) |
| `temporal_tables_16` | 1.2.2 | `el10.aarch64` | pigsty | 18.2 KiB | [temporal_tables_16-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/temporal_tables_16-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `temporal_tables_16` | 1.2.2 | `el10.aarch64` | pgdg | 24.0 KiB | [temporal_tables_16-1.2.2-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/temporal_tables_16-1.2.2-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-temporal-tables` | 1.2.2 | `d12.x86_64` | pigsty | 29.7 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-temporal-tables` | 1.2.2 | `d12.aarch64` | pigsty | 29.0 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-temporal-tables` | 1.2.2 | `u22.x86_64` | pigsty | 31.0 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-temporal-tables` | 1.2.2 | `u22.aarch64` | pigsty | 30.4 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-temporal-tables` | 1.2.2 | `u24.x86_64` | pigsty | 26.6 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-temporal-tables` | 1.2.2 | `u24.aarch64` | pigsty | 26.0 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `temporal_tables_15` | 1.2.2 | `el8.x86_64` | pigsty | 18.6 KiB | [temporal_tables_15-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/temporal_tables_15-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `temporal_tables_15` | 1.2.2 | `el8.x86_64` | pgdg | 23.7 KiB | [temporal_tables_15-1.2.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/temporal_tables_15-1.2.2-2PGDG.rhel8.x86_64.rpm) |
| `temporal_tables_15` | 1.2.2 | `el8.aarch64` | pigsty | 18.4 KiB | [temporal_tables_15-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/temporal_tables_15-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `temporal_tables_15` | 1.2.2 | `el8.aarch64` | pgdg | 23.2 KiB | [temporal_tables_15-1.2.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/temporal_tables_15-1.2.2-2PGDG.rhel8.aarch64.rpm) |
| `temporal_tables_15` | 1.2.2 | `el9.x86_64` | pigsty | 18.6 KiB | [temporal_tables_15-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/temporal_tables_15-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `temporal_tables_15` | 1.2.2 | `el9.x86_64` | pgdg | 23.4 KiB | [temporal_tables_15-1.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/temporal_tables_15-1.2.2-1PGDG.rhel9.x86_64.rpm) |
| `temporal_tables_15` | 1.2.2 | `el9.aarch64` | pigsty | 18.1 KiB | [temporal_tables_15-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/temporal_tables_15-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `temporal_tables_15` | 1.2.2 | `el9.aarch64` | pgdg | 22.5 KiB | [temporal_tables_15-1.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/temporal_tables_15-1.2.2-1PGDG.rhel9.aarch64.rpm) |
| `temporal_tables_15` | 1.2.2 | `el10.x86_64` | pigsty | 18.6 KiB | [temporal_tables_15-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/temporal_tables_15-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `temporal_tables_15` | 1.2.2 | `el10.x86_64` | pgdg | 24.6 KiB | [temporal_tables_15-1.2.2-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/temporal_tables_15-1.2.2-5PGDG.rhel10.x86_64.rpm) |
| `temporal_tables_15` | 1.2.2 | `el10.aarch64` | pigsty | 18.4 KiB | [temporal_tables_15-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/temporal_tables_15-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `temporal_tables_15` | 1.2.2 | `el10.aarch64` | pgdg | 24.0 KiB | [temporal_tables_15-1.2.2-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/temporal_tables_15-1.2.2-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-temporal-tables` | 1.2.2 | `d12.x86_64` | pigsty | 29.7 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-temporal-tables` | 1.2.2 | `d12.aarch64` | pigsty | 29.0 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-temporal-tables` | 1.2.2 | `u22.x86_64` | pigsty | 31.0 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-temporal-tables` | 1.2.2 | `u22.aarch64` | pigsty | 30.5 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-temporal-tables` | 1.2.2 | `u24.x86_64` | pigsty | 26.7 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-temporal-tables` | 1.2.2 | `u24.aarch64` | pigsty | 26.0 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `temporal_tables_14` | 1.2.2 | `el8.x86_64` | pigsty | 18.6 KiB | [temporal_tables_14-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/temporal_tables_14-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `temporal_tables_14` | 1.2.2 | `el8.x86_64` | pgdg | 23.7 KiB | [temporal_tables_14-1.2.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/temporal_tables_14-1.2.2-2PGDG.rhel8.x86_64.rpm) |
| `temporal_tables_14` | 1.2.2 | `el8.aarch64` | pigsty | 18.4 KiB | [temporal_tables_14-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/temporal_tables_14-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `temporal_tables_14` | 1.2.2 | `el8.aarch64` | pgdg | 23.2 KiB | [temporal_tables_14-1.2.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/temporal_tables_14-1.2.2-2PGDG.rhel8.aarch64.rpm) |
| `temporal_tables_14` | 1.2.2 | `el9.x86_64` | pigsty | 18.6 KiB | [temporal_tables_14-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/temporal_tables_14-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `temporal_tables_14` | 1.2.2 | `el9.x86_64` | pgdg | 23.4 KiB | [temporal_tables_14-1.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/temporal_tables_14-1.2.2-1PGDG.rhel9.x86_64.rpm) |
| `temporal_tables_14` | 1.2.2 | `el9.aarch64` | pigsty | 18.1 KiB | [temporal_tables_14-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/temporal_tables_14-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `temporal_tables_14` | 1.2.2 | `el9.aarch64` | pgdg | 22.5 KiB | [temporal_tables_14-1.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/temporal_tables_14-1.2.2-1PGDG.rhel9.aarch64.rpm) |
| `temporal_tables_14` | 1.2.2 | `el10.x86_64` | pigsty | 18.6 KiB | [temporal_tables_14-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/temporal_tables_14-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `temporal_tables_14` | 1.2.2 | `el10.x86_64` | pgdg | 24.6 KiB | [temporal_tables_14-1.2.2-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/temporal_tables_14-1.2.2-5PGDG.rhel10.x86_64.rpm) |
| `temporal_tables_14` | 1.2.2 | `el10.aarch64` | pigsty | 18.3 KiB | [temporal_tables_14-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/temporal_tables_14-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `temporal_tables_14` | 1.2.2 | `el10.aarch64` | pgdg | 24.0 KiB | [temporal_tables_14-1.2.2-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/temporal_tables_14-1.2.2-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-temporal-tables` | 1.2.2 | `d12.x86_64` | pigsty | 27.8 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-temporal-tables` | 1.2.2 | `d12.aarch64` | pigsty | 27.0 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-temporal-tables` | 1.2.2 | `u22.x86_64` | pigsty | 29.0 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-temporal-tables` | 1.2.2 | `u22.aarch64` | pigsty | 28.4 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-temporal-tables` | 1.2.2 | `u24.x86_64` | pigsty | 26.6 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-temporal-tables` | 1.2.2 | `u24.aarch64` | pigsty | 26.0 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `temporal_tables_13` | 1.2.2 | `el8.x86_64` | pigsty | 18.5 KiB | [temporal_tables_13-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/temporal_tables_13-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `temporal_tables_13` | 1.2.2 | `el8.x86_64` | pgdg | 23.6 KiB | [temporal_tables_13-1.2.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/temporal_tables_13-1.2.2-2PGDG.rhel8.x86_64.rpm) |
| `temporal_tables_13` | 1.2.2 | `el8.aarch64` | pigsty | 18.3 KiB | [temporal_tables_13-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/temporal_tables_13-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `temporal_tables_13` | 1.2.2 | `el8.aarch64` | pgdg | 23.2 KiB | [temporal_tables_13-1.2.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/temporal_tables_13-1.2.2-2PGDG.rhel8.aarch64.rpm) |
| `temporal_tables_13` | 1.2.2 | `el9.x86_64` | pigsty | 18.7 KiB | [temporal_tables_13-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/temporal_tables_13-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `temporal_tables_13` | 1.2.2 | `el9.x86_64` | pgdg | 23.5 KiB | [temporal_tables_13-1.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/temporal_tables_13-1.2.2-1PGDG.rhel9.x86_64.rpm) |
| `temporal_tables_13` | 1.2.2 | `el9.aarch64` | pigsty | 18.2 KiB | [temporal_tables_13-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/temporal_tables_13-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `temporal_tables_13` | 1.2.2 | `el9.aarch64` | pgdg | 22.5 KiB | [temporal_tables_13-1.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/temporal_tables_13-1.2.2-1PGDG.rhel9.aarch64.rpm) |
| `temporal_tables_13` | 1.2.2 | `el10.x86_64` | pigsty | 18.7 KiB | [temporal_tables_13-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/temporal_tables_13-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `temporal_tables_13` | 1.2.2 | `el10.x86_64` | pgdg | 24.7 KiB | [temporal_tables_13-1.2.2-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/temporal_tables_13-1.2.2-5PGDG.rhel10.x86_64.rpm) |
| `temporal_tables_13` | 1.2.2 | `el10.aarch64` | pigsty | 18.4 KiB | [temporal_tables_13-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/temporal_tables_13-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `temporal_tables_13` | 1.2.2 | `el10.aarch64` | pgdg | 24.1 KiB | [temporal_tables_13-1.2.2-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/temporal_tables_13-1.2.2-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-temporal-tables` | 1.2.2 | `d12.x86_64` | pigsty | 27.8 KiB | [postgresql-13-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-13-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-temporal-tables` | 1.2.2 | `d12.aarch64` | pigsty | 27.4 KiB | [postgresql-13-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-13-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-temporal-tables` | 1.2.2 | `u22.x86_64` | pigsty | 29.0 KiB | [postgresql-13-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-13-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-temporal-tables` | 1.2.2 | `u22.aarch64` | pigsty | 28.3 KiB | [postgresql-13-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-13-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-temporal-tables` | 1.2.2 | `u24.x86_64` | pigsty | 26.8 KiB | [postgresql-13-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-13-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-temporal-tables` | 1.2.2 | `u24.aarch64` | pigsty | 26.4 KiB | [postgresql-13-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-13-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://pgxn.org/dist/temporal_tables/" title="Repository" icon="link" subtitle="pgxn.org/dist/temporal_tables/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="temporal_tables-1.2.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get temporal_tables; # get temporal_tables source code
pig build dep temporal_tables; # install build dependencies
pig build pkg temporal_tables; # build extension rpm or deb
pig build ext temporal_tables; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install temporal_tables; # install by extension name, for the current active PG version
pig ext install temporal_tables; # install via package alias, for the active PG version
pig ext install temporal_tables -v 18;   # install for PG 18
pig ext install temporal_tables -v 17;   # install for PG 17
pig ext install temporal_tables -v 16;   # install for PG 16
pig ext install temporal_tables -v 15;   # install for PG 15
pig ext install temporal_tables -v 14;   # install for PG 14
pig ext install temporal_tables -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION temporal_tables;
```

