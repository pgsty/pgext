---
title: "plv8"
linkTitle: "plv8"
description: "PL/JavaScript (v8) trusted procedural language"
weight: 3010
categories: ["LANG"]
width: full
---

PL/JavaScript (v8) trusted procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3010** | {{< badge content="plv8" link="https://github.com/plv8/plv8" >}} | {{< ext "plv8" >}} | `3.2.4` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "pg_tle" >}} {{< ext "pllua" >}} {{< ext "plluau" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/plv8" >}} | `3.2.4` | {{< bg "18" "plv8_18*" "red" >}} {{< bg "17" "plv8_17*" "green" >}} {{< bg "16" "plv8_16*" "green" >}} {{< bg "15" "plv8_15*" "green" >}} {{< bg "14" "plv8_14*" "green" >}} | `plv8_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/plv8" >}} | `3.2.4` | {{< bg "18" "postgresql-18-plv8" "red" >}} {{< bg "17" "postgresql-17-plv8" "green" >}} {{< bg "16" "postgresql-16-plv8" "green" >}} {{< bg "15" "postgresql-15-plv8" "green" >}} {{< bg "14" "postgresql-14-plv8" "green" >}} | `postgresql-$v-plv8` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 3.2.4" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 3.2.4" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 3.2.4" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 3.2.4" "plv8_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "plv8_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-plv8 : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-plv8 : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-plv8 : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-plv8 : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-plv8 : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-plv8 : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.2.4" "postgresql-17-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-16-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-15-plv8 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.4" "postgresql-14-plv8 : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_18` | 3.2.4 | `el8.x86_64` | pigsty | 7.5 MiB | [plv8_18-3.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_18-3.2.4-1PIGSTY.el8.x86_64.rpm) |
| `plv8_18` | 3.2.4 | `el8.aarch64` | pigsty | 7.0 MiB | [plv8_18-3.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_18-3.2.4-1PIGSTY.el8.aarch64.rpm) |
| `plv8_18` | 3.2.4 | `el9.x86_64` | pigsty | 7.6 MiB | [plv8_18-3.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_18-3.2.4-1PIGSTY.el9.x86_64.rpm) |
| `plv8_18` | 3.2.4 | `el9.aarch64` | pigsty | 7.3 MiB | [plv8_18-3.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_18-3.2.4-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_17` | 3.2.4 | `el8.x86_64` | pigsty | 7.5 MiB | [plv8_17-3.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_17-3.2.4-1PIGSTY.el8.x86_64.rpm) |
| `plv8_17` | 3.2.4 | `el8.aarch64` | pigsty | 7.0 MiB | [plv8_17-3.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_17-3.2.4-1PIGSTY.el8.aarch64.rpm) |
| `plv8_17` | 3.2.4 | `el9.x86_64` | pigsty | 7.6 MiB | [plv8_17-3.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_17-3.2.4-1PIGSTY.el9.x86_64.rpm) |
| `plv8_17` | 3.2.4 | `el9.aarch64` | pigsty | 7.3 MiB | [plv8_17-3.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_17-3.2.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-plv8` | 3.2.4 | `d12.x86_64` | pigsty | 6.2 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-plv8` | 3.2.4 | `d12.aarch64` | pigsty | 5.6 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-plv8` | 3.2.4 | `u22.x86_64` | pigsty | 6.8 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-plv8` | 3.2.4 | `u22.aarch64` | pigsty | 6.6 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-plv8` | 3.2.4 | `u24.x86_64` | pigsty | 6.7 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-plv8` | 3.2.4 | `u24.aarch64` | pigsty | 6.6 MiB | [postgresql-17-plv8_3.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-17-plv8_3.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_16` | 3.2.4 | `el8.x86_64` | pigsty | 7.5 MiB | [plv8_16-3.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_16-3.2.4-1PIGSTY.el8.x86_64.rpm) |
| `plv8_16` | 3.2.4 | `el8.aarch64` | pigsty | 7.0 MiB | [plv8_16-3.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_16-3.2.4-1PIGSTY.el8.aarch64.rpm) |
| `plv8_16` | 3.2.4 | `el9.x86_64` | pigsty | 7.6 MiB | [plv8_16-3.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_16-3.2.4-1PIGSTY.el9.x86_64.rpm) |
| `plv8_16` | 3.2.4 | `el9.aarch64` | pigsty | 7.3 MiB | [plv8_16-3.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_16-3.2.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-plv8` | 3.2.4 | `d12.x86_64` | pigsty | 6.2 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-plv8` | 3.2.4 | `d12.aarch64` | pigsty | 5.6 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-plv8` | 3.2.4 | `u22.x86_64` | pigsty | 6.8 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-plv8` | 3.2.4 | `u22.aarch64` | pigsty | 6.7 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-plv8` | 3.2.4 | `u24.x86_64` | pigsty | 6.7 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-plv8` | 3.2.4 | `u24.aarch64` | pigsty | 6.6 MiB | [postgresql-16-plv8_3.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-16-plv8_3.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_15` | 3.2.4 | `el8.x86_64` | pigsty | 7.5 MiB | [plv8_15-3.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_15-3.2.4-1PIGSTY.el8.x86_64.rpm) |
| `plv8_15` | 3.2.4 | `el8.aarch64` | pigsty | 7.0 MiB | [plv8_15-3.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_15-3.2.4-1PIGSTY.el8.aarch64.rpm) |
| `plv8_15` | 3.2.4 | `el9.x86_64` | pigsty | 7.6 MiB | [plv8_15-3.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_15-3.2.4-1PIGSTY.el9.x86_64.rpm) |
| `plv8_15` | 3.2.4 | `el9.aarch64` | pigsty | 7.3 MiB | [plv8_15-3.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_15-3.2.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-plv8` | 3.2.4 | `d12.x86_64` | pigsty | 6.2 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-plv8` | 3.2.4 | `d12.aarch64` | pigsty | 5.6 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-plv8` | 3.2.4 | `u22.x86_64` | pigsty | 6.8 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-plv8` | 3.2.4 | `u22.aarch64` | pigsty | 6.7 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-plv8` | 3.2.4 | `u24.x86_64` | pigsty | 6.7 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-plv8` | 3.2.4 | `u24.aarch64` | pigsty | 6.6 MiB | [postgresql-15-plv8_3.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-15-plv8_3.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plv8_14` | 3.2.4 | `el8.x86_64` | pigsty | 7.5 MiB | [plv8_14-3.2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plv8_14-3.2.4-1PIGSTY.el8.x86_64.rpm) |
| `plv8_14` | 3.2.4 | `el8.aarch64` | pigsty | 7.0 MiB | [plv8_14-3.2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plv8_14-3.2.4-1PIGSTY.el8.aarch64.rpm) |
| `plv8_14` | 3.2.4 | `el9.x86_64` | pigsty | 7.6 MiB | [plv8_14-3.2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plv8_14-3.2.4-1PIGSTY.el9.x86_64.rpm) |
| `plv8_14` | 3.2.4 | `el9.aarch64` | pigsty | 7.3 MiB | [plv8_14-3.2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plv8_14-3.2.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-plv8` | 3.2.4 | `d12.x86_64` | pigsty | 6.2 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-plv8` | 3.2.4 | `d12.aarch64` | pigsty | 5.6 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-plv8` | 3.2.4 | `u22.x86_64` | pigsty | 6.8 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-plv8` | 3.2.4 | `u22.aarch64` | pigsty | 6.6 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-plv8` | 3.2.4 | `u24.x86_64` | pigsty | 6.7 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-plv8` | 3.2.4 | `u24.aarch64` | pigsty | 6.6 MiB | [postgresql-14-plv8_3.2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plv8/postgresql-14-plv8_3.2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/plv8/plv8" title="Repository" icon="github" subtitle="github.com/plv8/plv8" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plv8-3.2.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get plv8; # get plv8 source code
pig build dep plv8; # install build dependencies
pig build pkg plv8; # build extension rpm or deb
pig build ext plv8; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install plv8; # install by extension name, for the current active PG version
pig ext install plv8; # install via package alias, for the active PG version
pig ext install plv8 -v 18;   # install for PG 18
pig ext install plv8 -v 17;   # install for PG 17
pig ext install plv8 -v 16;   # install for PG 16
pig ext install plv8 -v 15;   # install for PG 15
pig ext install plv8 -v 14;   # install for PG 14
pig ext install plv8 -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION plv8 CASCADE SCHEMA pg_catalog;
```



--------

## Usage

```sql
CREATE EXTENSION plv8;

SELECT plv8_version();
SELECT plv8_info();

DO $$ plv8.elog(NOTICE, plv8.version); $$ LANGUAGE plv8;
```

Example:

```sql
CREATE FUNCTION plv8_test(keys TEXT[], vals TEXT[]) RETURNS JSON AS $$
    var o = {};
    for(var i=0; i<keys.length; i++){
        o[keys[i]] = vals[i];
    }
    return o;
$$ LANGUAGE plv8 IMMUTABLE STRICT;


SELECT plv8_test(ARRAY['name', 'age'], ARRAY['Tom', '29']);
```



