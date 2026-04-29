---
title: "temporal_tables"
linkTitle: "temporal_tables"
description: "temporal tables"
weight: 1040
categories: ["TIME"]
width: full
---

[**temporal_tables**](https://pgxn.org/dist/temporal_tables/) : temporal tables


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1040** | {{< badge content="temporal_tables" link="https://pgxn.org/dist/temporal_tables/" >}} | {{< ext "temporal_tables" >}} | `1.2.2` | {{< category "TIME" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "timescaledb_toolkit" >}} {{< ext "timescaledb" >}} {{< ext "timeseries" >}} {{< ext "periods" >}} {{< ext "emaj" >}} {{< ext "table_version" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} |

> [!Note] no pg17 on el8/9 pgdg repo


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `temporal_tables` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "temporal_tables_18" "green" >}} {{< bg "17" "temporal_tables_17" "green" >}} {{< bg "16" "temporal_tables_16" "green" >}} {{< bg "15" "temporal_tables_15" "green" >}} {{< bg "14" "temporal_tables_14" "green" >}} | `temporal_tables_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "postgresql-18-temporal-tables" "green" >}} {{< bg "17" "postgresql-17-temporal-tables" "green" >}} {{< bg "16" "postgresql-16-temporal-tables" "green" >}} {{< bg "15" "postgresql-15-temporal-tables" "green" >}} {{< bg "14" "postgresql-14-temporal-tables" "green" >}} | `postgresql-$v-temporal-tables` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_17 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_14 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_17 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_14 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "temporal_tables_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2.2" "temporal_tables_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-18-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-17-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-16-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-15-temporal-tables : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.2" "postgresql-14-temporal-tables : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-temporal-tables : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-temporal-tables : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-temporal-tables : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `temporal_tables_18` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.5 KiB | [temporal_tables_18-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/temporal_tables_18-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `temporal_tables_18` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.2 KiB | [temporal_tables_18-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/temporal_tables_18-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `temporal_tables_18` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.3 KiB | [temporal_tables_18-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/temporal_tables_18-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `temporal_tables_18` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.1 KiB | [temporal_tables_18-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/temporal_tables_18-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `temporal_tables_18` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.5 KiB | [temporal_tables_18-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/temporal_tables_18-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `temporal_tables_18` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.3 KiB | [temporal_tables_18-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/temporal_tables_18-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-temporal-tables` | `1.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.8 KiB | [postgresql-18-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-18-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-temporal-tables` | `1.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.2 KiB | [postgresql-18-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-18-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-temporal-tables` | `1.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.7 KiB | [postgresql-18-temporal-tables_1.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/temporal-tables/postgresql-18-temporal-tables_1.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-temporal-tables` | `1.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.3 KiB | [postgresql-18-temporal-tables_1.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/temporal-tables/postgresql-18-temporal-tables_1.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-temporal-tables` | `1.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.9 KiB | [postgresql-18-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-18-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-temporal-tables` | `1.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.3 KiB | [postgresql-18-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-18-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-temporal-tables` | `1.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.7 KiB | [postgresql-18-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-18-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-temporal-tables` | `1.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.2 KiB | [postgresql-18-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-18-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `temporal_tables_17` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.5 KiB | [temporal_tables_17-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/temporal_tables_17-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `temporal_tables_17` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.3 KiB | [temporal_tables_17-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/temporal_tables_17-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `temporal_tables_17` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.0 KiB | [temporal_tables_17-1.2.2-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/temporal_tables_17-1.2.2-4PGDG.rhel9.x86_64.rpm) |
| `temporal_tables_17` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.5 KiB | [temporal_tables_17-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/temporal_tables_17-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `temporal_tables_17` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [temporal_tables_17-1.2.2-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/temporal_tables_17-1.2.2-4PGDG.rhel9.aarch64.rpm) |
| `temporal_tables_17` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.1 KiB | [temporal_tables_17-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/temporal_tables_17-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `temporal_tables_17` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.5 KiB | [temporal_tables_17-1.2.2-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/temporal_tables_17-1.2.2-5PGDG.rhel10.x86_64.rpm) |
| `temporal_tables_17` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.5 KiB | [temporal_tables_17-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/temporal_tables_17-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `temporal_tables_17` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [temporal_tables_17-1.2.2-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/temporal_tables_17-1.2.2-5PGDG.rhel10.aarch64.rpm) |
| `temporal_tables_17` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.3 KiB | [temporal_tables_17-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/temporal_tables_17-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-temporal-tables` | `1.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.7 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-temporal-tables` | `1.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.3 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-temporal-tables` | `1.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.7 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-temporal-tables` | `1.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.3 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-temporal-tables` | `1.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.4 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-temporal-tables` | `1.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.8 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-temporal-tables` | `1.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.6 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-temporal-tables` | `1.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.2 KiB | [postgresql-17-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-17-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `temporal_tables_16` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.6 KiB | [temporal_tables_16-1.2.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/temporal_tables_16-1.2.2-2PGDG.rhel8.x86_64.rpm) |
| `temporal_tables_16` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.5 KiB | [temporal_tables_16-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/temporal_tables_16-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `temporal_tables_16` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.2 KiB | [temporal_tables_16-1.2.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/temporal_tables_16-1.2.2-2PGDG.rhel8.aarch64.rpm) |
| `temporal_tables_16` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.3 KiB | [temporal_tables_16-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/temporal_tables_16-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `temporal_tables_16` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.5 KiB | [temporal_tables_16-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/temporal_tables_16-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `temporal_tables_16` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.3 KiB | [temporal_tables_16-1.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/temporal_tables_16-1.2.2-1PGDG.rhel9.x86_64.rpm) |
| `temporal_tables_16` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.1 KiB | [temporal_tables_16-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/temporal_tables_16-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `temporal_tables_16` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.4 KiB | [temporal_tables_16-1.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/temporal_tables_16-1.2.2-1PGDG.rhel9.aarch64.rpm) |
| `temporal_tables_16` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.5 KiB | [temporal_tables_16-1.2.2-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/temporal_tables_16-1.2.2-5PGDG.rhel10.x86_64.rpm) |
| `temporal_tables_16` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.5 KiB | [temporal_tables_16-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/temporal_tables_16-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `temporal_tables_16` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [temporal_tables_16-1.2.2-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/temporal_tables_16-1.2.2-5PGDG.rhel10.aarch64.rpm) |
| `temporal_tables_16` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.2 KiB | [temporal_tables_16-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/temporal_tables_16-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-temporal-tables` | `1.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.7 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-temporal-tables` | `1.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.2 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-temporal-tables` | `1.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.7 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-temporal-tables` | `1.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.3 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-temporal-tables` | `1.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.0 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-temporal-tables` | `1.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.4 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-temporal-tables` | `1.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.6 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-temporal-tables` | `1.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.2 KiB | [postgresql-16-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-16-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `temporal_tables_15` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.7 KiB | [temporal_tables_15-1.2.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/temporal_tables_15-1.2.2-2PGDG.rhel8.x86_64.rpm) |
| `temporal_tables_15` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.6 KiB | [temporal_tables_15-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/temporal_tables_15-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `temporal_tables_15` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.2 KiB | [temporal_tables_15-1.2.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/temporal_tables_15-1.2.2-2PGDG.rhel8.aarch64.rpm) |
| `temporal_tables_15` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.4 KiB | [temporal_tables_15-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/temporal_tables_15-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `temporal_tables_15` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.6 KiB | [temporal_tables_15-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/temporal_tables_15-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `temporal_tables_15` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.4 KiB | [temporal_tables_15-1.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/temporal_tables_15-1.2.2-1PGDG.rhel9.x86_64.rpm) |
| `temporal_tables_15` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.1 KiB | [temporal_tables_15-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/temporal_tables_15-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `temporal_tables_15` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.5 KiB | [temporal_tables_15-1.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/temporal_tables_15-1.2.2-1PGDG.rhel9.aarch64.rpm) |
| `temporal_tables_15` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.6 KiB | [temporal_tables_15-1.2.2-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/temporal_tables_15-1.2.2-5PGDG.rhel10.x86_64.rpm) |
| `temporal_tables_15` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.6 KiB | [temporal_tables_15-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/temporal_tables_15-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `temporal_tables_15` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [temporal_tables_15-1.2.2-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/temporal_tables_15-1.2.2-5PGDG.rhel10.aarch64.rpm) |
| `temporal_tables_15` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.4 KiB | [temporal_tables_15-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/temporal_tables_15-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-temporal-tables` | `1.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.7 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-temporal-tables` | `1.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.2 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-temporal-tables` | `1.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.7 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-temporal-tables` | `1.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.3 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-temporal-tables` | `1.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.0 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-temporal-tables` | `1.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.5 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-temporal-tables` | `1.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.7 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-temporal-tables` | `1.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.2 KiB | [postgresql-15-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-15-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `temporal_tables_14` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.7 KiB | [temporal_tables_14-1.2.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/temporal_tables_14-1.2.2-2PGDG.rhel8.x86_64.rpm) |
| `temporal_tables_14` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.6 KiB | [temporal_tables_14-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/temporal_tables_14-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `temporal_tables_14` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.2 KiB | [temporal_tables_14-1.2.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/temporal_tables_14-1.2.2-2PGDG.rhel8.aarch64.rpm) |
| `temporal_tables_14` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.4 KiB | [temporal_tables_14-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/temporal_tables_14-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `temporal_tables_14` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.6 KiB | [temporal_tables_14-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/temporal_tables_14-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `temporal_tables_14` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.4 KiB | [temporal_tables_14-1.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/temporal_tables_14-1.2.2-1PGDG.rhel9.x86_64.rpm) |
| `temporal_tables_14` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.1 KiB | [temporal_tables_14-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/temporal_tables_14-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `temporal_tables_14` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.5 KiB | [temporal_tables_14-1.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/temporal_tables_14-1.2.2-1PGDG.rhel9.aarch64.rpm) |
| `temporal_tables_14` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.6 KiB | [temporal_tables_14-1.2.2-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/temporal_tables_14-1.2.2-5PGDG.rhel10.x86_64.rpm) |
| `temporal_tables_14` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.6 KiB | [temporal_tables_14-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/temporal_tables_14-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `temporal_tables_14` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [temporal_tables_14-1.2.2-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/temporal_tables_14-1.2.2-5PGDG.rhel10.aarch64.rpm) |
| `temporal_tables_14` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.3 KiB | [temporal_tables_14-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/temporal_tables_14-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-temporal-tables` | `1.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.6 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-temporal-tables` | `1.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.2 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-temporal-tables` | `1.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.6 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-temporal-tables` | `1.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.2 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-temporal-tables` | `1.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.0 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-temporal-tables` | `1.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 28.4 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-temporal-tables` | `1.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.7 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-temporal-tables` | `1.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.2 KiB | [postgresql-14-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/temporal-tables/postgresql-14-temporal-tables_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://pgxn.org/dist/temporal_tables/" title="Repository" icon="link" subtitle="pgxn.org/dist/temporal_tables/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="temporal_tables-1.2.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg temporal_tables;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install temporal_tables;		# install via package name, for the active PG version

pig install temporal_tables -v 18;   # install for PG 18
pig install temporal_tables -v 17;   # install for PG 17
pig install temporal_tables -v 16;   # install for PG 16
pig install temporal_tables -v 15;   # install for PG 15
pig install temporal_tables -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION temporal_tables;
```



