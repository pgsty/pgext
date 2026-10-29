---
title: "pgmb"
linkTitle: "pgmb"
description: "A simple PostgreSQL Message Broker system"
weight: 2670
categories: ["FEAT"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pgmb**](https://github.com/fraruiz/pgmb) : A simple PostgreSQL Message Broker system


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2670** | {{< badge content="pgmb" link="https://github.com/fraruiz/pgmb" >}} | {{< ext "pgmb" >}} | `1.0.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgmb` |
|   **Requires**    | {{< ext "pg_cron" >}} {{< ext "http" >}} |
|   **See Also**    | {{< ext "pgmq" >}} {{< ext "redis" >}} {{< ext "ulak" >}} {{< ext "pgq" >}} {{< ext "pgmqtt" >}} {{< ext "redis_fdw" >}} {{< ext "tcn" >}} {{< ext "pg_durable" >}} {{< ext "fsm_core" >}} {{< ext "kafka_fdw" >}} |

> [!Note] Archived upstream. Requires pg_cron and http; preload pg_cron in its configured database. RPM 1.0.0-2PGSTY fixes schema creation and requires pgsql_http; existing DEBs retain the upstream schema-creation failure. Delay arguments do not delay dispatch; retries have no exponential backoff.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgmb` | `pg_cron`, `http` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pgmb_18" "green" >}} {{< bg "17" "pgmb_17" "green" >}} {{< bg "16" "pgmb_16" "green" >}} {{< bg "15" "pgmb_15" "green" >}} {{< bg "14" "pgmb_14" "green" >}} | `pgmb_$v` | `pg_cron_$v`, `pgsql_http_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pgmb" "green" >}} {{< bg "17" "postgresql-17-pgmb" "green" >}} {{< bg "16" "postgresql-16-pgmb" "green" >}} {{< bg "15" "postgresql-15-pgmb" "green" >}} {{< bg "14" "postgresql-14-pgmb" "green" >}} | `postgresql-$v-pgmb` | `postgresql-$v-cron`, `postgresql-$v-http` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_14 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_14 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_14 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_14 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmb_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.4 KiB | [pgmb_18-1.0.0-2PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_18-1.0.0-2PGSTY.el8.noarch.rpm) |
| `pgmb_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.3 KiB | [pgmb_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmb_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.4 KiB | [pgmb_18-1.0.0-2PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_18-1.0.0-2PGSTY.el8.noarch.rpm) |
| `pgmb_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.3 KiB | [pgmb_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmb_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.4 KiB | [pgmb_18-1.0.0-2PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_18-1.0.0-2PGSTY.el9.noarch.rpm) |
| `pgmb_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.3 KiB | [pgmb_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmb_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.4 KiB | [pgmb_18-1.0.0-2PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_18-1.0.0-2PGSTY.el9.noarch.rpm) |
| `pgmb_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.2 KiB | [pgmb_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmb_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.5 KiB | [pgmb_18-1.0.0-2PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_18-1.0.0-2PGSTY.el10.noarch.rpm) |
| `pgmb_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 KiB | [pgmb_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmb_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.5 KiB | [pgmb_18-1.0.0-2PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_18-1.0.0-2PGSTY.el10.noarch.rpm) |
| `pgmb_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.4 KiB | [pgmb_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgmb` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.4 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.4 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmb_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.4 KiB | [pgmb_17-1.0.0-2PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_17-1.0.0-2PGSTY.el8.noarch.rpm) |
| `pgmb_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.3 KiB | [pgmb_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmb_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.4 KiB | [pgmb_17-1.0.0-2PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_17-1.0.0-2PGSTY.el8.noarch.rpm) |
| `pgmb_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.3 KiB | [pgmb_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmb_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.4 KiB | [pgmb_17-1.0.0-2PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_17-1.0.0-2PGSTY.el9.noarch.rpm) |
| `pgmb_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.3 KiB | [pgmb_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmb_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.4 KiB | [pgmb_17-1.0.0-2PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_17-1.0.0-2PGSTY.el9.noarch.rpm) |
| `pgmb_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.2 KiB | [pgmb_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmb_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.5 KiB | [pgmb_17-1.0.0-2PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_17-1.0.0-2PGSTY.el10.noarch.rpm) |
| `pgmb_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 KiB | [pgmb_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmb_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.5 KiB | [pgmb_17-1.0.0-2PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_17-1.0.0-2PGSTY.el10.noarch.rpm) |
| `pgmb_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.4 KiB | [pgmb_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgmb` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.8 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.4 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.4 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmb_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.4 KiB | [pgmb_16-1.0.0-2PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_16-1.0.0-2PGSTY.el8.noarch.rpm) |
| `pgmb_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.3 KiB | [pgmb_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmb_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.4 KiB | [pgmb_16-1.0.0-2PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_16-1.0.0-2PGSTY.el8.noarch.rpm) |
| `pgmb_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.3 KiB | [pgmb_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmb_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.4 KiB | [pgmb_16-1.0.0-2PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_16-1.0.0-2PGSTY.el9.noarch.rpm) |
| `pgmb_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.3 KiB | [pgmb_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmb_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.4 KiB | [pgmb_16-1.0.0-2PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_16-1.0.0-2PGSTY.el9.noarch.rpm) |
| `pgmb_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.2 KiB | [pgmb_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmb_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.5 KiB | [pgmb_16-1.0.0-2PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_16-1.0.0-2PGSTY.el10.noarch.rpm) |
| `pgmb_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 KiB | [pgmb_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmb_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.5 KiB | [pgmb_16-1.0.0-2PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_16-1.0.0-2PGSTY.el10.noarch.rpm) |
| `pgmb_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.4 KiB | [pgmb_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgmb` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.4 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.4 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmb_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.4 KiB | [pgmb_15-1.0.0-2PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_15-1.0.0-2PGSTY.el8.noarch.rpm) |
| `pgmb_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.3 KiB | [pgmb_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmb_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.4 KiB | [pgmb_15-1.0.0-2PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_15-1.0.0-2PGSTY.el8.noarch.rpm) |
| `pgmb_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.3 KiB | [pgmb_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmb_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.4 KiB | [pgmb_15-1.0.0-2PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_15-1.0.0-2PGSTY.el9.noarch.rpm) |
| `pgmb_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.3 KiB | [pgmb_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmb_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.4 KiB | [pgmb_15-1.0.0-2PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_15-1.0.0-2PGSTY.el9.noarch.rpm) |
| `pgmb_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.2 KiB | [pgmb_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmb_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.5 KiB | [pgmb_15-1.0.0-2PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_15-1.0.0-2PGSTY.el10.noarch.rpm) |
| `pgmb_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 KiB | [pgmb_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmb_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.5 KiB | [pgmb_15-1.0.0-2PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_15-1.0.0-2PGSTY.el10.noarch.rpm) |
| `pgmb_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.4 KiB | [pgmb_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgmb` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.8 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.8 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.4 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.4 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmb_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.4 KiB | [pgmb_14-1.0.0-2PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_14-1.0.0-2PGSTY.el8.noarch.rpm) |
| `pgmb_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.3 KiB | [pgmb_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmb_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.4 KiB | [pgmb_14-1.0.0-2PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_14-1.0.0-2PGSTY.el8.noarch.rpm) |
| `pgmb_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.3 KiB | [pgmb_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmb_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.4 KiB | [pgmb_14-1.0.0-2PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_14-1.0.0-2PGSTY.el9.noarch.rpm) |
| `pgmb_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.2 KiB | [pgmb_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmb_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.4 KiB | [pgmb_14-1.0.0-2PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_14-1.0.0-2PGSTY.el9.noarch.rpm) |
| `pgmb_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.2 KiB | [pgmb_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmb_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.5 KiB | [pgmb_14-1.0.0-2PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_14-1.0.0-2PGSTY.el10.noarch.rpm) |
| `pgmb_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 KiB | [pgmb_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmb_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.5 KiB | [pgmb_14-1.0.0-2PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_14-1.0.0-2PGSTY.el10.noarch.rpm) |
| `pgmb_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.4 KiB | [pgmb_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgmb` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.8 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.8 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 7.4 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 7.4 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/fraruiz/pgmb" title="Repository" icon="github" subtitle="github.com/fraruiz/pgmb" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmb-1.0.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgmb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgmb;		# install via package name, for the active PG version

pig install pgmb -v 18;   # install for PG 18
pig install pgmb -v 17;   # install for PG 17
pig install pgmb -v 16;   # install for PG 16
pig install pgmb -v 15;   # install for PG 15
pig install pgmb -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pgmb CASCADE; -- requires pg_cron, http
```

## Usage

Sources:
- [README](https://github.com/fraruiz/pgmb/blob/release/README.md)
- [Control](https://github.com/fraruiz/pgmb/blob/release/pgmb.control)
- [SQL](https://github.com/fraruiz/pgmb/blob/release/pgmb--1.0.0.sql)
- [META.json](https://github.com/fraruiz/pgmb/blob/release/META.json)

`pgmb` implements a SQL message broker with worker endpoints, routing-key patterns, queues and HTTP dispatch scheduled through `pg_cron`. The upstream repository is archived.

### Installation and compatibility

The RPM revision `1.0.0-2PGSTY` requires `pgsql_http_N` and removes a redundant schema-creation statement. Earlier RPMs and existing DEB packages retain the upstream SQL and can fail at `CREATE EXTENSION pgmb` because the control file and install script both create the schema. Package installation alone does not establish that the extension can be created.

Install as a superuser. The required extensions are `pg_cron` and `http`; `pgmb` itself does not require preload. Configure the scheduler for the database in which the extension will be used:

```ini
shared_preload_libraries = 'pg_cron'
cron.database_name = 'postgres'
```

Restart PostgreSQL, connect to that database, and create the dependencies before the extension:

```sql
CREATE EXTENSION pg_cron;
CREATE EXTENSION http;
CREATE EXTENSION pgmb;
```

### Register a queue and send a message

The worker endpoint receives HTTP requests from the database server. Use a trusted endpoint and restrict who can register workers.

```sql
SELECT pgmb.create(
    'orders', 'order.*', 3,
    pgmb.worker('order_worker', 'http://127.0.0.1:8080/worker', 100)
);

SELECT pgmb.send(
    gen_random_uuid(), 'order.created', '{"order_id": 123}'::jsonb
);

SELECT * FROM pgmb.workers;
SELECT * FROM pgmb.queues;
SELECT * FROM pgmb.orders_queue;
SELECT * FROM pgmb.orders_dead_letter_queue;
```

The argument order for `pgmb.create` is name, binding pattern, maximum retries, then worker UUID. A trigger routes messages to matching queues, and queue creation registers a scheduler job.

### API

| Function | Purpose |
|---|---|
| `pgmb.worker(name, endpoint, rps)` | Register a worker and return its UUID; rps limits the batch size per dispatch. |
| `pgmb.create(name, binding_key, max_retries, worker_id)` | Create a queue, its dead-letter table and scheduler job; return the queue UUID. |
| `pgmb.send(id, routing_key, body)` | Store a JSONB message and route it to matching queues. |
| `pgmb.send` overloads | Accept optional headers and an integer-seconds or timestamp delay. |
| `pgmb.dispatch_messages(queue_id)` | Dispatch a batch; mark successful HTTP responses acknowledged, retry failures and move exhausted messages to the dead-letter table. |

### Limits of version 1.0.0

- Delays populate `enqueued_at`, but the dispatcher does not filter out future timestamps; the overloads do not guarantee delayed delivery.
- Headers are stored with messages, but the dispatcher sends only the body using `http_post`; they are not forwarded as custom HTTP headers.
- Retries use a counter on subsequent dispatches. The SQL does not implement exponential backoff.
- HTTP effects are external to the database transaction. Review delivery and retention requirements before relying on this archived project.
