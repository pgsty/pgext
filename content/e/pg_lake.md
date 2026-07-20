---
title: "pg_lake"
linkTitle: "pg_lake"
description: "Data lake extension by Snowflake"
weight: 2560
categories: ["OLAP"]
width: full
---

[**pg_lake**](https://github.com/Snowflake-Labs/pg_lake) : Data lake extension by Snowflake


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2560** | {{< badge content="pg_lake" link="https://github.com/Snowflake-Labs/pg_lake" >}} | {{< ext "pg_lake" >}} | `3.4` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `lake` `pg_catalog` |
|   **Requires**    | {{< ext "pg_lake_copy" >}} {{< ext "pg_lake_table" >}} |
|   **See Also**    | {{< ext "duckdb_fdw" >}} {{< ext "pg_duckdb" >}} {{< ext "pg_ducklake" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_parquet" >}} |
|    **Siblings**   | {{< ext "pg_extension_base" >}} {{< ext "pg_extension_updater" >}} {{< ext "pg_map" >}} {{< ext "pg_lake_engine" >}} {{< ext "pg_lake_iceberg" >}} {{< ext "pg_lake_table" >}} {{< ext "pg_lake_copy" >}} |

> [!Note] Pigsty packages this release for PG16-18. Configure shared_preload_libraries=pg_extension_base and run the matching PG-major pgduck_server process. RPM supports EL9/EL10 only; EL8 is rejected because OpenSSL 3 is required. DEB supports Debian 12/13 and Ubuntu 22.04/24.04/26.04 on amd64/arm64. DuckDB and Avro are private per PG major. Co-installation with pg_duckdb, pg_mooncake, and duckdb_fdw is file-safe, but overlapping hooks and COPY behavior can be preload-order-sensitive.
Extension SQL/control version is 3.4; source and DEB/RPM package version is 3.4.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_lake` | `pg_lake_copy`, `pg_lake_table` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4.0` | {{< bg "18" "pg_lake_18" "green" >}} {{< bg "17" "pg_lake_17" "green" >}} {{< bg "16" "pg_lake_16" "green" >}} {{< bg "15" "pg_lake_15" "red" >}} {{< bg "14" "pg_lake_14" "red" >}} | `pg_lake_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4.0` | {{< bg "18" "postgresql-18-pg-lake" "green" >}} {{< bg "17" "postgresql-17-pg-lake" "green" >}} {{< bg "16" "postgresql-16-pg-lake" "green" >}} {{< bg "15" "postgresql-15-pg-lake" "red" >}} {{< bg "14" "postgresql-14-pg-lake" "red" >}} | `postgresql-$v-pg-lake` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_lake_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_lake_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lake_18` | `3.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.8 MiB | [pg_lake_18-3.4.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lake_18-3.4.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_lake_18` | `3.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.2 MiB | [pg_lake_18-3.4.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lake_18-3.4.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_lake_18` | `3.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.2 MiB | [pg_lake_18-3.4.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lake_18-3.4.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_lake_18` | `3.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.4 MiB | [pg_lake_18-3.4.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lake_18-3.4.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-lake` | `3.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.0 MiB | [postgresql-18-pg-lake_3.4.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lake/postgresql-18-pg-lake_3.4.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-lake` | `3.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.8 MiB | [postgresql-18-pg-lake_3.4.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lake/postgresql-18-pg-lake_3.4.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-lake` | `3.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.0 MiB | [postgresql-18-pg-lake_3.4.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lake/postgresql-18-pg-lake_3.4.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-lake` | `3.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.6 MiB | [postgresql-18-pg-lake_3.4.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lake/postgresql-18-pg-lake_3.4.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-lake` | `3.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.4 MiB | [postgresql-18-pg-lake_3.4.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lake/postgresql-18-pg-lake_3.4.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-lake` | `3.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.1 MiB | [postgresql-18-pg-lake_3.4.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lake/postgresql-18-pg-lake_3.4.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-lake` | `3.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.5 MiB | [postgresql-18-pg-lake_3.4.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lake/postgresql-18-pg-lake_3.4.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-lake` | `3.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.2 MiB | [postgresql-18-pg-lake_3.4.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lake/postgresql-18-pg-lake_3.4.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-lake` | `3.4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.0 MiB | [postgresql-18-pg-lake_3.4.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lake/postgresql-18-pg-lake_3.4.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-lake` | `3.4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.9 MiB | [postgresql-18-pg-lake_3.4.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lake/postgresql-18-pg-lake_3.4.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lake_17` | `3.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.9 MiB | [pg_lake_17-3.4.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lake_17-3.4.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_lake_17` | `3.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.2 MiB | [pg_lake_17-3.4.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lake_17-3.4.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_lake_17` | `3.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.2 MiB | [pg_lake_17-3.4.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lake_17-3.4.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_lake_17` | `3.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.3 MiB | [pg_lake_17-3.4.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lake_17-3.4.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-lake` | `3.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.0 MiB | [postgresql-17-pg-lake_3.4.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lake/postgresql-17-pg-lake_3.4.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-lake` | `3.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.8 MiB | [postgresql-17-pg-lake_3.4.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lake/postgresql-17-pg-lake_3.4.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-lake` | `3.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.0 MiB | [postgresql-17-pg-lake_3.4.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lake/postgresql-17-pg-lake_3.4.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-lake` | `3.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.5 MiB | [postgresql-17-pg-lake_3.4.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lake/postgresql-17-pg-lake_3.4.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-lake` | `3.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.6 MiB | [postgresql-17-pg-lake_3.4.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lake/postgresql-17-pg-lake_3.4.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-lake` | `3.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.3 MiB | [postgresql-17-pg-lake_3.4.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lake/postgresql-17-pg-lake_3.4.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-lake` | `3.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.5 MiB | [postgresql-17-pg-lake_3.4.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lake/postgresql-17-pg-lake_3.4.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-lake` | `3.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.2 MiB | [postgresql-17-pg-lake_3.4.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lake/postgresql-17-pg-lake_3.4.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-lake` | `3.4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.0 MiB | [postgresql-17-pg-lake_3.4.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lake/postgresql-17-pg-lake_3.4.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-lake` | `3.4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.8 MiB | [postgresql-17-pg-lake_3.4.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lake/postgresql-17-pg-lake_3.4.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lake_16` | `3.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.9 MiB | [pg_lake_16-3.4.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lake_16-3.4.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_lake_16` | `3.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.2 MiB | [pg_lake_16-3.4.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lake_16-3.4.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_lake_16` | `3.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.2 MiB | [pg_lake_16-3.4.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lake_16-3.4.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_lake_16` | `3.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.3 MiB | [pg_lake_16-3.4.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lake_16-3.4.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-lake` | `3.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.0 MiB | [postgresql-16-pg-lake_3.4.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lake/postgresql-16-pg-lake_3.4.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-lake` | `3.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.8 MiB | [postgresql-16-pg-lake_3.4.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lake/postgresql-16-pg-lake_3.4.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-lake` | `3.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.9 MiB | [postgresql-16-pg-lake_3.4.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lake/postgresql-16-pg-lake_3.4.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-lake` | `3.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.5 MiB | [postgresql-16-pg-lake_3.4.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lake/postgresql-16-pg-lake_3.4.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-lake` | `3.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.5 MiB | [postgresql-16-pg-lake_3.4.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lake/postgresql-16-pg-lake_3.4.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-lake` | `3.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.2 MiB | [postgresql-16-pg-lake_3.4.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lake/postgresql-16-pg-lake_3.4.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-lake` | `3.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.5 MiB | [postgresql-16-pg-lake_3.4.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lake/postgresql-16-pg-lake_3.4.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-lake` | `3.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.2 MiB | [postgresql-16-pg-lake_3.4.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lake/postgresql-16-pg-lake_3.4.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-lake` | `3.4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.0 MiB | [postgresql-16-pg-lake_3.4.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lake/postgresql-16-pg-lake_3.4.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-lake` | `3.4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.8 MiB | [postgresql-16-pg-lake_3.4.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lake/postgresql-16-pg-lake_3.4.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Snowflake-Labs/pg_lake" title="Repository" icon="github" subtitle="github.com/Snowflake-Labs/pg_lake" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_lake-3.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_lake;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_lake;		# install via package name, for the active PG version

pig install pg_lake -v 18;   # install for PG 18
pig install pg_lake -v 17;   # install for PG 17
pig install pg_lake -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_extension_base';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_lake CASCADE; -- requires pg_lake_copy, pg_lake_table
```

## Usage

Sources:

- [Official extension control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake/pg_lake.control)
- [Official upstream documentation](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/README.md)

`pg_lake` — Data lake extension by Snowflake

The reviewed catalog snapshot records version `3.4`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pg_lake_copy`, `pg_lake_table`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_lake";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lake';
```

The upstream project is associated with `Snowflake`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
