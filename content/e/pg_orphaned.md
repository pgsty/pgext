---
title: "pg_orphaned"
linkTitle: "pg_orphaned"
description: "Deal with orphaned files"
weight: 5200
categories: ["ADMIN"]
width: full
---

Deal with orphaned files


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5200** | {{< badge content="pg_orphaned" link="https://github.com/bdrouvot/pg_orphaned" >}} | {{< ext "pg_orphaned" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_dirtyread" >}} {{< ext "pg_surgery" >}} {{< ext "amcheck" >}} {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pg_checksums" >}} {{< ext "pg_catcheck" >}} {{< ext "pg_repack" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_orphaned" >}} | `1.0` | {{< bg "18" "pg_orphaned_18*" "green" >}} {{< bg "17" "pg_orphaned_17*" "green" >}} {{< bg "16" "pg_orphaned_16*" "green" >}} {{< bg "15" "pg_orphaned_15*" "green" >}} {{< bg "14" "pg_orphaned_14*" "green" >}} {{< bg "13" "pg_orphaned_13*" "green" >}} | `pg_orphaned_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_orphaned" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-orphaned" "red" >}} {{< bg "17" "postgresql-17-pg-orphaned" "green" >}} {{< bg "16" "postgresql-16-pg-orphaned" "green" >}} {{< bg "15" "postgresql-15-pg-orphaned" "green" >}} {{< bg "14" "postgresql-14-pg-orphaned" "green" >}} {{< bg "13" "postgresql-13-pg-orphaned" "green" >}} | `postgresql-$v-pg-orphaned` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.0" "pg_orphaned_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.0" "pg_orphaned_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.0" "pg_orphaned_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.0" "pg_orphaned_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.0" "pg_orphaned_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.0" "pg_orphaned_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-orphaned : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-orphaned : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-orphaned : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-orphaned : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-orphaned : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-orphaned : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-orphaned : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-orphaned : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-orphaned : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-orphaned : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-orphaned : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-orphaned : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-orphaned : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-orphaned : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_orphaned_18` | 1.0 | `el8.x86_64` | pigsty | 21.3 KiB | [pg_orphaned_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_18` | 1.0 | `el8.aarch64` | pigsty | 21.1 KiB | [pg_orphaned_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_18` | 1.0 | `el9.x86_64` | pigsty | 21.4 KiB | [pg_orphaned_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_18` | 1.0 | `el9.aarch64` | pigsty | 21.4 KiB | [pg_orphaned_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_18` | 1.0 | `el10.x86_64` | pigsty | 21.4 KiB | [pg_orphaned_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_orphaned_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_orphaned_18` | 1.0 | `el10.aarch64` | pigsty | 21.5 KiB | [pg_orphaned_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_orphaned_18-1.0-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_orphaned_17` | 1.0 | `el8.x86_64` | pigsty | 21.3 KiB | [pg_orphaned_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_17` | 1.0 | `el8.aarch64` | pigsty | 21.1 KiB | [pg_orphaned_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_17` | 1.0 | `el9.x86_64` | pigsty | 21.3 KiB | [pg_orphaned_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_17` | 1.0 | `el9.aarch64` | pigsty | 21.4 KiB | [pg_orphaned_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_17` | 1.0 | `el10.x86_64` | pigsty | 21.4 KiB | [pg_orphaned_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_orphaned_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_orphaned_17` | 1.0 | `el10.aarch64` | pigsty | 21.5 KiB | [pg_orphaned_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_orphaned_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-orphaned` | 1.0 | `d12.x86_64` | pigsty | 33.4 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-orphaned` | 1.0 | `d12.aarch64` | pigsty | 33.4 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-orphaned` | 1.0 | `u22.x86_64` | pigsty | 34.8 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-orphaned` | 1.0 | `u22.aarch64` | pigsty | 34.9 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-orphaned` | 1.0 | `u24.x86_64` | pigsty | 30.5 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-orphaned` | 1.0 | `u24.aarch64` | pigsty | 30.5 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_orphaned_16` | 1.0 | `el8.x86_64` | pigsty | 21.3 KiB | [pg_orphaned_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_16` | 1.0 | `el8.aarch64` | pigsty | 21.1 KiB | [pg_orphaned_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_16` | 1.0 | `el9.x86_64` | pigsty | 21.3 KiB | [pg_orphaned_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_16` | 1.0 | `el9.aarch64` | pigsty | 21.4 KiB | [pg_orphaned_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_16` | 1.0 | `el10.x86_64` | pigsty | 21.4 KiB | [pg_orphaned_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_orphaned_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_orphaned_16` | 1.0 | `el10.aarch64` | pigsty | 21.5 KiB | [pg_orphaned_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_orphaned_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-orphaned` | 1.0 | `d12.x86_64` | pigsty | 33.0 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-orphaned` | 1.0 | `d12.aarch64` | pigsty | 33.1 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-orphaned` | 1.0 | `u22.x86_64` | pigsty | 34.4 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-orphaned` | 1.0 | `u22.aarch64` | pigsty | 34.5 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-orphaned` | 1.0 | `u24.x86_64` | pigsty | 30.5 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-orphaned` | 1.0 | `u24.aarch64` | pigsty | 30.6 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_orphaned_15` | 1.0 | `el8.x86_64` | pigsty | 21.3 KiB | [pg_orphaned_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_15` | 1.0 | `el8.aarch64` | pigsty | 21.2 KiB | [pg_orphaned_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_15` | 1.0 | `el9.x86_64` | pigsty | 21.4 KiB | [pg_orphaned_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_15` | 1.0 | `el9.aarch64` | pigsty | 21.5 KiB | [pg_orphaned_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_15` | 1.0 | `el10.x86_64` | pigsty | 21.5 KiB | [pg_orphaned_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_orphaned_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_orphaned_15` | 1.0 | `el10.aarch64` | pigsty | 21.6 KiB | [pg_orphaned_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_orphaned_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-orphaned` | 1.0 | `d12.x86_64` | pigsty | 33.1 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-orphaned` | 1.0 | `d12.aarch64` | pigsty | 33.1 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-orphaned` | 1.0 | `u22.x86_64` | pigsty | 34.5 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-orphaned` | 1.0 | `u22.aarch64` | pigsty | 34.6 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-orphaned` | 1.0 | `u24.x86_64` | pigsty | 30.6 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-orphaned` | 1.0 | `u24.aarch64` | pigsty | 30.6 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_orphaned_14` | 1.0 | `el8.x86_64` | pigsty | 21.3 KiB | [pg_orphaned_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_14` | 1.0 | `el8.aarch64` | pigsty | 21.2 KiB | [pg_orphaned_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_14` | 1.0 | `el9.x86_64` | pigsty | 21.4 KiB | [pg_orphaned_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_14` | 1.0 | `el9.aarch64` | pigsty | 21.5 KiB | [pg_orphaned_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_14` | 1.0 | `el10.x86_64` | pigsty | 21.5 KiB | [pg_orphaned_14-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_orphaned_14-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_orphaned_14` | 1.0 | `el10.aarch64` | pigsty | 21.6 KiB | [pg_orphaned_14-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_orphaned_14-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-orphaned` | 1.0 | `d12.x86_64` | pigsty | 33.0 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-orphaned` | 1.0 | `d12.aarch64` | pigsty | 32.9 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-orphaned` | 1.0 | `u22.x86_64` | pigsty | 34.4 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-orphaned` | 1.0 | `u22.aarch64` | pigsty | 34.5 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-orphaned` | 1.0 | `u24.x86_64` | pigsty | 30.6 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-orphaned` | 1.0 | `u24.aarch64` | pigsty | 30.6 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_orphaned_13` | 1.0 | `el8.x86_64` | pigsty | 21.0 KiB | [pg_orphaned_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_13` | 1.0 | `el8.aarch64` | pigsty | 21.2 KiB | [pg_orphaned_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_13` | 1.0 | `el9.x86_64` | pigsty | 21.3 KiB | [pg_orphaned_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_13` | 1.0 | `el9.aarch64` | pigsty | 21.5 KiB | [pg_orphaned_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_13` | 1.0 | `el10.x86_64` | pigsty | 21.4 KiB | [pg_orphaned_13-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_orphaned_13-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_orphaned_13` | 1.0 | `el10.aarch64` | pigsty | 21.5 KiB | [pg_orphaned_13-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_orphaned_13-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-orphaned` | 1.0 | `d12.x86_64` | pigsty | 32.9 KiB | [postgresql-13-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-13-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-orphaned` | 1.0 | `d12.aarch64` | pigsty | 32.8 KiB | [postgresql-13-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-13-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-orphaned` | 1.0 | `u22.x86_64` | pigsty | 34.4 KiB | [postgresql-13-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-13-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-orphaned` | 1.0 | `u22.aarch64` | pigsty | 34.2 KiB | [postgresql-13-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-13-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-orphaned` | 1.0 | `u24.x86_64` | pigsty | 30.3 KiB | [postgresql-13-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-13-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-orphaned` | 1.0 | `u24.aarch64` | pigsty | 30.5 KiB | [postgresql-13-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-13-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bdrouvot/pg_orphaned" title="Repository" icon="github" subtitle="github.com/bdrouvot/pg_orphaned" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_orphaned-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_orphaned; # get pg_orphaned source code
pig build dep pg_orphaned; # install build dependencies
pig build pkg pg_orphaned; # build extension rpm or deb
pig build ext pg_orphaned; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_orphaned; # install by extension name, for the current active PG version
pig ext install pg_orphaned; # install via package alias, for the active PG version
pig ext install pg_orphaned -v 18;   # install for PG 18
pig ext install pg_orphaned -v 17;   # install for PG 17
pig ext install pg_orphaned -v 16;   # install for PG 16
pig ext install pg_orphaned -v 15;   # install for PG 15
pig ext install pg_orphaned -v 14;   # install for PG 14
pig ext install pg_orphaned -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_orphaned;
```