## Usage

> [temporal_tables: System-period temporal tables for PostgreSQL](https://github.com/arkhipov/temporal_tables)

A temporal table is a table that records the period of time when a row is valid. The system period is a column (or a pair of columns) with a system-maintained value that contains the period of time when a row is valid from a database perspective. When you insert a row into such table, the system automatically generates the values for the start and end of the period. When you update or delete a row from a system-period temporal table, the old row is archived into another table, which is called the history table.

There is [a fantastic tutorial](http://clarkdave.net/2015/02/historical-records-with-postgresql-and-temporal-tables-and-sql-2011/) on using and querying temporal tables in PostgreSQL with this extension.

### Creating a System-Period Temporal Table

The extension uses a general trigger function to maintain system-period temporal table behaviour:

```
versioning(<system_period_column_name>, <history_table_name>, <adjust>)
```

First, create a table and add a system period column:

```sql
CREATE TABLE employees (
  name text NOT NULL PRIMARY KEY,
  department text,
  salary numeric(20, 2)
);

ALTER TABLE employees ADD COLUMN sys_period tstzrange NOT NULL;
```

Then create a history table:

```sql
CREATE TABLE employees_history (LIKE employees);
```

A history table must contain a system period column with the same name and data type as in the original one. If both tables contain a column, the data type must be the same.

Finally, create a trigger to link it with the history table:

```sql
CREATE TRIGGER versioning_trigger
BEFORE INSERT OR UPDATE OR DELETE ON employees
FOR EACH ROW EXECUTE PROCEDURE versioning('sys_period',
                                          'employees_history',
                                          true);
```


## Inserting Data

Inserting data into a system-period temporal table is similar to inserting data into a regular table:

```sql
INSERT INTO employees (name, department, salary)
VALUES ('Bernard Marx', 'Hatchery and Conditioning Centre', 10000);

INSERT INTO employees (name, department, salary)
VALUES ('Lenina Crowne', 'Hatchery and Conditioning Centre', 7000);
```

The start of `sys_period` column represents the time when the row became current, generated by `CURRENT_TIMESTAMP`.


## Updating Data

When a user updates rows, the trigger inserts a copy of the old row into the history table. If a single transaction makes multiple updates to the same row, only one history row is generated:

```sql
UPDATE employees SET salary = 11200 WHERE name = 'Bernard Marx';
```

The history table now contains the previous version:

| name         | department                       | salary | sys_period              |
|--------------|----------------------------------|--------|-------------------------|
| Bernard Marx | Hatchery and Conditioning Centre | 10000  | [2006-08-08, 2007-02-27)|

### Update Conflicts and Time Adjustment

Update conflicts can occur when multiple transactions update the same row. When the `adjust` parameter is set to `true`, the start of `sys_period` is adjusted by adding a small delta (typically 1 microsecond) to avoid failures with SQLSTATE 22000.


## Deleting Data

When a user deletes data, the trigger adds rows to the history table:

```sql
DELETE FROM employees WHERE name = 'Helmholtz Watson';
```


## Advanced Usage

You can set a custom system time for versioning triggers, useful for creating a data warehouse from a system that recorded timestamps:

```sql
SELECT set_system_time('1985-08-08 06:42:00+08');
```

To revert to the default behaviour:

```sql
SELECT set_system_time(NULL);
```

If issued within a transaction that is later aborted, all changes are undone. If committed, changes persist until the end of the session.


## Examples and Hints

### Using Inheritance for History Tables

```sql
CREATE TABLE employees_history (
  name text NOT NULL,
  department text,
  salary numeric(20, 2),
  sys_period tstzrange NOT NULL
);

CREATE TABLE employees (PRIMARY KEY(name)) INHERITS (employees_history);
```

### Pruning History Tables

History tables are always growing. Several pruning strategies:

  1. Periodically delete old data from a history table.
  2. Use partitioning and detach old partitions from a history table.
  3. Retain only the latest N versions of a row.
  4. Prune rows when a corresponding row is deleted from the temporal table.
  5. Prune rows that satisfy specified business rules.

You can also set another tablespace for a history table to move it on cheaper storage.

### Data Audit

You can add triggers to save the user that modified or deleted the current row:

```sql
CREATE FUNCTION employees_modify()
RETURNS TRIGGER AS $$
BEGIN
  NEW.user_modified = SESSION_USER;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER employees_modify
BEFORE INSERT OR UPDATE ON employees
FOR EACH ROW EXECUTE PROCEDURE employees_modify();

CREATE FUNCTION employees_delete()
RETURNS TRIGGER AS $$
BEGIN
  NEW.user_deleted = SESSION_USER;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER employees_delete
BEFORE INSERT ON employees_history
FOR EACH ROW EXECUTE PROCEDURE employees_delete();
```
