---
title: "pg_similarity"
linkTitle: "pg_similarity"
description: "support similarity queries"
weight: 1840
categories: ["Rag"]
width: full
---

support similarity queries

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1840** | {{< badge content="pg_similarity" link="https://github.com/eulerto/pg_similarity" >}} | {{< ext "pg_similarity" "pg_similarity" >}} | `1.0` | {{< category "RAG" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "vector" >}} {{< ext "smlar" >}} {{< ext "fuzzystrmatch" >}} {{< ext "pg_trgm" >}} {{< ext "vchord" >}} {{< ext "pg_bigm" >}} {{< ext "citext" >}} {{< ext "unaccent" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_similarity" >}} | `1.0` | {{< badge content="18" color="red" alt="pg_similarity_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_similarity_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_similarity" >}} | `1.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-similarity` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_similarity_18" "1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_similarity_18-1.0-3PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_similarity_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_17-1.0-2PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_similarity_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_16-1.0-2PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_similarity_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_15-1.0-2PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_similarity_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_14-1.0-2PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_similarity_18" "1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_similarity_18-1.0-3PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_similarity_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_17-1.0-2PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_similarity_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_16-1.0-2PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_similarity_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_15-1.0-2PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_similarity_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_14-1.0-2PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_similarity_18" "1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_similarity_18-1.0-3PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_similarity_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_17-1.0-2PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_similarity_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_16-1.0-2PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_similarity_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_15-1.0-2PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_similarity_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_14-1.0-2PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_similarity_18" "1.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_similarity_18-1.0-3PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_similarity_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_17-1.0-2PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_similarity_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_16-1.0-2PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_similarity_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_15-1.0-2PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_similarity_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_14-1.0-2PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-similarity" "1.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_similarity_18` | 1.0 | `el8.x86_64` | pgdg | 43.2 KiB | [pg_similarity_18-1.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_similarity_18-1.0-3PGDG.rhel8.x86_64.rpm) |
| `pg_similarity_18` | 1.0 | `el8.aarch64` | pgdg | 40.7 KiB | [pg_similarity_18-1.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_similarity_18-1.0-3PGDG.rhel8.aarch64.rpm) |
| `pg_similarity_18` | 1.0 | `el8.x86_64` | pigsty | 41.6 KiB | [pg_similarity_18-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_18-1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_similarity_18` | 1.0 | `el9.x86_64` | pgdg | 42.6 KiB | [pg_similarity_18-1.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_similarity_18-1.0-3PGDG.rhel9.x86_64.rpm) |
| `pg_similarity_18` | 1.0 | `el9.x86_64` | pigsty | 41.2 KiB | [pg_similarity_18-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_18-1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_similarity_18` | 1.0 | `el9.aarch64` | pgdg | 41.2 KiB | [pg_similarity_18-1.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_similarity_18-1.0-3PGDG.rhel9.aarch64.rpm) |
| `pg_similarity_18` | 1.0 | `el9.aarch64` | pigsty | 39.6 KiB | [pg_similarity_18-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_18-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-18-similarity` | 1.0 | `d12.x86_64` | pgdg | 98.6 KiB | [postgresql-18-similarity_1.0-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg12+1_amd64.deb) |
| `postgresql-18-similarity` | 1.0 | `d12.aarch64` | pgdg | 96.2 KiB | [postgresql-18-similarity_1.0-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg12+1_arm64.deb) |
| `postgresql-18-similarity` | 1.0 | `u22.aarch64` | pgdg | 96.2 KiB | [postgresql-18-similarity_1.0-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg22.04+1_arm64.deb) |
| `postgresql-18-similarity` | 1.0 | `u22.x86_64` | pgdg | 98.4 KiB | [postgresql-18-similarity_1.0-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg22.04+1_amd64.deb) |
| `postgresql-18-similarity` | 1.0 | `u24.x86_64` | pgdg | 97.5 KiB | [postgresql-18-similarity_1.0-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg24.04+1_amd64.deb) |
| `postgresql-18-similarity` | 1.0 | `u24.aarch64` | pgdg | 94.9 KiB | [postgresql-18-similarity_1.0-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_similarity_17` | 1.0 | `el8.x86_64` | pigsty | 41.6 KiB | [pg_similarity_17-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_17-1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_similarity_17` | 1.0 | `el8.aarch64` | pigsty | 39.2 KiB | [pg_similarity_17-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_17-1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_similarity_17` | 1.0 | `el9.aarch64` | pigsty | 39.7 KiB | [pg_similarity_17-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_17-1.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_similarity_17` | 1.0 | `el9.x86_64` | pigsty | 41.2 KiB | [pg_similarity_17-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_17-1.0-2PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-similarity` | 1.0 | `d12.x86_64` | pgdg | 98.7 KiB | [postgresql-17-similarity_1.0-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg12+1_amd64.deb) |
| `postgresql-17-similarity` | 1.0 | `d12.aarch64` | pgdg | 96.1 KiB | [postgresql-17-similarity_1.0-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg12+1_arm64.deb) |
| `postgresql-17-similarity` | 1.0 | `u22.aarch64` | pgdg | 101.5 KiB | [postgresql-17-similarity_1.0-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg22.04+1_arm64.deb) |
| `postgresql-17-similarity` | 1.0 | `u22.x86_64` | pgdg | 103.7 KiB | [postgresql-17-similarity_1.0-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg22.04+1_amd64.deb) |
| `postgresql-17-similarity` | 1.0 | `u24.aarch64` | pgdg | 95.0 KiB | [postgresql-17-similarity_1.0-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg24.04+1_arm64.deb) |
| `postgresql-17-similarity` | 1.0 | `u24.x86_64` | pgdg | 97.6 KiB | [postgresql-17-similarity_1.0-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_similarity_16` | 1.0 | `el8.x86_64` | pigsty | 41.6 KiB | [pg_similarity_16-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_16-1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_similarity_16` | 1.0 | `el8.aarch64` | pigsty | 39.2 KiB | [pg_similarity_16-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_16-1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_similarity_16` | 1.0 | `el9.x86_64` | pigsty | 41.2 KiB | [pg_similarity_16-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_16-1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_similarity_16` | 1.0 | `el9.aarch64` | pigsty | 39.7 KiB | [pg_similarity_16-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_16-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-similarity` | 1.0 | `d12.x86_64` | pgdg | 98.7 KiB | [postgresql-16-similarity_1.0-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg12+1_amd64.deb) |
| `postgresql-16-similarity` | 1.0 | `d12.aarch64` | pgdg | 96.1 KiB | [postgresql-16-similarity_1.0-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg12+1_arm64.deb) |
| `postgresql-16-similarity` | 1.0 | `u22.x86_64` | pgdg | 103.6 KiB | [postgresql-16-similarity_1.0-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg22.04+1_amd64.deb) |
| `postgresql-16-similarity` | 1.0 | `u22.aarch64` | pgdg | 101.5 KiB | [postgresql-16-similarity_1.0-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg22.04+1_arm64.deb) |
| `postgresql-16-similarity` | 1.0 | `u24.aarch64` | pgdg | 95.1 KiB | [postgresql-16-similarity_1.0-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg24.04+1_arm64.deb) |
| `postgresql-16-similarity` | 1.0 | `u24.x86_64` | pgdg | 97.5 KiB | [postgresql-16-similarity_1.0-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_similarity_15` | 1.0 | `el8.x86_64` | pigsty | 42.7 KiB | [pg_similarity_15-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_15-1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_similarity_15` | 1.0 | `el8.aarch64` | pigsty | 40.1 KiB | [pg_similarity_15-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_15-1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_similarity_15` | 1.0 | `el9.x86_64` | pigsty | 43.4 KiB | [pg_similarity_15-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_15-1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_similarity_15` | 1.0 | `el9.aarch64` | pigsty | 41.7 KiB | [pg_similarity_15-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_15-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-similarity` | 1.0 | `d12.aarch64` | pgdg | 96.9 KiB | [postgresql-15-similarity_1.0-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg12+1_arm64.deb) |
| `postgresql-15-similarity` | 1.0 | `d12.x86_64` | pgdg | 99.5 KiB | [postgresql-15-similarity_1.0-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg12+1_amd64.deb) |
| `postgresql-15-similarity` | 1.0 | `u22.x86_64` | pgdg | 105.4 KiB | [postgresql-15-similarity_1.0-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg22.04+1_amd64.deb) |
| `postgresql-15-similarity` | 1.0 | `u22.aarch64` | pgdg | 103.0 KiB | [postgresql-15-similarity_1.0-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg22.04+1_arm64.deb) |
| `postgresql-15-similarity` | 1.0 | `u24.aarch64` | pgdg | 96.4 KiB | [postgresql-15-similarity_1.0-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg24.04+1_arm64.deb) |
| `postgresql-15-similarity` | 1.0 | `u24.x86_64` | pgdg | 99.3 KiB | [postgresql-15-similarity_1.0-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_similarity_14` | 1.0 | `el8.x86_64` | pigsty | 42.7 KiB | [pg_similarity_14-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_14-1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_similarity_14` | 1.0 | `el8.aarch64` | pigsty | 40.1 KiB | [pg_similarity_14-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_14-1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_similarity_14` | 1.0 | `el9.x86_64` | pigsty | 43.3 KiB | [pg_similarity_14-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_14-1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_similarity_14` | 1.0 | `el9.aarch64` | pigsty | 41.6 KiB | [pg_similarity_14-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_14-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-similarity` | 1.0 | `d12.aarch64` | pgdg | 96.8 KiB | [postgresql-14-similarity_1.0-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg12+1_arm64.deb) |
| `postgresql-14-similarity` | 1.0 | `d12.x86_64` | pgdg | 99.5 KiB | [postgresql-14-similarity_1.0-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg12+1_amd64.deb) |
| `postgresql-14-similarity` | 1.0 | `u22.aarch64` | pgdg | 102.8 KiB | [postgresql-14-similarity_1.0-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg22.04+1_arm64.deb) |
| `postgresql-14-similarity` | 1.0 | `u22.x86_64` | pgdg | 105.4 KiB | [postgresql-14-similarity_1.0-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg22.04+1_amd64.deb) |
| `postgresql-14-similarity` | 1.0 | `u24.x86_64` | pgdg | 99.3 KiB | [postgresql-14-similarity_1.0-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg24.04+1_amd64.deb) |
| `postgresql-14-similarity` | 1.0 | `u24.aarch64` | pgdg | 96.4 KiB | [postgresql-14-similarity_1.0-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_similarity_13` | 1.0 | `el8.aarch64` | pigsty | 40.0 KiB | [pg_similarity_13-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_13-1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_similarity_13` | 1.0 | `el8.x86_64` | pigsty | 42.3 KiB | [pg_similarity_13-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_13-1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_similarity_13` | 1.0 | `el9.x86_64` | pigsty | 43.5 KiB | [pg_similarity_13-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_13-1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_similarity_13` | 1.0 | `el9.aarch64` | pigsty | 41.6 KiB | [pg_similarity_13-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_13-1.0-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-similarity` | 1.0 | `d12.aarch64` | pgdg | 96.3 KiB | [postgresql-13-similarity_1.0-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg12+1_arm64.deb) |
| `postgresql-13-similarity` | 1.0 | `d12.x86_64` | pgdg | 99.0 KiB | [postgresql-13-similarity_1.0-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg12+1_amd64.deb) |
| `postgresql-13-similarity` | 1.0 | `u22.aarch64` | pgdg | 102.2 KiB | [postgresql-13-similarity_1.0-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg22.04+1_arm64.deb) |
| `postgresql-13-similarity` | 1.0 | `u22.x86_64` | pgdg | 104.9 KiB | [postgresql-13-similarity_1.0-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg22.04+1_amd64.deb) |
| `postgresql-13-similarity` | 1.0 | `u24.x86_64` | pgdg | 98.6 KiB | [postgresql-13-similarity_1.0-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg24.04+1_amd64.deb) |
| `postgresql-13-similarity` | 1.0 | `u24.aarch64` | pgdg | 96.0 KiB | [postgresql-13-similarity_1.0-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/eulerto/pg_similarity" title="Repository" icon="github" subtitle="github.com/eulerto/pg_similarity" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_similarity-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_similarity; # get pg_similarity source code
pig build dep pg_similarity; # install build dependencies
pig build pkg pg_similarity; # build extension rpm or deb
pig build ext pg_similarity; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_similarity; # install by extension name, for the current active PG version
pig ext install pg_similarity; # install via package alias, for the active PG version
pig ext install pg_similarity -v 18;   # install for PG 18
pig ext install pg_similarity -v 17;   # install for PG 17
pig ext install pg_similarity -v 16;   # install for PG 16
pig ext install pg_similarity -v 15;   # install for PG 15
pig ext install pg_similarity -v 14;   # install for PG 14
pig ext install pg_similarity -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_similarity;
```

