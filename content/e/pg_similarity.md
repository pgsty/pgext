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
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_similarity` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_similarity_18*" "green" >}} {{< bg "17" "pg_similarity_17*" "green" >}} {{< bg "16" "pg_similarity_16*" "green" >}} {{< bg "15" "pg_similarity_15*" "green" >}} {{< bg "14" "pg_similarity_14*" "green" >}} {{< bg "13" "pg_similarity_13*" "green" >}} | `pg_similarity_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "postgresql-18-similarity" "green" >}} {{< bg "17" "postgresql-17-similarity" "green" >}} {{< bg "16" "postgresql-16-similarity" "green" >}} {{< bg "15" "postgresql-15-similarity" "green" >}} {{< bg "14" "postgresql-14-similarity" "green" >}} {{< bg "13" "postgresql-13-similarity" "green" >}} | `postgresql-$v-similarity` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0" "pg_similarity_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0" "pg_similarity_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0" "pg_similarity_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0" "pg_similarity_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_similarity_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0" "pg_similarity_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0" "pg_similarity_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0" "pg_similarity_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-similarity : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-similarity : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-similarity : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-similarity : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-similarity : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-similarity : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-similarity : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-similarity : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-similarity : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
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

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_similarity_13` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 45.4 KiB | [pg_similarity_13-1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_similarity_13-1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_similarity_13` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.8 KiB | [pg_similarity_13-1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_similarity_13-1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_similarity_13` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 44.1 KiB | [pg_similarity_13-1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_similarity_13-1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_similarity_13` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 42.5 KiB | [pg_similarity_13-1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_similarity_13-1.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_similarity_13` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.6 KiB | [pg_similarity_13-1.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_similarity_13-1.0-3PGDG.rhel10.x86_64.rpm) |
| `pg_similarity_13` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 44.5 KiB | [pg_similarity_13-1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_similarity_13-1.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_similarity_13` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.1 KiB | [pg_similarity_13-1.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_similarity_13-1.0-3PGDG.rhel10.aarch64.rpm) |
| `pg_similarity_13` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.0 KiB | [pg_similarity_13-1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_similarity_13-1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-similarity` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 99.0 KiB | [postgresql-13-similarity_1.0-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg12+1_amd64.deb) |
| `postgresql-13-similarity` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 96.3 KiB | [postgresql-13-similarity_1.0-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg12+1_arm64.deb) |
| `postgresql-13-similarity` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 99.2 KiB | [postgresql-13-similarity_1.0-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg13+1_amd64.deb) |
| `postgresql-13-similarity` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 96.5 KiB | [postgresql-13-similarity_1.0-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg13+1_arm64.deb) |
| `postgresql-13-similarity` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 104.9 KiB | [postgresql-13-similarity_1.0-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg22.04+1_amd64.deb) |
| `postgresql-13-similarity` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 102.2 KiB | [postgresql-13-similarity_1.0-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg22.04+1_arm64.deb) |
| `postgresql-13-similarity` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.6 KiB | [postgresql-13-similarity_1.0-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg24.04+1_amd64.deb) |
| `postgresql-13-similarity` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.0 KiB | [postgresql-13-similarity_1.0-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-similarity/postgresql-13-similarity_1.0-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/eulerto/pg_similarity" title="Repository" icon="github" subtitle="github.com/eulerto/pg_similarity" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_similarity-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_similarity;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_similarity;		# install via package name, for the active PG version

pig install pg_similarity -v 18;   # install for PG 18
pig install pg_similarity -v 17;   # install for PG 17
pig install pg_similarity -v 16;   # install for PG 16
pig install pg_similarity -v 15;   # install for PG 15
pig install pg_similarity -v 14;   # install for PG 14
pig install pg_similarity -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_similarity;
```
