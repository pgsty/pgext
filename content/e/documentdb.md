---
title: "documentdb"
linkTitle: "documentdb"
description: "API surface for DocumentDB for PostgreSQL"
weight: 9000
categories: ["SIM"]
width: full
---

[**documentdb**](https://github.com/documentdb/documentdb) : API surface for DocumentDB for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9000** | {{< badge content="documentdb" link="https://github.com/documentdb/documentdb" >}} | {{< ext "documentdb" >}} | `0.113` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "documentdb_core" >}} {{< ext "pg_cron" >}} {{< ext "postgis" >}} {{< ext "tsm_system_rows" >}} {{< ext "vector" >}} |
|    **Need By**    | {{< ext "documentdb_distributed" >}} {{< ext "documentdb_extended_rum" >}} |
|   **See Also**    | {{< ext "mongo_fdw" >}} {{< ext "wal2mongo" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} |
|    **Siblings**   | {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} {{< ext "documentdb_extended_rum" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.113` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `documentdb` | `documentdb_core`, `pg_cron`, `postgis`, `tsm_system_rows`, `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.113` | {{< bg "18" "documentdb_18" "green" >}} {{< bg "17" "documentdb_17" "green" >}} {{< bg "16" "documentdb_16" "green" >}} {{< bg "15" "documentdb_15" "green" >}} {{< bg "14" "documentdb_14" "red" >}} | `documentdb_$v` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v`, `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.113` | {{< bg "18" "postgresql-18-documentdb" "green" >}} {{< bg "17" "postgresql-17-documentdb" "green" >}} {{< bg "16" "postgresql-16-documentdb" "green" >}} {{< bg "15" "postgresql-15-documentdb" "green" >}} {{< bg "14" "postgresql-14-documentdb" "red" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum`, `postgresql-$v-postgis-3` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.113" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.113" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.113" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.113" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.113" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.113" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.113" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.113" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.113" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.113" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.113" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.113" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.113" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_18` | `0.113` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.0 MiB | [documentdb_18-0.113-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_18-0.113-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_18` | `0.113` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.9 MiB | [documentdb_18-0.113-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_18-0.113-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_18` | `0.113` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [documentdb_18-0.113-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_18-0.113-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_18` | `0.113` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [documentdb_18-0.113-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_18-0.113-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_18` | `0.113` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [documentdb_18-0.113-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_18-0.113-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_18` | `0.113` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [documentdb_18-0.113-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_18-0.113-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-documentdb` | `0.113` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.2 MiB | [postgresql-18-documentdb_0.113-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-documentdb` | `0.113` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.0 MiB | [postgresql-18-documentdb_0.113-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-documentdb` | `0.114` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.1 MiB | [postgresql-18-documentdb_0.114-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-documentdb` | `0.113` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.0 MiB | [postgresql-18-documentdb_0.113-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-documentdb` | `0.113` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.2 MiB | [postgresql-18-documentdb_0.113-0PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0PIGSTY~trixie_amd64.deb) |
| `postgresql-18-documentdb` | `0.112` | [d13.x86_64](/os/d13.x86_64) | pgdg | 4.9 MiB | [postgresql-18-documentdb_0.112-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.112-0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-documentdb` | `0.114` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.0 MiB | [postgresql-18-documentdb_0.114-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-documentdb` | `0.113` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-18-documentdb_0.113-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-documentdb` | `0.113` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.0 MiB | [postgresql-18-documentdb_0.113-0PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0PIGSTY~trixie_arm64.deb) |
| `postgresql-18-documentdb` | `0.112` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-18-documentdb_0.112-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.112-0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-documentdb` | `0.113` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.5 MiB | [postgresql-18-documentdb_0.113-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0PIGSTY~jammy_amd64.deb) |
| `postgresql-18-documentdb` | `0.113` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.4 MiB | [postgresql-18-documentdb_0.113-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0PIGSTY~jammy_arm64.deb) |
| `postgresql-18-documentdb` | `0.113` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.4 MiB | [postgresql-18-documentdb_0.113-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0PIGSTY~noble_amd64.deb) |
| `postgresql-18-documentdb` | `0.113` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.3 MiB | [postgresql-18-documentdb_0.113-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0PIGSTY~noble_arm64.deb) |
| `postgresql-18-documentdb` | `0.114` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.0 MiB | [postgresql-18-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-documentdb` | `0.113` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.9 MiB | [postgresql-18-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-documentdb` | `0.113` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.4 MiB | [postgresql-18-documentdb_0.113-0PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0PIGSTY~resolute_amd64.deb) |
| `postgresql-18-documentdb` | `0.112` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.8 MiB | [postgresql-18-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-documentdb` | `0.114` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.9 MiB | [postgresql-18-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-documentdb` | `0.113` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-18-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-documentdb` | `0.113` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.3 MiB | [postgresql-18-documentdb_0.113-0PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0PIGSTY~resolute_arm64.deb) |
| `postgresql-18-documentdb` | `0.112` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-18-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_17` | `0.113` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.0 MiB | [documentdb_17-0.113-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_17-0.113-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_17` | `0.113` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.9 MiB | [documentdb_17-0.113-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_17-0.113-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_17` | `0.113` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [documentdb_17-0.113-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_17-0.113-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_17` | `0.113` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [documentdb_17-0.113-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_17-0.113-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_17` | `0.113` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [documentdb_17-0.113-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_17-0.113-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_17` | `0.113` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [documentdb_17-0.113-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_17-0.113-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-documentdb` | `0.113` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.2 MiB | [postgresql-17-documentdb_0.113-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-documentdb` | `0.113` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.0 MiB | [postgresql-17-documentdb_0.113-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-documentdb` | `0.114` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.1 MiB | [postgresql-17-documentdb_0.114-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-documentdb` | `0.113` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.0 MiB | [postgresql-17-documentdb_0.113-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-documentdb` | `0.113` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.2 MiB | [postgresql-17-documentdb_0.113-0PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0PIGSTY~trixie_amd64.deb) |
| `postgresql-17-documentdb` | `0.112` | [d13.x86_64](/os/d13.x86_64) | pgdg | 4.9 MiB | [postgresql-17-documentdb_0.112-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.112-0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-documentdb` | `0.114` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.0 MiB | [postgresql-17-documentdb_0.114-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-documentdb` | `0.113` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-17-documentdb_0.113-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-documentdb` | `0.113` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.0 MiB | [postgresql-17-documentdb_0.113-0PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0PIGSTY~trixie_arm64.deb) |
| `postgresql-17-documentdb` | `0.112` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-17-documentdb_0.112-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.112-0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-documentdb` | `0.113` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.9 MiB | [postgresql-17-documentdb_0.113-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0PIGSTY~jammy_amd64.deb) |
| `postgresql-17-documentdb` | `0.113` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.8 MiB | [postgresql-17-documentdb_0.113-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0PIGSTY~jammy_arm64.deb) |
| `postgresql-17-documentdb` | `0.113` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.4 MiB | [postgresql-17-documentdb_0.113-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0PIGSTY~noble_amd64.deb) |
| `postgresql-17-documentdb` | `0.113` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.3 MiB | [postgresql-17-documentdb_0.113-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0PIGSTY~noble_arm64.deb) |
| `postgresql-17-documentdb` | `0.114` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.0 MiB | [postgresql-17-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-documentdb` | `0.113` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.9 MiB | [postgresql-17-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-documentdb` | `0.113` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.4 MiB | [postgresql-17-documentdb_0.113-0PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0PIGSTY~resolute_amd64.deb) |
| `postgresql-17-documentdb` | `0.112` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.8 MiB | [postgresql-17-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-documentdb` | `0.114` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.9 MiB | [postgresql-17-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-documentdb` | `0.113` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-17-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-documentdb` | `0.113` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.3 MiB | [postgresql-17-documentdb_0.113-0PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0PIGSTY~resolute_arm64.deb) |
| `postgresql-17-documentdb` | `0.112` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-17-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_16` | `0.113` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.1 MiB | [documentdb_16-0.113-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_16-0.113-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_16` | `0.113` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.9 MiB | [documentdb_16-0.113-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_16-0.113-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_16` | `0.113` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [documentdb_16-0.113-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_16-0.113-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_16` | `0.113` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [documentdb_16-0.113-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_16-0.113-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_16` | `0.113` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [documentdb_16-0.113-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_16-0.113-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_16` | `0.113` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [documentdb_16-0.113-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_16-0.113-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-documentdb` | `0.113` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.2 MiB | [postgresql-16-documentdb_0.113-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-documentdb` | `0.113` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.0 MiB | [postgresql-16-documentdb_0.113-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-documentdb` | `0.114` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.1 MiB | [postgresql-16-documentdb_0.114-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-documentdb` | `0.113` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.0 MiB | [postgresql-16-documentdb_0.113-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-documentdb` | `0.113` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.2 MiB | [postgresql-16-documentdb_0.113-0PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0PIGSTY~trixie_amd64.deb) |
| `postgresql-16-documentdb` | `0.112` | [d13.x86_64](/os/d13.x86_64) | pgdg | 4.9 MiB | [postgresql-16-documentdb_0.112-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.112-0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-documentdb` | `0.114` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.0 MiB | [postgresql-16-documentdb_0.114-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-documentdb` | `0.113` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-16-documentdb_0.113-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-documentdb` | `0.113` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.0 MiB | [postgresql-16-documentdb_0.113-0PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0PIGSTY~trixie_arm64.deb) |
| `postgresql-16-documentdb` | `0.112` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-16-documentdb_0.112-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.112-0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-documentdb` | `0.113` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.9 MiB | [postgresql-16-documentdb_0.113-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0PIGSTY~jammy_amd64.deb) |
| `postgresql-16-documentdb` | `0.113` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.8 MiB | [postgresql-16-documentdb_0.113-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0PIGSTY~jammy_arm64.deb) |
| `postgresql-16-documentdb` | `0.113` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.4 MiB | [postgresql-16-documentdb_0.113-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0PIGSTY~noble_amd64.deb) |
| `postgresql-16-documentdb` | `0.113` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.3 MiB | [postgresql-16-documentdb_0.113-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0PIGSTY~noble_arm64.deb) |
| `postgresql-16-documentdb` | `0.114` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.0 MiB | [postgresql-16-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-documentdb` | `0.113` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.9 MiB | [postgresql-16-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-documentdb` | `0.113` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.4 MiB | [postgresql-16-documentdb_0.113-0PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0PIGSTY~resolute_amd64.deb) |
| `postgresql-16-documentdb` | `0.112` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.8 MiB | [postgresql-16-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-documentdb` | `0.114` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.9 MiB | [postgresql-16-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-documentdb` | `0.113` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-16-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-documentdb` | `0.113` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.3 MiB | [postgresql-16-documentdb_0.113-0PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0PIGSTY~resolute_arm64.deb) |
| `postgresql-16-documentdb` | `0.112` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-16-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_15` | `0.113` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.1 MiB | [documentdb_15-0.113-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_15-0.113-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_15` | `0.113` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.9 MiB | [documentdb_15-0.113-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_15-0.113-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_15` | `0.113` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [documentdb_15-0.113-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_15-0.113-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_15` | `0.113` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [documentdb_15-0.113-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_15-0.113-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_15` | `0.113` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [documentdb_15-0.113-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_15-0.113-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_15` | `0.113` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.8 MiB | [documentdb_15-0.113-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_15-0.113-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-documentdb` | `0.113` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.2 MiB | [postgresql-15-documentdb_0.113-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-documentdb` | `0.113` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.0 MiB | [postgresql-15-documentdb_0.113-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-documentdb` | `0.114` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.2 MiB | [postgresql-15-documentdb_0.114-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.113` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.0 MiB | [postgresql-15-documentdb_0.113-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.113` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.3 MiB | [postgresql-15-documentdb_0.113-0PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0PIGSTY~trixie_amd64.deb) |
| `postgresql-15-documentdb` | `0.112` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.0 MiB | [postgresql-15-documentdb_0.112-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.112-0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.114` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.0 MiB | [postgresql-15-documentdb_0.114-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-documentdb` | `0.113` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.9 MiB | [postgresql-15-documentdb_0.113-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-documentdb` | `0.113` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 MiB | [postgresql-15-documentdb_0.113-0PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0PIGSTY~trixie_arm64.deb) |
| `postgresql-15-documentdb` | `0.112` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-15-documentdb_0.112-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.112-0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-documentdb` | `0.113` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.0 MiB | [postgresql-15-documentdb_0.113-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0PIGSTY~jammy_amd64.deb) |
| `postgresql-15-documentdb` | `0.113` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.9 MiB | [postgresql-15-documentdb_0.113-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0PIGSTY~jammy_arm64.deb) |
| `postgresql-15-documentdb` | `0.113` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.4 MiB | [postgresql-15-documentdb_0.113-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0PIGSTY~noble_amd64.deb) |
| `postgresql-15-documentdb` | `0.113` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.3 MiB | [postgresql-15-documentdb_0.113-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0PIGSTY~noble_arm64.deb) |
| `postgresql-15-documentdb` | `0.114` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.1 MiB | [postgresql-15-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.113` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.9 MiB | [postgresql-15-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.113` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.4 MiB | [postgresql-15-documentdb_0.113-0PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0PIGSTY~resolute_amd64.deb) |
| `postgresql-15-documentdb` | `0.112` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.9 MiB | [postgresql-15-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.114` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.9 MiB | [postgresql-15-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-documentdb` | `0.113` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.8 MiB | [postgresql-15-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-documentdb` | `0.113` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.3 MiB | [postgresql-15-documentdb_0.113-0PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0PIGSTY~resolute_arm64.deb) |
| `postgresql-15-documentdb` | `0.112` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-15-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/documentdb/documentdb" title="Repository" icon="github" subtitle="github.com/documentdb/documentdb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="documentdb-0.113-0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg documentdb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install documentdb;		# install via package name, for the active PG version

pig install documentdb -v 18;   # install for PG 18
pig install documentdb -v 17;   # install for PG 17
pig install documentdb -v 16;   # install for PG 16
pig install documentdb -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_documentdb, pg_documentdb_core, pg_cron';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION documentdb CASCADE; -- requires documentdb_core, pg_cron, postgis, tsm_system_rows, vector
```




## Usage

Sources: [README](https://github.com/documentdb/documentdb/blob/v0.113-0/README.md), [CHANGELOG](https://github.com/documentdb/documentdb/blob/v0.113-0/CHANGELOG.md), [documentdb.control](https://github.com/documentdb/documentdb/blob/v0.113-0/pg_documentdb/documentdb.control), [documentdb_core.control](https://github.com/documentdb/documentdb/blob/v0.113-0/pg_documentdb_core/documentdb_core.control), [documentdb_extended_rum.control](https://github.com/documentdb/documentdb/blob/v0.113-0/pg_documentdb_extended_rum/documentdb_extended_rum.control)

`documentdb` is a MongoDB-compatible document database implemented as PostgreSQL extensions. It adds BSON storage and APIs in PostgreSQL, plus an optional gateway layer for MongoDB wire-protocol clients. FerretDB 2.0+ can use DocumentDB as its backend.

### Components

The public extension surface is split across related extensions:

- `documentdb_core`: BSON datatype and low-level BSON operations.
- `documentdb`: public API for document CRUD and query behavior.
- `documentdb_extended_rum`: extended RUM access method used by DocumentDB indexing.
- `pg_documentdb_gw`: gateway protocol layer used by the local Docker image and MongoDB-compatible clients.

Install the SQL extension in each database that needs the API:

```sql
CREATE EXTENSION IF NOT EXISTS documentdb CASCADE;
```

The `documentdb.control` file for `0.113-0` declares dependencies on `documentdb_core`, `pg_cron`, `tsm_system_rows`, `vector`, and `postgis`. The gateway container listens on a MongoDB-compatible port; the README examples use `10260` to avoid colliding with local MongoDB services.

### MongoDB Client Example

```python
import pymongo

client = pymongo.MongoClient(
    "mongodb://user:pass@localhost:10260/?tls=true&tlsAllowInvalidCertificates=true"
)

db = client["quickStartDatabase"]
coll = db.create_collection("quickStartCollection")

coll.insert_one({"name": "Alice", "email": "alice@example.com"})
print(coll.find_one({"name": "Alice"}))
```

The upstream README also demonstrates aggregation pipelines through normal MongoDB drivers:

```python
pipeline = [
    {"$match": {"name": "Alice"}},
    {"$project": {"_id": 0, "name": 1, "email": 1}},
]

for doc in coll.aggregate(pipeline):
    print(doc)
```

### Version Notes

This project's CSV tracks DocumentDB `0.113` for PostgreSQL 15-18. The upstream tag is `v0.113-0`; control files report `default_version = '0.113-0'`.

The `0.111` through `0.113` changelog entries are mostly query-planner, collation, and index correctness work:

- `0.113-0` adds opt-in collation support for non-unique ordered indexes with `$in` and `$nin`, and supports pruning dead index entries on ordered TTL indexes behind feature flags.
- `0.112-0` removes the legacy composite-returning `bson_update_document` UDF path, expands non-unique ordered-index collation support, and improves `$group` and accumulator execution.
- `0.111-0` adds background init job infrastructure, more `$group` validation, collation/index pushdown improvements, and several crash fixes.

### Caveats

- DocumentDB is a multi-extension stack; `CREATE EXTENSION documentdb CASCADE` is the normal entry point, but operational deployments also need the gateway/runtime pieces if MongoDB wire compatibility is required.
- Some features listed in the changelog are gated by `documentdb.*` feature flags. Verify flag defaults in the exact installed build before documenting behavior as always-on.
- `documentdb_extended_rum` is relocatable, but `documentdb` and `documentdb_core` are not.
