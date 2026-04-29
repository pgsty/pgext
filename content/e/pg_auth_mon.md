---
title: "pg_auth_mon"
linkTitle: "pg_auth_mon"
description: "monitor connection attempts per user"
weight: 7150
categories: ["SEC"]
width: full
---

[**pg_auth_mon**](https://github.com/RafiaSabih/pg_auth_mon) : monitor connection attempts per user


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7150** | {{< badge content="pg_auth_mon" link="https://github.com/RafiaSabih/pg_auth_mon" >}} | {{< ext "pg_auth_mon" >}} | `3.0` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "passwordcheck_cracklib" >}} {{< ext "pgaudit" >}} {{< ext "pgauditlogtofile" >}} {{< ext "login_hook" >}} {{< ext "auth_delay" >}} {{< ext "credcheck" >}} {{< ext "logerrors" >}} {{< ext "set_user" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_auth_mon` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.0` | {{< bg "18" "pg_auth_mon_18" "green" >}} {{< bg "17" "pg_auth_mon_17" "green" >}} {{< bg "16" "pg_auth_mon_16" "green" >}} {{< bg "15" "pg_auth_mon_15" "green" >}} {{< bg "14" "pg_auth_mon_14" "green" >}} | `pg_auth_mon_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.0` | {{< bg "18" "postgresql-18-pg-auth-mon" "green" >}} {{< bg "17" "postgresql-17-pg-auth-mon" "green" >}} {{< bg "16" "postgresql-16-pg-auth-mon" "green" >}} {{< bg "15" "postgresql-15-pg-auth-mon" "green" >}} {{< bg "14" "postgresql-14-pg-auth-mon" "green" >}} | `postgresql-$v-pg-auth-mon` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0" "pg_auth_mon_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0" "pg_auth_mon_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0" "pg_auth_mon_14 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0" "pg_auth_mon_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0" "pg_auth_mon_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0" "pg_auth_mon_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0" "pg_auth_mon_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0" "pg_auth_mon_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0" "pg_auth_mon_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0" "pg_auth_mon_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0" "pg_auth_mon_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0" "pg_auth_mon_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.0" "pg_auth_mon_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.0" "postgresql-18-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-17-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-16-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-15-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-14-pg-auth-mon : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.0" "postgresql-18-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-17-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-16-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-15-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-14-pg-auth-mon : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.0" "postgresql-18-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-17-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-16-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-15-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-14-pg-auth-mon : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.0" "postgresql-18-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-17-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-16-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-15-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-14-pg-auth-mon : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.0" "postgresql-18-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-17-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-16-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-15-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-14-pg-auth-mon : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.0" "postgresql-18-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-17-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-16-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-15-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-14-pg-auth-mon : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.0" "postgresql-18-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-17-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-16-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-15-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-14-pg-auth-mon : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.0" "postgresql-18-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-17-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-16-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-15-pg-auth-mon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0" "postgresql-14-pg-auth-mon : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-auth-mon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-auth-mon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-auth-mon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-auth-mon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-auth-mon : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-auth-mon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-auth-mon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-auth-mon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-auth-mon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-auth-mon : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auth_mon_18` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.9 KiB | [pg_auth_mon_18-3.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_auth_mon_18-3.0-3PGDG.rhel8.x86_64.rpm) |
| `pg_auth_mon_18` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.0 KiB | [pg_auth_mon_18-3.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_auth_mon_18-3.0-3PGDG.rhel8.aarch64.rpm) |
| `pg_auth_mon_18` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.9 KiB | [pg_auth_mon_18-3.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_auth_mon_18-3.0-3PGDG.rhel9.x86_64.rpm) |
| `pg_auth_mon_18` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.7 KiB | [pg_auth_mon_18-3.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_auth_mon_18-3.0-3PGDG.rhel9.aarch64.rpm) |
| `pg_auth_mon_18` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.2 KiB | [pg_auth_mon_18-3.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_auth_mon_18-3.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_auth_mon_18` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.4 KiB | [pg_auth_mon_18-3.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_auth_mon_18-3.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-auth-mon` | `3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.5 KiB | [postgresql-18-pg-auth-mon_3.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auth-mon/postgresql-18-pg-auth-mon_3.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-auth-mon` | `3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.7 KiB | [postgresql-18-pg-auth-mon_3.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auth-mon/postgresql-18-pg-auth-mon_3.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-auth-mon` | `3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.5 KiB | [postgresql-18-pg-auth-mon_3.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-auth-mon/postgresql-18-pg-auth-mon_3.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-auth-mon` | `3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.7 KiB | [postgresql-18-pg-auth-mon_3.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-auth-mon/postgresql-18-pg-auth-mon_3.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-auth-mon` | `3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.6 KiB | [postgresql-18-pg-auth-mon_3.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auth-mon/postgresql-18-pg-auth-mon_3.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-auth-mon` | `3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.5 KiB | [postgresql-18-pg-auth-mon_3.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auth-mon/postgresql-18-pg-auth-mon_3.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-auth-mon` | `3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.3 KiB | [postgresql-18-pg-auth-mon_3.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auth-mon/postgresql-18-pg-auth-mon_3.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-auth-mon` | `3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.4 KiB | [postgresql-18-pg-auth-mon_3.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auth-mon/postgresql-18-pg-auth-mon_3.0-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auth_mon_17` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.6 KiB | [pg_auth_mon_17-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_auth_mon_17-3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_auth_mon_17` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.7 KiB | [pg_auth_mon_17-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_auth_mon_17-3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_auth_mon_17` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.8 KiB | [pg_auth_mon_17-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_auth_mon_17-3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_auth_mon_17` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.7 KiB | [pg_auth_mon_17-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_auth_mon_17-3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_auth_mon_17` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.3 KiB | [pg_auth_mon_17-3.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_auth_mon_17-3.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_auth_mon_17` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.4 KiB | [pg_auth_mon_17-3.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_auth_mon_17-3.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-auth-mon` | `3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-17-pg-auth-mon_3.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auth-mon/postgresql-17-pg-auth-mon_3.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-auth-mon` | `3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.7 KiB | [postgresql-17-pg-auth-mon_3.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auth-mon/postgresql-17-pg-auth-mon_3.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-auth-mon` | `3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.4 KiB | [postgresql-17-pg-auth-mon_3.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-auth-mon/postgresql-17-pg-auth-mon_3.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-auth-mon` | `3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.7 KiB | [postgresql-17-pg-auth-mon_3.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-auth-mon/postgresql-17-pg-auth-mon_3.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-auth-mon` | `3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.5 KiB | [postgresql-17-pg-auth-mon_3.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auth-mon/postgresql-17-pg-auth-mon_3.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-auth-mon` | `3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.5 KiB | [postgresql-17-pg-auth-mon_3.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auth-mon/postgresql-17-pg-auth-mon_3.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-auth-mon` | `3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.3 KiB | [postgresql-17-pg-auth-mon_3.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auth-mon/postgresql-17-pg-auth-mon_3.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-auth-mon` | `3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.4 KiB | [postgresql-17-pg-auth-mon_3.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auth-mon/postgresql-17-pg-auth-mon_3.0-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auth_mon_16` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.3 KiB | [pg_auth_mon_16-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_auth_mon_16-2.0-1.rhel8.x86_64.rpm) |
| `pg_auth_mon_16` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.4 KiB | [pg_auth_mon_16-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_auth_mon_16-2.0-1.rhel8.aarch64.rpm) |
| `pg_auth_mon_16` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.3 KiB | [pg_auth_mon_16-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_auth_mon_16-2.0-1.rhel9.x86_64.rpm) |
| `pg_auth_mon_16` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.1 KiB | [pg_auth_mon_16-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_auth_mon_16-2.0-1.rhel9.aarch64.rpm) |
| `pg_auth_mon_16` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.3 KiB | [pg_auth_mon_16-3.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_auth_mon_16-3.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_auth_mon_16` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.4 KiB | [pg_auth_mon_16-3.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_auth_mon_16-3.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-auth-mon` | `3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-16-pg-auth-mon_3.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auth-mon/postgresql-16-pg-auth-mon_3.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-auth-mon` | `3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.7 KiB | [postgresql-16-pg-auth-mon_3.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auth-mon/postgresql-16-pg-auth-mon_3.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-auth-mon` | `3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.5 KiB | [postgresql-16-pg-auth-mon_3.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-auth-mon/postgresql-16-pg-auth-mon_3.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-auth-mon` | `3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.7 KiB | [postgresql-16-pg-auth-mon_3.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-auth-mon/postgresql-16-pg-auth-mon_3.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-auth-mon` | `3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.5 KiB | [postgresql-16-pg-auth-mon_3.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auth-mon/postgresql-16-pg-auth-mon_3.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-auth-mon` | `3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.4 KiB | [postgresql-16-pg-auth-mon_3.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auth-mon/postgresql-16-pg-auth-mon_3.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-auth-mon` | `3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.3 KiB | [postgresql-16-pg-auth-mon_3.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auth-mon/postgresql-16-pg-auth-mon_3.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-auth-mon` | `3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.4 KiB | [postgresql-16-pg-auth-mon_3.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auth-mon/postgresql-16-pg-auth-mon_3.0-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auth_mon_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.3 KiB | [pg_auth_mon_15-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_auth_mon_15-2.0-1.rhel8.x86_64.rpm) |
| `pg_auth_mon_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.5 KiB | [pg_auth_mon_15-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_auth_mon_15-1.0-1.rhel8.x86_64.rpm) |
| `pg_auth_mon_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.4 KiB | [pg_auth_mon_15-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_auth_mon_15-2.0-1.rhel8.aarch64.rpm) |
| `pg_auth_mon_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.5 KiB | [pg_auth_mon_15-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_auth_mon_15-1.0-1.rhel8.aarch64.rpm) |
| `pg_auth_mon_15` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.3 KiB | [pg_auth_mon_15-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_auth_mon_15-2.0-1.rhel9.x86_64.rpm) |
| `pg_auth_mon_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.9 KiB | [pg_auth_mon_15-1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_auth_mon_15-1.0-1.rhel9.x86_64.rpm) |
| `pg_auth_mon_15` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.2 KiB | [pg_auth_mon_15-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_auth_mon_15-2.0-1.rhel9.aarch64.rpm) |
| `pg_auth_mon_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.7 KiB | [pg_auth_mon_15-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_auth_mon_15-1.0-1.rhel9.aarch64.rpm) |
| `pg_auth_mon_15` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.3 KiB | [pg_auth_mon_15-3.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_auth_mon_15-3.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_auth_mon_15` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.4 KiB | [pg_auth_mon_15-3.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_auth_mon_15-3.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-auth-mon` | `3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.5 KiB | [postgresql-15-pg-auth-mon_3.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auth-mon/postgresql-15-pg-auth-mon_3.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-auth-mon` | `3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.7 KiB | [postgresql-15-pg-auth-mon_3.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auth-mon/postgresql-15-pg-auth-mon_3.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-auth-mon` | `3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.5 KiB | [postgresql-15-pg-auth-mon_3.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-auth-mon/postgresql-15-pg-auth-mon_3.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-auth-mon` | `3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.8 KiB | [postgresql-15-pg-auth-mon_3.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-auth-mon/postgresql-15-pg-auth-mon_3.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-auth-mon` | `3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.5 KiB | [postgresql-15-pg-auth-mon_3.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auth-mon/postgresql-15-pg-auth-mon_3.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-auth-mon` | `3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.4 KiB | [postgresql-15-pg-auth-mon_3.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auth-mon/postgresql-15-pg-auth-mon_3.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-auth-mon` | `3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.4 KiB | [postgresql-15-pg-auth-mon_3.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auth-mon/postgresql-15-pg-auth-mon_3.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-auth-mon` | `3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.4 KiB | [postgresql-15-pg-auth-mon_3.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auth-mon/postgresql-15-pg-auth-mon_3.0-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auth_mon_14` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.2 KiB | [pg_auth_mon_14-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_auth_mon_14-2.0-1.rhel8.x86_64.rpm) |
| `pg_auth_mon_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.9 KiB | [pg_auth_mon_14-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_auth_mon_14-1.0-1.rhel8.x86_64.rpm) |
| `pg_auth_mon_14` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.3 KiB | [pg_auth_mon_14-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_auth_mon_14-2.0-1.rhel8.aarch64.rpm) |
| `pg_auth_mon_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.5 KiB | [pg_auth_mon_14-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_auth_mon_14-1.0-1.rhel8.aarch64.rpm) |
| `pg_auth_mon_14` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.3 KiB | [pg_auth_mon_14-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_auth_mon_14-2.0-1.rhel9.x86_64.rpm) |
| `pg_auth_mon_14` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.1 KiB | [pg_auth_mon_14-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_auth_mon_14-2.0-1.rhel9.aarch64.rpm) |
| `pg_auth_mon_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.7 KiB | [pg_auth_mon_14-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_auth_mon_14-1.0-1.rhel9.aarch64.rpm) |
| `pg_auth_mon_14` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.2 KiB | [pg_auth_mon_14-3.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_auth_mon_14-3.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_auth_mon_14` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.4 KiB | [pg_auth_mon_14-3.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_auth_mon_14-3.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-auth-mon` | `3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.2 KiB | [postgresql-14-pg-auth-mon_3.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auth-mon/postgresql-14-pg-auth-mon_3.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-auth-mon` | `3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.7 KiB | [postgresql-14-pg-auth-mon_3.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auth-mon/postgresql-14-pg-auth-mon_3.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-auth-mon` | `3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.3 KiB | [postgresql-14-pg-auth-mon_3.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-auth-mon/postgresql-14-pg-auth-mon_3.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-auth-mon` | `3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.7 KiB | [postgresql-14-pg-auth-mon_3.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-auth-mon/postgresql-14-pg-auth-mon_3.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-auth-mon` | `3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.3 KiB | [postgresql-14-pg-auth-mon_3.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auth-mon/postgresql-14-pg-auth-mon_3.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-auth-mon` | `3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.1 KiB | [postgresql-14-pg-auth-mon_3.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auth-mon/postgresql-14-pg-auth-mon_3.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-auth-mon` | `3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.2 KiB | [postgresql-14-pg-auth-mon_3.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auth-mon/postgresql-14-pg-auth-mon_3.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-auth-mon` | `3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.3 KiB | [postgresql-14-pg-auth-mon_3.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auth-mon/postgresql-14-pg-auth-mon_3.0-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/RafiaSabih/pg_auth_mon" title="Repository" icon="github" subtitle="github.com/RafiaSabih/pg_auth_mon" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_auth_mon-3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_auth_mon;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_auth_mon;		# install via package name, for the active PG version

pig install pg_auth_mon -v 18;   # install for PG 18
pig install pg_auth_mon -v 17;   # install for PG 17
pig install pg_auth_mon -v 16;   # install for PG 16
pig install pg_auth_mon -v 15;   # install for PG 15
pig install pg_auth_mon -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_auth_mon';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_auth_mon;
```




## Usage

> [pg_auth_mon: Monitor authentication attempts](https://github.com/RafiaSabih/pg_auth_mon)

`pg_auth_mon` stores authentication attempt statistics per user, tracking successful and failed logins with timestamps.

```sql
CREATE EXTENSION pg_auth_mon;
```

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'pg_auth_mon'
```

Optional GUC:

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_auth_mon.log_period` | `0` | Dump stats to Postgres log every N minutes (0=off) |

### Querying Authentication Stats

```sql
SELECT * FROM pg_auth_mon;
```

The `pg_auth_mon` view provides per-user information:

| Column | Description |
|--------|-------------|
| `uid` | User OID (0 for invalid/non-existing usernames) |
| `successful_auth_count` | Total number of successful logins |
| `last_successful_auth` | Timestamp of last successful login |
| `failed_auth_count` | Total number of failed authentications |
| `last_failed_auth` | Timestamp of last failed login |
| `hba_conflicts_count` | Count of HBA file conflicts |

All login attempts with invalid usernames are aggregated into a single row with OID 0.
