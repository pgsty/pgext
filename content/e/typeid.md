---
title: "typeid"
linkTitle: "typeid"
description: "Allows to use TypeIDs in Postgres natively"
weight: 4580
categories: ["FUNC"]
languages: ["Rust"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_typeid**](https://github.com/blitss/typeid-postgres) : Allows to use TypeIDs in Postgres natively


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4580** | {{< badge content="typeid" link="https://github.com/blitss/typeid-postgres" >}} | {{< ext "typeid" "pg_typeid" >}} | `0.4.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_uuidv7" >}} {{< ext "sequential_uuids" >}} {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "snowflake" >}} {{< ext "pg_uuid_v8" >}} {{< ext "uuid-ossp" >}} {{< ext "pg_hashids" >}} {{< ext "pg_slug_gen" >}} {{< ext "base36" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_typeid` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.1` | {{< bg "18" "pg_typeid_18" "green" >}} {{< bg "17" "pg_typeid_17" "green" >}} {{< bg "16" "pg_typeid_16" "green" >}} {{< bg "15" "pg_typeid_15" "green" >}} {{< bg "14" "pg_typeid_14" "green" >}} | `pg_typeid_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.1` | {{< bg "18" "postgresql-18-typeid" "green" >}} {{< bg "17" "postgresql-17-typeid" "green" >}} {{< bg "16" "postgresql-16-typeid" "green" >}} {{< bg "15" "postgresql-15-typeid" "green" >}} {{< bg "14" "postgresql-14-typeid" "green" >}} | `postgresql-$v-typeid` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_typeid_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-typeid : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-typeid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-typeid : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_typeid_18` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 892.3 KiB | [pg_typeid_18-0.4.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_typeid_18-0.4.1-1PGSTY.el8.x86_64.rpm) |
| `pg_typeid_18` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 775.5 KiB | [pg_typeid_18-0.4.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_typeid_18-0.4.1-1PGSTY.el8.aarch64.rpm) |
| `pg_typeid_18` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 913.4 KiB | [pg_typeid_18-0.4.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_typeid_18-0.4.1-1PGSTY.el9.x86_64.rpm) |
| `pg_typeid_18` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 827.7 KiB | [pg_typeid_18-0.4.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_typeid_18-0.4.1-1PGSTY.el9.aarch64.rpm) |
| `pg_typeid_18` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 907.3 KiB | [pg_typeid_18-0.4.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_typeid_18-0.4.1-1PGSTY.el10.x86_64.rpm) |
| `pg_typeid_18` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 807.4 KiB | [pg_typeid_18-0.4.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_typeid_18-0.4.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-typeid` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 741.1 KiB | [postgresql-18-typeid_0.4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-18-typeid_0.4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-typeid` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 621.0 KiB | [postgresql-18-typeid_0.4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-18-typeid_0.4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-typeid` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 742.9 KiB | [postgresql-18-typeid_0.4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-18-typeid_0.4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-typeid` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 621.9 KiB | [postgresql-18-typeid_0.4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-18-typeid_0.4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-typeid` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 826.4 KiB | [postgresql-18-typeid_0.4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-18-typeid_0.4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-typeid` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 736.6 KiB | [postgresql-18-typeid_0.4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-18-typeid_0.4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-typeid` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 824.1 KiB | [postgresql-18-typeid_0.4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-18-typeid_0.4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-typeid` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 726.1 KiB | [postgresql-18-typeid_0.4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-18-typeid_0.4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-typeid` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 815.3 KiB | [postgresql-18-typeid_0.4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-18-typeid_0.4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-typeid` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 725.1 KiB | [postgresql-18-typeid_0.4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-18-typeid_0.4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_typeid_17` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 889.3 KiB | [pg_typeid_17-0.4.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_typeid_17-0.4.1-1PGSTY.el8.x86_64.rpm) |
| `pg_typeid_17` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 771.0 KiB | [pg_typeid_17-0.4.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_typeid_17-0.4.1-1PGSTY.el8.aarch64.rpm) |
| `pg_typeid_17` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 906.3 KiB | [pg_typeid_17-0.4.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_typeid_17-0.4.1-1PGSTY.el9.x86_64.rpm) |
| `pg_typeid_17` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 825.0 KiB | [pg_typeid_17-0.4.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_typeid_17-0.4.1-1PGSTY.el9.aarch64.rpm) |
| `pg_typeid_17` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 910.6 KiB | [pg_typeid_17-0.4.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_typeid_17-0.4.1-1PGSTY.el10.x86_64.rpm) |
| `pg_typeid_17` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 806.5 KiB | [pg_typeid_17-0.4.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_typeid_17-0.4.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-typeid` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 739.2 KiB | [postgresql-17-typeid_0.4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-17-typeid_0.4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-typeid` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 620.2 KiB | [postgresql-17-typeid_0.4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-17-typeid_0.4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-typeid` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 740.5 KiB | [postgresql-17-typeid_0.4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-17-typeid_0.4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-typeid` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 620.2 KiB | [postgresql-17-typeid_0.4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-17-typeid_0.4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-typeid` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 823.9 KiB | [postgresql-17-typeid_0.4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-17-typeid_0.4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-typeid` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 734.9 KiB | [postgresql-17-typeid_0.4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-17-typeid_0.4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-typeid` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 820.5 KiB | [postgresql-17-typeid_0.4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-17-typeid_0.4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-typeid` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 723.8 KiB | [postgresql-17-typeid_0.4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-17-typeid_0.4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-typeid` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 812.2 KiB | [postgresql-17-typeid_0.4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-17-typeid_0.4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-typeid` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 723.7 KiB | [postgresql-17-typeid_0.4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-17-typeid_0.4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_typeid_16` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 888.1 KiB | [pg_typeid_16-0.4.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_typeid_16-0.4.1-1PGSTY.el8.x86_64.rpm) |
| `pg_typeid_16` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 768.2 KiB | [pg_typeid_16-0.4.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_typeid_16-0.4.1-1PGSTY.el8.aarch64.rpm) |
| `pg_typeid_16` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 902.0 KiB | [pg_typeid_16-0.4.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_typeid_16-0.4.1-1PGSTY.el9.x86_64.rpm) |
| `pg_typeid_16` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 823.1 KiB | [pg_typeid_16-0.4.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_typeid_16-0.4.1-1PGSTY.el9.aarch64.rpm) |
| `pg_typeid_16` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 902.9 KiB | [pg_typeid_16-0.4.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_typeid_16-0.4.1-1PGSTY.el10.x86_64.rpm) |
| `pg_typeid_16` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 806.1 KiB | [pg_typeid_16-0.4.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_typeid_16-0.4.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-typeid` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 737.8 KiB | [postgresql-16-typeid_0.4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-16-typeid_0.4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-typeid` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 618.6 KiB | [postgresql-16-typeid_0.4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-16-typeid_0.4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-typeid` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 739.7 KiB | [postgresql-16-typeid_0.4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-16-typeid_0.4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-typeid` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 618.7 KiB | [postgresql-16-typeid_0.4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-16-typeid_0.4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-typeid` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 824.7 KiB | [postgresql-16-typeid_0.4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-16-typeid_0.4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-typeid` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 733.3 KiB | [postgresql-16-typeid_0.4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-16-typeid_0.4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-typeid` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 812.8 KiB | [postgresql-16-typeid_0.4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-16-typeid_0.4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-typeid` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 722.4 KiB | [postgresql-16-typeid_0.4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-16-typeid_0.4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-typeid` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 809.4 KiB | [postgresql-16-typeid_0.4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-16-typeid_0.4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-typeid` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 721.2 KiB | [postgresql-16-typeid_0.4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-16-typeid_0.4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_typeid_15` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 879.8 KiB | [pg_typeid_15-0.4.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_typeid_15-0.4.1-1PGSTY.el8.x86_64.rpm) |
| `pg_typeid_15` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 760.6 KiB | [pg_typeid_15-0.4.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_typeid_15-0.4.1-1PGSTY.el8.aarch64.rpm) |
| `pg_typeid_15` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 896.1 KiB | [pg_typeid_15-0.4.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_typeid_15-0.4.1-1PGSTY.el9.x86_64.rpm) |
| `pg_typeid_15` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 814.0 KiB | [pg_typeid_15-0.4.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_typeid_15-0.4.1-1PGSTY.el9.aarch64.rpm) |
| `pg_typeid_15` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 897.0 KiB | [pg_typeid_15-0.4.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_typeid_15-0.4.1-1PGSTY.el10.x86_64.rpm) |
| `pg_typeid_15` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 804.2 KiB | [pg_typeid_15-0.4.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_typeid_15-0.4.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-typeid` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 734.1 KiB | [postgresql-15-typeid_0.4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-15-typeid_0.4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-typeid` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 613.9 KiB | [postgresql-15-typeid_0.4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-15-typeid_0.4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-typeid` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 733.7 KiB | [postgresql-15-typeid_0.4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-15-typeid_0.4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-typeid` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 615.3 KiB | [postgresql-15-typeid_0.4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-15-typeid_0.4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-typeid` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 817.1 KiB | [postgresql-15-typeid_0.4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-15-typeid_0.4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-typeid` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 726.6 KiB | [postgresql-15-typeid_0.4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-15-typeid_0.4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-typeid` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 808.4 KiB | [postgresql-15-typeid_0.4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-15-typeid_0.4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-typeid` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 717.2 KiB | [postgresql-15-typeid_0.4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-15-typeid_0.4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-typeid` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 804.9 KiB | [postgresql-15-typeid_0.4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-15-typeid_0.4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-typeid` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 715.7 KiB | [postgresql-15-typeid_0.4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-15-typeid_0.4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_typeid_14` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 875.3 KiB | [pg_typeid_14-0.4.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_typeid_14-0.4.1-1PGSTY.el8.x86_64.rpm) |
| `pg_typeid_14` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 759.1 KiB | [pg_typeid_14-0.4.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_typeid_14-0.4.1-1PGSTY.el8.aarch64.rpm) |
| `pg_typeid_14` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 891.9 KiB | [pg_typeid_14-0.4.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_typeid_14-0.4.1-1PGSTY.el9.x86_64.rpm) |
| `pg_typeid_14` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 812.4 KiB | [pg_typeid_14-0.4.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_typeid_14-0.4.1-1PGSTY.el9.aarch64.rpm) |
| `pg_typeid_14` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 893.2 KiB | [pg_typeid_14-0.4.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_typeid_14-0.4.1-1PGSTY.el10.x86_64.rpm) |
| `pg_typeid_14` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 801.4 KiB | [pg_typeid_14-0.4.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_typeid_14-0.4.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-typeid` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 731.1 KiB | [postgresql-14-typeid_0.4.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-14-typeid_0.4.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-typeid` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 612.0 KiB | [postgresql-14-typeid_0.4.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-typeid/postgresql-14-typeid_0.4.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-typeid` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 732.1 KiB | [postgresql-14-typeid_0.4.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-14-typeid_0.4.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-typeid` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 612.7 KiB | [postgresql-14-typeid_0.4.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-typeid/postgresql-14-typeid_0.4.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-typeid` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 813.3 KiB | [postgresql-14-typeid_0.4.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-14-typeid_0.4.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-typeid` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 724.5 KiB | [postgresql-14-typeid_0.4.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-typeid/postgresql-14-typeid_0.4.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-typeid` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 807.1 KiB | [postgresql-14-typeid_0.4.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-14-typeid_0.4.1-1PGSTY~noble_amd64.deb) |
| `postgresql-14-typeid` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 715.2 KiB | [postgresql-14-typeid_0.4.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-typeid/postgresql-14-typeid_0.4.1-1PGSTY~noble_arm64.deb) |
| `postgresql-14-typeid` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 800.6 KiB | [postgresql-14-typeid_0.4.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-14-typeid_0.4.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-typeid` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 712.9 KiB | [postgresql-14-typeid_0.4.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-typeid/postgresql-14-typeid_0.4.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/blitss/typeid-postgres" title="Repository" icon="github" subtitle="github.com/blitss/typeid-postgres" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="typeid-postgres-0.4.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_typeid;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_typeid;		# install via package name, for the active PG version
pig install typeid;		# install by extension name, for the current active PG version

pig install typeid -v 18;   # install for PG 18
pig install typeid -v 17;   # install for PG 17
pig install typeid -v 16;   # install for PG 16
pig install typeid -v 15;   # install for PG 15
pig install typeid -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

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
{.wide .expand}

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
