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
| **3910** | {{< badge content="pgbson" link="https://github.com/buzzm/postgresbson" >}} | {{< ext "pgbson" >}} | `2.2.0` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgjq" >}} {{< ext "jsquery" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsonschema" >}} {{< ext "pg_projection" >}} {{< ext "hstore" >}} {{< ext "jsonb_plperl" >}} {{< ext "documentdb" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "jsonb_plperlu" >}} |

> [!Note] PGXN distribution bson and SQL extension pgbson; package version 2.2.0 installs SQL version 2.2.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgbson` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "postgresbson_18" "green" >}} {{< bg "17" "postgresbson_17" "green" >}} {{< bg "16" "postgresbson_16" "green" >}} {{< bg "15" "postgresbson_15" "green" >}} {{< bg "14" "postgresbson_14" "green" >}} | `postgresbson_$v` | `libbson` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "postgresql-18-pgbson" "green" >}} {{< bg "17" "postgresql-17-pgbson" "green" >}} {{< bg "16" "postgresql-16-pgbson" "green" >}} {{< bg "15" "postgresql-15-pgbson" "green" >}} {{< bg "14" "postgresql-14-pgbson" "green" >}} | `postgresql-$v-pgbson` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_14 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_14 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_14 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresbson_14 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgbson : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgbson : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_18` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 54.6 KiB | [postgresbson_18-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_18-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `postgresbson_18` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.0 KiB | [postgresbson_18-2.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/postgresbson_18-2.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `postgresbson_18` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.0 KiB | [postgresbson_18-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_18-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `postgresbson_18` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.2 KiB | [postgresbson_18-2.2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgresbson_18-2.2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `postgresbson_18` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.0 KiB | [postgresbson_18-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_18-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `postgresbson_18` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.9 KiB | [postgresbson_18-2.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgresbson_18-2.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `postgresbson_18` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.4 KiB | [postgresbson_18-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_18-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `postgresbson_18` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.3 KiB | [postgresbson_18-2.2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgresbson_18-2.2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `postgresbson_18` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.2 KiB | [postgresbson_18-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_18-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `postgresbson_18` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.1 KiB | [postgresbson_18-2.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgresbson_18-2.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `postgresbson_18` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.6 KiB | [postgresbson_18-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_18-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresbson_18` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.5 KiB | [postgresbson_18-2.2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgresbson_18-2.2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-pgbson` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 46.0 KiB | [postgresql-18-pgbson_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-18-pgbson_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgbson` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 45.4 KiB | [postgresql-18-pgbson_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-18-pgbson_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgbson` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 46.1 KiB | [postgresql-18-pgbson_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-18-pgbson_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pgbson` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 45.5 KiB | [postgresql-18-pgbson_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-18-pgbson_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pgbson` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 48.2 KiB | [postgresql-18-pgbson_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-18-pgbson_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pgbson` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 47.5 KiB | [postgresql-18-pgbson_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-18-pgbson_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pgbson` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 47.6 KiB | [postgresql-18-pgbson_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-18-pgbson_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pgbson` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 47.1 KiB | [postgresql-18-pgbson_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-18-pgbson_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pgbson` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 47.4 KiB | [postgresql-18-pgbson_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-18-pgbson_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pgbson` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 47.4 KiB | [postgresql-18-pgbson_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-18-pgbson_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_17` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 54.7 KiB | [postgresbson_17-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_17-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `postgresbson_17` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.0 KiB | [postgresbson_17-2.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgresbson_17-2.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `postgresbson_17` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.0 KiB | [postgresbson_17-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_17-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `postgresbson_17` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.2 KiB | [postgresbson_17-2.2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgresbson_17-2.2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `postgresbson_17` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.0 KiB | [postgresbson_17-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_17-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `postgresbson_17` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.9 KiB | [postgresbson_17-2.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgresbson_17-2.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `postgresbson_17` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.4 KiB | [postgresbson_17-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_17-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `postgresbson_17` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.3 KiB | [postgresbson_17-2.2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgresbson_17-2.2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `postgresbson_17` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.3 KiB | [postgresbson_17-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_17-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `postgresbson_17` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.1 KiB | [postgresbson_17-2.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgresbson_17-2.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `postgresbson_17` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.6 KiB | [postgresbson_17-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_17-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresbson_17` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.5 KiB | [postgresbson_17-2.2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgresbson_17-2.2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-pgbson` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.9 KiB | [postgresql-17-pgbson_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-17-pgbson_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgbson` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 45.4 KiB | [postgresql-17-pgbson_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-17-pgbson_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgbson` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 46.0 KiB | [postgresql-17-pgbson_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-17-pgbson_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pgbson` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 45.4 KiB | [postgresql-17-pgbson_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-17-pgbson_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pgbson` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 49.2 KiB | [postgresql-17-pgbson_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-17-pgbson_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pgbson` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.6 KiB | [postgresql-17-pgbson_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-17-pgbson_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pgbson` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 47.6 KiB | [postgresql-17-pgbson_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-17-pgbson_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pgbson` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 47.0 KiB | [postgresql-17-pgbson_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-17-pgbson_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pgbson` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 47.4 KiB | [postgresql-17-pgbson_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-17-pgbson_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pgbson` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 47.5 KiB | [postgresql-17-pgbson_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-17-pgbson_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_16` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 54.7 KiB | [postgresbson_16-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_16-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `postgresbson_16` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.0 KiB | [postgresbson_16-2.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgresbson_16-2.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `postgresbson_16` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.0 KiB | [postgresbson_16-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_16-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `postgresbson_16` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.2 KiB | [postgresbson_16-2.2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgresbson_16-2.2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `postgresbson_16` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.0 KiB | [postgresbson_16-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_16-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `postgresbson_16` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.8 KiB | [postgresbson_16-2.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgresbson_16-2.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `postgresbson_16` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.4 KiB | [postgresbson_16-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_16-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `postgresbson_16` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.3 KiB | [postgresbson_16-2.2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgresbson_16-2.2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `postgresbson_16` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.3 KiB | [postgresbson_16-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_16-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `postgresbson_16` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.1 KiB | [postgresbson_16-2.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgresbson_16-2.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `postgresbson_16` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.6 KiB | [postgresbson_16-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_16-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresbson_16` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.5 KiB | [postgresbson_16-2.2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgresbson_16-2.2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-pgbson` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.9 KiB | [postgresql-16-pgbson_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-16-pgbson_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgbson` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 45.4 KiB | [postgresql-16-pgbson_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-16-pgbson_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgbson` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 46.0 KiB | [postgresql-16-pgbson_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-16-pgbson_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pgbson` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 45.4 KiB | [postgresql-16-pgbson_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-16-pgbson_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pgbson` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 49.2 KiB | [postgresql-16-pgbson_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-16-pgbson_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pgbson` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.6 KiB | [postgresql-16-pgbson_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-16-pgbson_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pgbson` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 47.6 KiB | [postgresql-16-pgbson_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-16-pgbson_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pgbson` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.9 KiB | [postgresql-16-pgbson_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-16-pgbson_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pgbson` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 47.4 KiB | [postgresql-16-pgbson_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-16-pgbson_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pgbson` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 47.5 KiB | [postgresql-16-pgbson_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-16-pgbson_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_15` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 54.6 KiB | [postgresbson_15-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_15-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `postgresbson_15` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.0 KiB | [postgresbson_15-2.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgresbson_15-2.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `postgresbson_15` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 53.9 KiB | [postgresbson_15-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_15-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `postgresbson_15` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.1 KiB | [postgresbson_15-2.2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgresbson_15-2.2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `postgresbson_15` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.1 KiB | [postgresbson_15-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_15-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `postgresbson_15` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.9 KiB | [postgresbson_15-2.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgresbson_15-2.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `postgresbson_15` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.4 KiB | [postgresbson_15-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_15-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `postgresbson_15` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.3 KiB | [postgresbson_15-2.2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgresbson_15-2.2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `postgresbson_15` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.2 KiB | [postgresbson_15-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_15-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `postgresbson_15` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.0 KiB | [postgresbson_15-2.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgresbson_15-2.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `postgresbson_15` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.6 KiB | [postgresbson_15-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_15-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresbson_15` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.5 KiB | [postgresbson_15-2.2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgresbson_15-2.2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-pgbson` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.9 KiB | [postgresql-15-pgbson_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-15-pgbson_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgbson` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.0 KiB | [postgresql-15-pgbson_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-15-pgbson_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgbson` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.9 KiB | [postgresql-15-pgbson_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-15-pgbson_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pgbson` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 45.4 KiB | [postgresql-15-pgbson_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-15-pgbson_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pgbson` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 49.2 KiB | [postgresql-15-pgbson_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-15-pgbson_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pgbson` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.6 KiB | [postgresql-15-pgbson_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-15-pgbson_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pgbson` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 47.5 KiB | [postgresql-15-pgbson_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-15-pgbson_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pgbson` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.9 KiB | [postgresql-15-pgbson_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-15-pgbson_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pgbson` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 47.3 KiB | [postgresql-15-pgbson_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-15-pgbson_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pgbson` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 47.4 KiB | [postgresql-15-pgbson_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-15-pgbson_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresbson_14` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 54.5 KiB | [postgresbson_14-2.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/postgresbson_14-2.2.0-1PGSTY.el8.x86_64.rpm) |
| `postgresbson_14` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.0 KiB | [postgresbson_14-2.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgresbson_14-2.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `postgresbson_14` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 53.8 KiB | [postgresbson_14-2.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/postgresbson_14-2.2.0-1PGSTY.el8.aarch64.rpm) |
| `postgresbson_14` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.1 KiB | [postgresbson_14-2.2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgresbson_14-2.2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `postgresbson_14` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.9 KiB | [postgresbson_14-2.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/postgresbson_14-2.2.0-1PGSTY.el9.x86_64.rpm) |
| `postgresbson_14` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.9 KiB | [postgresbson_14-2.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresbson_14-2.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `postgresbson_14` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.3 KiB | [postgresbson_14-2.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/postgresbson_14-2.2.0-1PGSTY.el9.aarch64.rpm) |
| `postgresbson_14` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 36.2 KiB | [postgresbson_14-2.2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgresbson_14-2.2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `postgresbson_14` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.0 KiB | [postgresbson_14-2.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/postgresbson_14-2.2.0-1PGSTY.el10.x86_64.rpm) |
| `postgresbson_14` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 37.0 KiB | [postgresbson_14-2.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgresbson_14-2.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `postgresbson_14` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.5 KiB | [postgresbson_14-2.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/postgresbson_14-2.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresbson_14` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.5 KiB | [postgresbson_14-2.2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgresbson_14-2.2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-pgbson` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.8 KiB | [postgresql-14-pgbson_2.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-14-pgbson_2.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgbson` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 45.3 KiB | [postgresql-14-pgbson_2.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgresbson/postgresql-14-pgbson_2.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgbson` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.8 KiB | [postgresql-14-pgbson_2.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-14-pgbson_2.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pgbson` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 45.3 KiB | [postgresql-14-pgbson_2.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/postgresbson/postgresql-14-pgbson_2.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pgbson` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 49.1 KiB | [postgresql-14-pgbson_2.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-14-pgbson_2.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pgbson` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.4 KiB | [postgresql-14-pgbson_2.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgresbson/postgresql-14-pgbson_2.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pgbson` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 47.4 KiB | [postgresql-14-pgbson_2.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-14-pgbson_2.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pgbson` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.9 KiB | [postgresql-14-pgbson_2.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgresbson/postgresql-14-pgbson_2.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pgbson` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 47.2 KiB | [postgresql-14-pgbson_2.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-14-pgbson_2.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pgbson` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 47.3 KiB | [postgresql-14-pgbson_2.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/postgresbson/postgresql-14-pgbson_2.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/buzzm/postgresbson" title="Repository" icon="github" subtitle="github.com/buzzm/postgresbson" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresbson-2.2.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgbson;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgbson;		# install via package name, for the active PG version

pig install pgbson -v 18;   # install for PG 18
pig install pgbson -v 17;   # install for PG 17
pig install pgbson -v 16;   # install for PG 16
pig install pgbson -v 15;   # install for PG 15
pig install pgbson -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pgbson;
```

## Usage

Sources:

- [pgbson 2.2.0 README](https://api.pgxn.org/src/bson/bson-2.2.0/README.md)
- [pgbson 2.2 control file](https://api.pgxn.org/src/bson/bson-2.2.0/pgbson.control)
- [pgbson 2.2 SQL API](https://api.pgxn.org/src/bson/bson-2.2.0/pgbson--2.2.sql)

`pgbson` adds a BSON data type, typed dot-path accessors, JSON-style navigation, casts, comparison operators, and btree/hash indexing. The PGXN distribution release is `2.2.0`, while the SQL extension version is `2.2`. Use BSON when binary round-trip fidelity or BSON-specific scalar types matter; use `jsonb` when PostgreSQL-native JSON indexing is the primary requirement.

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
ALTER EXTENSION pgbson UPDATE TO '2.2';
```

- Installing a 2.1 shared library does not update an existing 2.0 extension's SQL objects; run the extension update after installing the files.
- The 2.1 shared library fixes a backend crash when `bson_get_bson()` or `->` resolves to a scalar endpoint. Earlier binaries should be replaced even when an application does not yet use the new 2.1 SQL function.
- BSON-to-JSON/JSONB casts use Extended JSON. BSON and JSONB have different type, equality, and ordering semantics, so conversion is not lossless for every workflow.
- In 2.1, `->>` on a BSON datetime includes the trailing `Z`; `bson_get_datetime()` is unchanged. Check clients that compare the old text form.
- BSON top-level values are documents, not bare arrays or scalars. `bson_get_value` uses its `_` wrapper to return any nested shape within that restriction.

### Binary Fixes in 2.1.1

Distribution 2.1.1 retains SQL version 2.1: install the new library and reconnect, with no SQL extension update for an existing 2.1 installation. It fixes the BSON binary receive function reading the wrong memory and crashing a backend, affecting binary COPY and drivers sending binary parameters; text input was unaffected. `bson_get_datetime` now returns NULL for dates outside PostgreSQL timestamp range. Truncation at embedded NUL in `bson_get_string` and `->>` is newly documented, not new behavior.

### Current Upgrade Boundary in 2.2.0

Distribution 2.2.0 installs SQL version 2.2. Comparison functions now use C implementations, removing comparison hijacking through same-name functions on `search_path`. Install the new library before applying the SQL update above; existing indexes do not need rebuilding for this change. The release includes 2.1.2 fixes for libbson compatibility, encoding conversion and invalid input handling. Invalid UTF-8 validation still depends on the libbson version: old key validation does not establish full string-value validation. Use matching library and SQL files; the earlier 2.1.1 no-SQL-update note does not apply to this upgrade.
