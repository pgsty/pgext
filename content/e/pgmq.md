---
title: "pgmq"
linkTitle: "pgmq"
description: "A lightweight message queue. Like AWS SQS and RSMQ but on Postgres."
weight: 2900
categories: ["FEAT"]
width: full
---

A lightweight message queue. Like AWS SQS and RSMQ but on Postgres.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2900** | {{< badge content="pgmq" link="https://github.com/pgmq/pgmq" >}} | {{< ext "pgmq" >}} | `1.5.1` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pg_later" >}} {{< ext "vectorize" >}} |
|   **See Also**    | {{< ext "kafka_fdw" >}} {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_net" >}} {{< ext "pg_background" >}} {{< ext "pgagent" >}} {{< ext "pg_jobmon" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgmq" >}} | `1.5.1` | {{< bg "18" "pgmq_18" "green" >}} {{< bg "17" "pgmq_17" "green" >}} {{< bg "16" "pgmq_16" "green" >}} {{< bg "15" "pgmq_15" "green" >}} {{< bg "14" "pgmq_14" "green" >}} | `pgmq_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgmq" >}} | `1.5.1` | {{< bg "18" "postgresql-18-pgmq" "green" >}} {{< bg "17" "postgresql-17-pgmq" "green" >}} {{< bg "16" "postgresql-16-pgmq" "green" >}} {{< bg "15" "postgresql-15-pgmq" "green" >}} {{< bg "14" "postgresql-14-pgmq" "green" >}} | `postgresql-$v-pgmq` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.5.1" "pgmq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_14 : AVAIL 2" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.5.1" "pgmq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.5.1" "pgmq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.5.1" "pgmq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.5.1" "pgmq_14 : AVAIL 2" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pgmq : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.5.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pgmq : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.5.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pgmq : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.5.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pgmq : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.5.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pgmq : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.5.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pgmq : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.5.1" "postgresql-17-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-16-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-15-pgmq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-14-pgmq : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_18` | 1.5.1 | `el8.x86_64` | pigsty | 27.6 KiB | [pgmq_18-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_18-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_18` | 1.5.1 | `el8.aarch64` | pigsty | 27.6 KiB | [pgmq_18-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_18-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_18` | 1.5.1 | `el9.x86_64` | pigsty | 26.8 KiB | [pgmq_18-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_18-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_18` | 1.5.1 | `el9.aarch64` | pigsty | 26.8 KiB | [pgmq_18-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_18-1.5.1-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_17` | 1.5.1 | `el8.x86_64` | pigsty | 27.7 KiB | [pgmq_17-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_17-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_17` | 1.4.4 | `el8.x86_64` | pigsty | 29.6 KiB | [pgmq_17-1.4.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_17-1.4.4-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_17` | 1.5.1 | `el8.aarch64` | pigsty | 27.6 KiB | [pgmq_17-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_17-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_17` | 1.5.1 | `el9.x86_64` | pigsty | 26.8 KiB | [pgmq_17-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_17-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_17` | 1.5.1 | `el9.aarch64` | pigsty | 26.8 KiB | [pgmq_17-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_17-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_17` | 1.4.4 | `el9.aarch64` | pigsty | 28.9 KiB | [pgmq_17-1.4.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_17-1.4.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pgmq` | 1.5.1 | `d12.x86_64` | pigsty | 18.3 KiB | [postgresql-17-pgmq_1.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-17-pgmq_1.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgmq` | 1.5.1 | `d12.aarch64` | pigsty | 18.3 KiB | [postgresql-17-pgmq_1.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-17-pgmq_1.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgmq` | 1.5.1 | `u22.x86_64` | pigsty | 18.8 KiB | [postgresql-17-pgmq_1.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-17-pgmq_1.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgmq` | 1.5.1 | `u22.aarch64` | pigsty | 18.8 KiB | [postgresql-17-pgmq_1.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-17-pgmq_1.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgmq` | 1.5.1 | `u24.x86_64` | pigsty | 18.7 KiB | [postgresql-17-pgmq_1.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-17-pgmq_1.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgmq` | 1.5.1 | `u24.aarch64` | pigsty | 18.7 KiB | [postgresql-17-pgmq_1.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-17-pgmq_1.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_16` | 1.5.1 | `el8.x86_64` | pigsty | 27.7 KiB | [pgmq_16-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_16-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_16` | 1.4.4 | `el8.x86_64` | pigsty | 29.6 KiB | [pgmq_16-1.4.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_16-1.4.4-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_16` | 1.5.1 | `el8.aarch64` | pigsty | 27.6 KiB | [pgmq_16-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_16-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_16` | 1.5.1 | `el9.x86_64` | pigsty | 26.8 KiB | [pgmq_16-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_16-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_16` | 1.5.1 | `el9.aarch64` | pigsty | 26.8 KiB | [pgmq_16-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_16-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_16` | 1.4.4 | `el9.aarch64` | pigsty | 28.9 KiB | [pgmq_16-1.4.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_16-1.4.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pgmq` | 1.5.1 | `d12.x86_64` | pigsty | 18.3 KiB | [postgresql-16-pgmq_1.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-16-pgmq_1.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgmq` | 1.5.1 | `d12.aarch64` | pigsty | 18.3 KiB | [postgresql-16-pgmq_1.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-16-pgmq_1.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgmq` | 1.5.1 | `u22.x86_64` | pigsty | 18.8 KiB | [postgresql-16-pgmq_1.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-16-pgmq_1.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgmq` | 1.5.1 | `u22.aarch64` | pigsty | 18.8 KiB | [postgresql-16-pgmq_1.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-16-pgmq_1.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgmq` | 1.5.1 | `u24.x86_64` | pigsty | 18.7 KiB | [postgresql-16-pgmq_1.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-16-pgmq_1.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgmq` | 1.5.1 | `u24.aarch64` | pigsty | 18.7 KiB | [postgresql-16-pgmq_1.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-16-pgmq_1.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_15` | 1.5.1 | `el8.x86_64` | pigsty | 27.7 KiB | [pgmq_15-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_15-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_15` | 1.4.4 | `el8.x86_64` | pigsty | 29.6 KiB | [pgmq_15-1.4.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_15-1.4.4-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_15` | 1.5.1 | `el8.aarch64` | pigsty | 27.6 KiB | [pgmq_15-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_15-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_15` | 1.5.1 | `el9.x86_64` | pigsty | 26.8 KiB | [pgmq_15-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_15-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_15` | 1.5.1 | `el9.aarch64` | pigsty | 26.8 KiB | [pgmq_15-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_15-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_15` | 1.4.4 | `el9.aarch64` | pigsty | 28.9 KiB | [pgmq_15-1.4.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_15-1.4.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgmq` | 1.5.1 | `d12.x86_64` | pigsty | 18.3 KiB | [postgresql-15-pgmq_1.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-15-pgmq_1.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgmq` | 1.5.1 | `d12.aarch64` | pigsty | 18.3 KiB | [postgresql-15-pgmq_1.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-15-pgmq_1.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgmq` | 1.5.1 | `u22.x86_64` | pigsty | 18.8 KiB | [postgresql-15-pgmq_1.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-15-pgmq_1.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgmq` | 1.5.1 | `u22.aarch64` | pigsty | 18.8 KiB | [postgresql-15-pgmq_1.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-15-pgmq_1.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgmq` | 1.5.1 | `u24.x86_64` | pigsty | 18.7 KiB | [postgresql-15-pgmq_1.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-15-pgmq_1.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgmq` | 1.5.1 | `u24.aarch64` | pigsty | 18.7 KiB | [postgresql-15-pgmq_1.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-15-pgmq_1.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmq_14` | 1.5.1 | `el8.x86_64` | pigsty | 27.7 KiB | [pgmq_14-1.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_14-1.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_14` | 1.4.4 | `el8.x86_64` | pigsty | 29.6 KiB | [pgmq_14-1.4.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmq_14-1.4.4-1PIGSTY.el8.x86_64.rpm) |
| `pgmq_14` | 1.5.1 | `el8.aarch64` | pigsty | 27.6 KiB | [pgmq_14-1.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmq_14-1.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmq_14` | 1.5.1 | `el9.x86_64` | pigsty | 26.8 KiB | [pgmq_14-1.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmq_14-1.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmq_14` | 1.5.1 | `el9.aarch64` | pigsty | 26.8 KiB | [pgmq_14-1.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_14-1.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmq_14` | 1.4.4 | `el9.aarch64` | pigsty | 28.9 KiB | [pgmq_14-1.4.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmq_14-1.4.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgmq` | 1.5.1 | `d12.x86_64` | pigsty | 18.3 KiB | [postgresql-14-pgmq_1.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-14-pgmq_1.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgmq` | 1.5.1 | `d12.aarch64` | pigsty | 18.3 KiB | [postgresql-14-pgmq_1.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmq/postgresql-14-pgmq_1.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgmq` | 1.5.1 | `u22.x86_64` | pigsty | 18.8 KiB | [postgresql-14-pgmq_1.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-14-pgmq_1.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgmq` | 1.5.1 | `u22.aarch64` | pigsty | 18.8 KiB | [postgresql-14-pgmq_1.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmq/postgresql-14-pgmq_1.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgmq` | 1.5.1 | `u24.x86_64` | pigsty | 18.7 KiB | [postgresql-14-pgmq_1.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-14-pgmq_1.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgmq` | 1.5.1 | `u24.aarch64` | pigsty | 18.7 KiB | [postgresql-14-pgmq_1.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmq/postgresql-14-pgmq_1.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgmq/pgmq" title="Repository" icon="github" subtitle="github.com/pgmq/pgmq" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmq-1.7.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgmq; # get pgmq source code
pig build dep pgmq; # install build dependencies
pig build pkg pgmq; # build extension rpm or deb
pig build ext pgmq; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgmq; # install by extension name, for the current active PG version
pig ext install pgmq; # install via package alias, for the active PG version
pig ext install pgmq -v 18;   # install for PG 18
pig ext install pgmq -v 17;   # install for PG 17
pig ext install pgmq -v 16;   # install for PG 16
pig ext install pgmq -v 15;   # install for PG 15
pig ext install pgmq -v 14;   # install for PG 14
pig ext install pgmq -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgmq CASCADE SCHEMA pgmq;
```

