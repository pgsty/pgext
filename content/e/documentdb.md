---
title: "documentdb"
linkTitle: "documentdb"
description: "API surface for DocumentDB for PostgreSQL"
weight: 9000
categories: ["SIM"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**documentdb**](https://github.com/documentdb/documentdb) : API surface for DocumentDB for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9000** | {{< badge content="documentdb" link="https://github.com/documentdb/documentdb" >}} | {{< ext "documentdb" >}} | `1.2` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `documentdb_api` `documentdb_api_internal` `documentdb_api_catalog` `documentdb_data` `documentdb_api_v2` `documentdb_api_internal_admin` `documentdb_api_internal_readwrite` `documentdb_api_internal_readonly` `documentdb_api_internal_bgworker` |
|   **Requires**    | {{< ext "documentdb_core" >}} {{< ext "pg_cron" >}} {{< ext "postgis" >}} {{< ext "tsm_system_rows" >}} {{< ext "vector" >}} |
|    **Need By**    | {{< ext "documentdb_extended_rum" >}} |
|   **See Also**    | {{< ext "pgbson" >}} {{< ext "jsquery" >}} {{< ext "pg_projection" >}} {{< ext "mongo_fdw" >}} {{< ext "pgjq" >}} {{< ext "pg_graphql" >}} {{< ext "omni_rest" >}} {{< ext "pg_jsonschema" >}} {{< ext "pg_net" >}} {{< ext "jsonschema" >}} |
|    **Siblings**   | {{< ext "documentdb_core" >}} {{< ext "documentdb_extended_rum" >}} |

> [!Note] SQL 1.0-0; RPM external RUM dependency only on PG15-17.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `documentdb` | `documentdb_core`, `pg_cron`, `postgis`, `tsm_system_rows`, `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2` | {{< bg "19" "documentdb_19" "green" >}} {{< bg "18" "documentdb_18" "green" >}} {{< bg "17" "documentdb_17" "green" >}} {{< bg "16" "documentdb_16" "green" >}} {{< bg "15" "documentdb_15" "green" >}} | `documentdb_$v` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v`, `postgis37_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2` | {{< bg "19" "postgresql-19-documentdb" "green" >}} {{< bg "18" "postgresql-18-documentdb" "green" >}} {{< bg "17" "postgresql-17-documentdb" "green" >}} {{< bg "16" "postgresql-16-documentdb" "green" >}} {{< bg "15" "postgresql-15-documentdb" "green" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum`, `postgresql-$v-postgis-3` |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "documentdb_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.0" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_15 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "documentdb_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.0" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_15 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "MISS" "documentdb_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_15 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "MISS" "documentdb_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_15 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "MISS" "documentdb_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_15 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "MISS" "documentdb_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_15 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "MISS" "postgresql-19-documentdb : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "postgresql-15-documentdb : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "MISS" "postgresql-19-documentdb : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "postgresql-15-documentdb : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "MISS" "postgresql-19-documentdb : MISS 0" "red" >}} | {{< bg "PGDG 1.2" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-15-documentdb : AVAIL 4" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "MISS" "postgresql-19-documentdb : MISS 0" "red" >}} | {{< bg "PGDG 1.2" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-15-documentdb : AVAIL 4" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "MISS" "postgresql-19-documentdb : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "postgresql-15-documentdb : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "MISS" "postgresql-19-documentdb : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "postgresql-15-documentdb : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "MISS" "postgresql-19-documentdb : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "postgresql-15-documentdb : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "MISS" "postgresql-19-documentdb : MISS 0" "red" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "postgresql-15-documentdb : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "MISS" "postgresql-19-documentdb : MISS 0" "red" >}} | {{< bg "PGDG 1.2" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-15-documentdb : AVAIL 4" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "MISS" "postgresql-19-documentdb : MISS 0" "red" >}} | {{< bg "PGDG 1.2" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-15-documentdb : AVAIL 4" "blue" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.0 MiB | [documentdb_18-1.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_18-1.0-2PGSTY.el8.x86_64.rpm) |
| `documentdb_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 5.8 MiB | [documentdb_18-1.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_18-1.0-2PGSTY.el8.aarch64.rpm) |
| `documentdb_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.9 MiB | [documentdb_18-1.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_18-1.0-2PGSTY.el9.x86_64.rpm) |
| `documentdb_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 5.7 MiB | [documentdb_18-1.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_18-1.0-2PGSTY.el9.aarch64.rpm) |
| `documentdb_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.9 MiB | [documentdb_18-1.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_18-1.0-2PGSTY.el10.x86_64.rpm) |
| `documentdb_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 5.8 MiB | [documentdb_18-1.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_18-1.0-2PGSTY.el10.aarch64.rpm) |
| `postgresql-18-documentdb` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.6 MiB | [postgresql-18-documentdb_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-documentdb` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.4 MiB | [postgresql-18-documentdb_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-documentdb` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.7 MiB | [postgresql-18-documentdb_1.2-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_1.2-0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-documentdb` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.5 MiB | [postgresql-18-documentdb_1.0-1-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1-1.pgdg13+2_amd64.deb) |
| `postgresql-18-documentdb` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.5 MiB | [postgresql-18-documentdb_1.0-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_1.0-0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-documentdb` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.6 MiB | [postgresql-18-documentdb_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-documentdb` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.5 MiB | [postgresql-18-documentdb_1.2-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_1.2-0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-documentdb` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.3 MiB | [postgresql-18-documentdb_1.0-1-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1-1.pgdg13+2_arm64.deb) |
| `postgresql-18-documentdb` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.3 MiB | [postgresql-18-documentdb_1.0-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_1.0-0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-documentdb` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.4 MiB | [postgresql-18-documentdb_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-documentdb` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.9 MiB | [postgresql-18-documentdb_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-documentdb` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.8 MiB | [postgresql-18-documentdb_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-documentdb` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.8 MiB | [postgresql-18-documentdb_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-documentdb` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.7 MiB | [postgresql-18-documentdb_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-documentdb` | `1.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.6 MiB | [postgresql-18-documentdb_1.2-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_1.2-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-documentdb` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.4 MiB | [postgresql-18-documentdb_1.0-1-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1-1.pgdg26.04+2_amd64.deb) |
| `postgresql-18-documentdb` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.4 MiB | [postgresql-18-documentdb_1.0-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_1.0-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-documentdb` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.8 MiB | [postgresql-18-documentdb_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-documentdb` | `1.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.4 MiB | [postgresql-18-documentdb_1.2-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_1.2-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-documentdb` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.2 MiB | [postgresql-18-documentdb_1.0-1-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1-1.pgdg26.04+2_arm64.deb) |
| `postgresql-18-documentdb` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.2 MiB | [postgresql-18-documentdb_1.0-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_1.0-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-documentdb` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-18-documentdb_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-18-documentdb_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.0 MiB | [documentdb_17-1.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_17-1.0-2PGSTY.el8.x86_64.rpm) |
| `documentdb_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 5.8 MiB | [documentdb_17-1.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_17-1.0-2PGSTY.el8.aarch64.rpm) |
| `documentdb_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.9 MiB | [documentdb_17-1.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_17-1.0-2PGSTY.el9.x86_64.rpm) |
| `documentdb_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 5.7 MiB | [documentdb_17-1.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_17-1.0-2PGSTY.el9.aarch64.rpm) |
| `documentdb_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.9 MiB | [documentdb_17-1.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_17-1.0-2PGSTY.el10.x86_64.rpm) |
| `documentdb_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 5.8 MiB | [documentdb_17-1.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_17-1.0-2PGSTY.el10.aarch64.rpm) |
| `postgresql-17-documentdb` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.6 MiB | [postgresql-17-documentdb_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-documentdb` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.4 MiB | [postgresql-17-documentdb_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-documentdb` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.7 MiB | [postgresql-17-documentdb_1.2-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_1.2-0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-documentdb` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.5 MiB | [postgresql-17-documentdb_1.0-1-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1-1.pgdg13+2_amd64.deb) |
| `postgresql-17-documentdb` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.5 MiB | [postgresql-17-documentdb_1.0-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_1.0-0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-documentdb` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.6 MiB | [postgresql-17-documentdb_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-documentdb` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.5 MiB | [postgresql-17-documentdb_1.2-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_1.2-0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-documentdb` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.3 MiB | [postgresql-17-documentdb_1.0-1-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1-1.pgdg13+2_arm64.deb) |
| `postgresql-17-documentdb` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.3 MiB | [postgresql-17-documentdb_1.0-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_1.0-0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-documentdb` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.4 MiB | [postgresql-17-documentdb_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-documentdb` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.4 MiB | [postgresql-17-documentdb_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-documentdb` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.2 MiB | [postgresql-17-documentdb_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-documentdb` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.8 MiB | [postgresql-17-documentdb_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-documentdb` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.7 MiB | [postgresql-17-documentdb_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-documentdb` | `1.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.6 MiB | [postgresql-17-documentdb_1.2-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_1.2-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-documentdb` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.4 MiB | [postgresql-17-documentdb_1.0-1-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1-1.pgdg26.04+2_amd64.deb) |
| `postgresql-17-documentdb` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.4 MiB | [postgresql-17-documentdb_1.0-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_1.0-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-documentdb` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.8 MiB | [postgresql-17-documentdb_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-documentdb` | `1.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.4 MiB | [postgresql-17-documentdb_1.2-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_1.2-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-documentdb` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.2 MiB | [postgresql-17-documentdb_1.0-1-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1-1.pgdg26.04+2_arm64.deb) |
| `postgresql-17-documentdb` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.2 MiB | [postgresql-17-documentdb_1.0-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_1.0-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-documentdb` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-17-documentdb_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-17-documentdb_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.0 MiB | [documentdb_16-1.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_16-1.0-2PGSTY.el8.x86_64.rpm) |
| `documentdb_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 5.8 MiB | [documentdb_16-1.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_16-1.0-2PGSTY.el8.aarch64.rpm) |
| `documentdb_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 5.9 MiB | [documentdb_16-1.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_16-1.0-2PGSTY.el9.x86_64.rpm) |
| `documentdb_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 5.8 MiB | [documentdb_16-1.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_16-1.0-2PGSTY.el9.aarch64.rpm) |
| `documentdb_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 5.9 MiB | [documentdb_16-1.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_16-1.0-2PGSTY.el10.x86_64.rpm) |
| `documentdb_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 5.8 MiB | [documentdb_16-1.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_16-1.0-2PGSTY.el10.aarch64.rpm) |
| `postgresql-16-documentdb` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.6 MiB | [postgresql-16-documentdb_1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-documentdb` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.4 MiB | [postgresql-16-documentdb_1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-documentdb` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.7 MiB | [postgresql-16-documentdb_1.2-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_1.2-0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-documentdb` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.5 MiB | [postgresql-16-documentdb_1.0-1-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1-1.pgdg13+2_amd64.deb) |
| `postgresql-16-documentdb` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.5 MiB | [postgresql-16-documentdb_1.0-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_1.0-0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-documentdb` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.6 MiB | [postgresql-16-documentdb_1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-documentdb` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.5 MiB | [postgresql-16-documentdb_1.2-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_1.2-0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-documentdb` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.3 MiB | [postgresql-16-documentdb_1.0-1-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1-1.pgdg13+2_arm64.deb) |
| `postgresql-16-documentdb` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.3 MiB | [postgresql-16-documentdb_1.0-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_1.0-0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-documentdb` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.4 MiB | [postgresql-16-documentdb_1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-documentdb` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.4 MiB | [postgresql-16-documentdb_1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-documentdb` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.2 MiB | [postgresql-16-documentdb_1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-documentdb` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.8 MiB | [postgresql-16-documentdb_1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-documentdb` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.7 MiB | [postgresql-16-documentdb_1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-documentdb` | `1.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.6 MiB | [postgresql-16-documentdb_1.2-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_1.2-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-documentdb` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.4 MiB | [postgresql-16-documentdb_1.0-1-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1-1.pgdg26.04+2_amd64.deb) |
| `postgresql-16-documentdb` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.4 MiB | [postgresql-16-documentdb_1.0-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_1.0-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-documentdb` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.8 MiB | [postgresql-16-documentdb_1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-documentdb` | `1.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.4 MiB | [postgresql-16-documentdb_1.2-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_1.2-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-documentdb` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.2 MiB | [postgresql-16-documentdb_1.0-1-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1-1.pgdg26.04+2_arm64.deb) |
| `postgresql-16-documentdb` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.2 MiB | [postgresql-16-documentdb_1.0-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_1.0-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-documentdb` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-16-documentdb_1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-16-documentdb_1.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.1 MiB | [documentdb_15-1.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_15-1.0-2PGSTY.el8.x86_64.rpm) |
| `documentdb_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 5.8 MiB | [documentdb_15-1.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_15-1.0-2PGSTY.el8.aarch64.rpm) |
| `documentdb_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.0 MiB | [documentdb_15-1.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_15-1.0-2PGSTY.el9.x86_64.rpm) |
| `documentdb_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 5.8 MiB | [documentdb_15-1.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_15-1.0-2PGSTY.el9.aarch64.rpm) |
| `documentdb_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.0 MiB | [documentdb_15-1.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_15-1.0-2PGSTY.el10.x86_64.rpm) |
| `documentdb_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 5.8 MiB | [documentdb_15-1.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_15-1.0-2PGSTY.el10.aarch64.rpm) |
| `postgresql-15-documentdb` | `0.117` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.6 MiB | [postgresql-15-documentdb_0.117-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.117-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-documentdb` | `0.117` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.3 MiB | [postgresql-15-documentdb_0.117-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.117-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-documentdb` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.7 MiB | [postgresql-15-documentdb_1.2-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_1.2-0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-documentdb` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.6 MiB | [postgresql-15-documentdb_1.0-1-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_1.0-1-1.pgdg13+2_amd64.deb) |
| `postgresql-15-documentdb` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.6 MiB | [postgresql-15-documentdb_1.0-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_1.0-0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.117` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.6 MiB | [postgresql-15-documentdb_0.117-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-15-documentdb_0.117-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-documentdb` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.5 MiB | [postgresql-15-documentdb_1.2-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_1.2-0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-documentdb` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.4 MiB | [postgresql-15-documentdb_1.0-1-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_1.0-1-1.pgdg13+2_arm64.deb) |
| `postgresql-15-documentdb` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.4 MiB | [postgresql-15-documentdb_1.0-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_1.0-0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-documentdb` | `0.117` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.4 MiB | [postgresql-15-documentdb_0.117-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-15-documentdb_0.117-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-documentdb` | `0.117` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.4 MiB | [postgresql-15-documentdb_0.117-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.117-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-documentdb` | `0.117` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.2 MiB | [postgresql-15-documentdb_0.117-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.117-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-documentdb` | `0.117` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.8 MiB | [postgresql-15-documentdb_0.117-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.117-1PGSTY~noble_amd64.deb) |
| `postgresql-15-documentdb` | `0.117` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.7 MiB | [postgresql-15-documentdb_0.117-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.117-1PGSTY~noble_arm64.deb) |
| `postgresql-15-documentdb` | `1.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.6 MiB | [postgresql-15-documentdb_1.2-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_1.2-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-documentdb` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.5 MiB | [postgresql-15-documentdb_1.0-1-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_1.0-1-1.pgdg26.04+2_amd64.deb) |
| `postgresql-15-documentdb` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.5 MiB | [postgresql-15-documentdb_1.0-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_1.0-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.117` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.8 MiB | [postgresql-15-documentdb_0.117-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-15-documentdb_0.117-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-documentdb` | `1.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.4 MiB | [postgresql-15-documentdb_1.2-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_1.2-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-documentdb` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.3 MiB | [postgresql-15-documentdb_1.0-1-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_1.0-1-1.pgdg26.04+2_arm64.deb) |
| `postgresql-15-documentdb` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 5.3 MiB | [postgresql-15-documentdb_1.0-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_1.0-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-documentdb` | `0.117` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-15-documentdb_0.117-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-15-documentdb_0.117-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/documentdb/documentdb" title="Repository" icon="github" subtitle="github.com/documentdb/documentdb" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="documentdb-1.2-0.tar.gz intelrdfpmath-applied-2.0u3-1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg documentdb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install documentdb;		# install via package name, for the active PG version

pig install documentdb -v 19;   # install for PG 19
pig install documentdb -v 18;   # install for PG 18
pig install documentdb -v 17;   # install for PG 17
pig install documentdb -v 16;   # install for PG 16
pig install documentdb -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_documentdb_core, pg_cron, pg_documentdb';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION documentdb CASCADE; -- requires documentdb_core, pg_cron, postgis, tsm_system_rows, vector
```

## Usage

Sources:

- [DocumentDB v1.0-0 README](https://github.com/documentdb/documentdb/blob/v1.0-0/README.md)
- [DocumentDB v1.0-0 changelog](https://github.com/documentdb/documentdb/blob/v1.0-0/CHANGELOG.md)
- [`documentdb` control file](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb/documentdb.control)
- [Official preload helper](https://github.com/documentdb/documentdb/blob/v1.0-0/scripts/preload_libraries.sh)
- [DocumentDB 1.0 release](https://github.com/documentdb/documentdb/releases/tag/v1.0-0)
- [DocumentDB 1.2 release](https://github.com/documentdb/documentdb/releases/tag/v1.2-0)
- [1.2 control](https://github.com/documentdb/documentdb/blob/v1.2-0/pg_documentdb/documentdb.control)
- [1.1 to 1.2 SQL](https://github.com/documentdb/documentdb/blob/v1.2-0/pg_documentdb/sql/documentdb--1.1-0--1.2-0.sql)
- [Official upgrade procedure](https://github.com/documentdb/documentdb/blob/v1.2-0/packaging/README.md)

`documentdb` is the public PostgreSQL API extension for DocumentDB, an open-source MongoDB-compatible document database built on PostgreSQL. It stores BSON documents and implements CRUD, aggregation, full-text, geospatial, and vector workflows. MongoDB drivers require the separate DocumentDB gateway; installing this extension alone exposes the PostgreSQL API, not a wire-protocol listener.

### Configure and Install

The official deployment helper preloads the core and API libraries with `pg_cron`. Restart PostgreSQL after changing this setting:

```ini
shared_preload_libraries = 'pg_cron, pg_documentdb_core, pg_documentdb, pg_documentdb_extended_rum'
```

Install the public extension and its declared dependencies:

```sql
CREATE EXTENSION documentdb CASCADE;
CREATE EXTENSION documentdb_extended_rum;
```

`CASCADE` can install `documentdb_core`, `pg_cron`, `tsm_system_rows`, `vector`, and `postgis` when their files are present. Installation is superuser-only and non-relocatable.

### Native SQL Workflow

The SQL surface uses a database name, collection name, and BSON command document:

```sql
SELECT documentdb_api.create_collection('appdb', 'people');

SELECT documentdb_api.insert_one(
  'appdb',
  'people',
  '{"_id": 1, "name": "Ada", "team": "storage"}',
  NULL
);

SELECT document
FROM documentdb_api_catalog.bson_aggregation_find(
  'appdb',
  '{"find":"people","filter":{"team":"storage"}}'
);
```

For application compatibility, run the gateway and use a supported MongoDB driver against its configured TLS endpoint. The gateway translates wire-protocol commands into this PostgreSQL API.

### Important Objects

- `documentdb_api` contains collection-management and command functions such as `create_collection` and `insert_one`.
- `documentdb_api_catalog.bson_aggregation_find` executes a MongoDB-style find specification and returns BSON documents.
- `documentdb_core.bson` is the storage and interchange type supplied by `documentdb_core`.
- DocumentDB roles and internal schemas separate public read/write operations from administrative and implementation objects.
- `documentdb.enableNonBlockingUniqueIndexBuild` controls the v0.114 path for background unique ordered-index builds and is enabled by default in that release.

### Version and Operational Notes

Release 1.0-0 adds the first generally available release line. It fixes permission checks for zero-batch reads and write permissions for merge/output targets, along with crash, geospatial ordering and index-vacuum bugs. Custom role CRUD and post-join nested lookup remain disabled by default behind their settings. Match all libraries and apply the release upgrade SQL.

The official DocumentDB Local container enables usage telemetry by default; native packages are unaffected. Set DOCUMENTDB_USAGE_TELEMETRY=false, NO_ANALYTICS=1 or DO_NOT_TRACK=1 to disable it in the container.

MongoDB compatibility is not identical to every MongoDB server version. Test operators, index behavior, transactions, schema validation, authentication, and driver behavior used by the application. Match `documentdb`, `documentdb_core`, gateway, and optional distributed/index components to the same release line.

### DocumentDB 1.2 changes

DocumentDB 1.2 uses SQL version `1.2-0`. Back up the data directory before upgrading; reverting upgraded data to 1.0 is not supported. The official 1.2 container updates extensions before starting the gateway and refuses startup without an explicit password, including on an existing volume. For an adopted PostgreSQL instance, follow the package guide's restart and update sequence rather than restarting or altering unrelated instances.

Collation remains disabled unless `documentdb_core.enableCollation = on`, even for explicit `simple` collation. Collated non-unique indexes additionally require `documentdb.enableCollationWithNonUniqueOrderedIndexes`; collated unique indexes also require `documentdb.enable_collated_unique_indexes` and `documentdb.enable_composite_unique_optional_key`. The release adds per-command/per-update collation behavior and fixes descending unique-index null/missing conflicts, `$dateTrunc` day boundaries and rescanned `$densify` results. Review affected query and index semantics after upgrading.
