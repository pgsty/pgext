---
title: "emaj"
linkTitle: "emaj"
description: "Enables fine-grained write logging and time travel on subsets of the database."
weight: 1050
categories: ["TIME"]
width: full
---

Enables fine-grained write logging and time travel on subsets of the database.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1050** | {{< badge content="emaj" link="https://github.com/dalibo/emaj" >}} | {{< ext "emaj" >}} | `4.7.1` | {{< category "TIME" >}} | {{< license "GPL-3.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "dblink" >}} {{< ext "btree_gist" >}} |
|   **See Also**    | {{< ext "timescaledb_toolkit" >}} {{< ext "timescaledb" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "table_version" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} {{< ext "timeseries" >}} |

> [!Note] max_prepared_transactions


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/emaj" >}} | `4.7.1` | {{< bg "18" "e-maj_18" "green" >}} {{< bg "17" "e-maj_17" "green" >}} {{< bg "16" "e-maj_16" "green" >}} {{< bg "15" "e-maj_15" "green" >}} {{< bg "14" "e-maj_14" "green" >}} {{< bg "13" "e-maj_13" "green" >}} | `e-maj_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/emaj" >}} | `4.7.1` | {{< bg "18" "postgresql-18-emaj" "green" >}} {{< bg "17" "postgresql-17-emaj" "green" >}} {{< bg "16" "postgresql-16-emaj" "green" >}} {{< bg "15" "postgresql-15-emaj" "green" >}} {{< bg "14" "postgresql-14-emaj" "green" >}} {{< bg "13" "postgresql-13-emaj" "green" >}} | `postgresql-$v-emaj` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_13 : AVAIL 10" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_13 : AVAIL 10" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_13 : AVAIL 10" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_13 : AVAIL 10" "blue" >}} |
|    `el10.x86_64`    | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_13 : AVAIL 3" "blue" >}} |
|    `el10.aarch64`    | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_13 : AVAIL 3" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-emaj : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.7.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-13-emaj : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-emaj : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.7.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-13-emaj : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-emaj : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-emaj : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-emaj : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-emaj : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.7.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-13-emaj : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-emaj : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.7.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-13-emaj : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-emaj : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.7.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-13-emaj : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-emaj : MISS 0" "red" >}}      | {{< bg "PIGSTY 4.7.0" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.0" "postgresql-13-emaj : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `e-maj_18` | 4.7.1 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_18` | 4.7.1 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_18` | 4.7.1 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_18` | 4.7.1 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_18` | 4.7.1 | `el10.x86_64` | pgdg | 5.1 MiB | [e-maj_18-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/e-maj_18-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_18` | 4.7.1 | `el10.aarch64` | pgdg | 5.1 MiB | [e-maj_18-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/e-maj_18-4.7.1-1PGDG.rhel10.noarch.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `e-maj_17` | 4.7.1 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.7.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_17-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.6.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_17-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.5.0 | `el8.x86_64` | pgdg | 5.2 MiB | [e-maj_17-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.7.1 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.7.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_17-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.6.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_17-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.5.0 | `el8.aarch64` | pgdg | 5.2 MiB | [e-maj_17-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.7.1 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.7.0 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_17-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.6.0 | `el9.x86_64` | pgdg | 4.4 MiB | [e-maj_17-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.5.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_17-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.7.1 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.7.0 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_17-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.6.0 | `el9.aarch64` | pgdg | 4.4 MiB | [e-maj_17-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.5.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_17-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.7.1 | `el10.x86_64` | pgdg | 5.1 MiB | [e-maj_17-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/e-maj_17-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_17` | 4.7.0 | `el10.x86_64` | pgdg | 5.1 MiB | [e-maj_17-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/e-maj_17-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_17` | 4.6.0 | `el10.x86_64` | pgdg | 4.4 MiB | [e-maj_17-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/e-maj_17-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_17` | 4.7.1 | `el10.aarch64` | pgdg | 5.1 MiB | [e-maj_17-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/e-maj_17-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_17` | 4.7.0 | `el10.aarch64` | pgdg | 5.1 MiB | [e-maj_17-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/e-maj_17-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_17` | 4.6.0 | `el10.aarch64` | pgdg | 4.4 MiB | [e-maj_17-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/e-maj_17-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-17-emaj` | 4.7.0 | `d12.x86_64` | pigsty | 212.0 KiB | [postgresql-17-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-emaj` | 4.7.0 | `d12.aarch64` | pigsty | 212.0 KiB | [postgresql-17-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-emaj` | 4.7.0 | `u22.x86_64` | pigsty | 191.7 KiB | [postgresql-17-emaj_4.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-emaj` | 4.7.0 | `u22.aarch64` | pigsty | 191.8 KiB | [postgresql-17-emaj_4.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-emaj` | 4.7.0 | `u24.x86_64` | pigsty | 191.9 KiB | [postgresql-17-emaj_4.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-emaj` | 4.7.0 | `u24.aarch64` | pigsty | 191.9 KiB | [postgresql-17-emaj_4.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `e-maj_16` | 4.7.1 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.7.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_16-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.6.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_16-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.5.0 | `el8.x86_64` | pgdg | 5.2 MiB | [e-maj_16-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.4.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_16-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.3.1 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_16-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_16-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_16-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_16` | 4.2.0 | `el8.x86_64` | pgdg | 4.5 MiB | [e-maj_16-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_16` | 4.7.1 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.7.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_16-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.6.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_16-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.5.0 | `el8.aarch64` | pgdg | 5.2 MiB | [e-maj_16-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.4.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_16-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.3.1 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_16-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_16-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_16-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_16` | 4.2.0 | `el8.aarch64` | pgdg | 4.5 MiB | [e-maj_16-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_16` | 4.7.1 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.7.0 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.6.0 | `el9.x86_64` | pgdg | 4.4 MiB | [e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.5.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.4.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.3.1 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_16` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.2.0 | `el9.x86_64` | pgdg | 4.1 MiB | [e-maj_16-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_16` | 4.7.1 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.7.0 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.6.0 | `el9.aarch64` | pgdg | 4.4 MiB | [e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.5.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.4.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.3.1 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_16` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.2.0 | `el9.aarch64` | pgdg | 4.1 MiB | [e-maj_16-4.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.2.0-1.rhel9.aarch64.rpm) |
| `e-maj_16` | 4.7.1 | `el10.x86_64` | pgdg | 5.1 MiB | [e-maj_16-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/e-maj_16-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_16` | 4.7.0 | `el10.x86_64` | pgdg | 5.1 MiB | [e-maj_16-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/e-maj_16-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_16` | 4.6.0 | `el10.x86_64` | pgdg | 4.4 MiB | [e-maj_16-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/e-maj_16-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_16` | 4.7.1 | `el10.aarch64` | pgdg | 5.1 MiB | [e-maj_16-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/e-maj_16-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_16` | 4.7.0 | `el10.aarch64` | pgdg | 5.1 MiB | [e-maj_16-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/e-maj_16-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_16` | 4.6.0 | `el10.aarch64` | pgdg | 4.4 MiB | [e-maj_16-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/e-maj_16-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-16-emaj` | 4.7.0 | `d12.x86_64` | pigsty | 212.0 KiB | [postgresql-16-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-emaj` | 4.7.0 | `d12.aarch64` | pigsty | 212.1 KiB | [postgresql-16-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-emaj` | 4.7.0 | `u22.x86_64` | pigsty | 191.7 KiB | [postgresql-16-emaj_4.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-emaj` | 4.7.0 | `u22.aarch64` | pigsty | 191.7 KiB | [postgresql-16-emaj_4.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-emaj` | 4.7.0 | `u24.x86_64` | pigsty | 192.0 KiB | [postgresql-16-emaj_4.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-emaj` | 4.7.0 | `u24.aarch64` | pigsty | 191.9 KiB | [postgresql-16-emaj_4.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `e-maj_15` | 4.7.1 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.7.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.6.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.5.0 | `el8.x86_64` | pgdg | 5.2 MiB | [e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.4.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.3.1 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_15` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.2.0 | `el8.x86_64` | pgdg | 4.5 MiB | [e-maj_15-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_15` | 4.1.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_15-4.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.1.0-1.rhel8.x86_64.rpm) |
| `e-maj_15` | 4.7.1 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.7.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.6.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.5.0 | `el8.aarch64` | pgdg | 5.2 MiB | [e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.4.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.3.1 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_15` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.2.0 | `el8.aarch64` | pgdg | 4.5 MiB | [e-maj_15-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_15` | 4.1.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_15-4.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.1.0-1.rhel8.aarch64.rpm) |
| `e-maj_15` | 4.7.1 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.7.0 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.6.0 | `el9.x86_64` | pgdg | 4.4 MiB | [e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.5.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.4.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.3.1 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_15` | 4.2.0 | `el9.x86_64` | pgdg | 4.1 MiB | [e-maj_15-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_15` | 4.1.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_15-4.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.1.0-1.rhel9.x86_64.rpm) |
| `e-maj_15` | 4.7.1 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.7.0 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.6.0 | `el9.aarch64` | pgdg | 4.4 MiB | [e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.5.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.4.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.3.1 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_15` | 4.2.0 | `el9.aarch64` | pgdg | 4.1 MiB | [e-maj_15-4.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.2.0-1.rhel9.aarch64.rpm) |
| `e-maj_15` | 4.1.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_15-4.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.1.0-1.rhel9.aarch64.rpm) |
| `e-maj_15` | 4.7.1 | `el10.x86_64` | pgdg | 5.1 MiB | [e-maj_15-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/e-maj_15-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_15` | 4.7.0 | `el10.x86_64` | pgdg | 5.1 MiB | [e-maj_15-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/e-maj_15-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_15` | 4.6.0 | `el10.x86_64` | pgdg | 4.4 MiB | [e-maj_15-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/e-maj_15-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_15` | 4.7.1 | `el10.aarch64` | pgdg | 5.1 MiB | [e-maj_15-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/e-maj_15-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_15` | 4.7.0 | `el10.aarch64` | pgdg | 5.1 MiB | [e-maj_15-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/e-maj_15-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_15` | 4.6.0 | `el10.aarch64` | pgdg | 4.4 MiB | [e-maj_15-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/e-maj_15-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-15-emaj` | 4.7.0 | `d12.x86_64` | pigsty | 212.1 KiB | [postgresql-15-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-emaj` | 4.7.0 | `d12.aarch64` | pigsty | 212.1 KiB | [postgresql-15-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-emaj` | 4.7.0 | `u22.x86_64` | pigsty | 191.8 KiB | [postgresql-15-emaj_4.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-emaj` | 4.7.0 | `u22.aarch64` | pigsty | 191.8 KiB | [postgresql-15-emaj_4.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-emaj` | 4.7.0 | `u24.x86_64` | pigsty | 191.8 KiB | [postgresql-15-emaj_4.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-emaj` | 4.7.0 | `u24.aarch64` | pigsty | 191.9 KiB | [postgresql-15-emaj_4.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `e-maj_14` | 4.7.1 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.7.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.6.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.5.0 | `el8.x86_64` | pgdg | 5.2 MiB | [e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.4.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.3.1 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_14` | 4.2.0 | `el8.x86_64` | pgdg | 4.5 MiB | [e-maj_14-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_14` | 4.1.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_14-4.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.1.0-1.rhel8.x86_64.rpm) |
| `e-maj_14` | 4.7.1 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.7.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.6.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.5.0 | `el8.aarch64` | pgdg | 5.2 MiB | [e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.4.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.3.1 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_14` | 4.2.0 | `el8.aarch64` | pgdg | 4.5 MiB | [e-maj_14-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_14` | 4.1.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_14-4.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.1.0-1.rhel8.aarch64.rpm) |
| `e-maj_14` | 4.7.1 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.7.0 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_14-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.6.0 | `el9.x86_64` | pgdg | 4.4 MiB | [e-maj_14-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.5.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_14-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.4.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_14-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.3.1 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_14-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_14` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.2.0 | `el9.x86_64` | pgdg | 4.1 MiB | [e-maj_14-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_14` | 4.1.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_14-4.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.1.0-1.rhel9.x86_64.rpm) |
| `e-maj_14` | 4.7.1 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.7.0 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_14-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.6.0 | `el9.aarch64` | pgdg | 4.4 MiB | [e-maj_14-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.5.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_14-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.4.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_14-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.3.1 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_14-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_14` | 4.2.0 | `el9.aarch64` | pgdg | 4.1 MiB | [e-maj_14-4.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.2.0-1.rhel9.aarch64.rpm) |
| `e-maj_14` | 4.1.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_14-4.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.1.0-1.rhel9.aarch64.rpm) |
| `e-maj_14` | 4.7.1 | `el10.x86_64` | pgdg | 5.1 MiB | [e-maj_14-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/e-maj_14-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_14` | 4.7.0 | `el10.x86_64` | pgdg | 5.1 MiB | [e-maj_14-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/e-maj_14-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_14` | 4.6.0 | `el10.x86_64` | pgdg | 4.4 MiB | [e-maj_14-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/e-maj_14-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_14` | 4.7.1 | `el10.aarch64` | pgdg | 5.1 MiB | [e-maj_14-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/e-maj_14-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_14` | 4.7.0 | `el10.aarch64` | pgdg | 5.1 MiB | [e-maj_14-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/e-maj_14-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_14` | 4.6.0 | `el10.aarch64` | pgdg | 4.4 MiB | [e-maj_14-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/e-maj_14-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-14-emaj` | 4.7.0 | `d12.x86_64` | pigsty | 212.1 KiB | [postgresql-14-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-emaj` | 4.7.0 | `d12.aarch64` | pigsty | 212.0 KiB | [postgresql-14-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-emaj` | 4.7.0 | `u22.x86_64` | pigsty | 191.8 KiB | [postgresql-14-emaj_4.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-emaj` | 4.7.0 | `u22.aarch64` | pigsty | 191.8 KiB | [postgresql-14-emaj_4.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-emaj` | 4.7.0 | `u24.x86_64` | pigsty | 192.0 KiB | [postgresql-14-emaj_4.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-emaj` | 4.7.0 | `u24.aarch64` | pigsty | 191.9 KiB | [postgresql-14-emaj_4.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `e-maj_13` | 4.7.1 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_13-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.7.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_13-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.6.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_13-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.5.0 | `el8.x86_64` | pgdg | 5.2 MiB | [e-maj_13-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.4.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_13-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.3.1 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_13-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_13-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_13` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_13-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.2.0 | `el8.x86_64` | pgdg | 4.5 MiB | [e-maj_13-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_13` | 4.1.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_13-4.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.1.0-1.rhel8.x86_64.rpm) |
| `e-maj_13` | 4.7.1 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_13-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.7.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_13-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.6.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_13-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.5.0 | `el8.aarch64` | pgdg | 5.2 MiB | [e-maj_13-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.4.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_13-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.3.1 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_13-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_13-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_13` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_13-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.2.0 | `el8.aarch64` | pgdg | 4.5 MiB | [e-maj_13-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_13` | 4.1.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_13-4.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.1.0-1.rhel8.aarch64.rpm) |
| `e-maj_13` | 4.7.1 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_13-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.7.0 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_13-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.6.0 | `el9.x86_64` | pgdg | 4.4 MiB | [e-maj_13-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.5.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_13-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.4.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_13-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.3.1 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_13-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_13-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_13-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_13` | 4.2.0 | `el9.x86_64` | pgdg | 4.1 MiB | [e-maj_13-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_13` | 4.1.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_13-4.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.1.0-1.rhel9.x86_64.rpm) |
| `e-maj_13` | 4.7.1 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_13-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.7.0 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_13-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.6.0 | `el9.aarch64` | pgdg | 4.4 MiB | [e-maj_13-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.5.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_13-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.4.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_13-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.3.1 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_13-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_13-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_13` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_13-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.2.0 | `el9.aarch64` | pgdg | 4.1 MiB | [e-maj_13-4.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.2.0-1.rhel9.aarch64.rpm) |
| `e-maj_13` | 4.1.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_13-4.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.1.0-1.rhel9.aarch64.rpm) |
| `e-maj_13` | 4.7.1 | `el10.x86_64` | pgdg | 5.1 MiB | [e-maj_13-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/e-maj_13-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_13` | 4.7.0 | `el10.x86_64` | pgdg | 5.1 MiB | [e-maj_13-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/e-maj_13-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_13` | 4.6.0 | `el10.x86_64` | pgdg | 4.4 MiB | [e-maj_13-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/e-maj_13-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_13` | 4.7.1 | `el10.aarch64` | pgdg | 5.1 MiB | [e-maj_13-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/e-maj_13-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_13` | 4.7.0 | `el10.aarch64` | pgdg | 5.1 MiB | [e-maj_13-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/e-maj_13-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_13` | 4.6.0 | `el10.aarch64` | pgdg | 4.4 MiB | [e-maj_13-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/e-maj_13-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-13-emaj` | 4.7.0 | `d12.x86_64` | pigsty | 212.0 KiB | [postgresql-13-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-13-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-emaj` | 4.7.0 | `d12.aarch64` | pigsty | 212.0 KiB | [postgresql-13-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-13-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-emaj` | 4.7.0 | `u22.x86_64` | pigsty | 191.7 KiB | [postgresql-13-emaj_4.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-13-emaj_4.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-emaj` | 4.7.0 | `u22.aarch64` | pigsty | 191.8 KiB | [postgresql-13-emaj_4.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-13-emaj_4.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-emaj` | 4.7.0 | `u24.x86_64` | pigsty | 192.0 KiB | [postgresql-13-emaj_4.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-13-emaj_4.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-emaj` | 4.7.0 | `u24.aarch64` | pigsty | 191.9 KiB | [postgresql-13-emaj_4.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-13-emaj_4.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dalibo/emaj" title="Repository" icon="github" subtitle="github.com/dalibo/emaj" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="emaj-4.7.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get emaj; # get emaj source code
pig build dep emaj; # install build dependencies
pig build pkg emaj; # build extension rpm or deb
pig build ext emaj; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install emaj; # install by extension name, for the current active PG version
pig ext install emaj; # install via package alias, for the active PG version
pig ext install emaj -v 18;   # install for PG 18
pig ext install emaj -v 17;   # install for PG 17
pig ext install emaj -v 16;   # install for PG 16
pig ext install emaj -v 15;   # install for PG 15
pig ext install emaj -v 14;   # install for PG 14
pig ext install emaj -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION emaj CASCADE SCHEMA emaj;
```

