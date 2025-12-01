---
title: "safeupdate"
linkTitle: "safeupdate"
description: "Require criteria for UPDATE and DELETE"
weight: 5820
categories: ["ADMIN"]
width: full
---

[**safeupdate**](https://github.com/eradman/pg-safeupdate) : Require criteria for UPDATE and DELETE


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5820** | {{< badge content="safeupdate" link="https://github.com/eradman/pg-safeupdate" >}} | {{< ext "safeupdate" >}} | `1.5` | {{< category "ADMIN" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_readonly" >}} {{< ext "pg_upless" >}} {{< ext "pg_savior" >}} {{< ext "pg_permissions" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "login_hook" >}} {{< ext "noset" >}} |

> [!Note] pg13 breaks on debian systems


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `safeupdate` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5` | {{< bg "18" "safeupdate_18*" "green" >}} {{< bg "17" "safeupdate_17*" "green" >}} {{< bg "16" "safeupdate_16*" "green" >}} {{< bg "15" "safeupdate_15*" "green" >}} {{< bg "14" "safeupdate_14*" "green" >}} {{< bg "13" "safeupdate_13*" "red" >}} | `safeupdate_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5` | {{< bg "18" "postgresql-18-pg-safeupdate" "green" >}} {{< bg "17" "postgresql-17-pg-safeupdate" "green" >}} {{< bg "16" "postgresql-16-pg-safeupdate" "green" >}} {{< bg "15" "postgresql-15-pg-safeupdate" "green" >}} {{< bg "14" "postgresql-14-pg-safeupdate" "green" >}} {{< bg "13" "postgresql-13-pg-safeupdate" "red" >}} | `postgresql-$v-pg-safeupdate` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.5" "safeupdate_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "safeupdate_13 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.5" "safeupdate_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.2" "safeupdate_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.5" "safeupdate_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.2" "safeupdate_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.5" "safeupdate_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.2" "safeupdate_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.5" "safeupdate_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "safeupdate_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.5" "safeupdate_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "safeupdate_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "safeupdate_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-safeupdate : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-safeupdate : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-safeupdate : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-safeupdate : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-safeupdate : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-safeupdate : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-safeupdate : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.5" "postgresql-18-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-17-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-16-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-15-pg-safeupdate : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5" "postgresql-14-pg-safeupdate : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-safeupdate : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `safeupdate_18` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.4 KiB | [safeupdate_18-1.5-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/safeupdate_18-1.5-2PGDG.rhel8.x86_64.rpm) |
| `safeupdate_18` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.4 KiB | [safeupdate_18-1.5-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/safeupdate_18-1.5-2PGDG.rhel8.aarch64.rpm) |
| `safeupdate_18` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.8 KiB | [safeupdate_18-1.5-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/safeupdate_18-1.5-2PGDG.rhel9.x86_64.rpm) |
| `safeupdate_18` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.6 KiB | [safeupdate_18-1.5-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/safeupdate_18-1.5-2PGDG.rhel9.aarch64.rpm) |
| `safeupdate_18` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.2 KiB | [safeupdate_18-1.5-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/safeupdate_18-1.5-2PGDG.rhel10.x86_64.rpm) |
| `safeupdate_18` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [safeupdate_18-1.5-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/safeupdate_18-1.5-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-safeupdate` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-18-pg-safeupdate_1.5-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-18-pg-safeupdate_1.5-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-safeupdate` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.8 KiB | [postgresql-18-pg-safeupdate_1.5-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-18-pg-safeupdate_1.5-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-safeupdate` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-18-pg-safeupdate_1.5-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-safeupdate/postgresql-18-pg-safeupdate_1.5-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-safeupdate` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.8 KiB | [postgresql-18-pg-safeupdate_1.5-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-safeupdate/postgresql-18-pg-safeupdate_1.5-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-safeupdate` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 8.8 KiB | [postgresql-18-pg-safeupdate_1.5-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-18-pg-safeupdate_1.5-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-safeupdate` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.9 KiB | [postgresql-18-pg-safeupdate_1.5-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-18-pg-safeupdate_1.5-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-safeupdate` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.0 KiB | [postgresql-18-pg-safeupdate_1.5-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-18-pg-safeupdate_1.5-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-safeupdate` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.0 KiB | [postgresql-18-pg-safeupdate_1.5-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-18-pg-safeupdate_1.5-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `safeupdate_17` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.3 KiB | [safeupdate_17-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/safeupdate_17-1.5-1PGDG.rhel8.x86_64.rpm) |
| `safeupdate_17` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.3 KiB | [safeupdate_17-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/safeupdate_17-1.5-1PGDG.rhel8.aarch64.rpm) |
| `safeupdate_17` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [safeupdate_17-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/safeupdate_17-1.5-1PGDG.rhel9.x86_64.rpm) |
| `safeupdate_17` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [safeupdate_17-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/safeupdate_17-1.5-1PGDG.rhel9.aarch64.rpm) |
| `safeupdate_17` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.2 KiB | [safeupdate_17-1.5-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/safeupdate_17-1.5-2PGDG.rhel10.x86_64.rpm) |
| `safeupdate_17` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [safeupdate_17-1.5-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/safeupdate_17-1.5-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-safeupdate` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-17-pg-safeupdate_1.5-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-safeupdate` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-17-pg-safeupdate_1.5-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-safeupdate` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-17-pg-safeupdate_1.5-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-safeupdate` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.8 KiB | [postgresql-17-pg-safeupdate_1.5-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-safeupdate` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.4 KiB | [postgresql-17-pg-safeupdate_1.5-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-safeupdate` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.5 KiB | [postgresql-17-pg-safeupdate_1.5-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-safeupdate` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.9 KiB | [postgresql-17-pg-safeupdate_1.5-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-safeupdate` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.0 KiB | [postgresql-17-pg-safeupdate_1.5-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-17-pg-safeupdate_1.5-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `safeupdate_16` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.3 KiB | [safeupdate_16-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/safeupdate_16-1.5-1PGDG.rhel8.x86_64.rpm) |
| `safeupdate_16` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.4 KiB | [safeupdate_16-1.4.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/safeupdate_16-1.4.2-2PGDG.rhel8.x86_64.rpm) |
| `safeupdate_16` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.3 KiB | [safeupdate_16-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/safeupdate_16-1.5-1PGDG.rhel8.aarch64.rpm) |
| `safeupdate_16` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.4 KiB | [safeupdate_16-1.4.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/safeupdate_16-1.4.2-2PGDG.rhel8.aarch64.rpm) |
| `safeupdate_16` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [safeupdate_16-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/safeupdate_16-1.5-1PGDG.rhel9.x86_64.rpm) |
| `safeupdate_16` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [safeupdate_16-1.4.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/safeupdate_16-1.4.2-2PGDG.rhel9.x86_64.rpm) |
| `safeupdate_16` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [safeupdate_16-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/safeupdate_16-1.5-1PGDG.rhel9.aarch64.rpm) |
| `safeupdate_16` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.6 KiB | [safeupdate_16-1.4.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/safeupdate_16-1.4.2-2PGDG.rhel9.aarch64.rpm) |
| `safeupdate_16` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.2 KiB | [safeupdate_16-1.5-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/safeupdate_16-1.5-2PGDG.rhel10.x86_64.rpm) |
| `safeupdate_16` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.2 KiB | [safeupdate_16-1.5-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/safeupdate_16-1.5-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-safeupdate` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-16-pg-safeupdate_1.5-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-safeupdate` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-16-pg-safeupdate_1.5-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-safeupdate` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-16-pg-safeupdate_1.5-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-safeupdate` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.8 KiB | [postgresql-16-pg-safeupdate_1.5-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-safeupdate` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.4 KiB | [postgresql-16-pg-safeupdate_1.5-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-safeupdate` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.5 KiB | [postgresql-16-pg-safeupdate_1.5-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-safeupdate` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.9 KiB | [postgresql-16-pg-safeupdate_1.5-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-safeupdate` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.9 KiB | [postgresql-16-pg-safeupdate_1.5-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-16-pg-safeupdate_1.5-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `safeupdate_15` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.3 KiB | [safeupdate_15-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/safeupdate_15-1.5-1PGDG.rhel8.x86_64.rpm) |
| `safeupdate_15` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.8 KiB | [safeupdate_15-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/safeupdate_15-1.4-1.rhel8.x86_64.rpm) |
| `safeupdate_15` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.3 KiB | [safeupdate_15-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/safeupdate_15-1.5-1PGDG.rhel8.aarch64.rpm) |
| `safeupdate_15` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.3 KiB | [safeupdate_15-1.4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/safeupdate_15-1.4.2-1.rhel8.aarch64.rpm) |
| `safeupdate_15` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.8 KiB | [safeupdate_15-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/safeupdate_15-1.4-1.rhel8.aarch64.rpm) |
| `safeupdate_15` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [safeupdate_15-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/safeupdate_15-1.5-1PGDG.rhel9.x86_64.rpm) |
| `safeupdate_15` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [safeupdate_15-1.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/safeupdate_15-1.4.2-1.rhel9.x86_64.rpm) |
| `safeupdate_15` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.9 KiB | [safeupdate_15-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/safeupdate_15-1.4-1.rhel9.x86_64.rpm) |
| `safeupdate_15` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [safeupdate_15-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/safeupdate_15-1.5-1PGDG.rhel9.aarch64.rpm) |
| `safeupdate_15` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.6 KiB | [safeupdate_15-1.4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/safeupdate_15-1.4.2-1.rhel9.aarch64.rpm) |
| `safeupdate_15` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.7 KiB | [safeupdate_15-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/safeupdate_15-1.4-1.rhel9.aarch64.rpm) |
| `safeupdate_15` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.3 KiB | [safeupdate_15-1.5-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/safeupdate_15-1.5-2PGDG.rhel10.x86_64.rpm) |
| `safeupdate_15` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.3 KiB | [safeupdate_15-1.5-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/safeupdate_15-1.5-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-safeupdate` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-15-pg-safeupdate_1.5-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-safeupdate` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-15-pg-safeupdate_1.5-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-safeupdate` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-15-pg-safeupdate_1.5-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-safeupdate` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.8 KiB | [postgresql-15-pg-safeupdate_1.5-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-safeupdate` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.4 KiB | [postgresql-15-pg-safeupdate_1.5-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-safeupdate` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.5 KiB | [postgresql-15-pg-safeupdate_1.5-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-safeupdate` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.9 KiB | [postgresql-15-pg-safeupdate_1.5-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-safeupdate` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.9 KiB | [postgresql-15-pg-safeupdate_1.5-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-15-pg-safeupdate_1.5-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `safeupdate_14` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.3 KiB | [safeupdate_14-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/safeupdate_14-1.5-1PGDG.rhel8.x86_64.rpm) |
| `safeupdate_14` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.3 KiB | [safeupdate_14-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/safeupdate_14-1.4-1.rhel8.x86_64.rpm) |
| `safeupdate_14` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.3 KiB | [safeupdate_14-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/safeupdate_14-1.5-1PGDG.rhel8.aarch64.rpm) |
| `safeupdate_14` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.2 KiB | [safeupdate_14-1.4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/safeupdate_14-1.4.2-1.rhel8.aarch64.rpm) |
| `safeupdate_14` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.7 KiB | [safeupdate_14-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/safeupdate_14-1.4-1.rhel8.aarch64.rpm) |
| `safeupdate_14` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.9 KiB | [safeupdate_14-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/safeupdate_14-1.5-1PGDG.rhel9.x86_64.rpm) |
| `safeupdate_14` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.8 KiB | [safeupdate_14-1.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/safeupdate_14-1.4.2-1.rhel9.x86_64.rpm) |
| `safeupdate_14` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [safeupdate_14-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/safeupdate_14-1.5-1PGDG.rhel9.aarch64.rpm) |
| `safeupdate_14` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.6 KiB | [safeupdate_14-1.4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/safeupdate_14-1.4.2-1.rhel9.aarch64.rpm) |
| `safeupdate_14` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.6 KiB | [safeupdate_14-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/safeupdate_14-1.4-1.rhel9.aarch64.rpm) |
| `safeupdate_14` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.3 KiB | [safeupdate_14-1.5-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/safeupdate_14-1.5-2PGDG.rhel10.x86_64.rpm) |
| `safeupdate_14` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.3 KiB | [safeupdate_14-1.5-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/safeupdate_14-1.5-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-safeupdate` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-14-pg-safeupdate_1.5-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-safeupdate` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-14-pg-safeupdate_1.5-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-safeupdate` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-14-pg-safeupdate_1.5-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-safeupdate` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.8 KiB | [postgresql-14-pg-safeupdate_1.5-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-safeupdate` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.3 KiB | [postgresql-14-pg-safeupdate_1.5-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-safeupdate` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.5 KiB | [postgresql-14-pg-safeupdate_1.5-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-safeupdate` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 8.9 KiB | [postgresql-14-pg-safeupdate_1.5-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-safeupdate` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.9 KiB | [postgresql-14-pg-safeupdate_1.5-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-safeupdate/postgresql-14-pg-safeupdate_1.5-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `safeupdate_13` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.9 KiB | [safeupdate_13-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/safeupdate_13-1.4-1.rhel8.x86_64.rpm) |
| `safeupdate_13` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.7 KiB | [safeupdate_13-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/safeupdate_13-1.3-1.rhel8.x86_64.rpm) |
| `safeupdate_13` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.2 KiB | [safeupdate_13-1.4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/safeupdate_13-1.4.2-1.rhel8.aarch64.rpm) |
| `safeupdate_13` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.6 KiB | [safeupdate_13-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/safeupdate_13-1.4-1.rhel8.aarch64.rpm) |
| `safeupdate_13` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.8 KiB | [safeupdate_13-1.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/safeupdate_13-1.4.2-1.rhel9.x86_64.rpm) |
| `safeupdate_13` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.6 KiB | [safeupdate_13-1.4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/safeupdate_13-1.4.2-1.rhel9.aarch64.rpm) |
| `safeupdate_13` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.4 KiB | [safeupdate_13-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/safeupdate_13-1.4-1.rhel9.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/eradman/pg-safeupdate" title="Repository" icon="github" subtitle="github.com/eradman/pg-safeupdate" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-safeupdate-1.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg safeupdate;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install safeupdate;		# install via package name, for the active PG version

pig install safeupdate -v 18;   # install for PG 18
pig install safeupdate -v 17;   # install for PG 17
pig install safeupdate -v 16;   # install for PG 16
pig install safeupdate -v 15;   # install for PG 15
pig install safeupdate -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'safeupdate';
```


This extension does not need `CREATE EXTENSION` DDL command


