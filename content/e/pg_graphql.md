---
title: "pg_graphql"
linkTitle: "pg_graphql"
description: "Add in-database GraphQL support"
weight: 2790
categories: ["FEAT"]
width: full
---

[**pg_graphql**](https://github.com/supabase/pg_graphql) : Add in-database GraphQL support


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2790** | {{< badge content="pg_graphql" link="https://github.com/supabase/pg_graphql" >}} | {{< ext "pg_graphql" >}} | `1.5.12` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `graphql` |
|   **See Also**    | {{< ext "age" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_net" >}} {{< ext "http" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "wrappers" >}} |

> [!Note] not an official release by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.12` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pg_graphql` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.12` | {{< bg "18" "pg_graphql_18" "green" >}} {{< bg "17" "pg_graphql_17" "green" >}} {{< bg "16" "pg_graphql_16" "green" >}} {{< bg "15" "pg_graphql_15" "green" >}} {{< bg "14" "pg_graphql_14" "green" >}} {{< bg "13" "pg_graphql_13" "red" >}} | `pg_graphql_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.12` | {{< bg "18" "postgresql-18-pg-graphql" "green" >}} {{< bg "17" "postgresql-17-pg-graphql" "green" >}} {{< bg "16" "postgresql-16-pg-graphql" "green" >}} {{< bg "15" "postgresql-15-pg-graphql" "green" >}} {{< bg "14" "postgresql-14-pg-graphql" "green" >}} {{< bg "13" "postgresql-13-pg-graphql" "red" >}} | `postgresql-$v-pg-graphql` | - |


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
| `pg_graphql_18` | `1.5.12` | [el8.x86_64](/os/el8.x86_64) | pigsty | 871.9 KiB | [pg_graphql_18-1.5.12-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_18-1.5.12-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_18` | `1.5.12` | [el8.aarch64](/os/el8.aarch64) | pigsty | 693.1 KiB | [pg_graphql_18-1.5.12-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_18-1.5.12-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_18` | `1.5.12` | [el9.x86_64](/os/el9.x86_64) | pigsty | 880.2 KiB | [pg_graphql_18-1.5.12-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_18-1.5.12-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_18` | `1.5.12` | [el9.aarch64](/os/el9.aarch64) | pigsty | 739.4 KiB | [pg_graphql_18-1.5.12-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_18-1.5.12-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_18` | `1.5.12` | [el10.x86_64](/os/el10.x86_64) | pigsty | 880.3 KiB | [pg_graphql_18-1.5.12-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_18-1.5.12-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_18` | `1.5.12` | [el10.aarch64](/os/el10.aarch64) | pigsty | 739.2 KiB | [pg_graphql_18-1.5.12-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_18-1.5.12-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-graphql` | `1.5.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 727.9 KiB | [postgresql-18-pg-graphql_1.5.12-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-graphql` | `1.5.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 566.2 KiB | [postgresql-18-pg-graphql_1.5.12-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-graphql` | `1.5.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 728.1 KiB | [postgresql-18-pg-graphql_1.5.12-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-graphql` | `1.5.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 564.2 KiB | [postgresql-18-pg-graphql_1.5.12-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-graphql` | `1.5.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 803.6 KiB | [postgresql-18-pg-graphql_1.5.12-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-graphql` | `1.5.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 661.7 KiB | [postgresql-18-pg-graphql_1.5.12-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-graphql` | `1.5.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 795.8 KiB | [postgresql-18-pg-graphql_1.5.12-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-graphql` | `1.5.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 654.6 KiB | [postgresql-18-pg-graphql_1.5.12-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.5.12-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_17` | `1.5.12` | [el8.x86_64](/os/el8.x86_64) | pigsty | 872.4 KiB | [pg_graphql_17-1.5.12-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_17-1.5.12-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_17` | `1.5.12` | [el8.aarch64](/os/el8.aarch64) | pigsty | 693.2 KiB | [pg_graphql_17-1.5.12-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_17-1.5.12-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_17` | `1.5.12` | [el9.x86_64](/os/el9.x86_64) | pigsty | 881.1 KiB | [pg_graphql_17-1.5.12-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_17-1.5.12-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_17` | `1.5.12` | [el9.aarch64](/os/el9.aarch64) | pigsty | 739.5 KiB | [pg_graphql_17-1.5.12-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_17-1.5.12-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_17` | `1.5.12` | [el10.x86_64](/os/el10.x86_64) | pigsty | 879.8 KiB | [pg_graphql_17-1.5.12-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_17-1.5.12-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_17` | `1.5.12` | [el10.aarch64](/os/el10.aarch64) | pigsty | 739.1 KiB | [pg_graphql_17-1.5.12-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_17-1.5.12-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-graphql` | `1.5.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 728.7 KiB | [postgresql-17-pg-graphql_1.5.12-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-graphql` | `1.5.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 566.3 KiB | [postgresql-17-pg-graphql_1.5.12-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-graphql` | `1.5.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 728.5 KiB | [postgresql-17-pg-graphql_1.5.12-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-graphql` | `1.5.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 564.3 KiB | [postgresql-17-pg-graphql_1.5.12-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-graphql` | `1.5.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 803.2 KiB | [postgresql-17-pg-graphql_1.5.12-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-graphql` | `1.5.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 661.7 KiB | [postgresql-17-pg-graphql_1.5.12-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-graphql` | `1.5.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 795.9 KiB | [postgresql-17-pg-graphql_1.5.12-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-graphql` | `1.5.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 654.7 KiB | [postgresql-17-pg-graphql_1.5.12-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.12-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_16` | `1.5.12` | [el8.x86_64](/os/el8.x86_64) | pigsty | 871.8 KiB | [pg_graphql_16-1.5.12-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_16-1.5.12-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_16` | `1.5.12` | [el8.aarch64](/os/el8.aarch64) | pigsty | 692.7 KiB | [pg_graphql_16-1.5.12-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_16-1.5.12-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_16` | `1.5.12` | [el9.x86_64](/os/el9.x86_64) | pigsty | 880.7 KiB | [pg_graphql_16-1.5.12-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_16-1.5.12-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_16` | `1.5.12` | [el9.aarch64](/os/el9.aarch64) | pigsty | 739.6 KiB | [pg_graphql_16-1.5.12-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_16-1.5.12-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_16` | `1.5.12` | [el10.x86_64](/os/el10.x86_64) | pigsty | 880.8 KiB | [pg_graphql_16-1.5.12-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_16-1.5.12-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_16` | `1.5.12` | [el10.aarch64](/os/el10.aarch64) | pigsty | 739.4 KiB | [pg_graphql_16-1.5.12-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_16-1.5.12-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-graphql` | `1.5.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 728.3 KiB | [postgresql-16-pg-graphql_1.5.12-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-graphql` | `1.5.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 563.9 KiB | [postgresql-16-pg-graphql_1.5.12-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-graphql` | `1.5.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 727.9 KiB | [postgresql-16-pg-graphql_1.5.12-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-graphql` | `1.5.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 564.0 KiB | [postgresql-16-pg-graphql_1.5.12-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-graphql` | `1.5.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 803.0 KiB | [postgresql-16-pg-graphql_1.5.12-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-graphql` | `1.5.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 661.5 KiB | [postgresql-16-pg-graphql_1.5.12-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-graphql` | `1.5.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 795.6 KiB | [postgresql-16-pg-graphql_1.5.12-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-graphql` | `1.5.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 654.5 KiB | [postgresql-16-pg-graphql_1.5.12-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.12-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_15` | `1.5.12` | [el8.x86_64](/os/el8.x86_64) | pigsty | 871.8 KiB | [pg_graphql_15-1.5.12-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_15-1.5.12-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_15` | `1.5.12` | [el8.aarch64](/os/el8.aarch64) | pigsty | 692.9 KiB | [pg_graphql_15-1.5.12-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_15-1.5.12-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_15` | `1.5.12` | [el9.x86_64](/os/el9.x86_64) | pigsty | 879.6 KiB | [pg_graphql_15-1.5.12-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_15-1.5.12-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_15` | `1.5.12` | [el9.aarch64](/os/el9.aarch64) | pigsty | 739.5 KiB | [pg_graphql_15-1.5.12-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_15-1.5.12-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_15` | `1.5.12` | [el10.x86_64](/os/el10.x86_64) | pigsty | 879.6 KiB | [pg_graphql_15-1.5.12-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_15-1.5.12-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_15` | `1.5.12` | [el10.aarch64](/os/el10.aarch64) | pigsty | 739.3 KiB | [pg_graphql_15-1.5.12-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_15-1.5.12-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-graphql` | `1.5.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 728.4 KiB | [postgresql-15-pg-graphql_1.5.12-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-graphql` | `1.5.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 564.0 KiB | [postgresql-15-pg-graphql_1.5.12-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-graphql` | `1.5.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 728.4 KiB | [postgresql-15-pg-graphql_1.5.12-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-graphql` | `1.5.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 564.0 KiB | [postgresql-15-pg-graphql_1.5.12-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-graphql` | `1.5.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 803.8 KiB | [postgresql-15-pg-graphql_1.5.12-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-graphql` | `1.5.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 661.4 KiB | [postgresql-15-pg-graphql_1.5.12-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-graphql` | `1.5.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 798.5 KiB | [postgresql-15-pg-graphql_1.5.12-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-graphql` | `1.5.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 654.3 KiB | [postgresql-15-pg-graphql_1.5.12-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.12-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_14` | `1.5.12` | [el8.x86_64](/os/el8.x86_64) | pigsty | 871.7 KiB | [pg_graphql_14-1.5.12-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_14-1.5.12-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_14` | `1.5.12` | [el8.aarch64](/os/el8.aarch64) | pigsty | 692.5 KiB | [pg_graphql_14-1.5.12-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_14-1.5.12-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_14` | `1.5.12` | [el9.x86_64](/os/el9.x86_64) | pigsty | 879.3 KiB | [pg_graphql_14-1.5.12-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_14-1.5.12-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_14` | `1.5.12` | [el9.aarch64](/os/el9.aarch64) | pigsty | 739.4 KiB | [pg_graphql_14-1.5.12-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_14-1.5.12-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_14` | `1.5.12` | [el10.x86_64](/os/el10.x86_64) | pigsty | 880.8 KiB | [pg_graphql_14-1.5.12-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_14-1.5.12-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_14` | `1.5.12` | [el10.aarch64](/os/el10.aarch64) | pigsty | 739.0 KiB | [pg_graphql_14-1.5.12-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_14-1.5.12-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-graphql` | `1.5.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 728.3 KiB | [postgresql-14-pg-graphql_1.5.12-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-graphql` | `1.5.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 563.6 KiB | [postgresql-14-pg-graphql_1.5.12-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-graphql` | `1.5.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 728.0 KiB | [postgresql-14-pg-graphql_1.5.12-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-graphql` | `1.5.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 566.1 KiB | [postgresql-14-pg-graphql_1.5.12-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-graphql` | `1.5.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 802.9 KiB | [postgresql-14-pg-graphql_1.5.12-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-graphql` | `1.5.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 661.5 KiB | [postgresql-14-pg-graphql_1.5.12-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-graphql` | `1.5.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 796.4 KiB | [postgresql-14-pg-graphql_1.5.12-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-graphql` | `1.5.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 654.5 KiB | [postgresql-14-pg-graphql_1.5.12-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.12-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/pg_graphql" title="Repository" icon="github" subtitle="github.com/supabase/pg_graphql" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_graphql-1.5.12.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_graphql;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_graphql;		# install via package name, for the active PG version

pig install pg_graphql -v 18;   # install for PG 18
pig install pg_graphql -v 17;   # install for PG 17
pig install pg_graphql -v 16;   # install for PG 16
pig install pg_graphql -v 15;   # install for PG 15
pig install pg_graphql -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_graphql;
```
