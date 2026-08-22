---
title: "pgbson"
linkTitle: "pgbson"
description: "BSON data type and accessor functions for PostgreSQL"
weight: 3910
categories: ["TYPE"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pgbson**](https://github.com/buzzm/postgresbson) : BSON data type and accessor functions for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3910** | {{< badge content="pgbson" link="https://github.com/buzzm/postgresbson" >}} | {{< ext "pgbson" >}} | `2.1.0` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgjq" >}} {{< ext "jsquery" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsonschema" >}} {{< ext "pg_projection" >}} {{< ext "hstore" >}} {{< ext "jsonb_plperl" >}} {{< ext "documentdb" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "jsonb_plperlu" >}} |

> [!Note] PGXN distribution name is bson, CREATE EXTENSION name is pgbson, source archive and RPM root are postgresbson, and the control default_version is 2.1 while the package release is 2.1.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgbson` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.0` | {{< bg "18" "postgresbson_18" "green" >}} {{< bg "17" "postgresbson_17" "green" >}} {{< bg "16" "postgresbson_16" "green" >}} {{< bg "15" "postgresbson_15" "green" >}} {{< bg "14" "postgresbson_14" "green" >}} | `postgresbson_$v` | `libbson` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.0` | {{< bg "18" "postgresql-18-pgbson" "green" >}} {{< bg "17" "postgresql-17-pgbson" "green" >}} {{< bg "16" "postgresql-16-pgbson" "green" >}} {{< bg "15" "postgresql-15-pgbson" "green" >}} {{< bg "14" "postgresql-14-pgbson" "green" >}} | `postgresql-$v-pgbson` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresbson_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_18` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.2 KiB | [postgresbson_18-2.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_18-2.1.0-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_18` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.6 KiB | [postgresbson_18-2.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_18-2.1.0-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_18` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.4 KiB | [postgresbson_18-2.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_18-2.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_18` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 34.1 KiB | [postgresbson_18-2.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_18-2.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_18` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.4 KiB | [postgresbson_18-2.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_18-2.1.0-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_18` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.2 KiB | [postgresbson_18-2.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_18-2.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgbson` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 43.2 KiB | [postgresql-18-pgbson_2.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-18-pgbson_2.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgbson` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.9 KiB | [postgresql-18-pgbson_2.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-18-pgbson_2.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgbson` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 43.2 KiB | [postgresql-18-pgbson_2.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-18-pgbson_2.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgbson` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 42.8 KiB | [postgresql-18-pgbson_2.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-18-pgbson_2.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgbson` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 45.2 KiB | [postgresql-18-pgbson_2.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-18-pgbson_2.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgbson` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 44.8 KiB | [postgresql-18-pgbson_2.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-18-pgbson_2.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgbson` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 44.4 KiB | [postgresql-18-pgbson_2.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-18-pgbson_2.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgbson` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 44.2 KiB | [postgresql-18-pgbson_2.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-18-pgbson_2.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgbson` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 44.2 KiB | [postgresql-18-pgbson_2.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-18-pgbson_2.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgbson` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 44.3 KiB | [postgresql-18-pgbson_2.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-18-pgbson_2.1.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_17` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.2 KiB | [postgresbson_17-2.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_17-2.1.0-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_17` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.6 KiB | [postgresbson_17-2.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_17-2.1.0-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_17` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.4 KiB | [postgresbson_17-2.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_17-2.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_17` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 34.1 KiB | [postgresbson_17-2.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_17-2.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_17` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.3 KiB | [postgresbson_17-2.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_17-2.1.0-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_17` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.2 KiB | [postgresbson_17-2.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_17-2.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgbson` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 43.3 KiB | [postgresql-17-pgbson_2.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-17-pgbson_2.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgbson` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.8 KiB | [postgresql-17-pgbson_2.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-17-pgbson_2.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgbson` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 43.3 KiB | [postgresql-17-pgbson_2.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-17-pgbson_2.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgbson` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 42.9 KiB | [postgresql-17-pgbson_2.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-17-pgbson_2.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgbson` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 46.4 KiB | [postgresql-17-pgbson_2.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-17-pgbson_2.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgbson` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.9 KiB | [postgresql-17-pgbson_2.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-17-pgbson_2.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgbson` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 44.5 KiB | [postgresql-17-pgbson_2.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-17-pgbson_2.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgbson` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 44.2 KiB | [postgresql-17-pgbson_2.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-17-pgbson_2.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgbson` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 44.3 KiB | [postgresql-17-pgbson_2.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-17-pgbson_2.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgbson` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 44.3 KiB | [postgresql-17-pgbson_2.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-17-pgbson_2.1.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_16` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.2 KiB | [postgresbson_16-2.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_16-2.1.0-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_16` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.6 KiB | [postgresbson_16-2.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_16-2.1.0-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_16` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.4 KiB | [postgresbson_16-2.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_16-2.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_16` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 34.1 KiB | [postgresbson_16-2.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_16-2.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_16` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.3 KiB | [postgresbson_16-2.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_16-2.1.0-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_16` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.2 KiB | [postgresbson_16-2.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_16-2.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgbson` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 43.3 KiB | [postgresql-16-pgbson_2.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-16-pgbson_2.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgbson` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.8 KiB | [postgresql-16-pgbson_2.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-16-pgbson_2.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgbson` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 43.3 KiB | [postgresql-16-pgbson_2.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-16-pgbson_2.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgbson` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 42.9 KiB | [postgresql-16-pgbson_2.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-16-pgbson_2.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgbson` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 46.4 KiB | [postgresql-16-pgbson_2.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-16-pgbson_2.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgbson` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.9 KiB | [postgresql-16-pgbson_2.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-16-pgbson_2.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgbson` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 44.5 KiB | [postgresql-16-pgbson_2.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-16-pgbson_2.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgbson` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 44.3 KiB | [postgresql-16-pgbson_2.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-16-pgbson_2.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgbson` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 44.3 KiB | [postgresql-16-pgbson_2.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-16-pgbson_2.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgbson` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 44.3 KiB | [postgresql-16-pgbson_2.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-16-pgbson_2.1.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_15` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.1 KiB | [postgresbson_15-2.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_15-2.1.0-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_15` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.7 KiB | [postgresbson_15-2.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_15-2.1.0-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_15` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.5 KiB | [postgresbson_15-2.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_15-2.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_15` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 34.1 KiB | [postgresbson_15-2.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_15-2.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_15` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.3 KiB | [postgresbson_15-2.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_15-2.1.0-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_15` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.2 KiB | [postgresbson_15-2.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_15-2.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgbson` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 43.3 KiB | [postgresql-15-pgbson_2.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-15-pgbson_2.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgbson` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.8 KiB | [postgresql-15-pgbson_2.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-15-pgbson_2.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgbson` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 43.2 KiB | [postgresql-15-pgbson_2.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-15-pgbson_2.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgbson` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 42.8 KiB | [postgresql-15-pgbson_2.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-15-pgbson_2.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgbson` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 46.4 KiB | [postgresql-15-pgbson_2.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-15-pgbson_2.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgbson` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.9 KiB | [postgresql-15-pgbson_2.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-15-pgbson_2.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgbson` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 44.5 KiB | [postgresql-15-pgbson_2.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-15-pgbson_2.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgbson` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 44.2 KiB | [postgresql-15-pgbson_2.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-15-pgbson_2.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgbson` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 44.4 KiB | [postgresql-15-pgbson_2.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-15-pgbson_2.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgbson` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 44.3 KiB | [postgresql-15-pgbson_2.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-15-pgbson_2.1.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_14` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.1 KiB | [postgresbson_14-2.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_14-2.1.0-1PIGSTY.el8.x86_64.rpm) |
| `postgresbson_14` | `2.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.7 KiB | [postgresbson_14-2.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_14-2.1.0-1PIGSTY.el8.aarch64.rpm) |
| `postgresbson_14` | `2.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 34.5 KiB | [postgresbson_14-2.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_14-2.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresbson_14` | `2.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 34.1 KiB | [postgresbson_14-2.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_14-2.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresbson_14` | `2.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 34.3 KiB | [postgresbson_14-2.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_14-2.1.0-1PIGSTY.el10.x86_64.rpm) |
| `postgresbson_14` | `2.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.2 KiB | [postgresbson_14-2.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_14-2.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgbson` | `2.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 43.3 KiB | [postgresql-14-pgbson_2.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-14-pgbson_2.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgbson` | `2.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.8 KiB | [postgresql-14-pgbson_2.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-14-pgbson_2.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgbson` | `2.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 43.2 KiB | [postgresql-14-pgbson_2.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-14-pgbson_2.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgbson` | `2.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 42.7 KiB | [postgresql-14-pgbson_2.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-14-pgbson_2.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgbson` | `2.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 46.3 KiB | [postgresql-14-pgbson_2.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-14-pgbson_2.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgbson` | `2.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.8 KiB | [postgresql-14-pgbson_2.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-14-pgbson_2.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgbson` | `2.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 44.5 KiB | [postgresql-14-pgbson_2.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-14-pgbson_2.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgbson` | `2.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 44.1 KiB | [postgresql-14-pgbson_2.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-14-pgbson_2.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgbson` | `2.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 44.4 KiB | [postgresql-14-pgbson_2.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-14-pgbson_2.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgbson` | `2.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 44.3 KiB | [postgresql-14-pgbson_2.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-14-pgbson_2.1.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/buzzm/postgresbson" title="Repository" icon="github" subtitle="github.com/buzzm/postgresbson" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresbson-2.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgbson;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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

- [pgbson 2.1.0 README](https://api.pgxn.org/src/bson/bson-2.1.0/README.md)
- [pgbson 2.1 control file](https://api.pgxn.org/src/bson/bson-2.1.0/pgbson.control)
- [pgbson 2.1 SQL API](https://api.pgxn.org/src/bson/bson-2.1.0/pgbson--2.1.sql)

`pgbson` adds a BSON data type, typed dot-path accessors, JSON-style navigation, casts, comparison operators, and btree/hash indexing. The PGXN distribution release is `2.1.0`, while the SQL extension version is `2.1`. Use BSON when binary round-trip fidelity or BSON-specific scalar types matter; use `jsonb` when PostgreSQL-native JSON indexing is the primary requirement.

### Install and Store BSON

```sql
CREATE EXTENSION pgbson;
SELECT pgbson_version();

CREATE TABLE events (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  payload bson NOT NULL
);

INSERT INTO events (payload)
VALUES ('{"user":{"name":"Ada"},"attempt":3}'::jsonb::bson);
```

The native module depends on `libbson`. The implicit `bytea`-to-`bson` cast validates BSON input, while the reverse cast preserves the binary representation.

### Extract Values

Typed accessors avoid materializing each intermediate document:

```sql
SELECT bson_get_string(payload, 'user.name'),
       bson_get_int32(payload, 'attempt')
FROM events;
```

Other typed getters cover 64-bit integers, doubles, decimals, datetimes, binary values, booleans, embedded BSON documents, and JSONB arrays. A missing path or a type mismatch returns `NULL`, so validate the expected BSON schema at ingestion when those cases must be distinguished.

Version 2.1 adds a type-agnostic terminal extractor:

```sql
SELECT bson_get_value(payload, 'user.name')
FROM events;
-- { "_" : "Ada" }
```

`bson_get_value` always wraps the selected scalar, array, or document under the key `_`. Remove exactly that one wrapper in the caller. It intentionally has no chainable `->` equivalent.

### Navigate, Compare, and Index

```sql
SELECT payload->'user'->>'name'
FROM events;

CREATE INDEX events_user_name_idx
ON events (bson_get_string(payload, 'user.name'));

CREATE INDEX events_payload_btree_idx ON events (payload);
CREATE INDEX events_payload_hash_idx ON events USING hash (payload);
```

Version 2.1 provides logical comparison operators `=`, `<>`, `<`, `<=`, `>`, and `>=`; `==` and `<<>>` perform binary equality and inequality. The default btree operator class uses logical BSON comparison, while the hash operator class uses binary equality. Choose intentionally when field order or byte identity matters.

### Upgrade and Caveats

```sql
ALTER EXTENSION pgbson UPDATE TO '2.1';
```

- Installing a 2.1 shared library does not update an existing 2.0 extension's SQL objects; run the extension update after installing the files.
- The 2.1 shared library fixes a backend crash when `bson_get_bson()` or `->` resolves to a scalar endpoint. Earlier binaries should be replaced even when an application does not yet use the new 2.1 SQL function.
- BSON-to-JSON/JSONB casts use Extended JSON. BSON and JSONB have different type, equality, and ordering semantics, so conversion is not lossless for every workflow.
- In 2.1, `->>` on a BSON datetime includes the trailing `Z`; `bson_get_datetime()` is unchanged. Check clients that compare the old text form.
- BSON top-level values are documents, not bare arrays or scalars. `bson_get_value` uses its `_` wrapper to return any nested shape within that restriction.
