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
|   **Requires**    | {{< ext "dblink" >}} {{< ext "btree_gist" >}} |
|   **See Also**    | {{< ext "timescaledb_toolkit" >}} {{< ext "timescaledb" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "table_version" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} {{< ext "timeseries" >}} |

> [!Note] max_prepared_transactions


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `4.7.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `emaj` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.7.1` | {{< bg "18" "e-maj_18" "green" >}} {{< bg "17" "e-maj_17" "green" >}} {{< bg "16" "e-maj_16" "green" >}} {{< bg "15" "e-maj_15" "green" >}} {{< bg "14" "e-maj_14" "green" >}} {{< bg "13" "e-maj_13" "green" >}} | `e-maj_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.7.1` | {{< bg "18" "postgresql-18-emaj" "green" >}} {{< bg "17" "postgresql-17-emaj" "green" >}} {{< bg "16" "postgresql-16-emaj" "green" >}} {{< bg "15" "postgresql-15-emaj" "green" >}} {{< bg "14" "postgresql-14-emaj" "green" >}} {{< bg "13" "postgresql-13-emaj" "green" >}} | `postgresql-$v-emaj` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_13 : AVAIL 10" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_13 : AVAIL 10" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_13 : AVAIL 10" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_13 : AVAIL 10" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.7.1" "e-maj_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.7.1" "e-maj_13 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-13-emaj : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-13-emaj : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-13-emaj : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-13-emaj : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-13-emaj : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-13-emaj : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-13-emaj : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-18-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-17-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-16-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-15-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-14-emaj : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.7.1" "postgresql-13-emaj : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
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
| `e-maj_16` | `4.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_16` | `4.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.1 MiB | [e-maj_16-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/e-maj_16-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_16` | `4.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.4 MiB | [e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.7 MiB | [e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.7 MiB | [e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_16` | `4.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_16` | `4.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/e-maj_16-4.3.0-1PGDG.rhel9.noarch.rpm) |
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
| `e-maj_15` | `4.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_15` | `4.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.5 MiB | [e-maj_15-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_15` | `4.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_15-4.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/e-maj_15-4.1.0-1.rhel8.x86_64.rpm) |
| `e-maj_15` | `4.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.2 MiB | [e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_15` | `4.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_15` | `4.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.5 MiB | [e-maj_15-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_15` | `4.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_15-4.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/e-maj_15-4.1.0-1.rhel8.aarch64.rpm) |
| `e-maj_15` | `4.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.4 MiB | [e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.7 MiB | [e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.7 MiB | [e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_15` | `4.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_15-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/e-maj_15-4.3.0-1PGDG.rhel9.x86_64.rpm) |
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
| `e-maj_14` | `4.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_14` | `4.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.5 MiB | [e-maj_14-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_14` | `4.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_14-4.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/e-maj_14-4.1.0-1.rhel8.x86_64.rpm) |
| `e-maj_14` | `4.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.2 MiB | [e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_14` | `4.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_14` | `4.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/e-maj_14-4.3.0-1PGDG.rhel8.noarch.rpm) |
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
| `e-maj_14` | `4.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_14` | `4.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/e-maj_14-4.3.0-1PGDG.rhel9.noarch.rpm) |
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

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `e-maj_13` | `4.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_13-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_13-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_13-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.2 MiB | [e-maj_13-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 5.3 MiB | [e-maj_13-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_13-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_13-4.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.3.0-1PGDG.rhel8.x86_64.rpm) |
| `e-maj_13` | `4.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_13-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.5 MiB | [e-maj_13-4.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.2.0-1.rhel8.x86_64.rpm) |
| `e-maj_13` | `4.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 4.6 MiB | [e-maj_13-4.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/e-maj_13-4.1.0-1.rhel8.x86_64.rpm) |
| `e-maj_13` | `4.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_13-4.7.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.7.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_13-4.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.7.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_13-4.6.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.6.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.2 MiB | [e-maj_13-4.5.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.5.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 5.3 MiB | [e-maj_13-4.4.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.4.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_13-4.3.1-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.3.1-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_13-4.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.3.0-1PGDG.rhel8.aarch64.rpm) |
| `e-maj_13` | `4.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_13-4.3.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.3.0-1PGDG.rhel8.noarch.rpm) |
| `e-maj_13` | `4.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.5 MiB | [e-maj_13-4.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.2.0-1.rhel8.aarch64.rpm) |
| `e-maj_13` | `4.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 4.6 MiB | [e-maj_13-4.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/e-maj_13-4.1.0-1.rhel8.aarch64.rpm) |
| `e-maj_13` | `4.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_13-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 5.1 MiB | [e-maj_13-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.4 MiB | [e-maj_13-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.7 MiB | [e-maj_13-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.7 MiB | [e-maj_13-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_13-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_13-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_13-4.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.3.0-1PGDG.rhel9.x86_64.rpm) |
| `e-maj_13` | `4.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.1 MiB | [e-maj_13-4.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.2.0-1.rhel9.x86_64.rpm) |
| `e-maj_13` | `4.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 4.2 MiB | [e-maj_13-4.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/e-maj_13-4.1.0-1.rhel9.x86_64.rpm) |
| `e-maj_13` | `4.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_13-4.7.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.7.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 5.1 MiB | [e-maj_13-4.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.7.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.4 MiB | [e-maj_13-4.6.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.6.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.7 MiB | [e-maj_13-4.5.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.5.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.7 MiB | [e-maj_13-4.4.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.4.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_13-4.3.1-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.3.1-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_13-4.3.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.3.0-1PGDG.rhel9.noarch.rpm) |
| `e-maj_13` | `4.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_13-4.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.3.0-1PGDG.rhel9.aarch64.rpm) |
| `e-maj_13` | `4.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.1 MiB | [e-maj_13-4.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.2.0-1.rhel9.aarch64.rpm) |
| `e-maj_13` | `4.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 4.2 MiB | [e-maj_13-4.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/e-maj_13-4.1.0-1.rhel9.aarch64.rpm) |
| `e-maj_13` | `4.7.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 5.1 MiB | [e-maj_13-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/e-maj_13-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_13` | `4.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 5.1 MiB | [e-maj_13-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/e-maj_13-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_13` | `4.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 4.4 MiB | [e-maj_13-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/e-maj_13-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_13` | `4.7.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 5.1 MiB | [e-maj_13-4.7.1-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/e-maj_13-4.7.1-1PGDG.rhel10.noarch.rpm) |
| `e-maj_13` | `4.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 5.1 MiB | [e-maj_13-4.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/e-maj_13-4.7.0-1PGDG.rhel10.noarch.rpm) |
| `e-maj_13` | `4.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 4.4 MiB | [e-maj_13-4.6.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/e-maj_13-4.6.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-13-emaj` | `4.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 214.0 KiB | [postgresql-13-emaj_4.7.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-13-emaj_4.7.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-emaj` | `4.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 214.0 KiB | [postgresql-13-emaj_4.7.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/emaj/postgresql-13-emaj_4.7.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-emaj` | `4.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 213.9 KiB | [postgresql-13-emaj_4.7.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-13-emaj_4.7.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-emaj` | `4.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 213.9 KiB | [postgresql-13-emaj_4.7.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/emaj/postgresql-13-emaj_4.7.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-emaj` | `4.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 193.8 KiB | [postgresql-13-emaj_4.7.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-13-emaj_4.7.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-emaj` | `4.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 193.8 KiB | [postgresql-13-emaj_4.7.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/emaj/postgresql-13-emaj_4.7.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-emaj` | `4.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 193.9 KiB | [postgresql-13-emaj_4.7.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-13-emaj_4.7.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-emaj` | `4.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 193.9 KiB | [postgresql-13-emaj_4.7.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/emaj/postgresql-13-emaj_4.7.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dalibo/emaj" title="Repository" icon="github" subtitle="github.com/dalibo/emaj" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="emaj-4.7.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg emaj;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install emaj;		# install via package name, for the active PG version

pig install emaj -v 18;   # install for PG 18
pig install emaj -v 17;   # install for PG 17
pig install emaj -v 16;   # install for PG 16
pig install emaj -v 15;   # install for PG 15
pig install emaj -v 14;   # install for PG 14
pig install emaj -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION emaj CASCADE; -- requires dblink, btree_gist
```
