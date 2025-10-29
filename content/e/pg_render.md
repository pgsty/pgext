---
title: "pg_render"
linkTitle: "pg_render"
description: "Render HTML in SQL"
weight: 4290
categories: ["UTIL"]
width: full
---

Render HTML in SQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4290** | {{< badge content="pg_render" link="https://github.com/mkaski/pg_render" >}} | {{< ext "pg_render" >}} | `0.1.3` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_html5_email_address" >}} {{< ext "pg_readme" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} |

> [!Note] pgrx=0.16.1, manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_render" >}} | `0.1.3` | {{< bg "18" "pg_render_18" "green" >}} {{< bg "17" "pg_render_17" "green" >}} {{< bg "16" "pg_render_16" "green" >}} {{< bg "15" "pg_render_15" "green" >}} {{< bg "14" "pg_render_14" "green" >}} {{< bg "13" "pg_render_13" "red" >}} | `pg_render_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_render" >}} | `0.1.3` | {{< bg "18" "postgresql-18-pg-render" "green" >}} {{< bg "17" "postgresql-17-pg-render" "green" >}} {{< bg "16" "postgresql-16-pg-render" "green" >}} {{< bg "15" "postgresql-15-pg-render" "green" >}} {{< bg "14" "postgresql-14-pg-render" "green" >}} {{< bg "13" "postgresql-13-pg-render" "red" >}} | `postgresql-$v-pg-render` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_render_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "pg_render_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_render_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_render_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_render_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_render_13 : MISS 0" "red" >}}      |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_render_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "pg_render_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_render_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_render_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_render_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_render_13 : MISS 0" "red" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_render_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "pg_render_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_render_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_render_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_render_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_render_13 : MISS 0" "red" >}}      |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_render_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "pg_render_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_render_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_render_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "pg_render_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_render_13 : MISS 0" "red" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "pg_render_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_render_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_render_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_render_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_render_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_render_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pg_render_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_render_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_render_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_render_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_render_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_render_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-render : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "postgresql-17-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-16-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-15-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-14-pg-render : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-render : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-render : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "postgresql-17-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-16-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-15-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-14-pg-render : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-render : MISS 0" "red" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-render : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-render : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-render : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-render : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-render : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-render : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-render : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-render : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-render : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-render : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-render : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-render : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-render : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "postgresql-17-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-16-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-15-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-14-pg-render : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-render : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-render : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "postgresql-17-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-16-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-15-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-14-pg-render : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-render : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-render : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "postgresql-17-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-16-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-15-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-14-pg-render : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-render : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-render : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.2" "postgresql-17-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-16-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-15-pg-render : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.2" "postgresql-14-pg-render : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-render : MISS 0" "red" >}}      |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_render_17` | 0.1.2 | `el8.x86_64` | pigsty | 1.0 MiB | [pg_render_17-0.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_render_17-0.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_render_17` | 0.1.2 | `el8.aarch64` | pigsty | 923.4 KiB | [pg_render_17-0.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_render_17-0.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_render_17` | 0.1.2 | `el9.x86_64` | pigsty | 1.0 MiB | [pg_render_17-0.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_render_17-0.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_render_17` | 0.1.2 | `el9.aarch64` | pigsty | 990.6 KiB | [pg_render_17-0.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_render_17-0.1.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-render` | 0.1.2 | `d12.x86_64` | pigsty | 848.9 KiB | [postgresql-17-pg-render_0.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-render/postgresql-17-pg-render_0.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-render` | 0.1.2 | `d12.aarch64` | pigsty | 723.8 KiB | [postgresql-17-pg-render_0.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-render/postgresql-17-pg-render_0.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-render` | 0.1.2 | `u22.x86_64` | pigsty | 952.2 KiB | [postgresql-17-pg-render_0.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-render/postgresql-17-pg-render_0.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-render` | 0.1.2 | `u22.aarch64` | pigsty | 870.8 KiB | [postgresql-17-pg-render_0.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-render/postgresql-17-pg-render_0.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-render` | 0.1.2 | `u24.x86_64` | pigsty | 944.6 KiB | [postgresql-17-pg-render_0.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-render/postgresql-17-pg-render_0.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-render` | 0.1.2 | `u24.aarch64` | pigsty | 854.8 KiB | [postgresql-17-pg-render_0.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-render/postgresql-17-pg-render_0.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_render_16` | 0.1.2 | `el8.x86_64` | pigsty | 1.0 MiB | [pg_render_16-0.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_render_16-0.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_render_16` | 0.1.2 | `el8.aarch64` | pigsty | 923.3 KiB | [pg_render_16-0.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_render_16-0.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_render_16` | 0.1.2 | `el9.x86_64` | pigsty | 1.0 MiB | [pg_render_16-0.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_render_16-0.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_render_16` | 0.1.2 | `el9.aarch64` | pigsty | 994.3 KiB | [pg_render_16-0.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_render_16-0.1.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-render` | 0.1.2 | `d12.x86_64` | pigsty | 849.1 KiB | [postgresql-16-pg-render_0.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-render/postgresql-16-pg-render_0.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-render` | 0.1.2 | `d12.aarch64` | pigsty | 723.0 KiB | [postgresql-16-pg-render_0.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-render/postgresql-16-pg-render_0.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-render` | 0.1.2 | `u22.x86_64` | pigsty | 952.3 KiB | [postgresql-16-pg-render_0.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-render/postgresql-16-pg-render_0.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-render` | 0.1.2 | `u22.aarch64` | pigsty | 870.8 KiB | [postgresql-16-pg-render_0.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-render/postgresql-16-pg-render_0.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-render` | 0.1.2 | `u24.x86_64` | pigsty | 940.4 KiB | [postgresql-16-pg-render_0.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-render/postgresql-16-pg-render_0.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-render` | 0.1.2 | `u24.aarch64` | pigsty | 858.3 KiB | [postgresql-16-pg-render_0.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-render/postgresql-16-pg-render_0.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_render_15` | 0.1.2 | `el8.x86_64` | pigsty | 1.0 MiB | [pg_render_15-0.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_render_15-0.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_render_15` | 0.1.2 | `el8.aarch64` | pigsty | 923.3 KiB | [pg_render_15-0.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_render_15-0.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_render_15` | 0.1.2 | `el9.x86_64` | pigsty | 1.0 MiB | [pg_render_15-0.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_render_15-0.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_render_15` | 0.1.2 | `el9.aarch64` | pigsty | 990.3 KiB | [pg_render_15-0.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_render_15-0.1.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-render` | 0.1.2 | `d12.x86_64` | pigsty | 848.8 KiB | [postgresql-15-pg-render_0.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-render/postgresql-15-pg-render_0.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-render` | 0.1.2 | `d12.aarch64` | pigsty | 723.3 KiB | [postgresql-15-pg-render_0.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-render/postgresql-15-pg-render_0.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-render` | 0.1.2 | `u22.x86_64` | pigsty | 952.4 KiB | [postgresql-15-pg-render_0.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-render/postgresql-15-pg-render_0.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-render` | 0.1.2 | `u22.aarch64` | pigsty | 870.1 KiB | [postgresql-15-pg-render_0.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-render/postgresql-15-pg-render_0.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-render` | 0.1.2 | `u24.x86_64` | pigsty | 944.5 KiB | [postgresql-15-pg-render_0.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-render/postgresql-15-pg-render_0.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-render` | 0.1.2 | `u24.aarch64` | pigsty | 859.2 KiB | [postgresql-15-pg-render_0.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-render/postgresql-15-pg-render_0.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_render_14` | 0.1.2 | `el8.x86_64` | pigsty | 1.0 MiB | [pg_render_14-0.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_render_14-0.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_render_14` | 0.1.2 | `el8.aarch64` | pigsty | 923.1 KiB | [pg_render_14-0.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_render_14-0.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_render_14` | 0.1.2 | `el9.x86_64` | pigsty | 1.0 MiB | [pg_render_14-0.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_render_14-0.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_render_14` | 0.1.2 | `el9.aarch64` | pigsty | 994.9 KiB | [pg_render_14-0.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_render_14-0.1.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-render` | 0.1.2 | `d12.x86_64` | pigsty | 848.7 KiB | [postgresql-14-pg-render_0.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-render/postgresql-14-pg-render_0.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-render` | 0.1.2 | `d12.aarch64` | pigsty | 724.0 KiB | [postgresql-14-pg-render_0.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-render/postgresql-14-pg-render_0.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-render` | 0.1.2 | `u22.x86_64` | pigsty | 951.2 KiB | [postgresql-14-pg-render_0.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-render/postgresql-14-pg-render_0.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-render` | 0.1.2 | `u22.aarch64` | pigsty | 869.6 KiB | [postgresql-14-pg-render_0.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-render/postgresql-14-pg-render_0.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-render` | 0.1.2 | `u24.x86_64` | pigsty | 944.9 KiB | [postgresql-14-pg-render_0.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-render/postgresql-14-pg-render_0.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-render` | 0.1.2 | `u24.aarch64` | pigsty | 858.3 KiB | [postgresql-14-pg-render_0.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-render/postgresql-14-pg-render_0.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/mkaski/pg_render" title="Repository" icon="github" subtitle="github.com/mkaski/pg_render" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_render-0.1.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_render; # get pg_render source code
pig build dep pg_render; # install build dependencies
pig build pkg pg_render; # build extension rpm or deb
pig build ext pg_render; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_render; # install by extension name, for the current active PG version
pig ext install pg_render; # install via package alias, for the active PG version
pig ext install pg_render -v 18;   # install for PG 18
pig ext install pg_render -v 17;   # install for PG 17
pig ext install pg_render -v 16;   # install for PG 16
pig ext install pg_render -v 15;   # install for PG 15
pig ext install pg_render -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_render;
```

