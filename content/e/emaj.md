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
| **1050** | {{< badge content="emaj" link="https://github.com/dalibo/emaj" >}} | {{< ext "emaj" >}} | `4.7.1` | {{< category "TIME" >}} | {{< license "GPL-3.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `emaj` |
|   **Requires**    | {{< ext "dblink" >}} {{< ext "btree_gist" >}} |
|   **See Also**    | {{< ext "timescaledb_toolkit" >}} {{< ext "timescaledb" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "table_version" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} {{< ext "timeseries" >}} |

> [!Note] max_prepared_transactions


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `4.7.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `emaj` | `dblink`, `btree_gist` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.7.1` | {{< bg "18" "e-maj_18" "green" >}} {{< bg "17" "e-maj_17" "green" >}} {{< bg "16" "e-maj_16" "green" >}} {{< bg "15" "e-maj_15" "green" >}} {{< bg "14" "e-maj_14" "green" >}} | `e-maj_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.7.1` | {{< bg "18" "postgresql-18-emaj" "green" >}} {{< bg "17" "postgresql-17-emaj" "green" >}} {{< bg "16" "postgresql-16-emaj" "green" >}} {{< bg "15" "postgresql-15-emaj" "green" >}} {{< bg "14" "postgresql-14-emaj" "green" >}} | `postgresql-$v-emaj` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 10" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 10" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 10" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 10" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-emaj : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-emaj : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `e-maj_18` | `4.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_18` | `4.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_18` | `4.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_18` | `4.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_18` | `4.7.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 5.1 MiB | [e-maj_18-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/e-maj_18-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_18` | `4.7.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 5.1 MiB | [e-maj_18-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/e-maj_18-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `postgresql-18-emaj` | `4.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 213.9 KiB | [postgresql-18-emaj_4.7.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-18-emaj_4.7.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-emaj` | `4.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 213.9 KiB | [postgresql-18-emaj_4.7.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-18-emaj_4.7.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-emaj` | `4.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 213.9 KiB | [postgresql-18-emaj_4.7.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-18-emaj_4.7.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-emaj` | `4.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 213.9 KiB | [postgresql-18-emaj_4.7.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-18-emaj_4.7.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-emaj` | `4.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.8 KiB | [postgresql-18-emaj_4.7.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-18-emaj_4.7.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-emaj` | `4.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 193.8 KiB | [postgresql-18-emaj_4.7.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-18-emaj_4.7.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-emaj` | `4.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 193.9 KiB | [postgresql-18-emaj_4.7.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-18-emaj_4.7.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-emaj` | `4.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 193.9 KiB | [postgresql-18-emaj_4.7.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-18-emaj_4.7.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `e-maj_17` | `4.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | `4.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_17-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | `4.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_17-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | `4.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.2 MiB | [e-maj_17-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | `4.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | `4.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_17-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | `4.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_17-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | `4.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.2 MiB | [e-maj_17-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | `4.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | `4.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_17-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | `4.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.4 MiB | [e-maj_17-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | `4.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.7 MiB | [e-maj_17-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | `4.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | `4.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_17-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | `4.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.4 MiB | [e-maj_17-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | `4.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.7 MiB | [e-maj_17-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | `4.7.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 5.1 MiB | [e-maj_17-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/e-maj_17-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_17` | `4.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 5.1 MiB | [e-maj_17-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/e-maj_17-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_17` | `4.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.4 MiB | [e-maj_17-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/e-maj_17-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_17` | `4.7.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 5.1 MiB | [e-maj_17-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/e-maj_17-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_17` | `4.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 5.1 MiB | [e-maj_17-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/e-maj_17-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_17` | `4.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.4 MiB | [e-maj_17-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/e-maj_17-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-17-emaj` | `4.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 214.0 KiB | [postgresql-17-emaj_4.7.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-17-emaj_4.7.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-emaj` | `4.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 214.0 KiB | [postgresql-17-emaj_4.7.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-17-emaj_4.7.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-emaj` | `4.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 214.0 KiB | [postgresql-17-emaj_4.7.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-17-emaj_4.7.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-emaj` | `4.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 214.0 KiB | [postgresql-17-emaj_4.7.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-17-emaj_4.7.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-emaj` | `4.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.8 KiB | [postgresql-17-emaj_4.7.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-17-emaj_4.7.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-emaj` | `4.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 193.8 KiB | [postgresql-17-emaj_4.7.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-17-emaj_4.7.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-emaj` | `4.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 193.9 KiB | [postgresql-17-emaj_4.7.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-17-emaj_4.7.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-emaj` | `4.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 193.9 KiB | [postgresql-17-emaj_4.7.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-17-emaj_4.7.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `e-maj_16` | `4.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_16-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_16-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.2 MiB | [e-maj_16-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_16-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_16-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_16-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_16` | `4.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_16-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.5 MiB | [e-maj_16-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_16` | `4.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_16-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_16-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.2 MiB | [e-maj_16-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_16-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_16-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_16-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | `4.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_16-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_16` | `4.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.5 MiB | [e-maj_16-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_16` | `4.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.4 MiB | [e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.7 MiB | [e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.7 MiB | [e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_16` | `4.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.1 MiB | [e-maj_16-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_16` | `4.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.4 MiB | [e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.7 MiB | [e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.7 MiB | [e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_16` | `4.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.1 MiB | [e-maj_16-4.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.2.0-1.rhel9.aarch64.rpm) |
| `e-maj_16` | `4.7.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 5.1 MiB | [e-maj_16-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/e-maj_16-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_16` | `4.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 5.1 MiB | [e-maj_16-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/e-maj_16-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_16` | `4.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.4 MiB | [e-maj_16-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/e-maj_16-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_16` | `4.7.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 5.1 MiB | [e-maj_16-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/e-maj_16-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_16` | `4.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 5.1 MiB | [e-maj_16-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/e-maj_16-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_16` | `4.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.4 MiB | [e-maj_16-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/e-maj_16-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-16-emaj` | `4.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 214.0 KiB | [postgresql-16-emaj_4.7.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-16-emaj_4.7.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-emaj` | `4.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 214.0 KiB | [postgresql-16-emaj_4.7.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-16-emaj_4.7.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-emaj` | `4.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 214.0 KiB | [postgresql-16-emaj_4.7.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-16-emaj_4.7.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-emaj` | `4.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 214.0 KiB | [postgresql-16-emaj_4.7.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-16-emaj_4.7.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-emaj` | `4.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.8 KiB | [postgresql-16-emaj_4.7.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-16-emaj_4.7.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-emaj` | `4.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 193.8 KiB | [postgresql-16-emaj_4.7.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-16-emaj_4.7.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-emaj` | `4.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 193.8 KiB | [postgresql-16-emaj_4.7.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-16-emaj_4.7.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-emaj` | `4.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 193.8 KiB | [postgresql-16-emaj_4.7.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-16-emaj_4.7.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `e-maj_15` | `4.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.2 MiB | [e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_15` | `4.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.5 MiB | [e-maj_15-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_15` | `4.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_15-4.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.1.0-1.rhel8.x86_64.rpm) |
| `e-maj_15` | `4.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.2 MiB | [e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_15` | `4.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.5 MiB | [e-maj_15-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_15` | `4.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_15-4.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.1.0-1.rhel8.aarch64.rpm) |
| `e-maj_15` | `4.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.4 MiB | [e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.7 MiB | [e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.7 MiB | [e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_15` | `4.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.1 MiB | [e-maj_15-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_15` | `4.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_15-4.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.1.0-1.rhel9.x86_64.rpm) |
| `e-maj_15` | `4.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.4 MiB | [e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.7 MiB | [e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.7 MiB | [e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_15` | `4.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.1 MiB | [e-maj_15-4.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.2.0-1.rhel9.aarch64.rpm) |
| `e-maj_15` | `4.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_15-4.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.1.0-1.rhel9.aarch64.rpm) |
| `e-maj_15` | `4.7.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 5.1 MiB | [e-maj_15-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/e-maj_15-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_15` | `4.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 5.1 MiB | [e-maj_15-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/e-maj_15-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_15` | `4.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.4 MiB | [e-maj_15-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/e-maj_15-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_15` | `4.7.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 5.1 MiB | [e-maj_15-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/e-maj_15-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_15` | `4.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 5.1 MiB | [e-maj_15-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/e-maj_15-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_15` | `4.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.4 MiB | [e-maj_15-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/e-maj_15-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-15-emaj` | `4.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 213.9 KiB | [postgresql-15-emaj_4.7.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-15-emaj_4.7.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-emaj` | `4.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 213.9 KiB | [postgresql-15-emaj_4.7.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-15-emaj_4.7.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-emaj` | `4.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 213.9 KiB | [postgresql-15-emaj_4.7.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-15-emaj_4.7.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-emaj` | `4.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 213.9 KiB | [postgresql-15-emaj_4.7.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-15-emaj_4.7.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-emaj` | `4.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.8 KiB | [postgresql-15-emaj_4.7.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-15-emaj_4.7.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-emaj` | `4.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 193.8 KiB | [postgresql-15-emaj_4.7.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-15-emaj_4.7.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-emaj` | `4.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 194.0 KiB | [postgresql-15-emaj_4.7.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-15-emaj_4.7.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-emaj` | `4.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 194.0 KiB | [postgresql-15-emaj_4.7.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-15-emaj_4.7.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `e-maj_14` | `4.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.2 MiB | [e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_14` | `4.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.5 MiB | [e-maj_14-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_14` | `4.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_14-4.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.1.0-1.rhel8.x86_64.rpm) |
| `e-maj_14` | `4.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.2 MiB | [e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_14` | `4.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.5 MiB | [e-maj_14-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_14` | `4.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_14-4.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.1.0-1.rhel8.aarch64.rpm) |
| `e-maj_14` | `4.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_14-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.4 MiB | [e-maj_14-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.7 MiB | [e-maj_14-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.7 MiB | [e-maj_14-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_14-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_14` | `4.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.1 MiB | [e-maj_14-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_14` | `4.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_14-4.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.1.0-1.rhel9.x86_64.rpm) |
| `e-maj_14` | `4.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_14-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.4 MiB | [e-maj_14-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.7 MiB | [e-maj_14-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.7 MiB | [e-maj_14-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_14-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | `4.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_14` | `4.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.1 MiB | [e-maj_14-4.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.2.0-1.rhel9.aarch64.rpm) |
| `e-maj_14` | `4.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_14-4.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.1.0-1.rhel9.aarch64.rpm) |
| `e-maj_14` | `4.7.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 5.1 MiB | [e-maj_14-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/e-maj_14-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_14` | `4.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 5.1 MiB | [e-maj_14-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/e-maj_14-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_14` | `4.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.4 MiB | [e-maj_14-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/e-maj_14-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_14` | `4.7.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 5.1 MiB | [e-maj_14-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/e-maj_14-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_14` | `4.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 5.1 MiB | [e-maj_14-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/e-maj_14-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_14` | `4.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.4 MiB | [e-maj_14-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/e-maj_14-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-14-emaj` | `4.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 214.0 KiB | [postgresql-14-emaj_4.7.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-14-emaj_4.7.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-emaj` | `4.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 214.0 KiB | [postgresql-14-emaj_4.7.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-14-emaj_4.7.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-emaj` | `4.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 214.0 KiB | [postgresql-14-emaj_4.7.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-14-emaj_4.7.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-emaj` | `4.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 214.0 KiB | [postgresql-14-emaj_4.7.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-14-emaj_4.7.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-emaj` | `4.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.8 KiB | [postgresql-14-emaj_4.7.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-14-emaj_4.7.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-emaj` | `4.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 193.8 KiB | [postgresql-14-emaj_4.7.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-14-emaj_4.7.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-emaj` | `4.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 193.9 KiB | [postgresql-14-emaj_4.7.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-14-emaj_4.7.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-emaj` | `4.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 193.9 KiB | [postgresql-14-emaj_4.7.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-14-emaj_4.7.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dalibo/emaj" title="Repository" icon="github" subtitle="github.com/dalibo/emaj" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="emaj-4.7.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg emaj;		# build deb
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
CREATE EXTENSION emaj CASCADE; -- requires dblink, btree_gist
```



## Usage

> [E-Maj: logs and rollbacks table content changes](https://github.com/dalibo/emaj)
>
> [**Documentation**](https://emaj.readthedocs.io/en/latest/) | [Emaj_web GUI](https://github.com/dalibo/emaj_web)

E-Maj logs table changes (Inserts, Updates, Deletes, Truncates) performed on one or several sets of tables, and can efficiently cancel these changes if needed, resetting a tables set to a predefined stable state.

In **development environments**, it helps testing applications by providing an easy way to rollback all updates generated by program execution, and replay processing as many times as needed.

In **production environments**, it provides:

- History of changes performed on tables for examination
- Inter-batch savepoints on groups of tables
- Easy "restore" of table groups to a stable state without stopping the cluster
- Multiple savepoints during batch windows, each usable as a restore point


## Concepts

### Tables Group

A tables group is a set of application tables that live at the same rhythm — their content can be restored as a whole if needed. A group can include tables and sequences across different schemas. Each group operates in one of two states: **LOGGING** or **IDLE**, and can be designated as:

- **ROLLBACKABLE** (standard) — supports both logging and rollback
- **AUDIT_ONLY** — allows change recording without rollback capability, even for tables without primary keys or UNLOGGED tables

### Mark

A mark represents a snapshot moment in a tables group's lifecycle, capturing a stable state across all group members. Marks have unique names within a group.

### Rollback

Rollback operations restore tables and sequences to their state when a specific mark was established:

- **Unlogged rollback** — cancelled updates leave no trace
- **Logged rollback** — cancellations are recorded, allowing subsequent reversal

Note: E-Maj rollback differs fundamentally from PostgreSQL's native transaction rollback.


## Main Functions

### Start a Tables Group

```sql
SELECT emaj.emaj_start_group('my_group', 'mark_1');
```

Activates update recording. The group must be in IDLE state. Automatically creates an initial mark.

### Set an Intermediate Mark

```sql
SELECT emaj.emaj_set_mark_group('my_group', 'mark_2');
```

Records a point-in-time snapshot of the application state. The group must be in LOGGING state.

### Rollback a Tables Group

Unlogged rollback (restores tables, no trace of cancellation):

```sql
SELECT * FROM emaj.emaj_rollback_group('my_group', 'mark_1');
```

Logged rollback (permits reverting the rollback itself):

```sql
SELECT * FROM emaj.emaj_logged_rollback_group('my_group', 'mark_1');
```

Both support the `'_EMAJ_LAST_MARK_'` keyword for targeting the most recent mark.

### Stop a Tables Group

```sql
SELECT emaj.emaj_stop_group('my_group');
```

Deactivates logging, changing the group state from LOGGING to IDLE.


## Multi-Group Operations

Functions support batch operations on multiple groups simultaneously:

```sql
SELECT emaj.emaj_start_groups('{"group1","group2"}', 'multi_mark');
SELECT emaj.emaj_set_mark_groups('{"group1","group2"}', 'common_mark');
SELECT * FROM emaj.emaj_rollback_groups('{"group1","group2"}', 'common_mark');
SELECT emaj.emaj_stop_groups('{"group1","group2"}');
```


## Examining Changes

E-Maj provides functions to count and examine data content changes between marks, and to generate SQL scripts that replay logged changes. This is useful for auditing and debugging.


## Emaj_web

**Emaj_web** is a PHP-based web GUI tool for user-friendly E-Maj administration. It is available on [GitHub](https://github.com/dalibo/emaj_web) and described in the [documentation](https://emaj.readthedocs.io/en/latest/webOverview.html).
