---
title: "vector"
linkTitle: "vector"
description: "vector data type and ivfflat and hnsw access methods"
weight: 1800
categories: ["RAG"]
width: full
---

[**pgvector**](https://github.com/pgvector/pgvector) : vector data type and ivfflat and hnsw access methods


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1800** | {{< badge content="vector" link="https://github.com/pgvector/pgvector" >}} | {{< ext "vector" "pgvector" >}} | `0.8.5` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "documentdb" >}} {{< ext "pgmnemo" >}} {{< ext "vchord" >}} {{< ext "vectorize" >}} {{< ext "vectorscale" >}} |
|   **See Also**    | {{< ext "pg_bestmatch" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "pg4ml" >}} {{< ext "pgml" >}} {{< ext "pg_similarity" >}} {{< ext "smlar" >}} {{< ext "pg_search" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.8.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgvector` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.8.5` | {{< bg "18" "pgvector_18" "green" >}} {{< bg "17" "pgvector_17" "green" >}} {{< bg "16" "pgvector_16" "green" >}} {{< bg "15" "pgvector_15" "green" >}} {{< bg "14" "pgvector_14" "green" >}} | `pgvector_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.8.4` | {{< bg "18" "postgresql-18-pgvector" "green" >}} {{< bg "17" "postgresql-17-pgvector" "green" >}} {{< bg "16" "postgresql-16-pgvector" "green" >}} {{< bg "15" "postgresql-15-pgvector" "green" >}} {{< bg "14" "postgresql-14-pgvector" "green" >}} | `postgresql-$v-pgvector` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.8.5" "pgvector_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_16 : AVAIL 18" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_15 : AVAIL 20" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_14 : AVAIL 20" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.8.5" "pgvector_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_16 : AVAIL 18" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_15 : AVAIL 20" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_14 : AVAIL 20" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.8.5" "pgvector_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_17 : AVAIL 10" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_16 : AVAIL 20" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_15 : AVAIL 22" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_14 : AVAIL 22" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.8.5" "pgvector_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_17 : AVAIL 10" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_16 : AVAIL 20" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_15 : AVAIL 22" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_14 : AVAIL 22" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.8.5" "pgvector_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_14 : AVAIL 9" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.8.5" "pgvector_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 0.8.5" "pgvector_14 : AVAIL 9" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.8.5" "postgresql-18-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-17-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-16-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-15-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-14-pgvector : AVAIL 4" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.8.5" "postgresql-18-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-17-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-16-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-15-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-14-pgvector : AVAIL 4" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.8.5" "postgresql-18-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-17-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-16-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-15-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-14-pgvector : AVAIL 4" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.8.5" "postgresql-18-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-17-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-16-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-15-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-14-pgvector : AVAIL 4" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.8.5" "postgresql-18-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-17-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-16-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-15-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-14-pgvector : AVAIL 4" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.8.5" "postgresql-18-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-17-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-16-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-15-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-14-pgvector : AVAIL 4" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.8.5" "postgresql-18-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-17-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-16-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-15-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-14-pgvector : AVAIL 4" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.8.5" "postgresql-18-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-17-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-16-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-15-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-14-pgvector : AVAIL 4" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 0.8.5" "postgresql-18-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-17-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-16-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-15-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-14-pgvector : AVAIL 4" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 0.8.5" "postgresql-18-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-17-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-16-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-15-pgvector : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.8.5" "postgresql-14-pgvector : AVAIL 4" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvector_18` | `0.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.5 KiB | [pgvector_18-0.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgvector_18-0.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_18` | `0.8.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 115.1 KiB | [pgvector_18-0.8.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvector_18-0.8.4-1PIGSTY.el8.x86_64.rpm) |
| `pgvector_18` | `0.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.3 KiB | [pgvector_18-0.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgvector_18-0.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_18` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 108.2 KiB | [pgvector_18-0.8.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgvector_18-0.8.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_18` | `0.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.3 KiB | [pgvector_18-0.8.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgvector_18-0.8.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_18` | `0.8.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.9 KiB | [pgvector_18-0.8.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgvector_18-0.8.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_18` | `0.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.0 KiB | [pgvector_18-0.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgvector_18-0.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_18` | `0.8.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.7 KiB | [pgvector_18-0.8.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvector_18-0.8.4-1PIGSTY.el8.aarch64.rpm) |
| `pgvector_18` | `0.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.8 KiB | [pgvector_18-0.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgvector_18-0.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_18` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.9 KiB | [pgvector_18-0.8.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgvector_18-0.8.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_18` | `0.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.4 KiB | [pgvector_18-0.8.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgvector_18-0.8.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_18` | `0.8.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.7 KiB | [pgvector_18-0.8.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgvector_18-0.8.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_18` | `0.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.9 KiB | [pgvector_18-0.8.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgvector_18-0.8.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_18` | `0.8.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 108.2 KiB | [pgvector_18-0.8.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvector_18-0.8.4-1PIGSTY.el9.x86_64.rpm) |
| `pgvector_18` | `0.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.6 KiB | [pgvector_18-0.8.4-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgvector_18-0.8.4-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_18` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.9 KiB | [pgvector_18-0.8.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgvector_18-0.8.3-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_18` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [pgvector_18-0.8.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgvector_18-0.8.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_18` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [pgvector_18-0.8.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgvector_18-0.8.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pgvector_18` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.8 KiB | [pgvector_18-0.8.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgvector_18-0.8.2-1PGDG.rhel9.6.x86_64.rpm) |
| `pgvector_18` | `0.8.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.5 KiB | [pgvector_18-0.8.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgvector_18-0.8.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_18` | `0.8.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 95.7 KiB | [pgvector_18-0.8.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgvector_18-0.8.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_18` | `0.8.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 98.0 KiB | [pgvector_18-0.8.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvector_18-0.8.4-1PIGSTY.el9.aarch64.rpm) |
| `pgvector_18` | `0.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 95.3 KiB | [pgvector_18-0.8.4-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgvector_18-0.8.4-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_18` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.8 KiB | [pgvector_18-0.8.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgvector_18-0.8.3-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_18` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.7 KiB | [pgvector_18-0.8.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgvector_18-0.8.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_18` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.8 KiB | [pgvector_18-0.8.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgvector_18-0.8.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pgvector_18` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.9 KiB | [pgvector_18-0.8.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgvector_18-0.8.2-1PGDG.rhel9.6.aarch64.rpm) |
| `pgvector_18` | `0.8.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.2 KiB | [pgvector_18-0.8.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgvector_18-0.8.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_18` | `0.8.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 106.1 KiB | [pgvector_18-0.8.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgvector_18-0.8.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_18` | `0.8.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 109.2 KiB | [pgvector_18-0.8.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvector_18-0.8.4-1PIGSTY.el10.x86_64.rpm) |
| `pgvector_18` | `0.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.9 KiB | [pgvector_18-0.8.4-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgvector_18-0.8.4-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_18` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.0 KiB | [pgvector_18-0.8.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgvector_18-0.8.3-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_18` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.1 KiB | [pgvector_18-0.8.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgvector_18-0.8.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_18` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.1 KiB | [pgvector_18-0.8.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgvector_18-0.8.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pgvector_18` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.7 KiB | [pgvector_18-0.8.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgvector_18-0.8.2-1PGDG.rhel10.0.x86_64.rpm) |
| `pgvector_18` | `0.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.9 KiB | [pgvector_18-0.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgvector_18-0.8.1-1PGDG.rhel10.x86_64.rpm) |
| `pgvector_18` | `0.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 98.0 KiB | [pgvector_18-0.8.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgvector_18-0.8.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_18` | `0.8.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 100.3 KiB | [pgvector_18-0.8.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvector_18-0.8.4-1PIGSTY.el10.aarch64.rpm) |
| `pgvector_18` | `0.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 97.7 KiB | [pgvector_18-0.8.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgvector_18-0.8.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_18` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.9 KiB | [pgvector_18-0.8.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgvector_18-0.8.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_18` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.9 KiB | [pgvector_18-0.8.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgvector_18-0.8.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_18` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.9 KiB | [pgvector_18-0.8.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgvector_18-0.8.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pgvector_18` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.9 KiB | [pgvector_18-0.8.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgvector_18-0.8.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pgvector_18` | `0.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.8 KiB | [pgvector_18-0.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgvector_18-0.8.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgvector` | `0.8.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 261.3 KiB | [postgresql-18-pgvector_0.8.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.5-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 261.0 KiB | [postgresql-18-pgvector_0.8.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 254.5 KiB | [postgresql-18-pgvector_0.8.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 258.4 KiB | [postgresql-18-pgvector_0.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 231.3 KiB | [postgresql-18-pgvector_0.8.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.5-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 231.0 KiB | [postgresql-18-pgvector_0.8.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 229.2 KiB | [postgresql-18-pgvector_0.8.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 228.9 KiB | [postgresql-18-pgvector_0.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 262.1 KiB | [postgresql-18-pgvector_0.8.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.5-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 261.9 KiB | [postgresql-18-pgvector_0.8.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 254.8 KiB | [postgresql-18-pgvector_0.8.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 259.3 KiB | [postgresql-18-pgvector_0.8.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.3-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 232.5 KiB | [postgresql-18-pgvector_0.8.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.5-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 232.3 KiB | [postgresql-18-pgvector_0.8.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 230.3 KiB | [postgresql-18-pgvector_0.8.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 229.9 KiB | [postgresql-18-pgvector_0.8.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.3-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 264.0 KiB | [postgresql-18-pgvector_0.8.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 264.0 KiB | [postgresql-18-pgvector_0.8.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 272.5 KiB | [postgresql-18-pgvector_0.8.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 262.0 KiB | [postgresql-18-pgvector_0.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 232.0 KiB | [postgresql-18-pgvector_0.8.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 231.7 KiB | [postgresql-18-pgvector_0.8.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 246.4 KiB | [postgresql-18-pgvector_0.8.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 230.0 KiB | [postgresql-18-pgvector_0.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 257.7 KiB | [postgresql-18-pgvector_0.8.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 257.7 KiB | [postgresql-18-pgvector_0.8.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 261.9 KiB | [postgresql-18-pgvector_0.8.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 255.2 KiB | [postgresql-18-pgvector_0.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 227.5 KiB | [postgresql-18-pgvector_0.8.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 227.6 KiB | [postgresql-18-pgvector_0.8.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 239.8 KiB | [postgresql-18-pgvector_0.8.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 225.4 KiB | [postgresql-18-pgvector_0.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 256.2 KiB | [postgresql-18-pgvector_0.8.5-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.5-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 255.9 KiB | [postgresql-18-pgvector_0.8.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 261.2 KiB | [postgresql-18-pgvector_0.8.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 253.6 KiB | [postgresql-18-pgvector_0.8.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgvector` | `0.8.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 227.2 KiB | [postgresql-18-pgvector_0.8.5-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.5-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 226.6 KiB | [postgresql-18-pgvector_0.8.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 238.8 KiB | [postgresql-18-pgvector_0.8.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvector/postgresql-18-pgvector_0.8.4-1PIGSTY~resolute_arm64.deb) |
| `postgresql-18-pgvector` | `0.8.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 224.4 KiB | [postgresql-18-pgvector_0.8.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-18-pgvector_0.8.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvector_17` | `0.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.5 KiB | [pgvector_17-0.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgvector_17-0.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_17` | `0.8.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 114.9 KiB | [pgvector_17-0.8.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvector_17-0.8.4-1PIGSTY.el8.x86_64.rpm) |
| `pgvector_17` | `0.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.2 KiB | [pgvector_17-0.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgvector_17-0.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_17` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 108.1 KiB | [pgvector_17-0.8.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgvector_17-0.8.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_17` | `0.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.4 KiB | [pgvector_17-0.8.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgvector_17-0.8.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_17` | `0.8.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.8 KiB | [pgvector_17-0.8.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgvector_17-0.8.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_17` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.9 KiB | [pgvector_17-0.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgvector_17-0.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_17` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.6 KiB | [pgvector_17-0.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgvector_17-0.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_17` | `0.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.9 KiB | [pgvector_17-0.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgvector_17-0.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_17` | `0.8.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.5 KiB | [pgvector_17-0.8.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvector_17-0.8.4-1PIGSTY.el8.aarch64.rpm) |
| `pgvector_17` | `0.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.6 KiB | [pgvector_17-0.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgvector_17-0.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_17` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.9 KiB | [pgvector_17-0.8.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgvector_17-0.8.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_17` | `0.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.1 KiB | [pgvector_17-0.8.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgvector_17-0.8.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_17` | `0.8.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.5 KiB | [pgvector_17-0.8.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgvector_17-0.8.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_17` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.6 KiB | [pgvector_17-0.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgvector_17-0.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_17` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.5 KiB | [pgvector_17-0.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgvector_17-0.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_17` | `0.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.8 KiB | [pgvector_17-0.8.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgvector_17-0.8.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_17` | `0.8.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 107.9 KiB | [pgvector_17-0.8.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvector_17-0.8.4-1PIGSTY.el9.x86_64.rpm) |
| `pgvector_17` | `0.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.5 KiB | [pgvector_17-0.8.4-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgvector_17-0.8.4-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_17` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.9 KiB | [pgvector_17-0.8.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgvector_17-0.8.3-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_17` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.6 KiB | [pgvector_17-0.8.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgvector_17-0.8.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_17` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.5 KiB | [pgvector_17-0.8.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgvector_17-0.8.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pgvector_17` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [pgvector_17-0.8.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgvector_17-0.8.2-1PGDG.rhel9.6.x86_64.rpm) |
| `pgvector_17` | `0.8.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.4 KiB | [pgvector_17-0.8.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgvector_17-0.8.1-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_17` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 107.6 KiB | [pgvector_17-0.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgvector_17-0.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_17` | `0.7.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.7 KiB | [pgvector_17-0.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgvector_17-0.7.4-1PGDG.rhel9.x86_64.rpm) |
| `pgvector_17` | `0.8.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 95.6 KiB | [pgvector_17-0.8.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgvector_17-0.8.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_17` | `0.8.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 98.0 KiB | [pgvector_17-0.8.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvector_17-0.8.4-1PIGSTY.el9.aarch64.rpm) |
| `pgvector_17` | `0.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 95.3 KiB | [pgvector_17-0.8.4-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgvector_17-0.8.4-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_17` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.7 KiB | [pgvector_17-0.8.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgvector_17-0.8.3-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_17` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.7 KiB | [pgvector_17-0.8.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgvector_17-0.8.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_17` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.7 KiB | [pgvector_17-0.8.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgvector_17-0.8.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pgvector_17` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.8 KiB | [pgvector_17-0.8.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgvector_17-0.8.2-1PGDG.rhel9.6.aarch64.rpm) |
| `pgvector_17` | `0.8.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.1 KiB | [pgvector_17-0.8.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgvector_17-0.8.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_17` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.8 KiB | [pgvector_17-0.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgvector_17-0.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_17` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 89.6 KiB | [pgvector_17-0.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgvector_17-0.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_17` | `0.8.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 106.2 KiB | [pgvector_17-0.8.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgvector_17-0.8.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_17` | `0.8.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 109.1 KiB | [pgvector_17-0.8.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvector_17-0.8.4-1PIGSTY.el10.x86_64.rpm) |
| `pgvector_17` | `0.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 106.0 KiB | [pgvector_17-0.8.4-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgvector_17-0.8.4-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_17` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.1 KiB | [pgvector_17-0.8.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgvector_17-0.8.3-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_17` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.0 KiB | [pgvector_17-0.8.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgvector_17-0.8.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_17` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.9 KiB | [pgvector_17-0.8.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgvector_17-0.8.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pgvector_17` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.4 KiB | [pgvector_17-0.8.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgvector_17-0.8.2-1PGDG.rhel10.0.x86_64.rpm) |
| `pgvector_17` | `0.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.1 KiB | [pgvector_17-0.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgvector_17-0.8.1-1PGDG.rhel10.x86_64.rpm) |
| `pgvector_17` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.7 KiB | [pgvector_17-0.8.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgvector_17-0.8.0-2PGDG.rhel10.x86_64.rpm) |
| `pgvector_17` | `0.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 97.9 KiB | [pgvector_17-0.8.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgvector_17-0.8.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_17` | `0.8.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 100.0 KiB | [pgvector_17-0.8.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvector_17-0.8.4-1PIGSTY.el10.aarch64.rpm) |
| `pgvector_17` | `0.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 97.5 KiB | [pgvector_17-0.8.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgvector_17-0.8.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_17` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.9 KiB | [pgvector_17-0.8.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgvector_17-0.8.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_17` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.9 KiB | [pgvector_17-0.8.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgvector_17-0.8.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_17` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.9 KiB | [pgvector_17-0.8.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgvector_17-0.8.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pgvector_17` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.9 KiB | [pgvector_17-0.8.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgvector_17-0.8.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pgvector_17` | `0.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.8 KiB | [pgvector_17-0.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgvector_17-0.8.1-1PGDG.rhel10.aarch64.rpm) |
| `pgvector_17` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.2 KiB | [pgvector_17-0.8.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgvector_17-0.8.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgvector` | `0.8.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 261.4 KiB | [postgresql-17-pgvector_0.8.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.5-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 261.0 KiB | [postgresql-17-pgvector_0.8.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 253.7 KiB | [postgresql-17-pgvector_0.8.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 258.4 KiB | [postgresql-17-pgvector_0.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 231.1 KiB | [postgresql-17-pgvector_0.8.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.5-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 230.7 KiB | [postgresql-17-pgvector_0.8.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 229.0 KiB | [postgresql-17-pgvector_0.8.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 228.5 KiB | [postgresql-17-pgvector_0.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 261.6 KiB | [postgresql-17-pgvector_0.8.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.5-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 261.3 KiB | [postgresql-17-pgvector_0.8.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 254.4 KiB | [postgresql-17-pgvector_0.8.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 258.8 KiB | [postgresql-17-pgvector_0.8.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.3-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 232.3 KiB | [postgresql-17-pgvector_0.8.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.5-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 231.6 KiB | [postgresql-17-pgvector_0.8.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 229.9 KiB | [postgresql-17-pgvector_0.8.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 229.6 KiB | [postgresql-17-pgvector_0.8.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.3-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 304.2 KiB | [postgresql-17-pgvector_0.8.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 302.6 KiB | [postgresql-17-pgvector_0.8.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 311.4 KiB | [postgresql-17-pgvector_0.8.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 301.4 KiB | [postgresql-17-pgvector_0.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 270.7 KiB | [postgresql-17-pgvector_0.8.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 270.1 KiB | [postgresql-17-pgvector_0.8.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 285.5 KiB | [postgresql-17-pgvector_0.8.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 268.0 KiB | [postgresql-17-pgvector_0.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 257.7 KiB | [postgresql-17-pgvector_0.8.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 257.2 KiB | [postgresql-17-pgvector_0.8.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 261.4 KiB | [postgresql-17-pgvector_0.8.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 254.8 KiB | [postgresql-17-pgvector_0.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 227.6 KiB | [postgresql-17-pgvector_0.8.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 227.4 KiB | [postgresql-17-pgvector_0.8.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 239.4 KiB | [postgresql-17-pgvector_0.8.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 224.9 KiB | [postgresql-17-pgvector_0.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 255.6 KiB | [postgresql-17-pgvector_0.8.5-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.5-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 255.5 KiB | [postgresql-17-pgvector_0.8.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 260.7 KiB | [postgresql-17-pgvector_0.8.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 253.2 KiB | [postgresql-17-pgvector_0.8.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgvector` | `0.8.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 226.7 KiB | [postgresql-17-pgvector_0.8.5-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.5-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 226.5 KiB | [postgresql-17-pgvector_0.8.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 238.7 KiB | [postgresql-17-pgvector_0.8.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvector/postgresql-17-pgvector_0.8.4-1PIGSTY~resolute_arm64.deb) |
| `postgresql-17-pgvector` | `0.8.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 224.3 KiB | [postgresql-17-pgvector_0.8.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-17-pgvector_0.8.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvector_16` | `0.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.7 KiB | [pgvector_16-0.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_16` | `0.8.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 115.2 KiB | [pgvector_16-0.8.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvector_16-0.8.4-1PIGSTY.el8.x86_64.rpm) |
| `pgvector_16` | `0.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.4 KiB | [pgvector_16-0.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_16` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 108.4 KiB | [pgvector_16-0.8.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.8.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_16` | `0.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.5 KiB | [pgvector_16-0.8.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.8.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_16` | `0.8.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.8 KiB | [pgvector_16-0.8.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.8.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.6 KiB | [pgvector_16-0.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.4 KiB | [pgvector_16-0.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.1 KiB | [pgvector_16-0.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.8 KiB | [pgvector_16-0.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.7.2-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.3 KiB | [pgvector_16-0.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 99.8 KiB | [pgvector_16-0.7.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.7.0-2PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.1 KiB | [pgvector_16-0.6.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.6.2-2PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.1 KiB | [pgvector_16-0.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.0 KiB | [pgvector_16-0.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.0 KiB | [pgvector_16-0.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.9 KiB | [pgvector_16-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.4 KiB | [pgvector_16-0.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgvector_16-0.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_16` | `0.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.8 KiB | [pgvector_16-0.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_16` | `0.8.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 106.4 KiB | [pgvector_16-0.8.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvector_16-0.8.4-1PIGSTY.el8.aarch64.rpm) |
| `pgvector_16` | `0.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.5 KiB | [pgvector_16-0.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_16` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.7 KiB | [pgvector_16-0.8.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.8.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_16` | `0.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.0 KiB | [pgvector_16-0.8.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.8.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_16` | `0.8.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.3 KiB | [pgvector_16-0.8.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.8.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.4 KiB | [pgvector_16-0.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.3 KiB | [pgvector_16-0.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.0 KiB | [pgvector_16-0.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.7 KiB | [pgvector_16-0.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.7.2-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.3 KiB | [pgvector_16-0.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 89.7 KiB | [pgvector_16-0.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 70.4 KiB | [pgvector_16-0.6.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.6.2-2PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 70.2 KiB | [pgvector_16-0.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 69.1 KiB | [pgvector_16-0.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 68.3 KiB | [pgvector_16-0.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 58.6 KiB | [pgvector_16-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 58.0 KiB | [pgvector_16-0.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgvector_16-0.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_16` | `0.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.8 KiB | [pgvector_16-0.8.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.8.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_16` | `0.8.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 107.8 KiB | [pgvector_16-0.8.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvector_16-0.8.4-1PIGSTY.el9.x86_64.rpm) |
| `pgvector_16` | `0.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.5 KiB | [pgvector_16-0.8.4-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.8.4-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_16` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [pgvector_16-0.8.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.8.3-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_16` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.5 KiB | [pgvector_16-0.8.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.8.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_16` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.5 KiB | [pgvector_16-0.8.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.8.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pgvector_16` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.6 KiB | [pgvector_16-0.8.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgvector_16-0.8.2-1PGDG.rhel9.6.x86_64.rpm) |
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
| `pgvector_16` | `0.8.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 95.4 KiB | [pgvector_16-0.8.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.8.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_16` | `0.8.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 97.7 KiB | [pgvector_16-0.8.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvector_16-0.8.4-1PIGSTY.el9.aarch64.rpm) |
| `pgvector_16` | `0.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 95.1 KiB | [pgvector_16-0.8.4-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.8.4-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_16` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.5 KiB | [pgvector_16-0.8.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.8.3-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_16` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.5 KiB | [pgvector_16-0.8.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.8.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_16` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.4 KiB | [pgvector_16-0.8.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.8.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pgvector_16` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.6 KiB | [pgvector_16-0.8.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.8.2-1PGDG.rhel9.6.aarch64.rpm) |
| `pgvector_16` | `0.8.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.9 KiB | [pgvector_16-0.8.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.8.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.5 KiB | [pgvector_16-0.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 89.6 KiB | [pgvector_16-0.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 89.2 KiB | [pgvector_16-0.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 88.6 KiB | [pgvector_16-0.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.7.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 88.4 KiB | [pgvector_16-0.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 87.6 KiB | [pgvector_16-0.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.2 KiB | [pgvector_16-0.6.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.6.2-2PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.8 KiB | [pgvector_16-0.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.6 KiB | [pgvector_16-0.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.0 KiB | [pgvector_16-0.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 59.8 KiB | [pgvector_16-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 59.6 KiB | [pgvector_16-0.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgvector_16-0.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_16` | `0.8.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 106.1 KiB | [pgvector_16-0.8.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgvector_16-0.8.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_16` | `0.8.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 109.1 KiB | [pgvector_16-0.8.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvector_16-0.8.4-1PIGSTY.el10.x86_64.rpm) |
| `pgvector_16` | `0.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.9 KiB | [pgvector_16-0.8.4-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgvector_16-0.8.4-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_16` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.0 KiB | [pgvector_16-0.8.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgvector_16-0.8.3-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_16` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.6 KiB | [pgvector_16-0.8.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgvector_16-0.8.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_16` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.6 KiB | [pgvector_16-0.8.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgvector_16-0.8.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pgvector_16` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.3 KiB | [pgvector_16-0.8.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgvector_16-0.8.2-1PGDG.rhel10.0.x86_64.rpm) |
| `pgvector_16` | `0.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.7 KiB | [pgvector_16-0.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgvector_16-0.8.1-1PGDG.rhel10.x86_64.rpm) |
| `pgvector_16` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.2 KiB | [pgvector_16-0.8.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgvector_16-0.8.0-2PGDG.rhel10.x86_64.rpm) |
| `pgvector_16` | `0.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 97.5 KiB | [pgvector_16-0.8.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgvector_16-0.8.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_16` | `0.8.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 99.9 KiB | [pgvector_16-0.8.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvector_16-0.8.4-1PIGSTY.el10.aarch64.rpm) |
| `pgvector_16` | `0.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 97.3 KiB | [pgvector_16-0.8.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgvector_16-0.8.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_16` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.6 KiB | [pgvector_16-0.8.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgvector_16-0.8.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_16` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.5 KiB | [pgvector_16-0.8.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgvector_16-0.8.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_16` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.5 KiB | [pgvector_16-0.8.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgvector_16-0.8.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pgvector_16` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.5 KiB | [pgvector_16-0.8.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgvector_16-0.8.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pgvector_16` | `0.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.3 KiB | [pgvector_16-0.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgvector_16-0.8.1-1PGDG.rhel10.aarch64.rpm) |
| `pgvector_16` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 95.8 KiB | [pgvector_16-0.8.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgvector_16-0.8.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgvector` | `0.8.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 261.2 KiB | [postgresql-16-pgvector_0.8.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.5-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 260.8 KiB | [postgresql-16-pgvector_0.8.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 254.2 KiB | [postgresql-16-pgvector_0.8.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 258.6 KiB | [postgresql-16-pgvector_0.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 230.8 KiB | [postgresql-16-pgvector_0.8.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.5-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 230.3 KiB | [postgresql-16-pgvector_0.8.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 228.5 KiB | [postgresql-16-pgvector_0.8.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 228.3 KiB | [postgresql-16-pgvector_0.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 261.7 KiB | [postgresql-16-pgvector_0.8.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.5-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 261.4 KiB | [postgresql-16-pgvector_0.8.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 254.4 KiB | [postgresql-16-pgvector_0.8.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 258.8 KiB | [postgresql-16-pgvector_0.8.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.3-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 231.9 KiB | [postgresql-16-pgvector_0.8.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.5-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 231.6 KiB | [postgresql-16-pgvector_0.8.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 230.2 KiB | [postgresql-16-pgvector_0.8.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 229.3 KiB | [postgresql-16-pgvector_0.8.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.3-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 294.6 KiB | [postgresql-16-pgvector_0.8.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 293.9 KiB | [postgresql-16-pgvector_0.8.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 302.8 KiB | [postgresql-16-pgvector_0.8.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 292.1 KiB | [postgresql-16-pgvector_0.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 261.6 KiB | [postgresql-16-pgvector_0.8.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 261.0 KiB | [postgresql-16-pgvector_0.8.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 276.3 KiB | [postgresql-16-pgvector_0.8.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 258.8 KiB | [postgresql-16-pgvector_0.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 256.6 KiB | [postgresql-16-pgvector_0.8.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 256.9 KiB | [postgresql-16-pgvector_0.8.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 261.4 KiB | [postgresql-16-pgvector_0.8.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 254.3 KiB | [postgresql-16-pgvector_0.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 227.0 KiB | [postgresql-16-pgvector_0.8.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 226.7 KiB | [postgresql-16-pgvector_0.8.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 239.0 KiB | [postgresql-16-pgvector_0.8.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 224.6 KiB | [postgresql-16-pgvector_0.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 255.4 KiB | [postgresql-16-pgvector_0.8.5-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.5-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 255.2 KiB | [postgresql-16-pgvector_0.8.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 260.5 KiB | [postgresql-16-pgvector_0.8.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 252.9 KiB | [postgresql-16-pgvector_0.8.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgvector` | `0.8.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 226.1 KiB | [postgresql-16-pgvector_0.8.5-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.5-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 226.2 KiB | [postgresql-16-pgvector_0.8.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 238.2 KiB | [postgresql-16-pgvector_0.8.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvector/postgresql-16-pgvector_0.8.4-1PIGSTY~resolute_arm64.deb) |
| `postgresql-16-pgvector` | `0.8.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 224.1 KiB | [postgresql-16-pgvector_0.8.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-16-pgvector_0.8.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvector_15` | `0.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 110.6 KiB | [pgvector_15-0.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_15` | `0.8.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 116.0 KiB | [pgvector_15-0.8.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvector_15-0.8.4-1PIGSTY.el8.x86_64.rpm) |
| `pgvector_15` | `0.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 110.3 KiB | [pgvector_15-0.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_15` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.3 KiB | [pgvector_15-0.8.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.8.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_15` | `0.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 108.4 KiB | [pgvector_15-0.8.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.8.2-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_15` | `0.8.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.8 KiB | [pgvector_15-0.8.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.8.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.7 KiB | [pgvector_15-0.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.4 KiB | [pgvector_15-0.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.0 KiB | [pgvector_15-0.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.7 KiB | [pgvector_15-0.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.7.2-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.3 KiB | [pgvector_15-0.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.8 KiB | [pgvector_15-0.7.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.7.0-2PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.9 KiB | [pgvector_15-0.6.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.6.2-2PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.9 KiB | [pgvector_15-0.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.8 KiB | [pgvector_15-0.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.1 KiB | [pgvector_15-0.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 64.1 KiB | [pgvector_15-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.5 KiB | [pgvector_15-0.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.9 KiB | [pgvector_15-0.4.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.4.4-1.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pgvector_15-0.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgvector_15-0.4.1-1.rhel8.x86_64.rpm) |
| `pgvector_15` | `0.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.7 KiB | [pgvector_15-0.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_15` | `0.8.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 107.1 KiB | [pgvector_15-0.8.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvector_15-0.8.4-1PIGSTY.el8.aarch64.rpm) |
| `pgvector_15` | `0.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.3 KiB | [pgvector_15-0.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_15` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.6 KiB | [pgvector_15-0.8.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.8.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_15` | `0.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.8 KiB | [pgvector_15-0.8.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.8.2-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_15` | `0.8.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.2 KiB | [pgvector_15-0.8.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.8.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.3 KiB | [pgvector_15-0.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 92.1 KiB | [pgvector_15-0.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.9 KiB | [pgvector_15-0.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.4 KiB | [pgvector_15-0.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.7.2-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.0 KiB | [pgvector_15-0.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.4 KiB | [pgvector_15-0.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.3 KiB | [pgvector_15-0.6.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.6.2-2PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.0 KiB | [pgvector_15-0.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 69.9 KiB | [pgvector_15-0.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 69.3 KiB | [pgvector_15-0.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.7 KiB | [pgvector_15-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.0 KiB | [pgvector_15-0.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.5 KiB | [pgvector_15-0.4.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.4.4-1.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.6 KiB | [pgvector_15-0.4.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgvector_15-0.4.1-1.rhel8.aarch64.rpm) |
| `pgvector_15` | `0.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.0 KiB | [pgvector_15-0.8.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.8.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_15` | `0.8.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 112.4 KiB | [pgvector_15-0.8.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvector_15-0.8.4-1PIGSTY.el9.x86_64.rpm) |
| `pgvector_15` | `0.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.7 KiB | [pgvector_15-0.8.4-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.8.4-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_15` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.2 KiB | [pgvector_15-0.8.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.8.3-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_15` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.2 KiB | [pgvector_15-0.8.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.8.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_15` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.1 KiB | [pgvector_15-0.8.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.8.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pgvector_15` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.2 KiB | [pgvector_15-0.8.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgvector_15-0.8.2-1PGDG.rhel9.6.x86_64.rpm) |
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
| `pgvector_15` | `0.8.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 100.0 KiB | [pgvector_15-0.8.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.8.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_15` | `0.8.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 102.4 KiB | [pgvector_15-0.8.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvector_15-0.8.4-1PIGSTY.el9.aarch64.rpm) |
| `pgvector_15` | `0.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.8 KiB | [pgvector_15-0.8.4-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.8.4-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_15` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.6 KiB | [pgvector_15-0.8.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.8.3-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_15` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.5 KiB | [pgvector_15-0.8.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.8.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_15` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.5 KiB | [pgvector_15-0.8.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.8.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pgvector_15` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.6 KiB | [pgvector_15-0.8.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.8.2-1PGDG.rhel9.6.aarch64.rpm) |
| `pgvector_15` | `0.8.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.8 KiB | [pgvector_15-0.8.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.8.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.5 KiB | [pgvector_15-0.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.2 KiB | [pgvector_15-0.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.8 KiB | [pgvector_15-0.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.5 KiB | [pgvector_15-0.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.7.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.1 KiB | [pgvector_15-0.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 92.3 KiB | [pgvector_15-0.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.5 KiB | [pgvector_15-0.6.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.6.2-2PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.1 KiB | [pgvector_15-0.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.0 KiB | [pgvector_15-0.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.3 KiB | [pgvector_15-0.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.5 KiB | [pgvector_15-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.9 KiB | [pgvector_15-0.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.4 KiB | [pgvector_15-0.4.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.4.4-1.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.6 KiB | [pgvector_15-0.4.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgvector_15-0.4.1-1.rhel9.aarch64.rpm) |
| `pgvector_15` | `0.8.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 110.2 KiB | [pgvector_15-0.8.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgvector_15-0.8.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_15` | `0.8.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 113.3 KiB | [pgvector_15-0.8.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvector_15-0.8.4-1PIGSTY.el10.x86_64.rpm) |
| `pgvector_15` | `0.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 109.9 KiB | [pgvector_15-0.8.4-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgvector_15-0.8.4-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_15` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 109.1 KiB | [pgvector_15-0.8.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgvector_15-0.8.3-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_15` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 108.8 KiB | [pgvector_15-0.8.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgvector_15-0.8.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_15` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 108.8 KiB | [pgvector_15-0.8.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgvector_15-0.8.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pgvector_15` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 109.3 KiB | [pgvector_15-0.8.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgvector_15-0.8.2-1PGDG.rhel10.0.x86_64.rpm) |
| `pgvector_15` | `0.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 108.7 KiB | [pgvector_15-0.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgvector_15-0.8.1-1PGDG.rhel10.x86_64.rpm) |
| `pgvector_15` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 108.5 KiB | [pgvector_15-0.8.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgvector_15-0.8.0-2PGDG.rhel10.x86_64.rpm) |
| `pgvector_15` | `0.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 102.4 KiB | [pgvector_15-0.8.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgvector_15-0.8.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_15` | `0.8.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 104.4 KiB | [pgvector_15-0.8.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvector_15-0.8.4-1PIGSTY.el10.aarch64.rpm) |
| `pgvector_15` | `0.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 102.1 KiB | [pgvector_15-0.8.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgvector_15-0.8.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_15` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.6 KiB | [pgvector_15-0.8.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgvector_15-0.8.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_15` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.6 KiB | [pgvector_15-0.8.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgvector_15-0.8.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_15` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.6 KiB | [pgvector_15-0.8.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgvector_15-0.8.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pgvector_15` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.6 KiB | [pgvector_15-0.8.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgvector_15-0.8.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pgvector_15` | `0.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.0 KiB | [pgvector_15-0.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgvector_15-0.8.1-1PGDG.rhel10.aarch64.rpm) |
| `pgvector_15` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 100.8 KiB | [pgvector_15-0.8.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgvector_15-0.8.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgvector` | `0.8.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 262.1 KiB | [postgresql-15-pgvector_0.8.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.5-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 261.3 KiB | [postgresql-15-pgvector_0.8.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 254.8 KiB | [postgresql-15-pgvector_0.8.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 259.4 KiB | [postgresql-15-pgvector_0.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 232.7 KiB | [postgresql-15-pgvector_0.8.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.5-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 232.4 KiB | [postgresql-15-pgvector_0.8.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 229.9 KiB | [postgresql-15-pgvector_0.8.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 230.0 KiB | [postgresql-15-pgvector_0.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 262.9 KiB | [postgresql-15-pgvector_0.8.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.5-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 262.7 KiB | [postgresql-15-pgvector_0.8.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 255.8 KiB | [postgresql-15-pgvector_0.8.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 260.5 KiB | [postgresql-15-pgvector_0.8.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.3-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 233.5 KiB | [postgresql-15-pgvector_0.8.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.5-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 233.6 KiB | [postgresql-15-pgvector_0.8.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 231.2 KiB | [postgresql-15-pgvector_0.8.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 231.3 KiB | [postgresql-15-pgvector_0.8.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.3-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 298.2 KiB | [postgresql-15-pgvector_0.8.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 296.7 KiB | [postgresql-15-pgvector_0.8.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 306.3 KiB | [postgresql-15-pgvector_0.8.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 294.1 KiB | [postgresql-15-pgvector_0.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 265.1 KiB | [postgresql-15-pgvector_0.8.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 264.6 KiB | [postgresql-15-pgvector_0.8.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 280.7 KiB | [postgresql-15-pgvector_0.8.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 262.7 KiB | [postgresql-15-pgvector_0.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 259.5 KiB | [postgresql-15-pgvector_0.8.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 259.2 KiB | [postgresql-15-pgvector_0.8.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 264.4 KiB | [postgresql-15-pgvector_0.8.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 256.9 KiB | [postgresql-15-pgvector_0.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.2 KiB | [postgresql-15-pgvector_0.8.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.5 KiB | [postgresql-15-pgvector_0.8.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 242.8 KiB | [postgresql-15-pgvector_0.8.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 228.3 KiB | [postgresql-15-pgvector_0.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 258.0 KiB | [postgresql-15-pgvector_0.8.5-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.5-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 257.8 KiB | [postgresql-15-pgvector_0.8.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 263.9 KiB | [postgresql-15-pgvector_0.8.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 255.7 KiB | [postgresql-15-pgvector_0.8.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgvector` | `0.8.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 229.8 KiB | [postgresql-15-pgvector_0.8.5-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.5-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 229.4 KiB | [postgresql-15-pgvector_0.8.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 241.9 KiB | [postgresql-15-pgvector_0.8.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvector/postgresql-15-pgvector_0.8.4-1PIGSTY~resolute_arm64.deb) |
| `postgresql-15-pgvector` | `0.8.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 227.4 KiB | [postgresql-15-pgvector_0.8.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-15-pgvector_0.8.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvector_14` | `0.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 110.5 KiB | [pgvector_14-0.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_14` | `0.8.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 115.9 KiB | [pgvector_14-0.8.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvector_14-0.8.4-1PIGSTY.el8.x86_64.rpm) |
| `pgvector_14` | `0.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 110.2 KiB | [pgvector_14-0.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_14` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.2 KiB | [pgvector_14-0.8.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.8.3-1PGDG.rhel8.10.x86_64.rpm) |
| `pgvector_14` | `0.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 108.3 KiB | [pgvector_14-0.8.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgvector_14-0.8.2-1PGDG.rhel8.10.x86_64.rpm) |
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
| `pgvector_14` | `0.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.6 KiB | [pgvector_14-0.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_14` | `0.8.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 107.1 KiB | [pgvector_14-0.8.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvector_14-0.8.4-1PIGSTY.el8.aarch64.rpm) |
| `pgvector_14` | `0.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.4 KiB | [pgvector_14-0.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_14` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.5 KiB | [pgvector_14-0.8.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.8.3-1PGDG.rhel8.10.aarch64.rpm) |
| `pgvector_14` | `0.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.7 KiB | [pgvector_14-0.8.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgvector_14-0.8.2-1PGDG.rhel8.10.aarch64.rpm) |
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
| `pgvector_14` | `0.8.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.9 KiB | [pgvector_14-0.8.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.8.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_14` | `0.8.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 112.3 KiB | [pgvector_14-0.8.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvector_14-0.8.4-1PIGSTY.el9.x86_64.rpm) |
| `pgvector_14` | `0.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.6 KiB | [pgvector_14-0.8.4-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.8.4-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_14` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.9 KiB | [pgvector_14-0.8.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.8.3-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_14` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.8 KiB | [pgvector_14-0.8.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.8.2-1PGDG.rhel9.8.x86_64.rpm) |
| `pgvector_14` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.9 KiB | [pgvector_14-0.8.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.8.2-1PGDG.rhel9.7.x86_64.rpm) |
| `pgvector_14` | `0.8.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 113.0 KiB | [pgvector_14-0.8.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgvector_14-0.8.2-1PGDG.rhel9.6.x86_64.rpm) |
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
| `pgvector_14` | `0.8.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 100.0 KiB | [pgvector_14-0.8.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.8.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_14` | `0.8.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 102.4 KiB | [pgvector_14-0.8.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvector_14-0.8.4-1PIGSTY.el9.aarch64.rpm) |
| `pgvector_14` | `0.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 100.0 KiB | [pgvector_14-0.8.4-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.8.4-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_14` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.4 KiB | [pgvector_14-0.8.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.8.3-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_14` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.1 KiB | [pgvector_14-0.8.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.8.2-1PGDG.rhel9.8.aarch64.rpm) |
| `pgvector_14` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.0 KiB | [pgvector_14-0.8.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.8.2-1PGDG.rhel9.7.aarch64.rpm) |
| `pgvector_14` | `0.8.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.2 KiB | [pgvector_14-0.8.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.8.2-1PGDG.rhel9.6.aarch64.rpm) |
| `pgvector_14` | `0.8.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.3 KiB | [pgvector_14-0.8.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.8.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 98.1 KiB | [pgvector_14-0.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.2 KiB | [pgvector_14-0.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.7.4-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.8 KiB | [pgvector_14-0.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.7.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.5 KiB | [pgvector_14-0.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.7.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.1 KiB | [pgvector_14-0.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 92.4 KiB | [pgvector_14-0.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.4 KiB | [pgvector_14-0.6.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.6.2-2PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.1 KiB | [pgvector_14-0.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.0 KiB | [pgvector_14-0.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.3 KiB | [pgvector_14-0.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.5 KiB | [pgvector_14-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.9 KiB | [pgvector_14-0.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.3 KiB | [pgvector_14-0.4.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.4.4-1.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.5 KiB | [pgvector_14-0.4.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgvector_14-0.4.1-1.rhel9.aarch64.rpm) |
| `pgvector_14` | `0.8.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 110.2 KiB | [pgvector_14-0.8.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgvector_14-0.8.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_14` | `0.8.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 113.1 KiB | [pgvector_14-0.8.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvector_14-0.8.4-1PIGSTY.el10.x86_64.rpm) |
| `pgvector_14` | `0.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 109.9 KiB | [pgvector_14-0.8.4-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgvector_14-0.8.4-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_14` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 109.3 KiB | [pgvector_14-0.8.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgvector_14-0.8.3-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_14` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 108.9 KiB | [pgvector_14-0.8.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgvector_14-0.8.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pgvector_14` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 109.0 KiB | [pgvector_14-0.8.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgvector_14-0.8.2-1PGDG.rhel10.1.x86_64.rpm) |
| `pgvector_14` | `0.8.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 109.6 KiB | [pgvector_14-0.8.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgvector_14-0.8.2-1PGDG.rhel10.0.x86_64.rpm) |
| `pgvector_14` | `0.8.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 108.9 KiB | [pgvector_14-0.8.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgvector_14-0.8.1-1PGDG.rhel10.x86_64.rpm) |
| `pgvector_14` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 108.6 KiB | [pgvector_14-0.8.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgvector_14-0.8.0-2PGDG.rhel10.x86_64.rpm) |
| `pgvector_14` | `0.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 102.6 KiB | [pgvector_14-0.8.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgvector_14-0.8.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_14` | `0.8.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 104.4 KiB | [pgvector_14-0.8.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvector_14-0.8.4-1PIGSTY.el10.aarch64.rpm) |
| `pgvector_14` | `0.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 102.2 KiB | [pgvector_14-0.8.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgvector_14-0.8.4-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_14` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.6 KiB | [pgvector_14-0.8.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgvector_14-0.8.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_14` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.5 KiB | [pgvector_14-0.8.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgvector_14-0.8.2-1PGDG.rhel10.2.aarch64.rpm) |
| `pgvector_14` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.3 KiB | [pgvector_14-0.8.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgvector_14-0.8.2-1PGDG.rhel10.1.aarch64.rpm) |
| `pgvector_14` | `0.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.5 KiB | [pgvector_14-0.8.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgvector_14-0.8.2-1PGDG.rhel10.0.aarch64.rpm) |
| `pgvector_14` | `0.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.2 KiB | [pgvector_14-0.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgvector_14-0.8.1-1PGDG.rhel10.aarch64.rpm) |
| `pgvector_14` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 100.7 KiB | [pgvector_14-0.8.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgvector_14-0.8.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgvector` | `0.8.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 262.4 KiB | [postgresql-14-pgvector_0.8.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.5-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 261.4 KiB | [postgresql-14-pgvector_0.8.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 255.2 KiB | [postgresql-14-pgvector_0.8.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 259.1 KiB | [postgresql-14-pgvector_0.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 232.1 KiB | [postgresql-14-pgvector_0.8.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.5-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 232.2 KiB | [postgresql-14-pgvector_0.8.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 230.1 KiB | [postgresql-14-pgvector_0.8.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 229.9 KiB | [postgresql-14-pgvector_0.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 262.9 KiB | [postgresql-14-pgvector_0.8.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.5-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 262.4 KiB | [postgresql-14-pgvector_0.8.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 255.1 KiB | [postgresql-14-pgvector_0.8.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 260.3 KiB | [postgresql-14-pgvector_0.8.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.3-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 233.3 KiB | [postgresql-14-pgvector_0.8.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.5-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 233.2 KiB | [postgresql-14-pgvector_0.8.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 231.3 KiB | [postgresql-14-pgvector_0.8.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 230.9 KiB | [postgresql-14-pgvector_0.8.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.3-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 295.8 KiB | [postgresql-14-pgvector_0.8.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 295.4 KiB | [postgresql-14-pgvector_0.8.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 305.1 KiB | [postgresql-14-pgvector_0.8.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 293.6 KiB | [postgresql-14-pgvector_0.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 263.5 KiB | [postgresql-14-pgvector_0.8.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 263.0 KiB | [postgresql-14-pgvector_0.8.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 279.1 KiB | [postgresql-14-pgvector_0.8.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 260.8 KiB | [postgresql-14-pgvector_0.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 259.5 KiB | [postgresql-14-pgvector_0.8.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 259.2 KiB | [postgresql-14-pgvector_0.8.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 264.3 KiB | [postgresql-14-pgvector_0.8.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 256.8 KiB | [postgresql-14-pgvector_0.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.3 KiB | [postgresql-14-pgvector_0.8.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 230.0 KiB | [postgresql-14-pgvector_0.8.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 242.6 KiB | [postgresql-14-pgvector_0.8.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 228.4 KiB | [postgresql-14-pgvector_0.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 257.8 KiB | [postgresql-14-pgvector_0.8.5-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.5-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 257.8 KiB | [postgresql-14-pgvector_0.8.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 263.8 KiB | [postgresql-14-pgvector_0.8.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 255.4 KiB | [postgresql-14-pgvector_0.8.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgvector` | `0.8.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 229.4 KiB | [postgresql-14-pgvector_0.8.5-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.5-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 229.3 KiB | [postgresql-14-pgvector_0.8.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 242.0 KiB | [postgresql-14-pgvector_0.8.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvector/postgresql-14-pgvector_0.8.4-1PIGSTY~resolute_arm64.deb) |
| `postgresql-14-pgvector` | `0.8.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 227.2 KiB | [postgresql-14-pgvector_0.8.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgvector/postgresql-14-pgvector_0.8.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgvector/pgvector" title="Repository" icon="github" subtitle="github.com/pgvector/pgvector" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgvector-0.8.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgvector;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgvector;		# install via package name, for the active PG version
pig install vector;		# install by extension name, for the current active PG version

pig install vector -v 18;   # install for PG 18
pig install vector -v 17;   # install for PG 17
pig install vector -v 16;   # install for PG 16
pig install vector -v 15;   # install for PG 15
pig install vector -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION vector;
```




## Usage

Sources:

- [pgvector v0.8.4 release](https://github.com/pgvector/pgvector/releases/tag/v0.8.4)
- [pgvector v0.8.4 README](https://github.com/pgvector/pgvector/blob/v0.8.4/README.md)
- [pgvector v0.8.4 CHANGELOG](https://github.com/pgvector/pgvector/blob/v0.8.4/CHANGELOG.md)

`pgvector` provides vector similarity search inside PostgreSQL. The extension name is `vector`, while Pigsty packages it as `pgvector`. It supports exact search, approximate nearest-neighbor search with HNSW and IVFFlat indexes, and multiple vector representations for dense, half-precision, binary, and sparse embeddings.

v0.8.4 is a maintenance release after the 0.8.x HNSW/vacuum fixes. Use it instead of older 0.8.x builds when maintaining HNSW indexes under write-heavy workloads.

### Create and Query Vectors

```sql
CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE items (
  id bigserial PRIMARY KEY,
  embedding vector(3)
);

INSERT INTO items (embedding)
VALUES ('[1,2,3]'), ('[4,5,6]');

SELECT *
FROM items
ORDER BY embedding <-> '[3,1,2]'
LIMIT 5;
```

Common distance operators:

- `<->` for L2 distance
- `<#>` for negative inner product
- `<=>` for cosine distance
- `<+>` for L1 distance
- `<~>` for Hamming distance on binary vectors
- `<%>` for Jaccard distance on binary vectors

Because PostgreSQL indexes scan in ascending order, `<#>` returns the negative inner product; multiply by `-1` when displaying the actual inner product.

### Vector Types

```sql
CREATE TABLE embeddings (
  id bigserial PRIMARY KEY,
  dense      vector(768),
  half_dense halfvec(768),
  binary_sig bit(1024),
  sparse     sparsevec(100000)
);
```

`vector` is the standard single-precision type. Use `halfvec` to reduce storage and memory pressure, `bit` for binary signatures, and `sparsevec` for high-dimensional sparse vectors.

Aggregates such as `avg()` and `sum()` can be used with vector columns:

```sql
SELECT avg(embedding) FROM items;
```

### HNSW Indexes

HNSW gives strong speed/recall tradeoffs and does not require a training step.

```sql
CREATE INDEX items_embedding_hnsw
ON items USING hnsw (embedding vector_l2_ops);

SET hnsw.ef_search = 100;

SELECT *
FROM items
ORDER BY embedding <-> '[3,1,2]'
LIMIT 10;
```

Choose the operator class that matches the distance:

```sql
CREATE INDEX ON items USING hnsw (embedding vector_ip_ops);
CREATE INDEX ON items USING hnsw (embedding vector_cosine_ops);
CREATE INDEX ON items USING hnsw (embedding vector_l1_ops);
CREATE INDEX ON embeddings USING hnsw (half_dense halfvec_l2_ops);
CREATE INDEX ON embeddings USING hnsw (sparse sparsevec_l2_ops);
CREATE INDEX ON embeddings USING hnsw (binary_sig bit_hamming_ops);
```

Useful tuning settings include `hnsw.ef_search`, `hnsw.iterative_scan`, `hnsw.max_scan_tuples`, and `hnsw.scan_mem_multiplier`.

### IVFFlat Indexes

IVFFlat requires representative data before index creation because it trains cluster lists at build time.

```sql
CREATE INDEX items_embedding_ivfflat
ON items USING ivfflat (embedding vector_l2_ops)
WITH (lists = 100);

SET ivfflat.probes = 10;

SELECT *
FROM items
ORDER BY embedding <-> '[3,1,2]'
LIMIT 10;
```

Increase `lists` for larger tables and increase `ivfflat.probes` for higher recall. For filtered queries, test whether an exact btree filter, a partial vector index, or partitioning gives better plans.

### Filtering and Hybrid Search

Normal PostgreSQL filters can be combined with vector ordering:

```sql
CREATE INDEX ON items (tenant_id);

SELECT *
FROM items
WHERE tenant_id = 42
ORDER BY embedding <=> '[0.1,0.2,0.3]'
LIMIT 20;
```

For hybrid search, combine `pgvector` with PostgreSQL full text search, trigram search, or an external ranking expression:

```sql
SELECT id,
       ts_rank_cd(text_tsv, plainto_tsquery('database')) AS text_score,
       1 - (embedding <=> '[0.1,0.2,0.3]') AS vector_score
FROM docs
WHERE text_tsv @@ plainto_tsquery('database')
ORDER BY vector_score DESC
LIMIT 20;
```

### Maintenance

```sql
VACUUM items;
REINDEX INDEX CONCURRENTLY items_embedding_hnsw;
ANALYZE items;
```

HNSW indexes can be large and expensive to build. Use `maintenance_work_mem` for builds, monitor build notices, and schedule `REINDEX` when index bloat or recall drift matters.

### Caveats

- Pigsty local metadata may lag this upstream version; this stub tracks upstream pgvector 0.8.4 while the local package row may still show an older package version until the package catalog is refreshed.
- Use the operator class that matches the query operator. A cosine index will not accelerate an L2 `ORDER BY`.
- Approximate indexes trade exact recall for speed. Validate recall with representative data and query filters.
- Build IVFFlat after loading data. If data distribution changes substantially, rebuild the index.
- Keep pgvector updated when using HNSW with heavy writes and vacuum activity; v0.8.x includes important HNSW maintenance fixes.
