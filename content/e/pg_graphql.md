---
title: "pg_graphql"
linkTitle: "pg_graphql"
description: "Add in-database GraphQL support"
weight: 2740
categories: ["FEAT"]
width: full
---

[**pg_graphql**](https://github.com/supabase/pg_graphql) : Add in-database GraphQL support


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2740** | {{< badge content="pg_graphql" link="https://github.com/supabase/pg_graphql" >}} | {{< ext "pg_graphql" >}} | `1.6.1` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `graphql` |
|   **See Also**    | {{< ext "age" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_net" >}} {{< ext "http" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "wrappers" >}} |

> [!Note] not an official release by Vonng.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.6.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_graphql` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.6.1` | {{< bg "18" "pg_graphql_18" "green" >}} {{< bg "17" "pg_graphql_17" "green" >}} {{< bg "16" "pg_graphql_16" "green" >}} {{< bg "15" "pg_graphql_15" "green" >}} {{< bg "14" "pg_graphql_14" "green" >}} | `pg_graphql_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.6.1` | {{< bg "18" "postgresql-18-pg-graphql" "green" >}} {{< bg "17" "postgresql-17-pg-graphql" "green" >}} {{< bg "16" "postgresql-16-pg-graphql" "green" >}} {{< bg "15" "postgresql-15-pg-graphql" "green" >}} {{< bg "14" "postgresql-14-pg-graphql" "green" >}} | `postgresql-$v-pg-graphql` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "pg_graphql_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-15-pg-graphql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-14-pg-graphql : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_18` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.3 MiB | [pg_graphql_18-1.6.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_18-1.6.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_18` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.2 MiB | [pg_graphql_18-1.6.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_18-1.6.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_18` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.3 MiB | [pg_graphql_18-1.6.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_18-1.6.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_18` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.2 MiB | [pg_graphql_18-1.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_18-1.6.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_18` | `1.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.3 MiB | [pg_graphql_18-1.6.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_18-1.6.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_18` | `1.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.2 MiB | [pg_graphql_18-1.6.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_18-1.6.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-graphql` | `1.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-18-pg-graphql_1.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-graphql` | `1.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 937.3 KiB | [postgresql-18-pg-graphql_1.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-graphql` | `1.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.1 MiB | [postgresql-18-pg-graphql_1.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-graphql` | `1.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 937.7 KiB | [postgresql-18-pg-graphql_1.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-graphql` | `1.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.2 MiB | [postgresql-18-pg-graphql_1.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-graphql` | `1.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-18-pg-graphql_1.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-graphql` | `1.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.2 MiB | [postgresql-18-pg-graphql_1.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-graphql` | `1.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-18-pg-graphql_1.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-graphql` | `1.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.2 MiB | [postgresql-18-pg-graphql_1.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-graphql` | `1.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.1 MiB | [postgresql-18-pg-graphql_1.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-graphql/postgresql-18-pg-graphql_1.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_17` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.3 MiB | [pg_graphql_17-1.6.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_17-1.6.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_17` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.2 MiB | [pg_graphql_17-1.6.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_17-1.6.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_17` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.3 MiB | [pg_graphql_17-1.6.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_17-1.6.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_17` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.2 MiB | [pg_graphql_17-1.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_17-1.6.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_17` | `1.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.3 MiB | [pg_graphql_17-1.6.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_17-1.6.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_17` | `1.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.2 MiB | [pg_graphql_17-1.6.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_17-1.6.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-graphql` | `1.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-17-pg-graphql_1.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-graphql` | `1.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 935.5 KiB | [postgresql-17-pg-graphql_1.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-graphql` | `1.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.1 MiB | [postgresql-17-pg-graphql_1.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-graphql` | `1.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 936.0 KiB | [postgresql-17-pg-graphql_1.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-graphql` | `1.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.2 MiB | [postgresql-17-pg-graphql_1.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-graphql` | `1.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-17-pg-graphql_1.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-graphql` | `1.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.2 MiB | [postgresql-17-pg-graphql_1.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-graphql` | `1.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-17-pg-graphql_1.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-graphql` | `1.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.2 MiB | [postgresql-17-pg-graphql_1.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-graphql` | `1.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.1 MiB | [postgresql-17-pg-graphql_1.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_16` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.3 MiB | [pg_graphql_16-1.6.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_16-1.6.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_16` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.2 MiB | [pg_graphql_16-1.6.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_16-1.6.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_16` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.3 MiB | [pg_graphql_16-1.6.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_16-1.6.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_16` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.2 MiB | [pg_graphql_16-1.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_16-1.6.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_16` | `1.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.3 MiB | [pg_graphql_16-1.6.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_16-1.6.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_16` | `1.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.2 MiB | [pg_graphql_16-1.6.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_16-1.6.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-graphql` | `1.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-16-pg-graphql_1.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-graphql` | `1.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 934.0 KiB | [postgresql-16-pg-graphql_1.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-graphql` | `1.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.1 MiB | [postgresql-16-pg-graphql_1.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-graphql` | `1.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 936.1 KiB | [postgresql-16-pg-graphql_1.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-graphql` | `1.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.2 MiB | [postgresql-16-pg-graphql_1.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-graphql` | `1.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-16-pg-graphql_1.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-graphql` | `1.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.2 MiB | [postgresql-16-pg-graphql_1.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-graphql` | `1.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-16-pg-graphql_1.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-graphql` | `1.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.2 MiB | [postgresql-16-pg-graphql_1.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-graphql` | `1.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.1 MiB | [postgresql-16-pg-graphql_1.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_15` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.3 MiB | [pg_graphql_15-1.6.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_15-1.6.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_15` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.2 MiB | [pg_graphql_15-1.6.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_15-1.6.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_15` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.3 MiB | [pg_graphql_15-1.6.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_15-1.6.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_15` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.2 MiB | [pg_graphql_15-1.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_15-1.6.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_15` | `1.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.3 MiB | [pg_graphql_15-1.6.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_15-1.6.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_15` | `1.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.2 MiB | [pg_graphql_15-1.6.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_15-1.6.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-graphql` | `1.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-15-pg-graphql_1.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-graphql` | `1.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 930.9 KiB | [postgresql-15-pg-graphql_1.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-graphql` | `1.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.1 MiB | [postgresql-15-pg-graphql_1.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-graphql` | `1.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 932.1 KiB | [postgresql-15-pg-graphql_1.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-graphql` | `1.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.2 MiB | [postgresql-15-pg-graphql_1.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-graphql` | `1.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-15-pg-graphql_1.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-graphql` | `1.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.2 MiB | [postgresql-15-pg-graphql_1.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-graphql` | `1.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-15-pg-graphql_1.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-graphql` | `1.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.2 MiB | [postgresql-15-pg-graphql_1.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-graphql` | `1.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.1 MiB | [postgresql-15-pg-graphql_1.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_graphql_14` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.3 MiB | [pg_graphql_14-1.6.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_14-1.6.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_14` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.2 MiB | [pg_graphql_14-1.6.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_14-1.6.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_14` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.3 MiB | [pg_graphql_14-1.6.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_14-1.6.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_14` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.2 MiB | [pg_graphql_14-1.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_14-1.6.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_14` | `1.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.3 MiB | [pg_graphql_14-1.6.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_graphql_14-1.6.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_graphql_14` | `1.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.2 MiB | [pg_graphql_14-1.6.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_graphql_14-1.6.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-graphql` | `1.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-14-pg-graphql_1.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-graphql` | `1.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 928.3 KiB | [postgresql-14-pg-graphql_1.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-graphql` | `1.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.1 MiB | [postgresql-14-pg-graphql_1.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-graphql` | `1.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 929.1 KiB | [postgresql-14-pg-graphql_1.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-graphql` | `1.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.2 MiB | [postgresql-14-pg-graphql_1.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-graphql` | `1.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-14-pg-graphql_1.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-graphql` | `1.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.2 MiB | [postgresql-14-pg-graphql_1.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-graphql` | `1.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-14-pg-graphql_1.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-graphql` | `1.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.2 MiB | [postgresql-14-pg-graphql_1.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-graphql` | `1.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.1 MiB | [postgresql-14-pg-graphql_1.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/pg_graphql" title="Repository" icon="github" subtitle="github.com/supabase/pg_graphql" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_graphql-1.6.1.tar.gz" >}}
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




## Usage

Sources: [README](https://github.com/supabase/pg_graphql/blob/v1.6.1/README.md), [documentation](https://supabase.github.io/pg_graphql), [v1.6.1 release](https://github.com/supabase/pg_graphql/releases/tag/v1.6.1)

`pg_graphql` reflects a GraphQL schema from your existing SQL schema, enabling GraphQL queries directly inside PostgreSQL without additional servers or middleware.

### Schema Reflection

Tables, foreign keys, and enums are automatically mapped to GraphQL types:

```sql
CREATE TABLE account (
    id serial PRIMARY KEY,
    email varchar(255) NOT NULL,
    created_at timestamp NOT NULL
);

CREATE TABLE blog (
    id serial PRIMARY KEY,
    owner_id integer NOT NULL REFERENCES account(id),
    name varchar(255) NOT NULL,
    description varchar(255)
);

CREATE TYPE blog_post_status AS ENUM ('PENDING', 'RELEASED');

CREATE TABLE blog_post (
    id uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    blog_id integer NOT NULL REFERENCES blog(id),
    title varchar(255) NOT NULL,
    body varchar(10000),
    status blog_post_status NOT NULL,
    created_at timestamp NOT NULL
);
```

This schema automatically generates GraphQL types (`Account`, `Blog`, `BlogPost`) with relationships derived from foreign keys.

### Name Inflection

Enable automatic snake_case to camelCase/PascalCase conversion:

```sql
COMMENT ON SCHEMA public IS e'@graphql({"inflect_names": true})';
```

### Querying

Execute a GraphQL query via the `graphql.resolve` function:

```sql
SELECT graphql.resolve($$
    {
      accountCollection(first: 1) {
        edges {
          node {
            id
            email
            blogCollection {
              edges {
                node {
                  name
                  blogPostCollection(filter: { status: { eq: RELEASED } }) {
                    edges {
                      node {
                        title
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
$$);
```

### Features

- Table queries appear as pageable collections on the root `Query` type
- Foreign key relationships create nested query fields automatically
- Mutations support bulk insert, update, and delete
- Filtering, ordering, and pagination are built in
- PostgreSQL Row-Level Security (RLS) policies are respected

### Version Notes

`pg_graphql` 1.6.1 fixes double `NON_NULL` wrapping in `byPk` argument types that could break GraphiQL introspection, fixes mixed introspection/data queries returning partial results incorrectly, and updates documentation around introspection opt-in notices.
