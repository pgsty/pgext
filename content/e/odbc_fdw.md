---
title: "odbc_fdw"
linkTitle: "odbc_fdw"
description: "Foreign data wrapper for accessing remote databases using ODBC"
weight: 8520
categories: ["FDW"]
width: full
---

[**odbc_fdw**](https://github.com/devrimgunduz/odbc_fdw) : Foreign data wrapper for accessing remote databases using ODBC


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8520** | {{< badge content="odbc_fdw" link="https://github.com/devrimgunduz/odbc_fdw" >}} | {{< ext "odbc_fdw" >}} | `0.5.2` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "jdbc_fdw" >}} {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "postgres_fdw" >}} |

> [!Note] Package/source version 0.6.1; SQL extension version 0.5.2. Live queries require an installed ODBC driver and configured DSN.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.5.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `odbc_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.6.1` | {{< bg "18" "odbc_fdw_18" "green" >}} {{< bg "17" "odbc_fdw_17" "green" >}} {{< bg "16" "odbc_fdw_16" "green" >}} {{< bg "15" "odbc_fdw_15" "green" >}} {{< bg "14" "odbc_fdw_14" "green" >}} | `odbc_fdw_$v` | `unixODBC` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.1` | {{< bg "18" "postgresql-18-odbc-fdw" "green" >}} {{< bg "17" "postgresql-17-odbc-fdw" "green" >}} {{< bg "16" "postgresql-16-odbc-fdw" "green" >}} {{< bg "15" "postgresql-15-odbc-fdw" "green" >}} {{< bg "14" "postgresql-14-odbc-fdw" "green" >}} | `postgresql-$v-odbc-fdw` | `libodbc2` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_14 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.6.1" "odbc_fdw_14 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-odbc-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-odbc-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-odbc-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-odbc-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-odbc-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-odbc-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-odbc-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-odbc-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-odbc-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-odbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-odbc-fdw : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `odbc_fdw_18` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.9 KiB | [odbc_fdw_18-0.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/odbc_fdw_18-0.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `odbc_fdw_18` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.1 KiB | [odbc_fdw_18-0.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/odbc_fdw_18-0.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `odbc_fdw_18` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.0 KiB | [odbc_fdw_18-0.6.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/odbc_fdw_18-0.6.1-1PGDG.rhel9.8.x86_64.rpm) |
| `odbc_fdw_18` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.4 KiB | [odbc_fdw_18-0.6.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/odbc_fdw_18-0.6.1-1PGDG.rhel9.8.aarch64.rpm) |
| `odbc_fdw_18` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.0 KiB | [odbc_fdw_18-0.6.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/odbc_fdw_18-0.6.1-1PGDG.rhel10.2.x86_64.rpm) |
| `odbc_fdw_18` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.8 KiB | [odbc_fdw_18-0.6.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/odbc_fdw_18-0.6.1-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-odbc-fdw` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 48.7 KiB | [postgresql-18-odbc-fdw_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/odbc-fdw/postgresql-18-odbc-fdw_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-odbc-fdw` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 47.9 KiB | [postgresql-18-odbc-fdw_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/odbc-fdw/postgresql-18-odbc-fdw_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-odbc-fdw` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 49.0 KiB | [postgresql-18-odbc-fdw_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/odbc-fdw/postgresql-18-odbc-fdw_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-odbc-fdw` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 47.8 KiB | [postgresql-18-odbc-fdw_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/odbc-fdw/postgresql-18-odbc-fdw_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-odbc-fdw` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 53.1 KiB | [postgresql-18-odbc-fdw_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/odbc-fdw/postgresql-18-odbc-fdw_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-odbc-fdw` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 52.5 KiB | [postgresql-18-odbc-fdw_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/odbc-fdw/postgresql-18-odbc-fdw_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-odbc-fdw` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 51.1 KiB | [postgresql-18-odbc-fdw_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/odbc-fdw/postgresql-18-odbc-fdw_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-odbc-fdw` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 50.7 KiB | [postgresql-18-odbc-fdw_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/odbc-fdw/postgresql-18-odbc-fdw_0.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-odbc-fdw` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 50.8 KiB | [postgresql-18-odbc-fdw_0.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/odbc-fdw/postgresql-18-odbc-fdw_0.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-odbc-fdw` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 50.3 KiB | [postgresql-18-odbc-fdw_0.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/odbc-fdw/postgresql-18-odbc-fdw_0.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `odbc_fdw_17` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.9 KiB | [odbc_fdw_17-0.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/odbc_fdw_17-0.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.4 KiB | [odbc_fdw_17-0.5.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/odbc_fdw_17-0.5.1-2PGDG.rhel8.x86_64.rpm) |
| `odbc_fdw_17` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.1 KiB | [odbc_fdw_17-0.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/odbc_fdw_17-0.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.6 KiB | [odbc_fdw_17-0.5.1-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/odbc_fdw_17-0.5.1-2PGDG.rhel8.aarch64.rpm) |
| `odbc_fdw_17` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.0 KiB | [odbc_fdw_17-0.6.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/odbc_fdw_17-0.6.1-1PGDG.rhel9.8.x86_64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.7 KiB | [odbc_fdw_17-0.5.1-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/odbc_fdw_17-0.5.1-5PGDG.rhel9.8.x86_64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.4 KiB | [odbc_fdw_17-0.5.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/odbc_fdw_17-0.5.1-2PGDG.rhel9.x86_64.rpm) |
| `odbc_fdw_17` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.4 KiB | [odbc_fdw_17-0.6.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/odbc_fdw_17-0.6.1-1PGDG.rhel9.8.aarch64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.1 KiB | [odbc_fdw_17-0.5.1-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/odbc_fdw_17-0.5.1-5PGDG.rhel9.8.aarch64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.9 KiB | [odbc_fdw_17-0.5.1-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/odbc_fdw_17-0.5.1-2PGDG.rhel9.aarch64.rpm) |
| `odbc_fdw_17` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.0 KiB | [odbc_fdw_17-0.6.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/odbc_fdw_17-0.6.1-1PGDG.rhel10.2.x86_64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.3 KiB | [odbc_fdw_17-0.5.1-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/odbc_fdw_17-0.5.1-5PGDG.rhel10.2.x86_64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [odbc_fdw_17-0.5.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/odbc_fdw_17-0.5.1-3PGDG.rhel10.x86_64.rpm) |
| `odbc_fdw_17` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.7 KiB | [odbc_fdw_17-0.6.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/odbc_fdw_17-0.6.1-1PGDG.rhel10.2.aarch64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.4 KiB | [odbc_fdw_17-0.5.1-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/odbc_fdw_17-0.5.1-5PGDG.rhel10.2.aarch64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [odbc_fdw_17-0.5.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/odbc_fdw_17-0.5.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-odbc-fdw` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 48.9 KiB | [postgresql-17-odbc-fdw_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/odbc-fdw/postgresql-17-odbc-fdw_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-odbc-fdw` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 47.9 KiB | [postgresql-17-odbc-fdw_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/odbc-fdw/postgresql-17-odbc-fdw_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-odbc-fdw` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 49.0 KiB | [postgresql-17-odbc-fdw_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/odbc-fdw/postgresql-17-odbc-fdw_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-odbc-fdw` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 47.7 KiB | [postgresql-17-odbc-fdw_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/odbc-fdw/postgresql-17-odbc-fdw_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-odbc-fdw` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 60.4 KiB | [postgresql-17-odbc-fdw_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/odbc-fdw/postgresql-17-odbc-fdw_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-odbc-fdw` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 59.6 KiB | [postgresql-17-odbc-fdw_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/odbc-fdw/postgresql-17-odbc-fdw_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-odbc-fdw` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 51.1 KiB | [postgresql-17-odbc-fdw_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/odbc-fdw/postgresql-17-odbc-fdw_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-odbc-fdw` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 50.4 KiB | [postgresql-17-odbc-fdw_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/odbc-fdw/postgresql-17-odbc-fdw_0.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-odbc-fdw` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 50.8 KiB | [postgresql-17-odbc-fdw_0.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/odbc-fdw/postgresql-17-odbc-fdw_0.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-odbc-fdw` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 50.2 KiB | [postgresql-17-odbc-fdw_0.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/odbc-fdw/postgresql-17-odbc-fdw_0.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `odbc_fdw_16` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.9 KiB | [odbc_fdw_16-0.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/odbc_fdw_16-0.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [odbc_fdw_16-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/odbc_fdw_16-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `odbc_fdw_16` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.1 KiB | [odbc_fdw_16-0.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/odbc_fdw_16-0.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.5 KiB | [odbc_fdw_16-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/odbc_fdw_16-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `odbc_fdw_16` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.0 KiB | [odbc_fdw_16-0.6.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/odbc_fdw_16-0.6.1-1PGDG.rhel9.8.x86_64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.7 KiB | [odbc_fdw_16-0.5.1-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/odbc_fdw_16-0.5.1-5PGDG.rhel9.8.x86_64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [odbc_fdw_16-0.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/odbc_fdw_16-0.5.1-1PGDG.rhel9.x86_64.rpm) |
| `odbc_fdw_16` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.4 KiB | [odbc_fdw_16-0.6.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/odbc_fdw_16-0.6.1-1PGDG.rhel9.8.aarch64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.1 KiB | [odbc_fdw_16-0.5.1-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/odbc_fdw_16-0.5.1-5PGDG.rhel9.8.aarch64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.8 KiB | [odbc_fdw_16-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/odbc_fdw_16-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `odbc_fdw_16` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.0 KiB | [odbc_fdw_16-0.6.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/odbc_fdw_16-0.6.1-1PGDG.rhel10.2.x86_64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.3 KiB | [odbc_fdw_16-0.5.1-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/odbc_fdw_16-0.5.1-5PGDG.rhel10.2.x86_64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [odbc_fdw_16-0.5.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/odbc_fdw_16-0.5.1-3PGDG.rhel10.x86_64.rpm) |
| `odbc_fdw_16` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.7 KiB | [odbc_fdw_16-0.6.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/odbc_fdw_16-0.6.1-1PGDG.rhel10.2.aarch64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.4 KiB | [odbc_fdw_16-0.5.1-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/odbc_fdw_16-0.5.1-5PGDG.rhel10.2.aarch64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [odbc_fdw_16-0.5.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/odbc_fdw_16-0.5.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-odbc-fdw` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 48.6 KiB | [postgresql-16-odbc-fdw_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/odbc-fdw/postgresql-16-odbc-fdw_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-odbc-fdw` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 47.7 KiB | [postgresql-16-odbc-fdw_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/odbc-fdw/postgresql-16-odbc-fdw_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-odbc-fdw` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 49.0 KiB | [postgresql-16-odbc-fdw_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/odbc-fdw/postgresql-16-odbc-fdw_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-odbc-fdw` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 47.7 KiB | [postgresql-16-odbc-fdw_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/odbc-fdw/postgresql-16-odbc-fdw_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-odbc-fdw` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 60.0 KiB | [postgresql-16-odbc-fdw_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/odbc-fdw/postgresql-16-odbc-fdw_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-odbc-fdw` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 59.6 KiB | [postgresql-16-odbc-fdw_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/odbc-fdw/postgresql-16-odbc-fdw_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-odbc-fdw` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 51.0 KiB | [postgresql-16-odbc-fdw_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/odbc-fdw/postgresql-16-odbc-fdw_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-odbc-fdw` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 50.5 KiB | [postgresql-16-odbc-fdw_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/odbc-fdw/postgresql-16-odbc-fdw_0.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-odbc-fdw` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 50.7 KiB | [postgresql-16-odbc-fdw_0.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/odbc-fdw/postgresql-16-odbc-fdw_0.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-odbc-fdw` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 50.3 KiB | [postgresql-16-odbc-fdw_0.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/odbc-fdw/postgresql-16-odbc-fdw_0.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `odbc_fdw_15` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.9 KiB | [odbc_fdw_15-0.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/odbc_fdw_15-0.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [odbc_fdw_15-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/odbc_fdw_15-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `odbc_fdw_15` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.1 KiB | [odbc_fdw_15-0.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/odbc_fdw_15-0.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.5 KiB | [odbc_fdw_15-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/odbc_fdw_15-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `odbc_fdw_15` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.0 KiB | [odbc_fdw_15-0.6.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/odbc_fdw_15-0.6.1-1PGDG.rhel9.8.x86_64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.7 KiB | [odbc_fdw_15-0.5.1-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/odbc_fdw_15-0.5.1-5PGDG.rhel9.8.x86_64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [odbc_fdw_15-0.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/odbc_fdw_15-0.5.1-1PGDG.rhel9.x86_64.rpm) |
| `odbc_fdw_15` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.4 KiB | [odbc_fdw_15-0.6.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/odbc_fdw_15-0.6.1-1PGDG.rhel9.8.aarch64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.1 KiB | [odbc_fdw_15-0.5.1-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/odbc_fdw_15-0.5.1-5PGDG.rhel9.8.aarch64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.7 KiB | [odbc_fdw_15-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/odbc_fdw_15-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `odbc_fdw_15` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.0 KiB | [odbc_fdw_15-0.6.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/odbc_fdw_15-0.6.1-1PGDG.rhel10.2.x86_64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.3 KiB | [odbc_fdw_15-0.5.1-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/odbc_fdw_15-0.5.1-5PGDG.rhel10.2.x86_64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [odbc_fdw_15-0.5.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/odbc_fdw_15-0.5.1-3PGDG.rhel10.x86_64.rpm) |
| `odbc_fdw_15` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.8 KiB | [odbc_fdw_15-0.6.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/odbc_fdw_15-0.6.1-1PGDG.rhel10.2.aarch64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.4 KiB | [odbc_fdw_15-0.5.1-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/odbc_fdw_15-0.5.1-5PGDG.rhel10.2.aarch64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [odbc_fdw_15-0.5.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/odbc_fdw_15-0.5.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-odbc-fdw` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 48.6 KiB | [postgresql-15-odbc-fdw_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/odbc-fdw/postgresql-15-odbc-fdw_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-odbc-fdw` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 47.9 KiB | [postgresql-15-odbc-fdw_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/odbc-fdw/postgresql-15-odbc-fdw_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-odbc-fdw` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 49.1 KiB | [postgresql-15-odbc-fdw_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/odbc-fdw/postgresql-15-odbc-fdw_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-odbc-fdw` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 47.5 KiB | [postgresql-15-odbc-fdw_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/odbc-fdw/postgresql-15-odbc-fdw_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-odbc-fdw` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 59.9 KiB | [postgresql-15-odbc-fdw_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/odbc-fdw/postgresql-15-odbc-fdw_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-odbc-fdw` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 59.6 KiB | [postgresql-15-odbc-fdw_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/odbc-fdw/postgresql-15-odbc-fdw_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-odbc-fdw` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 51.1 KiB | [postgresql-15-odbc-fdw_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/odbc-fdw/postgresql-15-odbc-fdw_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-odbc-fdw` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 50.1 KiB | [postgresql-15-odbc-fdw_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/odbc-fdw/postgresql-15-odbc-fdw_0.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-odbc-fdw` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 50.8 KiB | [postgresql-15-odbc-fdw_0.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/odbc-fdw/postgresql-15-odbc-fdw_0.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-odbc-fdw` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 50.3 KiB | [postgresql-15-odbc-fdw_0.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/odbc-fdw/postgresql-15-odbc-fdw_0.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `odbc_fdw_14` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.9 KiB | [odbc_fdw_14-0.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/odbc_fdw_14-0.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [odbc_fdw_14-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/odbc_fdw_14-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `odbc_fdw_14` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.1 KiB | [odbc_fdw_14-0.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/odbc_fdw_14-0.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.5 KiB | [odbc_fdw_14-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/odbc_fdw_14-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `odbc_fdw_14` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.0 KiB | [odbc_fdw_14-0.6.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/odbc_fdw_14-0.6.1-1PGDG.rhel9.8.x86_64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.7 KiB | [odbc_fdw_14-0.5.1-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/odbc_fdw_14-0.5.1-5PGDG.rhel9.8.x86_64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [odbc_fdw_14-0.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/odbc_fdw_14-0.5.1-1PGDG.rhel9.x86_64.rpm) |
| `odbc_fdw_14` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.4 KiB | [odbc_fdw_14-0.6.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/odbc_fdw_14-0.6.1-1PGDG.rhel9.8.aarch64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.1 KiB | [odbc_fdw_14-0.5.1-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/odbc_fdw_14-0.5.1-5PGDG.rhel9.8.aarch64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.8 KiB | [odbc_fdw_14-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/odbc_fdw_14-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `odbc_fdw_14` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.0 KiB | [odbc_fdw_14-0.6.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/odbc_fdw_14-0.6.1-1PGDG.rhel10.2.x86_64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.3 KiB | [odbc_fdw_14-0.5.1-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/odbc_fdw_14-0.5.1-5PGDG.rhel10.2.x86_64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [odbc_fdw_14-0.5.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/odbc_fdw_14-0.5.1-3PGDG.rhel10.x86_64.rpm) |
| `odbc_fdw_14` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 28.8 KiB | [odbc_fdw_14-0.6.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/odbc_fdw_14-0.6.1-1PGDG.rhel10.2.aarch64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [odbc_fdw_14-0.5.1-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/odbc_fdw_14-0.5.1-5PGDG.rhel10.2.aarch64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [odbc_fdw_14-0.5.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/odbc_fdw_14-0.5.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-odbc-fdw` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 48.5 KiB | [postgresql-14-odbc-fdw_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/odbc-fdw/postgresql-14-odbc-fdw_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-odbc-fdw` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 47.8 KiB | [postgresql-14-odbc-fdw_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/odbc-fdw/postgresql-14-odbc-fdw_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-odbc-fdw` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 48.9 KiB | [postgresql-14-odbc-fdw_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/odbc-fdw/postgresql-14-odbc-fdw_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-odbc-fdw` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 47.4 KiB | [postgresql-14-odbc-fdw_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/odbc-fdw/postgresql-14-odbc-fdw_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-odbc-fdw` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 59.1 KiB | [postgresql-14-odbc-fdw_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/odbc-fdw/postgresql-14-odbc-fdw_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-odbc-fdw` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 58.8 KiB | [postgresql-14-odbc-fdw_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/odbc-fdw/postgresql-14-odbc-fdw_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-odbc-fdw` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 51.0 KiB | [postgresql-14-odbc-fdw_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/odbc-fdw/postgresql-14-odbc-fdw_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-odbc-fdw` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 50.0 KiB | [postgresql-14-odbc-fdw_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/odbc-fdw/postgresql-14-odbc-fdw_0.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-odbc-fdw` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 50.7 KiB | [postgresql-14-odbc-fdw_0.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/odbc-fdw/postgresql-14-odbc-fdw_0.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-odbc-fdw` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 50.2 KiB | [postgresql-14-odbc-fdw_0.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/odbc-fdw/postgresql-14-odbc-fdw_0.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/devrimgunduz/odbc_fdw" title="Repository" icon="github" subtitle="github.com/devrimgunduz/odbc_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="odbc_fdw-0.6.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg odbc_fdw;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install odbc_fdw;		# install via package name, for the active PG version

pig install odbc_fdw -v 18;   # install for PG 18
pig install odbc_fdw -v 17;   # install for PG 17
pig install odbc_fdw -v 16;   # install for PG 16
pig install odbc_fdw -v 15;   # install for PG 15
pig install odbc_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION odbc_fdw;
```

## Usage

Sources:

- [odbc_fdw 0.6.1 README](https://github.com/devrimgunduz/odbc_fdw/blob/0.6.1/README.md)
- [odbc_fdw changelog](https://github.com/devrimgunduz/odbc_fdw/blob/0.6.1/NEWS.md)
- [Extension control file](https://github.com/devrimgunduz/odbc_fdw/blob/0.6.1/odbc_fdw.control)
- [0.6.0 to 0.6.1 comparison](https://github.com/devrimgunduz/odbc_fdw/compare/0.6.0...0.6.1)

`odbc_fdw` exposes tables or driver-specific queries from an ODBC data source as PostgreSQL foreign tables. It is primarily a read/query bridge across heterogeneous systems; validate data-type conversions and remote-driver behavior before relying on it for production queries.

### Core Workflow

```sql
CREATE EXTENSION odbc_fdw;

-- In 0.6.1 a superuser must set the server-level dsn or driver option.
CREATE SERVER warehouse_odbc
  FOREIGN DATA WRAPPER odbc_fdw
  OPTIONS (dsn 'warehouse');

CREATE USER MAPPING FOR analyst
  SERVER warehouse_odbc
  OPTIONS (odbc_UID 'reporter', odbc_PWD 'secret');

CREATE FOREIGN TABLE remote_customer (
  id bigint,
  name text,
  created_at timestamp
) SERVER warehouse_odbc
  OPTIONS (schema 'sales', table 'customer');

SELECT * FROM remote_customer WHERE id = 42;
```

Use `driver` instead of `dsn` for a DSN-less connection. Other driver attributes use the `odbc_` prefix and may be placed on the server, user mapping, or foreign table. Put credentials in a user mapping. Quote case-sensitive attribute names, and wrap values containing `=` or `;` in braces as required by the driver.

### Queries and Import

`sql_query` overrides `table`; pair it with `sql_count` when the FDW needs an explicit row-count query:

```sql
CREATE FOREIGN TABLE active_customer (
  id bigint,
  name text
) SERVER warehouse_odbc
  OPTIONS (
    sql_query 'SELECT id, name FROM sales.customer WHERE active = 1',
    sql_count 'SELECT count(*) FROM sales.customer WHERE active = 1'
  );

IMPORT FOREIGN SCHEMA sales
  FROM SERVER warehouse_odbc
  INTO imported
  OPTIONS (prefix 'odbc_');
```

### Important Objects and Options

- `dsn` or `driver` selects the ODBC data source; 0.6.1 restricts these server options to superusers because the driver manager loads shared libraries.
- `schema`, `table`, `sql_query`, and `sql_count` select the remote relation or query.
- `prefix` changes local names created by `IMPORT FOREIGN SCHEMA`.
- `ODBCTablesList(server_name, ...)` lists visible remote tables.
- `ODBCTableSize(server_name, table_name)` and `ODBCQuerySize(server_name, query)` return remote row counts.

Version 0.6.0 restores compatibility and fixes crashes on recent PostgreSQL releases. Version 0.6.1 escapes remote literals and identifiers to prevent SQL injection, restricts driver selection, and redacts common credential attributes in debug connection strings. Upgrade before allowing delegated FDW use, while retaining normal server ownership and user-mapping controls.

Only the ODBC types listed by the upstream README are fully supported. Identifier length, driver SQL dialect, encodings, null handling, and binary values can vary. The source/package release is 0.6.1, while the control file and install SQL continue to declare extension version 0.5.2.
