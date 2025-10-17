---
title: "http"
linkTitle: "http"
description: "HTTP client for PostgreSQL, allows web page retrieval inside the database."
weight: 4070
categories: ["Util"]
width: full
---

HTTP client for PostgreSQL, allows web page retrieval inside the database.

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4070** | {{< badge content="http" link="https://github.com/pramsey/pgsql-http" >}} | {{< ext "http" "pg_http" >}} | `1.7.0` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pgjwt" >}} {{< ext "pg_smtp_client" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "pgjq" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/http" >}} | `1.7.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgsql_http_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/http" >}} | `1.7.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-http` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pgsql_http_18" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsql_http_18-1.7.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsql_http_17" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_http_17-1.7.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsql_http_16" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_http_16-1.7.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsql_http_15" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_http_15-1.7.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsql_http_14" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_http_14-1.7.0-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pgsql_http_18" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsql_http_18-1.7.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsql_http_17" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_http_17-1.7.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsql_http_16" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_http_16-1.7.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsql_http_15" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_http_15-1.7.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsql_http_14" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_http_14-1.7.0-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pgsql_http_18" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsql_http_18-1.7.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsql_http_17" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_http_17-1.7.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsql_http_16" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_http_16-1.7.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsql_http_15" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_http_15-1.7.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsql_http_14" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_http_14-1.7.0-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pgsql_http_18" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsql_http_18-1.7.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsql_http_17" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_http_17-1.7.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsql_http_16" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.7.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsql_http_15" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.7.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsql_http_14" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.7.0-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-http" "1.7.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsql_http_18` | 1.7.0 | `el8.aarch64` | pgdg | 23.2 KiB | [pgsql_http_18-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsql_http_18-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_18` | 1.7.0 | `el8.x86_64` | pgdg | 24.1 KiB | [pgsql_http_18-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsql_http_18-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_18` | 1.7.0 | `el9.x86_64` | pgdg | 25.1 KiB | [pgsql_http_18-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsql_http_18-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_18` | 1.7.0 | `el9.aarch64` | pgdg | 23.6 KiB | [pgsql_http_18-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsql_http_18-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-http` | 1.7.0 | `d12.x86_64` | pgdg | 44.5 KiB | [postgresql-18-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-18-http` | 1.7.0 | `d12.aarch64` | pgdg | 43.1 KiB | [postgresql-18-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-18-http` | 1.7.0 | `u22.x86_64` | pgdg | 44.7 KiB | [postgresql-18-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-http` | 1.7.0 | `u22.aarch64` | pgdg | 43.1 KiB | [postgresql-18-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-http` | 1.7.0 | `u24.x86_64` | pgdg | 44.5 KiB | [postgresql-18-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-http` | 1.7.0 | `u24.aarch64` | pgdg | 43.0 KiB | [postgresql-18-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsql_http_17` | 1.7.0 | `el8.x86_64` | pgdg | 24.1 KiB | [pgsql_http_17-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_http_17-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_17` | 1.7.0 | `el8.aarch64` | pgdg | 23.2 KiB | [pgsql_http_17-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_http_17-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_17` | 1.6.3 | `el8.x86_64` | pgdg | 23.2 KiB | [pgsql_http_17-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_http_17-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_17` | 1.6.3 | `el8.aarch64` | pgdg | 22.3 KiB | [pgsql_http_17-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_http_17-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_17` | 1.6.2 | `el8.x86_64` | pgdg | 23.2 KiB | [pgsql_http_17-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_http_17-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_17` | 1.6.2 | `el8.aarch64` | pgdg | 22.5 KiB | [pgsql_http_17-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_http_17-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_17` | 1.6.0 | `el8.aarch64` | pgdg | 22.1 KiB | [pgsql_http_17-1.6.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_http_17-1.6.0-2PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_17` | 1.6.0 | `el8.x86_64` | pgdg | 22.8 KiB | [pgsql_http_17-1.6.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_http_17-1.6.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_17` | 1.7.0 | `el9.x86_64` | pgdg | 25.2 KiB | [pgsql_http_17-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_http_17-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_17` | 1.7.0 | `el9.aarch64` | pgdg | 23.7 KiB | [pgsql_http_17-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_http_17-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_17` | 1.6.3 | `el9.x86_64` | pgdg | 24.2 KiB | [pgsql_http_17-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_http_17-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_17` | 1.6.3 | `el9.aarch64` | pgdg | 22.9 KiB | [pgsql_http_17-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_http_17-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_17` | 1.6.2 | `el9.x86_64` | pgdg | 24.3 KiB | [pgsql_http_17-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_http_17-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_17` | 1.6.2 | `el9.aarch64` | pgdg | 23.0 KiB | [pgsql_http_17-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_http_17-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_17` | 1.6.0 | `el9.x86_64` | pgdg | 23.8 KiB | [pgsql_http_17-1.6.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_http_17-1.6.0-2PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_17` | 1.6.0 | `el9.aarch64` | pgdg | 22.4 KiB | [pgsql_http_17-1.6.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_http_17-1.6.0-2PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-http` | 1.7.0 | `d12.x86_64` | pgdg | 44.6 KiB | [postgresql-17-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-17-http` | 1.7.0 | `d12.aarch64` | pgdg | 43.1 KiB | [postgresql-17-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-17-http` | 1.7.0 | `u22.aarch64` | pgdg | 47.4 KiB | [postgresql-17-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-http` | 1.7.0 | `u22.x86_64` | pgdg | 48.9 KiB | [postgresql-17-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-http` | 1.7.0 | `u24.aarch64` | pgdg | 43.2 KiB | [postgresql-17-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-http` | 1.7.0 | `u24.x86_64` | pgdg | 44.6 KiB | [postgresql-17-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsql_http_16` | 1.7.0 | `el8.x86_64` | pgdg | 24.1 KiB | [pgsql_http_16-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_http_16-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_16` | 1.7.0 | `el8.aarch64` | pgdg | 23.2 KiB | [pgsql_http_16-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_http_16-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_16` | 1.6.3 | `el8.aarch64` | pgdg | 22.3 KiB | [pgsql_http_16-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_http_16-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_16` | 1.6.3 | `el8.x86_64` | pgdg | 23.2 KiB | [pgsql_http_16-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_http_16-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_16` | 1.6.2 | `el8.x86_64` | pgdg | 23.2 KiB | [pgsql_http_16-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_http_16-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_16` | 1.6.2 | `el8.aarch64` | pgdg | 22.4 KiB | [pgsql_http_16-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_http_16-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_16` | 1.6.0 | `el8.x86_64` | pgdg | 22.7 KiB | [pgsql_http_16-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_http_16-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_16` | 1.6.0 | `el8.aarch64` | pgdg | 21.9 KiB | [pgsql_http_16-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_http_16-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_16` | 1.6.0 | `el8.aarch64` | pgdg | 22.1 KiB | [pgsql_http_16-1.6.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_http_16-1.6.0-2PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_16` | 1.6.0 | `el8.x86_64` | pgdg | 22.8 KiB | [pgsql_http_16-1.6.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_http_16-1.6.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_16` | 1.7.0 | `el9.x86_64` | pgdg | 25.2 KiB | [pgsql_http_16-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_http_16-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_16` | 1.7.0 | `el9.aarch64` | pgdg | 23.6 KiB | [pgsql_http_16-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_16` | 1.6.3 | `el9.aarch64` | pgdg | 22.9 KiB | [pgsql_http_16-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_16` | 1.6.3 | `el9.x86_64` | pgdg | 24.2 KiB | [pgsql_http_16-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_http_16-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_16` | 1.6.2 | `el9.x86_64` | pgdg | 24.3 KiB | [pgsql_http_16-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_http_16-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_16` | 1.6.2 | `el9.aarch64` | pgdg | 23.0 KiB | [pgsql_http_16-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_16` | 1.6.0 | `el9.aarch64` | pgdg | 22.2 KiB | [pgsql_http_16-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_16` | 1.6.0 | `el9.aarch64` | pgdg | 22.4 KiB | [pgsql_http_16-1.6.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.6.0-2PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_16` | 1.6.0 | `el9.x86_64` | pgdg | 23.5 KiB | [pgsql_http_16-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_http_16-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_16` | 1.6.0 | `el9.x86_64` | pgdg | 23.7 KiB | [pgsql_http_16-1.6.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_http_16-1.6.0-2PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-http` | 1.7.0 | `d12.x86_64` | pgdg | 44.5 KiB | [postgresql-16-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-16-http` | 1.7.0 | `d12.aarch64` | pgdg | 43.1 KiB | [postgresql-16-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-16-http` | 1.7.0 | `u22.x86_64` | pgdg | 48.9 KiB | [postgresql-16-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-http` | 1.7.0 | `u22.aarch64` | pgdg | 47.3 KiB | [postgresql-16-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-http` | 1.7.0 | `u24.aarch64` | pgdg | 43.2 KiB | [postgresql-16-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg24.04+1_arm64.deb) |
| `postgresql-16-http` | 1.7.0 | `u24.x86_64` | pgdg | 44.6 KiB | [postgresql-16-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsql_http_15` | 1.7.0 | `el8.x86_64` | pgdg | 24.4 KiB | [pgsql_http_15-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_http_15-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_15` | 1.7.0 | `el8.aarch64` | pgdg | 23.5 KiB | [pgsql_http_15-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_http_15-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_15` | 1.6.3 | `el8.x86_64` | pgdg | 23.3 KiB | [pgsql_http_15-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_http_15-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_15` | 1.6.3 | `el8.aarch64` | pgdg | 22.4 KiB | [pgsql_http_15-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_http_15-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_15` | 1.6.2 | `el8.aarch64` | pgdg | 22.7 KiB | [pgsql_http_15-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_http_15-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_15` | 1.6.2 | `el8.x86_64` | pgdg | 23.5 KiB | [pgsql_http_15-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_http_15-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_15` | 1.6.0 | `el8.aarch64` | pgdg | 22.2 KiB | [pgsql_http_15-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_http_15-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_15` | 1.6.0 | `el8.aarch64` | pgdg | 22.3 KiB | [pgsql_http_15-1.6.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_http_15-1.6.0-2PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_15` | 1.6.0 | `el8.x86_64` | pgdg | 23.1 KiB | [pgsql_http_15-1.6.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_http_15-1.6.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_15` | 1.6.0 | `el8.x86_64` | pgdg | 22.9 KiB | [pgsql_http_15-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_http_15-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_15` | 1.7.0 | `el9.aarch64` | pgdg | 23.8 KiB | [pgsql_http_15-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_15` | 1.7.0 | `el9.x86_64` | pgdg | 25.5 KiB | [pgsql_http_15-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_http_15-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_15` | 1.6.3 | `el9.x86_64` | pgdg | 24.4 KiB | [pgsql_http_15-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_http_15-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_15` | 1.6.3 | `el9.aarch64` | pgdg | 23.1 KiB | [pgsql_http_15-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_15` | 1.6.2 | `el9.x86_64` | pgdg | 24.7 KiB | [pgsql_http_15-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_http_15-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_15` | 1.6.2 | `el9.aarch64` | pgdg | 23.7 KiB | [pgsql_http_15-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_15` | 1.6.0 | `el9.x86_64` | pgdg | 24.1 KiB | [pgsql_http_15-1.6.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_http_15-1.6.0-2PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_15` | 1.6.0 | `el9.aarch64` | pgdg | 22.8 KiB | [pgsql_http_15-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_15` | 1.6.0 | `el9.aarch64` | pgdg | 23.1 KiB | [pgsql_http_15-1.6.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.6.0-2PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_15` | 1.6.0 | `el9.x86_64` | pgdg | 23.9 KiB | [pgsql_http_15-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_http_15-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-15-http` | 1.7.0 | `d12.x86_64` | pgdg | 45.4 KiB | [postgresql-15-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-15-http` | 1.7.0 | `d12.aarch64` | pgdg | 44.1 KiB | [postgresql-15-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-15-http` | 1.7.0 | `u22.x86_64` | pgdg | 50.0 KiB | [postgresql-15-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-http` | 1.7.0 | `u22.aarch64` | pgdg | 48.1 KiB | [postgresql-15-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-http` | 1.7.0 | `u24.aarch64` | pgdg | 44.1 KiB | [postgresql-15-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg24.04+1_arm64.deb) |
| `postgresql-15-http` | 1.7.0 | `u24.x86_64` | pgdg | 45.5 KiB | [postgresql-15-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsql_http_14` | 1.7.0 | `el8.aarch64` | pgdg | 23.5 KiB | [pgsql_http_14-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_http_14-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_14` | 1.7.0 | `el8.x86_64` | pgdg | 24.4 KiB | [pgsql_http_14-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_http_14-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_14` | 1.6.3 | `el8.x86_64` | pgdg | 23.3 KiB | [pgsql_http_14-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_http_14-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_14` | 1.6.3 | `el8.aarch64` | pgdg | 22.4 KiB | [pgsql_http_14-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_http_14-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_14` | 1.6.2 | `el8.x86_64` | pgdg | 23.5 KiB | [pgsql_http_14-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_http_14-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_14` | 1.6.2 | `el8.aarch64` | pgdg | 22.7 KiB | [pgsql_http_14-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_http_14-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_14` | 1.6.0 | `el8.x86_64` | pgdg | 22.9 KiB | [pgsql_http_14-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_http_14-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_14` | 1.6.0 | `el8.x86_64` | pgdg | 23.0 KiB | [pgsql_http_14-1.6.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_http_14-1.6.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_14` | 1.6.0 | `el8.aarch64` | pgdg | 22.2 KiB | [pgsql_http_14-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_http_14-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_14` | 1.6.0 | `el8.aarch64` | pgdg | 22.3 KiB | [pgsql_http_14-1.6.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_http_14-1.6.0-2PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_14` | 1.7.0 | `el9.aarch64` | pgdg | 23.8 KiB | [pgsql_http_14-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_14` | 1.7.0 | `el9.x86_64` | pgdg | 25.5 KiB | [pgsql_http_14-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_http_14-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_14` | 1.6.3 | `el9.aarch64` | pgdg | 23.1 KiB | [pgsql_http_14-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_14` | 1.6.3 | `el9.x86_64` | pgdg | 24.4 KiB | [pgsql_http_14-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_http_14-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_14` | 1.6.2 | `el9.aarch64` | pgdg | 23.7 KiB | [pgsql_http_14-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_14` | 1.6.2 | `el9.x86_64` | pgdg | 24.7 KiB | [pgsql_http_14-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_http_14-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_14` | 1.6.0 | `el9.aarch64` | pgdg | 22.8 KiB | [pgsql_http_14-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_14` | 1.6.0 | `el9.x86_64` | pgdg | 23.9 KiB | [pgsql_http_14-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_http_14-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_14` | 1.6.0 | `el9.x86_64` | pgdg | 24.1 KiB | [pgsql_http_14-1.6.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_http_14-1.6.0-2PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_14` | 1.6.0 | `el9.aarch64` | pgdg | 23.1 KiB | [pgsql_http_14-1.6.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.6.0-2PGDG.rhel9.aarch64.rpm) |
| `postgresql-14-http` | 1.7.0 | `d12.x86_64` | pgdg | 45.3 KiB | [postgresql-14-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-14-http` | 1.7.0 | `d12.aarch64` | pgdg | 44.0 KiB | [postgresql-14-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-14-http` | 1.7.0 | `u22.aarch64` | pgdg | 48.1 KiB | [postgresql-14-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-http` | 1.7.0 | `u22.x86_64` | pgdg | 50.0 KiB | [postgresql-14-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-http` | 1.7.0 | `u24.x86_64` | pgdg | 45.6 KiB | [postgresql-14-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-http` | 1.7.0 | `u24.aarch64` | pgdg | 44.1 KiB | [postgresql-14-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsql_http_13` | 1.7.0 | `el8.aarch64` | pgdg | 23.4 KiB | [pgsql_http_13-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_http_13-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_13` | 1.7.0 | `el8.x86_64` | pgdg | 24.0 KiB | [pgsql_http_13-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_http_13-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_13` | 1.6.3 | `el8.x86_64` | pgdg | 23.0 KiB | [pgsql_http_13-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_http_13-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_13` | 1.6.3 | `el8.aarch64` | pgdg | 22.4 KiB | [pgsql_http_13-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_http_13-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_13` | 1.6.2 | `el8.x86_64` | pgdg | 23.1 KiB | [pgsql_http_13-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_http_13-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_13` | 1.6.2 | `el8.aarch64` | pgdg | 22.6 KiB | [pgsql_http_13-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_http_13-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_13` | 1.6.0 | `el8.aarch64` | pgdg | 22.1 KiB | [pgsql_http_13-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_http_13-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_13` | 1.6.0 | `el8.x86_64` | pgdg | 22.5 KiB | [pgsql_http_13-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_http_13-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_13` | 1.6.0 | `el8.aarch64` | pgdg | 22.2 KiB | [pgsql_http_13-1.6.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsql_http_13-1.6.0-2PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_13` | 1.6.0 | `el8.x86_64` | pgdg | 22.6 KiB | [pgsql_http_13-1.6.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsql_http_13-1.6.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_13` | 1.7.0 | `el9.aarch64` | pgdg | 24.3 KiB | [pgsql_http_13-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_http_13-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_13` | 1.7.0 | `el9.x86_64` | pgdg | 25.4 KiB | [pgsql_http_13-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_http_13-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_13` | 1.6.3 | `el9.aarch64` | pgdg | 23.3 KiB | [pgsql_http_13-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_http_13-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_13` | 1.6.3 | `el9.x86_64` | pgdg | 24.3 KiB | [pgsql_http_13-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_http_13-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_13` | 1.6.2 | `el9.aarch64` | pgdg | 23.6 KiB | [pgsql_http_13-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_http_13-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_13` | 1.6.2 | `el9.x86_64` | pgdg | 24.5 KiB | [pgsql_http_13-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_http_13-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_13` | 1.6.0 | `el9.x86_64` | pgdg | 23.7 KiB | [pgsql_http_13-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_http_13-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_13` | 1.6.0 | `el9.aarch64` | pgdg | 23.1 KiB | [pgsql_http_13-1.6.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_http_13-1.6.0-2PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_13` | 1.6.0 | `el9.aarch64` | pgdg | 22.7 KiB | [pgsql_http_13-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsql_http_13-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_13` | 1.6.0 | `el9.x86_64` | pgdg | 23.9 KiB | [pgsql_http_13-1.6.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsql_http_13-1.6.0-2PGDG.rhel9.x86_64.rpm) |
| `postgresql-13-http` | 1.7.0 | `d12.aarch64` | pgdg | 43.9 KiB | [postgresql-13-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-13-http` | 1.7.0 | `d12.x86_64` | pgdg | 45.3 KiB | [postgresql-13-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-13-http` | 1.7.0 | `u22.aarch64` | pgdg | 48.3 KiB | [postgresql-13-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-http` | 1.7.0 | `u22.x86_64` | pgdg | 49.6 KiB | [postgresql-13-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-http` | 1.7.0 | `u24.aarch64` | pgdg | 43.7 KiB | [postgresql-13-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg24.04+1_arm64.deb) |
| `postgresql-13-http` | 1.7.0 | `u24.x86_64` | pgdg | 45.4 KiB | [postgresql-13-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-13-http_1.7.0-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pramsey/pgsql-http" title="Repository" icon="github" subtitle="github.com/pramsey/pgsql-http" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgsql-http-1.6.3.tar.gz" >}}
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

