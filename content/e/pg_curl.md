---
title: "pg_curl"
linkTitle: "pg_curl"
description: "Run curl actions for data transfer in URL syntax"
weight: 4090
categories: ["Util"]
width: full
---

Run curl actions for data transfer in URL syntax

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4090** | {{< badge content="pg_curl" link="https://github.com/RekGRpth/pg_curl" >}} | {{< ext "pg_curl" "pg_curl" >}} | `2.4.5` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pgjwt" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "pgjq" >}} {{< ext "pg_smtp_client" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_curl" >}} | `2.4.5` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_curl_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_curl" >}} | `2.4.5` | {{< badge content="18" color="red" alt="postgresql-18-pg-curl" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-curl` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_curl_18" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_curl_18-2.4.4-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_curl_17" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_curl_17-2.4.4-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_curl_16" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_curl_16-2.4.4-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_curl_15" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_curl_15-2.4.4-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_curl_14" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_curl_14-2.4.4-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_curl_18" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_curl_18-2.4.4-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_curl_17" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_curl_17-2.4.4-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_curl_16" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_curl_16-2.4.4-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_curl_15" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_curl_15-2.4.4-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_curl_14" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_curl_14-2.4.4-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_curl_18" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_curl_18-2.4.4-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_curl_17" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_curl_17-2.4.4-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_curl_16" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_curl_16-2.4.4-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_curl_15" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_curl_15-2.4.4-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_curl_14" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_curl_14-2.4.4-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_curl_18" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_curl_18-2.4.4-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_curl_17" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_curl_17-2.4.4-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_curl_16" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_curl_16-2.4.4-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_curl_15" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_curl_15-2.4.4-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_curl_14" "2.4.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_curl_14-2.4.4-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-curl" >}}     | {{< pkg "postgresql-17-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-curl" >}}     | {{< pkg "postgresql-17-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-curl" >}}     | {{< pkg "postgresql-17-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-curl" >}}     | {{< pkg "postgresql-17-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-curl" >}}     | {{< pkg "postgresql-17-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-curl" >}}     | {{< pkg "postgresql-17-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-curl" "2.4.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_curl_18` | 2.4.4 | `el8.aarch64` | pgdg | 42.1 KiB | [pg_curl_18-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_curl_18-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_18` | 2.4.4 | `el8.x86_64` | pgdg | 43.8 KiB | [pg_curl_18-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_curl_18-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_18` | 2.4.4 | `el9.aarch64` | pgdg | 44.0 KiB | [pg_curl_18-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_curl_18-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_18` | 2.4.4 | `el9.x86_64` | pgdg | 45.5 KiB | [pg_curl_18-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_curl_18-2.4.4-1PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_curl_17` | 2.4.5 | `el8.x86_64` | pigsty | 45.5 KiB | [pg_curl_17-2.4.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_17-2.4.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_curl_17` | 2.4.5 | `el8.aarch64` | pigsty | 43.9 KiB | [pg_curl_17-2.4.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_17-2.4.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_curl_17` | 2.4.4 | `el8.aarch64` | pgdg | 42.1 KiB | [pg_curl_17-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_curl_17-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_17` | 2.4.4 | `el8.x86_64` | pgdg | 43.8 KiB | [pg_curl_17-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_curl_17-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_17` | 2.4.3 | `el8.aarch64` | pgdg | 41.9 KiB | [pg_curl_17-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_curl_17-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_17` | 2.4.3 | `el8.x86_64` | pgdg | 43.7 KiB | [pg_curl_17-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_curl_17-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_17` | 2.4 | `el8.aarch64` | pigsty | 45.9 KiB | [pg_curl_17-2.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_17-2.4-2PIGSTY.el8.aarch64.rpm) |
| `pg_curl_17` | 2.4 | `el8.x86_64` | pigsty | 47.7 KiB | [pg_curl_17-2.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_17-2.4-2PIGSTY.el8.x86_64.rpm) |
| `pg_curl_17` | 2.4.5 | `el9.aarch64` | pigsty | 45.7 KiB | [pg_curl_17-2.4.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_17-2.4.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_curl_17` | 2.4.5 | `el9.x86_64` | pigsty | 47.5 KiB | [pg_curl_17-2.4.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_17-2.4.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_curl_17` | 2.4.4 | `el9.x86_64` | pgdg | 45.6 KiB | [pg_curl_17-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_curl_17-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_17` | 2.4.4 | `el9.aarch64` | pgdg | 43.8 KiB | [pg_curl_17-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_curl_17-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_17` | 2.4.3 | `el9.x86_64` | pgdg | 45.6 KiB | [pg_curl_17-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_curl_17-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_17` | 2.4.3 | `el9.aarch64` | pgdg | 44.0 KiB | [pg_curl_17-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_curl_17-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_17` | 2.4 | `el9.aarch64` | pigsty | 47.6 KiB | [pg_curl_17-2.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_17-2.4-2PIGSTY.el9.aarch64.rpm) |
| `pg_curl_17` | 2.4 | `el9.x86_64` | pigsty | 49.0 KiB | [pg_curl_17-2.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_17-2.4-2PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-curl` | 2.4.5 | `d12.x86_64` | pigsty | 99.8 KiB | [postgresql-17-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-curl` | 2.4.5 | `d12.aarch64` | pigsty | 98.5 KiB | [postgresql-17-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-curl` | 2.4.5 | `u22.aarch64` | pigsty | 115.6 KiB | [postgresql-17-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-curl` | 2.4.5 | `u22.x86_64` | pigsty | 117.3 KiB | [postgresql-17-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-curl` | 2.4.5 | `u24.x86_64` | pigsty | 108.2 KiB | [postgresql-17-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-curl` | 2.4.5 | `u24.aarch64` | pigsty | 107.7 KiB | [postgresql-17-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_curl_16` | 2.4.5 | `el8.aarch64` | pigsty | 43.9 KiB | [pg_curl_16-2.4.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_16-2.4.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_curl_16` | 2.4.5 | `el8.x86_64` | pigsty | 45.5 KiB | [pg_curl_16-2.4.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_16-2.4.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_curl_16` | 2.4.4 | `el8.aarch64` | pgdg | 42.1 KiB | [pg_curl_16-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_curl_16-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_16` | 2.4.4 | `el8.x86_64` | pgdg | 43.8 KiB | [pg_curl_16-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_curl_16-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_16` | 2.4.3 | `el8.x86_64` | pgdg | 43.8 KiB | [pg_curl_16-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_curl_16-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_16` | 2.4.3 | `el8.aarch64` | pgdg | 41.9 KiB | [pg_curl_16-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_curl_16-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_16` | 2.4 | `el8.aarch64` | pigsty | 45.9 KiB | [pg_curl_16-2.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_16-2.4-2PIGSTY.el8.aarch64.rpm) |
| `pg_curl_16` | 2.4 | `el8.x86_64` | pigsty | 47.7 KiB | [pg_curl_16-2.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_16-2.4-2PIGSTY.el8.x86_64.rpm) |
| `pg_curl_16` | 2.4.5 | `el9.aarch64` | pigsty | 45.7 KiB | [pg_curl_16-2.4.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_16-2.4.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_curl_16` | 2.4.5 | `el9.x86_64` | pigsty | 47.5 KiB | [pg_curl_16-2.4.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_16-2.4.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_curl_16` | 2.4.4 | `el9.x86_64` | pgdg | 45.7 KiB | [pg_curl_16-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_curl_16-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_16` | 2.4.4 | `el9.aarch64` | pgdg | 44.0 KiB | [pg_curl_16-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_curl_16-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_16` | 2.4.3 | `el9.x86_64` | pgdg | 45.5 KiB | [pg_curl_16-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_curl_16-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_16` | 2.4.3 | `el9.aarch64` | pgdg | 44.1 KiB | [pg_curl_16-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_curl_16-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_16` | 2.4 | `el9.aarch64` | pigsty | 47.6 KiB | [pg_curl_16-2.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_16-2.4-2PIGSTY.el9.aarch64.rpm) |
| `pg_curl_16` | 2.4 | `el9.x86_64` | pigsty | 49.1 KiB | [pg_curl_16-2.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_16-2.4-2PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-pg-curl` | 2.4.5 | `d12.aarch64` | pigsty | 98.5 KiB | [postgresql-16-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-curl` | 2.4.5 | `d12.x86_64` | pigsty | 99.9 KiB | [postgresql-16-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-curl` | 2.4.5 | `u22.aarch64` | pigsty | 115.6 KiB | [postgresql-16-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-curl` | 2.4.5 | `u22.x86_64` | pigsty | 117.3 KiB | [postgresql-16-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-curl` | 2.4.5 | `u24.x86_64` | pigsty | 108.1 KiB | [postgresql-16-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-curl` | 2.4.5 | `u24.aarch64` | pigsty | 107.8 KiB | [postgresql-16-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_curl_15` | 2.4.5 | `el8.x86_64` | pigsty | 45.5 KiB | [pg_curl_15-2.4.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_15-2.4.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_curl_15` | 2.4.5 | `el8.aarch64` | pigsty | 43.9 KiB | [pg_curl_15-2.4.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_15-2.4.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_curl_15` | 2.4.4 | `el8.x86_64` | pgdg | 43.8 KiB | [pg_curl_15-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_curl_15-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_15` | 2.4.4 | `el8.aarch64` | pgdg | 42.1 KiB | [pg_curl_15-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_curl_15-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_15` | 2.4.3 | `el8.x86_64` | pgdg | 43.7 KiB | [pg_curl_15-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_curl_15-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_15` | 2.4.3 | `el8.aarch64` | pgdg | 41.9 KiB | [pg_curl_15-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_curl_15-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_15` | 2.4 | `el8.x86_64` | pigsty | 47.7 KiB | [pg_curl_15-2.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_15-2.4-2PIGSTY.el8.x86_64.rpm) |
| `pg_curl_15` | 2.4 | `el8.aarch64` | pigsty | 45.9 KiB | [pg_curl_15-2.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_15-2.4-2PIGSTY.el8.aarch64.rpm) |
| `pg_curl_15` | 2.4.5 | `el9.x86_64` | pigsty | 47.4 KiB | [pg_curl_15-2.4.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_15-2.4.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_curl_15` | 2.4.5 | `el9.aarch64` | pigsty | 45.6 KiB | [pg_curl_15-2.4.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_15-2.4.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_curl_15` | 2.4.4 | `el9.x86_64` | pgdg | 45.6 KiB | [pg_curl_15-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_curl_15-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_15` | 2.4.4 | `el9.aarch64` | pgdg | 44.0 KiB | [pg_curl_15-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_curl_15-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_15` | 2.4.3 | `el9.aarch64` | pgdg | 44.0 KiB | [pg_curl_15-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_curl_15-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_15` | 2.4.3 | `el9.x86_64` | pgdg | 45.6 KiB | [pg_curl_15-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_curl_15-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_15` | 2.4 | `el9.aarch64` | pigsty | 47.7 KiB | [pg_curl_15-2.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_15-2.4-2PIGSTY.el9.aarch64.rpm) |
| `pg_curl_15` | 2.4 | `el9.x86_64` | pigsty | 49.1 KiB | [pg_curl_15-2.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_15-2.4-2PIGSTY.el9.x86_64.rpm) |
| `postgresql-15-pg-curl` | 2.4.5 | `d12.aarch64` | pigsty | 98.1 KiB | [postgresql-15-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-curl` | 2.4.5 | `d12.x86_64` | pigsty | 99.6 KiB | [postgresql-15-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-curl` | 2.4.5 | `u22.x86_64` | pigsty | 117.3 KiB | [postgresql-15-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-curl` | 2.4.5 | `u22.aarch64` | pigsty | 115.7 KiB | [postgresql-15-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-curl` | 2.4.5 | `u24.x86_64` | pigsty | 107.8 KiB | [postgresql-15-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-curl` | 2.4.5 | `u24.aarch64` | pigsty | 107.7 KiB | [postgresql-15-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_curl_14` | 2.4.5 | `el8.x86_64` | pigsty | 45.5 KiB | [pg_curl_14-2.4.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_14-2.4.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_curl_14` | 2.4.5 | `el8.aarch64` | pigsty | 43.9 KiB | [pg_curl_14-2.4.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_14-2.4.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_curl_14` | 2.4.4 | `el8.aarch64` | pgdg | 42.1 KiB | [pg_curl_14-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_curl_14-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_14` | 2.4.4 | `el8.x86_64` | pgdg | 43.8 KiB | [pg_curl_14-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_curl_14-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_14` | 2.4.3 | `el8.aarch64` | pgdg | 41.9 KiB | [pg_curl_14-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_curl_14-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_14` | 2.4.3 | `el8.x86_64` | pgdg | 43.7 KiB | [pg_curl_14-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_curl_14-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_14` | 2.4 | `el8.x86_64` | pigsty | 47.7 KiB | [pg_curl_14-2.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_14-2.4-2PIGSTY.el8.x86_64.rpm) |
| `pg_curl_14` | 2.4 | `el8.aarch64` | pigsty | 45.9 KiB | [pg_curl_14-2.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_14-2.4-2PIGSTY.el8.aarch64.rpm) |
| `pg_curl_14` | 2.4.5 | `el9.x86_64` | pigsty | 47.7 KiB | [pg_curl_14-2.4.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_14-2.4.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_curl_14` | 2.4.5 | `el9.aarch64` | pigsty | 45.7 KiB | [pg_curl_14-2.4.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_14-2.4.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_curl_14` | 2.4.4 | `el9.aarch64` | pgdg | 43.8 KiB | [pg_curl_14-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_curl_14-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_14` | 2.4.4 | `el9.x86_64` | pgdg | 45.5 KiB | [pg_curl_14-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_curl_14-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_14` | 2.4.3 | `el9.x86_64` | pgdg | 45.5 KiB | [pg_curl_14-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_curl_14-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_14` | 2.4.3 | `el9.aarch64` | pgdg | 44.0 KiB | [pg_curl_14-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_curl_14-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_14` | 2.4 | `el9.x86_64` | pigsty | 49.0 KiB | [pg_curl_14-2.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_14-2.4-2PIGSTY.el9.x86_64.rpm) |
| `pg_curl_14` | 2.4 | `el9.aarch64` | pigsty | 47.6 KiB | [pg_curl_14-2.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_14-2.4-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-curl` | 2.4.5 | `d12.x86_64` | pigsty | 99.6 KiB | [postgresql-14-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-curl` | 2.4.5 | `d12.aarch64` | pigsty | 98.0 KiB | [postgresql-14-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-curl` | 2.4.5 | `u22.aarch64` | pigsty | 115.7 KiB | [postgresql-14-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-curl` | 2.4.5 | `u22.x86_64` | pigsty | 117.4 KiB | [postgresql-14-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-curl` | 2.4.5 | `u24.aarch64` | pigsty | 107.5 KiB | [postgresql-14-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-curl` | 2.4.5 | `u24.x86_64` | pigsty | 107.8 KiB | [postgresql-14-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_curl_13` | 2.4.5 | `el8.aarch64` | pigsty | 43.8 KiB | [pg_curl_13-2.4.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_13-2.4.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_curl_13` | 2.4.5 | `el8.x86_64` | pigsty | 44.9 KiB | [pg_curl_13-2.4.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_13-2.4.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_curl_13` | 2.4.4 | `el8.aarch64` | pgdg | 42.0 KiB | [pg_curl_13-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_curl_13-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_13` | 2.4.4 | `el8.x86_64` | pgdg | 43.1 KiB | [pg_curl_13-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_curl_13-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_13` | 2.4.3 | `el8.aarch64` | pgdg | 41.9 KiB | [pg_curl_13-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_curl_13-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_13` | 2.4.3 | `el8.x86_64` | pgdg | 43.0 KiB | [pg_curl_13-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_curl_13-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_13` | 2.4 | `el8.x86_64` | pigsty | 47.0 KiB | [pg_curl_13-2.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_13-2.4-2PIGSTY.el8.x86_64.rpm) |
| `pg_curl_13` | 2.4 | `el8.aarch64` | pigsty | 45.9 KiB | [pg_curl_13-2.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_13-2.4-2PIGSTY.el8.aarch64.rpm) |
| `pg_curl_13` | 2.4.5 | `el9.x86_64` | pigsty | 47.4 KiB | [pg_curl_13-2.4.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_13-2.4.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_curl_13` | 2.4.5 | `el9.aarch64` | pigsty | 45.7 KiB | [pg_curl_13-2.4.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_13-2.4.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_curl_13` | 2.4.4 | `el9.x86_64` | pgdg | 45.6 KiB | [pg_curl_13-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_curl_13-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_13` | 2.4.4 | `el9.aarch64` | pgdg | 44.2 KiB | [pg_curl_13-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_curl_13-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_13` | 2.4.3 | `el9.x86_64` | pgdg | 45.5 KiB | [pg_curl_13-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_curl_13-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_13` | 2.4.3 | `el9.aarch64` | pgdg | 44.1 KiB | [pg_curl_13-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_curl_13-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_13` | 2.4 | `el9.aarch64` | pigsty | 47.7 KiB | [pg_curl_13-2.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_13-2.4-2PIGSTY.el9.aarch64.rpm) |
| `pg_curl_13` | 2.4 | `el9.x86_64` | pigsty | 49.1 KiB | [pg_curl_13-2.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_13-2.4-2PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-curl` | 2.4.5 | `d12.aarch64` | pigsty | 98.4 KiB | [postgresql-13-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-13-pg-curl_2.4.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-curl` | 2.4.5 | `d12.x86_64` | pigsty | 100.1 KiB | [postgresql-13-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-13-pg-curl_2.4.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-curl` | 2.4.5 | `u22.aarch64` | pigsty | 115.2 KiB | [postgresql-13-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-13-pg-curl_2.4.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-curl` | 2.4.5 | `u22.x86_64` | pigsty | 117.1 KiB | [postgresql-13-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-13-pg-curl_2.4.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-curl` | 2.4.5 | `u24.aarch64` | pigsty | 107.5 KiB | [postgresql-13-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-13-pg-curl_2.4.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-curl` | 2.4.5 | `u24.x86_64` | pigsty | 107.5 KiB | [postgresql-13-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-13-pg-curl_2.4.5-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/RekGRpth/pg_curl" title="Repository" icon="github" subtitle="github.com/RekGRpth/pg_curl" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_curl-2.4.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_curl; # get pg_curl source code
pig build dep pg_curl; # install build dependencies
pig build pkg pg_curl; # build extension rpm or deb
pig build ext pg_curl; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_curl; # install by extension name, for the current active PG version
pig ext install pg_curl; # install via package alias, for the active PG version
pig ext install pg_curl -v 18;   # install for PG 18
pig ext install pg_curl -v 17;   # install for PG 17
pig ext install pg_curl -v 16;   # install for PG 16
pig ext install pg_curl -v 15;   # install for PG 15
pig ext install pg_curl -v 14;   # install for PG 14
pig ext install pg_curl -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_curl;
```



--------

## Usage


```sql
CREATE EXTENSION pg_curl;
```

Perform HTTP Get:

```sql
-- wrap curl http get
CREATE OR REPLACE FUNCTION get(url TEXT) RETURNS TEXT LANGUAGE SQL AS $BODY$
WITH s AS (SELECT
               curl_easy_reset(),
               curl_easy_setopt_url(url),
               curl_easy_perform(),
               curl_easy_getinfo_data_in()
) SELECT convert_from(curl_easy_getinfo_data_in, 'utf-8') FROM s;
$BODY$;


SELECT get('https://www.postgresql.org/');
```


Perform Email SMTP:

```bash
CREATE OR REPLACE FUNCTION email(url TEXT, username TEXT, password TEXT, subject TEXT, sender TEXT, recipient TEXT, body TEXT, type TEXT) RETURNS TEXT LANGUAGE SQL AS $BODY$
    WITH s AS (SELECT
        curl_easy_reset(),
        curl_easy_setopt_mail_from(sender),
        curl_easy_setopt_password(password),
        curl_easy_setopt_url(url),
        curl_easy_setopt_username(username),
        curl_header_append('From', sender),
        curl_header_append('Subject', subject),
        curl_header_append('To', recipient),
        curl_mime_data(body, type:=type),
        curl_recipient_append(recipient),
        curl_easy_perform(),
        curl_easy_getinfo_header_in()
    ) SELECT curl_easy_getinfo_header_in FROM s;
$BODY$;
```

Perform FTP download:

```sql
CREATE OR REPLACE FUNCTION download(url TEXT, username TEXT, password TEXT) RETURNS BYTEA LANGUAGE SQL AS $BODY$
    WITH s AS (SELECT
        curl_easy_reset(),
        curl_easy_setopt_password(password),
        curl_easy_setopt_url(url),
        curl_easy_setopt_username(username),
        curl_easy_perform(),
        curl_easy_getinfo_data_in()
    ) SELECT curl_easy_getinfo_data_in FROM s;
$BODY$;
```
