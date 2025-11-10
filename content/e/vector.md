---
title: "vector"
linkTitle: "vector"
description: "vector data type and ivfflat and hnsw access methods"
weight: 1800
categories: ["RAG"]
width: full
---

[**pgvector**](https://github.com/pgvector/pgvector)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1800** | {{< badge content="vector" link="https://github.com/pgvector/pgvector" >}} | {{< ext "vector" "pgvector" >}} | `0.8.1` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "documentdb" >}} {{< ext "vchord" >}} {{< ext "vectorize" >}} {{< ext "vectorscale" >}} |
|   **See Also**    | {{< ext "pg_bestmatch" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "pg4ml" >}} {{< ext "pgml" >}} {{< ext "pg_similarity" >}} {{< ext "smlar" >}} {{< ext "pg_search" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/vector" >}} | `0.8.1` | {{< bg "18" "pgvector_18*" "green" >}} {{< bg "17" "pgvector_17*" "green" >}} {{< bg "16" "pgvector_16*" "green" >}} {{< bg "15" "pgvector_15*" "green" >}} {{< bg "14" "pgvector_14*" "green" >}} {{< bg "13" "pgvector_13*" "green" >}} | `pgvector_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/vector" >}} | `0.8.1` | {{< bg "18" "postgresql-18-pgvector" "green" >}} {{< bg "17" "postgresql-17-pgvector" "green" >}} {{< bg "16" "postgresql-16-pgvector" "green" >}} {{< bg "15" "postgresql-15-pgvector" "green" >}} {{< bg "14" "postgresql-14-pgvector" "green" >}} {{< bg "13" "postgresql-13-pgvector" "green" >}} | `postgresql-$v-pgvector` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.8.1" "pgvector_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_16 : AVAIL 13" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_15 : AVAIL 15" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_14 : AVAIL 15" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_13 : AVAIL 15" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.8.1" "pgvector_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_16 : AVAIL 13" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_15 : AVAIL 15" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_14 : AVAIL 15" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_13 : AVAIL 15" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.8.1" "pgvector_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_16 : AVAIL 13" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_15 : AVAIL 15" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_14 : AVAIL 15" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_13 : AVAIL 15" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.8.1" "pgvector_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_16 : AVAIL 13" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_15 : AVAIL 15" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_14 : AVAIL 15" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_13 : AVAIL 15" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.8.1" "pgvector_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.8.1" "pgvector_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.8.1" "pgvector_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.8.1" "postgresql-18-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-17-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-16-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-15-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-14-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-13-pgvector : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.8.1" "postgresql-18-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-17-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-16-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-15-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-14-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-13-pgvector : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.8.1" "postgresql-18-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-17-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-16-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-15-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-14-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-13-pgvector : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.8.1" "postgresql-18-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-17-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-16-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-15-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-14-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-13-pgvector : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.8.1" "postgresql-18-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-17-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-16-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-15-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-14-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-13-pgvector : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.8.1" "postgresql-18-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-17-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-16-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-15-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-14-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-13-pgvector : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.8.1" "postgresql-18-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-17-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-16-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-15-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-14-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-13-pgvector : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.8.1" "postgresql-18-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-17-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-16-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-15-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-14-pgvector : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.8.1" "postgresql-13-pgvector : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvector_18` | `0.8.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.9 KiB | [pgvector_18-0.8.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgvector_18-0.8.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_18` | `0.8.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 95.0 KiB | [pgvector_18-0.8.1-1PGDG.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvector_18-0.8.1-1PGDG.el8.aarch64.rpm) |
| `pgvector_18` | `0.8.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.7 KiB | [pgvector_18-0.8.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgvector_18-0.8.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_18` | `0.8.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 103.1 KiB | [pgvector_18-0.8.1-1PGDG.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvector_18-0.8.1-1PGDG.el9.x86_64.rpm) |
| `pgvector_18` | `0.8.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.5 KiB | [pgvector_18-0.8.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgvector_18-0.8.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_18` | `0.8.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 92.8 KiB | [pgvector_18-0.8.1-1PGDG.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvector_18-0.8.1-1PGDG.el9.aarch64.rpm) |
| `pgvector_18` | `0.8.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.2 KiB | [pgvector_18-0.8.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgvector_18-0.8.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_18` | `0.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.9 KiB | [pgvector_18-0.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgvector_18-0.8.1-1PGDG.rhel10.x86_64.rpm) |
| `pgvector_18` | `0.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.8 KiB | [pgvector_18-0.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgvector_18-0.8.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgvector` | `0.8.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 256.3 KiB | [postgresql-18-pgvector_0.8.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.1-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 226.8 KiB | [postgresql-18-pgvector_0.8.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.1-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 256.9 KiB | [postgresql-18-pgvector_0.8.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.1-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 228.3 KiB | [postgresql-18-pgvector_0.8.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.1-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 259.4 KiB | [postgresql-18-pgvector_0.8.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 227.6 KiB | [postgresql-18-pgvector_0.8.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 252.9 KiB | [postgresql-18-pgvector_0.8.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 223.5 KiB | [postgresql-18-pgvector_0.8.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvector_17` | `0.8.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.8 KiB | [pgvector_17-0.8.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgvector_17-0.8.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_17` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.9 KiB | [pgvector_17-0.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgvector_17-0.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_17` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.6 KiB | [pgvector_17-0.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgvector_17-0.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_17` | `0.8.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.5 KiB | [pgvector_17-0.8.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgvector_17-0.8.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_17` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.6 KiB | [pgvector_17-0.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgvector_17-0.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_17` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.5 KiB | [pgvector_17-0.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgvector_17-0.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_17` | `0.8.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.4 KiB | [pgvector_17-0.8.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgvector_17-0.8.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_17` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.6 KiB | [pgvector_17-0.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgvector_17-0.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_17` | `0.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.7 KiB | [pgvector_17-0.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgvector_17-0.7.4-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_17` | `0.8.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.1 KiB | [pgvector_17-0.8.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgvector_17-0.8.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_17` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.8 KiB | [pgvector_17-0.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgvector_17-0.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_17` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 89.6 KiB | [pgvector_17-0.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgvector_17-0.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_17` | `0.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.1 KiB | [pgvector_17-0.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgvector_17-0.8.1-1PGDG.rhel10.x86_64.rpm) |
| `pgvector_17` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.7 KiB | [pgvector_17-0.8.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgvector_17-0.8.0-2PGDG.rhel10.x86_64.rpm) |
| `pgvector_17` | `0.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.8 KiB | [pgvector_17-0.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgvector_17-0.8.1-1PGDG.rhel10.aarch64.rpm) |
| `pgvector_17` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.2 KiB | [pgvector_17-0.8.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgvector_17-0.8.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgvector` | `0.8.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 255.9 KiB | [postgresql-17-pgvector_0.8.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.1-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 226.5 KiB | [postgresql-17-pgvector_0.8.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.1-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 256.5 KiB | [postgresql-17-pgvector_0.8.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.1-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 227.7 KiB | [postgresql-17-pgvector_0.8.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.1-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 298.2 KiB | [postgresql-17-pgvector_0.8.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 265.9 KiB | [postgresql-17-pgvector_0.8.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 252.4 KiB | [postgresql-17-pgvector_0.8.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 222.8 KiB | [postgresql-17-pgvector_0.8.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvector_16` | `0.8.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.8 KiB | [pgvector_16-0.8.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.8.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.6 KiB | [pgvector_16-0.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.4 KiB | [pgvector_16-0.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.1 KiB | [pgvector_16-0.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.8 KiB | [pgvector_16-0.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.7.2-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.3 KiB | [pgvector_16-0.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 99.8 KiB | [pgvector_16-0.7.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.7.0-2PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.1 KiB | [pgvector_16-0.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.1 KiB | [pgvector_16-0.6.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.6.2-2PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.0 KiB | [pgvector_16-0.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.0 KiB | [pgvector_16-0.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.9 KiB | [pgvector_16-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.4 KiB | [pgvector_16-0.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.8.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.3 KiB | [pgvector_16-0.8.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.8.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.4 KiB | [pgvector_16-0.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.3 KiB | [pgvector_16-0.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.0 KiB | [pgvector_16-0.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.7 KiB | [pgvector_16-0.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.7.2-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.3 KiB | [pgvector_16-0.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 89.7 KiB | [pgvector_16-0.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 70.2 KiB | [pgvector_16-0.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 70.4 KiB | [pgvector_16-0.6.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.6.2-2PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 69.1 KiB | [pgvector_16-0.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 68.3 KiB | [pgvector_16-0.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 58.6 KiB | [pgvector_16-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 58.0 KiB | [pgvector_16-0.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.8.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.0 KiB | [pgvector_16-0.8.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.8.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.3 KiB | [pgvector_16-0.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.4 KiB | [pgvector_16-0.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.7.4-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.0 KiB | [pgvector_16-0.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.7.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.6 KiB | [pgvector_16-0.7.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.7.2-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.2 KiB | [pgvector_16-0.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 101.4 KiB | [pgvector_16-0.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.8 KiB | [pgvector_16-0.6.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.6.2-2PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.9 KiB | [pgvector_16-0.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.8 KiB | [pgvector_16-0.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.1 KiB | [pgvector_16-0.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.4 KiB | [pgvector_16-0.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.1 KiB | [pgvector_16-0.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_16` | `0.8.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.9 KiB | [pgvector_16-0.8.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.8.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.5 KiB | [pgvector_16-0.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 89.6 KiB | [pgvector_16-0.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 89.2 KiB | [pgvector_16-0.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 88.6 KiB | [pgvector_16-0.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.7.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 88.4 KiB | [pgvector_16-0.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 87.6 KiB | [pgvector_16-0.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.8 KiB | [pgvector_16-0.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.2 KiB | [pgvector_16-0.6.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.6.2-2PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.6 KiB | [pgvector_16-0.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.0 KiB | [pgvector_16-0.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 59.8 KiB | [pgvector_16-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 59.6 KiB | [pgvector_16-0.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.7 KiB | [pgvector_16-0.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgvector_16-0.8.1-1PGDG.rhel10.x86_64.rpm) |
| `pgvector_16` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.2 KiB | [pgvector_16-0.8.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgvector_16-0.8.0-2PGDG.rhel10.x86_64.rpm) |
| `pgvector_16` | `0.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.3 KiB | [pgvector_16-0.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgvector_16-0.8.1-1PGDG.rhel10.aarch64.rpm) |
| `pgvector_16` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 95.8 KiB | [pgvector_16-0.8.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgvector_16-0.8.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgvector` | `0.8.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 255.5 KiB | [postgresql-16-pgvector_0.8.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.1-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 226.2 KiB | [postgresql-16-pgvector_0.8.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.1-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 256.3 KiB | [postgresql-16-pgvector_0.8.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.1-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 227.1 KiB | [postgresql-16-pgvector_0.8.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.1-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 288.7 KiB | [postgresql-16-pgvector_0.8.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 256.5 KiB | [postgresql-16-pgvector_0.8.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 251.5 KiB | [postgresql-16-pgvector_0.8.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 222.2 KiB | [postgresql-16-pgvector_0.8.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvector_15` | `0.8.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.8 KiB | [pgvector_15-0.8.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.8.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.7 KiB | [pgvector_15-0.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.4 KiB | [pgvector_15-0.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.0 KiB | [pgvector_15-0.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.7 KiB | [pgvector_15-0.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.7.2-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.3 KiB | [pgvector_15-0.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.8 KiB | [pgvector_15-0.7.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.7.0-2PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.9 KiB | [pgvector_15-0.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.9 KiB | [pgvector_15-0.6.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.6.2-2PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.8 KiB | [pgvector_15-0.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.1 KiB | [pgvector_15-0.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 64.1 KiB | [pgvector_15-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.5 KiB | [pgvector_15-0.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.9 KiB | [pgvector_15-0.4.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.4.4-1.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pgvector_15-0.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.4.1-1.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.8.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.2 KiB | [pgvector_15-0.8.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.8.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.3 KiB | [pgvector_15-0.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 92.1 KiB | [pgvector_15-0.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.9 KiB | [pgvector_15-0.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.4 KiB | [pgvector_15-0.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.7.2-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.0 KiB | [pgvector_15-0.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.4 KiB | [pgvector_15-0.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.0 KiB | [pgvector_15-0.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.3 KiB | [pgvector_15-0.6.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.6.2-2PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 69.9 KiB | [pgvector_15-0.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 69.3 KiB | [pgvector_15-0.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.7 KiB | [pgvector_15-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.0 KiB | [pgvector_15-0.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.5 KiB | [pgvector_15-0.4.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.4.4-1.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.6 KiB | [pgvector_15-0.4.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.4.1-1.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.8.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.6 KiB | [pgvector_15-0.8.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.8.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.0 KiB | [pgvector_15-0.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.9 KiB | [pgvector_15-0.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.7.4-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.3 KiB | [pgvector_15-0.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.7.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.9 KiB | [pgvector_15-0.7.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.7.2-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.7 KiB | [pgvector_15-0.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.1 KiB | [pgvector_15-0.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.7 KiB | [pgvector_15-0.6.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.6.2-2PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 78.6 KiB | [pgvector_15-0.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.5 KiB | [pgvector_15-0.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.7 KiB | [pgvector_15-0.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 66.0 KiB | [pgvector_15-0.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.4 KiB | [pgvector_15-0.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.8 KiB | [pgvector_15-0.4.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.4.4-1.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.0 KiB | [pgvector_15-0.4.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.4.1-1.rhel9.x86_64.rpm) |
| `pgvector_15` | `0.8.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.8 KiB | [pgvector_15-0.8.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.8.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.5 KiB | [pgvector_15-0.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.2 KiB | [pgvector_15-0.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.8 KiB | [pgvector_15-0.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.5 KiB | [pgvector_15-0.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.7.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.1 KiB | [pgvector_15-0.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 92.3 KiB | [pgvector_15-0.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.1 KiB | [pgvector_15-0.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.5 KiB | [pgvector_15-0.6.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.6.2-2PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.0 KiB | [pgvector_15-0.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.3 KiB | [pgvector_15-0.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.5 KiB | [pgvector_15-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.9 KiB | [pgvector_15-0.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.4 KiB | [pgvector_15-0.4.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.4.4-1.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.6 KiB | [pgvector_15-0.4.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.4.1-1.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 108.7 KiB | [pgvector_15-0.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgvector_15-0.8.1-1PGDG.rhel10.x86_64.rpm) |
| `pgvector_15` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 108.5 KiB | [pgvector_15-0.8.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgvector_15-0.8.0-2PGDG.rhel10.x86_64.rpm) |
| `pgvector_15` | `0.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.0 KiB | [pgvector_15-0.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgvector_15-0.8.1-1PGDG.rhel10.aarch64.rpm) |
| `pgvector_15` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 100.8 KiB | [pgvector_15-0.8.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgvector_15-0.8.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgvector` | `0.8.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 256.2 KiB | [postgresql-15-pgvector_0.8.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.1-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 227.7 KiB | [postgresql-15-pgvector_0.8.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.1-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 257.7 KiB | [postgresql-15-pgvector_0.8.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.1-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 228.8 KiB | [postgresql-15-pgvector_0.8.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.1-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 291.9 KiB | [postgresql-15-pgvector_0.8.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 259.9 KiB | [postgresql-15-pgvector_0.8.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 254.2 KiB | [postgresql-15-pgvector_0.8.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 226.0 KiB | [postgresql-15-pgvector_0.8.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvector_14` | `0.8.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.6 KiB | [pgvector_14-0.8.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.8.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.5 KiB | [pgvector_14-0.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.3 KiB | [pgvector_14-0.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.9 KiB | [pgvector_14-0.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.5 KiB | [pgvector_14-0.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.7.2-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.2 KiB | [pgvector_14-0.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.7 KiB | [pgvector_14-0.7.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.7.0-2PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.8 KiB | [pgvector_14-0.6.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.6.2-2PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.8 KiB | [pgvector_14-0.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.6 KiB | [pgvector_14-0.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.1 KiB | [pgvector_14-0.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 64.2 KiB | [pgvector_14-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.4 KiB | [pgvector_14-0.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.8 KiB | [pgvector_14-0.4.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.4.4-1.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.7 KiB | [pgvector_14-0.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.4.1-1.rhel8.x86_64.rpm) |
| `pgvector_14` | `0.8.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.1 KiB | [pgvector_14-0.8.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.8.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.1 KiB | [pgvector_14-0.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 92.1 KiB | [pgvector_14-0.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.7 KiB | [pgvector_14-0.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.3 KiB | [pgvector_14-0.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.7.2-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.9 KiB | [pgvector_14-0.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.2 KiB | [pgvector_14-0.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.2 KiB | [pgvector_14-0.6.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.6.2-2PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 70.9 KiB | [pgvector_14-0.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 69.8 KiB | [pgvector_14-0.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 69.2 KiB | [pgvector_14-0.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.6 KiB | [pgvector_14-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.0 KiB | [pgvector_14-0.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.5 KiB | [pgvector_14-0.4.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.4.4-1.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.5 KiB | [pgvector_14-0.4.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.4.1-1.rhel8.aarch64.rpm) |
| `pgvector_14` | `0.8.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.4 KiB | [pgvector_14-0.8.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.8.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.7 KiB | [pgvector_14-0.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.8 KiB | [pgvector_14-0.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.7.4-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.1 KiB | [pgvector_14-0.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.7.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.0 KiB | [pgvector_14-0.7.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.7.2-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.5 KiB | [pgvector_14-0.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.0 KiB | [pgvector_14-0.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.8 KiB | [pgvector_14-0.6.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.6.2-2PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 78.5 KiB | [pgvector_14-0.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.4 KiB | [pgvector_14-0.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.8 KiB | [pgvector_14-0.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.8 KiB | [pgvector_14-0.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.4 KiB | [pgvector_14-0.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.9 KiB | [pgvector_14-0.4.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.4.4-1.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.9 KiB | [pgvector_14-0.4.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.4.1-1.rhel9.x86_64.rpm) |
| `pgvector_14` | `0.8.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.3 KiB | [pgvector_14-0.8.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.8.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.1 KiB | [pgvector_14-0.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.2 KiB | [pgvector_14-0.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.8 KiB | [pgvector_14-0.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.5 KiB | [pgvector_14-0.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.7.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.1 KiB | [pgvector_14-0.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 92.4 KiB | [pgvector_14-0.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.1 KiB | [pgvector_14-0.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.4 KiB | [pgvector_14-0.6.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.6.2-2PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.0 KiB | [pgvector_14-0.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.3 KiB | [pgvector_14-0.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.5 KiB | [pgvector_14-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.9 KiB | [pgvector_14-0.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.3 KiB | [pgvector_14-0.4.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.4.4-1.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.5 KiB | [pgvector_14-0.4.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.4.1-1.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 108.9 KiB | [pgvector_14-0.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgvector_14-0.8.1-1PGDG.rhel10.x86_64.rpm) |
| `pgvector_14` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 108.6 KiB | [pgvector_14-0.8.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgvector_14-0.8.0-2PGDG.rhel10.x86_64.rpm) |
| `pgvector_14` | `0.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.2 KiB | [pgvector_14-0.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgvector_14-0.8.1-1PGDG.rhel10.aarch64.rpm) |
| `pgvector_14` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 100.7 KiB | [pgvector_14-0.8.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgvector_14-0.8.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgvector` | `0.8.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 256.3 KiB | [postgresql-14-pgvector_0.8.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.1-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 227.9 KiB | [postgresql-14-pgvector_0.8.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.1-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 257.4 KiB | [postgresql-14-pgvector_0.8.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.1-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 228.9 KiB | [postgresql-14-pgvector_0.8.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.1-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 289.8 KiB | [postgresql-14-pgvector_0.8.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 258.1 KiB | [postgresql-14-pgvector_0.8.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 254.1 KiB | [postgresql-14-pgvector_0.8.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 225.6 KiB | [postgresql-14-pgvector_0.8.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvector_13` | `0.8.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.1 KiB | [pgvector_13-0.8.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.8.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.0 KiB | [pgvector_13-0.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.7 KiB | [pgvector_13-0.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.4 KiB | [pgvector_13-0.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.1 KiB | [pgvector_13-0.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.7.2-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 99.6 KiB | [pgvector_13-0.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 99.2 KiB | [pgvector_13-0.7.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.7.0-2PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.8 KiB | [pgvector_13-0.6.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.6.2-2PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.6 KiB | [pgvector_13-0.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.5 KiB | [pgvector_13-0.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.9 KiB | [pgvector_13-0.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.5 KiB | [pgvector_13-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.9 KiB | [pgvector_13-0.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.3 KiB | [pgvector_13-0.4.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.4.4-1.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.1 KiB | [pgvector_13-0.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgvector_13-0.4.1-1.rhel8.x86_64.rpm) |
| `pgvector_13` | `0.8.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.6 KiB | [pgvector_13-0.8.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.8.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.7 KiB | [pgvector_13-0.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.5 KiB | [pgvector_13-0.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.2 KiB | [pgvector_13-0.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.8 KiB | [pgvector_13-0.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.7.2-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.4 KiB | [pgvector_13-0.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 89.8 KiB | [pgvector_13-0.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 70.7 KiB | [pgvector_13-0.6.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.6.2-2PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 70.4 KiB | [pgvector_13-0.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 69.3 KiB | [pgvector_13-0.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 68.7 KiB | [pgvector_13-0.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.5 KiB | [pgvector_13-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.0 KiB | [pgvector_13-0.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.4 KiB | [pgvector_13-0.4.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.4.4-1.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.4 KiB | [pgvector_13-0.4.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgvector_13-0.4.1-1.rhel8.aarch64.rpm) |
| `pgvector_13` | `0.8.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.9 KiB | [pgvector_13-0.8.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.8.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.1 KiB | [pgvector_13-0.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.2 KiB | [pgvector_13-0.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.7.4-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.8 KiB | [pgvector_13-0.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.7.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 106.4 KiB | [pgvector_13-0.7.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.7.2-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 105.7 KiB | [pgvector_13-0.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 105.2 KiB | [pgvector_13-0.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.2 KiB | [pgvector_13-0.6.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.6.2-2PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 78.1 KiB | [pgvector_13-0.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.9 KiB | [pgvector_13-0.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.7 KiB | [pgvector_13-0.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.5 KiB | [pgvector_13-0.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.3 KiB | [pgvector_13-0.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.8 KiB | [pgvector_13-0.4.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.4.4-1.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.7 KiB | [pgvector_13-0.4.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgvector_13-0.4.1-1.rhel9.x86_64.rpm) |
| `pgvector_13` | `0.8.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.6 KiB | [pgvector_13-0.8.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.8.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.3 KiB | [pgvector_13-0.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.1 KiB | [pgvector_13-0.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.8 KiB | [pgvector_13-0.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.4 KiB | [pgvector_13-0.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.7.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.1 KiB | [pgvector_13-0.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 92.4 KiB | [pgvector_13-0.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.8 KiB | [pgvector_13-0.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.2 KiB | [pgvector_13-0.6.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.6.2-2PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.7 KiB | [pgvector_13-0.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.0 KiB | [pgvector_13-0.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.5 KiB | [pgvector_13-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.8 KiB | [pgvector_13-0.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.4 KiB | [pgvector_13-0.4.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.4.4-1.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.4 KiB | [pgvector_13-0.4.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgvector_13-0.4.1-1.rhel9.aarch64.rpm) |
| `pgvector_13` | `0.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 108.1 KiB | [pgvector_13-0.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgvector_13-0.8.1-1PGDG.rhel10.x86_64.rpm) |
| `pgvector_13` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 107.7 KiB | [pgvector_13-0.8.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgvector_13-0.8.0-2PGDG.rhel10.x86_64.rpm) |
| `pgvector_13` | `0.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.1 KiB | [pgvector_13-0.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgvector_13-0.8.1-1PGDG.rhel10.aarch64.rpm) |
| `pgvector_13` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 100.6 KiB | [pgvector_13-0.8.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgvector_13-0.8.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pgvector` | `0.8.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 255.1 KiB | [postgresql-13-pgvector_0.8.1-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-13-pgvector_0.8.1-2.pgdg12+1_amd64.deb) |
| `postgresql-13-pgvector` | `0.8.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 227.0 KiB | [postgresql-13-pgvector_0.8.1-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-13-pgvector_0.8.1-2.pgdg12+1_arm64.deb) |
| `postgresql-13-pgvector` | `0.8.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 256.2 KiB | [postgresql-13-pgvector_0.8.1-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-13-pgvector_0.8.1-2.pgdg13+1_amd64.deb) |
| `postgresql-13-pgvector` | `0.8.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 227.7 KiB | [postgresql-13-pgvector_0.8.1-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-13-pgvector_0.8.1-2.pgdg13+1_arm64.deb) |
| `postgresql-13-pgvector` | `0.8.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 288.1 KiB | [postgresql-13-pgvector_0.8.1-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-13-pgvector_0.8.1-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pgvector` | `0.8.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 256.7 KiB | [postgresql-13-pgvector_0.8.1-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-13-pgvector_0.8.1-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pgvector` | `0.8.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 253.8 KiB | [postgresql-13-pgvector_0.8.1-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-13-pgvector_0.8.1-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pgvector` | `0.8.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 225.3 KiB | [postgresql-13-pgvector_0.8.1-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-13-pgvector_0.8.1-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgvector/pgvector" title="Repository" icon="github" subtitle="github.com/pgvector/pgvector" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgvector-0.8.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get vector; # get vector source code
pig build dep vector; # install build dependencies
pig build pkg vector; # build extension rpm or deb
pig build ext vector; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install vector; # install by extension name, for the current active PG version
pig ext install pgvector; # install via package alias, for the active PG version
pig ext install vector -v 18;   # install for PG 18
pig ext install vector -v 17;   # install for PG 17
pig ext install vector -v 16;   # install for PG 16
pig ext install vector -v 15;   # install for PG 15
pig ext install vector -v 14;   # install for PG 14
pig ext install vector -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION vector;
```



--------

## Usage

###  Getting Started

Create a vector column with 3 dimensions

```sql
CREATE TABLE items (id bigserial PRIMARY KEY, embedding vector(3));
```

Insert vectors

```sql
INSERT INTO items (embedding) VALUES ('[1,2,3]'), ('[4,5,6]');
```

Get the nearest neighbors by L2 distance

```sql
SELECT * FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

Also supports inner product (`<#>`), cosine distance (`<=>`), and L1 distance (`<+>`)

Note: `<#>` returns the negative inner product since Postgres only supports `ASC` order index scans on operators


--------

### CRUD

Create a new table with a vector column

```sql
CREATE TABLE items (id bigserial PRIMARY KEY, embedding vector(3));
```

Or add a vector column to an existing table

```sql
ALTER TABLE items ADD COLUMN embedding vector(3);
```

Also supports [half-precision](https://github.com/pgvector/pgvector#half-precision-vectors), [binary](https://github.com/pgvector/pgvector#binary-vectors), and [sparse](https://github.com/pgvector/pgvector#sparse-vectors) vectors

**Insert vectors**

```sql
INSERT INTO items (embedding) VALUES ('[1,2,3]'), ('[4,5,6]');
```

Or load vectors in bulk using `COPY` ([example](https://github.com/pgvector/pgvector-python/blob/master/examples/loading/example.py))

```sql
COPY items (embedding) FROM STDIN WITH (FORMAT BINARY);
```

**Upsert vectors**

```sql
INSERT INTO items (id, embedding) VALUES (1, '[1,2,3]'), (2, '[4,5,6]')
    ON CONFLICT (id) DO UPDATE SET embedding = EXCLUDED.embedding;
```

**Update vectors**

```sql
UPDATE items SET embedding = '[1,2,3]' WHERE id = 1;
```

**Delete vectors**

```sql
DELETE FROM items WHERE id = 1;
```


--------

### Querying

Get the nearest neighbors to a vector

```sql
SELECT * FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

Supported distance functions are:

- `<->` - L2 distance
- `<#>` - (negative) inner product
- `<=>` - cosine distance
- `<+>` - L1 distance
- `<~>` - Hamming distance (binary vectors)
- `<%>` - Jaccard distance (binary vectors)

Get the nearest neighbors to a row

```sql
SELECT * FROM items WHERE id != 1 ORDER BY embedding <-> (SELECT embedding FROM items WHERE id = 1) LIMIT 5;
```

Get rows within a certain distance

```sql
SELECT * FROM items WHERE embedding <-> '[3,1,2]' < 5;
```

Note: Combine with `ORDER BY` and `LIMIT` to use an index


#### Distances

Get the distance

```sql
SELECT embedding <-> '[3,1,2]' AS distance FROM items;
```

For inner product, multiply by -1 (since `<#>` returns the negative inner product)

```sql
SELECT (embedding <#> '[3,1,2]') * -1 AS inner_product FROM items;
```

For cosine similarity, use 1 - cosine distance

```sql
SELECT 1 - (embedding <=> '[3,1,2]') AS cosine_similarity FROM items;
```

#### Aggregates

Average vectors

```sql
SELECT AVG(embedding) FROM items;
```

Average groups of vectors

```sql
SELECT category_id, AVG(embedding) FROM items GROUP BY category_id;
```


--------

### HNSW Indexing

By default, pgvector performs exact nearest neighbor search, which provides perfect recall.

You can add an index to use approximate nearest neighbor search, which trades some recall for speed. Unlike typical indexes, you will see different results for queries after adding an approximate index.

Supported index types are:

- [HNSW](https://github.com/pgvector/pgvector#hnsw)
- [IVFFlat](https://github.com/pgvector/pgvector#ivfflat)

An HNSW index creates a multilayer graph. It has better query performance than IVFFlat (in terms of speed-recall tradeoff), but has slower build times and uses more memory. Also, an index can be created without any data in the table since there isn’t a training step like IVFFlat.

Add an index for each distance function you want to use.

L2 distance

```sql
CREATE INDEX ON items USING hnsw (embedding vector_l2_ops);
```

Note: Use `halfvec_l2_ops` for `halfvec` and `sparsevec_l2_ops` for `sparsevec` (and similar with the other distance functions)

**Inner product**

```sql
CREATE INDEX ON items USING hnsw (embedding vector_ip_ops);
```

**Cosine distance**

```sql
CREATE INDEX ON items USING hnsw (embedding vector_cosine_ops);
```

**L1 distance**

```sql
CREATE INDEX ON items USING hnsw (embedding vector_l1_ops);
```

**Hamming distance**

```sql
CREATE INDEX ON items USING hnsw (embedding bit_hamming_ops);
```

**Jaccard distance**

```sql
CREATE INDEX ON items USING hnsw (embedding bit_jaccard_ops);
```

Supported types are:

- `vector` - up to 2,000 dimensions
- `halfvec` - up to 4,000 dimensions
- `bit` - up to 64,000 dimensions
- `sparsevec` - up to 1,000 non-zero elements

#### Index Options

Specify HNSW parameters

- `m` - the max number of connections per layer (16 by default)
- `ef_construction` - the size of the dynamic candidate list for constructing the graph (64 by default)

```sql
CREATE INDEX ON items USING hnsw (embedding vector_l2_ops) WITH (m = 16, ef_construction = 64);
```

A higher value of `ef_construction` provides better recall at the cost of index build time / insert speed.

#### Query Options

Specify the size of the dynamic candidate list for search (40 by default)

```sql
SET hnsw.ef_search = 100;
```

A higher value provides better recall at the cost of speed.

Use `SET LOCAL` inside a transaction to set it for a single query

```sql
BEGIN;
SET LOCAL hnsw.ef_search = 100;
SELECT ...
COMMIT;
```

#### Index Build Time

Indexes build significantly faster when the graph fits into `maintenance_work_mem`

```sql
SET maintenance_work_mem = '8GB';
```

A notice is shown when the graph no longer fits

```
NOTICE:  hnsw graph no longer fits into maintenance_work_mem after 100000 tuples
DETAIL:  Building will take significantly more time.
HINT:  Increase maintenance_work_mem to speed up builds.
```


Note: Do not set `maintenance_work_mem` so high that it exhausts the memory on the server

Like other index types, it’s faster to create an index after loading your initial data

You can also speed up index creation by increasing the number of parallel workers (2 by default)

```sql
SET max_parallel_maintenance_workers = 7; -- plus leader
```

For a large number of workers, you may need to increase `max_parallel_workers` (8 by default)

The [index options](https://github.com/pgvector/pgvector#index-options) also have a significant impact on build time (use the defaults unless seeing low recall)

#### Indexing Progress

Check [indexing progress](https://www.postgresql.org/docs/current/progress-reporting.html#CREATE-INDEX-PROGRESS-REPORTING)

```sql
SELECT phase, round(100.0 * blocks_done / nullif(blocks_total, 0), 1) AS "%" FROM pg_stat_progress_create_index;
```

The phases for HNSW are:

1. `initializing`
2. `loading tuples`
