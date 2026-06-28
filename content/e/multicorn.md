---
title: "multicorn"
linkTitle: "multicorn"
description: "Fetch foreign data in Python in your PostgreSQL server."
weight: 8510
categories: ["FDW"]
width: full
---

[**multicorn**](https://github.com/pgsql-io/multicorn2) : Fetch foreign data in Python in your PostgreSQL server.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8510** | {{< badge content="multicorn" link="https://github.com/pgsql-io/multicorn2" >}} | {{< ext "multicorn" >}} | `3.2` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wrappers" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} {{< ext "pgspider_ext" >}} {{< ext "mysql_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "mongo_fdw" >}} {{< ext "redis_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `multicorn` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.2` | {{< bg "18" "multicorn2_18" "green" >}} {{< bg "17" "multicorn2_17" "green" >}} {{< bg "16" "multicorn2_16" "green" >}} {{< bg "15" "multicorn2_15" "green" >}} {{< bg "14" "multicorn2_14" "green" >}} | `multicorn2_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2` | {{< bg "18" "postgresql-18-multicorn" "green" >}} {{< bg "17" "postgresql-17-multicorn" "green" >}} {{< bg "16" "postgresql-16-multicorn" "green" >}} {{< bg "15" "postgresql-15-multicorn" "green" >}} {{< bg "14" "postgresql-14-multicorn" "green" >}} | `postgresql-$v-multicorn` | `python3-multicorn` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 3.2" "multicorn2_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_14 : AVAIL 6" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 3.2" "multicorn2_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_14 : AVAIL 6" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.2" "multicorn2_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_14 : AVAIL 7" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.2" "multicorn2_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_14 : AVAIL 7" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.2" "multicorn2_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.2" "multicorn2_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.2" "multicorn2_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-multicorn : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-multicorn : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-multicorn : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-multicorn : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-multicorn : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-multicorn : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-multicorn : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-multicorn : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-multicorn : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.2" "postgresql-18-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-17-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-16-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-15-multicorn : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2" "postgresql-14-multicorn : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `multicorn2_18` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 138.0 KiB | [multicorn2_18-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/multicorn2_18-3.2-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_18` | `3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.5 KiB | [multicorn2_18-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/multicorn2_18-3.1-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_18` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 136.3 KiB | [multicorn2_18-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/multicorn2_18-3.2-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_18` | `3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.9 KiB | [multicorn2_18-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/multicorn2_18-3.1-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_18` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 134.6 KiB | [multicorn2_18-3.2-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/multicorn2_18-3.2-3PGDG.rhel9.8.x86_64.rpm) |
| `multicorn2_18` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 134.6 KiB | [multicorn2_18-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/multicorn2_18-3.2-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_18` | `3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.0 KiB | [multicorn2_18-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/multicorn2_18-3.1-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_18` | `3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 133.8 KiB | [multicorn2_18-3.2-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/multicorn2_18-3.2-3PGDG.rhel9.8.aarch64.rpm) |
| `multicorn2_18` | `3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 135.2 KiB | [multicorn2_18-3.2-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/multicorn2_18-3.2-3PGDG.rhel10.2.x86_64.rpm) |
| `multicorn2_18` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 134.4 KiB | [multicorn2_18-3.2-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/multicorn2_18-3.2-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-multicorn` | `3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 81.5 KiB | [postgresql-18-multicorn_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/multicorn/postgresql-18-multicorn_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-multicorn` | `3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 79.7 KiB | [postgresql-18-multicorn_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/multicorn/postgresql-18-multicorn_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-multicorn` | `3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 82.5 KiB | [postgresql-18-multicorn_3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/multicorn/postgresql-18-multicorn_3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-multicorn` | `3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 80.8 KiB | [postgresql-18-multicorn_3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/multicorn/postgresql-18-multicorn_3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-multicorn` | `3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 87.7 KiB | [postgresql-18-multicorn_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/multicorn/postgresql-18-multicorn_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-multicorn` | `3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 86.9 KiB | [postgresql-18-multicorn_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/multicorn/postgresql-18-multicorn_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-multicorn` | `3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 85.5 KiB | [postgresql-18-multicorn_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/multicorn/postgresql-18-multicorn_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-multicorn` | `3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 84.3 KiB | [postgresql-18-multicorn_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/multicorn/postgresql-18-multicorn_3.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-multicorn` | `3.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 84.8 KiB | [postgresql-18-multicorn_3.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/multicorn/postgresql-18-multicorn_3.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-multicorn` | `3.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 83.7 KiB | [postgresql-18-multicorn_3.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/multicorn/postgresql-18-multicorn_3.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `multicorn2_17` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 138.0 KiB | [multicorn2_17-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/multicorn2_17-3.2-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_17` | `3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.4 KiB | [multicorn2_17-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/multicorn2_17-3.1-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_17` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.6 KiB | [multicorn2_17-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/multicorn2_17-3.0-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_17` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 136.3 KiB | [multicorn2_17-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/multicorn2_17-3.2-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_17` | `3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.9 KiB | [multicorn2_17-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/multicorn2_17-3.1-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_17` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 113.1 KiB | [multicorn2_17-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/multicorn2_17-3.0-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_17` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 134.8 KiB | [multicorn2_17-3.2-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/multicorn2_17-3.2-3PGDG.rhel9.8.x86_64.rpm) |
| `multicorn2_17` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 134.6 KiB | [multicorn2_17-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/multicorn2_17-3.2-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_17` | `3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.0 KiB | [multicorn2_17-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/multicorn2_17-3.1-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_17` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 110.5 KiB | [multicorn2_17-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/multicorn2_17-3.0-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_17` | `3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 133.7 KiB | [multicorn2_17-3.2-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/multicorn2_17-3.2-3PGDG.rhel9.8.aarch64.rpm) |
| `multicorn2_17` | `3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 135.3 KiB | [multicorn2_17-3.2-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/multicorn2_17-3.2-3PGDG.rhel10.2.x86_64.rpm) |
| `multicorn2_17` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 134.4 KiB | [multicorn2_17-3.2-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/multicorn2_17-3.2-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-multicorn` | `3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 81.4 KiB | [postgresql-17-multicorn_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/multicorn/postgresql-17-multicorn_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-multicorn` | `3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 79.4 KiB | [postgresql-17-multicorn_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/multicorn/postgresql-17-multicorn_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-multicorn` | `3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 82.4 KiB | [postgresql-17-multicorn_3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/multicorn/postgresql-17-multicorn_3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-multicorn` | `3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 80.7 KiB | [postgresql-17-multicorn_3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/multicorn/postgresql-17-multicorn_3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-multicorn` | `3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 106.5 KiB | [postgresql-17-multicorn_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/multicorn/postgresql-17-multicorn_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-multicorn` | `3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 105.6 KiB | [postgresql-17-multicorn_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/multicorn/postgresql-17-multicorn_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-multicorn` | `3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 85.4 KiB | [postgresql-17-multicorn_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/multicorn/postgresql-17-multicorn_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-multicorn` | `3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 83.9 KiB | [postgresql-17-multicorn_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/multicorn/postgresql-17-multicorn_3.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-multicorn` | `3.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 84.7 KiB | [postgresql-17-multicorn_3.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/multicorn/postgresql-17-multicorn_3.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-multicorn` | `3.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 83.4 KiB | [postgresql-17-multicorn_3.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/multicorn/postgresql-17-multicorn_3.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `multicorn2_16` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 138.0 KiB | [multicorn2_16-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/multicorn2_16-3.2-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_16` | `3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.4 KiB | [multicorn2_16-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/multicorn2_16-3.1-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_16` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.6 KiB | [multicorn2_16-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/multicorn2_16-3.0-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_16` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 136.4 KiB | [multicorn2_16-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/multicorn2_16-3.2-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_16` | `3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 134.0 KiB | [multicorn2_16-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/multicorn2_16-3.1-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_16` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 113.2 KiB | [multicorn2_16-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/multicorn2_16-3.0-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_16` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 134.6 KiB | [multicorn2_16-3.2-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/multicorn2_16-3.2-3PGDG.rhel9.8.x86_64.rpm) |
| `multicorn2_16` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 134.6 KiB | [multicorn2_16-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/multicorn2_16-3.2-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_16` | `3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.0 KiB | [multicorn2_16-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/multicorn2_16-3.1-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_16` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 110.6 KiB | [multicorn2_16-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/multicorn2_16-3.0-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_16` | `3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 133.8 KiB | [multicorn2_16-3.2-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/multicorn2_16-3.2-3PGDG.rhel9.8.aarch64.rpm) |
| `multicorn2_16` | `3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 135.3 KiB | [multicorn2_16-3.2-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/multicorn2_16-3.2-3PGDG.rhel10.2.x86_64.rpm) |
| `multicorn2_16` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 134.7 KiB | [multicorn2_16-3.2-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/multicorn2_16-3.2-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-multicorn` | `3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 81.4 KiB | [postgresql-16-multicorn_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/multicorn/postgresql-16-multicorn_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-multicorn` | `3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 79.4 KiB | [postgresql-16-multicorn_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/multicorn/postgresql-16-multicorn_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-multicorn` | `3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 82.6 KiB | [postgresql-16-multicorn_3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/multicorn/postgresql-16-multicorn_3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-multicorn` | `3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 80.7 KiB | [postgresql-16-multicorn_3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/multicorn/postgresql-16-multicorn_3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-multicorn` | `3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 106.4 KiB | [postgresql-16-multicorn_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/multicorn/postgresql-16-multicorn_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-multicorn` | `3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 105.2 KiB | [postgresql-16-multicorn_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/multicorn/postgresql-16-multicorn_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-multicorn` | `3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 85.6 KiB | [postgresql-16-multicorn_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/multicorn/postgresql-16-multicorn_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-multicorn` | `3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 84.3 KiB | [postgresql-16-multicorn_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/multicorn/postgresql-16-multicorn_3.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-multicorn` | `3.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 84.9 KiB | [postgresql-16-multicorn_3.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/multicorn/postgresql-16-multicorn_3.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-multicorn` | `3.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 83.6 KiB | [postgresql-16-multicorn_3.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/multicorn/postgresql-16-multicorn_3.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `multicorn2_15` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 139.3 KiB | [multicorn2_15-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/multicorn2_15-3.2-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_15` | `3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 136.8 KiB | [multicorn2_15-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/multicorn2_15-3.1-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_15` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.9 KiB | [multicorn2_15-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/multicorn2_15-3.0-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_15` | `2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 111.5 KiB | [multicorn2_15-2.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/multicorn2_15-2.4-2.rhel8.x86_64.rpm) |
| `multicorn2_15` | `2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.8 KiB | [multicorn2_15-2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/multicorn2_15-2.4-1.rhel8.x86_64.rpm) |
| `multicorn2_15` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 137.7 KiB | [multicorn2_15-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/multicorn2_15-3.2-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_15` | `3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 135.2 KiB | [multicorn2_15-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/multicorn2_15-3.1-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_15` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 114.2 KiB | [multicorn2_15-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/multicorn2_15-3.0-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_15` | `2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 110.0 KiB | [multicorn2_15-2.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/multicorn2_15-2.4-2.rhel8.aarch64.rpm) |
| `multicorn2_15` | `2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.4 KiB | [multicorn2_15-2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/multicorn2_15-2.4-1.rhel8.aarch64.rpm) |
| `multicorn2_15` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 138.5 KiB | [multicorn2_15-3.2-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/multicorn2_15-3.2-3PGDG.rhel9.8.x86_64.rpm) |
| `multicorn2_15` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 138.5 KiB | [multicorn2_15-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/multicorn2_15-3.2-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_15` | `3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.8 KiB | [multicorn2_15-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/multicorn2_15-3.1-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_15` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.6 KiB | [multicorn2_15-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/multicorn2_15-3.0-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_15` | `2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.9 KiB | [multicorn2_15-2.4-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/multicorn2_15-2.4-2.rhel9.x86_64.rpm) |
| `multicorn2_15` | `2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.4 KiB | [multicorn2_15-2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/multicorn2_15-2.4-1.rhel9.x86_64.rpm) |
| `multicorn2_15` | `3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 136.9 KiB | [multicorn2_15-3.2-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/multicorn2_15-3.2-3PGDG.rhel9.8.aarch64.rpm) |
| `multicorn2_15` | `3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 139.0 KiB | [multicorn2_15-3.2-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/multicorn2_15-3.2-3PGDG.rhel10.2.x86_64.rpm) |
| `multicorn2_15` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 137.6 KiB | [multicorn2_15-3.2-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/multicorn2_15-3.2-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-multicorn` | `3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.5 KiB | [postgresql-15-multicorn_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/multicorn/postgresql-15-multicorn_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-multicorn` | `3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 80.4 KiB | [postgresql-15-multicorn_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/multicorn/postgresql-15-multicorn_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-multicorn` | `3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 83.8 KiB | [postgresql-15-multicorn_3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/multicorn/postgresql-15-multicorn_3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-multicorn` | `3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 81.8 KiB | [postgresql-15-multicorn_3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/multicorn/postgresql-15-multicorn_3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-multicorn` | `3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 108.1 KiB | [postgresql-15-multicorn_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/multicorn/postgresql-15-multicorn_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-multicorn` | `3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 106.8 KiB | [postgresql-15-multicorn_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/multicorn/postgresql-15-multicorn_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-multicorn` | `3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 87.4 KiB | [postgresql-15-multicorn_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/multicorn/postgresql-15-multicorn_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-multicorn` | `3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.7 KiB | [postgresql-15-multicorn_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/multicorn/postgresql-15-multicorn_3.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-multicorn` | `3.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 86.7 KiB | [postgresql-15-multicorn_3.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/multicorn/postgresql-15-multicorn_3.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-multicorn` | `3.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 85.2 KiB | [postgresql-15-multicorn_3.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/multicorn/postgresql-15-multicorn_3.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `multicorn2_14` | `3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 139.4 KiB | [multicorn2_14-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-3.2-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_14` | `3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 136.8 KiB | [multicorn2_14-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-3.1-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_14` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.9 KiB | [multicorn2_14-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-3.0-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_14` | `2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 111.5 KiB | [multicorn2_14-2.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-2.4-2.rhel8.x86_64.rpm) |
| `multicorn2_14` | `2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.8 KiB | [multicorn2_14-2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-2.4-1.rhel8.x86_64.rpm) |
| `multicorn2_14` | `2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.3 KiB | [multicorn2_14-2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-2.3-1.rhel8.x86_64.rpm) |
| `multicorn2_14` | `3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 137.7 KiB | [multicorn2_14-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-3.2-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_14` | `3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 135.2 KiB | [multicorn2_14-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-3.1-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_14` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 114.2 KiB | [multicorn2_14-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-3.0-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_14` | `2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 110.0 KiB | [multicorn2_14-2.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-2.4-2.rhel8.aarch64.rpm) |
| `multicorn2_14` | `2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.4 KiB | [multicorn2_14-2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-2.4-1.rhel8.aarch64.rpm) |
| `multicorn2_14` | `2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 113.9 KiB | [multicorn2_14-2.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-2.3-1.rhel8.aarch64.rpm) |
| `multicorn2_14` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 138.4 KiB | [multicorn2_14-3.2-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-3.2-3PGDG.rhel9.8.x86_64.rpm) |
| `multicorn2_14` | `3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 138.5 KiB | [multicorn2_14-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-3.2-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_14` | `3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.8 KiB | [multicorn2_14-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-3.1-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_14` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.2 KiB | [multicorn2_14-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-3.0-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_14` | `2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.9 KiB | [multicorn2_14-2.4-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-2.4-2.rhel9.x86_64.rpm) |
| `multicorn2_14` | `2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.6 KiB | [multicorn2_14-2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-2.4-1.rhel9.x86_64.rpm) |
| `multicorn2_14` | `2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.0 KiB | [multicorn2_14-2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-2.3-1.rhel9.x86_64.rpm) |
| `multicorn2_14` | `3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 136.9 KiB | [multicorn2_14-3.2-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-3.2-3PGDG.rhel9.8.aarch64.rpm) |
| `multicorn2_14` | `3.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 136.7 KiB | [multicorn2_14-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-3.2-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_14` | `3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 113.3 KiB | [multicorn2_14-3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-3.1-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_14` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 112.8 KiB | [multicorn2_14-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-3.0-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_14` | `2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 108.3 KiB | [multicorn2_14-2.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-2.4-2.rhel9.aarch64.rpm) |
| `multicorn2_14` | `2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.7 KiB | [multicorn2_14-2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-2.4-1.rhel9.aarch64.rpm) |
| `multicorn2_14` | `2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 113.1 KiB | [multicorn2_14-2.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-2.3-1.rhel9.aarch64.rpm) |
| `multicorn2_14` | `3.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 139.0 KiB | [multicorn2_14-3.2-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/multicorn2_14-3.2-3PGDG.rhel10.2.x86_64.rpm) |
| `multicorn2_14` | `3.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 137.6 KiB | [multicorn2_14-3.2-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/multicorn2_14-3.2-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-multicorn` | `3.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.5 KiB | [postgresql-14-multicorn_3.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/multicorn/postgresql-14-multicorn_3.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-multicorn` | `3.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 80.5 KiB | [postgresql-14-multicorn_3.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/multicorn/postgresql-14-multicorn_3.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-multicorn` | `3.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 83.6 KiB | [postgresql-14-multicorn_3.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/multicorn/postgresql-14-multicorn_3.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-multicorn` | `3.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 81.6 KiB | [postgresql-14-multicorn_3.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/multicorn/postgresql-14-multicorn_3.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-multicorn` | `3.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 108.2 KiB | [postgresql-14-multicorn_3.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/multicorn/postgresql-14-multicorn_3.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-multicorn` | `3.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 107.0 KiB | [postgresql-14-multicorn_3.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/multicorn/postgresql-14-multicorn_3.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-multicorn` | `3.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 87.4 KiB | [postgresql-14-multicorn_3.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/multicorn/postgresql-14-multicorn_3.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-multicorn` | `3.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.7 KiB | [postgresql-14-multicorn_3.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/multicorn/postgresql-14-multicorn_3.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-multicorn` | `3.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 86.6 KiB | [postgresql-14-multicorn_3.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/multicorn/postgresql-14-multicorn_3.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-multicorn` | `3.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 85.1 KiB | [postgresql-14-multicorn_3.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/multicorn/postgresql-14-multicorn_3.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgsql-io/multicorn2" title="Repository" icon="github" subtitle="github.com/pgsql-io/multicorn2" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="multicorn2-3.2.tar.gz" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install multicorn;		# install via package name, for the active PG version

pig install multicorn -v 18;   # install for PG 18
pig install multicorn -v 17;   # install for PG 17
pig install multicorn -v 16;   # install for PG 16
pig install multicorn -v 15;   # install for PG 15
pig install multicorn -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION multicorn;
```




## Usage

> [multicorn: Fetch foreign data in Python in your PostgreSQL server](https://github.com/pgsql-io/multicorn2)

Multicorn2 allows you to write Foreign Data Wrappers in Python. You implement a Python class that inherits from `multicorn.ForeignDataWrapper`, and Multicorn handles bridging it to PostgreSQL's FDW interface.

### Define a Python FDW Class

Create a Python module (e.g., `myfdw.py`) accessible to the PostgreSQL process:

```python
from multicorn import ForeignDataWrapper

class MyFDW(ForeignDataWrapper):
    def __init__(self, options, columns):
        super().__init__(options, columns)
        self.options = options
        self.columns = columns

    def execute(self, quals, columns):
        """Yield rows as dictionaries. quals contains WHERE pushdown info."""
        yield {"id": 1, "name": "example"}

    def insert(self, new_values):
        """Handle INSERT operations."""
        pass

    def update(self, old_values, new_values):
        """Handle UPDATE operations."""
        pass

    def delete(self, old_values):
        """Handle DELETE operations."""
        pass
```

### Create Server and Foreign Table

```sql
CREATE EXTENSION multicorn;

CREATE SERVER multicorn_srv FOREIGN DATA WRAPPER multicorn
  OPTIONS (wrapper 'myfdw.MyFDW');

CREATE FOREIGN TABLE my_table (
  id integer,
  name text
)
SERVER multicorn_srv
OPTIONS (
  option1 'value1'
);

SELECT * FROM my_table;
```

The `wrapper` option specifies the fully qualified Python class name. Any additional options are passed to the class constructor's `options` parameter.

### Built-in FDW Examples

Multicorn ships with several example FDW implementations that can be used directly or as reference:

- **CsvFdw** -- read CSV files
- **ProcessFdw** -- execute system commands and parse output
- **GCalFdw** -- access Google Calendar
- **ImapFdw** -- query IMAP mailboxes
- **RssFdw** -- read RSS/Atom feeds

```sql
CREATE SERVER csv_srv FOREIGN DATA WRAPPER multicorn
  OPTIONS (wrapper 'multicorn.csvfdw.CsvFdw');

CREATE FOREIGN TABLE csvtest (
  col1 text,
  col2 text
)
SERVER csv_srv
OPTIONS (
  filename '/tmp/data.csv',
  skip_header '1',
  delimiter ','
);
```
