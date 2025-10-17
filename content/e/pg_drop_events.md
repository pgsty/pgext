---
title: "pg_drop_events"
linkTitle: "pg_drop_events"
description: "logs transaction ids of drop table, drop column, drop materialized view statements"
weight: 5830
categories: ["Admin"]
width: full
---

logs transaction ids of drop table, drop column, drop materialized view statements

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5830** | {{< badge content="pg_drop_events" link="https://github.com/bolajiwahab/pg_drop_events" >}} | {{< ext "pg_drop_events" "pg_drop_events" >}} | `0.1.0` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pg_savior" >}} {{< ext "table_log" >}} {{< ext "pgaudit" >}} {{< ext "pg_auditor" >}} {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "pg_upless" >}} {{< ext "pgauditlogtofile" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_drop_events" >}} | `0.1.0` | {{< badge content="18" color="red" alt="pg_drop_events_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_drop_events_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_drop_events" >}} | `0.1.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-drop-events" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-drop-events` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_drop_events_18" >}}     | {{< pkg "pg_drop_events_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_drop_events_17-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_drop_events_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_drop_events_16-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_drop_events_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_drop_events_15-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_drop_events_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_drop_events_14-0.1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_drop_events_18" >}}     | {{< pkg "pg_drop_events_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_drop_events_17-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_drop_events_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_drop_events_16-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_drop_events_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_drop_events_15-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_drop_events_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_drop_events_14-0.1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_drop_events_18" >}}     | {{< pkg "pg_drop_events_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_drop_events_17-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_drop_events_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_drop_events_16-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_drop_events_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_drop_events_15-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_drop_events_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_drop_events_14-0.1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_drop_events_18" >}}     | {{< pkg "pg_drop_events_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_drop_events_17-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_drop_events_16" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_drop_events_16-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_drop_events_15" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_drop_events_15-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_drop_events_14" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_drop_events_14-0.1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-drop-events" >}}     | {{< pkg "postgresql-17-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-17-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-16-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-15-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-14-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-drop-events" >}}     | {{< pkg "postgresql-17-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-17-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-16-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-15-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-14-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-drop-events" >}}     | {{< pkg "postgresql-17-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-17-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-16-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-15-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-14-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-drop-events" >}}     | {{< pkg "postgresql-17-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-17-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-16-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-15-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-14-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-drop-events" >}}     | {{< pkg "postgresql-17-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-17-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-16-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-15-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-14-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-drop-events" >}}     | {{< pkg "postgresql-17-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-17-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-16-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-15-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-drop-events" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-14-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_drop_events_17` | 0.1.0 | `el8.x86_64` | pigsty | 10.9 KiB | [pg_drop_events_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_drop_events_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_drop_events_17` | 0.1.0 | `el8.aarch64` | pigsty | 10.9 KiB | [pg_drop_events_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_drop_events_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_drop_events_17` | 0.1.0 | `el9.aarch64` | pigsty | 10.8 KiB | [pg_drop_events_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_drop_events_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_drop_events_17` | 0.1.0 | `el9.x86_64` | pigsty | 10.9 KiB | [pg_drop_events_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_drop_events_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-drop-events` | 0.1.0 | `d12.x86_64` | pigsty | 7.8 KiB | [postgresql-17-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-17-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-drop-events` | 0.1.0 | `d12.aarch64` | pigsty | 7.8 KiB | [postgresql-17-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-17-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-drop-events` | 0.1.0 | `u22.x86_64` | pigsty | 7.8 KiB | [postgresql-17-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-17-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-drop-events` | 0.1.0 | `u22.aarch64` | pigsty | 7.8 KiB | [postgresql-17-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-17-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-drop-events` | 0.1.0 | `u24.x86_64` | pigsty | 7.8 KiB | [postgresql-17-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-17-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-drop-events` | 0.1.0 | `u24.aarch64` | pigsty | 7.8 KiB | [postgresql-17-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-17-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_drop_events_16` | 0.1.0 | `el8.x86_64` | pigsty | 10.9 KiB | [pg_drop_events_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_drop_events_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_drop_events_16` | 0.1.0 | `el8.aarch64` | pigsty | 10.9 KiB | [pg_drop_events_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_drop_events_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_drop_events_16` | 0.1.0 | `el9.x86_64` | pigsty | 10.9 KiB | [pg_drop_events_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_drop_events_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_drop_events_16` | 0.1.0 | `el9.aarch64` | pigsty | 10.8 KiB | [pg_drop_events_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_drop_events_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-drop-events` | 0.1.0 | `d12.x86_64` | pigsty | 7.8 KiB | [postgresql-16-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-16-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-drop-events` | 0.1.0 | `d12.aarch64` | pigsty | 7.8 KiB | [postgresql-16-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-16-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-drop-events` | 0.1.0 | `u22.aarch64` | pigsty | 7.8 KiB | [postgresql-16-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-16-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-drop-events` | 0.1.0 | `u22.x86_64` | pigsty | 7.8 KiB | [postgresql-16-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-16-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-drop-events` | 0.1.0 | `u24.x86_64` | pigsty | 7.8 KiB | [postgresql-16-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-16-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-drop-events` | 0.1.0 | `u24.aarch64` | pigsty | 7.8 KiB | [postgresql-16-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-16-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_drop_events_15` | 0.1.0 | `el8.x86_64` | pigsty | 10.9 KiB | [pg_drop_events_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_drop_events_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_drop_events_15` | 0.1.0 | `el8.aarch64` | pigsty | 10.9 KiB | [pg_drop_events_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_drop_events_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_drop_events_15` | 0.1.0 | `el9.x86_64` | pigsty | 10.9 KiB | [pg_drop_events_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_drop_events_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_drop_events_15` | 0.1.0 | `el9.aarch64` | pigsty | 10.8 KiB | [pg_drop_events_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_drop_events_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-drop-events` | 0.1.0 | `d12.aarch64` | pigsty | 7.8 KiB | [postgresql-15-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-15-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-drop-events` | 0.1.0 | `d12.x86_64` | pigsty | 7.8 KiB | [postgresql-15-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-15-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-drop-events` | 0.1.0 | `u22.aarch64` | pigsty | 7.8 KiB | [postgresql-15-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-15-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-drop-events` | 0.1.0 | `u22.x86_64` | pigsty | 7.8 KiB | [postgresql-15-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-15-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-drop-events` | 0.1.0 | `u24.x86_64` | pigsty | 7.8 KiB | [postgresql-15-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-15-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-drop-events` | 0.1.0 | `u24.aarch64` | pigsty | 7.8 KiB | [postgresql-15-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-15-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_drop_events_14` | 0.1.0 | `el8.x86_64` | pigsty | 10.9 KiB | [pg_drop_events_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_drop_events_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_drop_events_14` | 0.1.0 | `el8.aarch64` | pigsty | 10.9 KiB | [pg_drop_events_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_drop_events_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_drop_events_14` | 0.1.0 | `el9.x86_64` | pigsty | 10.9 KiB | [pg_drop_events_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_drop_events_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_drop_events_14` | 0.1.0 | `el9.aarch64` | pigsty | 10.8 KiB | [pg_drop_events_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_drop_events_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-drop-events` | 0.1.0 | `d12.x86_64` | pigsty | 7.8 KiB | [postgresql-14-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-14-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-drop-events` | 0.1.0 | `d12.aarch64` | pigsty | 7.8 KiB | [postgresql-14-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-14-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-drop-events` | 0.1.0 | `u22.x86_64` | pigsty | 7.8 KiB | [postgresql-14-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-14-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-drop-events` | 0.1.0 | `u22.aarch64` | pigsty | 7.8 KiB | [postgresql-14-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-14-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-drop-events` | 0.1.0 | `u24.x86_64` | pigsty | 7.8 KiB | [postgresql-14-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-14-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-drop-events` | 0.1.0 | `u24.aarch64` | pigsty | 7.8 KiB | [postgresql-14-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-14-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_drop_events_13` | 0.1.0 | `el8.aarch64` | pigsty | 10.9 KiB | [pg_drop_events_13-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_drop_events_13-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_drop_events_13` | 0.1.0 | `el8.x86_64` | pigsty | 10.9 KiB | [pg_drop_events_13-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_drop_events_13-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_drop_events_13` | 0.1.0 | `el9.aarch64` | pigsty | 10.8 KiB | [pg_drop_events_13-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_drop_events_13-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_drop_events_13` | 0.1.0 | `el9.x86_64` | pigsty | 10.9 KiB | [pg_drop_events_13-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_drop_events_13-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-drop-events` | 0.1.0 | `d12.aarch64` | pigsty | 7.8 KiB | [postgresql-13-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-13-pg-drop-events_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-drop-events` | 0.1.0 | `d12.x86_64` | pigsty | 7.8 KiB | [postgresql-13-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-drop-events/postgresql-13-pg-drop-events_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-drop-events` | 0.1.0 | `u22.aarch64` | pigsty | 7.8 KiB | [postgresql-13-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-13-pg-drop-events_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-drop-events` | 0.1.0 | `u22.x86_64` | pigsty | 7.8 KiB | [postgresql-13-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-drop-events/postgresql-13-pg-drop-events_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-drop-events` | 0.1.0 | `u24.aarch64` | pigsty | 7.8 KiB | [postgresql-13-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-13-pg-drop-events_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-drop-events` | 0.1.0 | `u24.x86_64` | pigsty | 7.8 KiB | [postgresql-13-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-drop-events/postgresql-13-pg-drop-events_0.1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bolajiwahab/pg_drop_events" title="Repository" icon="github" subtitle="github.com/bolajiwahab/pg_drop_events" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_drop_events-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_drop_events; # get pg_drop_events source code
pig build dep pg_drop_events; # install build dependencies
pig build pkg pg_drop_events; # build extension rpm or deb
pig build ext pg_drop_events; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_drop_events; # install by extension name, for the current active PG version
pig ext install pg_drop_events; # install via package alias, for the active PG version
pig ext install pg_drop_events -v 18;   # install for PG 18
pig ext install pg_drop_events -v 17;   # install for PG 17
pig ext install pg_drop_events -v 16;   # install for PG 16
pig ext install pg_drop_events -v 15;   # install for PG 15
pig ext install pg_drop_events -v 14;   # install for PG 14
pig ext install pg_drop_events -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_drop_events CASCADE SCHEMA public;
```

