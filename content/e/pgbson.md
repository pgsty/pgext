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

> [!Note] PGXN dist name is bson, but CREATE EXTENSION name is pgbson; RPM package root is postgresbson; RPM runtime dependency is libbson.


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


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

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

> Syntax:
>
> ```sql
> CREATE EXTENSION pgbson;
> SELECT bson_get_datetime(bson_column, 'msg.header.event.ts') FROM my_table;
> SELECT (bson_column->'msg'->'header'->'event'->>'ts')::timestamp FROM my_table;
> ```
>
> Source: [README](https://github.com/buzzm/postgresbson)

`pgbson` adds a BSON data type to PostgreSQL together with functions and operators for creating, inspecting, and querying BSON documents. The upstream README positions it as a binary, richly typed alternative to JSON/JSONB with round-trip fidelity and first-class support for datetimes, numeric subtypes, and raw bytes.

## Why BSON

The README highlights several BSON advantages over JSON:

- datetimes are first-class values
- numeric types remain distinct (`int32`, `int64`, `float`, `decimal`)
- raw byte arrays are first-class
- round-tripping preserves exact binary representation
- native SDK support exists across many languages

## Access Patterns

The extension exposes two styles of access:

### Dotpath Accessors

These are the high-performance typed accessors documented upstream:

```sql
SELECT bson_get_datetime(bson_column, 'msg.header.event.ts') FROM my_table;
SELECT bson_get_bson(bson_column, 'msg.header.event') FROM my_table;
```

The README argues these are more memory-efficient than repeated arrow dereferences because they walk the BSON structure directly and materialize only the terminal value.

### Arrow Operators

It also supports JSON-like operators:

```sql
SELECT (bson_column->'msg'->'header'->'event'->>'ts')::timestamp
FROM my_table;
```

## JSON Interop

The BSON type can be cast to JSON using Extended JSON (EJSON) so type fidelity is preserved. This allows BSON values to be fed into JSON/JSONB functions and operators when needed:

```sql
SELECT (bson_get_bson(bson_column, 'msg.header.event')::jsonb) ?& ARRAY['id','type']
FROM my_table;
```

## Notes

The README includes examples of end-to-end BSON round-tripping across Java, Kafka, Python, and PostgreSQL, emphasizing that the stored BSON payload can be retrieved byte-for-byte unchanged when cast back to `bytea`.
