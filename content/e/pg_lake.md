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
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_lake_18 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_lake_18 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lake_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lake_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lake_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lake_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |


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

- [Official pg_lake README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake/pg_lake.control)
- [Official build and startup guide](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/building-from-source.md)
- [Official project documentation index](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/README.md)

`pg_lake` is the top-level extension for Snowflake's PostgreSQL lakehouse stack. It installs the table, Iceberg, copy, query-engine, extension-base, and map components needed to query object-store files and create transactional Iceberg tables. The PostgreSQL extensions orchestrate planning and transactions while a separate local `pgduck_server` process executes vectorized work with DuckDB.

### Start the Stack

Version `3.4` supports PostgreSQL 16 through 18. Preload the common extension infrastructure, restart PostgreSQL, and start `pgduck_server` on the database host:

```conf
shared_preload_libraries = 'pg_extension_base'
```

```shell
pgduck_server --cache_dir /var/cache/pg_lake
```

Create the complete dependency tree in the target database:

```sql
CREATE EXTENSION pg_lake CASCADE;
SELECT lake.version();
```

Configure object-store credentials for `pgduck_server`, then choose the managed Iceberg location:

```sql
SET pg_lake_iceberg.default_location_prefix =
    's3://analytics-bucket/warehouse';
```

### Core Workflows

Create and modify a transactional Iceberg table:

```sql
CREATE TABLE measurements (
    station_name text NOT NULL,
    measured_at timestamptz NOT NULL,
    value double precision
) USING iceberg;

INSERT INTO measurements VALUES
    ('Istanbul', now(), 18.5),
    ('Haarlem', now(), 9.3);
```

Import or export Parquet, CSV, or newline-delimited JSON through `COPY`:

```sql
COPY (SELECT * FROM measurements)
TO 's3://analytics-bucket/export/measurements.parquet';

COPY measurements
FROM 's3://analytics-bucket/import/measurements.parquet';
```

Query files without loading them into PostgreSQL:

```sql
CREATE FOREIGN TABLE external_events ()
SERVER pg_lake
OPTIONS (path 's3://analytics-bucket/events/*.parquet');

SELECT count(*) FROM external_events;
```

### Component Index

- `pg_lake`: meta-extension and `lake.version()`.
- `pg_lake_table`: data-lake FDW, Iceberg table syntax, file utilities, and table catalogs.
- `pg_lake_iceberg`: Iceberg metadata, snapshots, manifests, and catalog integration.
- `pg_lake_copy`: `COPY` interception for object-store files and lake formats.
- `pg_lake_engine`: shared query rewrite, type conversion, cleanup, and `pgduck_server` client layer.
- `pg_extension_base`: preload and lifecycle-worker infrastructure.
- `pg_map`: generated PostgreSQL map types used for nested lake data.

### Operational Caveats

- `pgduck_server` is required for lake queries and must have working object-store credentials and local socket connectivity from PostgreSQL.
- S3 and compatible credentials are resolved by the DuckDB secrets/credential chain. Grant only the bucket permissions required by the workload.
- Iceberg writes create Parquet files per statement. Batch inserts and run regular `VACUUM` to avoid many small files.
- The PostgreSQL extensions, `pgduck_server`, object-store data, and Iceberg catalog form one deployment unit. Back up and upgrade them as separate evidence layers; creating the extension alone does not prove the external services are usable.
