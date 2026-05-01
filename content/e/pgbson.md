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
| **3910** | {{< badge content="pgbson" link="https://github.com/buzzm/postgresbson" >}} | {{< ext "pgbson" >}} | `2.0.2` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "jsonb_plperl" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "mongo_fdw" >}} {{< ext "documentdb" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} |

> [!Note] Release tag 2.0.2 still ships extension SQL version 2.0; PGXN dist name is bson, CREATE EXTENSION name is pgbson, RPM package root is postgresbson, and the runtime dependency is libbson.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgbson` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.2` | {{< bg "18" "postgresbson_18" "green" >}} {{< bg "17" "postgresbson_17" "green" >}} {{< bg "16" "postgresbson_16" "green" >}} {{< bg "15" "postgresbson_15" "green" >}} {{< bg "14" "postgresbson_14" "green" >}} | `postgresbson_$v` | `libbson` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.2` | {{< bg "18" "postgresql-18-pgbson" "green" >}} {{< bg "17" "postgresql-17-pgbson" "green" >}} {{< bg "16" "postgresql-16-pgbson" "green" >}} {{< bg "15" "postgresql-15-pgbson" "green" >}} {{< bg "14" "postgresql-14-pgbson" "green" >}} | `postgresql-$v-pgbson` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.2" "postgresql-14-pgbson : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_18` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.3 KiB | [postgresbson_18-2.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_18-2.0.2-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_18` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.9 KiB | [postgresbson_18-2.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_18-2.0.2-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_18` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.7 KiB | [postgresbson_18-2.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_18-2.0.2-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_18` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.5 KiB | [postgresbson_18-2.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_18-2.0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_18` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.8 KiB | [postgresbson_18-2.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_18-2.0.2-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_18` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.6 KiB | [postgresbson_18-2.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_18-2.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgbson` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.6 KiB | [postgresql-18-pgbson_2.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgbson` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.3 KiB | [postgresql-18-pgbson_2.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgbson` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.7 KiB | [postgresql-18-pgbson_2.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgbson` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.3 KiB | [postgresql-18-pgbson_2.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgbson` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 39.3 KiB | [postgresql-18-pgbson_2.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgbson` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 38.9 KiB | [postgresql-18-pgbson_2.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgbson` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.6 KiB | [postgresql-18-pgbson_2.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgbson` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.6 KiB | [postgresql-18-pgbson_2.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgbson` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.7 KiB | [postgresql-18-pgbson_2.0.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgbson` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.2 KiB | [postgresql-18-pgbson_2.0.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-18-pgbson_2.0.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_17` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.3 KiB | [postgresbson_17-2.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_17-2.0.2-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_17` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.9 KiB | [postgresbson_17-2.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_17-2.0.2-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_17` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.7 KiB | [postgresbson_17-2.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_17-2.0.2-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_17` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.5 KiB | [postgresbson_17-2.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_17-2.0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_17` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.8 KiB | [postgresbson_17-2.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_17-2.0.2-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_17` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.6 KiB | [postgresbson_17-2.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_17-2.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgbson` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.6 KiB | [postgresql-17-pgbson_2.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgbson` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.3 KiB | [postgresql-17-pgbson_2.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgbson` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.8 KiB | [postgresql-17-pgbson_2.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgbson` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.3 KiB | [postgresql-17-pgbson_2.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgbson` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 40.3 KiB | [postgresql-17-pgbson_2.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgbson` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 39.9 KiB | [postgresql-17-pgbson_2.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgbson` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.7 KiB | [postgresql-17-pgbson_2.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgbson` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.6 KiB | [postgresql-17-pgbson_2.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgbson` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.7 KiB | [postgresql-17-pgbson_2.0.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgbson` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.4 KiB | [postgresql-17-pgbson_2.0.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-17-pgbson_2.0.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_16` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.3 KiB | [postgresbson_16-2.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_16-2.0.2-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_16` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.9 KiB | [postgresbson_16-2.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_16-2.0.2-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_16` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.7 KiB | [postgresbson_16-2.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_16-2.0.2-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_16` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.5 KiB | [postgresbson_16-2.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_16-2.0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_16` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.8 KiB | [postgresbson_16-2.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_16-2.0.2-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_16` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.6 KiB | [postgresbson_16-2.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_16-2.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgbson` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.6 KiB | [postgresql-16-pgbson_2.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgbson` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.3 KiB | [postgresql-16-pgbson_2.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgbson` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.7 KiB | [postgresql-16-pgbson_2.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgbson` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.3 KiB | [postgresql-16-pgbson_2.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgbson` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 40.3 KiB | [postgresql-16-pgbson_2.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgbson` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 39.9 KiB | [postgresql-16-pgbson_2.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgbson` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.8 KiB | [postgresql-16-pgbson_2.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgbson` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.6 KiB | [postgresql-16-pgbson_2.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgbson` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.7 KiB | [postgresql-16-pgbson_2.0.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgbson` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.4 KiB | [postgresql-16-pgbson_2.0.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-16-pgbson_2.0.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_15` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.3 KiB | [postgresbson_15-2.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_15-2.0.2-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_15` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.9 KiB | [postgresbson_15-2.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_15-2.0.2-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_15` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [postgresbson_15-2.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_15-2.0.2-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_15` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.5 KiB | [postgresbson_15-2.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_15-2.0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_15` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.8 KiB | [postgresbson_15-2.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_15-2.0.2-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_15` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.6 KiB | [postgresbson_15-2.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_15-2.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgbson` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.6 KiB | [postgresql-15-pgbson_2.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgbson` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.2 KiB | [postgresql-15-pgbson_2.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgbson` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.7 KiB | [postgresql-15-pgbson_2.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgbson` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.2 KiB | [postgresql-15-pgbson_2.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgbson` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 40.3 KiB | [postgresql-15-pgbson_2.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgbson` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 39.8 KiB | [postgresql-15-pgbson_2.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgbson` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.7 KiB | [postgresql-15-pgbson_2.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgbson` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.6 KiB | [postgresql-15-pgbson_2.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgbson` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.8 KiB | [postgresql-15-pgbson_2.0.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgbson` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.5 KiB | [postgresql-15-pgbson_2.0.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-15-pgbson_2.0.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_14` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.2 KiB | [postgresbson_14-2.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_14-2.0.2-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_14` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.9 KiB | [postgresbson_14-2.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_14-2.0.2-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_14` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.8 KiB | [postgresbson_14-2.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_14-2.0.2-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_14` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.5 KiB | [postgresbson_14-2.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_14-2.0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_14` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.8 KiB | [postgresbson_14-2.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_14-2.0.2-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_14` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.6 KiB | [postgresbson_14-2.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_14-2.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgbson` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 37.5 KiB | [postgresql-14-pgbson_2.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgbson` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 37.2 KiB | [postgresql-14-pgbson_2.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgbson` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 37.6 KiB | [postgresql-14-pgbson_2.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgbson` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 37.2 KiB | [postgresql-14-pgbson_2.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgbson` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 40.3 KiB | [postgresql-14-pgbson_2.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgbson` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 39.8 KiB | [postgresql-14-pgbson_2.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgbson` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.6 KiB | [postgresql-14-pgbson_2.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgbson` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 38.6 KiB | [postgresql-14-pgbson_2.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgbson` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.8 KiB | [postgresql-14-pgbson_2.0.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgbson` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.3 KiB | [postgresql-14-pgbson_2.0.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-14-pgbson_2.0.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/buzzm/postgresbson" title="Repository" icon="github" subtitle="github.com/buzzm/postgresbson" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresbson-2.0.2.tar.gz" >}}
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

Sources: [README](https://github.com/buzzm/postgresbson/blob/master/README.md), [META.json 2.0.2](https://github.com/buzzm/postgresbson/blob/master/META.json), [pgbson.control](https://github.com/buzzm/postgresbson/blob/master/pgbson.control)

`pgbson` adds a BSON data type plus BSON-aware accessors and operators. Upstream documents the package release as `2.0.2`, while the extension control file still exposes SQL default version `2.0`; this matches the packaging note that the dist version is ahead of the extension SQL version.

```sql
CREATE EXTENSION pgbson;
```

### Core Access Patterns

Typed dotpath accessors walk the BSON structure directly and are the upstream-recommended fast path:

```sql
SELECT bson_get_datetime(bson_column, 'msg.header.event.ts') FROM my_table;
SELECT bson_get_bson(bson_column, 'msg.header.event') FROM my_table;
SELECT bson_get_string(bson_column, 'data.payload.product.definition.id') FROM my_table;
```

JSON-style operators are also supported:

```sql
SELECT (bson_column->'msg'->'header'->'event'->>'ts')::timestamp
FROM my_table;
```

### Main Functions and Operators

- Typed getters such as `bson_get_string`, `bson_get_int32`, `bson_get_int64`, `bson_get_double`, `bson_get_decimal`, `bson_get_datetime`, `bson_get_binary`, and `bson_get_boolean`.
- `bson_get_bson` to return a BSON subdocument.
- `bson_get_jsonb_array` when a path resolves to an array and you want native `jsonb` array operators afterward.
- Arrow operators `->` and `->>` similar to PostgreSQL JSON types.
- Casts to `json`/`jsonb` using Extended JSON so type fidelity is preserved.

### Interop and Indexing

Cast BSON to `jsonb` when you want PostgreSQL JSON operators:

```sql
SELECT (bson_get_bson(bson_column, 'msg.header.event')::jsonb) ?& ARRAY['id', 'type']
FROM my_table;
```

Build expression indexes on extracted paths:

```sql
CREATE INDEX ON data_collection (bson_get_string(data, 'd.recordId'));
```

The README also notes BSON values can round-trip byte-for-byte through `bytea` casts.

### Caveats

- Dotpath accessors are usually faster and more memory-efficient than long `->` chains because they avoid materializing intermediate substructures.
- `bson_get_bson()` returns `NULL` for scalar endpoints because simple scalars are not BSON documents.
- Upstream explicitly calls out array handling and wrong-type accessor behavior as rough edges that still need better ergonomics.
