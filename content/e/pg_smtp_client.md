---
title: "pg_smtp_client"
linkTitle: "pg_smtp_client"
description: "PostgreSQL extension to send email using SMTP"
weight: 4170
categories: ["UTIL"]
width: full
---

PostgreSQL extension to send email using SMTP


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4170** | {{< badge content="pg_smtp_client" link="https://github.com/brianpursley/pg_smtp_client" >}} | {{< ext "pg_smtp_client" >}} | `0.2.0` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_html5_email_address" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "pg_curl" >}} {{< ext "pgjq" >}} |

> [!Note] pgrx=0.12.7


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_smtp_client" >}} | `0.2.0` | {{< bg "18" "pg_smtp_client_18" "red" >}} {{< bg "17" "pg_smtp_client_17" "green" >}} {{< bg "16" "pg_smtp_client_16" "green" >}} {{< bg "15" "pg_smtp_client_15" "green" >}} {{< bg "14" "pg_smtp_client_14" "green" >}} | `pg_smtp_client_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_smtp_client" >}} | `0.2.0` | {{< bg "18" "postgresql-18-pg-smtp-client" "red" >}} {{< bg "17" "postgresql-17-pg-smtp-client" "green" >}} {{< bg "16" "postgresql-16-pg-smtp-client" "green" >}} {{< bg "15" "postgresql-15-pg-smtp-client" "green" >}} {{< bg "14" "postgresql-14-pg-smtp-client" "green" >}} | `postgresql-$v-pg-smtp-client` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_smtp_client_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_smtp_client_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_smtp_client_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_smtp_client_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_smtp_client_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-smtp-client : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-smtp-client : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-smtp-client : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-smtp-client : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-smtp-client : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-smtp-client : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-smtp-client : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-smtp-client : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-smtp-client : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-smtp-client : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-smtp-client : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-smtp-client : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-smtp-client : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_smtp_client_17` | 0.2.0 | `el8.x86_64` | pigsty | 422.5 KiB | [pg_smtp_client_17-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_smtp_client_17-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_smtp_client_17` | 0.2.0 | `el8.aarch64` | pigsty | 390.1 KiB | [pg_smtp_client_17-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_smtp_client_17-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_smtp_client_17` | 0.2.0 | `el9.x86_64` | pigsty | 427.8 KiB | [pg_smtp_client_17-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_smtp_client_17-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_smtp_client_17` | 0.2.0 | `el9.aarch64` | pigsty | 415.5 KiB | [pg_smtp_client_17-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_smtp_client_17-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-smtp-client` | 0.2.0 | `d12.x86_64` | pigsty | 345.0 KiB | [postgresql-17-pg-smtp-client_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-smtp-client/postgresql-17-pg-smtp-client_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-smtp-client` | 0.2.0 | `d12.aarch64` | pigsty | 305.7 KiB | [postgresql-17-pg-smtp-client_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-smtp-client/postgresql-17-pg-smtp-client_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-smtp-client` | 0.2.0 | `u22.x86_64` | pigsty | 380.0 KiB | [postgresql-17-pg-smtp-client_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-smtp-client/postgresql-17-pg-smtp-client_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-smtp-client` | 0.2.0 | `u22.aarch64` | pigsty | 357.9 KiB | [postgresql-17-pg-smtp-client_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-smtp-client/postgresql-17-pg-smtp-client_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-smtp-client` | 0.2.0 | `u24.x86_64` | pigsty | 376.0 KiB | [postgresql-17-pg-smtp-client_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-smtp-client/postgresql-17-pg-smtp-client_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-smtp-client` | 0.2.0 | `u24.aarch64` | pigsty | 355.0 KiB | [postgresql-17-pg-smtp-client_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-smtp-client/postgresql-17-pg-smtp-client_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_smtp_client_16` | 0.2.0 | `el8.x86_64` | pigsty | 422.4 KiB | [pg_smtp_client_16-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_smtp_client_16-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_smtp_client_16` | 0.2.0 | `el8.aarch64` | pigsty | 390.1 KiB | [pg_smtp_client_16-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_smtp_client_16-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_smtp_client_16` | 0.2.0 | `el9.x86_64` | pigsty | 427.6 KiB | [pg_smtp_client_16-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_smtp_client_16-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_smtp_client_16` | 0.2.0 | `el9.aarch64` | pigsty | 415.5 KiB | [pg_smtp_client_16-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_smtp_client_16-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-smtp-client` | 0.2.0 | `d12.x86_64` | pigsty | 345.0 KiB | [postgresql-16-pg-smtp-client_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-smtp-client/postgresql-16-pg-smtp-client_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-smtp-client` | 0.2.0 | `d12.aarch64` | pigsty | 305.8 KiB | [postgresql-16-pg-smtp-client_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-smtp-client/postgresql-16-pg-smtp-client_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-smtp-client` | 0.2.0 | `u22.x86_64` | pigsty | 380.4 KiB | [postgresql-16-pg-smtp-client_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-smtp-client/postgresql-16-pg-smtp-client_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-smtp-client` | 0.2.0 | `u22.aarch64` | pigsty | 357.8 KiB | [postgresql-16-pg-smtp-client_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-smtp-client/postgresql-16-pg-smtp-client_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-smtp-client` | 0.2.0 | `u24.x86_64` | pigsty | 376.3 KiB | [postgresql-16-pg-smtp-client_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-smtp-client/postgresql-16-pg-smtp-client_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-smtp-client` | 0.2.0 | `u24.aarch64` | pigsty | 354.5 KiB | [postgresql-16-pg-smtp-client_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-smtp-client/postgresql-16-pg-smtp-client_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_smtp_client_15` | 0.2.0 | `el8.x86_64` | pigsty | 422.3 KiB | [pg_smtp_client_15-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_smtp_client_15-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_smtp_client_15` | 0.2.0 | `el8.aarch64` | pigsty | 390.0 KiB | [pg_smtp_client_15-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_smtp_client_15-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_smtp_client_15` | 0.2.0 | `el9.x86_64` | pigsty | 428.3 KiB | [pg_smtp_client_15-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_smtp_client_15-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_smtp_client_15` | 0.2.0 | `el9.aarch64` | pigsty | 415.6 KiB | [pg_smtp_client_15-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_smtp_client_15-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-smtp-client` | 0.2.0 | `d12.x86_64` | pigsty | 345.1 KiB | [postgresql-15-pg-smtp-client_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-smtp-client/postgresql-15-pg-smtp-client_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-smtp-client` | 0.2.0 | `d12.aarch64` | pigsty | 305.5 KiB | [postgresql-15-pg-smtp-client_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-smtp-client/postgresql-15-pg-smtp-client_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-smtp-client` | 0.2.0 | `u22.x86_64` | pigsty | 380.4 KiB | [postgresql-15-pg-smtp-client_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-smtp-client/postgresql-15-pg-smtp-client_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-smtp-client` | 0.2.0 | `u22.aarch64` | pigsty | 357.9 KiB | [postgresql-15-pg-smtp-client_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-smtp-client/postgresql-15-pg-smtp-client_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-smtp-client` | 0.2.0 | `u24.x86_64` | pigsty | 375.9 KiB | [postgresql-15-pg-smtp-client_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-smtp-client/postgresql-15-pg-smtp-client_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-smtp-client` | 0.2.0 | `u24.aarch64` | pigsty | 354.5 KiB | [postgresql-15-pg-smtp-client_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-smtp-client/postgresql-15-pg-smtp-client_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_smtp_client_14` | 0.2.0 | `el8.x86_64` | pigsty | 422.3 KiB | [pg_smtp_client_14-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_smtp_client_14-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_smtp_client_14` | 0.2.0 | `el8.aarch64` | pigsty | 390.0 KiB | [pg_smtp_client_14-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_smtp_client_14-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_smtp_client_14` | 0.2.0 | `el9.x86_64` | pigsty | 427.6 KiB | [pg_smtp_client_14-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_smtp_client_14-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_smtp_client_14` | 0.2.0 | `el9.aarch64` | pigsty | 415.5 KiB | [pg_smtp_client_14-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_smtp_client_14-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-smtp-client` | 0.2.0 | `d12.x86_64` | pigsty | 344.8 KiB | [postgresql-14-pg-smtp-client_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-smtp-client/postgresql-14-pg-smtp-client_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-smtp-client` | 0.2.0 | `d12.aarch64` | pigsty | 305.8 KiB | [postgresql-14-pg-smtp-client_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-smtp-client/postgresql-14-pg-smtp-client_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-smtp-client` | 0.2.0 | `u22.x86_64` | pigsty | 380.1 KiB | [postgresql-14-pg-smtp-client_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-smtp-client/postgresql-14-pg-smtp-client_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-smtp-client` | 0.2.0 | `u22.aarch64` | pigsty | 357.7 KiB | [postgresql-14-pg-smtp-client_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-smtp-client/postgresql-14-pg-smtp-client_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-smtp-client` | 0.2.0 | `u24.x86_64` | pigsty | 376.3 KiB | [postgresql-14-pg-smtp-client_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-smtp-client/postgresql-14-pg-smtp-client_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-smtp-client` | 0.2.0 | `u24.aarch64` | pigsty | 354.4 KiB | [postgresql-14-pg-smtp-client_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-smtp-client/postgresql-14-pg-smtp-client_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/brianpursley/pg_smtp_client" title="Repository" icon="github" subtitle="github.com/brianpursley/pg_smtp_client" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_smtp_client-0.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_smtp_client; # get pg_smtp_client source code
pig build dep pg_smtp_client; # install build dependencies
pig build pkg pg_smtp_client; # build extension rpm or deb
pig build ext pg_smtp_client; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_smtp_client; # install by extension name, for the current active PG version
pig ext install pg_smtp_client; # install via package alias, for the active PG version
pig ext install pg_smtp_client -v 17;   # install for PG 17
pig ext install pg_smtp_client -v 16;   # install for PG 16
pig ext install pg_smtp_client -v 15;   # install for PG 15
pig ext install pg_smtp_client -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_smtp_client CASCADE SCHEMA smtp_client;
```



--------

## Usage

> https://github.com/frectonz/pglite-fusion/blob/main/README.md

### Enabling the extension

Connect to postgres and run the following command.

```sql
CREATE EXTENSION IF NOT EXISTS pg_smtp_client CASCADE;
```

## Usage

Use the `smtp_client.send_email()` function to send an email.

### Function Parameters

| Parameter     | Type    | Description                                           | System Configuration (Optional) |
|---------------|---------|-------------------------------------------------------|---------------------------------|
| subject       | text    | The subject of the email                              |                                 |
| body          | text    | The body of the email                                 |                                 |
| html          | boolean | Whether the body is HTML (true) or plain text (false) |                                 |
| from_address  | text    | The from email address                                | `smtp_client.from_address`      |
| recipients    | text[]  | The email addresses of the recipients                 |                                 |
| ccs           | text[]  | The email addresses to CCs                            |                                 |
| bccs          | text[]  | The email addresses to BCCs                           |                                 |
| smtp_server   | text    | The SMTP server to use                                | `smtp_client.server`            |
| smtp_port     | integer | The port of the SMTP server                           | `smtp_client.port`              |
| smtp_tls      | boolean | Whether to use TLS                                    | `smtp_client.tls`               |
| smtp_username | text    | The username for the SMTP server                      | `smtp_client.username`          |
| smtp_password | text    | The password for the SMTP server                      | `smtp_client.password`          |

### Default Configuration

You can configure the following system-wide default values for some of the parameters (as indiciated in the table above) like this:

```
ALTER SYSTEM SET smtp_client.server TO 'smtp.example.com';
ALTER SYSTEM SET smtp_client.port TO 587;
ALTER SYSTEM SET smtp_client.tls TO true;
ALTER SYSTEM SET smtp_client.username TO 'MySmtpUsername';
ALTER SYSTEM SET smtp_client.password TO 'MySmtpPassword';
ALTER SYSTEM SET smtp_client.from_address TO 'from@example.com';
SELECT pg_reload_conf();
```

### Usage Examples

Send an email:
```sql
SELECT smtp_client.send_email('test subject', 'test body', false, 'from@example.com', array['to@example.com'], null, null, 'smtp.example.com', 587, true, 'username', 'password');
```

Send an email using configured default values:
```sql
SELECT smtp_client.send_email('test subject', 'test body', false, null, array['to@example.com']);
```

Or, using named arguments:
```sql
SELECT smtp_client.send_email('test subject', 'test body', recipients => array['to@example.com']);
```