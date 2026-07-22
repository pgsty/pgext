---
title: "pg_when"
linkTitle: "pg_when"
description: "Natural language time parsing for PostgreSQL"
weight: 1120
categories: ["TIME"]
width: full
---

[**pg_when**](https://github.com/frectonz/pg-when) : Natural language time parsing for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1120** | {{< badge content="pg_when" link="https://github.com/frectonz/pg-when" >}} | {{< ext "pg_when" >}} | `0.1.9` | {{< category "TIME" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.9` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_when` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.9` | {{< bg "18" "pg_when_18" "green" >}} {{< bg "17" "pg_when_17" "green" >}} {{< bg "16" "pg_when_16" "green" >}} {{< bg "15" "pg_when_15" "green" >}} {{< bg "14" "pg_when_14" "green" >}} | `pg_when_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.9` | {{< bg "18" "postgresql-18-pg-when" "green" >}} {{< bg "17" "postgresql-17-pg-when" "green" >}} {{< bg "16" "postgresql-16-pg-when" "green" >}} {{< bg "15" "postgresql-15-pg-when" "green" >}} {{< bg "14" "postgresql-14-pg-when" "green" >}} | `postgresql-$v-pg-when` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_18` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_when_18-0.1.9-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_18-0.1.9-3PIGSTY.el8.x86_64.rpm) |
| `pg_when_18` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 981.7 KiB | [pg_when_18-0.1.9-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_18-0.1.9-3PIGSTY.el8.aarch64.rpm) |
| `pg_when_18` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_when_18-0.1.9-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_18-0.1.9-3PIGSTY.el9.x86_64.rpm) |
| `pg_when_18` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_when_18-0.1.9-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_18-0.1.9-3PIGSTY.el9.aarch64.rpm) |
| `pg_when_18` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_when_18-0.1.9-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_18-0.1.9-3PIGSTY.el10.x86_64.rpm) |
| `pg_when_18` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1017.5 KiB | [pg_when_18-0.1.9-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_18-0.1.9-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-when` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 872.3 KiB | [postgresql-18-pg-when_0.1.9-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 744.8 KiB | [postgresql-18-pg-when_0.1.9-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 873.1 KiB | [postgresql-18-pg-when_0.1.9-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 744.8 KiB | [postgresql-18-pg-when_0.1.9-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 964.0 KiB | [postgresql-18-pg-when_0.1.9-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 873.9 KiB | [postgresql-18-pg-when_0.1.9-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 956.6 KiB | [postgresql-18-pg-when_0.1.9-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 863.8 KiB | [postgresql-18-pg-when_0.1.9-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-3PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 952.5 KiB | [postgresql-18-pg-when_0.1.9-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-3PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 861.5 KiB | [postgresql-18-pg-when_0.1.9-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_17` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_when_17-0.1.9-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_17-0.1.9-3PIGSTY.el8.x86_64.rpm) |
| `pg_when_17` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 978.9 KiB | [pg_when_17-0.1.9-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_17-0.1.9-3PIGSTY.el8.aarch64.rpm) |
| `pg_when_17` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_when_17-0.1.9-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_17-0.1.9-3PIGSTY.el9.x86_64.rpm) |
| `pg_when_17` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_when_17-0.1.9-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_17-0.1.9-3PIGSTY.el9.aarch64.rpm) |
| `pg_when_17` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_when_17-0.1.9-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_17-0.1.9-3PIGSTY.el10.x86_64.rpm) |
| `pg_when_17` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1017.3 KiB | [pg_when_17-0.1.9-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_17-0.1.9-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-when` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 870.5 KiB | [postgresql-17-pg-when_0.1.9-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 743.3 KiB | [postgresql-17-pg-when_0.1.9-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 870.5 KiB | [postgresql-17-pg-when_0.1.9-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 742.9 KiB | [postgresql-17-pg-when_0.1.9-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 963.3 KiB | [postgresql-17-pg-when_0.1.9-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 871.8 KiB | [postgresql-17-pg-when_0.1.9-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 954.5 KiB | [postgresql-17-pg-when_0.1.9-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 861.8 KiB | [postgresql-17-pg-when_0.1.9-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-3PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 950.3 KiB | [postgresql-17-pg-when_0.1.9-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-3PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 859.5 KiB | [postgresql-17-pg-when_0.1.9-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_16` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_when_16-0.1.9-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_16-0.1.9-3PIGSTY.el8.x86_64.rpm) |
| `pg_when_16` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 976.9 KiB | [pg_when_16-0.1.9-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_16-0.1.9-3PIGSTY.el8.aarch64.rpm) |
| `pg_when_16` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_when_16-0.1.9-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_16-0.1.9-3PIGSTY.el9.x86_64.rpm) |
| `pg_when_16` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_when_16-0.1.9-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_16-0.1.9-3PIGSTY.el9.aarch64.rpm) |
| `pg_when_16` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_when_16-0.1.9-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_16-0.1.9-3PIGSTY.el10.x86_64.rpm) |
| `pg_when_16` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1016.9 KiB | [pg_when_16-0.1.9-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_16-0.1.9-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-when` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 869.6 KiB | [postgresql-16-pg-when_0.1.9-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 742.0 KiB | [postgresql-16-pg-when_0.1.9-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 869.5 KiB | [postgresql-16-pg-when_0.1.9-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 742.5 KiB | [postgresql-16-pg-when_0.1.9-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 962.5 KiB | [postgresql-16-pg-when_0.1.9-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 870.9 KiB | [postgresql-16-pg-when_0.1.9-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 953.7 KiB | [postgresql-16-pg-when_0.1.9-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 860.9 KiB | [postgresql-16-pg-when_0.1.9-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-3PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 948.2 KiB | [postgresql-16-pg-when_0.1.9-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-3PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 858.7 KiB | [postgresql-16-pg-when_0.1.9-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_15` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.0 MiB | [pg_when_15-0.1.9-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_15-0.1.9-3PIGSTY.el8.x86_64.rpm) |
| `pg_when_15` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 968.0 KiB | [pg_when_15-0.1.9-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_15-0.1.9-3PIGSTY.el8.aarch64.rpm) |
| `pg_when_15` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.0 MiB | [pg_when_15-0.1.9-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_15-0.1.9-3PIGSTY.el9.x86_64.rpm) |
| `pg_when_15` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_when_15-0.1.9-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_15-0.1.9-3PIGSTY.el9.aarch64.rpm) |
| `pg_when_15` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_when_15-0.1.9-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_15-0.1.9-3PIGSTY.el10.x86_64.rpm) |
| `pg_when_15` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1013.5 KiB | [pg_when_15-0.1.9-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_15-0.1.9-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-when` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 864.0 KiB | [postgresql-15-pg-when_0.1.9-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 737.1 KiB | [postgresql-15-pg-when_0.1.9-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 864.0 KiB | [postgresql-15-pg-when_0.1.9-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 737.3 KiB | [postgresql-15-pg-when_0.1.9-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 954.2 KiB | [postgresql-15-pg-when_0.1.9-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 866.3 KiB | [postgresql-15-pg-when_0.1.9-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 945.6 KiB | [postgresql-15-pg-when_0.1.9-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 854.9 KiB | [postgresql-15-pg-when_0.1.9-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-3PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 941.8 KiB | [postgresql-15-pg-when_0.1.9-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-3PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 853.4 KiB | [postgresql-15-pg-when_0.1.9-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_14` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.0 MiB | [pg_when_14-0.1.9-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_14-0.1.9-3PIGSTY.el8.x86_64.rpm) |
| `pg_when_14` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 965.7 KiB | [pg_when_14-0.1.9-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_14-0.1.9-3PIGSTY.el8.aarch64.rpm) |
| `pg_when_14` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.0 MiB | [pg_when_14-0.1.9-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_14-0.1.9-3PIGSTY.el9.x86_64.rpm) |
| `pg_when_14` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1021.5 KiB | [pg_when_14-0.1.9-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_14-0.1.9-3PIGSTY.el9.aarch64.rpm) |
| `pg_when_14` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_when_14-0.1.9-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_14-0.1.9-3PIGSTY.el10.x86_64.rpm) |
| `pg_when_14` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1011.0 KiB | [pg_when_14-0.1.9-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_14-0.1.9-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-when` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 862.2 KiB | [postgresql-14-pg-when_0.1.9-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 735.6 KiB | [postgresql-14-pg-when_0.1.9-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 862.2 KiB | [postgresql-14-pg-when_0.1.9-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 736.5 KiB | [postgresql-14-pg-when_0.1.9-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 952.6 KiB | [postgresql-14-pg-when_0.1.9-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 863.7 KiB | [postgresql-14-pg-when_0.1.9-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 943.5 KiB | [postgresql-14-pg-when_0.1.9-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 854.3 KiB | [postgresql-14-pg-when_0.1.9-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-3PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 939.8 KiB | [postgresql-14-pg-when_0.1.9-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-3PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 850.8 KiB | [postgresql-14-pg-when_0.1.9-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/frectonz/pg-when" title="Repository" icon="github" subtitle="github.com/frectonz/pg-when" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_when-0.1.9.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_when;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_when;		# install via package name, for the active PG version

pig install pg_when -v 18;   # install for PG 18
pig install pg_when -v 17;   # install for PG 17
pig install pg_when -v 16;   # install for PG 16
pig install pg_when -v 15;   # install for PG 15
pig install pg_when -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_when;
```




## Usage

Sources: [README](https://github.com/frectonz/pg-when/blob/main/README.md), [Cargo.toml version 0.1.9](https://github.com/frectonz/pg-when/blob/main/Cargo.toml), [META.json](https://github.com/frectonz/pg-when/blob/main/META.json)

`pg-when` parses a constrained natural-language time expression and returns either a PostgreSQL timestamp with time zone or Unix epoch values at different resolutions.

```sql
CREATE EXTENSION pg_when;

SELECT when_is('next friday at 8:00 pm in America/New_York');
SELECT seconds_at('next friday at 8:00 pm in America/New_York');
SELECT millis_at('next friday at 8:00 pm in America/New_York');
SELECT micros_at('next friday at 8:00 pm in America/New_York');
SELECT nanos_at('next friday at 8:00 pm in America/New_York');
```

### Supported Query Shape

The parser accepts up to three parts:

```sql
SELECT when_is('<date> at <time> in <timezone>');
SELECT when_is('<date>');
SELECT when_is('<time> in <timezone>');
SELECT when_is('<date> at <time>');
```

If no timezone is provided, upstream says the default is UTC.

### Common Inputs

- relative dates: `today`, `tomorrow`, `last month`, `this friday`, `5 days ago`, `in 2 years`
- exact dates: `YYYY-MM-DD`, `DD/MM/YYYY`, `January 10, 2004`, `10 Jan 2004`
- relative times: `noon`, `midnight`, `morning`, `evening`, `next hour`
- exact times: `8:30 pm`, `15:45`
- time zones: `America/New_York`, `Europe/London`, `UTC-08:00`, `UTC+05:30`

### Examples

```sql
SELECT when_is('5 days ago at this hour in Asia/Tokyo');
SELECT when_is('in 2 months at midnight in UTC-8');
SELECT when_is('December 31, 2026 at evening');
```

### Caveats

- The extension is aimed at the documented grammar above, not arbitrary English.
- Upstream still lists source/runtime support and Docker image examples from PostgreSQL 13 through 18, but this repo's package matrix is PostgreSQL 14 through 18 only; do not assume Pigsty packages for PostgreSQL 13.
- Upstream `Cargo.toml` currently pins `pgrx` 0.15.0; this repo's package metadata records a manual `pgrx` 0.17.0 upgrade.
