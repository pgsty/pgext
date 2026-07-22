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
| **9000** | {{< badge content="documentdb" link="https://github.com/documentdb/documentdb" >}} | {{< ext "documentdb" >}} | `0.114` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


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
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.114` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `documentdb` | `documentdb_core`, `pg_cron`, `postgis`, `tsm_system_rows`, `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.114` | {{< bg "18" "documentdb_18" "green" >}} {{< bg "17" "documentdb_17" "green" >}} {{< bg "16" "documentdb_16" "green" >}} {{< bg "15" "documentdb_15" "green" >}} {{< bg "14" "documentdb_14" "red" >}} | `documentdb_$v` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v`, `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.114` | {{< bg "18" "postgresql-18-documentdb" "green" >}} {{< bg "17" "postgresql-17-documentdb" "green" >}} {{< bg "16" "postgresql-16-documentdb" "green" >}} {{< bg "15" "postgresql-15-documentdb" "green" >}} {{< bg "14" "postgresql-14-documentdb" "red" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum`, `postgresql-$v-postgis-3` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_18` | `0.114` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.1 MiB | [documentdb_18-0.114-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_18-0.114-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_18` | `0.114` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [documentdb_18-0.114-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_18-0.114-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_18` | `0.114` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [documentdb_18-0.114-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_18-0.114-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_18` | `0.114` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [documentdb_18-0.114-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_18-0.114-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_18` | `0.114` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.0 MiB | [documentdb_18-0.114-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_18-0.114-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_18` | `0.114` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.9 MiB | [documentdb_18-0.114-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_18-0.114-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-documentdb` | `0.114` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.4 MiB | [postgresql-18-documentdb_0.114-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-documentdb` | `0.114` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 MiB | [postgresql-18-documentdb_0.114-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-documentdb` | `0.114` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.1 MiB | [postgresql-18-documentdb_0.114-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-documentdb` | `0.114` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.4 MiB | [postgresql-18-documentdb_0.114-0PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0PIGSTY~trixie_amd64.deb) |
| `postgresql-18-documentdb` | `0.113` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.0 MiB | [postgresql-18-documentdb_0.113-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-documentdb` | `0.112` | [d13.x86_64](/os/d13.x86_64) | pgdg | 4.9 MiB | [postgresql-18-documentdb_0.112-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.112-0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-documentdb` | `0.114` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.0 MiB | [postgresql-18-documentdb_0.114-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-documentdb` | `0.114` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 MiB | [postgresql-18-documentdb_0.114-0PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0PIGSTY~trixie_arm64.deb) |
| `postgresql-18-documentdb` | `0.113` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-18-documentdb_0.113-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-documentdb` | `0.112` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-18-documentdb_0.112-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.112-0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-documentdb` | `0.114` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.7 MiB | [postgresql-18-documentdb_0.114-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0PIGSTY~jammy_amd64.deb) |
| `postgresql-18-documentdb` | `0.114` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.6 MiB | [postgresql-18-documentdb_0.114-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0PIGSTY~jammy_arm64.deb) |
| `postgresql-18-documentdb` | `0.114` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.6 MiB | [postgresql-18-documentdb_0.114-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0PIGSTY~noble_amd64.deb) |
| `postgresql-18-documentdb` | `0.114` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.5 MiB | [postgresql-18-documentdb_0.114-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0PIGSTY~noble_arm64.deb) |
| `postgresql-18-documentdb` | `0.114` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.0 MiB | [postgresql-18-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-documentdb` | `0.114` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.6 MiB | [postgresql-18-documentdb_0.114-0PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0PIGSTY~resolute_amd64.deb) |
| `postgresql-18-documentdb` | `0.113` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.9 MiB | [postgresql-18-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-documentdb` | `0.112` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.8 MiB | [postgresql-18-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-documentdb` | `0.114` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.9 MiB | [postgresql-18-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-documentdb` | `0.114` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.5 MiB | [postgresql-18-documentdb_0.114-0PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-18-documentdb_0.114-0PIGSTY~resolute_arm64.deb) |
| `postgresql-18-documentdb` | `0.113` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-18-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-documentdb` | `0.112` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-18-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-18-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_17` | `0.114` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.1 MiB | [documentdb_17-0.114-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_17-0.114-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_17` | `0.114` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [documentdb_17-0.114-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_17-0.114-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_17` | `0.114` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [documentdb_17-0.114-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_17-0.114-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_17` | `0.114` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [documentdb_17-0.114-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_17-0.114-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_17` | `0.114` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.0 MiB | [documentdb_17-0.114-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_17-0.114-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_17` | `0.114` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.9 MiB | [documentdb_17-0.114-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_17-0.114-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-documentdb` | `0.114` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.4 MiB | [postgresql-17-documentdb_0.114-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-documentdb` | `0.114` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 MiB | [postgresql-17-documentdb_0.114-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-documentdb` | `0.114` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.1 MiB | [postgresql-17-documentdb_0.114-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-documentdb` | `0.114` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.4 MiB | [postgresql-17-documentdb_0.114-0PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0PIGSTY~trixie_amd64.deb) |
| `postgresql-17-documentdb` | `0.113` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.0 MiB | [postgresql-17-documentdb_0.113-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-documentdb` | `0.112` | [d13.x86_64](/os/d13.x86_64) | pgdg | 4.9 MiB | [postgresql-17-documentdb_0.112-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.112-0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-documentdb` | `0.114` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.0 MiB | [postgresql-17-documentdb_0.114-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-documentdb` | `0.114` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 MiB | [postgresql-17-documentdb_0.114-0PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0PIGSTY~trixie_arm64.deb) |
| `postgresql-17-documentdb` | `0.113` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-17-documentdb_0.113-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-documentdb` | `0.112` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-17-documentdb_0.112-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.112-0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-documentdb` | `0.114` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.1 MiB | [postgresql-17-documentdb_0.114-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0PIGSTY~jammy_amd64.deb) |
| `postgresql-17-documentdb` | `0.114` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.0 MiB | [postgresql-17-documentdb_0.114-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0PIGSTY~jammy_arm64.deb) |
| `postgresql-17-documentdb` | `0.114` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.6 MiB | [postgresql-17-documentdb_0.114-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0PIGSTY~noble_amd64.deb) |
| `postgresql-17-documentdb` | `0.114` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.5 MiB | [postgresql-17-documentdb_0.114-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0PIGSTY~noble_arm64.deb) |
| `postgresql-17-documentdb` | `0.114` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.0 MiB | [postgresql-17-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-documentdb` | `0.114` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.6 MiB | [postgresql-17-documentdb_0.114-0PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0PIGSTY~resolute_amd64.deb) |
| `postgresql-17-documentdb` | `0.113` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.9 MiB | [postgresql-17-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-documentdb` | `0.112` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.8 MiB | [postgresql-17-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-documentdb` | `0.114` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.9 MiB | [postgresql-17-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-documentdb` | `0.114` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.5 MiB | [postgresql-17-documentdb_0.114-0PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-17-documentdb_0.114-0PIGSTY~resolute_arm64.deb) |
| `postgresql-17-documentdb` | `0.113` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-17-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-documentdb` | `0.112` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-17-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-17-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_16` | `0.114` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.1 MiB | [documentdb_16-0.114-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_16-0.114-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_16` | `0.114` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [documentdb_16-0.114-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_16-0.114-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_16` | `0.114` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [documentdb_16-0.114-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_16-0.114-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_16` | `0.114` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [documentdb_16-0.114-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_16-0.114-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_16` | `0.114` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.0 MiB | [documentdb_16-0.114-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_16-0.114-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_16` | `0.114` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.9 MiB | [documentdb_16-0.114-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_16-0.114-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-documentdb` | `0.114` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.4 MiB | [postgresql-16-documentdb_0.114-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-documentdb` | `0.114` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 MiB | [postgresql-16-documentdb_0.114-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-documentdb` | `0.114` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.1 MiB | [postgresql-16-documentdb_0.114-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-documentdb` | `0.114` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.4 MiB | [postgresql-16-documentdb_0.114-0PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0PIGSTY~trixie_amd64.deb) |
| `postgresql-16-documentdb` | `0.113` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.0 MiB | [postgresql-16-documentdb_0.113-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-documentdb` | `0.112` | [d13.x86_64](/os/d13.x86_64) | pgdg | 4.9 MiB | [postgresql-16-documentdb_0.112-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.112-0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-documentdb` | `0.114` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.0 MiB | [postgresql-16-documentdb_0.114-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-documentdb` | `0.114` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 MiB | [postgresql-16-documentdb_0.114-0PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0PIGSTY~trixie_arm64.deb) |
| `postgresql-16-documentdb` | `0.113` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-16-documentdb_0.113-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-documentdb` | `0.112` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-16-documentdb_0.112-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.112-0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-documentdb` | `0.114` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.1 MiB | [postgresql-16-documentdb_0.114-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0PIGSTY~jammy_amd64.deb) |
| `postgresql-16-documentdb` | `0.114` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.0 MiB | [postgresql-16-documentdb_0.114-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0PIGSTY~jammy_arm64.deb) |
| `postgresql-16-documentdb` | `0.114` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.6 MiB | [postgresql-16-documentdb_0.114-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0PIGSTY~noble_amd64.deb) |
| `postgresql-16-documentdb` | `0.114` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.5 MiB | [postgresql-16-documentdb_0.114-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0PIGSTY~noble_arm64.deb) |
| `postgresql-16-documentdb` | `0.114` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.0 MiB | [postgresql-16-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-documentdb` | `0.114` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.6 MiB | [postgresql-16-documentdb_0.114-0PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0PIGSTY~resolute_amd64.deb) |
| `postgresql-16-documentdb` | `0.113` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.9 MiB | [postgresql-16-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-documentdb` | `0.112` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.8 MiB | [postgresql-16-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-documentdb` | `0.114` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.9 MiB | [postgresql-16-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-documentdb` | `0.114` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.5 MiB | [postgresql-16-documentdb_0.114-0PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-16-documentdb_0.114-0PIGSTY~resolute_arm64.deb) |
| `postgresql-16-documentdb` | `0.113` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-16-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-documentdb` | `0.112` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-16-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-16-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_15` | `0.114` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.2 MiB | [documentdb_15-0.114-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_15-0.114-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_15` | `0.114` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [documentdb_15-0.114-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_15-0.114-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_15` | `0.114` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [documentdb_15-0.114-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_15-0.114-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_15` | `0.114` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.8 MiB | [documentdb_15-0.114-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_15-0.114-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_15` | `0.114` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.0 MiB | [documentdb_15-0.114-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_15-0.114-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_15` | `0.114` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.9 MiB | [documentdb_15-0.114-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_15-0.114-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-documentdb` | `0.114` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.4 MiB | [postgresql-15-documentdb_0.114-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-documentdb` | `0.114` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 MiB | [postgresql-15-documentdb_0.114-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-documentdb` | `0.114` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.2 MiB | [postgresql-15-documentdb_0.114-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.114` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.4 MiB | [postgresql-15-documentdb_0.114-0PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0PIGSTY~trixie_amd64.deb) |
| `postgresql-15-documentdb` | `0.113` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.0 MiB | [postgresql-15-documentdb_0.113-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.112` | [d13.x86_64](/os/d13.x86_64) | pgdg | 5.0 MiB | [postgresql-15-documentdb_0.112-0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.112-0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.114` | [d13.aarch64](/os/d13.aarch64) | pgdg | 5.0 MiB | [postgresql-15-documentdb_0.114-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-documentdb` | `0.114` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 MiB | [postgresql-15-documentdb_0.114-0PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0PIGSTY~trixie_arm64.deb) |
| `postgresql-15-documentdb` | `0.113` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.9 MiB | [postgresql-15-documentdb_0.113-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-documentdb` | `0.112` | [d13.aarch64](/os/d13.aarch64) | pgdg | 4.8 MiB | [postgresql-15-documentdb_0.112-0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.112-0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-documentdb` | `0.114` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.2 MiB | [postgresql-15-documentdb_0.114-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0PIGSTY~jammy_amd64.deb) |
| `postgresql-15-documentdb` | `0.114` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.0 MiB | [postgresql-15-documentdb_0.114-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0PIGSTY~jammy_arm64.deb) |
| `postgresql-15-documentdb` | `0.114` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.6 MiB | [postgresql-15-documentdb_0.114-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0PIGSTY~noble_amd64.deb) |
| `postgresql-15-documentdb` | `0.114` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.5 MiB | [postgresql-15-documentdb_0.114-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0PIGSTY~noble_arm64.deb) |
| `postgresql-15-documentdb` | `0.114` | [u26.x86_64](/os/u26.x86_64) | pgdg | 5.1 MiB | [postgresql-15-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.114` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.6 MiB | [postgresql-15-documentdb_0.114-0PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0PIGSTY~resolute_amd64.deb) |
| `postgresql-15-documentdb` | `0.113` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.9 MiB | [postgresql-15-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.112` | [u26.x86_64](/os/u26.x86_64) | pgdg | 4.9 MiB | [postgresql-15-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.112-0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-documentdb` | `0.114` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.9 MiB | [postgresql-15-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-documentdb` | `0.114` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.5 MiB | [postgresql-15-documentdb_0.114-0PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/documentdb/postgresql-15-documentdb_0.114-0PIGSTY~resolute_arm64.deb) |
| `postgresql-15-documentdb` | `0.113` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.8 MiB | [postgresql-15-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.113-0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-documentdb` | `0.112` | [u26.aarch64](/os/u26.aarch64) | pgdg | 4.7 MiB | [postgresql-15-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/d/documentdb/postgresql-15-documentdb_0.112-0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/documentdb/documentdb" title="Repository" icon="github" subtitle="github.com/documentdb/documentdb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="documentdb-0.114-0.tar.gz" >}}
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

Sources:

- [DocumentDB v0.114-0 README](https://github.com/documentdb/documentdb/blob/v0.114-0/README.md)
- [DocumentDB v0.114-0 changelog](https://github.com/documentdb/documentdb/blob/v0.114-0/CHANGELOG.md)
- [`documentdb` control file](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb/documentdb.control)
- [Official preload helper](https://github.com/documentdb/documentdb/blob/v0.114-0/scripts/preload_libraries.sh)

`documentdb` is the public PostgreSQL API extension for DocumentDB, an open-source MongoDB-compatible document database built on PostgreSQL. It stores BSON documents and implements CRUD, aggregation, full-text, geospatial, and vector workflows. MongoDB drivers require the separate DocumentDB gateway; installing this extension alone exposes the PostgreSQL API, not a wire-protocol listener.

### Configure and Install

The official deployment helper preloads the core and API libraries with `pg_cron`. Restart PostgreSQL after changing this setting:

```conf
shared_preload_libraries = 'pg_cron, pg_documentdb_core, pg_documentdb'
```

Install the public extension and its declared dependencies:

```sql
CREATE EXTENSION documentdb CASCADE;
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

The v0.114-0 tagged changelog enables schema validation by default, fixes validator propagation and caching, and enables non-blocking unique ordered-index builds. It also records gateway configuration, connectivity-check, TLS, and credential-handling improvements. Two RUM optimizations in that changelog remain feature-flagged and disabled by default; do not describe them as active behavior.

MongoDB compatibility is not identical to every MongoDB server version. Test operators, index behavior, transactions, schema validation, authentication, and driver behavior used by the application. Match `documentdb`, `documentdb_core`, gateway, and optional distributed/index components to the same release line.
