---
title: "firebird_fdw"
linkTitle: "firebird_fdw"
description: "Foreign data wrapper for Firebird"
weight: 8750
categories: ["FDW"]
width: full
---

[**firebird_fdw**](https://github.com/ibarwick/firebird_fdw) : Foreign data wrapper for Firebird


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8750** | {{< badge content="firebird_fdw" link="https://github.com/ibarwick/firebird_fdw" >}} | {{< ext "firebird_fdw" >}} | `1.4.2` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "wrappers" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} {{< ext "postgres_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `firebird_fdw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.2` | {{< bg "18" "firebird_fdw_18" "green" >}} {{< bg "17" "firebird_fdw_17" "green" >}} {{< bg "16" "firebird_fdw_16" "green" >}} {{< bg "15" "firebird_fdw_15" "green" >}} {{< bg "14" "firebird_fdw_14" "green" >}} | `firebird_fdw_$v` | `libfq` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.2` | {{< bg "18" "postgresql-18-firebird-fdw" "green" >}} {{< bg "17" "postgresql-17-firebird-fdw" "green" >}} {{< bg "16" "postgresql-16-firebird-fdw" "green" >}} {{< bg "15" "postgresql-15-firebird-fdw" "green" >}} {{< bg "14" "postgresql-14-firebird-fdw" "green" >}} | `postgresql-$v-firebird-fdw` | `libfq` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_14 : AVAIL 4" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_18 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_16 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_15 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_14 : AVAIL 10" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_14 : AVAIL 10" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_14 : AVAIL 3" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.4.2" "firebird_fdw_14 : AVAIL 3" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-18-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-18-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-18-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-18-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-18-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-18-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-18-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-18-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-18-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-18-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-17-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-16-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-15-firebird-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.2" "postgresql-14-firebird-fdw : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `firebird_fdw_18` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.5 KiB | [firebird_fdw_18-1.4.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/firebird_fdw_18-1.4.2-1PIGSTY.el8.x86_64.rpm) |
| `firebird_fdw_18` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 51.8 KiB | [firebird_fdw_18-1.4.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/firebird_fdw_18-1.4.2-1PIGSTY.el8.aarch64.rpm) |
| `firebird_fdw_18` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 53.4 KiB | [firebird_fdw_18-1.4.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/firebird_fdw_18-1.4.2-1PIGSTY.el9.x86_64.rpm) |
| `firebird_fdw_18` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.1 KiB | [firebird_fdw_18-1.4.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/firebird_fdw_18-1.4.2-1PGDG.rhel9.8.x86_64.rpm) |
| `firebird_fdw_18` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.1 KiB | [firebird_fdw_18-1.4.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/firebird_fdw_18-1.4.2-1PGDG.rhel9.7.x86_64.rpm) |
| `firebird_fdw_18` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.2 KiB | [firebird_fdw_18-1.4.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/firebird_fdw_18-1.4.2-1PGDG.rhel9.6.x86_64.rpm) |
| `firebird_fdw_18` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.7 KiB | [firebird_fdw_18-1.4.1-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/firebird_fdw_18-1.4.1-3PGDG.rhel9.8.x86_64.rpm) |
| `firebird_fdw_18` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.7 KiB | [firebird_fdw_18-1.4.1-3PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/firebird_fdw_18-1.4.1-3PGDG.rhel9.7.x86_64.rpm) |
| `firebird_fdw_18` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.8 KiB | [firebird_fdw_18-1.4.1-3PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/firebird_fdw_18-1.4.1-3PGDG.rhel9.6.x86_64.rpm) |
| `firebird_fdw_18` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.5 KiB | [firebird_fdw_18-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/firebird_fdw_18-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `firebird_fdw_18` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 52.3 KiB | [firebird_fdw_18-1.4.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/firebird_fdw_18-1.4.2-1PIGSTY.el9.aarch64.rpm) |
| `firebird_fdw_18` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.1 KiB | [firebird_fdw_18-1.4.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/firebird_fdw_18-1.4.2-1PGDG.rhel9.8.aarch64.rpm) |
| `firebird_fdw_18` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.5 KiB | [firebird_fdw_18-1.4.1-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/firebird_fdw_18-1.4.1-3PGDG.rhel9.8.aarch64.rpm) |
| `firebird_fdw_18` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.2 KiB | [firebird_fdw_18-1.4.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/firebird_fdw_18-1.4.2-1PIGSTY.el10.x86_64.rpm) |
| `firebird_fdw_18` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.1 KiB | [firebird_fdw_18-1.4.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/firebird_fdw_18-1.4.2-1PGDG.rhel10.2.x86_64.rpm) |
| `firebird_fdw_18` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.7 KiB | [firebird_fdw_18-1.4.1-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/firebird_fdw_18-1.4.1-3PGDG.rhel10.2.x86_64.rpm) |
| `firebird_fdw_18` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 52.7 KiB | [firebird_fdw_18-1.4.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/firebird_fdw_18-1.4.2-1PIGSTY.el10.aarch64.rpm) |
| `firebird_fdw_18` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.5 KiB | [firebird_fdw_18-1.4.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/firebird_fdw_18-1.4.2-1PGDG.rhel10.2.aarch64.rpm) |
| `firebird_fdw_18` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.0 KiB | [firebird_fdw_18-1.4.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/firebird_fdw_18-1.4.1-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-firebird-fdw` | `1.4.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 139.1 KiB | [postgresql-18-firebird-fdw_1.4.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-18-firebird-fdw_1.4.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-firebird-fdw` | `1.4.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 136.3 KiB | [postgresql-18-firebird-fdw_1.4.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-18-firebird-fdw_1.4.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-firebird-fdw` | `1.4.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 139.6 KiB | [postgresql-18-firebird-fdw_1.4.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/firebird-fdw/postgresql-18-firebird-fdw_1.4.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-firebird-fdw` | `1.4.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 136.4 KiB | [postgresql-18-firebird-fdw_1.4.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/firebird-fdw/postgresql-18-firebird-fdw_1.4.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-firebird-fdw` | `1.4.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 148.3 KiB | [postgresql-18-firebird-fdw_1.4.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-18-firebird-fdw_1.4.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-firebird-fdw` | `1.4.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 146.8 KiB | [postgresql-18-firebird-fdw_1.4.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-18-firebird-fdw_1.4.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-firebird-fdw` | `1.4.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 142.6 KiB | [postgresql-18-firebird-fdw_1.4.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-18-firebird-fdw_1.4.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-firebird-fdw` | `1.4.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 141.5 KiB | [postgresql-18-firebird-fdw_1.4.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-18-firebird-fdw_1.4.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-firebird-fdw` | `1.4.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 141.1 KiB | [postgresql-18-firebird-fdw_1.4.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/firebird-fdw/postgresql-18-firebird-fdw_1.4.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-firebird-fdw` | `1.4.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 139.4 KiB | [postgresql-18-firebird-fdw_1.4.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/firebird-fdw/postgresql-18-firebird-fdw_1.4.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `firebird_fdw_17` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.4 KiB | [firebird_fdw_17-1.4.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/firebird_fdw_17-1.4.2-1PIGSTY.el8.x86_64.rpm) |
| `firebird_fdw_17` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 51.6 KiB | [firebird_fdw_17-1.4.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/firebird_fdw_17-1.4.2-1PIGSTY.el8.aarch64.rpm) |
| `firebird_fdw_17` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 53.5 KiB | [firebird_fdw_17-1.4.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/firebird_fdw_17-1.4.2-1PIGSTY.el9.x86_64.rpm) |
| `firebird_fdw_17` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.4 KiB | [firebird_fdw_17-1.4.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/firebird_fdw_17-1.4.2-1PGDG.rhel9.8.x86_64.rpm) |
| `firebird_fdw_17` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.0 KiB | [firebird_fdw_17-1.4.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/firebird_fdw_17-1.4.2-1PGDG.rhel9.7.x86_64.rpm) |
| `firebird_fdw_17` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.2 KiB | [firebird_fdw_17-1.4.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/firebird_fdw_17-1.4.2-1PGDG.rhel9.6.x86_64.rpm) |
| `firebird_fdw_17` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.6 KiB | [firebird_fdw_17-1.4.1-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/firebird_fdw_17-1.4.1-3PGDG.rhel9.8.x86_64.rpm) |
| `firebird_fdw_17` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.6 KiB | [firebird_fdw_17-1.4.1-3PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/firebird_fdw_17-1.4.1-3PGDG.rhel9.7.x86_64.rpm) |
| `firebird_fdw_17` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.7 KiB | [firebird_fdw_17-1.4.1-3PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/firebird_fdw_17-1.4.1-3PGDG.rhel9.6.x86_64.rpm) |
| `firebird_fdw_17` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.7 KiB | [firebird_fdw_17-1.4.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/firebird_fdw_17-1.4.0-3PGDG.rhel9.x86_64.rpm) |
| `firebird_fdw_17` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 52.4 KiB | [firebird_fdw_17-1.4.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/firebird_fdw_17-1.4.2-1PIGSTY.el9.aarch64.rpm) |
| `firebird_fdw_17` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.1 KiB | [firebird_fdw_17-1.4.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/firebird_fdw_17-1.4.2-1PGDG.rhel9.8.aarch64.rpm) |
| `firebird_fdw_17` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.6 KiB | [firebird_fdw_17-1.4.1-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/firebird_fdw_17-1.4.1-3PGDG.rhel9.8.aarch64.rpm) |
| `firebird_fdw_17` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.1 KiB | [firebird_fdw_17-1.4.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/firebird_fdw_17-1.4.2-1PIGSTY.el10.x86_64.rpm) |
| `firebird_fdw_17` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.0 KiB | [firebird_fdw_17-1.4.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/firebird_fdw_17-1.4.2-1PGDG.rhel10.2.x86_64.rpm) |
| `firebird_fdw_17` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.7 KiB | [firebird_fdw_17-1.4.1-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/firebird_fdw_17-1.4.1-3PGDG.rhel10.2.x86_64.rpm) |
| `firebird_fdw_17` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 52.5 KiB | [firebird_fdw_17-1.4.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/firebird_fdw_17-1.4.2-1PIGSTY.el10.aarch64.rpm) |
| `firebird_fdw_17` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.4 KiB | [firebird_fdw_17-1.4.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/firebird_fdw_17-1.4.2-1PGDG.rhel10.2.aarch64.rpm) |
| `firebird_fdw_17` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.8 KiB | [firebird_fdw_17-1.4.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/firebird_fdw_17-1.4.1-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-firebird-fdw` | `1.4.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 139.0 KiB | [postgresql-17-firebird-fdw_1.4.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-firebird-fdw` | `1.4.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 136.0 KiB | [postgresql-17-firebird-fdw_1.4.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-firebird-fdw` | `1.4.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 139.1 KiB | [postgresql-17-firebird-fdw_1.4.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-firebird-fdw` | `1.4.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 136.2 KiB | [postgresql-17-firebird-fdw_1.4.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-firebird-fdw` | `1.4.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 167.6 KiB | [postgresql-17-firebird-fdw_1.4.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-firebird-fdw` | `1.4.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 166.1 KiB | [postgresql-17-firebird-fdw_1.4.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-firebird-fdw` | `1.4.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 142.3 KiB | [postgresql-17-firebird-fdw_1.4.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-firebird-fdw` | `1.4.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 141.2 KiB | [postgresql-17-firebird-fdw_1.4.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-firebird-fdw` | `1.4.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 140.6 KiB | [postgresql-17-firebird-fdw_1.4.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-firebird-fdw` | `1.4.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 139.2 KiB | [postgresql-17-firebird-fdw_1.4.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/firebird-fdw/postgresql-17-firebird-fdw_1.4.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `firebird_fdw_16` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.5 KiB | [firebird_fdw_16-1.4.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/firebird_fdw_16-1.4.2-1PIGSTY.el8.x86_64.rpm) |
| `firebird_fdw_16` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 51.6 KiB | [firebird_fdw_16-1.4.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/firebird_fdw_16-1.4.2-1PIGSTY.el8.aarch64.rpm) |
| `firebird_fdw_16` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 53.5 KiB | [firebird_fdw_16-1.4.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/firebird_fdw_16-1.4.2-1PIGSTY.el9.x86_64.rpm) |
| `firebird_fdw_16` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.0 KiB | [firebird_fdw_16-1.4.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/firebird_fdw_16-1.4.2-1PGDG.rhel9.8.x86_64.rpm) |
| `firebird_fdw_16` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.0 KiB | [firebird_fdw_16-1.4.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/firebird_fdw_16-1.4.2-1PGDG.rhel9.7.x86_64.rpm) |
| `firebird_fdw_16` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.2 KiB | [firebird_fdw_16-1.4.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/firebird_fdw_16-1.4.2-1PGDG.rhel9.6.x86_64.rpm) |
| `firebird_fdw_16` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.6 KiB | [firebird_fdw_16-1.4.1-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/firebird_fdw_16-1.4.1-3PGDG.rhel9.8.x86_64.rpm) |
| `firebird_fdw_16` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.6 KiB | [firebird_fdw_16-1.4.1-3PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/firebird_fdw_16-1.4.1-3PGDG.rhel9.7.x86_64.rpm) |
| `firebird_fdw_16` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.8 KiB | [firebird_fdw_16-1.4.1-3PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/firebird_fdw_16-1.4.1-3PGDG.rhel9.6.x86_64.rpm) |
| `firebird_fdw_16` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.2 KiB | [firebird_fdw_16-1.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/firebird_fdw_16-1.3.1-1PGDG.rhel9.x86_64.rpm) |
| `firebird_fdw_16` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 52.3 KiB | [firebird_fdw_16-1.4.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/firebird_fdw_16-1.4.2-1PIGSTY.el9.aarch64.rpm) |
| `firebird_fdw_16` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.1 KiB | [firebird_fdw_16-1.4.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/firebird_fdw_16-1.4.2-1PGDG.rhel9.8.aarch64.rpm) |
| `firebird_fdw_16` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.6 KiB | [firebird_fdw_16-1.4.1-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/firebird_fdw_16-1.4.1-3PGDG.rhel9.8.aarch64.rpm) |
| `firebird_fdw_16` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.1 KiB | [firebird_fdw_16-1.4.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/firebird_fdw_16-1.4.2-1PIGSTY.el10.x86_64.rpm) |
| `firebird_fdw_16` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.0 KiB | [firebird_fdw_16-1.4.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/firebird_fdw_16-1.4.2-1PGDG.rhel10.2.x86_64.rpm) |
| `firebird_fdw_16` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.7 KiB | [firebird_fdw_16-1.4.1-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/firebird_fdw_16-1.4.1-3PGDG.rhel10.2.x86_64.rpm) |
| `firebird_fdw_16` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 52.6 KiB | [firebird_fdw_16-1.4.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/firebird_fdw_16-1.4.2-1PIGSTY.el10.aarch64.rpm) |
| `firebird_fdw_16` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.4 KiB | [firebird_fdw_16-1.4.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/firebird_fdw_16-1.4.2-1PGDG.rhel10.2.aarch64.rpm) |
| `firebird_fdw_16` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.9 KiB | [firebird_fdw_16-1.4.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/firebird_fdw_16-1.4.1-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-firebird-fdw` | `1.4.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 139.2 KiB | [postgresql-16-firebird-fdw_1.4.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-firebird-fdw` | `1.4.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 136.2 KiB | [postgresql-16-firebird-fdw_1.4.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-firebird-fdw` | `1.4.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 139.2 KiB | [postgresql-16-firebird-fdw_1.4.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-firebird-fdw` | `1.4.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 136.2 KiB | [postgresql-16-firebird-fdw_1.4.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-firebird-fdw` | `1.4.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 167.3 KiB | [postgresql-16-firebird-fdw_1.4.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-firebird-fdw` | `1.4.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 165.7 KiB | [postgresql-16-firebird-fdw_1.4.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-firebird-fdw` | `1.4.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 142.1 KiB | [postgresql-16-firebird-fdw_1.4.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-firebird-fdw` | `1.4.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 141.2 KiB | [postgresql-16-firebird-fdw_1.4.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-firebird-fdw` | `1.4.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 140.8 KiB | [postgresql-16-firebird-fdw_1.4.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-firebird-fdw` | `1.4.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 139.2 KiB | [postgresql-16-firebird-fdw_1.4.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/firebird-fdw/postgresql-16-firebird-fdw_1.4.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `firebird_fdw_15` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.9 KiB | [firebird_fdw_15-1.4.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/firebird_fdw_15-1.4.2-1PIGSTY.el8.x86_64.rpm) |
| `firebird_fdw_15` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 52.1 KiB | [firebird_fdw_15-1.4.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/firebird_fdw_15-1.4.2-1PIGSTY.el8.aarch64.rpm) |
| `firebird_fdw_15` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.0 KiB | [firebird_fdw_15-1.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/firebird_fdw_15-1.3.0-1.rhel8.aarch64.rpm) |
| `firebird_fdw_15` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.4 KiB | [firebird_fdw_15-1.4.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/firebird_fdw_15-1.4.2-1PIGSTY.el9.x86_64.rpm) |
| `firebird_fdw_15` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.1 KiB | [firebird_fdw_15-1.4.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/firebird_fdw_15-1.4.2-1PGDG.rhel9.8.x86_64.rpm) |
| `firebird_fdw_15` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.1 KiB | [firebird_fdw_15-1.4.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/firebird_fdw_15-1.4.2-1PGDG.rhel9.7.x86_64.rpm) |
| `firebird_fdw_15` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.2 KiB | [firebird_fdw_15-1.4.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/firebird_fdw_15-1.4.2-1PGDG.rhel9.6.x86_64.rpm) |
| `firebird_fdw_15` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.8 KiB | [firebird_fdw_15-1.4.1-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/firebird_fdw_15-1.4.1-3PGDG.rhel9.8.x86_64.rpm) |
| `firebird_fdw_15` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.8 KiB | [firebird_fdw_15-1.4.1-3PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/firebird_fdw_15-1.4.1-3PGDG.rhel9.7.x86_64.rpm) |
| `firebird_fdw_15` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.9 KiB | [firebird_fdw_15-1.4.1-3PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/firebird_fdw_15-1.4.1-3PGDG.rhel9.6.x86_64.rpm) |
| `firebird_fdw_15` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.1 KiB | [firebird_fdw_15-1.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/firebird_fdw_15-1.3.0-1.rhel9.x86_64.rpm) |
| `firebird_fdw_15` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.5 KiB | [firebird_fdw_15-1.4.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/firebird_fdw_15-1.4.2-1PIGSTY.el9.aarch64.rpm) |
| `firebird_fdw_15` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.2 KiB | [firebird_fdw_15-1.4.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/firebird_fdw_15-1.4.2-1PGDG.rhel9.8.aarch64.rpm) |
| `firebird_fdw_15` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.8 KiB | [firebird_fdw_15-1.4.1-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/firebird_fdw_15-1.4.1-3PGDG.rhel9.8.aarch64.rpm) |
| `firebird_fdw_15` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.2 KiB | [firebird_fdw_15-1.4.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/firebird_fdw_15-1.4.2-1PIGSTY.el10.x86_64.rpm) |
| `firebird_fdw_15` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.9 KiB | [firebird_fdw_15-1.4.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/firebird_fdw_15-1.4.2-1PGDG.rhel10.2.x86_64.rpm) |
| `firebird_fdw_15` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.7 KiB | [firebird_fdw_15-1.4.1-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/firebird_fdw_15-1.4.1-3PGDG.rhel10.2.x86_64.rpm) |
| `firebird_fdw_15` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 53.9 KiB | [firebird_fdw_15-1.4.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/firebird_fdw_15-1.4.2-1PIGSTY.el10.aarch64.rpm) |
| `firebird_fdw_15` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.6 KiB | [firebird_fdw_15-1.4.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/firebird_fdw_15-1.4.2-1PGDG.rhel10.2.aarch64.rpm) |
| `firebird_fdw_15` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.2 KiB | [firebird_fdw_15-1.4.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/firebird_fdw_15-1.4.1-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-firebird-fdw` | `1.4.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 139.7 KiB | [postgresql-15-firebird-fdw_1.4.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-firebird-fdw` | `1.4.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 136.5 KiB | [postgresql-15-firebird-fdw_1.4.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-firebird-fdw` | `1.4.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 139.7 KiB | [postgresql-15-firebird-fdw_1.4.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-firebird-fdw` | `1.4.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 136.7 KiB | [postgresql-15-firebird-fdw_1.4.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-firebird-fdw` | `1.4.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 168.4 KiB | [postgresql-15-firebird-fdw_1.4.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-firebird-fdw` | `1.4.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 166.8 KiB | [postgresql-15-firebird-fdw_1.4.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-firebird-fdw` | `1.4.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 143.1 KiB | [postgresql-15-firebird-fdw_1.4.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-firebird-fdw` | `1.4.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 142.2 KiB | [postgresql-15-firebird-fdw_1.4.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-firebird-fdw` | `1.4.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 141.6 KiB | [postgresql-15-firebird-fdw_1.4.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-firebird-fdw` | `1.4.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 140.2 KiB | [postgresql-15-firebird-fdw_1.4.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/firebird-fdw/postgresql-15-firebird-fdw_1.4.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `firebird_fdw_14` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 54.5 KiB | [firebird_fdw_14-1.4.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/firebird_fdw_14-1.4.2-1PIGSTY.el8.x86_64.rpm) |
| `firebird_fdw_14` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.7 KiB | [firebird_fdw_14-1.2.3-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/firebird_fdw_14-1.2.3-2.rhel8.x86_64.rpm) |
| `firebird_fdw_14` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.6 KiB | [firebird_fdw_14-1.2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/firebird_fdw_14-1.2.3-1.rhel8.x86_64.rpm) |
| `firebird_fdw_14` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.4 KiB | [firebird_fdw_14-1.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/firebird_fdw_14-1.2.2-1.rhel8.x86_64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 52.6 KiB | [firebird_fdw_14-1.4.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/firebird_fdw_14-1.4.2-1PIGSTY.el8.aarch64.rpm) |
| `firebird_fdw_14` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.1 KiB | [firebird_fdw_14-1.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/firebird_fdw_14-1.3.0-1.rhel8.aarch64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.0 KiB | [firebird_fdw_14-1.4.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/firebird_fdw_14-1.4.2-1PIGSTY.el9.x86_64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.6 KiB | [firebird_fdw_14-1.4.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/firebird_fdw_14-1.4.2-1PGDG.rhel9.8.x86_64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.6 KiB | [firebird_fdw_14-1.4.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/firebird_fdw_14-1.4.2-1PGDG.rhel9.7.x86_64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.7 KiB | [firebird_fdw_14-1.4.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/firebird_fdw_14-1.4.2-1PGDG.rhel9.6.x86_64.rpm) |
| `firebird_fdw_14` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.8 KiB | [firebird_fdw_14-1.4.1-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/firebird_fdw_14-1.4.1-3PGDG.rhel9.8.x86_64.rpm) |
| `firebird_fdw_14` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.8 KiB | [firebird_fdw_14-1.4.1-3PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/firebird_fdw_14-1.4.1-3PGDG.rhel9.7.x86_64.rpm) |
| `firebird_fdw_14` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.9 KiB | [firebird_fdw_14-1.4.1-3PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/firebird_fdw_14-1.4.1-3PGDG.rhel9.6.x86_64.rpm) |
| `firebird_fdw_14` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.1 KiB | [firebird_fdw_14-1.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/firebird_fdw_14-1.3.0-1.rhel9.x86_64.rpm) |
| `firebird_fdw_14` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 153.9 KiB | [firebird_fdw_14-1.2.3-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/firebird_fdw_14-1.2.3-2.rhel9.x86_64.rpm) |
| `firebird_fdw_14` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 153.8 KiB | [firebird_fdw_14-1.2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/firebird_fdw_14-1.2.3-1.rhel9.x86_64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.1 KiB | [firebird_fdw_14-1.4.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/firebird_fdw_14-1.4.2-1PIGSTY.el9.aarch64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.7 KiB | [firebird_fdw_14-1.4.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/firebird_fdw_14-1.4.2-1PGDG.rhel9.8.aarch64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.7 KiB | [firebird_fdw_14-1.4.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/firebird_fdw_14-1.4.2-1PGDG.rhel9.7.aarch64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.8 KiB | [firebird_fdw_14-1.4.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/firebird_fdw_14-1.4.2-1PGDG.rhel9.6.aarch64.rpm) |
| `firebird_fdw_14` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.8 KiB | [firebird_fdw_14-1.4.1-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/firebird_fdw_14-1.4.1-3PGDG.rhel9.8.aarch64.rpm) |
| `firebird_fdw_14` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.8 KiB | [firebird_fdw_14-1.4.1-3PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/firebird_fdw_14-1.4.1-3PGDG.rhel9.7.aarch64.rpm) |
| `firebird_fdw_14` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.9 KiB | [firebird_fdw_14-1.4.1-3PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/firebird_fdw_14-1.4.1-3PGDG.rhel9.6.aarch64.rpm) |
| `firebird_fdw_14` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.7 KiB | [firebird_fdw_14-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/firebird_fdw_14-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `firebird_fdw_14` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.0 KiB | [firebird_fdw_14-1.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/firebird_fdw_14-1.3.0-1.rhel9.aarch64.rpm) |
| `firebird_fdw_14` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 152.3 KiB | [firebird_fdw_14-1.2.3-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/firebird_fdw_14-1.2.3-3.rhel9.aarch64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.8 KiB | [firebird_fdw_14-1.4.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/firebird_fdw_14-1.4.2-1PIGSTY.el10.x86_64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.5 KiB | [firebird_fdw_14-1.4.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/firebird_fdw_14-1.4.2-1PGDG.rhel10.2.x86_64.rpm) |
| `firebird_fdw_14` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.7 KiB | [firebird_fdw_14-1.4.1-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/firebird_fdw_14-1.4.1-3PGDG.rhel10.2.x86_64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.4 KiB | [firebird_fdw_14-1.4.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/firebird_fdw_14-1.4.2-1PIGSTY.el10.aarch64.rpm) |
| `firebird_fdw_14` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.1 KiB | [firebird_fdw_14-1.4.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/firebird_fdw_14-1.4.2-1PGDG.rhel10.2.aarch64.rpm) |
| `firebird_fdw_14` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.3 KiB | [firebird_fdw_14-1.4.1-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/firebird_fdw_14-1.4.1-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-firebird-fdw` | `1.4.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 141.2 KiB | [postgresql-14-firebird-fdw_1.4.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-firebird-fdw` | `1.4.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 138.1 KiB | [postgresql-14-firebird-fdw_1.4.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-firebird-fdw` | `1.4.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 141.3 KiB | [postgresql-14-firebird-fdw_1.4.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-firebird-fdw` | `1.4.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 138.1 KiB | [postgresql-14-firebird-fdw_1.4.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-firebird-fdw` | `1.4.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 169.4 KiB | [postgresql-14-firebird-fdw_1.4.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-firebird-fdw` | `1.4.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 167.8 KiB | [postgresql-14-firebird-fdw_1.4.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-firebird-fdw` | `1.4.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 144.8 KiB | [postgresql-14-firebird-fdw_1.4.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-firebird-fdw` | `1.4.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 143.9 KiB | [postgresql-14-firebird-fdw_1.4.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-firebird-fdw` | `1.4.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 143.1 KiB | [postgresql-14-firebird-fdw_1.4.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-firebird-fdw` | `1.4.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 141.6 KiB | [postgresql-14-firebird-fdw_1.4.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/f/firebird-fdw/postgresql-14-firebird-fdw_1.4.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ibarwick/firebird_fdw" title="Repository" icon="github" subtitle="github.com/ibarwick/firebird_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="firebird_fdw-1.4.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg firebird_fdw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install firebird_fdw;		# install via package name, for the active PG version

pig install firebird_fdw -v 18;   # install for PG 18
pig install firebird_fdw -v 17;   # install for PG 17
pig install firebird_fdw -v 16;   # install for PG 16
pig install firebird_fdw -v 15;   # install for PG 15
pig install firebird_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION firebird_fdw;
```

## Usage

Sources: [upstream README](https://github.com/ibarwick/firebird_fdw), [1.4.2 README](https://github.com/ibarwick/firebird_fdw/blob/REL_1_4_STABLE/README.md), [1.4.2 tag](https://github.com/ibarwick/firebird_fdw/tree/1.4.2).

`firebird_fdw` connects PostgreSQL to Firebird databases through the foreign data wrapper API. It supports reads, writes, `IMPORT FOREIGN SCHEMA`, predicate pushdown for supported expressions, connection caching, and foreign-table `TRUNCATE` on PostgreSQL 14+.

### Create Server

```sql
CREATE EXTENSION firebird_fdw;

CREATE SERVER firebird_server FOREIGN DATA WRAPPER firebird_fdw
  OPTIONS (address 'localhost', database '/path/to/database.fdb');
```

Server options include:

- `address`, default `localhost`.
- `port`, default `3050`.
- `database`, the Firebird database name or path.
- `updatable`, default `true`; table-level settings can override it.
- `disable_pushdowns`, useful for debugging or benchmarking.
- `quote_identifiers`, to quote table and column identifiers by default.
- `implicit_bool_type`, for integer-backed Firebird boolean columns.
- `batch_size`, for multi-row inserts on PostgreSQL 14+.

### Create User Mapping

```sql
CREATE USER MAPPING FOR CURRENT_USER SERVER firebird_server
  OPTIONS (username 'sysdba', password 'masterke');
```

### Create Foreign Table

```sql
CREATE FOREIGN TABLE fb_test (
  id smallint,
  val varchar(2048)
)
SERVER firebird_server
OPTIONS (table_name 'fdw_test');
```

With column name mapping:

```sql
CREATE FOREIGN TABLE fb_mapped (
  id smallint OPTIONS (column_name 'test_id'),
  val varchar(2048) OPTIONS (column_name 'test_val')
)
SERVER firebird_server
OPTIONS (table_name 'fdw_test');
```

With a custom query, the foreign table is read-only:

```sql
CREATE FOREIGN TABLE fb_query (
  id smallint,
  val varchar(2048)
)
SERVER firebird_server
OPTIONS (query $$ SELECT id, val FROM fdw_test WHERE id > 10 $$);
```

Table options include `table_name`, `query`, `updatable`, `estimated_row_count`, `quote_identifier`, and `batch_size`. Column options include `column_name`, `quote_identifier`, and `implicit_bool_type`.

### Import Foreign Schema

```sql
IMPORT FOREIGN SCHEMA someschema
  FROM SERVER firebird_server
  INTO public
  OPTIONS (import_views 'true', verbose 'true');
```

Import options include `import_not_null`, `import_views`, `updatable`, and `verbose`. The schema name has no special Firebird meaning and can be any value accepted by PostgreSQL's `IMPORT FOREIGN SCHEMA` syntax.

### CRUD Operations

```sql
SELECT * FROM fb_test WHERE id > 5;
INSERT INTO fb_test VALUES (10, 'new record');
UPDATE fb_test SET val = 'updated' WHERE id = 10;
DELETE FROM fb_test WHERE id = 10;
TRUNCATE fb_test;
```

`UPDATE` and `DELETE` use Firebird's `RDB$DB_KEY`. `TRUNCATE` is implemented as an unqualified Firebird `DELETE` because Firebird has no native `TRUNCATE` statement.

### Utility Functions

- `firebird_fdw_version()` returns the FDW version as an integer.
- `firebird_fdw_close_connections()` closes cached Firebird connections for the current PostgreSQL session.
- `firebird_fdw_server_options(servername text)` shows effective server option values and whether each was explicitly provided.
- `firebird_fdw_diag()` returns diagnostic key/value data, including FDW and `libfq` versions.
- `firebird_version()` reports Firebird server versions for configured foreign servers and may open connections to do so.

### Caveats

- Pigsty packages `firebird_fdw` 1.4.2 for PostgreSQL 14-18. Upstream documents compatibility with PostgreSQL 10-19, with newer FDW API features available only on newer PostgreSQL releases.
- Upstream supports Firebird 2.5 and later; older Firebird versions are not a tested target.
- `firebird_fdw` and `libfq` are developed together, so package compatibility depends on matching those libraries.
- Custom-query foreign tables cannot be updated.
- The `implicit_bool_type` feature is experimental and is designed around integer columns representing boolean values.
