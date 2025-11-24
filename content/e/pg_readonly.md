---
title: "pg_readonly"
linkTitle: "pg_readonly"
description: "cluster database read only"
weight: 5120
categories: ["ADMIN"]
width: full
---

[**pg_readonly**](https://github.com/pierreforstmann/pg_readonly) : cluster database read only


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5120** | {{< badge content="pg_readonly" link="https://github.com/pierreforstmann/pg_readonly" >}} | {{< ext "pg_readonly" >}} | `1.0.3` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_permissions" >}} {{< ext "pg_upless" >}} {{< ext "safeupdate" >}} {{< ext "set_user" >}} {{< ext "pgaudit" >}} {{< ext "noset" >}} {{< ext "sepgsql" >}} {{< ext "login_hook" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_readonly` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.3` | {{< bg "18" "pg_readonly_18*" "green" >}} {{< bg "17" "pg_readonly_17*" "green" >}} {{< bg "16" "pg_readonly_16*" "green" >}} {{< bg "15" "pg_readonly_15*" "green" >}} {{< bg "14" "pg_readonly_14*" "green" >}} {{< bg "13" "pg_readonly_13*" "green" >}} | `pg_readonly_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.3` | {{< bg "18" "postgresql-18-pg-readonly" "green" >}} {{< bg "17" "postgresql-17-pg-readonly" "green" >}} {{< bg "16" "postgresql-16-pg-readonly" "green" >}} {{< bg "15" "postgresql-15-pg-readonly" "green" >}} {{< bg "14" "postgresql-14-pg-readonly" "green" >}} {{< bg "13" "postgresql-13-pg-readonly" "green" >}} | `postgresql-$v-pg-readonly` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_13 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.3" "pg_readonly_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-pg-readonly : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-pg-readonly : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-pg-readonly : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-pg-readonly : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-pg-readonly : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-pg-readonly : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-pg-readonly : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-pg-readonly : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-13-pg-readonly : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readonly_18` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.6 KiB | [pg_readonly_18-1.0.3-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_readonly_18-1.0.3-5PGDG.rhel8.x86_64.rpm) |
| `pg_readonly_18` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.5 KiB | [pg_readonly_18-1.0.3-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_readonly_18-1.0.3-5PGDG.rhel8.aarch64.rpm) |
| `pg_readonly_18` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.5 KiB | [pg_readonly_18-1.0.3-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_readonly_18-1.0.3-5PGDG.rhel9.x86_64.rpm) |
| `pg_readonly_18` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.0 KiB | [pg_readonly_18-1.0.3-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_readonly_18-1.0.3-5PGDG.rhel9.aarch64.rpm) |
| `pg_readonly_18` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.8 KiB | [pg_readonly_18-1.0.3-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_readonly_18-1.0.3-5PGDG.rhel10.x86_64.rpm) |
| `pg_readonly_18` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.7 KiB | [pg_readonly_18-1.0.3-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10.0-aarch64/pg_readonly_18-1.0.3-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-readonly` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.4 KiB | [postgresql-18-pg-readonly_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-18-pg-readonly_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-readonly` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.1 KiB | [postgresql-18-pg-readonly_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-18-pg-readonly_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-readonly` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.3 KiB | [postgresql-18-pg-readonly_1.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readonly/postgresql-18-pg-readonly_1.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-readonly` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.2 KiB | [postgresql-18-pg-readonly_1.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readonly/postgresql-18-pg-readonly_1.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-readonly` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.9 KiB | [postgresql-18-pg-readonly_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-18-pg-readonly_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-readonly` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.7 KiB | [postgresql-18-pg-readonly_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-18-pg-readonly_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-readonly` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.0 KiB | [postgresql-18-pg-readonly_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-18-pg-readonly_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-readonly` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.8 KiB | [postgresql-18-pg-readonly_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-18-pg-readonly_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readonly_17` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.6 KiB | [pg_readonly_17-1.0.3-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_readonly_17-1.0.3-4PGDG.rhel8.x86_64.rpm) |
| `pg_readonly_17` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.4 KiB | [pg_readonly_17-1.0.3-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_readonly_17-1.0.3-4PGDG.rhel8.aarch64.rpm) |
| `pg_readonly_17` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.7 KiB | [pg_readonly_17-1.0.3-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_readonly_17-1.0.3-4PGDG.rhel9.x86_64.rpm) |
| `pg_readonly_17` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.4 KiB | [pg_readonly_17-1.0.3-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_readonly_17-1.0.3-4PGDG.rhel9.aarch64.rpm) |
| `pg_readonly_17` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.0 KiB | [pg_readonly_17-1.0.3-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_readonly_17-1.0.3-5PGDG.rhel10.x86_64.rpm) |
| `pg_readonly_17` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.9 KiB | [pg_readonly_17-1.0.3-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10.0-aarch64/pg_readonly_17-1.0.3-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-readonly` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.4 KiB | [postgresql-17-pg-readonly_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-readonly` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.2 KiB | [postgresql-17-pg-readonly_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-readonly` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.4 KiB | [postgresql-17-pg-readonly_1.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-readonly` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.2 KiB | [postgresql-17-pg-readonly_1.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-readonly` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.1 KiB | [postgresql-17-pg-readonly_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-readonly` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.0 KiB | [postgresql-17-pg-readonly_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-readonly` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.2 KiB | [postgresql-17-pg-readonly_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-readonly` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.9 KiB | [postgresql-17-pg-readonly_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readonly_16` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.4 KiB | [pg_readonly_16-1.0.3-2.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_readonly_16-1.0.3-2.rhel8.1.x86_64.rpm) |
| `pg_readonly_16` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.2 KiB | [pg_readonly_16-1.0.3-2.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_readonly_16-1.0.3-2.rhel8.1.aarch64.rpm) |
| `pg_readonly_16` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.5 KiB | [pg_readonly_16-1.0.3-2.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_readonly_16-1.0.3-2.rhel9.1.x86_64.rpm) |
| `pg_readonly_16` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.1 KiB | [pg_readonly_16-1.0.3-2.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_readonly_16-1.0.3-2.rhel9.1.aarch64.rpm) |
| `pg_readonly_16` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.0 KiB | [pg_readonly_16-1.0.3-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_readonly_16-1.0.3-5PGDG.rhel10.x86_64.rpm) |
| `pg_readonly_16` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.9 KiB | [pg_readonly_16-1.0.3-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10.0-aarch64/pg_readonly_16-1.0.3-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-readonly` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.4 KiB | [postgresql-16-pg-readonly_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-readonly` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.1 KiB | [postgresql-16-pg-readonly_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-readonly` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.4 KiB | [postgresql-16-pg-readonly_1.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-readonly` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.2 KiB | [postgresql-16-pg-readonly_1.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-readonly` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.0 KiB | [postgresql-16-pg-readonly_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-readonly` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.9 KiB | [postgresql-16-pg-readonly_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-readonly` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.2 KiB | [postgresql-16-pg-readonly_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-readonly` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.9 KiB | [postgresql-16-pg-readonly_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readonly_15` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.7 KiB | [pg_readonly_15-1.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_readonly_15-1.0.3-1.rhel8.x86_64.rpm) |
| `pg_readonly_15` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.2 KiB | [pg_readonly_15-1.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_readonly_15-1.0.1-1.rhel8.x86_64.rpm) |
| `pg_readonly_15` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.4 KiB | [pg_readonly_15-1.0.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_readonly_15-1.0.3-1.rhel8.aarch64.rpm) |
| `pg_readonly_15` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.3 KiB | [pg_readonly_15-1.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_readonly_15-1.0.3-1.rhel9.x86_64.rpm) |
| `pg_readonly_15` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [pg_readonly_15-1.0.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_readonly_15-1.0.1-1.rhel9.x86_64.rpm) |
| `pg_readonly_15` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.0 KiB | [pg_readonly_15-1.0.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_readonly_15-1.0.3-1.rhel9.aarch64.rpm) |
| `pg_readonly_15` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.0 KiB | [pg_readonly_15-1.0.3-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_readonly_15-1.0.3-5PGDG.rhel10.x86_64.rpm) |
| `pg_readonly_15` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.9 KiB | [pg_readonly_15-1.0.3-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10.0-aarch64/pg_readonly_15-1.0.3-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-readonly` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.4 KiB | [postgresql-15-pg-readonly_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-readonly` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.2 KiB | [postgresql-15-pg-readonly_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-readonly` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.3 KiB | [postgresql-15-pg-readonly_1.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-readonly` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.2 KiB | [postgresql-15-pg-readonly_1.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-readonly` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.0 KiB | [postgresql-15-pg-readonly_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-readonly` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.9 KiB | [postgresql-15-pg-readonly_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-readonly` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.2 KiB | [postgresql-15-pg-readonly_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-readonly` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.9 KiB | [postgresql-15-pg-readonly_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readonly_14` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.5 KiB | [pg_readonly_14-1.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_readonly_14-1.0.3-1.rhel8.x86_64.rpm) |
| `pg_readonly_14` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.7 KiB | [pg_readonly_14-1.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_readonly_14-1.0.1-1.rhel8.x86_64.rpm) |
| `pg_readonly_14` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.1 KiB | [pg_readonly_14-1.0.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_readonly_14-1.0.3-1.rhel8.aarch64.rpm) |
| `pg_readonly_14` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.1 KiB | [pg_readonly_14-1.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_readonly_14-1.0.3-1.rhel9.x86_64.rpm) |
| `pg_readonly_14` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [pg_readonly_14-1.0.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_readonly_14-1.0.1-1.rhel9.x86_64.rpm) |
| `pg_readonly_14` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.7 KiB | [pg_readonly_14-1.0.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_readonly_14-1.0.3-1.rhel9.aarch64.rpm) |
| `pg_readonly_14` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.9 KiB | [pg_readonly_14-1.0.3-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_readonly_14-1.0.3-5PGDG.rhel10.x86_64.rpm) |
| `pg_readonly_14` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.8 KiB | [pg_readonly_14-1.0.3-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10.0-aarch64/pg_readonly_14-1.0.3-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-readonly` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.1 KiB | [postgresql-14-pg-readonly_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-readonly` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.9 KiB | [postgresql-14-pg-readonly_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-readonly` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.1 KiB | [postgresql-14-pg-readonly_1.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-readonly` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.0 KiB | [postgresql-14-pg-readonly_1.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-readonly` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.8 KiB | [postgresql-14-pg-readonly_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-readonly` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.7 KiB | [postgresql-14-pg-readonly_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-readonly` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.9 KiB | [postgresql-14-pg-readonly_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-readonly` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.7 KiB | [postgresql-14-pg-readonly_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readonly_13` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.2 KiB | [pg_readonly_13-1.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_readonly_13-1.0.3-1.rhel8.x86_64.rpm) |
| `pg_readonly_13` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.3 KiB | [pg_readonly_13-1.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_readonly_13-1.0.1-1.rhel8.x86_64.rpm) |
| `pg_readonly_13` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.9 KiB | [pg_readonly_13-1.0.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_readonly_13-1.0.3-1.rhel8.aarch64.rpm) |
| `pg_readonly_13` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.7 KiB | [pg_readonly_13-1.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_readonly_13-1.0.3-1.rhel9.x86_64.rpm) |
| `pg_readonly_13` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.6 KiB | [pg_readonly_13-1.0.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_readonly_13-1.0.1-1.rhel9.x86_64.rpm) |
| `pg_readonly_13` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.4 KiB | [pg_readonly_13-1.0.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_readonly_13-1.0.3-1.rhel9.aarch64.rpm) |
| `pg_readonly_13` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.8 KiB | [pg_readonly_13-1.0.3-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_readonly_13-1.0.3-5PGDG.rhel10.x86_64.rpm) |
| `pg_readonly_13` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.8 KiB | [pg_readonly_13-1.0.3-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10.0-aarch64/pg_readonly_13-1.0.3-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-readonly` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.9 KiB | [postgresql-13-pg-readonly_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-readonly` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.8 KiB | [postgresql-13-pg-readonly_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-readonly` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.9 KiB | [postgresql-13-pg-readonly_1.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-readonly` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.9 KiB | [postgresql-13-pg-readonly_1.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-readonly` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.5 KiB | [postgresql-13-pg-readonly_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-readonly` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.4 KiB | [postgresql-13-pg-readonly_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-readonly` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.7 KiB | [postgresql-13-pg-readonly_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-readonly` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.6 KiB | [postgresql-13-pg-readonly_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pierreforstmann/pg_readonly" title="Repository" icon="github" subtitle="github.com/pierreforstmann/pg_readonly" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_readonly-1.0.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_readonly;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_readonly;		# install via package name, for the active PG version

pig install pg_readonly -v 18;   # install for PG 18
pig install pg_readonly -v 17;   # install for PG 17
pig install pg_readonly -v 16;   # install for PG 16
pig install pg_readonly -v 15;   # install for PG 15
pig install pg_readonly -v 14;   # install for PG 14
pig install pg_readonly -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_readonly;
```
