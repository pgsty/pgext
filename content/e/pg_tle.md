---
title: "pg_tle"
linkTitle: "pg_tle"
description: "Trusted Language Extensions for PostgreSQL"
weight: 3000
categories: ["LANG"]
width: full
---

[**pg_tle**](https://github.com/aws/pg_tle) : Trusted Language Extensions for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3000** | {{< badge content="pg_tle" link="https://github.com/aws/pg_tle" >}} | {{< ext "pg_tle" >}} | `1.5.2` | {{< category "LANG" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "plv8" >}} {{< ext "pllua" >}} {{< ext "pljava" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "plprql" >}} {{< ext "plsh" >}} |

> [!Note] require bison flex to build


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_tle` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5.2` | {{< bg "18" "pg_tle_18*" "green" >}} {{< bg "17" "pg_tle_17*" "green" >}} {{< bg "16" "pg_tle_16*" "green" >}} {{< bg "15" "pg_tle_15*" "green" >}} {{< bg "14" "pg_tle_14*" "green" >}} {{< bg "13" "pg_tle_13*" "green" >}} | `pg_tle_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.2` | {{< bg "18" "postgresql-18-pg-tle" "green" >}} {{< bg "17" "postgresql-17-pg-tle" "green" >}} {{< bg "16" "postgresql-16-pg-tle" "green" >}} {{< bg "15" "postgresql-15-pg-tle" "green" >}} {{< bg "14" "postgresql-14-pg-tle" "green" >}} {{< bg "13" "postgresql-13-pg-tle" "green" >}} | `postgresql-$v-pg-tle` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_13 : AVAIL 4" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_13 : AVAIL 4" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_13 : AVAIL 4" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_13 : AVAIL 4" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_13 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.2" "pg_tle_13 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-18-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-17-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-16-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-15-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-14-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-13-pg-tle : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-18-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-17-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-16-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-15-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-14-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-13-pg-tle : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-18-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-17-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-16-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-15-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-14-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-13-pg-tle : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-18-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-17-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-16-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-15-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-14-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-13-pg-tle : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-18-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-17-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-16-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-15-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-14-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-13-pg-tle : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-18-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-17-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-16-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-15-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-14-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-13-pg-tle : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-18-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-17-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-16-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-15-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-14-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-13-pg-tle : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-18-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-17-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-16-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-15-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-14-pg-tle : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.2" "postgresql-13-pg-tle : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tle_18` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 68.8 KiB | [pg_tle_18-1.5.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_18-1.5.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_tle_18` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.7 KiB | [pg_tle_18-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_tle_18-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_18` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 65.4 KiB | [pg_tle_18-1.5.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_18-1.5.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_tle_18` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.3 KiB | [pg_tle_18-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_tle_18-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_18` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 63.0 KiB | [pg_tle_18-1.5.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_18-1.5.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_tle_18` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.1 KiB | [pg_tle_18-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_tle_18-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_18` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 59.9 KiB | [pg_tle_18-1.5.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_18-1.5.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_tle_18` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.1 KiB | [pg_tle_18-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_tle_18-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_18` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 63.3 KiB | [pg_tle_18-1.5.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tle_18-1.5.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_tle_18` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.8 KiB | [pg_tle_18-1.5.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_tle_18-1.5.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_tle_18` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.6 KiB | [pg_tle_18-1.5.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tle_18-1.5.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_tle_18` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.3 KiB | [pg_tle_18-1.5.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_tle_18-1.5.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-tle` | `1.5.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 159.9 KiB | [postgresql-18-pg-tle_1.5.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-18-pg-tle_1.5.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-tle` | `1.5.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 155.1 KiB | [postgresql-18-pg-tle_1.5.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-18-pg-tle_1.5.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-tle` | `1.5.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 160.1 KiB | [postgresql-18-pg-tle_1.5.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tle/postgresql-18-pg-tle_1.5.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-tle` | `1.5.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 155.4 KiB | [postgresql-18-pg-tle_1.5.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tle/postgresql-18-pg-tle_1.5.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-tle` | `1.5.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 168.8 KiB | [postgresql-18-pg-tle_1.5.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-18-pg-tle_1.5.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-tle` | `1.5.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 165.3 KiB | [postgresql-18-pg-tle_1.5.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-18-pg-tle_1.5.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-tle` | `1.5.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 162.5 KiB | [postgresql-18-pg-tle_1.5.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-18-pg-tle_1.5.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-tle` | `1.5.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 159.8 KiB | [postgresql-18-pg-tle_1.5.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-18-pg-tle_1.5.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tle_17` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 68.8 KiB | [pg_tle_17-1.5.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_17-1.5.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_tle_17` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.2 KiB | [pg_tle_17-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_tle_17-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_17` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.1 KiB | [pg_tle_17-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_tle_17-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_17` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.1 KiB | [pg_tle_17-1.2.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_tle_17-1.2.0-2PGDG.rhel8.x86_64.rpm) |
| `pg_tle_17` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 65.4 KiB | [pg_tle_17-1.5.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_17-1.5.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_tle_17` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.8 KiB | [pg_tle_17-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_tle_17-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_17` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.7 KiB | [pg_tle_17-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_tle_17-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_17` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.4 KiB | [pg_tle_17-1.2.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_tle_17-1.2.0-2PGDG.rhel8.aarch64.rpm) |
| `pg_tle_17` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 63.1 KiB | [pg_tle_17-1.5.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_17-1.5.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_tle_17` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.7 KiB | [pg_tle_17-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_tle_17-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_17` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.5 KiB | [pg_tle_17-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_tle_17-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_17` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 59.4 KiB | [pg_tle_17-1.2.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_tle_17-1.2.0-2PGDG.rhel9.x86_64.rpm) |
| `pg_tle_17` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 60.0 KiB | [pg_tle_17-1.5.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_17-1.5.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_tle_17` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.8 KiB | [pg_tle_17-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_tle_17-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_17` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.7 KiB | [pg_tle_17-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_tle_17-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_17` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.2 KiB | [pg_tle_17-1.2.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_tle_17-1.2.0-2PGDG.rhel9.aarch64.rpm) |
| `pg_tle_17` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 63.2 KiB | [pg_tle_17-1.5.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tle_17-1.5.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_tle_17` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.1 KiB | [pg_tle_17-1.5.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_tle_17-1.5.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_tle_17` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.6 KiB | [pg_tle_17-1.5.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tle_17-1.5.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_tle_17` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 62.8 KiB | [pg_tle_17-1.5.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_tle_17-1.5.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-tle` | `1.5.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 159.6 KiB | [postgresql-17-pg-tle_1.5.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-tle` | `1.5.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 154.8 KiB | [postgresql-17-pg-tle_1.5.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-tle` | `1.5.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 159.8 KiB | [postgresql-17-pg-tle_1.5.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-tle` | `1.5.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 155.2 KiB | [postgresql-17-pg-tle_1.5.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-tle` | `1.5.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 183.7 KiB | [postgresql-17-pg-tle_1.5.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-tle` | `1.5.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 179.8 KiB | [postgresql-17-pg-tle_1.5.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-tle` | `1.5.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 162.3 KiB | [postgresql-17-pg-tle_1.5.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-tle` | `1.5.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 159.8 KiB | [postgresql-17-pg-tle_1.5.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tle_16` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 68.8 KiB | [pg_tle_16-1.5.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_16-1.5.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_tle_16` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.2 KiB | [pg_tle_16-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_tle_16-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_16` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.1 KiB | [pg_tle_16-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_tle_16-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.0 KiB | [pg_tle_16-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_tle_16-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_16` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 65.4 KiB | [pg_tle_16-1.5.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_16-1.5.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_tle_16` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.8 KiB | [pg_tle_16-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_tle_16-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_16` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.7 KiB | [pg_tle_16-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_tle_16-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.3 KiB | [pg_tle_16-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_tle_16-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_16` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 63.1 KiB | [pg_tle_16-1.5.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_16-1.5.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_tle_16` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.7 KiB | [pg_tle_16-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_tle_16-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_16` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.6 KiB | [pg_tle_16-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_tle_16-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 59.2 KiB | [pg_tle_16-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_tle_16-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_16` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 60.0 KiB | [pg_tle_16-1.5.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_16-1.5.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_tle_16` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.9 KiB | [pg_tle_16-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_tle_16-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_16` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.7 KiB | [pg_tle_16-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_tle_16-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.2 KiB | [pg_tle_16-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_tle_16-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_16` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 63.7 KiB | [pg_tle_16-1.5.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tle_16-1.5.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_tle_16` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.2 KiB | [pg_tle_16-1.5.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_tle_16-1.5.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_tle_16` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.6 KiB | [pg_tle_16-1.5.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tle_16-1.5.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_tle_16` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 62.8 KiB | [pg_tle_16-1.5.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_tle_16-1.5.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-tle` | `1.5.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 159.6 KiB | [postgresql-16-pg-tle_1.5.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-tle` | `1.5.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 155.1 KiB | [postgresql-16-pg-tle_1.5.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-tle` | `1.5.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 160.0 KiB | [postgresql-16-pg-tle_1.5.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-tle` | `1.5.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 155.4 KiB | [postgresql-16-pg-tle_1.5.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-tle` | `1.5.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 183.5 KiB | [postgresql-16-pg-tle_1.5.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-tle` | `1.5.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 179.7 KiB | [postgresql-16-pg-tle_1.5.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-tle` | `1.5.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 162.4 KiB | [postgresql-16-pg-tle_1.5.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-tle` | `1.5.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 159.9 KiB | [postgresql-16-pg-tle_1.5.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tle_15` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 69.9 KiB | [pg_tle_15-1.5.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_15-1.5.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_tle_15` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.1 KiB | [pg_tle_15-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_tle_15-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_15` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.0 KiB | [pg_tle_15-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_tle_15-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.9 KiB | [pg_tle_15-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_tle_15-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_15` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 66.4 KiB | [pg_tle_15-1.5.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_15-1.5.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_tle_15` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.5 KiB | [pg_tle_15-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_tle_15-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_15` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.4 KiB | [pg_tle_15-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_tle_15-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 60.2 KiB | [pg_tle_15-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_tle_15-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_15` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 69.1 KiB | [pg_tle_15-1.5.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_15-1.5.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_tle_15` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 70.5 KiB | [pg_tle_15-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_tle_15-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_15` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 70.8 KiB | [pg_tle_15-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_tle_15-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.5 KiB | [pg_tle_15-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_tle_15-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_15` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 66.5 KiB | [pg_tle_15-1.5.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_15-1.5.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_tle_15` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.1 KiB | [pg_tle_15-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_tle_15-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_15` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.9 KiB | [pg_tle_15-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_tle_15-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.3 KiB | [pg_tle_15-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_tle_15-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_15` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 69.8 KiB | [pg_tle_15-1.5.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tle_15-1.5.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_tle_15` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 71.7 KiB | [pg_tle_15-1.5.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_tle_15-1.5.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_tle_15` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 67.0 KiB | [pg_tle_15-1.5.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tle_15-1.5.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_tle_15` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 68.8 KiB | [pg_tle_15-1.5.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_tle_15-1.5.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-tle` | `1.5.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 161.2 KiB | [postgresql-15-pg-tle_1.5.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-tle` | `1.5.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 156.2 KiB | [postgresql-15-pg-tle_1.5.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-tle` | `1.5.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 161.8 KiB | [postgresql-15-pg-tle_1.5.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-tle` | `1.5.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 156.5 KiB | [postgresql-15-pg-tle_1.5.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-tle` | `1.5.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 189.9 KiB | [postgresql-15-pg-tle_1.5.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-tle` | `1.5.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 186.0 KiB | [postgresql-15-pg-tle_1.5.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-tle` | `1.5.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 168.5 KiB | [postgresql-15-pg-tle_1.5.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-tle` | `1.5.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 166.2 KiB | [postgresql-15-pg-tle_1.5.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tle_14` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 70.0 KiB | [pg_tle_14-1.5.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_14-1.5.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_tle_14` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.3 KiB | [pg_tle_14-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_tle_14-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_14` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.1 KiB | [pg_tle_14-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_tle_14-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 64.1 KiB | [pg_tle_14-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_tle_14-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_14` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 66.5 KiB | [pg_tle_14-1.5.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_14-1.5.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_tle_14` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.7 KiB | [pg_tle_14-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_tle_14-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_14` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.5 KiB | [pg_tle_14-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_tle_14-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_14` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 60.3 KiB | [pg_tle_14-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_tle_14-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_14` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 69.9 KiB | [pg_tle_14-1.5.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_14-1.5.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_tle_14` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 70.7 KiB | [pg_tle_14-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_tle_14-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_14` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 70.6 KiB | [pg_tle_14-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_tle_14-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_14` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.6 KiB | [pg_tle_14-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_tle_14-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_14` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 66.5 KiB | [pg_tle_14-1.5.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_14-1.5.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_tle_14` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.3 KiB | [pg_tle_14-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_tle_14-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_14` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.1 KiB | [pg_tle_14-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_tle_14-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_14` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.5 KiB | [pg_tle_14-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_tle_14-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_14` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 70.2 KiB | [pg_tle_14-1.5.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tle_14-1.5.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_tle_14` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 71.9 KiB | [pg_tle_14-1.5.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_tle_14-1.5.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_tle_14` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 67.1 KiB | [pg_tle_14-1.5.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tle_14-1.5.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_tle_14` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 69.1 KiB | [pg_tle_14-1.5.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_tle_14-1.5.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-tle` | `1.5.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 162.1 KiB | [postgresql-14-pg-tle_1.5.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-tle` | `1.5.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 156.9 KiB | [postgresql-14-pg-tle_1.5.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-tle` | `1.5.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 162.1 KiB | [postgresql-14-pg-tle_1.5.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-tle` | `1.5.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 157.0 KiB | [postgresql-14-pg-tle_1.5.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-tle` | `1.5.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 189.9 KiB | [postgresql-14-pg-tle_1.5.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-tle` | `1.5.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 186.0 KiB | [postgresql-14-pg-tle_1.5.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-tle` | `1.5.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 169.3 KiB | [postgresql-14-pg-tle_1.5.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-tle` | `1.5.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 167.0 KiB | [postgresql-14-pg-tle_1.5.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tle_13` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 69.0 KiB | [pg_tle_13-1.5.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_13-1.5.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_tle_13` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.4 KiB | [pg_tle_13-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_tle_13-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_13` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.2 KiB | [pg_tle_13-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_tle_13-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_13` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.2 KiB | [pg_tle_13-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_tle_13-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_13` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 66.6 KiB | [pg_tle_13-1.5.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_13-1.5.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_tle_13` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.7 KiB | [pg_tle_13-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_tle_13-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_13` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 65.6 KiB | [pg_tle_13-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_tle_13-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_13` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 60.4 KiB | [pg_tle_13-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_tle_13-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_13` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 69.5 KiB | [pg_tle_13-1.5.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_13-1.5.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_tle_13` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 71.3 KiB | [pg_tle_13-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_tle_13-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_13` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 71.2 KiB | [pg_tle_13-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_tle_13-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_13` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.8 KiB | [pg_tle_13-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_tle_13-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_13` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 66.8 KiB | [pg_tle_13-1.5.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_13-1.5.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_tle_13` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.5 KiB | [pg_tle_13-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_tle_13-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_13` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.4 KiB | [pg_tle_13-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_tle_13-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_13` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.9 KiB | [pg_tle_13-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_tle_13-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_13` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 70.1 KiB | [pg_tle_13-1.5.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tle_13-1.5.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_tle_13` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 72.0 KiB | [pg_tle_13-1.5.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_tle_13-1.5.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_tle_13` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 67.3 KiB | [pg_tle_13-1.5.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tle_13-1.5.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_tle_13` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 69.4 KiB | [pg_tle_13-1.5.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_tle_13-1.5.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-tle` | `1.5.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 162.2 KiB | [postgresql-13-pg-tle_1.5.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-tle` | `1.5.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 157.1 KiB | [postgresql-13-pg-tle_1.5.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-tle` | `1.5.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 162.5 KiB | [postgresql-13-pg-tle_1.5.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-tle` | `1.5.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 157.7 KiB | [postgresql-13-pg-tle_1.5.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-tle` | `1.5.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 189.5 KiB | [postgresql-13-pg-tle_1.5.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-tle` | `1.5.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 186.0 KiB | [postgresql-13-pg-tle_1.5.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-tle` | `1.5.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 169.7 KiB | [postgresql-13-pg-tle_1.5.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-tle` | `1.5.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 167.2 KiB | [postgresql-13-pg-tle_1.5.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aws/pg_tle" title="Repository" icon="github" subtitle="github.com/aws/pg_tle" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tle-1.5.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_tle;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_tle;		# install via package name, for the active PG version

pig install pg_tle -v 18;   # install for PG 18
pig install pg_tle -v 17;   # install for PG 17
pig install pg_tle -v 16;   # install for PG 16
pig install pg_tle -v 15;   # install for PG 15
pig install pg_tle -v 14;   # install for PG 14
pig install pg_tle -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_tle';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_tle;
```
