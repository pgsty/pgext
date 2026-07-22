---
title: "jdbc_fdw"
linkTitle: "jdbc_fdw"
description: "foreign-data wrapper for remote servers available over JDBC"
weight: 8530
categories: ["FDW"]
width: full
---

[**jdbc_fdw**](https://github.com/pgspider/jdbc_fdw) : foreign-data wrapper for remote servers available over JDBC


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8530** | {{< badge content="jdbc_fdw" link="https://github.com/pgspider/jdbc_fdw" >}} | {{< ext "jdbc_fdw" >}} | `1.2` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "odbc_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "mysql_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "postgres_fdw" >}} |

> [!Note] Package/source version 0.5.0; SQL extension version 1.2. PGDG RPM is PG14-16 and missing EL aarch64; PIGSTY DEB is PG14-18 on amd64 and arm64, with a downstream PG18 compatibility patch. Live queries require a JDBC driver and remote database.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `jdbc_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.4.0` | {{< bg "18" "jdbc_fdw_18" "red" >}} {{< bg "17" "jdbc_fdw_17" "red" >}} {{< bg "16" "jdbc_fdw_16" "green" >}} {{< bg "15" "jdbc_fdw_15" "green" >}} {{< bg "14" "jdbc_fdw_14" "green" >}} | `jdbc_fdw_$v` | `java-11-openjdk-headless` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.0` | {{< bg "18" "postgresql-18-jdbc-fdw" "green" >}} {{< bg "17" "postgresql-17-jdbc-fdw" "green" >}} {{< bg "16" "postgresql-16-jdbc-fdw" "green" >}} {{< bg "15" "postgresql-15-jdbc-fdw" "green" >}} {{< bg "14" "postgresql-14-jdbc-fdw" "green" >}} | `postgresql-$v-jdbc-fdw` | `default-jre-headless`, `libpq5` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "MISS" "jdbc_fdw_18 : MISS 0" "red" >}} | {{< bg "MISS" "jdbc_fdw_17 : MISS 0" "red" >}} | {{< bg "PGDG 0.4.0" "jdbc_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.4.0" "jdbc_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.4.0" "jdbc_fdw_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "jdbc_fdw_18 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_17 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_16 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_15 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "MISS" "jdbc_fdw_18 : MISS 0" "red" >}} | {{< bg "MISS" "jdbc_fdw_17 : MISS 0" "red" >}} | {{< bg "PGDG 0.4.0" "jdbc_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.4.0" "jdbc_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.4.0" "jdbc_fdw_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "jdbc_fdw_18 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_17 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_16 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_15 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "jdbc_fdw_18 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_17 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_16 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_15 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "jdbc_fdw_18 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_17 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_16 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_15 : N/A 0" "gray" >}} | {{< bg "N/A" "jdbc_fdw_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-jdbc-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-jdbc-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-jdbc-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-jdbc-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-jdbc-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-jdbc-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-jdbc-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-jdbc-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-jdbc-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-jdbc-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-jdbc-fdw : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-jdbc-fdw` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 129.5 KiB | [postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jdbc-fdw/postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-jdbc-fdw` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 126.7 KiB | [postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jdbc-fdw/postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-jdbc-fdw` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 129.5 KiB | [postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jdbc-fdw/postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-jdbc-fdw` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 126.5 KiB | [postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jdbc-fdw/postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-jdbc-fdw` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 140.1 KiB | [postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jdbc-fdw/postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-jdbc-fdw` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 138.4 KiB | [postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jdbc-fdw/postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-jdbc-fdw` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 134.9 KiB | [postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jdbc-fdw/postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-jdbc-fdw` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 133.9 KiB | [postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jdbc-fdw/postgresql-18-jdbc-fdw_0.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-jdbc-fdw` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 133.8 KiB | [postgresql-18-jdbc-fdw_0.5.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jdbc-fdw/postgresql-18-jdbc-fdw_0.5.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-18-jdbc-fdw` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 132.1 KiB | [postgresql-18-jdbc-fdw_0.5.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jdbc-fdw/postgresql-18-jdbc-fdw_0.5.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-17-jdbc-fdw` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 129.1 KiB | [postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jdbc-fdw/postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-jdbc-fdw` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 126.3 KiB | [postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jdbc-fdw/postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-jdbc-fdw` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 129.2 KiB | [postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jdbc-fdw/postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-jdbc-fdw` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 126.4 KiB | [postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jdbc-fdw/postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-jdbc-fdw` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 161.0 KiB | [postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jdbc-fdw/postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-jdbc-fdw` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 159.5 KiB | [postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jdbc-fdw/postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-jdbc-fdw` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 134.6 KiB | [postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jdbc-fdw/postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-jdbc-fdw` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 133.4 KiB | [postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jdbc-fdw/postgresql-17-jdbc-fdw_0.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-jdbc-fdw` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 133.3 KiB | [postgresql-17-jdbc-fdw_0.5.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jdbc-fdw/postgresql-17-jdbc-fdw_0.5.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-17-jdbc-fdw` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 131.6 KiB | [postgresql-17-jdbc-fdw_0.5.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jdbc-fdw/postgresql-17-jdbc-fdw_0.5.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jdbc_fdw_16` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.8 KiB | [jdbc_fdw_16-0.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/jdbc_fdw_16-0.4.0-1PGDG.rhel8.x86_64.rpm) |
| `jdbc_fdw_16` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [jdbc_fdw_16-0.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/jdbc_fdw_16-0.4.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-jdbc-fdw` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 129.0 KiB | [postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jdbc-fdw/postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-jdbc-fdw` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 126.3 KiB | [postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jdbc-fdw/postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-jdbc-fdw` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 129.1 KiB | [postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jdbc-fdw/postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-jdbc-fdw` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 126.4 KiB | [postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jdbc-fdw/postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-jdbc-fdw` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 159.5 KiB | [postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jdbc-fdw/postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-jdbc-fdw` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 157.9 KiB | [postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jdbc-fdw/postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-jdbc-fdw` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 134.5 KiB | [postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jdbc-fdw/postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-jdbc-fdw` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 133.5 KiB | [postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jdbc-fdw/postgresql-16-jdbc-fdw_0.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-jdbc-fdw` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 133.3 KiB | [postgresql-16-jdbc-fdw_0.5.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jdbc-fdw/postgresql-16-jdbc-fdw_0.5.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-16-jdbc-fdw` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 131.6 KiB | [postgresql-16-jdbc-fdw_0.5.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jdbc-fdw/postgresql-16-jdbc-fdw_0.5.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jdbc_fdw_15` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.5 KiB | [jdbc_fdw_15-0.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/jdbc_fdw_15-0.4.0-1PGDG.rhel8.x86_64.rpm) |
| `jdbc_fdw_15` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.0 KiB | [jdbc_fdw_15-0.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/jdbc_fdw_15-0.4.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-15-jdbc-fdw` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 129.9 KiB | [postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jdbc-fdw/postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-jdbc-fdw` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 127.2 KiB | [postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jdbc-fdw/postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-jdbc-fdw` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 130.0 KiB | [postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jdbc-fdw/postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-jdbc-fdw` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 127.2 KiB | [postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jdbc-fdw/postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-jdbc-fdw` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 161.1 KiB | [postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jdbc-fdw/postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-jdbc-fdw` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 159.8 KiB | [postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jdbc-fdw/postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-jdbc-fdw` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 136.2 KiB | [postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jdbc-fdw/postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-jdbc-fdw` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 136.0 KiB | [postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jdbc-fdw/postgresql-15-jdbc-fdw_0.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-jdbc-fdw` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 135.2 KiB | [postgresql-15-jdbc-fdw_0.5.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jdbc-fdw/postgresql-15-jdbc-fdw_0.5.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-15-jdbc-fdw` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 133.7 KiB | [postgresql-15-jdbc-fdw_0.5.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jdbc-fdw/postgresql-15-jdbc-fdw_0.5.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jdbc_fdw_14` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.5 KiB | [jdbc_fdw_14-0.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/jdbc_fdw_14-0.4.0-1PGDG.rhel8.x86_64.rpm) |
| `jdbc_fdw_14` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.0 KiB | [jdbc_fdw_14-0.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/jdbc_fdw_14-0.4.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-14-jdbc-fdw` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 129.9 KiB | [postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jdbc-fdw/postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-jdbc-fdw` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 127.0 KiB | [postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jdbc-fdw/postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-jdbc-fdw` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 130.0 KiB | [postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jdbc-fdw/postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-jdbc-fdw` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 127.1 KiB | [postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jdbc-fdw/postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-jdbc-fdw` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 160.2 KiB | [postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jdbc-fdw/postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-jdbc-fdw` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 158.9 KiB | [postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jdbc-fdw/postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-jdbc-fdw` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 136.2 KiB | [postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jdbc-fdw/postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-jdbc-fdw` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 135.9 KiB | [postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jdbc-fdw/postgresql-14-jdbc-fdw_0.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-jdbc-fdw` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 135.1 KiB | [postgresql-14-jdbc-fdw_0.5.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jdbc-fdw/postgresql-14-jdbc-fdw_0.5.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-14-jdbc-fdw` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 133.6 KiB | [postgresql-14-jdbc-fdw_0.5.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jdbc-fdw/postgresql-14-jdbc-fdw_0.5.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgspider/jdbc_fdw" title="Repository" icon="github" subtitle="github.com/pgspider/jdbc_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="jdbc_fdw-0.5.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg jdbc_fdw;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install jdbc_fdw;		# install via package name, for the active PG version

pig install jdbc_fdw -v 18;   # install for PG 18
pig install jdbc_fdw -v 17;   # install for PG 17
pig install jdbc_fdw -v 16;   # install for PG 16
pig install jdbc_fdw -v 15;   # install for PG 15
pig install jdbc_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION jdbc_fdw;
```

## Usage

Sources:

- [jdbc_fdw v0.5.0 README](https://github.com/pgspider/jdbc_fdw/blob/v0.5.0/README.md)
- [Extension control file](https://github.com/pgspider/jdbc_fdw/blob/v0.5.0/jdbc_fdw.control)
- [jdbc_fdw v0.5.0 release](https://github.com/pgspider/jdbc_fdw/releases/tag/v0.5.0)

`jdbc_fdw` exposes a JDBC data source as PostgreSQL foreign tables and can execute remote SQL through a helper function. Use it when a suitable JDBC driver exists but no more specialized FDW is available; the JVM, driver JAR, credentials, and remote query behavior all run inside a PostgreSQL backend process.

### Core Workflow

```sql
CREATE EXTENSION jdbc_fdw;

CREATE SERVER reporting_jdbc
  FOREIGN DATA WRAPPER jdbc_fdw
  OPTIONS (
    drivername 'org.postgresql.Driver',
    url 'jdbc:postgresql://db.example/reporting',
    jarfile '/opt/jdbc/postgresql.jar',
    querytimeout '10',
    maxheapsize '256'
  );

CREATE USER MAPPING FOR app_user
  SERVER reporting_jdbc
  OPTIONS (username 'reader', password 'secret');

CREATE FOREIGN TABLE remote_orders (
  id bigint OPTIONS (key 'true'),
  created_at timestamptz,
  total numeric
) SERVER reporting_jdbc;

SELECT * FROM remote_orders WHERE id = 42;
```

There are no table-level options in v0.5.0. Foreign columns map by name. Mark the remote primary-key column with `OPTIONS (key 'true')` when `UPDATE` or `DELETE` needs row identity.

### Import and Direct SQL

```sql
IMPORT FOREIGN SCHEMA public
  FROM SERVER reporting_jdbc
  INTO jdbc_import
  OPTIONS (recreate 'true');

SELECT *
FROM jdbc_exec('reporting_jdbc', 'SELECT id, name FROM customer')
  AS t(id bigint, name text);
```

The upstream README says `IMPORT FOREIGN SCHEMA` currently works only with GridDB. `jdbc_exec` returns `record`, so queries returning columns require a column definition list.

### Important Options and Limits

- Server options: required `drivername` and `url`, absolute `jarfile`, plus `querytimeout` and JVM `maxheapsize`.
- User-mapping options: `username` and `password`.
- Column option: `key = true` identifies primary-key columns for writable operations.
- `jdbc_exec(connname, sql)` executes driver-specific SQL and can return a defined record set.

Version 0.5.0 supports predicate, column, and aggregate pushdown according to the upstream project, but not remote `RETURNING`, `GROUP BY`, `ORDER BY`, casts, or transaction-control statements. Arrays and foreign-table `TRUNCATE` are not implemented. Test type conversion and write semantics with the selected driver.

Protect JAR paths and server definitions from untrusted users, keep passwords in user mappings, and bound the JVM heap and remote query time. The source/package release is 0.5.0 while `jdbc_fdw.control` continues to declare SQL extension version 1.2; use `pg_extension.extversion` rather than assuming those version spaces are identical.
