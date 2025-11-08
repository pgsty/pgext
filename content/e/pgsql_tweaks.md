---
title: "pgsql_tweaks"
linkTitle: "pgsql_tweaks"
description: "Some functions and views for daily usage"
weight: 4200
categories: ["UTIL"]
width: full
---

Some functions and views for daily usage


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4200** | {{< badge content="pgsql_tweaks" link="https://codeberg.org/pgsql_tweaks/pgsql_tweaks" >}} | {{< ext "pgsql_tweaks" >}} | `1.0.2` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_extra_time" >}} {{< ext "extra_window_functions" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgsql_tweaks" >}} | `0.11.3` | {{< bg "18" "pgsql_tweaks_18" "green" >}} {{< bg "17" "pgsql_tweaks_17" "green" >}} {{< bg "16" "pgsql_tweaks_16" "green" >}} {{< bg "15" "pgsql_tweaks_15" "green" >}} {{< bg "14" "pgsql_tweaks_14" "green" >}} {{< bg "13" "pgsql_tweaks_13" "green" >}} | `pgsql_tweaks_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgsql_tweaks" >}} | `1.0.2` | {{< bg "18" "postgresql-18-pgsql-tweaks" "green" >}} {{< bg "17" "postgresql-17-pgsql-tweaks" "green" >}} {{< bg "16" "postgresql-16-pgsql-tweaks" "green" >}} {{< bg "15" "postgresql-15-pgsql-tweaks" "green" >}} {{< bg "14" "postgresql-14-pgsql-tweaks" "green" >}} {{< bg "13" "postgresql-13-pgsql-tweaks" "green" >}} | `postgresql-$v-pgsql-tweaks` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_14 : AVAIL 11" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_13 : AVAIL 11" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_14 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_13 : AVAIL 9" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_13 : AVAIL 10" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_14 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_13 : AVAIL 9" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_13 : AVAIL 4" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.0.2" "pgsql_tweaks_13 : AVAIL 4" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-pgsql-tweaks : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgsql-tweaks : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-13-pgsql-tweaks : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_tweaks_18` | 1.0.2 | `el8.x86_64` | pgdg | 29.1 KiB | [pgsql_tweaks_18-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsql_tweaks_18-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_18` | 1.0.2 | `el8.aarch64` | pgdg | 29.1 KiB | [pgsql_tweaks_18-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsql_tweaks_18-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_18` | 1.0.2 | `el9.x86_64` | pgdg | 27.8 KiB | [pgsql_tweaks_18-1.0.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsql_tweaks_18-1.0.2-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_18` | 1.0.2 | `el9.aarch64` | pgdg | 27.7 KiB | [pgsql_tweaks_18-1.0.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsql_tweaks_18-1.0.2-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_18` | 1.0.2 | `el10.x86_64` | pgdg | 28.3 KiB | [pgsql_tweaks_18-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsql_tweaks_18-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_18` | 1.0.2 | `el10.aarch64` | pgdg | 28.2 KiB | [pgsql_tweaks_18-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsql_tweaks_18-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pgsql-tweaks` | 1.0.2 | `d12.x86_64` | pigsty | 20.4 KiB | [postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgsql-tweaks` | 1.0.2 | `d12.aarch64` | pigsty | 20.4 KiB | [postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgsql-tweaks` | 1.0.2 | `d13.x86_64` | pigsty | 20.4 KiB | [postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgsql-tweaks` | 1.0.2 | `d13.aarch64` | pigsty | 20.4 KiB | [postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgsql-tweaks` | 1.0.2 | `u22.x86_64` | pigsty | 20.7 KiB | [postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgsql-tweaks` | 1.0.2 | `u22.aarch64` | pigsty | 20.7 KiB | [postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgsql-tweaks` | 1.0.2 | `u24.x86_64` | pigsty | 20.7 KiB | [postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgsql-tweaks` | 1.0.2 | `u24.aarch64` | pigsty | 20.7 KiB | [postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-18-pgsql-tweaks_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_tweaks_17` | 1.0.2 | `el8.x86_64` | pgdg | 29.1 KiB | [pgsql_tweaks_17-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 1.0.0 | `el8.x86_64` | pgdg | 29.1 KiB | [pgsql_tweaks_17-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 0.11.3 | `el8.x86_64` | pgdg | 27.1 KiB | [pgsql_tweaks_17-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 0.11.0 | `el8.x86_64` | pgdg | 27.0 KiB | [pgsql_tweaks_17-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 0.10.7 | `el8.x86_64` | pgdg | 24.7 KiB | [pgsql_tweaks_17-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 0.10.6 | `el8.x86_64` | pgdg | 24.0 KiB | [pgsql_tweaks_17-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 0.10.3 | `el8.x86_64` | pgdg | 23.7 KiB | [pgsql_tweaks_17-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_tweaks_17-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 1.0.2 | `el8.aarch64` | pgdg | 29.1 KiB | [pgsql_tweaks_17-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 1.0.0 | `el8.aarch64` | pgdg | 29.0 KiB | [pgsql_tweaks_17-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 0.11.3 | `el8.aarch64` | pgdg | 27.0 KiB | [pgsql_tweaks_17-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 0.11.0 | `el8.aarch64` | pgdg | 27.0 KiB | [pgsql_tweaks_17-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 0.10.7 | `el8.aarch64` | pgdg | 24.7 KiB | [pgsql_tweaks_17-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 0.10.6 | `el8.aarch64` | pgdg | 24.0 KiB | [pgsql_tweaks_17-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 0.10.3 | `el8.aarch64` | pgdg | 23.6 KiB | [pgsql_tweaks_17-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_tweaks_17-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_17` | 1.0.2 | `el9.x86_64` | pgdg | 27.8 KiB | [pgsql_tweaks_17-1.0.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_tweaks_17-1.0.2-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 1.0.0 | `el9.x86_64` | pgdg | 27.7 KiB | [pgsql_tweaks_17-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_tweaks_17-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 0.11.3 | `el9.x86_64` | pgdg | 26.1 KiB | [pgsql_tweaks_17-0.11.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_tweaks_17-0.11.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 0.11.0 | `el9.x86_64` | pgdg | 26.0 KiB | [pgsql_tweaks_17-0.11.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_tweaks_17-0.11.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 0.10.7 | `el9.x86_64` | pgdg | 24.0 KiB | [pgsql_tweaks_17-0.10.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_tweaks_17-0.10.7-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 0.10.6 | `el9.x86_64` | pgdg | 23.3 KiB | [pgsql_tweaks_17-0.10.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_tweaks_17-0.10.6-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 0.10.3 | `el9.x86_64` | pgdg | 23.1 KiB | [pgsql_tweaks_17-0.10.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_tweaks_17-0.10.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 1.0.2 | `el9.aarch64` | pgdg | 27.7 KiB | [pgsql_tweaks_17-1.0.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_tweaks_17-1.0.2-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 1.0.0 | `el9.aarch64` | pgdg | 27.7 KiB | [pgsql_tweaks_17-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_tweaks_17-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 0.11.3 | `el9.aarch64` | pgdg | 26.0 KiB | [pgsql_tweaks_17-0.11.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_tweaks_17-0.11.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 0.11.0 | `el9.aarch64` | pgdg | 25.9 KiB | [pgsql_tweaks_17-0.11.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_tweaks_17-0.11.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 0.10.7 | `el9.aarch64` | pgdg | 23.9 KiB | [pgsql_tweaks_17-0.10.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_tweaks_17-0.10.7-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 0.10.6 | `el9.aarch64` | pgdg | 23.2 KiB | [pgsql_tweaks_17-0.10.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_tweaks_17-0.10.6-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 0.10.3 | `el9.aarch64` | pgdg | 23.0 KiB | [pgsql_tweaks_17-0.10.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_tweaks_17-0.10.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_17` | 1.0.2 | `el10.x86_64` | pgdg | 28.3 KiB | [pgsql_tweaks_17-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsql_tweaks_17-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_17` | 1.0.0 | `el10.x86_64` | pgdg | 28.2 KiB | [pgsql_tweaks_17-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsql_tweaks_17-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_17` | 0.11.3 | `el10.x86_64` | pgdg | 26.6 KiB | [pgsql_tweaks_17-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsql_tweaks_17-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_17` | 0.11.0 | `el10.x86_64` | pgdg | 26.5 KiB | [pgsql_tweaks_17-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsql_tweaks_17-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_17` | 1.0.2 | `el10.aarch64` | pgdg | 28.2 KiB | [pgsql_tweaks_17-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_17` | 1.0.0 | `el10.aarch64` | pgdg | 28.2 KiB | [pgsql_tweaks_17-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_17` | 0.11.3 | `el10.aarch64` | pgdg | 26.5 KiB | [pgsql_tweaks_17-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_17` | 0.11.0 | `el10.aarch64` | pgdg | 26.5 KiB | [pgsql_tweaks_17-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_tweaks_17-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pgsql-tweaks` | 1.0.2 | `d12.x86_64` | pigsty | 20.4 KiB | [postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgsql-tweaks` | 1.0.2 | `d12.aarch64` | pigsty | 20.4 KiB | [postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgsql-tweaks` | 1.0.2 | `d13.x86_64` | pigsty | 20.4 KiB | [postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgsql-tweaks` | 1.0.2 | `d13.aarch64` | pigsty | 20.4 KiB | [postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgsql-tweaks` | 1.0.2 | `u22.x86_64` | pigsty | 20.7 KiB | [postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgsql-tweaks` | 1.0.2 | `u22.aarch64` | pigsty | 20.7 KiB | [postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgsql-tweaks` | 1.0.2 | `u24.x86_64` | pigsty | 20.7 KiB | [postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgsql-tweaks` | 1.0.2 | `u24.aarch64` | pigsty | 20.7 KiB | [postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-17-pgsql-tweaks_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_tweaks_16` | 1.0.2 | `el8.x86_64` | pgdg | 29.1 KiB | [pgsql_tweaks_16-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 1.0.0 | `el8.x86_64` | pgdg | 29.1 KiB | [pgsql_tweaks_16-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 0.11.3 | `el8.x86_64` | pgdg | 27.1 KiB | [pgsql_tweaks_16-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 0.11.0 | `el8.x86_64` | pgdg | 27.0 KiB | [pgsql_tweaks_16-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.7 | `el8.x86_64` | pgdg | 24.7 KiB | [pgsql_tweaks_16-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.6 | `el8.x86_64` | pgdg | 24.0 KiB | [pgsql_tweaks_16-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.3 | `el8.x86_64` | pgdg | 23.7 KiB | [pgsql_tweaks_16-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.2 | `el8.x86_64` | pgdg | 23.6 KiB | [pgsql_tweaks_16-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.1 | `el8.x86_64` | pgdg | 23.5 KiB | [pgsql_tweaks_16-0.10.1-2.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_tweaks_16-0.10.1-2.rhel7.noarch.rpm) |
| `pgsql_tweaks_16` | 1.0.2 | `el8.aarch64` | pgdg | 29.1 KiB | [pgsql_tweaks_16-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 1.0.0 | `el8.aarch64` | pgdg | 29.0 KiB | [pgsql_tweaks_16-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 0.11.3 | `el8.aarch64` | pgdg | 27.0 KiB | [pgsql_tweaks_16-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 0.11.0 | `el8.aarch64` | pgdg | 27.0 KiB | [pgsql_tweaks_16-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.7 | `el8.aarch64` | pgdg | 24.7 KiB | [pgsql_tweaks_16-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.6 | `el8.aarch64` | pgdg | 24.0 KiB | [pgsql_tweaks_16-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.3 | `el8.aarch64` | pgdg | 23.6 KiB | [pgsql_tweaks_16-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.2 | `el8.aarch64` | pgdg | 23.6 KiB | [pgsql_tweaks_16-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.1 | `el8.aarch64` | pgdg | 23.5 KiB | [pgsql_tweaks_16-0.10.1-2.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_tweaks_16-0.10.1-2.rhel7.noarch.rpm) |
| `pgsql_tweaks_16` | 1.0.2 | `el9.x86_64` | pgdg | 27.8 KiB | [pgsql_tweaks_16-1.0.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_tweaks_16-1.0.2-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 1.0.0 | `el9.x86_64` | pgdg | 27.7 KiB | [pgsql_tweaks_16-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_tweaks_16-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 0.11.3 | `el9.x86_64` | pgdg | 26.1 KiB | [pgsql_tweaks_16-0.11.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_tweaks_16-0.11.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 0.11.0 | `el9.x86_64` | pgdg | 26.0 KiB | [pgsql_tweaks_16-0.11.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_tweaks_16-0.11.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.7 | `el9.x86_64` | pgdg | 24.0 KiB | [pgsql_tweaks_16-0.10.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_tweaks_16-0.10.7-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.6 | `el9.x86_64` | pgdg | 23.3 KiB | [pgsql_tweaks_16-0.10.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_tweaks_16-0.10.6-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.3 | `el9.x86_64` | pgdg | 23.1 KiB | [pgsql_tweaks_16-0.10.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_tweaks_16-0.10.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.2 | `el9.x86_64` | pgdg | 23.0 KiB | [pgsql_tweaks_16-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_tweaks_16-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.1 | `el9.x86_64` | pgdg | 23.1 KiB | [pgsql_tweaks_16-0.10.1-2.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_tweaks_16-0.10.1-2.rhel7.noarch.rpm) |
| `pgsql_tweaks_16` | 1.0.2 | `el9.aarch64` | pgdg | 27.7 KiB | [pgsql_tweaks_16-1.0.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_tweaks_16-1.0.2-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 1.0.0 | `el9.aarch64` | pgdg | 27.7 KiB | [pgsql_tweaks_16-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_tweaks_16-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 0.11.3 | `el9.aarch64` | pgdg | 26.0 KiB | [pgsql_tweaks_16-0.11.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_tweaks_16-0.11.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 0.11.0 | `el9.aarch64` | pgdg | 25.9 KiB | [pgsql_tweaks_16-0.11.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_tweaks_16-0.11.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.7 | `el9.aarch64` | pgdg | 23.9 KiB | [pgsql_tweaks_16-0.10.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_tweaks_16-0.10.7-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.6 | `el9.aarch64` | pgdg | 23.2 KiB | [pgsql_tweaks_16-0.10.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_tweaks_16-0.10.6-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.3 | `el9.aarch64` | pgdg | 23.0 KiB | [pgsql_tweaks_16-0.10.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_tweaks_16-0.10.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.2 | `el9.aarch64` | pgdg | 22.8 KiB | [pgsql_tweaks_16-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_tweaks_16-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_16` | 0.10.1 | `el9.aarch64` | pgdg | 22.9 KiB | [pgsql_tweaks_16-0.10.1-2.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_tweaks_16-0.10.1-2.rhel7.noarch.rpm) |
| `pgsql_tweaks_16` | 1.0.2 | `el10.x86_64` | pgdg | 28.3 KiB | [pgsql_tweaks_16-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsql_tweaks_16-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_16` | 1.0.0 | `el10.x86_64` | pgdg | 28.2 KiB | [pgsql_tweaks_16-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsql_tweaks_16-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_16` | 0.11.3 | `el10.x86_64` | pgdg | 26.6 KiB | [pgsql_tweaks_16-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsql_tweaks_16-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_16` | 0.11.0 | `el10.x86_64` | pgdg | 26.5 KiB | [pgsql_tweaks_16-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsql_tweaks_16-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_16` | 1.0.2 | `el10.aarch64` | pgdg | 28.2 KiB | [pgsql_tweaks_16-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_16` | 1.0.0 | `el10.aarch64` | pgdg | 28.2 KiB | [pgsql_tweaks_16-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_16` | 0.11.3 | `el10.aarch64` | pgdg | 26.5 KiB | [pgsql_tweaks_16-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_16` | 0.11.0 | `el10.aarch64` | pgdg | 26.5 KiB | [pgsql_tweaks_16-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_tweaks_16-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pgsql-tweaks` | 1.0.2 | `d12.x86_64` | pigsty | 20.4 KiB | [postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgsql-tweaks` | 1.0.2 | `d12.aarch64` | pigsty | 20.4 KiB | [postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgsql-tweaks` | 1.0.2 | `d13.x86_64` | pigsty | 20.4 KiB | [postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgsql-tweaks` | 1.0.2 | `d13.aarch64` | pigsty | 20.4 KiB | [postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgsql-tweaks` | 1.0.2 | `u22.x86_64` | pigsty | 20.7 KiB | [postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgsql-tweaks` | 1.0.2 | `u22.aarch64` | pigsty | 20.7 KiB | [postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgsql-tweaks` | 1.0.2 | `u24.x86_64` | pigsty | 20.7 KiB | [postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgsql-tweaks` | 1.0.2 | `u24.aarch64` | pigsty | 20.7 KiB | [postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-16-pgsql-tweaks_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_tweaks_15` | 1.0.2 | `el8.x86_64` | pgdg | 29.1 KiB | [pgsql_tweaks_15-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 1.0.0 | `el8.x86_64` | pgdg | 29.1 KiB | [pgsql_tweaks_15-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 0.11.3 | `el8.x86_64` | pgdg | 27.1 KiB | [pgsql_tweaks_15-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 0.11.0 | `el8.x86_64` | pgdg | 27.0 KiB | [pgsql_tweaks_15-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.7 | `el8.x86_64` | pgdg | 24.7 KiB | [pgsql_tweaks_15-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.6 | `el8.x86_64` | pgdg | 24.0 KiB | [pgsql_tweaks_15-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.3 | `el8.x86_64` | pgdg | 23.7 KiB | [pgsql_tweaks_15-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.2 | `el8.x86_64` | pgdg | 23.6 KiB | [pgsql_tweaks_15-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.0 | `el8.x86_64` | pgdg | 23.4 KiB | [pgsql_tweaks_15-0.10.0-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_tweaks_15-0.10.0-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_15` | 1.0.2 | `el8.aarch64` | pgdg | 29.1 KiB | [pgsql_tweaks_15-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 1.0.0 | `el8.aarch64` | pgdg | 29.0 KiB | [pgsql_tweaks_15-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 0.11.3 | `el8.aarch64` | pgdg | 27.0 KiB | [pgsql_tweaks_15-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 0.11.0 | `el8.aarch64` | pgdg | 26.9 KiB | [pgsql_tweaks_15-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.7 | `el8.aarch64` | pgdg | 24.7 KiB | [pgsql_tweaks_15-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.6 | `el8.aarch64` | pgdg | 24.0 KiB | [pgsql_tweaks_15-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.3 | `el8.aarch64` | pgdg | 23.6 KiB | [pgsql_tweaks_15-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.2 | `el8.aarch64` | pgdg | 23.6 KiB | [pgsql_tweaks_15-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.1 | `el8.aarch64` | pgdg | 23.4 KiB | [pgsql_tweaks_15-0.10.1-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_tweaks_15-0.10.1-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_15` | 1.0.2 | `el9.x86_64` | pgdg | 27.8 KiB | [pgsql_tweaks_15-1.0.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_tweaks_15-1.0.2-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 1.0.0 | `el9.x86_64` | pgdg | 27.7 KiB | [pgsql_tweaks_15-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_tweaks_15-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 0.11.3 | `el9.x86_64` | pgdg | 26.1 KiB | [pgsql_tweaks_15-0.11.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_tweaks_15-0.11.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 0.11.0 | `el9.x86_64` | pgdg | 26.0 KiB | [pgsql_tweaks_15-0.11.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_tweaks_15-0.11.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.7 | `el9.x86_64` | pgdg | 24.0 KiB | [pgsql_tweaks_15-0.10.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_tweaks_15-0.10.7-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.6 | `el9.x86_64` | pgdg | 23.3 KiB | [pgsql_tweaks_15-0.10.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_tweaks_15-0.10.6-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.3 | `el9.x86_64` | pgdg | 23.1 KiB | [pgsql_tweaks_15-0.10.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_tweaks_15-0.10.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.2 | `el9.x86_64` | pgdg | 23.0 KiB | [pgsql_tweaks_15-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_tweaks_15-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.0 | `el9.x86_64` | pgdg | 22.9 KiB | [pgsql_tweaks_15-0.10.0-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_tweaks_15-0.10.0-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_15` | 1.0.2 | `el9.aarch64` | pgdg | 27.7 KiB | [pgsql_tweaks_15-1.0.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_tweaks_15-1.0.2-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 1.0.0 | `el9.aarch64` | pgdg | 27.7 KiB | [pgsql_tweaks_15-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_tweaks_15-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 0.11.3 | `el9.aarch64` | pgdg | 26.0 KiB | [pgsql_tweaks_15-0.11.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_tweaks_15-0.11.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 0.11.0 | `el9.aarch64` | pgdg | 25.9 KiB | [pgsql_tweaks_15-0.11.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_tweaks_15-0.11.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.7 | `el9.aarch64` | pgdg | 23.9 KiB | [pgsql_tweaks_15-0.10.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_tweaks_15-0.10.7-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.6 | `el9.aarch64` | pgdg | 23.2 KiB | [pgsql_tweaks_15-0.10.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_tweaks_15-0.10.6-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.3 | `el9.aarch64` | pgdg | 23.0 KiB | [pgsql_tweaks_15-0.10.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_tweaks_15-0.10.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.2 | `el9.aarch64` | pgdg | 22.8 KiB | [pgsql_tweaks_15-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_tweaks_15-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_15` | 0.10.1 | `el9.aarch64` | pgdg | 22.8 KiB | [pgsql_tweaks_15-0.10.1-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_tweaks_15-0.10.1-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_15` | 1.0.2 | `el10.x86_64` | pgdg | 28.3 KiB | [pgsql_tweaks_15-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsql_tweaks_15-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_15` | 1.0.0 | `el10.x86_64` | pgdg | 28.2 KiB | [pgsql_tweaks_15-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsql_tweaks_15-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_15` | 0.11.3 | `el10.x86_64` | pgdg | 26.6 KiB | [pgsql_tweaks_15-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsql_tweaks_15-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_15` | 0.11.0 | `el10.x86_64` | pgdg | 26.5 KiB | [pgsql_tweaks_15-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsql_tweaks_15-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_15` | 1.0.2 | `el10.aarch64` | pgdg | 28.2 KiB | [pgsql_tweaks_15-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_15` | 1.0.0 | `el10.aarch64` | pgdg | 28.2 KiB | [pgsql_tweaks_15-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_15` | 0.11.3 | `el10.aarch64` | pgdg | 26.5 KiB | [pgsql_tweaks_15-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_15` | 0.11.0 | `el10.aarch64` | pgdg | 26.5 KiB | [pgsql_tweaks_15-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_tweaks_15-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pgsql-tweaks` | 1.0.2 | `d12.x86_64` | pigsty | 20.4 KiB | [postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgsql-tweaks` | 1.0.2 | `d12.aarch64` | pigsty | 20.4 KiB | [postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgsql-tweaks` | 1.0.2 | `d13.x86_64` | pigsty | 20.4 KiB | [postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgsql-tweaks` | 1.0.2 | `d13.aarch64` | pigsty | 20.4 KiB | [postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgsql-tweaks` | 1.0.2 | `u22.x86_64` | pigsty | 20.7 KiB | [postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgsql-tweaks` | 1.0.2 | `u22.aarch64` | pigsty | 20.7 KiB | [postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgsql-tweaks` | 1.0.2 | `u24.x86_64` | pigsty | 20.7 KiB | [postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgsql-tweaks` | 1.0.2 | `u24.aarch64` | pigsty | 20.7 KiB | [postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-15-pgsql-tweaks_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_tweaks_14` | 1.0.2 | `el8.x86_64` | pgdg | 29.1 KiB | [pgsql_tweaks_14-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_tweaks_14-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 1.0.0 | `el8.x86_64` | pgdg | 29.1 KiB | [pgsql_tweaks_14-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_tweaks_14-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 0.9.1 | `el8.x86_64` | pgdg | 22.8 KiB | [pgsql_tweaks_14-0.9.1-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_tweaks_14-0.9.1-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_14` | 0.8.0 | `el8.x86_64` | pgdg | 22.7 KiB | [pgsql_tweaks_14-0.8.0-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_tweaks_14-0.8.0-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_14` | 0.11.3 | `el8.x86_64` | pgdg | 27.1 KiB | [pgsql_tweaks_14-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_tweaks_14-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 0.11.0 | `el8.x86_64` | pgdg | 27.0 KiB | [pgsql_tweaks_14-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_tweaks_14-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.7 | `el8.x86_64` | pgdg | 24.7 KiB | [pgsql_tweaks_14-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_tweaks_14-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.6 | `el8.x86_64` | pgdg | 24.0 KiB | [pgsql_tweaks_14-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_tweaks_14-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.3 | `el8.x86_64` | pgdg | 23.7 KiB | [pgsql_tweaks_14-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_tweaks_14-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.2 | `el8.x86_64` | pgdg | 23.6 KiB | [pgsql_tweaks_14-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_tweaks_14-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.0 | `el8.x86_64` | pgdg | 23.4 KiB | [pgsql_tweaks_14-0.10.0-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_tweaks_14-0.10.0-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_14` | 1.0.2 | `el8.aarch64` | pgdg | 29.1 KiB | [pgsql_tweaks_14-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_tweaks_14-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 1.0.0 | `el8.aarch64` | pgdg | 29.0 KiB | [pgsql_tweaks_14-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_tweaks_14-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 0.11.3 | `el8.aarch64` | pgdg | 27.0 KiB | [pgsql_tweaks_14-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_tweaks_14-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 0.11.0 | `el8.aarch64` | pgdg | 26.9 KiB | [pgsql_tweaks_14-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_tweaks_14-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.7 | `el8.aarch64` | pgdg | 24.7 KiB | [pgsql_tweaks_14-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_tweaks_14-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.6 | `el8.aarch64` | pgdg | 24.0 KiB | [pgsql_tweaks_14-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_tweaks_14-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.3 | `el8.aarch64` | pgdg | 23.6 KiB | [pgsql_tweaks_14-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_tweaks_14-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.2 | `el8.aarch64` | pgdg | 23.6 KiB | [pgsql_tweaks_14-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_tweaks_14-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.1 | `el8.aarch64` | pgdg | 23.4 KiB | [pgsql_tweaks_14-0.10.1-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_tweaks_14-0.10.1-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_14` | 1.0.2 | `el9.x86_64` | pgdg | 27.8 KiB | [pgsql_tweaks_14-1.0.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_tweaks_14-1.0.2-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 1.0.0 | `el9.x86_64` | pgdg | 27.8 KiB | [pgsql_tweaks_14-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_tweaks_14-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 0.9.1 | `el9.x86_64` | pgdg | 22.4 KiB | [pgsql_tweaks_14-0.9.1-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_tweaks_14-0.9.1-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_14` | 0.11.3 | `el9.x86_64` | pgdg | 26.1 KiB | [pgsql_tweaks_14-0.11.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_tweaks_14-0.11.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 0.11.0 | `el9.x86_64` | pgdg | 26.0 KiB | [pgsql_tweaks_14-0.11.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_tweaks_14-0.11.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.7 | `el9.x86_64` | pgdg | 24.0 KiB | [pgsql_tweaks_14-0.10.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_tweaks_14-0.10.7-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.6 | `el9.x86_64` | pgdg | 23.3 KiB | [pgsql_tweaks_14-0.10.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_tweaks_14-0.10.6-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.3 | `el9.x86_64` | pgdg | 23.1 KiB | [pgsql_tweaks_14-0.10.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_tweaks_14-0.10.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.2 | `el9.x86_64` | pgdg | 23.0 KiB | [pgsql_tweaks_14-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_tweaks_14-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.0 | `el9.x86_64` | pgdg | 22.9 KiB | [pgsql_tweaks_14-0.10.0-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_tweaks_14-0.10.0-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_14` | 1.0.2 | `el9.aarch64` | pgdg | 27.7 KiB | [pgsql_tweaks_14-1.0.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_tweaks_14-1.0.2-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 1.0.0 | `el9.aarch64` | pgdg | 27.7 KiB | [pgsql_tweaks_14-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_tweaks_14-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 0.11.3 | `el9.aarch64` | pgdg | 26.0 KiB | [pgsql_tweaks_14-0.11.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_tweaks_14-0.11.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 0.11.0 | `el9.aarch64` | pgdg | 25.9 KiB | [pgsql_tweaks_14-0.11.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_tweaks_14-0.11.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.7 | `el9.aarch64` | pgdg | 23.9 KiB | [pgsql_tweaks_14-0.10.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_tweaks_14-0.10.7-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.6 | `el9.aarch64` | pgdg | 23.2 KiB | [pgsql_tweaks_14-0.10.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_tweaks_14-0.10.6-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.3 | `el9.aarch64` | pgdg | 23.0 KiB | [pgsql_tweaks_14-0.10.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_tweaks_14-0.10.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.2 | `el9.aarch64` | pgdg | 22.8 KiB | [pgsql_tweaks_14-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_tweaks_14-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_14` | 0.10.1 | `el9.aarch64` | pgdg | 22.8 KiB | [pgsql_tweaks_14-0.10.1-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_tweaks_14-0.10.1-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_14` | 1.0.2 | `el10.x86_64` | pgdg | 28.3 KiB | [pgsql_tweaks_14-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsql_tweaks_14-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_14` | 1.0.0 | `el10.x86_64` | pgdg | 28.2 KiB | [pgsql_tweaks_14-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsql_tweaks_14-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_14` | 0.11.3 | `el10.x86_64` | pgdg | 26.6 KiB | [pgsql_tweaks_14-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsql_tweaks_14-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_14` | 0.11.0 | `el10.x86_64` | pgdg | 26.5 KiB | [pgsql_tweaks_14-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsql_tweaks_14-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_14` | 1.0.2 | `el10.aarch64` | pgdg | 28.2 KiB | [pgsql_tweaks_14-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsql_tweaks_14-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_14` | 1.0.0 | `el10.aarch64` | pgdg | 28.2 KiB | [pgsql_tweaks_14-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsql_tweaks_14-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_14` | 0.11.3 | `el10.aarch64` | pgdg | 26.5 KiB | [pgsql_tweaks_14-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsql_tweaks_14-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_14` | 0.11.0 | `el10.aarch64` | pgdg | 26.5 KiB | [pgsql_tweaks_14-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsql_tweaks_14-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-14-pgsql-tweaks` | 1.0.2 | `d12.x86_64` | pigsty | 20.4 KiB | [postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgsql-tweaks` | 1.0.2 | `d12.aarch64` | pigsty | 20.4 KiB | [postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgsql-tweaks` | 1.0.2 | `d13.x86_64` | pigsty | 20.4 KiB | [postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgsql-tweaks` | 1.0.2 | `d13.aarch64` | pigsty | 20.4 KiB | [postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgsql-tweaks` | 1.0.2 | `u22.x86_64` | pigsty | 20.7 KiB | [postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgsql-tweaks` | 1.0.2 | `u22.aarch64` | pigsty | 20.7 KiB | [postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgsql-tweaks` | 1.0.2 | `u24.x86_64` | pigsty | 20.7 KiB | [postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgsql-tweaks` | 1.0.2 | `u24.aarch64` | pigsty | 20.7 KiB | [postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-14-pgsql-tweaks_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_tweaks_13` | 1.0.2 | `el8.x86_64` | pgdg | 29.1 KiB | [pgsql_tweaks_13-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_tweaks_13-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 1.0.0 | `el8.x86_64` | pgdg | 29.1 KiB | [pgsql_tweaks_13-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_tweaks_13-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 0.9.1 | `el8.x86_64` | pgdg | 22.8 KiB | [pgsql_tweaks_13-0.9.1-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_tweaks_13-0.9.1-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_13` | 0.8.0 | `el8.x86_64` | pgdg | 22.7 KiB | [pgsql_tweaks_13-0.8.0-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_tweaks_13-0.8.0-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_13` | 0.11.3 | `el8.x86_64` | pgdg | 27.1 KiB | [pgsql_tweaks_13-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_tweaks_13-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 0.11.0 | `el8.x86_64` | pgdg | 27.0 KiB | [pgsql_tweaks_13-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_tweaks_13-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.7 | `el8.x86_64` | pgdg | 24.7 KiB | [pgsql_tweaks_13-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_tweaks_13-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.6 | `el8.x86_64` | pgdg | 24.0 KiB | [pgsql_tweaks_13-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_tweaks_13-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.3 | `el8.x86_64` | pgdg | 23.7 KiB | [pgsql_tweaks_13-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_tweaks_13-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.2 | `el8.x86_64` | pgdg | 23.6 KiB | [pgsql_tweaks_13-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_tweaks_13-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.0 | `el8.x86_64` | pgdg | 23.4 KiB | [pgsql_tweaks_13-0.10.0-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_tweaks_13-0.10.0-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_13` | 1.0.2 | `el8.aarch64` | pgdg | 29.1 KiB | [pgsql_tweaks_13-1.0.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_tweaks_13-1.0.2-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 1.0.0 | `el8.aarch64` | pgdg | 29.0 KiB | [pgsql_tweaks_13-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_tweaks_13-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 0.11.3 | `el8.aarch64` | pgdg | 27.0 KiB | [pgsql_tweaks_13-0.11.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_tweaks_13-0.11.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 0.11.0 | `el8.aarch64` | pgdg | 26.9 KiB | [pgsql_tweaks_13-0.11.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_tweaks_13-0.11.0-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.7 | `el8.aarch64` | pgdg | 24.7 KiB | [pgsql_tweaks_13-0.10.7-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_tweaks_13-0.10.7-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.6 | `el8.aarch64` | pgdg | 24.0 KiB | [pgsql_tweaks_13-0.10.6-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_tweaks_13-0.10.6-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.3 | `el8.aarch64` | pgdg | 23.6 KiB | [pgsql_tweaks_13-0.10.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_tweaks_13-0.10.3-1PGDG.rhel8.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.2 | `el8.aarch64` | pgdg | 23.6 KiB | [pgsql_tweaks_13-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_tweaks_13-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.1 | `el8.aarch64` | pgdg | 23.4 KiB | [pgsql_tweaks_13-0.10.1-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_tweaks_13-0.10.1-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_13` | 1.0.2 | `el9.x86_64` | pgdg | 27.8 KiB | [pgsql_tweaks_13-1.0.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_tweaks_13-1.0.2-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 1.0.0 | `el9.x86_64` | pgdg | 27.7 KiB | [pgsql_tweaks_13-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_tweaks_13-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 0.9.1 | `el9.x86_64` | pgdg | 22.4 KiB | [pgsql_tweaks_13-0.9.1-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_tweaks_13-0.9.1-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_13` | 0.11.3 | `el9.x86_64` | pgdg | 26.1 KiB | [pgsql_tweaks_13-0.11.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_tweaks_13-0.11.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 0.11.0 | `el9.x86_64` | pgdg | 26.0 KiB | [pgsql_tweaks_13-0.11.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_tweaks_13-0.11.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.7 | `el9.x86_64` | pgdg | 23.9 KiB | [pgsql_tweaks_13-0.10.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_tweaks_13-0.10.7-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.6 | `el9.x86_64` | pgdg | 23.3 KiB | [pgsql_tweaks_13-0.10.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_tweaks_13-0.10.6-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.3 | `el9.x86_64` | pgdg | 23.1 KiB | [pgsql_tweaks_13-0.10.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_tweaks_13-0.10.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.2 | `el9.x86_64` | pgdg | 23.0 KiB | [pgsql_tweaks_13-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_tweaks_13-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.0 | `el9.x86_64` | pgdg | 22.9 KiB | [pgsql_tweaks_13-0.10.0-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_tweaks_13-0.10.0-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_13` | 1.0.2 | `el9.aarch64` | pgdg | 27.7 KiB | [pgsql_tweaks_13-1.0.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_tweaks_13-1.0.2-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 1.0.0 | `el9.aarch64` | pgdg | 27.7 KiB | [pgsql_tweaks_13-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_tweaks_13-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 0.11.3 | `el9.aarch64` | pgdg | 26.0 KiB | [pgsql_tweaks_13-0.11.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_tweaks_13-0.11.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 0.11.0 | `el9.aarch64` | pgdg | 25.9 KiB | [pgsql_tweaks_13-0.11.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_tweaks_13-0.11.0-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.7 | `el9.aarch64` | pgdg | 23.9 KiB | [pgsql_tweaks_13-0.10.7-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_tweaks_13-0.10.7-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.6 | `el9.aarch64` | pgdg | 23.2 KiB | [pgsql_tweaks_13-0.10.6-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_tweaks_13-0.10.6-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.3 | `el9.aarch64` | pgdg | 23.0 KiB | [pgsql_tweaks_13-0.10.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_tweaks_13-0.10.3-1PGDG.rhel9.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.2 | `el9.aarch64` | pgdg | 22.8 KiB | [pgsql_tweaks_13-0.10.2-1PGDG.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_tweaks_13-0.10.2-1PGDG.rhel7.noarch.rpm) |
| `pgsql_tweaks_13` | 0.10.1 | `el9.aarch64` | pgdg | 22.8 KiB | [pgsql_tweaks_13-0.10.1-1.rhel7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_tweaks_13-0.10.1-1.rhel7.noarch.rpm) |
| `pgsql_tweaks_13` | 1.0.2 | `el10.x86_64` | pgdg | 28.3 KiB | [pgsql_tweaks_13-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgsql_tweaks_13-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_13` | 1.0.0 | `el10.x86_64` | pgdg | 28.2 KiB | [pgsql_tweaks_13-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgsql_tweaks_13-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_13` | 0.11.3 | `el10.x86_64` | pgdg | 26.6 KiB | [pgsql_tweaks_13-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgsql_tweaks_13-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_13` | 0.11.0 | `el10.x86_64` | pgdg | 26.5 KiB | [pgsql_tweaks_13-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgsql_tweaks_13-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_13` | 1.0.2 | `el10.aarch64` | pgdg | 28.2 KiB | [pgsql_tweaks_13-1.0.2-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgsql_tweaks_13-1.0.2-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_13` | 1.0.0 | `el10.aarch64` | pgdg | 28.2 KiB | [pgsql_tweaks_13-1.0.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgsql_tweaks_13-1.0.0-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_13` | 0.11.3 | `el10.aarch64` | pgdg | 26.5 KiB | [pgsql_tweaks_13-0.11.3-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgsql_tweaks_13-0.11.3-1PGDG.rhel10.noarch.rpm) |
| `pgsql_tweaks_13` | 0.11.0 | `el10.aarch64` | pgdg | 26.5 KiB | [pgsql_tweaks_13-0.11.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgsql_tweaks_13-0.11.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-13-pgsql-tweaks` | 1.0.2 | `d12.x86_64` | pigsty | 20.4 KiB | [postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgsql-tweaks` | 1.0.2 | `d12.aarch64` | pigsty | 20.4 KiB | [postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-tweaks/postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgsql-tweaks` | 1.0.2 | `d13.x86_64` | pigsty | 20.4 KiB | [postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pgsql-tweaks` | 1.0.2 | `d13.aarch64` | pigsty | 20.4 KiB | [postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-tweaks/postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pgsql-tweaks` | 1.0.2 | `u22.x86_64` | pigsty | 20.7 KiB | [postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgsql-tweaks` | 1.0.2 | `u22.aarch64` | pigsty | 20.7 KiB | [postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-tweaks/postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgsql-tweaks` | 1.0.2 | `u24.x86_64` | pigsty | 20.7 KiB | [postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgsql-tweaks` | 1.0.2 | `u24.aarch64` | pigsty | 20.7 KiB | [postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-tweaks/postgresql-13-pgsql-tweaks_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/pgsql_tweaks/pgsql_tweaks" title="Repository" icon="link" subtitle="codeberg.org/pgsql_tweaks/pgsql_tweaks" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsql_tweaks-v1.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgsql_tweaks; # get pgsql_tweaks source code
pig build dep pgsql_tweaks; # install build dependencies
pig build pkg pgsql_tweaks; # build extension rpm or deb
pig build ext pgsql_tweaks; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgsql_tweaks; # install by extension name, for the current active PG version
pig ext install pgsql_tweaks; # install via package alias, for the active PG version
pig ext install pgsql_tweaks -v 18;   # install for PG 18
pig ext install pgsql_tweaks -v 17;   # install for PG 17
pig ext install pgsql_tweaks -v 16;   # install for PG 16
pig ext install pgsql_tweaks -v 15;   # install for PG 15
pig ext install pgsql_tweaks -v 14;   # install for PG 14
pig ext install pgsql_tweaks -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgsql_tweaks;
```

