---
title: "pgmq"
linkTitle: "pgmq"
description: "A lightweight message queue. Like AWS SQS and RSMQ but on Postgres."
weight: 2890
categories: ["FEAT"]
width: full
---

[**pgmq**](https://github.com/pgmq/pgmq) : A lightweight message queue. Like AWS SQS and RSMQ but on Postgres.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2890** | {{< badge content="pgmq" link="https://github.com/pgmq/pgmq" >}} | {{< ext "pgmq" >}} | `1.8.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pg_later" >}} {{< ext "vectorize" >}} |
|   **See Also**    | {{< ext "kafka_fdw" >}} {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_net" >}} {{< ext "pg_background" >}} {{< ext "pgagent" >}} {{< ext "pg_jobmon" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.8.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pgmq` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.8.0` | {{< bg "18" "pgmq_18" "green" >}} {{< bg "17" "pgmq_17" "green" >}} {{< bg "16" "pgmq_16" "green" >}} {{< bg "15" "pgmq_15" "green" >}} {{< bg "14" "pgmq_14" "green" >}} {{< bg "13" "pgmq_13" "green" >}} | `pgmq_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.8.0` | {{< bg "18" "postgresql-18-pgmq" "green" >}} {{< bg "17" "postgresql-17-pgmq" "green" >}} {{< bg "16" "postgresql-16-pgmq" "green" >}} {{< bg "15" "postgresql-15-pgmq" "green" >}} {{< bg "14" "postgresql-14-pgmq" "green" >}} {{< bg "13" "postgresql-13-pgmq" "green" >}} | `postgresql-$v-pgmq` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "pgmq_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-13-pgmq : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-13-pgmq : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-13-pgmq : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-13-pgmq : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-13-pgmq : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-13-pgmq : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-13-pgmq : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-18-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-14-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.8.0" "postgresql-13-pgmq : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_18` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.8 KiB | [pgmq_18-1.8.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_18-1.8.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_18` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.8 KiB | [pgmq_18-1.8.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_18-1.8.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_18` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [pgmq_18-1.8.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_18-1.8.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_18` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.8 KiB | [pgmq_18-1.8.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_18-1.8.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_18` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [pgmq_18-1.8.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_18-1.8.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_18` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.0 KiB | [pgmq_18-1.8.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_18-1.8.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgmq` | `1.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.7 KiB | [postgresql-18-pgmq_1.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-18-pgmq_1.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgmq` | `1.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.7 KiB | [postgresql-18-pgmq_1.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-18-pgmq_1.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgmq` | `1.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.7 KiB | [postgresql-18-pgmq_1.8.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-18-pgmq_1.8.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgmq` | `1.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.7 KiB | [postgresql-18-pgmq_1.8.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-18-pgmq_1.8.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgmq` | `1.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.1 KiB | [postgresql-18-pgmq_1.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-18-pgmq_1.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgmq` | `1.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.1 KiB | [postgresql-18-pgmq_1.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-18-pgmq_1.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgmq` | `1.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.1 KiB | [postgresql-18-pgmq_1.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-18-pgmq_1.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgmq` | `1.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.1 KiB | [postgresql-18-pgmq_1.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-18-pgmq_1.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_17` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.8 KiB | [pgmq_17-1.8.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_17-1.8.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_17` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.8 KiB | [pgmq_17-1.8.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_17-1.8.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_17` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [pgmq_17-1.8.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_17-1.8.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_17` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.8 KiB | [pgmq_17-1.8.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_17-1.8.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_17` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [pgmq_17-1.8.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_17-1.8.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_17` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.9 KiB | [pgmq_17-1.8.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_17-1.8.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgmq` | `1.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.7 KiB | [postgresql-17-pgmq_1.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-17-pgmq_1.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgmq` | `1.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.7 KiB | [postgresql-17-pgmq_1.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-17-pgmq_1.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgmq` | `1.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.7 KiB | [postgresql-17-pgmq_1.8.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-17-pgmq_1.8.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgmq` | `1.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.7 KiB | [postgresql-17-pgmq_1.8.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-17-pgmq_1.8.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgmq` | `1.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.1 KiB | [postgresql-17-pgmq_1.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-17-pgmq_1.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgmq` | `1.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.2 KiB | [postgresql-17-pgmq_1.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-17-pgmq_1.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgmq` | `1.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.1 KiB | [postgresql-17-pgmq_1.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-17-pgmq_1.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgmq` | `1.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.1 KiB | [postgresql-17-pgmq_1.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-17-pgmq_1.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_16` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.8 KiB | [pgmq_16-1.8.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_16-1.8.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_16` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.8 KiB | [pgmq_16-1.8.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_16-1.8.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_16` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [pgmq_16-1.8.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_16-1.8.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_16` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.8 KiB | [pgmq_16-1.8.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_16-1.8.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_16` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [pgmq_16-1.8.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_16-1.8.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_16` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.0 KiB | [pgmq_16-1.8.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_16-1.8.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgmq` | `1.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.7 KiB | [postgresql-16-pgmq_1.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-16-pgmq_1.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgmq` | `1.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.7 KiB | [postgresql-16-pgmq_1.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-16-pgmq_1.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgmq` | `1.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.7 KiB | [postgresql-16-pgmq_1.8.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-16-pgmq_1.8.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgmq` | `1.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.6 KiB | [postgresql-16-pgmq_1.8.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-16-pgmq_1.8.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgmq` | `1.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.2 KiB | [postgresql-16-pgmq_1.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-16-pgmq_1.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgmq` | `1.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.2 KiB | [postgresql-16-pgmq_1.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-16-pgmq_1.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgmq` | `1.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.1 KiB | [postgresql-16-pgmq_1.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-16-pgmq_1.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgmq` | `1.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.1 KiB | [postgresql-16-pgmq_1.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-16-pgmq_1.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_15` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.8 KiB | [pgmq_15-1.8.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_15-1.8.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_15` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.8 KiB | [pgmq_15-1.8.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_15-1.8.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_15` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [pgmq_15-1.8.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_15-1.8.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_15` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.8 KiB | [pgmq_15-1.8.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_15-1.8.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_15` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [pgmq_15-1.8.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_15-1.8.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_15` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.0 KiB | [pgmq_15-1.8.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_15-1.8.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgmq` | `1.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.7 KiB | [postgresql-15-pgmq_1.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-15-pgmq_1.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgmq` | `1.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.7 KiB | [postgresql-15-pgmq_1.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-15-pgmq_1.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgmq` | `1.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.7 KiB | [postgresql-15-pgmq_1.8.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-15-pgmq_1.8.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgmq` | `1.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.7 KiB | [postgresql-15-pgmq_1.8.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-15-pgmq_1.8.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgmq` | `1.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.2 KiB | [postgresql-15-pgmq_1.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-15-pgmq_1.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgmq` | `1.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.2 KiB | [postgresql-15-pgmq_1.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-15-pgmq_1.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgmq` | `1.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.1 KiB | [postgresql-15-pgmq_1.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-15-pgmq_1.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgmq` | `1.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.1 KiB | [postgresql-15-pgmq_1.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-15-pgmq_1.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_14` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.8 KiB | [pgmq_14-1.8.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_14-1.8.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_14` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.8 KiB | [pgmq_14-1.8.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_14-1.8.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_14` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [pgmq_14-1.8.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_14-1.8.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_14` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.8 KiB | [pgmq_14-1.8.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_14-1.8.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_14` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [pgmq_14-1.8.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_14-1.8.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_14` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.0 KiB | [pgmq_14-1.8.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_14-1.8.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgmq` | `1.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.7 KiB | [postgresql-14-pgmq_1.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-14-pgmq_1.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgmq` | `1.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.7 KiB | [postgresql-14-pgmq_1.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-14-pgmq_1.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgmq` | `1.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.7 KiB | [postgresql-14-pgmq_1.8.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-14-pgmq_1.8.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgmq` | `1.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.6 KiB | [postgresql-14-pgmq_1.8.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-14-pgmq_1.8.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgmq` | `1.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.1 KiB | [postgresql-14-pgmq_1.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-14-pgmq_1.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgmq` | `1.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.2 KiB | [postgresql-14-pgmq_1.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-14-pgmq_1.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgmq` | `1.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.1 KiB | [postgresql-14-pgmq_1.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-14-pgmq_1.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgmq` | `1.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.1 KiB | [postgresql-14-pgmq_1.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-14-pgmq_1.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_13` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.8 KiB | [pgmq_13-1.8.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_13-1.8.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_13` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.8 KiB | [pgmq_13-1.8.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_13-1.8.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_13` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [pgmq_13-1.8.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_13-1.8.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_13` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.8 KiB | [pgmq_13-1.8.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_13-1.8.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_13` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [pgmq_13-1.8.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmq_13-1.8.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmq_13` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.0 KiB | [pgmq_13-1.8.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmq_13-1.8.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pgmq` | `1.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.7 KiB | [postgresql-13-pgmq_1.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-13-pgmq_1.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgmq` | `1.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 20.7 KiB | [postgresql-13-pgmq_1.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-13-pgmq_1.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgmq` | `1.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.7 KiB | [postgresql-13-pgmq_1.8.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-13-pgmq_1.8.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pgmq` | `1.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.7 KiB | [postgresql-13-pgmq_1.8.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmq/postgresql-13-pgmq_1.8.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pgmq` | `1.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.2 KiB | [postgresql-13-pgmq_1.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-13-pgmq_1.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgmq` | `1.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.2 KiB | [postgresql-13-pgmq_1.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-13-pgmq_1.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgmq` | `1.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.1 KiB | [postgresql-13-pgmq_1.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-13-pgmq_1.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgmq` | `1.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.1 KiB | [postgresql-13-pgmq_1.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-13-pgmq_1.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgmq/pgmq" title="Repository" icon="github" subtitle="github.com/pgmq/pgmq" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmq-1.8.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgmq;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgmq;		# install via package name, for the active PG version

pig install pgmq -v 18;   # install for PG 18
pig install pgmq -v 17;   # install for PG 17
pig install pgmq -v 16;   # install for PG 16
pig install pgmq -v 15;   # install for PG 15
pig install pgmq -v 14;   # install for PG 14
pig install pgmq -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgmq;
```
