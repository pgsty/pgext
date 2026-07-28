---
title: "emaj"
linkTitle: "emaj"
description: "Enables fine-grained write logging and time travel on subsets of the database."
weight: 1050
categories: ["TIME"]
width: full
---

[**emaj**](https://github.com/dalibo/emaj) : Enables fine-grained write logging and time travel on subsets of the database.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1050** | {{< badge content="emaj" link="https://github.com/dalibo/emaj" >}} | {{< ext "emaj" >}} | `5.0.0` | {{< category "TIME" >}} | {{< license "GPL-3.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `emaj` |
|   **Requires**    | {{< ext "btree_gist" >}} {{< ext "dblink" >}} |
|   **See Also**    | {{< ext "timescaledb_toolkit" >}} {{< ext "timescaledb" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "table_version" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} {{< ext "timeseries" >}} |

> [!Note] Requires max_prepared_transactions


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `emaj` | `btree_gist`, `dblink` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.0.0` | {{< bg "18" "emaj_18" "green" >}} {{< bg "17" "emaj_17" "green" >}} {{< bg "16" "emaj_16" "green" >}} {{< bg "15" "emaj_15" "green" >}} {{< bg "14" "emaj_14" "green" >}} | `emaj_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.0.0` | {{< bg "18" "postgresql-18-emaj" "green" >}} {{< bg "17" "postgresql-17-emaj" "green" >}} {{< bg "16" "postgresql-16-emaj" "green" >}} {{< bg "15" "postgresql-15-emaj" "green" >}} {{< bg "14" "postgresql-14-emaj" "green" >}} | `postgresql-$v-emaj` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 5.0.0" "emaj_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 5.0.0" "emaj_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 5.0.0" "emaj_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 5.0.0" "emaj_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 5.0.0" "emaj_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 5.0.0" "emaj_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "emaj_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.0.0" "postgresql-14-emaj : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `emaj_18` | `5.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 314.5 KiB | [emaj_18-5.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/emaj_18-5.0.0-1PIGSTY.el8.noarch.rpm) |
| `emaj_18` | `5.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 314.5 KiB | [emaj_18-5.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/emaj_18-5.0.0-1PIGSTY.el8.noarch.rpm) |
| `emaj_18` | `5.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 219.6 KiB | [emaj_18-5.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/emaj_18-5.0.0-1PIGSTY.el9.noarch.rpm) |
| `emaj_18` | `5.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 219.5 KiB | [emaj_18-5.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/emaj_18-5.0.0-1PIGSTY.el9.noarch.rpm) |
| `emaj_18` | `5.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 219.8 KiB | [emaj_18-5.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/emaj_18-5.0.0-1PIGSTY.el10.noarch.rpm) |
| `emaj_18` | `5.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 219.7 KiB | [emaj_18-5.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/emaj_18-5.0.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-emaj` | `5.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 232.0 KiB | [postgresql-18-emaj_5.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-18-emaj_5.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-emaj` | `5.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 232.0 KiB | [postgresql-18-emaj_5.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-18-emaj_5.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-emaj` | `5.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 232.1 KiB | [postgresql-18-emaj_5.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-18-emaj_5.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-emaj` | `5.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 232.1 KiB | [postgresql-18-emaj_5.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-18-emaj_5.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-emaj` | `5.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 209.9 KiB | [postgresql-18-emaj_5.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-18-emaj_5.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-emaj` | `5.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 209.9 KiB | [postgresql-18-emaj_5.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-18-emaj_5.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-emaj` | `5.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 210.0 KiB | [postgresql-18-emaj_5.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-18-emaj_5.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-emaj` | `5.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 210.0 KiB | [postgresql-18-emaj_5.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-18-emaj_5.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-emaj` | `5.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 209.7 KiB | [postgresql-18-emaj_5.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/emaj/postgresql-18-emaj_5.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-18-emaj` | `5.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 209.7 KiB | [postgresql-18-emaj_5.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/emaj/postgresql-18-emaj_5.0.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `emaj_17` | `5.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 314.5 KiB | [emaj_17-5.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/emaj_17-5.0.0-1PIGSTY.el8.noarch.rpm) |
| `emaj_17` | `5.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 314.5 KiB | [emaj_17-5.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/emaj_17-5.0.0-1PIGSTY.el8.noarch.rpm) |
| `emaj_17` | `5.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 219.3 KiB | [emaj_17-5.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/emaj_17-5.0.0-1PIGSTY.el9.noarch.rpm) |
| `emaj_17` | `5.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 219.3 KiB | [emaj_17-5.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/emaj_17-5.0.0-1PIGSTY.el9.noarch.rpm) |
| `emaj_17` | `5.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 219.8 KiB | [emaj_17-5.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/emaj_17-5.0.0-1PIGSTY.el10.noarch.rpm) |
| `emaj_17` | `5.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 219.7 KiB | [emaj_17-5.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/emaj_17-5.0.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-emaj` | `5.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 232.1 KiB | [postgresql-17-emaj_5.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-17-emaj_5.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-emaj` | `5.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 232.1 KiB | [postgresql-17-emaj_5.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-17-emaj_5.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-emaj` | `5.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 232.1 KiB | [postgresql-17-emaj_5.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-17-emaj_5.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-emaj` | `5.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 232.1 KiB | [postgresql-17-emaj_5.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-17-emaj_5.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-emaj` | `5.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 209.9 KiB | [postgresql-17-emaj_5.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-17-emaj_5.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-emaj` | `5.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 209.9 KiB | [postgresql-17-emaj_5.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-17-emaj_5.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-emaj` | `5.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 210.0 KiB | [postgresql-17-emaj_5.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-17-emaj_5.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-emaj` | `5.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 210.0 KiB | [postgresql-17-emaj_5.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-17-emaj_5.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-emaj` | `5.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 209.7 KiB | [postgresql-17-emaj_5.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/emaj/postgresql-17-emaj_5.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-17-emaj` | `5.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 209.7 KiB | [postgresql-17-emaj_5.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/emaj/postgresql-17-emaj_5.0.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `emaj_16` | `5.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 314.5 KiB | [emaj_16-5.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/emaj_16-5.0.0-1PIGSTY.el8.noarch.rpm) |
| `emaj_16` | `5.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 314.5 KiB | [emaj_16-5.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/emaj_16-5.0.0-1PIGSTY.el8.noarch.rpm) |
| `emaj_16` | `5.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 219.4 KiB | [emaj_16-5.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/emaj_16-5.0.0-1PIGSTY.el9.noarch.rpm) |
| `emaj_16` | `5.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 219.3 KiB | [emaj_16-5.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/emaj_16-5.0.0-1PIGSTY.el9.noarch.rpm) |
| `emaj_16` | `5.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 219.8 KiB | [emaj_16-5.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/emaj_16-5.0.0-1PIGSTY.el10.noarch.rpm) |
| `emaj_16` | `5.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 219.7 KiB | [emaj_16-5.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/emaj_16-5.0.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-emaj` | `5.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 232.1 KiB | [postgresql-16-emaj_5.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-16-emaj_5.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-emaj` | `5.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 232.1 KiB | [postgresql-16-emaj_5.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-16-emaj_5.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-emaj` | `5.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 232.1 KiB | [postgresql-16-emaj_5.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-16-emaj_5.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-emaj` | `5.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 232.1 KiB | [postgresql-16-emaj_5.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-16-emaj_5.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-emaj` | `5.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 210.0 KiB | [postgresql-16-emaj_5.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-16-emaj_5.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-emaj` | `5.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 210.0 KiB | [postgresql-16-emaj_5.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-16-emaj_5.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-emaj` | `5.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 210.0 KiB | [postgresql-16-emaj_5.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-16-emaj_5.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-emaj` | `5.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 210.0 KiB | [postgresql-16-emaj_5.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-16-emaj_5.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-emaj` | `5.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 209.7 KiB | [postgresql-16-emaj_5.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/emaj/postgresql-16-emaj_5.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-16-emaj` | `5.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 209.7 KiB | [postgresql-16-emaj_5.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/emaj/postgresql-16-emaj_5.0.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `emaj_15` | `5.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 314.5 KiB | [emaj_15-5.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/emaj_15-5.0.0-1PIGSTY.el8.noarch.rpm) |
| `emaj_15` | `5.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 314.5 KiB | [emaj_15-5.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/emaj_15-5.0.0-1PIGSTY.el8.noarch.rpm) |
| `emaj_15` | `5.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 219.4 KiB | [emaj_15-5.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/emaj_15-5.0.0-1PIGSTY.el9.noarch.rpm) |
| `emaj_15` | `5.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 219.3 KiB | [emaj_15-5.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/emaj_15-5.0.0-1PIGSTY.el9.noarch.rpm) |
| `emaj_15` | `5.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 219.8 KiB | [emaj_15-5.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/emaj_15-5.0.0-1PIGSTY.el10.noarch.rpm) |
| `emaj_15` | `5.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 219.7 KiB | [emaj_15-5.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/emaj_15-5.0.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-emaj` | `5.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 232.1 KiB | [postgresql-15-emaj_5.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-15-emaj_5.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-emaj` | `5.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 232.1 KiB | [postgresql-15-emaj_5.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-15-emaj_5.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-emaj` | `5.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 232.1 KiB | [postgresql-15-emaj_5.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-15-emaj_5.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-emaj` | `5.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 232.1 KiB | [postgresql-15-emaj_5.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-15-emaj_5.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-emaj` | `5.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 210.0 KiB | [postgresql-15-emaj_5.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-15-emaj_5.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-emaj` | `5.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 210.0 KiB | [postgresql-15-emaj_5.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-15-emaj_5.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-emaj` | `5.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 210.0 KiB | [postgresql-15-emaj_5.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-15-emaj_5.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-emaj` | `5.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 210.0 KiB | [postgresql-15-emaj_5.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-15-emaj_5.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-emaj` | `5.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 209.7 KiB | [postgresql-15-emaj_5.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/emaj/postgresql-15-emaj_5.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-15-emaj` | `5.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 209.7 KiB | [postgresql-15-emaj_5.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/emaj/postgresql-15-emaj_5.0.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `emaj_14` | `5.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 314.5 KiB | [emaj_14-5.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/emaj_14-5.0.0-1PIGSTY.el8.noarch.rpm) |
| `emaj_14` | `5.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 314.5 KiB | [emaj_14-5.0.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/emaj_14-5.0.0-1PIGSTY.el8.noarch.rpm) |
| `emaj_14` | `5.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 219.4 KiB | [emaj_14-5.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/emaj_14-5.0.0-1PIGSTY.el9.noarch.rpm) |
| `emaj_14` | `5.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 219.3 KiB | [emaj_14-5.0.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/emaj_14-5.0.0-1PIGSTY.el9.noarch.rpm) |
| `emaj_14` | `5.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 219.8 KiB | [emaj_14-5.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/emaj_14-5.0.0-1PIGSTY.el10.noarch.rpm) |
| `emaj_14` | `5.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 219.7 KiB | [emaj_14-5.0.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/emaj_14-5.0.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-emaj` | `5.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 232.1 KiB | [postgresql-14-emaj_5.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-14-emaj_5.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-emaj` | `5.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 232.1 KiB | [postgresql-14-emaj_5.0.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-14-emaj_5.0.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-emaj` | `5.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 232.0 KiB | [postgresql-14-emaj_5.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-14-emaj_5.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-emaj` | `5.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 232.0 KiB | [postgresql-14-emaj_5.0.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-14-emaj_5.0.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-emaj` | `5.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 210.0 KiB | [postgresql-14-emaj_5.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-14-emaj_5.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-emaj` | `5.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 210.0 KiB | [postgresql-14-emaj_5.0.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-14-emaj_5.0.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-emaj` | `5.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 209.9 KiB | [postgresql-14-emaj_5.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-14-emaj_5.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-emaj` | `5.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 209.9 KiB | [postgresql-14-emaj_5.0.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-14-emaj_5.0.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-emaj` | `5.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 209.6 KiB | [postgresql-14-emaj_5.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/emaj/postgresql-14-emaj_5.0.0-1PIGSTY~resolute_all.deb) |
| `postgresql-14-emaj` | `5.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 209.6 KiB | [postgresql-14-emaj_5.0.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/emaj/postgresql-14-emaj_5.0.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dalibo/emaj" title="Repository" icon="github" subtitle="github.com/dalibo/emaj" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="emaj-5.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg emaj;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install emaj;		# install via package name, for the active PG version

pig install emaj -v 18;   # install for PG 18
pig install emaj -v 17;   # install for PG 17
pig install emaj -v 16;   # install for PG 16
pig install emaj -v 15;   # install for PG 15
pig install emaj -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION emaj CASCADE; -- requires btree_gist, dblink
```

## Usage

Sources:

- [E-Maj 5.0.0 README](https://github.com/dalibo/emaj/blob/v5.0.0/README.md)
- [E-Maj 5.0.0 changelog](https://github.com/dalibo/emaj/blob/v5.0.0/CHANGES.md)
- [E-Maj quick start](https://github.com/dalibo/emaj/blob/v5.0.0/docs/en/quickStart.rst)
- [E-Maj upgrade guide](https://github.com/dalibo/emaj/blob/v5.0.0/docs/en/upgrade.rst)
- [E-Maj setup guide](https://github.com/dalibo/emaj/blob/v5.0.0/docs/en/setup.rst)

The canonical extension name is `emaj`; E-Maj records table and sequence changes for a coordinated table group and can roll that group back to a named mark. It is useful for repeatable test runs, batch savepoints, change inspection, and targeted recovery, but an E-Maj rollback is not a replacement for PostgreSQL transaction rollback or backup.

### Core Workflow

```sql
CREATE EXTENSION emaj CASCADE;
GRANT emaj_adm TO app_admin;

SELECT emaj.emaj_create_group('my_group', true);
SELECT emaj.emaj_assign_table('app', 'orders', 'my_group');
SELECT emaj.emaj_assign_sequences('app', '.*', '', 'my_group');

SELECT emaj.emaj_start_group('my_group', 'mark_1');
-- Run application changes.
SELECT emaj.emaj_set_mark_group('my_group', 'mark_2');
-- Run more application changes.

SELECT emaj.emaj_rollback_group('my_group', 'mark_1');
SELECT emaj.emaj_stop_group('my_group');
SELECT emaj.emaj_drop_group('my_group');
```

A rollbackable table group can contain tables and sequences from several schemas, but each table must have a primary key. Audit-only groups can record changes for objects that are not rollbackable. Starting and stopping a group takes locks on its application tables, so plan these operations around concurrent traffic.

### Important Objects

- `emaj_create_group` and assignment functions define table groups.
- `emaj_start_group`, `emaj_set_mark_group`, and `emaj_stop_group` manage logging sessions and marks.
- `emaj_rollback_group` performs an unlogged rollback; `emaj_logged_rollback_group` records the compensating changes.
- Multi-group variants operate on arrays of group names at one common point in time.
- Statistics and change-dump functions inspect changes between marks or generate SQL for replay.
- `emaj_set_param` changes or resets an E-Maj parameter without direct writes to the internal parameter table.
- `emaj_drop_extension()` is the supported full-removal helper.

### Version 5.0 Upgrade

For an E-Maj extension installed at version 2.3.1 or later, install the new package files and run:

```sql
ALTER EXTENSION emaj UPDATE;
```

The documented extension upgrade preserves logs and can run while groups remain in the LOGGING state. Review these 5.0 compatibility changes before cutover:

- PostgreSQL 14 through 19 are supported; PostgreSQL 12 and 13 are no longer supported.
- Direct `INSERT`, `UPDATE`, or `DELETE` against `emaj_param` must be replaced by `emaj_set_param`.
- Idempotent start and stop calls have new allow-already-active or allow-already-idle parameters; named-argument callers must review renamed parameters.
- The PHP command-line clients and `emaj_uninstall.sql` were removed.

Installations made with the standalone SQL script do not have the same in-place extension upgrade path; follow the official delete-and-reinstall procedure.

### Requirements and Caveats

The standard `CREATE EXTENSION` path requires superuser privileges and installs `dblink` plus `btree_gist` through `CASCADE`. E-Maj also supports a limited non-superuser script installation, with capability restrictions tied to the installer role.

`max_prepared_transactions` is required only for the parallel rollback client and must be at least the intended session count; changing it requires a restart. Large groups can also require a higher `max_locks_per_transaction`. Treat E-Maj log tables as operational data: size retention deliberately, monitor their growth, and keep ordinary backups for disaster recovery.
