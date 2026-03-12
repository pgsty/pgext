---
title: "pg_similarity"
linkTitle: "pg_similarity"
description: "support similarity queries"
weight: 1840
categories: ["RAG"]
width: full
---

[**pg_similarity**](https://github.com/eulerto/pg_similarity) : support similarity queries


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1840** | {{< badge content="pg_similarity" link="https://github.com/eulerto/pg_similarity" >}} | {{< ext "pg_similarity" >}} | `1.0` | {{< category "RAG" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "vector" >}} {{< ext "smlar" >}} {{< ext "fuzzystrmatch" >}} {{< ext "pg_trgm" >}} {{< ext "vchord" >}} {{< ext "pg_bigm" >}} {{< ext "citext" >}} {{< ext "unaccent" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_similarity` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_similarity_18" "green" >}} {{< bg "17" "pg_similarity_17" "green" >}} {{< bg "16" "pg_similarity_16" "green" >}} {{< bg "15" "pg_similarity_15" "green" >}} {{< bg "14" "pg_similarity_14" "green" >}} | `pg_similarity_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "postgresql-18-similarity" "green" >}} {{< bg "17" "postgresql-17-similarity" "green" >}} {{< bg "16" "postgresql-16-similarity" "green" >}} {{< bg "15" "postgresql-15-similarity" "green" >}} {{< bg "14" "postgresql-14-similarity" "green" >}} | `postgresql-$v-similarity` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0" "pg_similarity_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0" "pg_similarity_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0" "pg_similarity_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0" "pg_similarity_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0" "pg_similarity_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0" "pg_similarity_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_similarity_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.2 KiB | [pg_similarity_18-1.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_similarity_18-1.0-3PGDG.rhel8.x86_64.rpm) |
| `pg_similarity_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.9 KiB | [pg_similarity_18-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_18-1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_similarity_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [pg_similarity_18-1.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_similarity_18-1.0-3PGDG.rhel8.aarch64.rpm) |
| `pg_similarity_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.0 KiB | [pg_similarity_18-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_18-1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_similarity_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.6 KiB | [pg_similarity_18-1.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_similarity_18-1.0-3PGDG.rhel9.x86_64.rpm) |
| `pg_similarity_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.9 KiB | [pg_similarity_18-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_18-1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_similarity_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.2 KiB | [pg_similarity_18-1.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_similarity_18-1.0-3PGDG.rhel9.aarch64.rpm) |
| `pg_similarity_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 40.6 KiB | [pg_similarity_18-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_18-1.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_similarity_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.6 KiB | [pg_similarity_18-1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_similarity_18-1.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_similarity_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.3 KiB | [pg_similarity_18-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_similarity_18-1.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_similarity_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.3 KiB | [pg_similarity_18-1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_similarity_18-1.0-3PGDG.rhel10.aarch64.rpm) |
| `pg_similarity_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.0 KiB | [pg_similarity_18-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_similarity_18-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-similarity` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 98.6 KiB | [postgresql-18-similarity_1.0-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg12+1_amd64.deb) |
| `postgresql-18-similarity` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 96.2 KiB | [postgresql-18-similarity_1.0-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg12+1_arm64.deb) |
| `postgresql-18-similarity` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 98.5 KiB | [postgresql-18-similarity_1.0-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg13+1_amd64.deb) |
| `postgresql-18-similarity` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 96.2 KiB | [postgresql-18-similarity_1.0-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg13+1_arm64.deb) |
| `postgresql-18-similarity` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 98.4 KiB | [postgresql-18-similarity_1.0-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg22.04+1_amd64.deb) |
| `postgresql-18-similarity` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 96.2 KiB | [postgresql-18-similarity_1.0-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg22.04+1_arm64.deb) |
| `postgresql-18-similarity` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 97.5 KiB | [postgresql-18-similarity_1.0-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg24.04+1_amd64.deb) |
| `postgresql-18-similarity` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 94.9 KiB | [postgresql-18-similarity_1.0-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-18-similarity_1.0-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_similarity_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.8 KiB | [pg_similarity_17-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_17-1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_similarity_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.9 KiB | [pg_similarity_17-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_17-1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_similarity_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.0 KiB | [pg_similarity_17-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_17-1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_similarity_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 40.6 KiB | [pg_similarity_17-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_17-1.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_similarity_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.6 KiB | [pg_similarity_17-1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_similarity_17-1.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_similarity_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.3 KiB | [pg_similarity_17-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_similarity_17-1.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_similarity_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.3 KiB | [pg_similarity_17-1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_similarity_17-1.0-3PGDG.rhel10.aarch64.rpm) |
| `pg_similarity_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.0 KiB | [pg_similarity_17-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_similarity_17-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-similarity` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 98.7 KiB | [postgresql-17-similarity_1.0-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg12+1_amd64.deb) |
| `postgresql-17-similarity` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 96.1 KiB | [postgresql-17-similarity_1.0-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg12+1_arm64.deb) |
| `postgresql-17-similarity` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 98.8 KiB | [postgresql-17-similarity_1.0-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg13+1_amd64.deb) |
| `postgresql-17-similarity` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 96.0 KiB | [postgresql-17-similarity_1.0-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg13+1_arm64.deb) |
| `postgresql-17-similarity` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 103.7 KiB | [postgresql-17-similarity_1.0-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg22.04+1_amd64.deb) |
| `postgresql-17-similarity` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 101.5 KiB | [postgresql-17-similarity_1.0-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg22.04+1_arm64.deb) |
| `postgresql-17-similarity` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 97.6 KiB | [postgresql-17-similarity_1.0-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg24.04+1_amd64.deb) |
| `postgresql-17-similarity` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 95.0 KiB | [postgresql-17-similarity_1.0-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-17-similarity_1.0-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_similarity_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.8 KiB | [pg_similarity_16-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_16-1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_similarity_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.0 KiB | [pg_similarity_16-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_16-1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_similarity_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 41.9 KiB | [pg_similarity_16-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_16-1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_similarity_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 40.5 KiB | [pg_similarity_16-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_16-1.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_similarity_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.6 KiB | [pg_similarity_16-1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_similarity_16-1.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_similarity_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 42.3 KiB | [pg_similarity_16-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_similarity_16-1.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_similarity_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.3 KiB | [pg_similarity_16-1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_similarity_16-1.0-3PGDG.rhel10.aarch64.rpm) |
| `pg_similarity_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.0 KiB | [pg_similarity_16-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_similarity_16-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-similarity` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 98.7 KiB | [postgresql-16-similarity_1.0-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg12+1_amd64.deb) |
| `postgresql-16-similarity` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 96.1 KiB | [postgresql-16-similarity_1.0-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg12+1_arm64.deb) |
| `postgresql-16-similarity` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 98.6 KiB | [postgresql-16-similarity_1.0-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg13+1_amd64.deb) |
| `postgresql-16-similarity` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 96.2 KiB | [postgresql-16-similarity_1.0-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg13+1_arm64.deb) |
| `postgresql-16-similarity` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 103.6 KiB | [postgresql-16-similarity_1.0-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg22.04+1_amd64.deb) |
| `postgresql-16-similarity` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 101.5 KiB | [postgresql-16-similarity_1.0-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg22.04+1_arm64.deb) |
| `postgresql-16-similarity` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 97.5 KiB | [postgresql-16-similarity_1.0-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg24.04+1_amd64.deb) |
| `postgresql-16-similarity` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 95.1 KiB | [postgresql-16-similarity_1.0-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-16-similarity_1.0-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_similarity_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 45.9 KiB | [pg_similarity_15-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_15-1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_similarity_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.9 KiB | [pg_similarity_15-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_15-1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_similarity_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 44.1 KiB | [pg_similarity_15-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_15-1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_similarity_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 42.5 KiB | [pg_similarity_15-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_15-1.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_similarity_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.8 KiB | [pg_similarity_15-1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_similarity_15-1.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_similarity_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 44.5 KiB | [pg_similarity_15-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_similarity_15-1.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_similarity_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.2 KiB | [pg_similarity_15-1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_similarity_15-1.0-3PGDG.rhel10.aarch64.rpm) |
| `pg_similarity_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.0 KiB | [pg_similarity_15-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_similarity_15-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-similarity` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 99.5 KiB | [postgresql-15-similarity_1.0-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg12+1_amd64.deb) |
| `postgresql-15-similarity` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 96.9 KiB | [postgresql-15-similarity_1.0-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg12+1_arm64.deb) |
| `postgresql-15-similarity` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 99.7 KiB | [postgresql-15-similarity_1.0-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg13+1_amd64.deb) |
| `postgresql-15-similarity` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 96.9 KiB | [postgresql-15-similarity_1.0-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg13+1_arm64.deb) |
| `postgresql-15-similarity` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 105.4 KiB | [postgresql-15-similarity_1.0-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg22.04+1_amd64.deb) |
| `postgresql-15-similarity` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 103.0 KiB | [postgresql-15-similarity_1.0-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg22.04+1_arm64.deb) |
| `postgresql-15-similarity` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 99.3 KiB | [postgresql-15-similarity_1.0-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg24.04+1_amd64.deb) |
| `postgresql-15-similarity` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.4 KiB | [postgresql-15-similarity_1.0-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-15-similarity_1.0-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_similarity_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 45.9 KiB | [pg_similarity_14-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_14-1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_similarity_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.9 KiB | [pg_similarity_14-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_14-1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_similarity_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 44.2 KiB | [pg_similarity_14-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_14-1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_similarity_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 42.5 KiB | [pg_similarity_14-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_14-1.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_similarity_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.7 KiB | [pg_similarity_14-1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_similarity_14-1.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_similarity_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 44.5 KiB | [pg_similarity_14-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_similarity_14-1.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_similarity_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.2 KiB | [pg_similarity_14-1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_similarity_14-1.0-3PGDG.rhel10.aarch64.rpm) |
| `pg_similarity_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.0 KiB | [pg_similarity_14-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_similarity_14-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-similarity` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 99.5 KiB | [postgresql-14-similarity_1.0-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg12+1_amd64.deb) |
| `postgresql-14-similarity` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 96.8 KiB | [postgresql-14-similarity_1.0-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg12+1_arm64.deb) |
| `postgresql-14-similarity` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 99.4 KiB | [postgresql-14-similarity_1.0-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg13+1_amd64.deb) |
| `postgresql-14-similarity` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 96.9 KiB | [postgresql-14-similarity_1.0-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg13+1_arm64.deb) |
| `postgresql-14-similarity` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 105.4 KiB | [postgresql-14-similarity_1.0-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg22.04+1_amd64.deb) |
| `postgresql-14-similarity` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 102.8 KiB | [postgresql-14-similarity_1.0-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg22.04+1_arm64.deb) |
| `postgresql-14-similarity` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 99.3 KiB | [postgresql-14-similarity_1.0-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg24.04+1_amd64.deb) |
| `postgresql-14-similarity` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.4 KiB | [postgresql-14-similarity_1.0-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-14-similarity_1.0-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/eulerto/pg_similarity" title="Repository" icon="github" subtitle="github.com/eulerto/pg_similarity" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_similarity-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_similarity;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_similarity;		# install via package name, for the active PG version

pig install pg_similarity -v 18;   # install for PG 18
pig install pg_similarity -v 17;   # install for PG 17
pig install pg_similarity -v 16;   # install for PG 16
pig install pg_similarity -v 15;   # install for PG 15
pig install pg_similarity -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_similarity;
```



## Usage

> [pg_similarity](https://github.com/eulerto/pg_similarity): Support similarity queries on PostgreSQL.
> Source: [README.md](https://raw.githubusercontent.com/eulerto/pg_similarity/master/README.md)

**pg_similarity** is an extension to support similarity queries on PostgreSQL. The implementation is tightly integrated in the RDBMS in the sense that it defines operators so instead of the traditional operators (`=` and `<>`) you can use `~~~` and `~!~` (any of these operators represents a similarity function).

**pg_similarity** has three main components:

- **Functions**: a set of functions that implements similarity algorithms available in the literature. These functions can be used as UDFs and will be the base for implementing the similarity operators;
- **Operators**: a set of operators defined at the top of similarity functions. They use similarity functions to obtain the similarity threshold and compare its value to a user-defined threshold to decide if it is a match or not;
- **Session Variables**: a set of variables that store similarity function parameters. These variables can be defined at run time.


--------

## Functions and Operators

This extension supports a set of similarity algorithms. The most known algorithms are covered by this extension. You must be aware that each algorithm is suited for a specific domain. The following algorithms are provided:

- L1 Distance (as known as City Block or Manhattan Distance)
- Cosine Distance
- Dice Coefficient
- Euclidean Distance
- Hamming Distance
- Jaccard Coefficient
- Jaro Distance
- Jaro-Winkler Distance
- Levenshtein Distance
- Matching Coefficient
- Monge-Elkan Coefficient
- Needleman-Wunsch Coefficient
- Overlap Coefficient
- Q-Gram Distance
- Smith-Waterman Coefficient
- Smith-Waterman-Gotoh Coefficient
- Soundex Distance

| Algorithm | Function | Operator | Use Index? | Parameters |
|---|---|---|---|---|
| L1 Distance | `block(text, text) returns float8` | `~++` | yes | `pg_similarity.block_tokenizer`, `pg_similarity.block_threshold`, `pg_similarity.block_is_normalized` |
| Cosine Distance | `cosine(text, text) returns float8` | `~##` | yes | `pg_similarity.cosine_tokenizer`, `pg_similarity.cosine_threshold`, `pg_similarity.cosine_is_normalized` |
| Dice Coefficient | `dice(text, text) returns float8` | `~-~` | yes | `pg_similarity.dice_tokenizer`, `pg_similarity.dice_threshold`, `pg_similarity.dice_is_normalized` |
| Euclidean Distance | `euclidean(text, text) returns float8` | `~!!` | yes | `pg_similarity.euclidean_tokenizer`, `pg_similarity.euclidean_threshold`, `pg_similarity.euclidean_is_normalized` |
| Hamming Distance | `hamming(bit varying, bit varying) returns float8` / `hamming_text(text, text) returns float8` | `~@~` | no | `pg_similarity.hamming_threshold`, `pg_similarity.hamming_is_normalized` |
| Jaccard Coefficient | `jaccard(text, text) returns float8` | `~??` | yes | `pg_similarity.jaccard_tokenizer`, `pg_similarity.jaccard_threshold`, `pg_similarity.jaccard_is_normalized` |
| Jaro Distance | `jaro(text, text) returns float8` | `~%%` | no | `pg_similarity.jaro_threshold`, `pg_similarity.jaro_is_normalized` |
| Jaro-Winkler Distance | `jarowinkler(text, text) returns float8` | `~@@` | no | `pg_similarity.jarowinkler_threshold`, `pg_similarity.jarowinkler_is_normalized` |
| Levenshtein Distance | `lev(text, text) returns float8` | `~==` | no | `pg_similarity.levenshtein_threshold`, `pg_similarity.levenshtein_is_normalized` |
| Matching Coefficient | `matchingcoefficient(text, text) returns float8` | `~^^` | yes | `pg_similarity.matching_tokenizer`, `pg_similarity.matching_threshold`, `pg_similarity.matching_is_normalized` |
| Monge-Elkan Coefficient | `mongeelkan(text, text) returns float8` | `~\|\|` | no | `pg_similarity.mongeelkan_tokenizer`, `pg_similarity.mongeelkan_threshold`, `pg_similarity.mongeelkan_is_normalized` |
| Needleman-Wunsch Coefficient | `needlemanwunsch(text, text) returns float8` | `~#~` | no | `pg_similarity.nw_threshold`, `pg_similarity.nw_is_normalized` |
| Overlap Coefficient | `overlapcoefficient(text, text) returns float8` | `~**` | yes | `pg_similarity.overlap_tokenizer`, `pg_similarity.overlap_threshold`, `pg_similarity.overlap_is_normalized` |
| Q-Gram Distance | `qgram(text, text) returns float8` | `~~~` | yes | `pg_similarity.qgram_threshold`, `pg_similarity.qgram_is_normalized` |
| Smith-Waterman Coefficient | `smithwaterman(text, text) returns float8` | `~=~` | no | `pg_similarity.sw_threshold`, `pg_similarity.sw_is_normalized` |
| Smith-Waterman-Gotoh Coefficient | `smithwatermangotoh(text, text) returns float8` | `~!~` | no | `pg_similarity.swg_threshold`, `pg_similarity.swg_is_normalized` |
| Soundex Distance | `soundex(text, text) returns float8` | `~*~` | no | |


--------

## Parameters

The several parameters control the behavior of the pg_similarity functions and operators. They can be classified in three classes:

- **tokenizer**: controls how the strings are tokenized. Valid values are **alnum**, **gram**, **word**, and **camelcase**. All tokens are lowercase. Default is **alnum**.
  - **alnum**: delimiters are any non-alphanumeric characters.
  - **gram**: an n-gram is a subsequence of length n, extracted using sliding-by-one technique.
  - **word**: delimiters are white space characters.
  - **camelcase**: delimiters are capitalized characters but they are also included as first token characters.
- **threshold**: controls how flexible the result set will be. Values range from **0.0** to **1.0**. Default is **0.7**.
- **normalized**: controls whether the similarity coefficient/distance is normalized (between 0.0 and 1.0) or not. Default is **true**.


--------

## Examples

Set parameters at run time:

```sql
SHOW pg_similarity.levenshtein_threshold;
-- 0.7

SET pg_similarity.levenshtein_threshold TO 0.5;

SET pg_similarity.cosine_tokenizer TO camelcase;

SET pg_similarity.euclidean_is_normalized TO false;
```

Simple tables for examples:

```sql
CREATE TABLE foo (a text);
INSERT INTO foo VALUES('Euler'),('Oiler'),('Euler Taveira de Oliveira'),('Maria Taveira dos Santos'),('Carlos Santos Silva');

CREATE TABLE bar (b text);
INSERT INTO bar VALUES('Euler T. de Oliveira'),('Euller'),('Oliveira, Euler Taveira'),('Sr. Oliveira');
```

### Using similarity functions

```sql
SELECT a, b, cosine(a,b), jaro(a, b), euclidean(a, b) FROM foo, bar;
```

### Using the levenshtein operator (~==)

```sql
SHOW pg_similarity.levenshtein_threshold;
-- 0.7

SELECT a, b, lev(a,b) FROM foo, bar WHERE a ~== b;
--              a             |          b           |   lev
-- ---------------------------+----------------------+----------
--  Euler                     | Euller               | 0.833333
--  Euler Taveira de Oliveira | Euler T. de Oliveira |     0.76

SET pg_similarity.levenshtein_threshold TO 0.5;

SELECT a, b, lev(a,b) FROM foo, bar WHERE a ~== b;
--              a             |          b           |   lev
-- ---------------------------+----------------------+----------
--  Euler                     | Euller               | 0.833333
--  Oiler                     | Euller               |      0.5
--  Euler Taveira de Oliveira | Euler T. de Oliveira |     0.76
```

### Using the qgram operator (~~~)

```sql
SET pg_similarity.qgram_threshold TO 0.7;

SELECT a, b, qgram(a, b) FROM foo, bar WHERE a ~~~ b;
--              a             |            b            |  qgram
-- ---------------------------+-------------------------+----------
--  Euler                     | Euller                  |      0.8
--  Euler Taveira de Oliveira | Euler T. de Oliveira    |  0.77551
--  Euler Taveira de Oliveira | Oliveira, Euler Taveira | 0.807692
```

### Comparing different operators

```sql
SELECT * FROM bar WHERE b ~@@ 'euler'; -- jaro-winkler operator
SELECT * FROM bar WHERE b ~~~ 'euler'; -- qgram operator
SELECT * FROM bar WHERE b ~== 'euler'; -- levenshtein operator
SELECT * FROM bar WHERE b ~## 'euler'; -- cosine operator
```
