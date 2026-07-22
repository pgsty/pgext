---
title: "pgbson"
linkTitle: "pgbson"
description: "BSON data type and accessor functions for PostgreSQL"
weight: 3910
categories: ["TYPE"]
width: full
---

[**pgbson**](https://github.com/buzzm/postgresbson) : BSON data type and accessor functions for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3910** | {{< badge content="pgbson" link="https://github.com/buzzm/postgresbson" >}} | {{< ext "pgbson" >}} | `2.0.4` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "jsonb_plperl" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "mongo_fdw" >}} {{< ext "documentdb" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} |

> [!Note] PGXN distribution name is bson; CREATE EXTENSION name is pgbson; package release 2.0.4 still installs extension SQL version 2.0; RPM package root is postgresbson and requires libbson.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgbson` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.4` | {{< bg "18" "postgresbson_18" "green" >}} {{< bg "17" "postgresbson_17" "green" >}} {{< bg "16" "postgresbson_16" "green" >}} {{< bg "15" "postgresbson_15" "green" >}} {{< bg "14" "postgresbson_14" "green" >}} | `postgresbson_$v` | `libbson` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.4` | {{< bg "18" "postgresql-18-pgbson" "green" >}} {{< bg "17" "postgresql-17-pgbson" "green" >}} {{< bg "16" "postgresql-16-pgbson" "green" >}} {{< bg "15" "postgresql-15-pgbson" "green" >}} {{< bg "14" "postgresql-14-pgbson" "green" >}} | `postgresql-$v-pgbson` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.4" "postgresql-14-pgbson : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_18` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.4 KiB | [postgresbson_18-2.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_18-2.0.4-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_18` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.0 KiB | [postgresbson_18-2.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_18-2.0.4-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_18` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [postgresbson_18-2.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_18-2.0.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_18` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.6 KiB | [postgresbson_18-2.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_18-2.0.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_18` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [postgresbson_18-2.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_18-2.0.4-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_18` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.8 KiB | [postgresbson_18-2.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_18-2.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgbson` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.7 KiB | [postgresql-18-pgbson_2.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgbson` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.4 KiB | [postgresql-18-pgbson_2.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgbson` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.8 KiB | [postgresql-18-pgbson_2.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgbson` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.5 KiB | [postgresql-18-pgbson_2.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgbson` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 39.4 KiB | [postgresql-18-pgbson_2.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgbson` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 39.0 KiB | [postgresql-18-pgbson_2.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgbson` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.7 KiB | [postgresql-18-pgbson_2.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgbson` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.7 KiB | [postgresql-18-pgbson_2.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgbson` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.6 KiB | [postgresql-18-pgbson_2.0.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgbson` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.5 KiB | [postgresql-18-pgbson_2.0.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_17` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.4 KiB | [postgresbson_17-2.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_17-2.0.4-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_17` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.0 KiB | [postgresbson_17-2.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_17-2.0.4-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_17` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.8 KiB | [postgresbson_17-2.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_17-2.0.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_17` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.6 KiB | [postgresbson_17-2.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_17-2.0.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_17` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.9 KiB | [postgresbson_17-2.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_17-2.0.4-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_17` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.8 KiB | [postgresbson_17-2.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_17-2.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgbson` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.7 KiB | [postgresql-17-pgbson_2.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgbson` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.4 KiB | [postgresql-17-pgbson_2.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgbson` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.8 KiB | [postgresql-17-pgbson_2.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgbson` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.4 KiB | [postgresql-17-pgbson_2.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgbson` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 40.4 KiB | [postgresql-17-pgbson_2.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgbson` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 40.0 KiB | [postgresql-17-pgbson_2.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgbson` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.9 KiB | [postgresql-17-pgbson_2.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgbson` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.7 KiB | [postgresql-17-pgbson_2.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgbson` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.8 KiB | [postgresql-17-pgbson_2.0.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgbson` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.6 KiB | [postgresql-17-pgbson_2.0.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_16` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.4 KiB | [postgresbson_16-2.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_16-2.0.4-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_16` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.0 KiB | [postgresbson_16-2.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_16-2.0.4-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_16` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.8 KiB | [postgresbson_16-2.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_16-2.0.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_16` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.6 KiB | [postgresbson_16-2.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_16-2.0.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_16` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.9 KiB | [postgresbson_16-2.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_16-2.0.4-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_16` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.8 KiB | [postgresbson_16-2.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_16-2.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgbson` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.7 KiB | [postgresql-16-pgbson_2.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgbson` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.4 KiB | [postgresql-16-pgbson_2.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgbson` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.8 KiB | [postgresql-16-pgbson_2.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgbson` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.4 KiB | [postgresql-16-pgbson_2.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgbson` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 40.4 KiB | [postgresql-16-pgbson_2.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgbson` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 40.0 KiB | [postgresql-16-pgbson_2.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgbson` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.9 KiB | [postgresql-16-pgbson_2.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgbson` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.7 KiB | [postgresql-16-pgbson_2.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgbson` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.8 KiB | [postgresql-16-pgbson_2.0.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgbson` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.6 KiB | [postgresql-16-pgbson_2.0.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_15` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.4 KiB | [postgresbson_15-2.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_15-2.0.4-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_15` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.1 KiB | [postgresbson_15-2.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_15-2.0.4-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_15` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [postgresbson_15-2.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_15-2.0.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_15` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.6 KiB | [postgresbson_15-2.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_15-2.0.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_15` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.9 KiB | [postgresbson_15-2.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_15-2.0.4-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_15` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.7 KiB | [postgresbson_15-2.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_15-2.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgbson` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.7 KiB | [postgresql-15-pgbson_2.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgbson` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.3 KiB | [postgresql-15-pgbson_2.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgbson` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.8 KiB | [postgresql-15-pgbson_2.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgbson` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.3 KiB | [postgresql-15-pgbson_2.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgbson` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 40.4 KiB | [postgresql-15-pgbson_2.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgbson` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 39.9 KiB | [postgresql-15-pgbson_2.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgbson` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.8 KiB | [postgresql-15-pgbson_2.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgbson` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.7 KiB | [postgresql-15-pgbson_2.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgbson` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.8 KiB | [postgresql-15-pgbson_2.0.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgbson` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.7 KiB | [postgresql-15-pgbson_2.0.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_14` | `2.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.4 KiB | [postgresbson_14-2.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_14-2.0.4-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_14` | `2.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.1 KiB | [postgresbson_14-2.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_14-2.0.4-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_14` | `2.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.0 KiB | [postgresbson_14-2.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_14-2.0.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_14` | `2.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.6 KiB | [postgresbson_14-2.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_14-2.0.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_14` | `2.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.9 KiB | [postgresbson_14-2.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_14-2.0.4-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_14` | `2.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.7 KiB | [postgresbson_14-2.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_14-2.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgbson` | `2.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.6 KiB | [postgresql-14-pgbson_2.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgbson` | `2.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.3 KiB | [postgresql-14-pgbson_2.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgbson` | `2.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.7 KiB | [postgresql-14-pgbson_2.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgbson` | `2.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.3 KiB | [postgresql-14-pgbson_2.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgbson` | `2.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 40.4 KiB | [postgresql-14-pgbson_2.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgbson` | `2.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 39.9 KiB | [postgresql-14-pgbson_2.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgbson` | `2.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.8 KiB | [postgresql-14-pgbson_2.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgbson` | `2.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.7 KiB | [postgresql-14-pgbson_2.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgbson` | `2.0.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.8 KiB | [postgresql-14-pgbson_2.0.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgbson` | `2.0.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.6 KiB | [postgresql-14-pgbson_2.0.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/buzzm/postgresbson" title="Repository" icon="github" subtitle="github.com/buzzm/postgresbson" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresbson-2.0.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgbson;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgbson;		# install via package name, for the active PG version

pig install pgbson -v 18;   # install for PG 18
pig install pgbson -v 17;   # install for PG 17
pig install pgbson -v 16;   # install for PG 16
pig install pgbson -v 15;   # install for PG 15
pig install pgbson -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgbson;
```

## Usage

Sources:

- [postgresbson README at the 2.0.4 revision](https://github.com/buzzm/postgresbson/blob/ec71d314511d484a99ed510f480919dd0509fbe9/README.md)
- [META.json version 2.0.4](https://github.com/buzzm/postgresbson/blob/ec71d314511d484a99ed510f480919dd0509fbe9/META.json)
- [pgbson control file](https://github.com/buzzm/postgresbson/blob/ec71d314511d484a99ed510f480919dd0509fbe9/pgbson.control)
- [Version 2.0 SQL API](https://github.com/buzzm/postgresbson/blob/ec71d314511d484a99ed510f480919dd0509fbe9/pgbson--2.0.sql)

pgbson adds a BSON data type, typed path accessors, JSON-style operators, casts, and expression-index support. Use it when binary BSON must be stored without first converting every value to JSONB, especially when BSON type fidelity or byte-level round trips matter.

The distribution release is 2.0.4 while the extension control and SQL API version remain 2.0.

### Create the Extension

    CREATE EXTENSION pgbson;

The native module uses libbson. Install a package built against compatible PostgreSQL and libbson versions.

### Store and Validate BSON

The bytea-to-bson cast validates input when a value is written. Version 2.0.4 documents that reads can then assume the stored BSON is valid. Do not bypass the type's input or cast path with unsafe low-level writes.

### Extract Values

Typed dot-path accessors avoid materializing every intermediate object:

    SELECT bson_get_datetime(payload, 'msg.header.event.ts'),
           bson_get_string(payload, 'data.customer.name')
    FROM events;

Use bson_get_bson for a subdocument:

    SELECT bson_get_bson(payload, 'msg.header.event')
    FROM events;

JSON-style navigation is also available:

    SELECT payload->'msg'->'header'->'event'->>'ts'
    FROM events;

### Function and Operator Index

- bson_get_string, bson_get_int32, bson_get_int64, bson_get_double, bson_get_decimal: typed scalar accessors.
- bson_get_datetime, bson_get_binary, bson_get_boolean: accessors for additional BSON types.
- bson_get_bson: return an embedded BSON document.
- bson_get_jsonb_array: convert an array endpoint to a PostgreSQL jsonb array.
- -> and ->>: navigate values with JSON-like syntax.
- bson casts to json and jsonb: expose Extended JSON for PostgreSQL JSON processing.
- bson and bytea casts: preserve the BSON binary representation.

### Index and Interoperate

Create expression indexes on frequently queried paths:

    CREATE INDEX events_customer_id_idx
    ON events (bson_get_string(payload, 'data.customer.id'));

Cast a subdocument to jsonb when PostgreSQL's JSON operators are more convenient:

    SELECT bson_get_bson(payload, 'msg.header')::jsonb ? 'event'
    FROM events;

### Caveats

- A typed getter returns useful data only when the endpoint has the expected BSON type. Make type expectations explicit in ingestion code.
- bson_get_bson returns NULL for scalar endpoints because a scalar is not a BSON document.
- Dot-path accessors are generally preferable to long operator chains for repeated extraction because they avoid intermediate BSON values.
- BSON and JSONB have different type and ordering semantics. A cast can be useful but is not a lossless replacement for every BSON workflow.
