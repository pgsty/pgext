---
title: "firebird_fdw"
linkTitle: "firebird_fdw"
description: "Foreign data wrapper for Firebird"
weight: 8750
categories: ["FDW"]
width: full
---

Foreign data wrapper for Firebird


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8750** | {{< badge content="firebird_fdw" link="https://github.com/ibarwick/firebird_fdw" >}} | {{< ext "firebird_fdw" >}} | `1.4.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "wrappers" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} {{< ext "postgres_fdw" >}} |

> [!Note] pg18 breaks


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/firebird_fdw" >}} | `1.4.0` | {{< bg "18" "firebird_fdw_18" "red" >}} {{< bg "17" "firebird_fdw_17" "green" >}} {{< bg "16" "firebird_fdw_16" "green" >}} {{< bg "15" "firebird_fdw_15" "green" >}} {{< bg "14" "firebird_fdw_14" "green" >}} | `firebird_fdw_$v` | `libfq` |
| **Debian** | {{< badge content="PIGSTY" link="/e/firebird_fdw" >}} | `1.4.0` | {{< bg "18" "postgresql-18-firebird-fdw" "red" >}} {{< bg "17" "postgresql-17-firebird-fdw" "green" >}} {{< bg "16" "postgresql-16-firebird-fdw" "green" >}} {{< bg "15" "postgresql-15-firebird-fdw" "green" >}} {{< bg "14" "postgresql-14-firebird-fdw" "green" >}} | `postgresql-$v-firebird-fdw` | `libfq` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "firebird_fdw_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "firebird_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "firebird_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "firebird_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "firebird_fdw_14 : AVAIL 4" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "firebird_fdw_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "firebird_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "firebird_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "firebird_fdw_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4.0" "firebird_fdw_14 : AVAIL 2" "green" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.4.1" "firebird_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "firebird_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.4.0" "firebird_fdw_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4.0" "firebird_fdw_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4.0" "firebird_fdw_14 : AVAIL 4" "green" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.4.1" "firebird_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "firebird_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.4.0" "firebird_fdw_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.0" "firebird_fdw_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.0" "firebird_fdw_14 : AVAIL 4" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-firebird-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-firebird-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-firebird-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-firebird-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-firebird-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-firebird-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `firebird_fdw_18` | 1.4.1 | `el9.x86_64` | pgdg | 52.5 KiB | [firebird_fdw_18-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/firebird_fdw_18-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `firebird_fdw_18` | 1.4.1 | `el9.aarch64` | pgdg | 51.3 KiB | [firebird_fdw_18-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/firebird_fdw_18-1.4.1-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `firebird_fdw_17` | 1.4.0 | `el8.x86_64` | pigsty | 51.4 KiB | [firebird_fdw_17-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/firebird_fdw_17-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `firebird_fdw_17` | 1.4.0 | `el8.aarch64` | pigsty | 49.5 KiB | [firebird_fdw_17-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/firebird_fdw_17-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `firebird_fdw_17` | 1.4.0 | `el9.x86_64` | pigsty | 52.4 KiB | [firebird_fdw_17-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/firebird_fdw_17-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `firebird_fdw_17` | 1.4.0 | `el9.x86_64` | pgdg | 52.7 KiB | [firebird_fdw_17-1.4.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/firebird_fdw_17-1.4.0-3PGDG.rhel9.x86_64.rpm) |
| `firebird_fdw_17` | 1.4.0 | `el9.aarch64` | pigsty | 51.3 KiB | [firebird_fdw_17-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/firebird_fdw_17-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `firebird_fdw_17` | 1.4.0 | `el9.aarch64` | pgdg | 51.6 KiB | [firebird_fdw_17-1.4.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/firebird_fdw_17-1.4.0-3PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-firebird-fdw` | 1.4.0 | `d12.x86_64` | pigsty | 162.7 KiB | [postgresql-17-firebird-fdw_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-firebird-fdw` | 1.4.0 | `d12.aarch64` | pigsty | 159.1 KiB | [postgresql-17-firebird-fdw_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-firebird-fdw` | 1.4.0 | `u22.x86_64` | pigsty | 169.5 KiB | [postgresql-17-firebird-fdw_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-firebird-fdw` | 1.4.0 | `u22.aarch64` | pigsty | 167.8 KiB | [postgresql-17-firebird-fdw_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-firebird-fdw` | 1.4.0 | `u24.x86_64` | pigsty | 146.6 KiB | [postgresql-17-firebird-fdw_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-firebird-fdw` | 1.4.0 | `u24.aarch64` | pigsty | 145.0 KiB | [postgresql-17-firebird-fdw_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `firebird_fdw_16` | 1.4.0 | `el8.x86_64` | pigsty | 51.4 KiB | [firebird_fdw_16-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/firebird_fdw_16-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `firebird_fdw_16` | 1.4.0 | `el8.aarch64` | pigsty | 49.5 KiB | [firebird_fdw_16-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/firebird_fdw_16-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `firebird_fdw_16` | 1.4.0 | `el9.x86_64` | pigsty | 52.4 KiB | [firebird_fdw_16-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/firebird_fdw_16-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `firebird_fdw_16` | 1.3.1 | `el9.x86_64` | pgdg | 51.2 KiB | [firebird_fdw_16-1.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/firebird_fdw_16-1.3.1-1PGDG.rhel9.x86_64.rpm) |
| `firebird_fdw_16` | 1.4.0 | `el9.aarch64` | pigsty | 51.3 KiB | [firebird_fdw_16-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/firebird_fdw_16-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `firebird_fdw_16` | 1.4.0 | `el9.aarch64` | pgdg | 51.4 KiB | [firebird_fdw_16-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/firebird_fdw_16-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `firebird_fdw_16` | 1.3.1 | `el9.aarch64` | pgdg | 49.9 KiB | [firebird_fdw_16-1.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/firebird_fdw_16-1.3.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-firebird-fdw` | 1.4.0 | `d12.x86_64` | pigsty | 162.6 KiB | [postgresql-16-firebird-fdw_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-firebird-fdw` | 1.4.0 | `d12.aarch64` | pigsty | 159.2 KiB | [postgresql-16-firebird-fdw_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-firebird-fdw` | 1.4.0 | `u22.x86_64` | pigsty | 169.3 KiB | [postgresql-16-firebird-fdw_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-firebird-fdw` | 1.4.0 | `u22.aarch64` | pigsty | 167.5 KiB | [postgresql-16-firebird-fdw_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-firebird-fdw` | 1.4.0 | `u24.x86_64` | pigsty | 146.7 KiB | [postgresql-16-firebird-fdw_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-firebird-fdw` | 1.4.0 | `u24.aarch64` | pigsty | 145.1 KiB | [postgresql-16-firebird-fdw_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `firebird_fdw_15` | 1.4.0 | `el8.x86_64` | pigsty | 51.9 KiB | [firebird_fdw_15-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/firebird_fdw_15-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `firebird_fdw_15` | 1.4.0 | `el8.aarch64` | pigsty | 49.9 KiB | [firebird_fdw_15-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/firebird_fdw_15-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `firebird_fdw_15` | 1.3.0 | `el8.aarch64` | pgdg | 49.0 KiB | [firebird_fdw_15-1.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/firebird_fdw_15-1.3.0-1.rhel8.aarch64.rpm) |
| `firebird_fdw_15` | 1.4.0 | `el9.x86_64` | pigsty | 53.5 KiB | [firebird_fdw_15-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/firebird_fdw_15-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `firebird_fdw_15` | 1.3.0 | `el9.x86_64` | pgdg | 52.1 KiB | [firebird_fdw_15-1.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/firebird_fdw_15-1.3.0-1.rhel9.x86_64.rpm) |
| `firebird_fdw_15` | 1.4.0 | `el9.aarch64` | pigsty | 52.5 KiB | [firebird_fdw_15-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/firebird_fdw_15-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `firebird_fdw_15` | 1.4.0 | `el9.aarch64` | pgdg | 52.6 KiB | [firebird_fdw_15-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/firebird_fdw_15-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `firebird_fdw_15` | 1.3.0 | `el9.aarch64` | pgdg | 51.0 KiB | [firebird_fdw_15-1.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/firebird_fdw_15-1.3.0-1.rhel9.aarch64.rpm) |
| `postgresql-15-firebird-fdw` | 1.4.0 | `d12.x86_64` | pigsty | 163.1 KiB | [postgresql-15-firebird-fdw_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-firebird-fdw` | 1.4.0 | `d12.aarch64` | pigsty | 159.4 KiB | [postgresql-15-firebird-fdw_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-firebird-fdw` | 1.4.0 | `u22.x86_64` | pigsty | 170.2 KiB | [postgresql-15-firebird-fdw_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-firebird-fdw` | 1.4.0 | `u22.aarch64` | pigsty | 168.6 KiB | [postgresql-15-firebird-fdw_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-firebird-fdw` | 1.4.0 | `u24.x86_64` | pigsty | 147.7 KiB | [postgresql-15-firebird-fdw_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-firebird-fdw` | 1.4.0 | `u24.aarch64` | pigsty | 146.2 KiB | [postgresql-15-firebird-fdw_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `firebird_fdw_14` | 1.4.0 | `el8.x86_64` | pigsty | 51.9 KiB | [firebird_fdw_14-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/firebird_fdw_14-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `firebird_fdw_14` | 1.2.3 | `el8.x86_64` | pgdg | 151.7 KiB | [firebird_fdw_14-1.2.3-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/firebird_fdw_14-1.2.3-2.rhel8.x86_64.rpm) |
| `firebird_fdw_14` | 1.2.3 | `el8.x86_64` | pgdg | 151.6 KiB | [firebird_fdw_14-1.2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/firebird_fdw_14-1.2.3-1.rhel8.x86_64.rpm) |
| `firebird_fdw_14` | 1.2.2 | `el8.x86_64` | pgdg | 151.4 KiB | [firebird_fdw_14-1.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/firebird_fdw_14-1.2.2-1.rhel8.x86_64.rpm) |
| `firebird_fdw_14` | 1.4.0 | `el8.aarch64` | pigsty | 49.9 KiB | [firebird_fdw_14-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/firebird_fdw_14-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `firebird_fdw_14` | 1.3.0 | `el8.aarch64` | pgdg | 49.1 KiB | [firebird_fdw_14-1.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/firebird_fdw_14-1.3.0-1.rhel8.aarch64.rpm) |
| `firebird_fdw_14` | 1.4.0 | `el9.x86_64` | pigsty | 53.6 KiB | [firebird_fdw_14-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/firebird_fdw_14-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `firebird_fdw_14` | 1.3.0 | `el9.x86_64` | pgdg | 52.1 KiB | [firebird_fdw_14-1.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/firebird_fdw_14-1.3.0-1.rhel9.x86_64.rpm) |
| `firebird_fdw_14` | 1.2.3 | `el9.x86_64` | pgdg | 153.8 KiB | [firebird_fdw_14-1.2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/firebird_fdw_14-1.2.3-1.rhel9.x86_64.rpm) |
| `firebird_fdw_14` | 1.2.3 | `el9.x86_64` | pgdg | 153.9 KiB | [firebird_fdw_14-1.2.3-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/firebird_fdw_14-1.2.3-2.rhel9.x86_64.rpm) |
| `firebird_fdw_14` | 1.4.0 | `el9.aarch64` | pigsty | 52.5 KiB | [firebird_fdw_14-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/firebird_fdw_14-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `firebird_fdw_14` | 1.4.0 | `el9.aarch64` | pgdg | 52.7 KiB | [firebird_fdw_14-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/firebird_fdw_14-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `firebird_fdw_14` | 1.3.0 | `el9.aarch64` | pgdg | 51.0 KiB | [firebird_fdw_14-1.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/firebird_fdw_14-1.3.0-1.rhel9.aarch64.rpm) |
| `firebird_fdw_14` | 1.2.3 | `el9.aarch64` | pgdg | 152.3 KiB | [firebird_fdw_14-1.2.3-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/firebird_fdw_14-1.2.3-3.rhel9.aarch64.rpm) |
| `postgresql-14-firebird-fdw` | 1.4.0 | `d12.x86_64` | pigsty | 162.3 KiB | [postgresql-14-firebird-fdw_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-firebird-fdw` | 1.4.0 | `d12.aarch64` | pigsty | 158.7 KiB | [postgresql-14-firebird-fdw_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-firebird-fdw` | 1.4.0 | `u22.x86_64` | pigsty | 169.4 KiB | [postgresql-14-firebird-fdw_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-firebird-fdw` | 1.4.0 | `u22.aarch64` | pigsty | 167.8 KiB | [postgresql-14-firebird-fdw_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-firebird-fdw` | 1.4.0 | `u24.x86_64` | pigsty | 147.6 KiB | [postgresql-14-firebird-fdw_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-firebird-fdw` | 1.4.0 | `u24.aarch64` | pigsty | 146.1 KiB | [postgresql-14-firebird-fdw_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ibarwick/firebird_fdw" title="Repository" icon="github" subtitle="github.com/ibarwick/firebird_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="firebird_fdw-1.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get firebird_fdw; # get firebird_fdw source code
pig build dep firebird_fdw; # install build dependencies
pig build pkg firebird_fdw; # build extension rpm or deb
pig build ext firebird_fdw; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install firebird_fdw; # install by extension name, for the current active PG version
pig ext install firebird_fdw; # install via package alias, for the active PG version
pig ext install firebird_fdw -v 17;   # install for PG 17
pig ext install firebird_fdw -v 16;   # install for PG 16
pig ext install firebird_fdw -v 15;   # install for PG 15
pig ext install firebird_fdw -v 14;   # install for PG 14
pig ext install firebird_fdw -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION firebird_fdw;
```

