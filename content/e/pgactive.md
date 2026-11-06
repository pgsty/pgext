---
title: "pgactive"
linkTitle: "pgactive"
description: "Active-Active Replication Extension for PostgreSQL"
weight: 9560
categories: ["ETL"]
languages: ["C"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pgactive**](https://github.com/aws/pgactive) : Active-Active Replication Extension for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9560** | {{< badge content="pgactive" link="https://github.com/aws/pgactive" >}} | {{< ext "pgactive" >}} | `2.1.9` | {{< category "ETL" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "pglogical" >}} {{< ext "spock" >}} {{< ext "mimeo" >}} {{< ext "pgoutput" >}} {{< ext "postgres_fdw" >}} {{< ext "repmgr" >}} {{< ext "pg_failover_slots" >}} {{< ext "citus" >}} {{< ext "pgautofailover" >}} {{< ext "bgw_replstatus" >}} |

> [!Note] SQL 2.1.8.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.9` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgactive` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.9` | {{< bg "18" "pgactive_18" "green" >}} {{< bg "17" "pgactive_17" "green" >}} {{< bg "16" "pgactive_16" "green" >}} {{< bg "15" "pgactive_15" "green" >}} {{< bg "14" "pgactive_14" "green" >}} | `pgactive_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.9` | {{< bg "18" "postgresql-18-pgactive" "green" >}} {{< bg "17" "postgresql-17-pgactive" "green" >}} {{< bg "16" "postgresql-16-pgactive" "green" >}} {{< bg "15" "postgresql-15-pgactive" "green" >}} {{< bg "14" "postgresql-14-pgactive" "green" >}} | `postgresql-$v-pgactive` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "pgactive_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-18-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-14-pgactive : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-18-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-14-pgactive : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-18-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-14-pgactive : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-18-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-14-pgactive : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-18-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-14-pgactive : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-18-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-14-pgactive : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-18-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-14-pgactive : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-18-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-14-pgactive : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-18-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-14-pgactive : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-18-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-17-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-16-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-15-pgactive : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.9" "postgresql-14-pgactive : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgactive_18` | `2.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 642.6 KiB | [pgactive_18-2.1.9-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_18-2.1.9-1PGSTY.el8.x86_64.rpm) |
| `pgactive_18` | `2.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 623.1 KiB | [pgactive_18-2.1.9-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_18-2.1.9-1PGSTY.el8.aarch64.rpm) |
| `pgactive_18` | `2.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 634.4 KiB | [pgactive_18-2.1.9-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_18-2.1.9-1PGSTY.el9.x86_64.rpm) |
| `pgactive_18` | `2.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 623.1 KiB | [pgactive_18-2.1.9-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_18-2.1.9-1PGSTY.el9.aarch64.rpm) |
| `pgactive_18` | `2.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 638.9 KiB | [pgactive_18-2.1.9-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgactive_18-2.1.9-1PGSTY.el10.x86_64.rpm) |
| `pgactive_18` | `2.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 628.1 KiB | [pgactive_18-2.1.9-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgactive_18-2.1.9-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgactive` | `2.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 592.4 KiB | [postgresql-18-pgactive_2.1.9-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-18-pgactive_2.1.9-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgactive` | `2.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 564.9 KiB | [postgresql-18-pgactive_2.1.9-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-18-pgactive_2.1.9-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgactive` | `2.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 594.2 KiB | [postgresql-18-pgactive_2.1.9-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgactive/postgresql-18-pgactive_2.1.9-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pgactive` | `2.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 568.1 KiB | [postgresql-18-pgactive_2.1.9-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgactive/postgresql-18-pgactive_2.1.9-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pgactive` | `2.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 644.5 KiB | [postgresql-18-pgactive_2.1.9-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-18-pgactive_2.1.9-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pgactive` | `2.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 633.0 KiB | [postgresql-18-pgactive_2.1.9-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-18-pgactive_2.1.9-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pgactive` | `2.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 624.2 KiB | [postgresql-18-pgactive_2.1.9-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-18-pgactive_2.1.9-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pgactive` | `2.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 618.5 KiB | [postgresql-18-pgactive_2.1.9-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-18-pgactive_2.1.9-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pgactive` | `2.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 621.3 KiB | [postgresql-18-pgactive_2.1.9-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgactive/postgresql-18-pgactive_2.1.9-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pgactive` | `2.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 612.8 KiB | [postgresql-18-pgactive_2.1.9-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgactive/postgresql-18-pgactive_2.1.9-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgactive_17` | `2.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 634.3 KiB | [pgactive_17-2.1.9-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_17-2.1.9-1PGSTY.el8.x86_64.rpm) |
| `pgactive_17` | `2.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 615.4 KiB | [pgactive_17-2.1.9-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_17-2.1.9-1PGSTY.el8.aarch64.rpm) |
| `pgactive_17` | `2.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 626.0 KiB | [pgactive_17-2.1.9-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_17-2.1.9-1PGSTY.el9.x86_64.rpm) |
| `pgactive_17` | `2.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 615.7 KiB | [pgactive_17-2.1.9-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_17-2.1.9-1PGSTY.el9.aarch64.rpm) |
| `pgactive_17` | `2.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 630.4 KiB | [pgactive_17-2.1.9-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgactive_17-2.1.9-1PGSTY.el10.x86_64.rpm) |
| `pgactive_17` | `2.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 620.2 KiB | [pgactive_17-2.1.9-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgactive_17-2.1.9-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgactive` | `2.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 585.4 KiB | [postgresql-17-pgactive_2.1.9-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-17-pgactive_2.1.9-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgactive` | `2.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 559.7 KiB | [postgresql-17-pgactive_2.1.9-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-17-pgactive_2.1.9-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgactive` | `2.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 586.8 KiB | [postgresql-17-pgactive_2.1.9-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgactive/postgresql-17-pgactive_2.1.9-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pgactive` | `2.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 561.5 KiB | [postgresql-17-pgactive_2.1.9-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgactive/postgresql-17-pgactive_2.1.9-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pgactive` | `2.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 710.6 KiB | [postgresql-17-pgactive_2.1.9-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-17-pgactive_2.1.9-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pgactive` | `2.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 701.2 KiB | [postgresql-17-pgactive_2.1.9-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-17-pgactive_2.1.9-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pgactive` | `2.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 617.2 KiB | [postgresql-17-pgactive_2.1.9-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-17-pgactive_2.1.9-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pgactive` | `2.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 610.8 KiB | [postgresql-17-pgactive_2.1.9-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-17-pgactive_2.1.9-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pgactive` | `2.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 613.5 KiB | [postgresql-17-pgactive_2.1.9-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgactive/postgresql-17-pgactive_2.1.9-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pgactive` | `2.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 605.3 KiB | [postgresql-17-pgactive_2.1.9-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgactive/postgresql-17-pgactive_2.1.9-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgactive_16` | `2.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 628.8 KiB | [pgactive_16-2.1.9-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_16-2.1.9-1PGSTY.el8.x86_64.rpm) |
| `pgactive_16` | `2.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 609.6 KiB | [pgactive_16-2.1.9-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_16-2.1.9-1PGSTY.el8.aarch64.rpm) |
| `pgactive_16` | `2.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 620.9 KiB | [pgactive_16-2.1.9-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_16-2.1.9-1PGSTY.el9.x86_64.rpm) |
| `pgactive_16` | `2.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 610.2 KiB | [pgactive_16-2.1.9-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_16-2.1.9-1PGSTY.el9.aarch64.rpm) |
| `pgactive_16` | `2.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 625.8 KiB | [pgactive_16-2.1.9-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgactive_16-2.1.9-1PGSTY.el10.x86_64.rpm) |
| `pgactive_16` | `2.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 615.1 KiB | [pgactive_16-2.1.9-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgactive_16-2.1.9-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgactive` | `2.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 580.3 KiB | [postgresql-16-pgactive_2.1.9-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-16-pgactive_2.1.9-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgactive` | `2.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 554.3 KiB | [postgresql-16-pgactive_2.1.9-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-16-pgactive_2.1.9-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgactive` | `2.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 581.8 KiB | [postgresql-16-pgactive_2.1.9-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgactive/postgresql-16-pgactive_2.1.9-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pgactive` | `2.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 556.4 KiB | [postgresql-16-pgactive_2.1.9-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgactive/postgresql-16-pgactive_2.1.9-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pgactive` | `2.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 699.9 KiB | [postgresql-16-pgactive_2.1.9-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-16-pgactive_2.1.9-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pgactive` | `2.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 690.5 KiB | [postgresql-16-pgactive_2.1.9-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-16-pgactive_2.1.9-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pgactive` | `2.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 612.6 KiB | [postgresql-16-pgactive_2.1.9-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-16-pgactive_2.1.9-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pgactive` | `2.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 611.5 KiB | [postgresql-16-pgactive_2.1.9-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-16-pgactive_2.1.9-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pgactive` | `2.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 609.5 KiB | [postgresql-16-pgactive_2.1.9-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgactive/postgresql-16-pgactive_2.1.9-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pgactive` | `2.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 600.3 KiB | [postgresql-16-pgactive_2.1.9-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgactive/postgresql-16-pgactive_2.1.9-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgactive_15` | `2.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 621.9 KiB | [pgactive_15-2.1.9-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_15-2.1.9-1PGSTY.el8.x86_64.rpm) |
| `pgactive_15` | `2.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 602.0 KiB | [pgactive_15-2.1.9-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_15-2.1.9-1PGSTY.el8.aarch64.rpm) |
| `pgactive_15` | `2.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 617.7 KiB | [pgactive_15-2.1.9-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_15-2.1.9-1PGSTY.el9.x86_64.rpm) |
| `pgactive_15` | `2.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 606.7 KiB | [pgactive_15-2.1.9-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_15-2.1.9-1PGSTY.el9.aarch64.rpm) |
| `pgactive_15` | `2.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 622.1 KiB | [pgactive_15-2.1.9-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgactive_15-2.1.9-1PGSTY.el10.x86_64.rpm) |
| `pgactive_15` | `2.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 611.4 KiB | [pgactive_15-2.1.9-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgactive_15-2.1.9-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgactive` | `2.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 576.3 KiB | [postgresql-15-pgactive_2.1.9-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-15-pgactive_2.1.9-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgactive` | `2.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 629.0 KiB | [postgresql-15-pgactive_2.1.9-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-15-pgactive_2.1.9-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgactive` | `2.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 577.4 KiB | [postgresql-15-pgactive_2.1.9-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgactive/postgresql-15-pgactive_2.1.9-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pgactive` | `2.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 551.7 KiB | [postgresql-15-pgactive_2.1.9-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgactive/postgresql-15-pgactive_2.1.9-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pgactive` | `2.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 697.5 KiB | [postgresql-15-pgactive_2.1.9-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-15-pgactive_2.1.9-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pgactive` | `2.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 689.0 KiB | [postgresql-15-pgactive_2.1.9-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-15-pgactive_2.1.9-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pgactive` | `2.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 609.0 KiB | [postgresql-15-pgactive_2.1.9-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-15-pgactive_2.1.9-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pgactive` | `2.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 604.0 KiB | [postgresql-15-pgactive_2.1.9-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-15-pgactive_2.1.9-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pgactive` | `2.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 606.3 KiB | [postgresql-15-pgactive_2.1.9-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgactive/postgresql-15-pgactive_2.1.9-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pgactive` | `2.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 598.4 KiB | [postgresql-15-pgactive_2.1.9-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgactive/postgresql-15-pgactive_2.1.9-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgactive_14` | `2.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 622.5 KiB | [pgactive_14-2.1.9-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgactive_14-2.1.9-1PGSTY.el8.x86_64.rpm) |
| `pgactive_14` | `2.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 603.3 KiB | [pgactive_14-2.1.9-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgactive_14-2.1.9-1PGSTY.el8.aarch64.rpm) |
| `pgactive_14` | `2.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 621.5 KiB | [pgactive_14-2.1.9-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgactive_14-2.1.9-1PGSTY.el9.x86_64.rpm) |
| `pgactive_14` | `2.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 609.1 KiB | [pgactive_14-2.1.9-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgactive_14-2.1.9-1PGSTY.el9.aarch64.rpm) |
| `pgactive_14` | `2.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 627.3 KiB | [pgactive_14-2.1.9-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgactive_14-2.1.9-1PGSTY.el10.x86_64.rpm) |
| `pgactive_14` | `2.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 614.3 KiB | [pgactive_14-2.1.9-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgactive_14-2.1.9-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgactive` | `2.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 579.4 KiB | [postgresql-14-pgactive_2.1.9-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-14-pgactive_2.1.9-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgactive` | `2.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 552.3 KiB | [postgresql-14-pgactive_2.1.9-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgactive/postgresql-14-pgactive_2.1.9-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgactive` | `2.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 581.0 KiB | [postgresql-14-pgactive_2.1.9-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgactive/postgresql-14-pgactive_2.1.9-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pgactive` | `2.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 554.0 KiB | [postgresql-14-pgactive_2.1.9-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgactive/postgresql-14-pgactive_2.1.9-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pgactive` | `2.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 699.2 KiB | [postgresql-14-pgactive_2.1.9-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-14-pgactive_2.1.9-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pgactive` | `2.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 691.3 KiB | [postgresql-14-pgactive_2.1.9-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgactive/postgresql-14-pgactive_2.1.9-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pgactive` | `2.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 612.6 KiB | [postgresql-14-pgactive_2.1.9-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-14-pgactive_2.1.9-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pgactive` | `2.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 605.4 KiB | [postgresql-14-pgactive_2.1.9-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgactive/postgresql-14-pgactive_2.1.9-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pgactive` | `2.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 607.8 KiB | [postgresql-14-pgactive_2.1.9-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgactive/postgresql-14-pgactive_2.1.9-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pgactive` | `2.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 599.5 KiB | [postgresql-14-pgactive_2.1.9-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgactive/postgresql-14-pgactive_2.1.9-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aws/pgactive" title="Repository" icon="github" subtitle="github.com/aws/pgactive" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgactive-2.1.9.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgactive;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgactive;		# install via package name, for the active PG version

pig install pgactive -v 18;   # install for PG 18
pig install pgactive -v 17;   # install for PG 17
pig install pgactive -v 16;   # install for PG 16
pig install pgactive -v 15;   # install for PG 15
pig install pgactive -v 14;   # install for PG 14

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgactive';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pgactive;
```




## Usage

> [pgactive: Active-Active Replication Extension for PostgreSQL](https://github.com/aws/pgactive)

The `pgactive` extension provides active-active (multi-master) replication for PostgreSQL, where multiple database instances can accept writes and replicate changes to each other.

### Enabling

```sql
CREATE EXTENSION pgactive;
```

### Overview

pgactive enables active-active replication using logical replication as its foundation. Multiple databases in a cluster can accept changes and replicate them bidirectionally.

### Key Concepts

- **Active-Active**: All nodes accept reads and writes simultaneously
- **Asynchronous**: Changes are replicated asynchronously between nodes
- **Conflict Resolution**: Applications must handle conflicting changes from multiple writers
- **Logical Replication**: Uses PostgreSQL's logical decoding to interpret and apply changes

### Use Cases

- Multi-region high availability database clusters
- Reducing write latency between applications and databases
- Blue/green application updates
- Data migration between systems that must both remain writable

### Design Considerations

Applications using pgactive must be designed to handle:

- **Write conflicts**: Concurrent modifications to the same row on different nodes
- **Replication lag**: Changes may not be immediately visible on all nodes
- **Feature limitations**: Some features like auto-incrementing sequences require special handling across nodes

### Notes

- Built on PostgreSQL's native logical replication infrastructure
- Requires PostgreSQL 10+ (native logical replication support)
- Refer to the project documentation for detailed setup and conflict resolution strategies
