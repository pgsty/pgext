---
title: "http"
linkTitle: "http"
description: "HTTP client for PostgreSQL, allows web page retrieval inside the database."
weight: 4070
categories: ["UTIL"]
width: full
---

HTTP client for PostgreSQL, allows web page retrieval inside the database.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4070** | {{< badge content="http" link="https://github.com/pramsey/pgsql-http" >}} | {{< ext "http" "pg_http" >}} | `1.7.0` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pgjwt" >}} {{< ext "pg_smtp_client" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "pgjq" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/http" >}} | `1.7.0` | {{< bg "18" "pg_http_18*" "green" >}} {{< bg "17" "pg_http_17*" "green" >}} {{< bg "16" "pg_http_16*" "green" >}} {{< bg "15" "pg_http_15*" "green" >}} {{< bg "14" "pg_http_14*" "green" >}} {{< bg "13" "pg_http_13*" "green" >}} | `pg_http_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/http" >}} | `1.7.0` | {{< bg "18" "postgresql-18-http" "green" >}} {{< bg "17" "postgresql-17-http" "green" >}} {{< bg "16" "postgresql-16-http" "green" >}} {{< bg "15" "postgresql-15-http" "green" >}} {{< bg "14" "postgresql-14-http" "green" >}} {{< bg "13" "postgresql-13-http" "green" >}} | `postgresql-$v-http` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.7.0" "pg_http_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.7.0" "pg_http_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.7.0" "pg_http_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.7.0" "pg_http_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.7.0" "pg_http_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.7.0" "pg_http_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_http_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 1.7.0" "postgresql-18-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-17-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-16-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-15-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-14-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-13-http : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 1.7.0" "postgresql-18-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-17-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-16-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-15-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-14-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-13-http : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 1.7.0" "postgresql-18-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-17-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-16-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-15-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-14-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-13-http : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 1.7.0" "postgresql-18-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-17-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-16-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-15-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-14-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-13-http : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 1.7.0" "postgresql-18-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-17-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-16-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-15-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-14-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-13-http : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 1.7.0" "postgresql-18-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-17-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-16-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-15-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-14-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-13-http : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 1.7.0" "postgresql-18-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-17-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-16-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-15-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-14-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-13-http : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 1.7.0" "postgresql-18-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-17-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-16-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-15-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-14-http : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.0" "postgresql-13-http : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_http_18` | 1.7.0 | `el8.x86_64` | pigsty | 29.0 KiB | [pg_http_18-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_http_18-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_http_18` | 1.7.0 | `el8.aarch64` | pigsty | 28.3 KiB | [pg_http_18-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_http_18-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_http_18` | 1.7.0 | `el9.x86_64` | pigsty | 29.1 KiB | [pg_http_18-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_http_18-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_http_18` | 1.7.0 | `el9.aarch64` | pigsty | 28.0 KiB | [pg_http_18-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_http_18-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_http_18` | 1.7.0 | `el10.x86_64` | pigsty | 29.5 KiB | [pg_http_18-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_http_18-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_http_18` | 1.7.0 | `el10.aarch64` | pigsty | 28.2 KiB | [pg_http_18-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_http_18-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-http` | 1.7.0 | `d12.x86_64` | pgdg | 44.5 KiB | [postgresql-18-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-18-http` | 1.7.0 | `d12.aarch64` | pgdg | 43.1 KiB | [postgresql-18-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-18-http` | 1.7.0 | `d13.x86_64` | pgdg | 44.6 KiB | [postgresql-18-http_1.7.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg13+1_amd64.deb) |
| `postgresql-18-http` | 1.7.0 | `d13.aarch64` | pgdg | 43.2 KiB | [postgresql-18-http_1.7.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg13+1_arm64.deb) |
| `postgresql-18-http` | 1.7.0 | `u22.x86_64` | pgdg | 44.7 KiB | [postgresql-18-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-http` | 1.7.0 | `u22.aarch64` | pgdg | 43.1 KiB | [postgresql-18-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-http` | 1.7.0 | `u24.x86_64` | pgdg | 44.5 KiB | [postgresql-18-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-http` | 1.7.0 | `u24.aarch64` | pgdg | 43.0 KiB | [postgresql-18-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_http_17` | 1.7.0 | `el8.x86_64` | pigsty | 29.1 KiB | [pg_http_17-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_http_17-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_http_17` | 1.7.0 | `el8.aarch64` | pigsty | 28.3 KiB | [pg_http_17-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_http_17-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_http_17` | 1.7.0 | `el9.x86_64` | pigsty | 29.2 KiB | [pg_http_17-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_http_17-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_http_17` | 1.7.0 | `el9.aarch64` | pigsty | 28.1 KiB | [pg_http_17-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_http_17-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_http_17` | 1.7.0 | `el10.x86_64` | pigsty | 29.5 KiB | [pg_http_17-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_http_17-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_http_17` | 1.7.0 | `el10.aarch64` | pigsty | 28.4 KiB | [pg_http_17-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_http_17-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-http` | 1.7.0 | `d12.x86_64` | pgdg | 44.6 KiB | [postgresql-17-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-17-http` | 1.7.0 | `d12.aarch64` | pgdg | 43.1 KiB | [postgresql-17-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-17-http` | 1.7.0 | `d13.x86_64` | pgdg | 44.6 KiB | [postgresql-17-http_1.7.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg13+1_amd64.deb) |
| `postgresql-17-http` | 1.7.0 | `d13.aarch64` | pgdg | 43.3 KiB | [postgresql-17-http_1.7.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg13+1_arm64.deb) |
| `postgresql-17-http` | 1.7.0 | `u22.x86_64` | pgdg | 48.9 KiB | [postgresql-17-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-http` | 1.7.0 | `u22.aarch64` | pgdg | 47.4 KiB | [postgresql-17-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-http` | 1.7.0 | `u24.x86_64` | pgdg | 44.6 KiB | [postgresql-17-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-http` | 1.7.0 | `u24.aarch64` | pgdg | 43.2 KiB | [postgresql-17-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_http_16` | 1.7.0 | `el8.x86_64` | pigsty | 29.1 KiB | [pg_http_16-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_http_16-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_http_16` | 1.7.0 | `el8.aarch64` | pigsty | 28.3 KiB | [pg_http_16-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_http_16-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_http_16` | 1.7.0 | `el9.x86_64` | pigsty | 29.2 KiB | [pg_http_16-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_http_16-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_http_16` | 1.7.0 | `el9.aarch64` | pigsty | 28.1 KiB | [pg_http_16-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_http_16-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_http_16` | 1.7.0 | `el10.x86_64` | pigsty | 29.5 KiB | [pg_http_16-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_http_16-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_http_16` | 1.7.0 | `el10.aarch64` | pigsty | 28.4 KiB | [pg_http_16-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_http_16-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-http` | 1.7.0 | `d12.x86_64` | pgdg | 44.5 KiB | [postgresql-16-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-16-http` | 1.7.0 | `d12.aarch64` | pgdg | 43.1 KiB | [postgresql-16-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-16-http` | 1.7.0 | `d13.x86_64` | pgdg | 44.6 KiB | [postgresql-16-http_1.7.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg13+1_amd64.deb) |
| `postgresql-16-http` | 1.7.0 | `d13.aarch64` | pgdg | 43.3 KiB | [postgresql-16-http_1.7.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg13+1_arm64.deb) |
| `postgresql-16-http` | 1.7.0 | `u22.x86_64` | pgdg | 48.9 KiB | [postgresql-16-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-http` | 1.7.0 | `u22.aarch64` | pgdg | 47.3 KiB | [postgresql-16-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-http` | 1.7.0 | `u24.x86_64` | pgdg | 44.6 KiB | [postgresql-16-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-http` | 1.7.0 | `u24.aarch64` | pgdg | 43.2 KiB | [postgresql-16-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_http_15` | 1.7.0 | `el8.x86_64` | pigsty | 29.4 KiB | [pg_http_15-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_http_15-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_http_15` | 1.7.0 | `el8.aarch64` | pigsty | 28.6 KiB | [pg_http_15-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_http_15-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_http_15` | 1.7.0 | `el9.x86_64` | pigsty | 29.5 KiB | [pg_http_15-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_http_15-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_http_15` | 1.7.0 | `el9.aarch64` | pigsty | 28.4 KiB | [pg_http_15-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_http_15-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_http_15` | 1.7.0 | `el10.x86_64` | pigsty | 29.9 KiB | [pg_http_15-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_http_15-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_http_15` | 1.7.0 | `el10.aarch64` | pigsty | 28.7 KiB | [pg_http_15-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_http_15-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-http` | 1.7.0 | `d12.x86_64` | pgdg | 45.4 KiB | [postgresql-15-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-15-http` | 1.7.0 | `d12.aarch64` | pgdg | 44.1 KiB | [postgresql-15-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-15-http` | 1.7.0 | `d13.x86_64` | pgdg | 45.7 KiB | [postgresql-15-http_1.7.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg13+1_amd64.deb) |
| `postgresql-15-http` | 1.7.0 | `d13.aarch64` | pgdg | 44.1 KiB | [postgresql-15-http_1.7.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg13+1_arm64.deb) |
| `postgresql-15-http` | 1.7.0 | `u22.x86_64` | pgdg | 50.0 KiB | [postgresql-15-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-http` | 1.7.0 | `u22.aarch64` | pgdg | 48.1 KiB | [postgresql-15-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-http` | 1.7.0 | `u24.x86_64` | pgdg | 45.5 KiB | [postgresql-15-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-http` | 1.7.0 | `u24.aarch64` | pgdg | 44.1 KiB | [postgresql-15-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_http_14` | 1.7.0 | `el8.x86_64` | pigsty | 29.4 KiB | [pg_http_14-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_http_14-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_http_14` | 1.7.0 | `el8.aarch64` | pigsty | 28.6 KiB | [pg_http_14-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_http_14-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_http_14` | 1.7.0 | `el9.x86_64` | pigsty | 29.5 KiB | [pg_http_14-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_http_14-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_http_14` | 1.7.0 | `el9.aarch64` | pigsty | 28.4 KiB | [pg_http_14-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_http_14-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_http_14` | 1.7.0 | `el10.x86_64` | pigsty | 29.9 KiB | [pg_http_14-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_http_14-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_http_14` | 1.7.0 | `el10.aarch64` | pigsty | 28.7 KiB | [pg_http_14-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_http_14-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-http` | 1.7.0 | `d12.x86_64` | pgdg | 45.3 KiB | [postgresql-14-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-14-http` | 1.7.0 | `d12.aarch64` | pgdg | 44.0 KiB | [postgresql-14-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-14-http` | 1.7.0 | `d13.x86_64` | pgdg | 45.6 KiB | [postgresql-14-http_1.7.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg13+1_amd64.deb) |
| `postgresql-14-http` | 1.7.0 | `d13.aarch64` | pgdg | 44.1 KiB | [postgresql-14-http_1.7.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg13+1_arm64.deb) |
| `postgresql-14-http` | 1.7.0 | `u22.x86_64` | pgdg | 50.0 KiB | [postgresql-14-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-http` | 1.7.0 | `u22.aarch64` | pgdg | 48.1 KiB | [postgresql-14-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-http` | 1.7.0 | `u24.x86_64` | pgdg | 45.6 KiB | [postgresql-14-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-http` | 1.7.0 | `u24.aarch64` | pgdg | 44.1 KiB | [postgresql-14-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_http_13` | 1.7.0 | `el8.x86_64` | pigsty | 28.9 KiB | [pg_http_13-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_http_13-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_http_13` | 1.7.0 | `el8.aarch64` | pigsty | 28.5 KiB | [pg_http_13-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_http_13-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_http_13` | 1.7.0 | `el9.x86_64` | pigsty | 29.4 KiB | [pg_http_13-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_http_13-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_http_13` | 1.7.0 | `el9.aarch64` | pigsty | 28.7 KiB | [pg_http_13-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_http_13-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_http_13` | 1.7.0 | `el10.x86_64` | pigsty | 29.7 KiB | [pg_http_13-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_http_13-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_http_13` | 1.7.0 | `el10.aarch64` | pigsty | 28.6 KiB | [pg_http_13-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_http_13-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-http` | 1.7.0 | `d12.x86_64` | pgdg | 45.3 KiB | [postgresql-13-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-13-http` | 1.7.0 | `d12.aarch64` | pgdg | 43.9 KiB | [postgresql-13-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-13-http` | 1.7.0 | `d13.x86_64` | pgdg | 45.5 KiB | [postgresql-13-http_1.7.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg13+1_amd64.deb) |
| `postgresql-13-http` | 1.7.0 | `d13.aarch64` | pgdg | 43.8 KiB | [postgresql-13-http_1.7.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg13+1_arm64.deb) |
| `postgresql-13-http` | 1.7.0 | `u22.x86_64` | pgdg | 49.6 KiB | [postgresql-13-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-http` | 1.7.0 | `u22.aarch64` | pgdg | 48.3 KiB | [postgresql-13-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-http` | 1.7.0 | `u24.x86_64` | pgdg | 45.4 KiB | [postgresql-13-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-http` | 1.7.0 | `u24.aarch64` | pgdg | 43.7 KiB | [postgresql-13-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pramsey/pgsql-http" title="Repository" icon="github" subtitle="github.com/pramsey/pgsql-http" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsql-http-1.7.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get http; # get http source code
pig build dep http; # install build dependencies
pig build pkg http; # build extension rpm or deb
pig build ext http; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install http; # install by extension name, for the current active PG version
pig ext install pg_http; # install via package alias, for the active PG version
pig ext install http -v 18;   # install for PG 18
pig ext install http -v 17;   # install for PG 17
pig ext install http -v 16;   # install for PG 16
pig ext install http -v 15;   # install for PG 15
pig ext install http -v 14;   # install for PG 14
pig ext install http -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION http;
```




--------

## Usage

https://github.com/pramsey/pgsql-http

Request / Response Schema:

```
     Composite type "public.http_request"
    Column    |       Type        | Modifiers
--------------+-------------------+-----------
 method       | http_method       |
 uri          | character varying |
 headers      | http_header[]     |
 content_type | character varying |
 content      | character varying |

    Composite type "public.http_response"
    Column    |       Type        | Modifiers
--------------+-------------------+-----------
 status       | integer           |
 content_type | character varying |
 headers      | http_header[]     |
 content      | character varying |
```


### Examples

Sending HTTP GET requests with SQL

```sql
CREATE EXTENSION http;

-- get content
SELECT content FROM http_get('http://httpbun.com/');

-- get status and content_type
SELECT status, content_type FROM http_get('http://httpbun.com/');

--  status |       content_type
-- --------+--------------------------
--     200 | text/html; charset=utf-8

-- get headers
SELECT (unnest(headers)).* FROM http_get('http://httpbun.com/');

--             field           |                      value
--  ---------------------------+--------------------------------------------------
--  Location                  | https://httpbun.com/
--  Date                      | Mon, 04 Nov 2024 09:00:36 GMT
--  Content-Length            | 0
--  Connection                | close
--  alt-svc                   | h3=":443"; ma=2592000
--  content-security-policy   | frame-ancestors 'none'
--  content-type              | text/html
--  date                      | Mon, 04 Nov 2024 09:00:37 GMT
--  strict-transport-security | max-age=31536000; includeSubDomains; preload
--  x-content-type-options    | nosniff
--  x-powered-by              | httpbun/af040d24038613575a85f74c2283ae79f8169927
--  (11 rows)
```


```sql
SELECT status, content::json->'form' AS form FROM http_post('http://httpbun.com/post', jsonb_build_object('myvar','myval','foo','bar'));
```

Issue http put requests:

```sql
SELECT status, content_type, content::json->>'data' AS data
  FROM http_put('http://httpbun.com/put', 'some text', 'text/plain');


--  status |   content_type   |   data
-- --------+------------------+-----------
--  200 | application/json | some text
```

Issue http post request:

