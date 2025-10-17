---
title: "session_variable"
linkTitle: "session_variable"
description: "Registration and manipulation of session variables and constants"
weight: 9120
categories: ["Sim"]
width: full
---

Registration and manipulation of session variables and constants

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9120** | {{< badge content="session_variable" link="https://github.com/splendiddata/session_variable" >}} | {{< ext "session_variable" "session_variable" >}} | `3.4` | {{< category "SIM" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "orafce" >}} {{< ext "pgtt" >}} {{< ext "pg_statement_rollback" >}} {{< ext "plpgsql" >}} {{< ext "set_user" >}} {{< ext "oracle_fdw" >}} {{< ext "pg_dbms_lock" >}} {{< ext "babelfishpg_common" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/session_variable" >}} | `3.4` | {{< badge content="18" color="red" alt="session_variable_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `session_variable_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/session_variable" >}} | `3.4` | {{< badge content="18" color="red" alt="postgresql-18-session-variable" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-session-variable` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "session_variable_18" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/session_variable_18-3.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "session_variable_17" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/session_variable_17-3.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "session_variable_16" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/session_variable_16-3.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "session_variable_15" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/session_variable_15-3.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "session_variable_14" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/session_variable_14-3.4-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "session_variable_18" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/session_variable_18-3.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "session_variable_17" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/session_variable_17-3.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "session_variable_16" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/session_variable_16-3.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "session_variable_15" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/session_variable_15-3.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "session_variable_14" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/session_variable_14-3.4-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "session_variable_18" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/session_variable_18-3.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "session_variable_17" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/session_variable_17-3.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "session_variable_16" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/session_variable_16-3.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "session_variable_15" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/session_variable_15-3.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "session_variable_14" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/session_variable_14-3.4-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "session_variable_18" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/session_variable_18-3.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "session_variable_17" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/session_variable_17-3.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "session_variable_16" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/session_variable_16-3.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "session_variable_15" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/session_variable_15-3.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "session_variable_14" "3.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/session_variable_14-3.4-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-session-variable" >}}     | {{< pkg "postgresql-17-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-17-session-variable_3.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-16-session-variable_3.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-15-session-variable_3.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-14-session-variable_3.4-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-session-variable" >}}     | {{< pkg "postgresql-17-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-17-session-variable_3.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-16-session-variable_3.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-15-session-variable_3.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-14-session-variable_3.4-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-session-variable" >}}     | {{< pkg "postgresql-17-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-17-session-variable_3.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-16-session-variable_3.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-15-session-variable_3.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-14-session-variable_3.4-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-session-variable" >}}     | {{< pkg "postgresql-17-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-17-session-variable_3.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-16-session-variable_3.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-15-session-variable_3.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-14-session-variable_3.4-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-session-variable" >}}     | {{< pkg "postgresql-17-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-17-session-variable_3.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-16-session-variable_3.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-15-session-variable_3.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-14-session-variable_3.4-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-session-variable" >}}     | {{< pkg "postgresql-17-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-17-session-variable_3.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-16-session-variable_3.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-15-session-variable_3.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-session-variable" "3.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-14-session-variable_3.4-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `session_variable_18` | 3.4 | `el8.aarch64` | pigsty | 33.3 KiB | [session_variable_18-3.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/session_variable_18-3.4-1PIGSTY.el8.aarch64.rpm) |
| `session_variable_18` | 3.4 | `el8.x86_64` | pigsty | 34.1 KiB | [session_variable_18-3.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/session_variable_18-3.4-1PIGSTY.el8.x86_64.rpm) |
| `session_variable_18` | 3.4 | `el9.x86_64` | pigsty | 34.4 KiB | [session_variable_18-3.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/session_variable_18-3.4-1PIGSTY.el9.x86_64.rpm) |
| `session_variable_18` | 3.4 | `el9.aarch64` | pigsty | 33.3 KiB | [session_variable_18-3.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/session_variable_18-3.4-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `session_variable_17` | 3.4 | `el8.aarch64` | pigsty | 33.3 KiB | [session_variable_17-3.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/session_variable_17-3.4-1PIGSTY.el8.aarch64.rpm) |
| `session_variable_17` | 3.4 | `el8.x86_64` | pigsty | 34.1 KiB | [session_variable_17-3.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/session_variable_17-3.4-1PIGSTY.el8.x86_64.rpm) |
| `session_variable_17` | 3.4 | `el9.aarch64` | pigsty | 33.4 KiB | [session_variable_17-3.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/session_variable_17-3.4-1PIGSTY.el9.aarch64.rpm) |
| `session_variable_17` | 3.4 | `el9.x86_64` | pigsty | 34.8 KiB | [session_variable_17-3.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/session_variable_17-3.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-session-variable` | 3.4 | `d12.aarch64` | pigsty | 67.7 KiB | [postgresql-17-session-variable_3.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-17-session-variable_3.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-session-variable` | 3.4 | `d12.x86_64` | pigsty | 68.9 KiB | [postgresql-17-session-variable_3.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-17-session-variable_3.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-session-variable` | 3.4 | `u22.aarch64` | pigsty | 71.7 KiB | [postgresql-17-session-variable_3.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-17-session-variable_3.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-session-variable` | 3.4 | `u22.x86_64` | pigsty | 72.3 KiB | [postgresql-17-session-variable_3.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-17-session-variable_3.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-session-variable` | 3.4 | `u24.x86_64` | pigsty | 65.6 KiB | [postgresql-17-session-variable_3.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-17-session-variable_3.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-session-variable` | 3.4 | `u24.aarch64` | pigsty | 64.9 KiB | [postgresql-17-session-variable_3.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-17-session-variable_3.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `session_variable_16` | 3.4 | `el8.aarch64` | pigsty | 33.3 KiB | [session_variable_16-3.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/session_variable_16-3.4-1PIGSTY.el8.aarch64.rpm) |
| `session_variable_16` | 3.4 | `el8.x86_64` | pigsty | 34.1 KiB | [session_variable_16-3.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/session_variable_16-3.4-1PIGSTY.el8.x86_64.rpm) |
| `session_variable_16` | 3.4 | `el9.x86_64` | pigsty | 34.8 KiB | [session_variable_16-3.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/session_variable_16-3.4-1PIGSTY.el9.x86_64.rpm) |
| `session_variable_16` | 3.4 | `el9.aarch64` | pigsty | 33.5 KiB | [session_variable_16-3.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/session_variable_16-3.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-session-variable` | 3.4 | `d12.x86_64` | pigsty | 68.3 KiB | [postgresql-16-session-variable_3.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-16-session-variable_3.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-session-variable` | 3.4 | `d12.aarch64` | pigsty | 67.2 KiB | [postgresql-16-session-variable_3.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-16-session-variable_3.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-session-variable` | 3.4 | `u22.aarch64` | pigsty | 71.2 KiB | [postgresql-16-session-variable_3.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-16-session-variable_3.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-session-variable` | 3.4 | `u22.x86_64` | pigsty | 71.8 KiB | [postgresql-16-session-variable_3.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-16-session-variable_3.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-session-variable` | 3.4 | `u24.aarch64` | pigsty | 64.9 KiB | [postgresql-16-session-variable_3.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-16-session-variable_3.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-session-variable` | 3.4 | `u24.x86_64` | pigsty | 65.6 KiB | [postgresql-16-session-variable_3.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-16-session-variable_3.4-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `session_variable_15` | 3.4 | `el8.x86_64` | pigsty | 34.5 KiB | [session_variable_15-3.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/session_variable_15-3.4-1PIGSTY.el8.x86_64.rpm) |
| `session_variable_15` | 3.4 | `el8.aarch64` | pigsty | 33.7 KiB | [session_variable_15-3.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/session_variable_15-3.4-1PIGSTY.el8.aarch64.rpm) |
| `session_variable_15` | 3.4 | `el9.x86_64` | pigsty | 34.7 KiB | [session_variable_15-3.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/session_variable_15-3.4-1PIGSTY.el9.x86_64.rpm) |
| `session_variable_15` | 3.4 | `el9.aarch64` | pigsty | 33.8 KiB | [session_variable_15-3.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/session_variable_15-3.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-session-variable` | 3.4 | `d12.x86_64` | pigsty | 68.5 KiB | [postgresql-15-session-variable_3.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-15-session-variable_3.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-session-variable` | 3.4 | `d12.aarch64` | pigsty | 67.4 KiB | [postgresql-15-session-variable_3.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-15-session-variable_3.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-session-variable` | 3.4 | `u22.x86_64` | pigsty | 71.9 KiB | [postgresql-15-session-variable_3.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-15-session-variable_3.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-session-variable` | 3.4 | `u22.aarch64` | pigsty | 71.5 KiB | [postgresql-15-session-variable_3.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-15-session-variable_3.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-session-variable` | 3.4 | `u24.aarch64` | pigsty | 65.1 KiB | [postgresql-15-session-variable_3.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-15-session-variable_3.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-session-variable` | 3.4 | `u24.x86_64` | pigsty | 65.7 KiB | [postgresql-15-session-variable_3.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-15-session-variable_3.4-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `session_variable_14` | 3.4 | `el8.x86_64` | pigsty | 34.5 KiB | [session_variable_14-3.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/session_variable_14-3.4-1PIGSTY.el8.x86_64.rpm) |
| `session_variable_14` | 3.4 | `el8.aarch64` | pigsty | 33.6 KiB | [session_variable_14-3.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/session_variable_14-3.4-1PIGSTY.el8.aarch64.rpm) |
| `session_variable_14` | 3.4 | `el9.aarch64` | pigsty | 33.8 KiB | [session_variable_14-3.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/session_variable_14-3.4-1PIGSTY.el9.aarch64.rpm) |
| `session_variable_14` | 3.4 | `el9.x86_64` | pigsty | 34.7 KiB | [session_variable_14-3.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/session_variable_14-3.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-14-session-variable` | 3.4 | `d12.x86_64` | pigsty | 67.1 KiB | [postgresql-14-session-variable_3.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-14-session-variable_3.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-session-variable` | 3.4 | `d12.aarch64` | pigsty | 66.0 KiB | [postgresql-14-session-variable_3.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-14-session-variable_3.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-session-variable` | 3.4 | `u22.x86_64` | pigsty | 70.6 KiB | [postgresql-14-session-variable_3.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-14-session-variable_3.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-session-variable` | 3.4 | `u22.aarch64` | pigsty | 70.2 KiB | [postgresql-14-session-variable_3.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-14-session-variable_3.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-session-variable` | 3.4 | `u24.aarch64` | pigsty | 65.2 KiB | [postgresql-14-session-variable_3.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-14-session-variable_3.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-session-variable` | 3.4 | `u24.x86_64` | pigsty | 65.7 KiB | [postgresql-14-session-variable_3.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-14-session-variable_3.4-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `session_variable_13` | 3.4 | `el8.x86_64` | pigsty | 34.1 KiB | [session_variable_13-3.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/session_variable_13-3.4-1PIGSTY.el8.x86_64.rpm) |
| `session_variable_13` | 3.4 | `el8.aarch64` | pigsty | 33.6 KiB | [session_variable_13-3.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/session_variable_13-3.4-1PIGSTY.el8.aarch64.rpm) |
| `session_variable_13` | 3.4 | `el9.aarch64` | pigsty | 33.8 KiB | [session_variable_13-3.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/session_variable_13-3.4-1PIGSTY.el9.aarch64.rpm) |
| `session_variable_13` | 3.4 | `el9.x86_64` | pigsty | 34.5 KiB | [session_variable_13-3.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/session_variable_13-3.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-session-variable` | 3.4 | `d12.x86_64` | pigsty | 66.7 KiB | [postgresql-13-session-variable_3.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-13-session-variable_3.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-session-variable` | 3.4 | `d12.aarch64` | pigsty | 65.8 KiB | [postgresql-13-session-variable_3.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/session-variable/postgresql-13-session-variable_3.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-session-variable` | 3.4 | `u22.aarch64` | pigsty | 69.8 KiB | [postgresql-13-session-variable_3.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-13-session-variable_3.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-session-variable` | 3.4 | `u22.x86_64` | pigsty | 70.4 KiB | [postgresql-13-session-variable_3.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/session-variable/postgresql-13-session-variable_3.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-session-variable` | 3.4 | `u24.x86_64` | pigsty | 65.4 KiB | [postgresql-13-session-variable_3.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-13-session-variable_3.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-session-variable` | 3.4 | `u24.aarch64` | pigsty | 65.0 KiB | [postgresql-13-session-variable_3.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/session-variable/postgresql-13-session-variable_3.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/splendiddata/session_variable" title="Repository" icon="github" subtitle="github.com/splendiddata/session_variable" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="session_variable-3.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get session_variable; # get session_variable source code
pig build dep session_variable; # install build dependencies
pig build pkg session_variable; # build extension rpm or deb
pig build ext session_variable; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install session_variable; # install by extension name, for the current active PG version
pig ext install session_variable; # install via package alias, for the active PG version
pig ext install session_variable -v 18;   # install for PG 18
pig ext install session_variable -v 17;   # install for PG 17
pig ext install session_variable -v 16;   # install for PG 16
pig ext install session_variable -v 15;   # install for PG 15
pig ext install session_variable -v 14;   # install for PG 14
pig ext install session_variable -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION session_variable;
```

