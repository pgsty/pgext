---
title: "typeid"
linkTitle: "typeid"
description: "Allows to use TypeIDs in Postgres natively"
weight: 4580
categories: ["FUNC"]
width: full
---

[**pg_typeid**](https://github.com/blitss/typeid-postgres) : Allows to use TypeIDs in Postgres natively


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4580** | {{< badge content="typeid" link="https://github.com/blitss/typeid-postgres" >}} | {{< ext "typeid" "pg_typeid" >}} | `0.3.0` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pg_uuidv7" >}} {{< ext "pgx_ulid" >}} {{< ext "uuid-ossp" >}} {{< ext "pg_hashids" >}} {{< ext "permuteseq" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_typeid` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "pg_typeid_18" "green" >}} {{< bg "17" "pg_typeid_17" "green" >}} {{< bg "16" "pg_typeid_16" "green" >}} {{< bg "15" "pg_typeid_15" "green" >}} {{< bg "14" "pg_typeid_14" "green" >}} | `pg_typeid_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-typeid" "green" >}} {{< bg "17" "postgresql-17-typeid" "green" >}} {{< bg "16" "postgresql-16-typeid" "green" >}} {{< bg "15" "postgresql-15-typeid" "green" >}} {{< bg "14" "postgresql-14-typeid" "green" >}} | `postgresql-$v-typeid` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_typeid_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-typeid : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_typeid_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 909.2 KiB | [pg_typeid_18-0.3.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_typeid_18-0.3.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_typeid_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 818.3 KiB | [pg_typeid_18-0.3.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_typeid_18-0.3.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_typeid_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 918.4 KiB | [pg_typeid_18-0.3.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_typeid_18-0.3.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_typeid_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 866.2 KiB | [pg_typeid_18-0.3.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_typeid_18-0.3.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_typeid_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 919.4 KiB | [pg_typeid_18-0.3.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_typeid_18-0.3.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_typeid_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 845.9 KiB | [pg_typeid_18-0.3.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_typeid_18-0.3.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-typeid` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 724.7 KiB | [postgresql-18-typeid_0.3.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-18-typeid_0.3.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-typeid` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 606.9 KiB | [postgresql-18-typeid_0.3.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-18-typeid_0.3.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-typeid` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 724.4 KiB | [postgresql-18-typeid_0.3.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-18-typeid_0.3.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-typeid` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 607.1 KiB | [postgresql-18-typeid_0.3.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-18-typeid_0.3.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-typeid` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 808.8 KiB | [postgresql-18-typeid_0.3.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-18-typeid_0.3.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-typeid` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 716.6 KiB | [postgresql-18-typeid_0.3.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-18-typeid_0.3.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-typeid` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 799.2 KiB | [postgresql-18-typeid_0.3.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-18-typeid_0.3.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-typeid` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 707.4 KiB | [postgresql-18-typeid_0.3.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-18-typeid_0.3.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-18-typeid` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 795.0 KiB | [postgresql-18-typeid_0.3.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-18-typeid_0.3.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-18-typeid` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 706.1 KiB | [postgresql-18-typeid_0.3.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-18-typeid_0.3.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_typeid_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 906.3 KiB | [pg_typeid_17-0.3.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_typeid_17-0.3.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_typeid_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 815.5 KiB | [pg_typeid_17-0.3.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_typeid_17-0.3.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_typeid_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 914.9 KiB | [pg_typeid_17-0.3.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_typeid_17-0.3.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_typeid_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 863.5 KiB | [pg_typeid_17-0.3.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_typeid_17-0.3.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_typeid_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 914.6 KiB | [pg_typeid_17-0.3.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_typeid_17-0.3.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_typeid_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 845.3 KiB | [pg_typeid_17-0.3.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_typeid_17-0.3.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-typeid` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 722.7 KiB | [postgresql-17-typeid_0.3.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-17-typeid_0.3.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-typeid` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 606.0 KiB | [postgresql-17-typeid_0.3.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-17-typeid_0.3.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-typeid` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 722.9 KiB | [postgresql-17-typeid_0.3.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-17-typeid_0.3.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-typeid` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 606.4 KiB | [postgresql-17-typeid_0.3.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-17-typeid_0.3.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-typeid` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 805.7 KiB | [postgresql-17-typeid_0.3.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-17-typeid_0.3.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-typeid` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 715.3 KiB | [postgresql-17-typeid_0.3.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-17-typeid_0.3.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-typeid` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 796.5 KiB | [postgresql-17-typeid_0.3.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-17-typeid_0.3.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-typeid` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 705.9 KiB | [postgresql-17-typeid_0.3.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-17-typeid_0.3.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-17-typeid` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 792.2 KiB | [postgresql-17-typeid_0.3.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-17-typeid_0.3.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-17-typeid` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 704.1 KiB | [postgresql-17-typeid_0.3.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-17-typeid_0.3.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_typeid_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 904.5 KiB | [pg_typeid_16-0.3.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_typeid_16-0.3.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_typeid_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 813.5 KiB | [pg_typeid_16-0.3.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_typeid_16-0.3.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_typeid_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 912.9 KiB | [pg_typeid_16-0.3.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_typeid_16-0.3.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_typeid_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 862.0 KiB | [pg_typeid_16-0.3.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_typeid_16-0.3.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_typeid_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 913.1 KiB | [pg_typeid_16-0.3.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_typeid_16-0.3.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_typeid_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 844.6 KiB | [pg_typeid_16-0.3.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_typeid_16-0.3.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-typeid` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 721.2 KiB | [postgresql-16-typeid_0.3.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-16-typeid_0.3.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-typeid` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 604.6 KiB | [postgresql-16-typeid_0.3.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-16-typeid_0.3.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-typeid` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 721.4 KiB | [postgresql-16-typeid_0.3.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-16-typeid_0.3.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-typeid` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 605.2 KiB | [postgresql-16-typeid_0.3.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-16-typeid_0.3.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-typeid` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 804.7 KiB | [postgresql-16-typeid_0.3.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-16-typeid_0.3.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-typeid` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 714.3 KiB | [postgresql-16-typeid_0.3.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-16-typeid_0.3.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-typeid` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 795.7 KiB | [postgresql-16-typeid_0.3.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-16-typeid_0.3.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-typeid` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 705.0 KiB | [postgresql-16-typeid_0.3.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-16-typeid_0.3.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-16-typeid` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 790.3 KiB | [postgresql-16-typeid_0.3.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-16-typeid_0.3.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-16-typeid` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 703.0 KiB | [postgresql-16-typeid_0.3.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-16-typeid_0.3.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_typeid_15` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 895.0 KiB | [pg_typeid_15-0.3.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_typeid_15-0.3.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_typeid_15` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 805.0 KiB | [pg_typeid_15-0.3.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_typeid_15-0.3.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_typeid_15` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 903.0 KiB | [pg_typeid_15-0.3.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_typeid_15-0.3.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_typeid_15` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 852.5 KiB | [pg_typeid_15-0.3.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_typeid_15-0.3.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_typeid_15` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 903.8 KiB | [pg_typeid_15-0.3.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_typeid_15-0.3.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_typeid_15` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 840.2 KiB | [pg_typeid_15-0.3.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_typeid_15-0.3.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-typeid` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 715.6 KiB | [postgresql-15-typeid_0.3.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-15-typeid_0.3.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-typeid` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 600.6 KiB | [postgresql-15-typeid_0.3.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-15-typeid_0.3.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-typeid` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 714.9 KiB | [postgresql-15-typeid_0.3.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-15-typeid_0.3.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-typeid` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 600.4 KiB | [postgresql-15-typeid_0.3.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-15-typeid_0.3.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-typeid` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 795.1 KiB | [postgresql-15-typeid_0.3.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-15-typeid_0.3.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-typeid` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 708.8 KiB | [postgresql-15-typeid_0.3.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-15-typeid_0.3.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-typeid` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 787.4 KiB | [postgresql-15-typeid_0.3.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-15-typeid_0.3.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-typeid` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 699.1 KiB | [postgresql-15-typeid_0.3.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-15-typeid_0.3.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-15-typeid` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 783.8 KiB | [postgresql-15-typeid_0.3.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-15-typeid_0.3.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-15-typeid` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 697.1 KiB | [postgresql-15-typeid_0.3.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-15-typeid_0.3.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_typeid_14` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 892.1 KiB | [pg_typeid_14-0.3.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_typeid_14-0.3.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_typeid_14` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 802.5 KiB | [pg_typeid_14-0.3.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_typeid_14-0.3.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_typeid_14` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 900.8 KiB | [pg_typeid_14-0.3.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_typeid_14-0.3.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_typeid_14` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 849.7 KiB | [pg_typeid_14-0.3.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_typeid_14-0.3.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_typeid_14` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 899.9 KiB | [pg_typeid_14-0.3.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_typeid_14-0.3.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_typeid_14` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 839.0 KiB | [pg_typeid_14-0.3.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_typeid_14-0.3.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-typeid` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 713.4 KiB | [postgresql-14-typeid_0.3.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-14-typeid_0.3.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-typeid` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 598.3 KiB | [postgresql-14-typeid_0.3.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-14-typeid_0.3.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-typeid` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 712.7 KiB | [postgresql-14-typeid_0.3.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-14-typeid_0.3.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-typeid` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 599.5 KiB | [postgresql-14-typeid_0.3.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-14-typeid_0.3.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-typeid` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 795.6 KiB | [postgresql-14-typeid_0.3.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-14-typeid_0.3.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-typeid` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 708.0 KiB | [postgresql-14-typeid_0.3.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-14-typeid_0.3.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-typeid` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 784.9 KiB | [postgresql-14-typeid_0.3.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-14-typeid_0.3.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-typeid` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 697.8 KiB | [postgresql-14-typeid_0.3.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-14-typeid_0.3.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-14-typeid` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 781.3 KiB | [postgresql-14-typeid_0.3.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-14-typeid_0.3.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-14-typeid` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 695.8 KiB | [postgresql-14-typeid_0.3.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-14-typeid_0.3.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/blitss/typeid-postgres" title="Repository" icon="github" subtitle="github.com/blitss/typeid-postgres" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="typeid-postgres-0.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_typeid;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_typeid;		# install via package name, for the active PG version
pig install typeid;		# install by extension name, for the current active PG version

pig install typeid -v 18;   # install for PG 18
pig install typeid -v 17;   # install for PG 17
pig install typeid -v 16;   # install for PG 16
pig install typeid -v 15;   # install for PG 15
pig install typeid -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION typeid;
```




## Usage

> [typeid: TypeID support for PostgreSQL - type-safe, sortable UUIDs with a prefix](https://github.com/blitss/typeid-postgres)

TypeID is an extension of UUIDv7 with a type prefix, stored internally as a prefix + binary UUID.

```sql
CREATE EXTENSION typeid;
```

### Functions

| Function | Return Type | Description |
|---|---|---|
| `typeid_generate(prefix TEXT)` | `typeid` | Generate a new TypeID with the given prefix |
| `typeid_generate_nil()` | `typeid` | Generate a TypeID with an empty prefix |
| `typeid_is_valid(input TEXT)` | `BOOLEAN` | Check if a TypeID string is valid |
| `typeid_prefix(typeid)` | `TEXT` | Extract the prefix from a TypeID |
| `typeid_to_uuid(typeid)` | `UUID` | Convert a TypeID to a UUID |
| `uuid_to_typeid(prefix TEXT, uuid UUID)` | `typeid` | Convert a UUID to a TypeID |
| `typeid_uuid_generate_v7()` | `UUID` | Generate a UUID v7 |
| `typeid_has_prefix(typeid, prefix TEXT)` | `BOOLEAN` | Check if a TypeID has a specific prefix |
| `typeid_is_nil_prefix(typeid)` | `BOOLEAN` | Check if a TypeID has a nil prefix |
| `typeid_generate_batch(prefix TEXT, count INTEGER)` | `SETOF typeid` | Generate a batch of TypeIDs |

### Operators

- `<`, `<=`, `=`, `>=`, `>`, `<>` for comparing TypeIDs
- `@>` for checking if a TypeID has a certain prefix (e.g. `id @> 'user'`)
- B-tree operator class for indexing

### Examples

```sql
-- Create table with TypeID primary key
CREATE TABLE users (
  id typeid DEFAULT typeid_generate('user') NOT NULL PRIMARY KEY,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Insert data
INSERT INTO users (id) SELECT typeid_generate('user') FROM generate_series(1, 100);

-- Extract prefix
SELECT typeid_prefix(id) FROM users LIMIT 1;  -- 'user'

-- Check prefix with operator
SELECT * FROM users WHERE id @> 'user';

-- Convert to UUID
SELECT typeid_to_uuid(id) FROM users LIMIT 1;
```
