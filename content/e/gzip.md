---
title: "gzip"
linkTitle: "gzip"
description: "gzip and gunzip functions."
weight: 4010
categories: ["Util"]
width: full
---

gzip and gunzip functions.

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4010** | {{< badge content="gzip" link="https://github.com/pramsey/pgsql-gzip" >}} | {{< ext "gzip" "pg_gzip" >}} | `1.0.1` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pgjq" >}} {{< ext "pgjwt" >}} {{< ext "pg_smtp_client" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/gzip" >}} | `1.0.1` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgsql_gzip_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/gzip" >}} | `1.0.1` | {{< badge content="18" color="red" alt="postgresql-18-gzip" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-gzip` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pgsql_gzip_18" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsql_gzip_18-1.0.0-6PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsql_gzip_17" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_gzip_17-1.0.0-3PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsql_gzip_16" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_gzip_16-1.0.0-2PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsql_gzip_15" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_gzip_15-1.0.0-2PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsql_gzip_14" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_gzip_14-1.0.0-2PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pgsql_gzip_18" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsql_gzip_18-1.0.0-6PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsql_gzip_17" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_gzip_17-1.0.0-3PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsql_gzip_16" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_gzip_16-1.0.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsql_gzip_15" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_gzip_15-1.0.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsql_gzip_14" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_gzip_14-1.0.0-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pgsql_gzip_18" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsql_gzip_18-1.0.0-6PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsql_gzip_17" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_gzip_17-1.0.0-3PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsql_gzip_16" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_gzip_16-1.0.0-2PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsql_gzip_15" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_gzip_15-1.0.0-2PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsql_gzip_14" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_gzip_14-1.0.0-2PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pgsql_gzip_18" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsql_gzip_18-1.0.0-6PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsql_gzip_17" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_gzip_17-1.0.0-3PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsql_gzip_16" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_gzip_16-1.0.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsql_gzip_15" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_gzip_15-1.0.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsql_gzip_14" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_gzip_14-1.0.0-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-gzip" >}}     | {{< pkg "postgresql-17-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-gzip" >}}     | {{< pkg "postgresql-17-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-gzip" >}}     | {{< pkg "postgresql-17-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-gzip" >}}     | {{< pkg "postgresql-17-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-gzip" >}}     | {{< pkg "postgresql-17-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-gzip" >}}     | {{< pkg "postgresql-17-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-gzip" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsql_gzip_18` | 1.0.0 | `el8.x86_64` | pgdg | 12.9 KiB | [pgsql_gzip_18-1.0.0-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsql_gzip_18-1.0.0-6PGDG.rhel8.x86_64.rpm) |
| `pgsql_gzip_18` | 1.0.0 | `el8.aarch64` | pgdg | 12.9 KiB | [pgsql_gzip_18-1.0.0-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsql_gzip_18-1.0.0-6PGDG.rhel8.aarch64.rpm) |
| `pgsql_gzip_18` | 1.0.0 | `el9.x86_64` | pgdg | 12.9 KiB | [pgsql_gzip_18-1.0.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsql_gzip_18-1.0.0-6PGDG.rhel9.x86_64.rpm) |
| `pgsql_gzip_18` | 1.0.0 | `el9.aarch64` | pgdg | 12.6 KiB | [pgsql_gzip_18-1.0.0-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsql_gzip_18-1.0.0-6PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsql_gzip_17` | 1.0.0 | `el8.x86_64` | pgdg | 12.7 KiB | [pgsql_gzip_17-1.0.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_gzip_17-1.0.0-3PGDG.rhel8.x86_64.rpm) |
| `pgsql_gzip_17` | 1.0.0 | `el8.aarch64` | pgdg | 12.7 KiB | [pgsql_gzip_17-1.0.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_gzip_17-1.0.0-3PGDG.rhel8.aarch64.rpm) |
| `pgsql_gzip_17` | 1.0.0 | `el8.aarch64` | pigsty | 13.9 KiB | [pgsql_gzip_17-1.0.0-9PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_gzip_17-1.0.0-9PIGSTY.el8.aarch64.rpm) |
| `pgsql_gzip_17` | 1.0.0 | `el9.aarch64` | pigsty | 13.7 KiB | [pgsql_gzip_17-1.0.0-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_gzip_17-1.0.0-4PIGSTY.el9.aarch64.rpm) |
| `pgsql_gzip_17` | 1.0.0 | `el9.x86_64` | pgdg | 12.7 KiB | [pgsql_gzip_17-1.0.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_gzip_17-1.0.0-3PGDG.rhel9.x86_64.rpm) |
| `pgsql_gzip_17` | 1.0.0 | `el9.aarch64` | pgdg | 12.5 KiB | [pgsql_gzip_17-1.0.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_gzip_17-1.0.0-3PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-gzip` | 1.0.1 | `d12.aarch64` | pigsty | 12.9 KiB | [postgresql-17-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-gzip` | 1.0.1 | `d12.x86_64` | pigsty | 12.9 KiB | [postgresql-17-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-gzip` | 1.0.1 | `u22.aarch64` | pigsty | 13.2 KiB | [postgresql-17-gzip_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-gzip` | 1.0.1 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-17-gzip_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-gzip` | 1.0.1 | `u24.aarch64` | pigsty | 12.8 KiB | [postgresql-17-gzip_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-gzip` | 1.0.1 | `u24.x86_64` | pigsty | 12.9 KiB | [postgresql-17-gzip_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsql_gzip_16` | 1.0.0 | `el8.x86_64` | pgdg | 12.6 KiB | [pgsql_gzip_16-1.0.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_gzip_16-1.0.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_gzip_16` | 1.0.0 | `el8.aarch64` | pigsty | 13.9 KiB | [pgsql_gzip_16-1.0.0-9PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_gzip_16-1.0.0-9PIGSTY.el8.aarch64.rpm) |
| `pgsql_gzip_16` | 1.0.0 | `el8.aarch64` | pgdg | 12.5 KiB | [pgsql_gzip_16-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_gzip_16-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_gzip_16` | 1.0.0 | `el9.x86_64` | pgdg | 12.6 KiB | [pgsql_gzip_16-1.0.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_gzip_16-1.0.0-2PGDG.rhel9.x86_64.rpm) |
| `pgsql_gzip_16` | 1.0.0 | `el9.aarch64` | pigsty | 13.7 KiB | [pgsql_gzip_16-1.0.0-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_gzip_16-1.0.0-4PIGSTY.el9.aarch64.rpm) |
| `pgsql_gzip_16` | 1.0.0 | `el9.aarch64` | pgdg | 12.2 KiB | [pgsql_gzip_16-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_gzip_16-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-gzip` | 1.0.1 | `d12.x86_64` | pigsty | 12.9 KiB | [postgresql-16-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-gzip` | 1.0.1 | `d12.aarch64` | pigsty | 12.9 KiB | [postgresql-16-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-gzip` | 1.0.1 | `u22.aarch64` | pigsty | 13.2 KiB | [postgresql-16-gzip_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-gzip` | 1.0.1 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-16-gzip_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-gzip` | 1.0.1 | `u24.x86_64` | pigsty | 12.9 KiB | [postgresql-16-gzip_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-gzip` | 1.0.1 | `u24.aarch64` | pigsty | 12.8 KiB | [postgresql-16-gzip_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsql_gzip_15` | 1.0.0 | `el8.x86_64` | pgdg | 12.6 KiB | [pgsql_gzip_15-1.0.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_gzip_15-1.0.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_gzip_15` | 1.0.0 | `el8.aarch64` | pgdg | 12.5 KiB | [pgsql_gzip_15-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_gzip_15-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_gzip_15` | 1.0.0 | `el8.aarch64` | pigsty | 13.9 KiB | [pgsql_gzip_15-1.0.0-9PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_gzip_15-1.0.0-9PIGSTY.el8.aarch64.rpm) |
| `pgsql_gzip_15` | 1.0.0 | `el9.x86_64` | pgdg | 12.6 KiB | [pgsql_gzip_15-1.0.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_gzip_15-1.0.0-2PGDG.rhel9.x86_64.rpm) |
| `pgsql_gzip_15` | 1.0.0 | `el9.aarch64` | pigsty | 13.7 KiB | [pgsql_gzip_15-1.0.0-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_gzip_15-1.0.0-4PIGSTY.el9.aarch64.rpm) |
| `pgsql_gzip_15` | 1.0.0 | `el9.aarch64` | pgdg | 12.2 KiB | [pgsql_gzip_15-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_gzip_15-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-gzip` | 1.0.1 | `d12.aarch64` | pigsty | 12.9 KiB | [postgresql-15-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-gzip` | 1.0.1 | `d12.x86_64` | pigsty | 12.9 KiB | [postgresql-15-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-gzip` | 1.0.1 | `u22.aarch64` | pigsty | 13.2 KiB | [postgresql-15-gzip_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-gzip` | 1.0.1 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-15-gzip_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-gzip` | 1.0.1 | `u24.x86_64` | pigsty | 12.9 KiB | [postgresql-15-gzip_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-gzip` | 1.0.1 | `u24.aarch64` | pigsty | 12.8 KiB | [postgresql-15-gzip_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsql_gzip_14` | 1.0.0 | `el8.x86_64` | pgdg | 12.6 KiB | [pgsql_gzip_14-1.0.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_gzip_14-1.0.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_gzip_14` | 1.0.0 | `el8.aarch64` | pigsty | 13.9 KiB | [pgsql_gzip_14-1.0.0-9PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_gzip_14-1.0.0-9PIGSTY.el8.aarch64.rpm) |
| `pgsql_gzip_14` | 1.0.0 | `el8.aarch64` | pgdg | 12.5 KiB | [pgsql_gzip_14-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_gzip_14-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_gzip_14` | 1.0.0 | `el9.aarch64` | pigsty | 13.7 KiB | [pgsql_gzip_14-1.0.0-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_gzip_14-1.0.0-4PIGSTY.el9.aarch64.rpm) |
| `pgsql_gzip_14` | 1.0.0 | `el9.x86_64` | pgdg | 12.6 KiB | [pgsql_gzip_14-1.0.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_gzip_14-1.0.0-2PGDG.rhel9.x86_64.rpm) |
| `pgsql_gzip_14` | 1.0.0 | `el9.aarch64` | pgdg | 12.2 KiB | [pgsql_gzip_14-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_gzip_14-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-14-gzip` | 1.0.1 | `d12.aarch64` | pigsty | 12.9 KiB | [postgresql-14-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-gzip` | 1.0.1 | `d12.x86_64` | pigsty | 12.8 KiB | [postgresql-14-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-gzip` | 1.0.1 | `u22.aarch64` | pigsty | 13.1 KiB | [postgresql-14-gzip_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-gzip` | 1.0.1 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-14-gzip_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-gzip` | 1.0.1 | `u24.aarch64` | pigsty | 12.8 KiB | [postgresql-14-gzip_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-gzip` | 1.0.1 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-14-gzip_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsql_gzip_13` | 1.0.0 | `el8.aarch64` | pigsty | 13.9 KiB | [pgsql_gzip_13-1.0.0-9PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_gzip_13-1.0.0-9PIGSTY.el8.aarch64.rpm) |
| `pgsql_gzip_13` | 1.0.0 | `el8.aarch64` | pgdg | 12.5 KiB | [pgsql_gzip_13-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_gzip_13-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_gzip_13` | 1.0.0 | `el8.x86_64` | pgdg | 12.5 KiB | [pgsql_gzip_13-1.0.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_gzip_13-1.0.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_gzip_13` | 1.0.0 | `el9.aarch64` | pigsty | 13.7 KiB | [pgsql_gzip_13-1.0.0-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_gzip_13-1.0.0-4PIGSTY.el9.aarch64.rpm) |
| `pgsql_gzip_13` | 1.0.0 | `el9.x86_64` | pgdg | 12.6 KiB | [pgsql_gzip_13-1.0.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_gzip_13-1.0.0-2PGDG.rhel9.x86_64.rpm) |
| `pgsql_gzip_13` | 1.0.0 | `el9.aarch64` | pgdg | 12.2 KiB | [pgsql_gzip_13-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_gzip_13-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-13-gzip` | 1.0.1 | `d12.x86_64` | pigsty | 12.6 KiB | [postgresql-13-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-gzip` | 1.0.1 | `d12.aarch64` | pigsty | 12.7 KiB | [postgresql-13-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-gzip` | 1.0.1 | `u22.x86_64` | pigsty | 13.0 KiB | [postgresql-13-gzip_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-gzip` | 1.0.1 | `u22.aarch64` | pigsty | 13.1 KiB | [postgresql-13-gzip_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-gzip` | 1.0.1 | `u24.x86_64` | pigsty | 12.7 KiB | [postgresql-13-gzip_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-gzip` | 1.0.1 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-13-gzip_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pramsey/pgsql-gzip" title="Repository" icon="github" subtitle="github.com/pramsey/pgsql-gzip" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgsql-gzip-1.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get gzip; # get gzip source code
pig build dep gzip; # install build dependencies
pig build pkg gzip; # build extension rpm or deb
pig build ext gzip; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install gzip; # install by extension name, for the current active PG version
pig ext install pg_gzip; # install via package alias, for the active PG version
pig ext install gzip -v 18;   # install for PG 18
pig ext install gzip -v 17;   # install for PG 17
pig ext install gzip -v 16;   # install for PG 16
pig ext install gzip -v 15;   # install for PG 15
pig ext install gzip -v 14;   # install for PG 14
pig ext install gzip -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION gzip;
```




--------

## Usage

Sometimes you just need to compress your `bytea` object before you return it to the client.

Sometimes you receive a compressed `bytea` from the client, and you have to uncompress it before you can work with it.

This extension is for that.

This extension is **not** for storage compression. PostgreSQL already does [tuple compression](https://www.postgresql.org/docs/current/storage-toast.html) on the fly if your tuple gets large enough, manually pre-compressing your data using this function won't make things smaller.


* `gzip(uncompressed BYTEA, [compression_level INTEGER])` returns `BYTEA`
* `gzip(uncompressed TEXT, [compression_level INTEGER])` returns `BYTEA`
* `gunzip(compressed BYTEA)` returns `BYTEA`


### Examples

    > SELECT gzip('this is my this is my this is my this is my text');

                                       gzip
    --------------------------------------------------------------------------
     \x1f8b08000000000000132bc9c82c5600a2dc4a851282ccd48a12002e7a22ff30000000

Wait, what, the compressed output is longer?!? No, it only **looks** that way, because in hex every byte is represented with two hex digits. The original string looks like this in hex:

    > SELECT 'this is my this is my this is my this is my text'::bytea;

                                                   bytea
    ----------------------------------------------------------------------------------------------------
     \x74686973206973206d792074686973206973206d792074686973206973206d792074686973206973206d792074657874

For really long, repetitive things, compression naturally works like a charm:

    > SELECT gzip(repeat('this is my ', 100));

                                                   bytea
    ----------------------------------------------------------------------------------------------------
     \x1f8b08000000000000132bc9c82c5600a2dc4a859251e628739439ca24970900d1341c5c4c040000

To convert a `bytea` back into an equivalent `text` you must use the `encode()` function with the `escape` encoding.

    > SELECT encode('test text'::bytea, 'escape');
       encode
    -----------
     test text

    > SELECT encode(gunzip(gzip('this text has been compressed and then decompressed')), 'escape')

                          encode
    -----------------------------------------------------
     this text has been compressed and then decompressed

