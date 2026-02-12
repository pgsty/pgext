---
title: "pgmb"
linkTitle: "pgmb"
description: "A simple PostgreSQL Message Broker system"
weight: 2870
categories: ["FEAT"]
width: full
---

[**pgmb**](https://github.com/fraruiz/pgmb) : A simple PostgreSQL Message Broker system


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2870** | {{< badge content="pgmb" link="https://github.com/fraruiz/pgmb" >}} | {{< ext "pgmb" >}} | `1.0.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgmb` |
|   **Requires**    | {{< ext "pg_cron" >}} {{< ext "http" >}} |
|   **See Also**    | {{< ext "pgmq" >}} {{< ext "pgq" >}} {{< ext "pg_task" >}} {{< ext "pg_cron" >}} {{< ext "pg_background" >}} {{< ext "pg_later" >}} {{< ext "pg_net" >}} {{< ext "kafka_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pgmb` | `pg_cron`, `http` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pgmb_18" "green" >}} {{< bg "17" "pgmb_17" "green" >}} {{< bg "16" "pgmb_16" "green" >}} {{< bg "15" "pgmb_15" "green" >}} {{< bg "14" "pgmb_14" "green" >}} {{< bg "13" "pgmb_13" "green" >}} | `pgmb_$v` | `pg_cron_$v`, `pg_http_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pgmb" "green" >}} {{< bg "17" "postgresql-17-pgmb" "green" >}} {{< bg "16" "postgresql-16-pgmb" "green" >}} {{< bg "15" "postgresql-15-pgmb" "green" >}} {{< bg "14" "postgresql-14-pgmb" "green" >}} {{< bg "13" "postgresql-13-pgmb" "green" >}} | `postgresql-$v-pgmb` | `postgresql-$v-cron`, `postgresql-$v-http` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmb_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmb : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmb : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmb : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmb : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmb : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmb : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmb : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmb : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmb_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.3 KiB | [pgmb_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmb_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.3 KiB | [pgmb_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmb_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.3 KiB | [pgmb_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmb_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.2 KiB | [pgmb_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmb_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 KiB | [pgmb_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmb_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.4 KiB | [pgmb_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgmb` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgmb` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-18-pgmb_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-18-pgmb_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmb_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.3 KiB | [pgmb_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmb_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.3 KiB | [pgmb_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmb_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.3 KiB | [pgmb_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmb_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.2 KiB | [pgmb_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmb_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 KiB | [pgmb_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmb_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.4 KiB | [pgmb_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgmb` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.8 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgmb` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-17-pgmb_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-17-pgmb_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmb_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.3 KiB | [pgmb_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmb_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.3 KiB | [pgmb_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmb_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.3 KiB | [pgmb_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmb_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.2 KiB | [pgmb_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmb_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 KiB | [pgmb_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmb_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.4 KiB | [pgmb_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgmb` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgmb` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-16-pgmb_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-16-pgmb_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmb_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.3 KiB | [pgmb_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmb_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.3 KiB | [pgmb_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmb_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.3 KiB | [pgmb_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmb_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.2 KiB | [pgmb_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmb_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 KiB | [pgmb_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmb_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.4 KiB | [pgmb_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgmb` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.8 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.8 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgmb` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-15-pgmb_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-15-pgmb_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmb_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.3 KiB | [pgmb_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmb_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.3 KiB | [pgmb_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmb_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.2 KiB | [pgmb_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmb_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.2 KiB | [pgmb_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmb_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 KiB | [pgmb_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmb_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.4 KiB | [pgmb_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgmb` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.8 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.8 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgmb` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-14-pgmb_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-14-pgmb_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmb_13` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.3 KiB | [pgmb_13-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmb_13-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmb_13` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.3 KiB | [pgmb_13-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmb_13-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmb_13` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.3 KiB | [pgmb_13-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmb_13-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmb_13` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.2 KiB | [pgmb_13-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmb_13-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmb_13` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.4 KiB | [pgmb_13-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmb_13-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmb_13` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.4 KiB | [pgmb_13-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmb_13-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pgmb` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.8 KiB | [postgresql-13-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-13-pgmb_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgmb` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.7 KiB | [postgresql-13-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmb/postgresql-13-pgmb_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgmb` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.7 KiB | [postgresql-13-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-13-pgmb_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pgmb` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.7 KiB | [postgresql-13-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmb/postgresql-13-pgmb_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pgmb` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-13-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-13-pgmb_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgmb` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-13-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmb/postgresql-13-pgmb_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgmb` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-13-pgmb_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-13-pgmb_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgmb` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-13-pgmb_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmb/postgresql-13-pgmb_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/fraruiz/pgmb" title="Repository" icon="github" subtitle="github.com/fraruiz/pgmb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmb-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgmb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgmb;		# install via package name, for the active PG version

pig install pgmb -v 18;   # install for PG 18
pig install pgmb -v 17;   # install for PG 17
pig install pgmb -v 16;   # install for PG 16
pig install pgmb -v 15;   # install for PG 15
pig install pgmb -v 14;   # install for PG 14
pig install pgmb -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgmb CASCADE; -- requires pg_cron, http
```
