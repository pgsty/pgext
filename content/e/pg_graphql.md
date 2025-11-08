---
title: "pg_graphql"
linkTitle: "pg_graphql"
description: "Add in-database GraphQL support"
weight: 2790
categories: ["FEAT"]
width: full
---

Add in-database GraphQL support


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2790** | {{< badge content="pg_graphql" link="https://github.com/supabase/pg_graphql" >}} | {{< ext "pg_graphql" >}} | `1.5.12` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "age" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_net" >}} {{< ext "http" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "wrappers" >}} |

> [!Note] not an official release by Vonng


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_graphql" >}} | `1.5.12` | {{< bg "18" "pg_graphql_18" "green" >}} {{< bg "17" "pg_graphql_17" "green" >}} {{< bg "16" "pg_graphql_16" "green" >}} {{< bg "15" "pg_graphql_15" "green" >}} {{< bg "14" "pg_graphql_14" "green" >}} {{< bg "13" "pg_graphql_13" "red" >}} | `pg_graphql_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_graphql" >}} | `1.5.12` | {{< bg "18" "postgresql-18-pg-graphql" "green" >}} {{< bg "17" "postgresql-17-pg-graphql" "green" >}} {{< bg "16" "postgresql-16-pg-graphql" "green" >}} {{< bg "15" "postgresql-15-pg-graphql" "green" >}} {{< bg "14" "postgresql-14-pg-graphql" "green" >}} {{< bg "13" "postgresql-13-pg-graphql" "red" >}} | `postgresql-$v-pg-graphql` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_graphql_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_graphql_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_graphql_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_graphql_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_graphql_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "pg_graphql_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_graphql_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-graphql : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-graphql : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-graphql : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-graphql : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-graphql : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-graphql : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-graphql : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.12" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-graphql : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_18` | 1.5.12 | `el8.x86_64` | pigsty | 878.2 KiB | [pg_graphql_18-1.5.12-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_18-1.5.12-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_18` | 1.5.12 | `el8.aarch64` | pigsty | 704.9 KiB | [pg_graphql_18-1.5.12-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_18-1.5.12-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_18` | 1.5.12 | `el9.x86_64` | pigsty | 885.1 KiB | [pg_graphql_18-1.5.12-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_18-1.5.12-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_18` | 1.5.12 | `el9.aarch64` | pigsty | 752.4 KiB | [pg_graphql_18-1.5.12-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_18-1.5.12-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_18` | 1.5.12 | `el10.x86_64` | pigsty | 882.0 KiB | [pg_graphql_18-1.5.12-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_18-1.5.12-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_18` | 1.5.12 | `el10.aarch64` | pigsty | 753.5 KiB | [pg_graphql_18-1.5.12-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_18-1.5.12-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-graphql` | 1.5.12 | `d12.x86_64` | pigsty | 743.5 KiB | [postgresql-18-pg-graphql_1.5.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-graphql` | 1.5.12 | `d12.aarch64` | pigsty | 582.4 KiB | [postgresql-18-pg-graphql_1.5.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-graphql` | 1.5.12 | `d13.x86_64` | pigsty | 743.6 KiB | [postgresql-18-pg-graphql_1.5.12-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-graphql` | 1.5.12 | `d13.aarch64` | pigsty | 581.7 KiB | [postgresql-18-pg-graphql_1.5.12-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-graphql` | 1.5.12 | `u22.x86_64` | pigsty | 819.5 KiB | [postgresql-18-pg-graphql_1.5.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-graphql` | 1.5.12 | `u22.aarch64` | pigsty | 678.9 KiB | [postgresql-18-pg-graphql_1.5.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-graphql` | 1.5.12 | `u24.x86_64` | pigsty | 811.8 KiB | [postgresql-18-pg-graphql_1.5.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-graphql` | 1.5.12 | `u24.aarch64` | pigsty | 672.2 KiB | [postgresql-18-pg-graphql_1.5.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_17` | 1.5.12 | `el8.x86_64` | pigsty | 878.0 KiB | [pg_graphql_17-1.5.12-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_17-1.5.12-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_17` | 1.5.12 | `el8.aarch64` | pigsty | 704.0 KiB | [pg_graphql_17-1.5.12-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_17-1.5.12-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_17` | 1.5.12 | `el9.x86_64` | pigsty | 885.1 KiB | [pg_graphql_17-1.5.12-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_17-1.5.12-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_17` | 1.5.12 | `el9.aarch64` | pigsty | 751.7 KiB | [pg_graphql_17-1.5.12-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_17-1.5.12-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_17` | 1.5.12 | `el10.x86_64` | pigsty | 881.8 KiB | [pg_graphql_17-1.5.12-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_17-1.5.12-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_17` | 1.5.12 | `el10.aarch64` | pigsty | 753.1 KiB | [pg_graphql_17-1.5.12-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_17-1.5.12-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-graphql` | 1.5.12 | `d12.x86_64` | pigsty | 743.7 KiB | [postgresql-17-pg-graphql_1.5.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-graphql` | 1.5.12 | `d12.aarch64` | pigsty | 580.3 KiB | [postgresql-17-pg-graphql_1.5.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-graphql` | 1.5.12 | `d13.x86_64` | pigsty | 743.5 KiB | [postgresql-17-pg-graphql_1.5.12-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-graphql` | 1.5.12 | `d13.aarch64` | pigsty | 580.7 KiB | [postgresql-17-pg-graphql_1.5.12-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-graphql` | 1.5.12 | `u22.x86_64` | pigsty | 819.3 KiB | [postgresql-17-pg-graphql_1.5.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-graphql` | 1.5.12 | `u22.aarch64` | pigsty | 678.8 KiB | [postgresql-17-pg-graphql_1.5.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-graphql` | 1.5.12 | `u24.x86_64` | pigsty | 811.2 KiB | [postgresql-17-pg-graphql_1.5.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-graphql` | 1.5.12 | `u24.aarch64` | pigsty | 672.2 KiB | [postgresql-17-pg-graphql_1.5.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_16` | 1.5.12 | `el8.x86_64` | pigsty | 877.9 KiB | [pg_graphql_16-1.5.12-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_16-1.5.12-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_16` | 1.5.12 | `el8.aarch64` | pigsty | 704.2 KiB | [pg_graphql_16-1.5.12-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_16-1.5.12-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_16` | 1.5.12 | `el9.x86_64` | pigsty | 885.5 KiB | [pg_graphql_16-1.5.12-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_16-1.5.12-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_16` | 1.5.12 | `el9.aarch64` | pigsty | 751.9 KiB | [pg_graphql_16-1.5.12-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_16-1.5.12-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_16` | 1.5.12 | `el10.x86_64` | pigsty | 881.6 KiB | [pg_graphql_16-1.5.12-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_16-1.5.12-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_16` | 1.5.12 | `el10.aarch64` | pigsty | 753.7 KiB | [pg_graphql_16-1.5.12-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_16-1.5.12-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-graphql` | 1.5.12 | `d12.x86_64` | pigsty | 743.3 KiB | [postgresql-16-pg-graphql_1.5.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-graphql` | 1.5.12 | `d12.aarch64` | pigsty | 580.5 KiB | [postgresql-16-pg-graphql_1.5.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-graphql` | 1.5.12 | `d13.x86_64` | pigsty | 743.3 KiB | [postgresql-16-pg-graphql_1.5.12-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-graphql` | 1.5.12 | `d13.aarch64` | pigsty | 580.8 KiB | [postgresql-16-pg-graphql_1.5.12-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-graphql` | 1.5.12 | `u22.x86_64` | pigsty | 820.2 KiB | [postgresql-16-pg-graphql_1.5.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-graphql` | 1.5.12 | `u22.aarch64` | pigsty | 678.9 KiB | [postgresql-16-pg-graphql_1.5.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-graphql` | 1.5.12 | `u24.x86_64` | pigsty | 812.3 KiB | [postgresql-16-pg-graphql_1.5.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-graphql` | 1.5.12 | `u24.aarch64` | pigsty | 672.4 KiB | [postgresql-16-pg-graphql_1.5.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_15` | 1.5.12 | `el8.x86_64` | pigsty | 878.0 KiB | [pg_graphql_15-1.5.12-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_15-1.5.12-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_15` | 1.5.12 | `el8.aarch64` | pigsty | 704.0 KiB | [pg_graphql_15-1.5.12-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_15-1.5.12-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_15` | 1.5.12 | `el9.x86_64` | pigsty | 885.4 KiB | [pg_graphql_15-1.5.12-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_15-1.5.12-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_15` | 1.5.12 | `el9.aarch64` | pigsty | 751.8 KiB | [pg_graphql_15-1.5.12-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_15-1.5.12-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_15` | 1.5.12 | `el10.x86_64` | pigsty | 880.9 KiB | [pg_graphql_15-1.5.12-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_15-1.5.12-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_15` | 1.5.12 | `el10.aarch64` | pigsty | 753.5 KiB | [pg_graphql_15-1.5.12-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_15-1.5.12-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-graphql` | 1.5.12 | `d12.x86_64` | pigsty | 744.2 KiB | [postgresql-15-pg-graphql_1.5.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-graphql` | 1.5.12 | `d12.aarch64` | pigsty | 582.3 KiB | [postgresql-15-pg-graphql_1.5.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-graphql` | 1.5.12 | `d13.x86_64` | pigsty | 743.4 KiB | [postgresql-15-pg-graphql_1.5.12-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-graphql` | 1.5.12 | `d13.aarch64` | pigsty | 580.7 KiB | [postgresql-15-pg-graphql_1.5.12-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-graphql` | 1.5.12 | `u22.x86_64` | pigsty | 820.3 KiB | [postgresql-15-pg-graphql_1.5.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-graphql` | 1.5.12 | `u22.aarch64` | pigsty | 678.7 KiB | [postgresql-15-pg-graphql_1.5.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-graphql` | 1.5.12 | `u24.x86_64` | pigsty | 812.6 KiB | [postgresql-15-pg-graphql_1.5.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-graphql` | 1.5.12 | `u24.aarch64` | pigsty | 672.3 KiB | [postgresql-15-pg-graphql_1.5.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_14` | 1.5.12 | `el8.x86_64` | pigsty | 877.8 KiB | [pg_graphql_14-1.5.12-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_14-1.5.12-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_14` | 1.5.12 | `el8.aarch64` | pigsty | 705.0 KiB | [pg_graphql_14-1.5.12-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_14-1.5.12-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_14` | 1.5.12 | `el9.x86_64` | pigsty | 884.4 KiB | [pg_graphql_14-1.5.12-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_14-1.5.12-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_14` | 1.5.12 | `el9.aarch64` | pigsty | 751.5 KiB | [pg_graphql_14-1.5.12-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_14-1.5.12-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_14` | 1.5.12 | `el10.x86_64` | pigsty | 881.1 KiB | [pg_graphql_14-1.5.12-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_14-1.5.12-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_14` | 1.5.12 | `el10.aarch64` | pigsty | 753.3 KiB | [pg_graphql_14-1.5.12-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_14-1.5.12-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-graphql` | 1.5.12 | `d12.x86_64` | pigsty | 744.0 KiB | [postgresql-14-pg-graphql_1.5.12-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-graphql` | 1.5.12 | `d12.aarch64` | pigsty | 581.6 KiB | [postgresql-14-pg-graphql_1.5.12-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-graphql` | 1.5.12 | `d13.x86_64` | pigsty | 742.8 KiB | [postgresql-14-pg-graphql_1.5.12-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-graphql` | 1.5.12 | `d13.aarch64` | pigsty | 581.7 KiB | [postgresql-14-pg-graphql_1.5.12-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-graphql` | 1.5.12 | `u22.x86_64` | pigsty | 819.2 KiB | [postgresql-14-pg-graphql_1.5.12-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-graphql` | 1.5.12 | `u22.aarch64` | pigsty | 678.4 KiB | [postgresql-14-pg-graphql_1.5.12-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-graphql` | 1.5.12 | `u24.x86_64` | pigsty | 811.3 KiB | [postgresql-14-pg-graphql_1.5.12-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-graphql` | 1.5.12 | `u24.aarch64` | pigsty | 672.3 KiB | [postgresql-14-pg-graphql_1.5.12-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/pg_graphql" title="Repository" icon="github" subtitle="github.com/supabase/pg_graphql" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_graphql-1.5.12.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_graphql; # get pg_graphql source code
pig build dep pg_graphql; # install build dependencies
pig build pkg pg_graphql; # build extension rpm or deb
pig build ext pg_graphql; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_graphql; # install by extension name, for the current active PG version
pig ext install pg_graphql; # install via package alias, for the active PG version
pig ext install pg_graphql -v 18;   # install for PG 18
pig ext install pg_graphql -v 17;   # install for PG 17
pig ext install pg_graphql -v 16;   # install for PG 16
pig ext install pg_graphql -v 15;   # install for PG 15
pig ext install pg_graphql -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_graphql CASCADE SCHEMA graphql;
```

