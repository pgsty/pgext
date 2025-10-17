---
title: "emaj"
linkTitle: "emaj"
description: "Enables fine-grained write logging and time travel on subsets of the database."
weight: 1050
categories: ["Time"]
width: full
---

Enables fine-grained write logging and time travel on subsets of the database.

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1050** | {{< badge content="emaj" link="https://github.com/dalibo/emaj" >}} | {{< ext "emaj" "emaj" >}} | `4.7.0` | {{< category "TIME" >}} | {{< license "GPL-3.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "dblink" >}} {{< ext "btree_gist" >}} |
|   **See Also**    | {{< ext "timescaledb_toolkit" >}} {{< ext "timescaledb" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "table_version" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} {{< ext "timeseries" >}} |

> [!Note] max_prepared_transactions


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/emaj" >}} | `4.7.0` | {{< badge content="18" color="red" alt="e-maj_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `e-maj_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/emaj" >}} | `4.7.0` | {{< badge content="18" color="red" alt="postgresql-18-emaj" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-emaj` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "e-maj_18" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "e-maj_17" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "e-maj_16" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "e-maj_15" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "e-maj_14" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "e-maj_18" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "e-maj_17" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "e-maj_16" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "e-maj_15" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "e-maj_14" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "e-maj_18" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "e-maj_17" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "e-maj_16" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "e-maj_15" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "e-maj_14" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "e-maj_18" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "e-maj_17" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "e-maj_16" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "e-maj_15" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "e-maj_14" "4.7.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-emaj" >}}     | {{< pkg "postgresql-17-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-emaj" >}}     | {{< pkg "postgresql-17-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-emaj" >}}     | {{< pkg "postgresql-17-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-emaj" >}}     | {{< pkg "postgresql-17-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-emaj" >}}     | {{< pkg "postgresql-17-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-emaj" >}}     | {{< pkg "postgresql-17-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-emaj" "4.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `e-maj_18` | 4.7.1 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_18` | 4.7.1 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/e-maj_18-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_18` | 4.7.1 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_18` | 4.7.1 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/e-maj_18-4.7.1-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `e-maj_17` | 4.7.1 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.7.1 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.7.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_17-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.7.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_17-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.6.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_17-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.6.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_17-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.5.0 | `el8.x86_64` | pgdg | 5.2 MiB | [e-maj_17-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/e-maj_17-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.5.0 | `el8.aarch64` | pgdg | 5.2 MiB | [e-maj_17-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/e-maj_17-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_17` | 4.7.1 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.7.1 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.7.0 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_17-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.7.0 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_17-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.6.0 | `el9.aarch64` | pgdg | 4.4 MiB | [e-maj_17-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.6.0 | `el9.x86_64` | pgdg | 4.4 MiB | [e-maj_17-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.5.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_17-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/e-maj_17-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_17` | 4.5.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_17-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/e-maj_17-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `postgresql-17-emaj` | 4.7.0 | `d12.aarch64` | pigsty | 212.0 KiB | [postgresql-17-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-emaj` | 4.7.0 | `d12.x86_64` | pigsty | 212.0 KiB | [postgresql-17-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-emaj` | 4.7.0 | `u22.aarch64` | pigsty | 191.8 KiB | [postgresql-17-emaj_4.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-emaj` | 4.7.0 | `u22.x86_64` | pigsty | 191.7 KiB | [postgresql-17-emaj_4.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-emaj` | 4.7.0 | `u24.x86_64` | pigsty | 191.9 KiB | [postgresql-17-emaj_4.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-emaj` | 4.7.0 | `u24.aarch64` | pigsty | 191.9 KiB | [postgresql-17-emaj_4.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-17-emaj_4.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `e-maj_16` | 4.7.1 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.7.1 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.7.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_16-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.7.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_16-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.6.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_16-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.6.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_16-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.5.0 | `el8.aarch64` | pgdg | 5.2 MiB | [e-maj_16-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.5.0 | `el8.x86_64` | pgdg | 5.2 MiB | [e-maj_16-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.4.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_16-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.4.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_16-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.3.1 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_16-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.3.1 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_16-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_16-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_16-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_16` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_16-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_16` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_16-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_16` | 4.2.0 | `el8.aarch64` | pgdg | 4.5 MiB | [e-maj_16-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/e-maj_16-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_16` | 4.2.0 | `el8.x86_64` | pgdg | 4.5 MiB | [e-maj_16-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/e-maj_16-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_16` | 4.7.1 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.7.1 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.7.0 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.7.0 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.6.0 | `el9.aarch64` | pgdg | 4.4 MiB | [e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.6.0 | `el9.x86_64` | pgdg | 4.4 MiB | [e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.5.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.5.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.4.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.4.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.3.1 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.3.1 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_16` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_16` | 4.2.0 | `el9.x86_64` | pgdg | 4.1 MiB | [e-maj_16-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_16` | 4.2.0 | `el9.aarch64` | pgdg | 4.1 MiB | [e-maj_16-4.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.2.0-1.rhel9.aarch64.rpm) |
| `postgresql-16-emaj` | 4.7.0 | `d12.x86_64` | pigsty | 212.0 KiB | [postgresql-16-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-emaj` | 4.7.0 | `d12.aarch64` | pigsty | 212.1 KiB | [postgresql-16-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-emaj` | 4.7.0 | `u22.x86_64` | pigsty | 191.7 KiB | [postgresql-16-emaj_4.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-emaj` | 4.7.0 | `u22.aarch64` | pigsty | 191.7 KiB | [postgresql-16-emaj_4.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-emaj` | 4.7.0 | `u24.x86_64` | pigsty | 192.0 KiB | [postgresql-16-emaj_4.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-emaj` | 4.7.0 | `u24.aarch64` | pigsty | 191.9 KiB | [postgresql-16-emaj_4.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-16-emaj_4.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `e-maj_15` | 4.7.1 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.7.1 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.7.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.7.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.6.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.6.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.5.0 | `el8.aarch64` | pgdg | 5.2 MiB | [e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.5.0 | `el8.x86_64` | pgdg | 5.2 MiB | [e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.4.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.4.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.3.1 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.3.1 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_15` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_15` | 4.2.0 | `el8.aarch64` | pgdg | 4.5 MiB | [e-maj_15-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_15` | 4.2.0 | `el8.x86_64` | pgdg | 4.5 MiB | [e-maj_15-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_15` | 4.1.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_15-4.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.1.0-1.rhel8.x86_64.rpm) |
| `e-maj_15` | 4.1.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_15-4.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.1.0-1.rhel8.aarch64.rpm) |
| `e-maj_15` | 4.7.1 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.7.1 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.7.0 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.7.0 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.6.0 | `el9.aarch64` | pgdg | 4.4 MiB | [e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.6.0 | `el9.x86_64` | pgdg | 4.4 MiB | [e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.5.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.5.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.4.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.4.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.3.1 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.3.1 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_15` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_15` | 4.2.0 | `el9.x86_64` | pgdg | 4.1 MiB | [e-maj_15-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_15` | 4.2.0 | `el9.aarch64` | pgdg | 4.1 MiB | [e-maj_15-4.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.2.0-1.rhel9.aarch64.rpm) |
| `e-maj_15` | 4.1.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_15-4.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.1.0-1.rhel9.x86_64.rpm) |
| `e-maj_15` | 4.1.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_15-4.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/e-maj_15-4.1.0-1.rhel9.aarch64.rpm) |
| `postgresql-15-emaj` | 4.7.0 | `d12.x86_64` | pigsty | 212.1 KiB | [postgresql-15-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-emaj` | 4.7.0 | `d12.aarch64` | pigsty | 212.1 KiB | [postgresql-15-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-emaj` | 4.7.0 | `u22.x86_64` | pigsty | 191.8 KiB | [postgresql-15-emaj_4.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-emaj` | 4.7.0 | `u22.aarch64` | pigsty | 191.8 KiB | [postgresql-15-emaj_4.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-emaj` | 4.7.0 | `u24.x86_64` | pigsty | 191.8 KiB | [postgresql-15-emaj_4.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-emaj` | 4.7.0 | `u24.aarch64` | pigsty | 191.9 KiB | [postgresql-15-emaj_4.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-15-emaj_4.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `e-maj_14` | 4.7.1 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.7.1 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.7.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.7.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.6.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.6.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.5.0 | `el8.aarch64` | pgdg | 5.2 MiB | [e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.5.0 | `el8.x86_64` | pgdg | 5.2 MiB | [e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.4.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.4.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.3.1 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.3.1 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_14` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_14` | 4.2.0 | `el8.aarch64` | pgdg | 4.5 MiB | [e-maj_14-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_14` | 4.2.0 | `el8.x86_64` | pgdg | 4.5 MiB | [e-maj_14-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_14` | 4.1.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_14-4.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.1.0-1.rhel8.x86_64.rpm) |
| `e-maj_14` | 4.1.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_14-4.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.1.0-1.rhel8.aarch64.rpm) |
| `e-maj_14` | 4.7.1 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.7.1 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.7.0 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_14-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.7.0 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_14-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.6.0 | `el9.x86_64` | pgdg | 4.4 MiB | [e-maj_14-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.6.0 | `el9.aarch64` | pgdg | 4.4 MiB | [e-maj_14-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.5.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_14-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.5.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_14-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.4.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_14-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.4.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_14-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.3.1 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_14-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.3.1 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_14-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_14` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_14` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_14` | 4.2.0 | `el9.x86_64` | pgdg | 4.1 MiB | [e-maj_14-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_14` | 4.2.0 | `el9.aarch64` | pgdg | 4.1 MiB | [e-maj_14-4.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.2.0-1.rhel9.aarch64.rpm) |
| `e-maj_14` | 4.1.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_14-4.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/e-maj_14-4.1.0-1.rhel9.x86_64.rpm) |
| `e-maj_14` | 4.1.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_14-4.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.1.0-1.rhel9.aarch64.rpm) |
| `postgresql-14-emaj` | 4.7.0 | `d12.aarch64` | pigsty | 212.0 KiB | [postgresql-14-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-emaj` | 4.7.0 | `d12.x86_64` | pigsty | 212.1 KiB | [postgresql-14-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-emaj` | 4.7.0 | `u22.aarch64` | pigsty | 191.8 KiB | [postgresql-14-emaj_4.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-emaj` | 4.7.0 | `u22.x86_64` | pigsty | 191.8 KiB | [postgresql-14-emaj_4.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-emaj` | 4.7.0 | `u24.aarch64` | pigsty | 191.9 KiB | [postgresql-14-emaj_4.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-emaj` | 4.7.0 | `u24.x86_64` | pigsty | 192.0 KiB | [postgresql-14-emaj_4.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-14-emaj_4.7.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `e-maj_13` | 4.7.1 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_13-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.7.1 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_13-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.7.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_13-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.7.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_13-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.6.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_13-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.6.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_13-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.5.0 | `el8.aarch64` | pgdg | 5.2 MiB | [e-maj_13-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.5.0 | `el8.x86_64` | pgdg | 5.2 MiB | [e-maj_13-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.4.0 | `el8.x86_64` | pgdg | 5.3 MiB | [e-maj_13-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.4.0 | `el8.aarch64` | pgdg | 5.3 MiB | [e-maj_13-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.3.1 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_13-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.3.1 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_13-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_13-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_13` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_13-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.3.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_13-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | 4.3.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_13-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_13` | 4.2.0 | `el8.aarch64` | pgdg | 4.5 MiB | [e-maj_13-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_13` | 4.2.0 | `el8.x86_64` | pgdg | 4.5 MiB | [e-maj_13-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_13` | 4.1.0 | `el8.x86_64` | pgdg | 4.6 MiB | [e-maj_13-4.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.1.0-1.rhel8.x86_64.rpm) |
| `e-maj_13` | 4.1.0 | `el8.aarch64` | pgdg | 4.6 MiB | [e-maj_13-4.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.1.0-1.rhel8.aarch64.rpm) |
| `e-maj_13` | 4.7.1 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_13-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.7.1 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_13-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.7.0 | `el9.aarch64` | pgdg | 5.1 MiB | [e-maj_13-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.7.0 | `el9.x86_64` | pgdg | 5.1 MiB | [e-maj_13-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.6.0 | `el9.x86_64` | pgdg | 4.4 MiB | [e-maj_13-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.6.0 | `el9.aarch64` | pgdg | 4.4 MiB | [e-maj_13-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.5.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_13-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.5.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_13-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.4.0 | `el9.x86_64` | pgdg | 4.7 MiB | [e-maj_13-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.4.0 | `el9.aarch64` | pgdg | 4.7 MiB | [e-maj_13-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.3.1 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_13-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.3.1 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_13-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_13-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_13` | 4.3.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_13-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_13-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_13` | 4.3.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_13-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | 4.2.0 | `el9.x86_64` | pgdg | 4.1 MiB | [e-maj_13-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_13` | 4.2.0 | `el9.aarch64` | pgdg | 4.1 MiB | [e-maj_13-4.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.2.0-1.rhel9.aarch64.rpm) |
| `e-maj_13` | 4.1.0 | `el9.x86_64` | pgdg | 4.2 MiB | [e-maj_13-4.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.1.0-1.rhel9.x86_64.rpm) |
| `e-maj_13` | 4.1.0 | `el9.aarch64` | pgdg | 4.2 MiB | [e-maj_13-4.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.1.0-1.rhel9.aarch64.rpm) |
| `postgresql-13-emaj` | 4.7.0 | `d12.aarch64` | pigsty | 212.0 KiB | [postgresql-13-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-13-emaj_4.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-emaj` | 4.7.0 | `d12.x86_64` | pigsty | 212.0 KiB | [postgresql-13-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-13-emaj_4.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-emaj` | 4.7.0 | `u22.aarch64` | pigsty | 191.8 KiB | [postgresql-13-emaj_4.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-13-emaj_4.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-emaj` | 4.7.0 | `u22.x86_64` | pigsty | 191.7 KiB | [postgresql-13-emaj_4.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-13-emaj_4.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-emaj` | 4.7.0 | `u24.x86_64` | pigsty | 192.0 KiB | [postgresql-13-emaj_4.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-13-emaj_4.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-emaj` | 4.7.0 | `u24.aarch64` | pigsty | 191.9 KiB | [postgresql-13-emaj_4.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-13-emaj_4.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dalibo/emaj" title="Repository" icon="github" subtitle="github.com/dalibo/emaj" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="emaj-4.7.0.tar.gz" >}}
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

