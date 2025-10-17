---
title: "sparql"
linkTitle: "sparql"
description: "Query SPARQL datasource with SQL"
weight: 4470
categories: ["Util"]
width: full
---

Query SPARQL datasource with SQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4470** | {{< badge content="sparql" link="https://github.com/lacanoid/pgsparql" >}} | {{< ext "sparql" "pgsparql" >}} | `1.0` | {{< category "UTIL" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plperl" >}} {{< ext "plperlu" >}} |
|   **See Also**    | {{< ext "pgjq" >}} {{< ext "pgjwt" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/sparql" >}} | `1.0` | {{< badge content="18" color="red" alt="pgsparql_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgsparql_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/sparql" >}} | `1.0` | {{< badge content="18" color="red" alt="postgresql-18-pgsparql" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgsparql` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pgsparql_18" >}}     | {{< pkg "pgsparql_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsparql_17-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgsparql_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsparql_16-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgsparql_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsparql_15-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgsparql_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsparql_14-1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pgsparql_18" >}}     | {{< pkg "pgsparql_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsparql_17-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgsparql_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsparql_16-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgsparql_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsparql_15-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgsparql_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsparql_14-1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pgsparql_18" >}}     | {{< pkg "pgsparql_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsparql_17-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgsparql_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsparql_16-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgsparql_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsparql_15-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgsparql_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsparql_14-1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pgsparql_18" >}}     | {{< pkg "pgsparql_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsparql_17-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgsparql_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsparql_16-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgsparql_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsparql_15-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgsparql_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsparql_14-1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pgsparql" >}}     | {{< pkg "postgresql-17-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-17-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-16-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-15-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-14-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pgsparql" >}}     | {{< pkg "postgresql-17-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-17-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-16-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-15-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-14-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pgsparql" >}}     | {{< pkg "postgresql-17-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-17-pgsparql_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-16-pgsparql_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-15-pgsparql_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-14-pgsparql_1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pgsparql" >}}     | {{< pkg "postgresql-17-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-17-pgsparql_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-16-pgsparql_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-15-pgsparql_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-14-pgsparql_1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pgsparql" >}}     | {{< pkg "postgresql-17-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-17-pgsparql_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-16-pgsparql_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-15-pgsparql_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-14-pgsparql_1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pgsparql" >}}     | {{< pkg "postgresql-17-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-17-pgsparql_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-16-pgsparql_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-15-pgsparql_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pgsparql" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-14-pgsparql_1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsparql_17` | 1.0 | `el8.x86_64` | pigsty | 17.2 KiB | [pgsparql_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsparql_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsparql_17` | 1.0 | `el8.aarch64` | pigsty | 17.2 KiB | [pgsparql_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsparql_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsparql_17` | 1.0 | `el9.aarch64` | pigsty | 17.0 KiB | [pgsparql_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsparql_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgsparql_17` | 1.0 | `el9.x86_64` | pigsty | 17.0 KiB | [pgsparql_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsparql_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pgsparql` | 1.0 | `d12.x86_64` | pigsty | 10.8 KiB | [postgresql-17-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-17-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgsparql` | 1.0 | `d12.aarch64` | pigsty | 10.8 KiB | [postgresql-17-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-17-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgsparql` | 1.0 | `u22.x86_64` | pigsty | 10.5 KiB | [postgresql-17-pgsparql_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-17-pgsparql_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgsparql` | 1.0 | `u22.aarch64` | pigsty | 10.5 KiB | [postgresql-17-pgsparql_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-17-pgsparql_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgsparql` | 1.0 | `u24.x86_64` | pigsty | 10.5 KiB | [postgresql-17-pgsparql_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-17-pgsparql_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgsparql` | 1.0 | `u24.aarch64` | pigsty | 10.5 KiB | [postgresql-17-pgsparql_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-17-pgsparql_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsparql_16` | 1.0 | `el8.x86_64` | pigsty | 17.2 KiB | [pgsparql_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsparql_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsparql_16` | 1.0 | `el8.aarch64` | pigsty | 17.2 KiB | [pgsparql_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsparql_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsparql_16` | 1.0 | `el9.x86_64` | pigsty | 17.0 KiB | [pgsparql_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsparql_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsparql_16` | 1.0 | `el9.aarch64` | pigsty | 17.0 KiB | [pgsparql_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsparql_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pgsparql` | 1.0 | `d12.x86_64` | pigsty | 10.8 KiB | [postgresql-16-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-16-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgsparql` | 1.0 | `d12.aarch64` | pigsty | 10.8 KiB | [postgresql-16-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-16-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgsparql` | 1.0 | `u22.aarch64` | pigsty | 10.5 KiB | [postgresql-16-pgsparql_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-16-pgsparql_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgsparql` | 1.0 | `u22.x86_64` | pigsty | 10.5 KiB | [postgresql-16-pgsparql_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-16-pgsparql_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgsparql` | 1.0 | `u24.x86_64` | pigsty | 10.5 KiB | [postgresql-16-pgsparql_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-16-pgsparql_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgsparql` | 1.0 | `u24.aarch64` | pigsty | 10.5 KiB | [postgresql-16-pgsparql_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-16-pgsparql_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsparql_15` | 1.0 | `el8.x86_64` | pigsty | 17.2 KiB | [pgsparql_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsparql_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsparql_15` | 1.0 | `el8.aarch64` | pigsty | 17.2 KiB | [pgsparql_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsparql_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsparql_15` | 1.0 | `el9.x86_64` | pigsty | 17.0 KiB | [pgsparql_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsparql_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsparql_15` | 1.0 | `el9.aarch64` | pigsty | 17.0 KiB | [pgsparql_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsparql_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgsparql` | 1.0 | `d12.aarch64` | pigsty | 10.8 KiB | [postgresql-15-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-15-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgsparql` | 1.0 | `d12.x86_64` | pigsty | 10.8 KiB | [postgresql-15-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-15-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgsparql` | 1.0 | `u22.aarch64` | pigsty | 10.5 KiB | [postgresql-15-pgsparql_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-15-pgsparql_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgsparql` | 1.0 | `u22.x86_64` | pigsty | 10.5 KiB | [postgresql-15-pgsparql_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-15-pgsparql_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgsparql` | 1.0 | `u24.x86_64` | pigsty | 10.5 KiB | [postgresql-15-pgsparql_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-15-pgsparql_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgsparql` | 1.0 | `u24.aarch64` | pigsty | 10.5 KiB | [postgresql-15-pgsparql_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-15-pgsparql_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsparql_14` | 1.0 | `el8.x86_64` | pigsty | 17.2 KiB | [pgsparql_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsparql_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsparql_14` | 1.0 | `el8.aarch64` | pigsty | 17.2 KiB | [pgsparql_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsparql_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsparql_14` | 1.0 | `el9.x86_64` | pigsty | 17.0 KiB | [pgsparql_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsparql_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsparql_14` | 1.0 | `el9.aarch64` | pigsty | 17.0 KiB | [pgsparql_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsparql_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgsparql` | 1.0 | `d12.x86_64` | pigsty | 10.8 KiB | [postgresql-14-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-14-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgsparql` | 1.0 | `d12.aarch64` | pigsty | 10.8 KiB | [postgresql-14-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-14-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgsparql` | 1.0 | `u22.x86_64` | pigsty | 10.5 KiB | [postgresql-14-pgsparql_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-14-pgsparql_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgsparql` | 1.0 | `u22.aarch64` | pigsty | 10.5 KiB | [postgresql-14-pgsparql_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-14-pgsparql_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgsparql` | 1.0 | `u24.x86_64` | pigsty | 10.5 KiB | [postgresql-14-pgsparql_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-14-pgsparql_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgsparql` | 1.0 | `u24.aarch64` | pigsty | 10.5 KiB | [postgresql-14-pgsparql_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-14-pgsparql_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsparql_13` | 1.0 | `el8.aarch64` | pigsty | 17.2 KiB | [pgsparql_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsparql_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsparql_13` | 1.0 | `el8.x86_64` | pigsty | 17.2 KiB | [pgsparql_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsparql_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsparql_13` | 1.0 | `el9.aarch64` | pigsty | 17.0 KiB | [pgsparql_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsparql_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgsparql_13` | 1.0 | `el9.x86_64` | pigsty | 17.0 KiB | [pgsparql_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsparql_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pgsparql` | 1.0 | `d12.aarch64` | pigsty | 10.8 KiB | [postgresql-13-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-13-pgsparql_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgsparql` | 1.0 | `d12.x86_64` | pigsty | 10.8 KiB | [postgresql-13-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsparql/postgresql-13-pgsparql_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgsparql` | 1.0 | `u22.aarch64` | pigsty | 10.5 KiB | [postgresql-13-pgsparql_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-13-pgsparql_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgsparql` | 1.0 | `u22.x86_64` | pigsty | 10.5 KiB | [postgresql-13-pgsparql_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsparql/postgresql-13-pgsparql_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgsparql` | 1.0 | `u24.aarch64` | pigsty | 10.5 KiB | [postgresql-13-pgsparql_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-13-pgsparql_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pgsparql` | 1.0 | `u24.x86_64` | pigsty | 10.5 KiB | [postgresql-13-pgsparql_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsparql/postgresql-13-pgsparql_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/lacanoid/pgsparql" title="Repository" icon="github" subtitle="github.com/lacanoid/pgsparql" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgsparql-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get sparql; # get sparql source code
pig build dep sparql; # install build dependencies
pig build pkg sparql; # build extension rpm or deb
pig build ext sparql; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install sparql; # install by extension name, for the current active PG version
pig ext install pgsparql; # install via package alias, for the active PG version
pig ext install sparql -v 18;   # install for PG 18
pig ext install sparql -v 17;   # install for PG 17
pig ext install sparql -v 16;   # install for PG 16
pig ext install sparql -v 15;   # install for PG 15
pig ext install sparql -v 14;   # install for PG 14
pig ext install sparql -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION sparql CASCADE SCHEMA sparql;
```

