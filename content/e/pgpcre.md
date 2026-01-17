---
title: "pgpcre"
linkTitle: "pgpcre"
description: "Perl Compatible Regular Expression functions"
weight: 4230
categories: ["UTIL"]
width: full
---

[**pgpcre**](https://github.com/petere/pgpcre) : Perl Compatible Regular Expression functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4230** | {{< badge content="pgpcre" link="https://github.com/petere/pgpcre" >}} | {{< ext "pgpcre" >}} | `0.20190509` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "icu_ext" >}} {{< ext "fuzzystrmatch" >}} {{< ext "pg_trgm" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.20190509` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pgpcre` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.20190509` | {{< bg "18" "pgpcre_18" "green" >}} {{< bg "17" "pgpcre_17" "green" >}} {{< bg "16" "pgpcre_16" "green" >}} {{< bg "15" "pgpcre_15" "green" >}} {{< bg "14" "pgpcre_14" "green" >}} {{< bg "13" "pgpcre_13" "green" >}} | `pgpcre_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.20190509` | {{< bg "18" "postgresql-18-pgpcre" "green" >}} {{< bg "17" "postgresql-17-pgpcre" "green" >}} {{< bg "16" "postgresql-16-pgpcre" "green" >}} {{< bg "15" "postgresql-15-pgpcre" "green" >}} {{< bg "14" "postgresql-14-pgpcre" "green" >}} {{< bg "13" "postgresql-13-pgpcre" "green" >}} | `postgresql-$v-pgpcre` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.20190509" "pgpcre_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_13 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.20190509" "pgpcre_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_13 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.20190509" "pgpcre_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.20190509" "pgpcre_18 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.20190509" "pgpcre_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.20190509" "pgpcre_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.20190509" "pgpcre_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.20190509" "pgpcre_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.20190509" "pgpcre_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.20190509" "pgpcre_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.20190509" "pgpcre_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.20190509" "pgpcre_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.20190509" "pgpcre_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.20190509" "pgpcre_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.20190509" "pgpcre_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.20190509" "pgpcre_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.20190509" "pgpcre_13 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.20190509" "postgresql-18-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-17-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-16-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-15-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-14-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-13-pgpcre : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.20190509" "postgresql-18-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-17-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-16-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-15-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-14-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-13-pgpcre : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.20190509" "postgresql-18-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-17-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-16-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-15-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-14-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-13-pgpcre : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.20190509" "postgresql-18-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-17-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-16-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-15-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-14-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-13-pgpcre : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.20190509" "postgresql-18-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-17-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-16-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-15-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-14-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-13-pgpcre : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.20190509" "postgresql-18-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-17-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-16-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-15-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-14-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-13-pgpcre : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.20190509" "postgresql-18-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-17-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-16-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-15-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-14-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-13-pgpcre : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.20190509" "postgresql-18-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-17-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-16-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-15-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-14-pgpcre : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20190509" "postgresql-13-pgpcre : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgpcre_18` | `0.20190509` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.4 KiB | [pgpcre_18-0.20190509-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgpcre_18-0.20190509-3PGDG.rhel8.x86_64.rpm) |
| `pgpcre_18` | `0.20190509` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.4 KiB | [pgpcre_18-0.20190509-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpcre_18-0.20190509-1PIGSTY.el8.x86_64.rpm) |
| `pgpcre_18` | `0.20190509` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.2 KiB | [pgpcre_18-0.20190509-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgpcre_18-0.20190509-3PGDG.rhel8.aarch64.rpm) |
| `pgpcre_18` | `0.20190509` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.5 KiB | [pgpcre_18-0.20190509-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpcre_18-0.20190509-1PIGSTY.el8.aarch64.rpm) |
| `pgpcre_18` | `0.20190509` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.6 KiB | [pgpcre_18-0.20190509-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgpcre_18-0.20190509-3PGDG.rhel9.x86_64.rpm) |
| `pgpcre_18` | `0.20190509` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.3 KiB | [pgpcre_18-0.20190509-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpcre_18-0.20190509-1PIGSTY.el9.x86_64.rpm) |
| `pgpcre_18` | `0.20190509` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.2 KiB | [pgpcre_18-0.20190509-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgpcre_18-0.20190509-3PGDG.rhel9.aarch64.rpm) |
| `pgpcre_18` | `0.20190509` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.2 KiB | [pgpcre_18-0.20190509-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpcre_18-0.20190509-1PIGSTY.el9.aarch64.rpm) |
| `pgpcre_18` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.1 KiB | [pgpcre_18-0.20190509-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgpcre_18-0.20190509-4PGDG.rhel10.x86_64.rpm) |
| `pgpcre_18` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.0 KiB | [pgpcre_18-0.20190509-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgpcre_18-0.20190509-3PGDG.rhel10.x86_64.rpm) |
| `pgpcre_18` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 18.0 KiB | [pgpcre_18-0.20190509-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgpcre_18-0.20190509-4PGDG.rhel10.aarch64.rpm) |
| `pgpcre_18` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgpcre_18-0.20190509-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgpcre_18-0.20190509-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgpcre` | `0.20190509` | [d12.x86_64](/os/d12.x86_64) | pgdg | 18.1 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg12+1_amd64.deb) |
| `postgresql-18-pgpcre` | `0.20190509` | [d12.aarch64](/os/d12.aarch64) | pgdg | 18.2 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg12+1_arm64.deb) |
| `postgresql-18-pgpcre` | `0.20190509` | [d13.x86_64](/os/d13.x86_64) | pgdg | 18.2 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg13+1_amd64.deb) |
| `postgresql-18-pgpcre` | `0.20190509` | [d13.aarch64](/os/d13.aarch64) | pgdg | 18.2 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg13+1_arm64.deb) |
| `postgresql-18-pgpcre` | `0.20190509` | [u22.x86_64](/os/u22.x86_64) | pgdg | 18.3 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgpcre` | `0.20190509` | [u22.aarch64](/os/u22.aarch64) | pgdg | 18.0 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgpcre` | `0.20190509` | [u24.x86_64](/os/u24.x86_64) | pgdg | 18.2 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgpcre` | `0.20190509` | [u24.aarch64](/os/u24.aarch64) | pgdg | 18.2 KiB | [postgresql-18-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-18-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgpcre_17` | `0.20190509` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.4 KiB | [pgpcre_17-0.20190509-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpcre_17-0.20190509-1PIGSTY.el8.x86_64.rpm) |
| `pgpcre_17` | `0.20190509` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.1 KiB | [pgpcre_17-0.20190509-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgpcre_17-0.20190509-1PGDG.rhel8.x86_64.rpm) |
| `pgpcre_17` | `0.20190509` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.5 KiB | [pgpcre_17-0.20190509-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpcre_17-0.20190509-1PIGSTY.el8.aarch64.rpm) |
| `pgpcre_17` | `0.20190509` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.0 KiB | [pgpcre_17-0.20190509-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgpcre_17-0.20190509-1PGDG.rhel8.aarch64.rpm) |
| `pgpcre_17` | `0.20190509` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.3 KiB | [pgpcre_17-0.20190509-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpcre_17-0.20190509-1PIGSTY.el9.x86_64.rpm) |
| `pgpcre_17` | `0.20190509` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pgpcre_17-0.20190509-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgpcre_17-0.20190509-1PGDG.rhel9.x86_64.rpm) |
| `pgpcre_17` | `0.20190509` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.2 KiB | [pgpcre_17-0.20190509-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpcre_17-0.20190509-1PIGSTY.el9.aarch64.rpm) |
| `pgpcre_17` | `0.20190509` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.1 KiB | [pgpcre_17-0.20190509-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgpcre_17-0.20190509-1PGDG.rhel9.aarch64.rpm) |
| `pgpcre_17` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.1 KiB | [pgpcre_17-0.20190509-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgpcre_17-0.20190509-4PGDG.rhel10.x86_64.rpm) |
| `pgpcre_17` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.0 KiB | [pgpcre_17-0.20190509-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgpcre_17-0.20190509-3PGDG.rhel10.x86_64.rpm) |
| `pgpcre_17` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.0 KiB | [pgpcre_17-0.20190509-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgpcre_17-0.20190509-2PGDG.rhel10.x86_64.rpm) |
| `pgpcre_17` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 18.0 KiB | [pgpcre_17-0.20190509-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgpcre_17-0.20190509-4PGDG.rhel10.aarch64.rpm) |
| `pgpcre_17` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgpcre_17-0.20190509-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgpcre_17-0.20190509-3PGDG.rhel10.aarch64.rpm) |
| `pgpcre_17` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pgpcre_17-0.20190509-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgpcre_17-0.20190509-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgpcre` | `0.20190509` | [d12.x86_64](/os/d12.x86_64) | pgdg | 18.0 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg12+1_amd64.deb) |
| `postgresql-17-pgpcre` | `0.20190509` | [d12.aarch64](/os/d12.aarch64) | pgdg | 18.1 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg12+1_arm64.deb) |
| `postgresql-17-pgpcre` | `0.20190509` | [d13.x86_64](/os/d13.x86_64) | pgdg | 18.2 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg13+1_amd64.deb) |
| `postgresql-17-pgpcre` | `0.20190509` | [d13.aarch64](/os/d13.aarch64) | pgdg | 18.2 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg13+1_arm64.deb) |
| `postgresql-17-pgpcre` | `0.20190509` | [u22.x86_64](/os/u22.x86_64) | pgdg | 18.9 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgpcre` | `0.20190509` | [u22.aarch64](/os/u22.aarch64) | pgdg | 18.6 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgpcre` | `0.20190509` | [u24.x86_64](/os/u24.x86_64) | pgdg | 18.1 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgpcre` | `0.20190509` | [u24.aarch64](/os/u24.aarch64) | pgdg | 18.2 KiB | [postgresql-17-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-17-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgpcre_16` | `0.20190509` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.4 KiB | [pgpcre_16-0.20190509-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpcre_16-0.20190509-1PIGSTY.el8.x86_64.rpm) |
| `pgpcre_16` | `0.20190509` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.1 KiB | [pgpcre_16-0.20190509-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgpcre_16-0.20190509-1PGDG.rhel8.x86_64.rpm) |
| `pgpcre_16` | `0.20190509` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.5 KiB | [pgpcre_16-0.20190509-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpcre_16-0.20190509-1PIGSTY.el8.aarch64.rpm) |
| `pgpcre_16` | `0.20190509` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.0 KiB | [pgpcre_16-0.20190509-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgpcre_16-0.20190509-1PGDG.rhel8.aarch64.rpm) |
| `pgpcre_16` | `0.20190509` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.3 KiB | [pgpcre_16-0.20190509-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpcre_16-0.20190509-1PIGSTY.el9.x86_64.rpm) |
| `pgpcre_16` | `0.20190509` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pgpcre_16-0.20190509-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgpcre_16-0.20190509-1PGDG.rhel9.x86_64.rpm) |
| `pgpcre_16` | `0.20190509` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.2 KiB | [pgpcre_16-0.20190509-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpcre_16-0.20190509-1PIGSTY.el9.aarch64.rpm) |
| `pgpcre_16` | `0.20190509` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.1 KiB | [pgpcre_16-0.20190509-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgpcre_16-0.20190509-1PGDG.rhel9.aarch64.rpm) |
| `pgpcre_16` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.1 KiB | [pgpcre_16-0.20190509-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgpcre_16-0.20190509-4PGDG.rhel10.x86_64.rpm) |
| `pgpcre_16` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgpcre_16-0.20190509-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgpcre_16-0.20190509-3PGDG.rhel10.x86_64.rpm) |
| `pgpcre_16` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgpcre_16-0.20190509-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgpcre_16-0.20190509-2PGDG.rhel10.x86_64.rpm) |
| `pgpcre_16` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 18.0 KiB | [pgpcre_16-0.20190509-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgpcre_16-0.20190509-4PGDG.rhel10.aarch64.rpm) |
| `pgpcre_16` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgpcre_16-0.20190509-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgpcre_16-0.20190509-3PGDG.rhel10.aarch64.rpm) |
| `pgpcre_16` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pgpcre_16-0.20190509-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgpcre_16-0.20190509-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgpcre` | `0.20190509` | [d12.x86_64](/os/d12.x86_64) | pgdg | 18.0 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg12+1_amd64.deb) |
| `postgresql-16-pgpcre` | `0.20190509` | [d12.aarch64](/os/d12.aarch64) | pgdg | 18.1 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg12+1_arm64.deb) |
| `postgresql-16-pgpcre` | `0.20190509` | [d13.x86_64](/os/d13.x86_64) | pgdg | 18.2 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg13+1_amd64.deb) |
| `postgresql-16-pgpcre` | `0.20190509` | [d13.aarch64](/os/d13.aarch64) | pgdg | 18.2 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg13+1_arm64.deb) |
| `postgresql-16-pgpcre` | `0.20190509` | [u22.x86_64](/os/u22.x86_64) | pgdg | 18.9 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgpcre` | `0.20190509` | [u22.aarch64](/os/u22.aarch64) | pgdg | 18.6 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgpcre` | `0.20190509` | [u24.x86_64](/os/u24.x86_64) | pgdg | 18.1 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgpcre` | `0.20190509` | [u24.aarch64](/os/u24.aarch64) | pgdg | 18.2 KiB | [postgresql-16-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-16-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgpcre_15` | `0.20190509` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.4 KiB | [pgpcre_15-0.20190509-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpcre_15-0.20190509-1PIGSTY.el8.x86_64.rpm) |
| `pgpcre_15` | `0.20190509` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.1 KiB | [pgpcre_15-0.20190509-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgpcre_15-0.20190509-1PGDG.rhel8.x86_64.rpm) |
| `pgpcre_15` | `0.20190509` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.5 KiB | [pgpcre_15-0.20190509-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpcre_15-0.20190509-1PIGSTY.el8.aarch64.rpm) |
| `pgpcre_15` | `0.20190509` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.0 KiB | [pgpcre_15-0.20190509-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgpcre_15-0.20190509-1PGDG.rhel8.aarch64.rpm) |
| `pgpcre_15` | `0.20190509` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.3 KiB | [pgpcre_15-0.20190509-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpcre_15-0.20190509-1PIGSTY.el9.x86_64.rpm) |
| `pgpcre_15` | `0.20190509` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pgpcre_15-0.20190509-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgpcre_15-0.20190509-1PGDG.rhel9.x86_64.rpm) |
| `pgpcre_15` | `0.20190509` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.3 KiB | [pgpcre_15-0.20190509-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpcre_15-0.20190509-1PIGSTY.el9.aarch64.rpm) |
| `pgpcre_15` | `0.20190509` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.1 KiB | [pgpcre_15-0.20190509-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgpcre_15-0.20190509-1PGDG.rhel9.aarch64.rpm) |
| `pgpcre_15` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.1 KiB | [pgpcre_15-0.20190509-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgpcre_15-0.20190509-4PGDG.rhel10.x86_64.rpm) |
| `pgpcre_15` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.0 KiB | [pgpcre_15-0.20190509-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgpcre_15-0.20190509-3PGDG.rhel10.x86_64.rpm) |
| `pgpcre_15` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.0 KiB | [pgpcre_15-0.20190509-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgpcre_15-0.20190509-2PGDG.rhel10.x86_64.rpm) |
| `pgpcre_15` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 18.0 KiB | [pgpcre_15-0.20190509-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgpcre_15-0.20190509-4PGDG.rhel10.aarch64.rpm) |
| `pgpcre_15` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgpcre_15-0.20190509-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgpcre_15-0.20190509-3PGDG.rhel10.aarch64.rpm) |
| `pgpcre_15` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pgpcre_15-0.20190509-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgpcre_15-0.20190509-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgpcre` | `0.20190509` | [d12.x86_64](/os/d12.x86_64) | pgdg | 18.0 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg12+1_amd64.deb) |
| `postgresql-15-pgpcre` | `0.20190509` | [d12.aarch64](/os/d12.aarch64) | pgdg | 18.1 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg12+1_arm64.deb) |
| `postgresql-15-pgpcre` | `0.20190509` | [d13.x86_64](/os/d13.x86_64) | pgdg | 18.2 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg13+1_amd64.deb) |
| `postgresql-15-pgpcre` | `0.20190509` | [d13.aarch64](/os/d13.aarch64) | pgdg | 18.2 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg13+1_arm64.deb) |
| `postgresql-15-pgpcre` | `0.20190509` | [u22.x86_64](/os/u22.x86_64) | pgdg | 18.9 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgpcre` | `0.20190509` | [u22.aarch64](/os/u22.aarch64) | pgdg | 18.6 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgpcre` | `0.20190509` | [u24.x86_64](/os/u24.x86_64) | pgdg | 18.1 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgpcre` | `0.20190509` | [u24.aarch64](/os/u24.aarch64) | pgdg | 18.2 KiB | [postgresql-15-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-15-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgpcre_14` | `0.20190509` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.4 KiB | [pgpcre_14-0.20190509-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpcre_14-0.20190509-1PIGSTY.el8.x86_64.rpm) |
| `pgpcre_14` | `0.20190509` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.0 KiB | [pgpcre_14-0.20190509-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgpcre_14-0.20190509-1PGDG.rhel8.x86_64.rpm) |
| `pgpcre_14` | `0.20190509` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.5 KiB | [pgpcre_14-0.20190509-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpcre_14-0.20190509-1PIGSTY.el8.aarch64.rpm) |
| `pgpcre_14` | `0.20190509` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.0 KiB | [pgpcre_14-0.20190509-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgpcre_14-0.20190509-1PGDG.rhel8.aarch64.rpm) |
| `pgpcre_14` | `0.20190509` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.3 KiB | [pgpcre_14-0.20190509-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpcre_14-0.20190509-1PIGSTY.el9.x86_64.rpm) |
| `pgpcre_14` | `0.20190509` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pgpcre_14-0.20190509-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgpcre_14-0.20190509-1PGDG.rhel9.x86_64.rpm) |
| `pgpcre_14` | `0.20190509` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.2 KiB | [pgpcre_14-0.20190509-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpcre_14-0.20190509-1PIGSTY.el9.aarch64.rpm) |
| `pgpcre_14` | `0.20190509` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.1 KiB | [pgpcre_14-0.20190509-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgpcre_14-0.20190509-1PGDG.rhel9.aarch64.rpm) |
| `pgpcre_14` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.0 KiB | [pgpcre_14-0.20190509-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgpcre_14-0.20190509-4PGDG.rhel10.x86_64.rpm) |
| `pgpcre_14` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgpcre_14-0.20190509-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgpcre_14-0.20190509-3PGDG.rhel10.x86_64.rpm) |
| `pgpcre_14` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.8 KiB | [pgpcre_14-0.20190509-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgpcre_14-0.20190509-2PGDG.rhel10.x86_64.rpm) |
| `pgpcre_14` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 18.0 KiB | [pgpcre_14-0.20190509-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgpcre_14-0.20190509-4PGDG.rhel10.aarch64.rpm) |
| `pgpcre_14` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgpcre_14-0.20190509-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgpcre_14-0.20190509-3PGDG.rhel10.aarch64.rpm) |
| `pgpcre_14` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pgpcre_14-0.20190509-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgpcre_14-0.20190509-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgpcre` | `0.20190509` | [d12.x86_64](/os/d12.x86_64) | pgdg | 18.0 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg12+1_amd64.deb) |
| `postgresql-14-pgpcre` | `0.20190509` | [d12.aarch64](/os/d12.aarch64) | pgdg | 18.0 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg12+1_arm64.deb) |
| `postgresql-14-pgpcre` | `0.20190509` | [d13.x86_64](/os/d13.x86_64) | pgdg | 18.1 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg13+1_amd64.deb) |
| `postgresql-14-pgpcre` | `0.20190509` | [d13.aarch64](/os/d13.aarch64) | pgdg | 18.1 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg13+1_arm64.deb) |
| `postgresql-14-pgpcre` | `0.20190509` | [u22.x86_64](/os/u22.x86_64) | pgdg | 18.8 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgpcre` | `0.20190509` | [u22.aarch64](/os/u22.aarch64) | pgdg | 18.5 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgpcre` | `0.20190509` | [u24.x86_64](/os/u24.x86_64) | pgdg | 18.1 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgpcre` | `0.20190509` | [u24.aarch64](/os/u24.aarch64) | pgdg | 18.1 KiB | [postgresql-14-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-14-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgpcre_13` | `0.20190509` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.3 KiB | [pgpcre_13-0.20190509-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgpcre_13-0.20190509-1PIGSTY.el8.x86_64.rpm) |
| `pgpcre_13` | `0.20190509` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.0 KiB | [pgpcre_13-0.20190509-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgpcre_13-0.20190509-1PGDG.rhel8.x86_64.rpm) |
| `pgpcre_13` | `0.20190509` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.5 KiB | [pgpcre_13-0.20190509-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgpcre_13-0.20190509-1PIGSTY.el8.aarch64.rpm) |
| `pgpcre_13` | `0.20190509` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.0 KiB | [pgpcre_13-0.20190509-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgpcre_13-0.20190509-1PGDG.rhel8.aarch64.rpm) |
| `pgpcre_13` | `0.20190509` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.3 KiB | [pgpcre_13-0.20190509-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgpcre_13-0.20190509-1PIGSTY.el9.x86_64.rpm) |
| `pgpcre_13` | `0.20190509` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pgpcre_13-0.20190509-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgpcre_13-0.20190509-1PGDG.rhel9.x86_64.rpm) |
| `pgpcre_13` | `0.20190509` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.2 KiB | [pgpcre_13-0.20190509-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgpcre_13-0.20190509-1PIGSTY.el9.aarch64.rpm) |
| `pgpcre_13` | `0.20190509` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.0 KiB | [pgpcre_13-0.20190509-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgpcre_13-0.20190509-1PGDG.rhel9.aarch64.rpm) |
| `pgpcre_13` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.1 KiB | [pgpcre_13-0.20190509-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgpcre_13-0.20190509-4PGDG.rhel10.x86_64.rpm) |
| `pgpcre_13` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 18.0 KiB | [pgpcre_13-0.20190509-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgpcre_13-0.20190509-3PGDG.rhel10.x86_64.rpm) |
| `pgpcre_13` | `0.20190509` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgpcre_13-0.20190509-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgpcre_13-0.20190509-2PGDG.rhel10.x86_64.rpm) |
| `pgpcre_13` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 18.0 KiB | [pgpcre_13-0.20190509-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgpcre_13-0.20190509-4PGDG.rhel10.aarch64.rpm) |
| `pgpcre_13` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgpcre_13-0.20190509-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgpcre_13-0.20190509-3PGDG.rhel10.aarch64.rpm) |
| `pgpcre_13` | `0.20190509` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pgpcre_13-0.20190509-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgpcre_13-0.20190509-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pgpcre` | `0.20190509` | [d12.x86_64](/os/d12.x86_64) | pgdg | 17.9 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg12+1_amd64.deb) |
| `postgresql-13-pgpcre` | `0.20190509` | [d12.aarch64](/os/d12.aarch64) | pgdg | 17.9 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg12+1_arm64.deb) |
| `postgresql-13-pgpcre` | `0.20190509` | [d13.x86_64](/os/d13.x86_64) | pgdg | 17.9 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg13+1_amd64.deb) |
| `postgresql-13-pgpcre` | `0.20190509` | [d13.aarch64](/os/d13.aarch64) | pgdg | 17.9 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg13+1_arm64.deb) |
| `postgresql-13-pgpcre` | `0.20190509` | [u22.x86_64](/os/u22.x86_64) | pgdg | 18.4 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pgpcre` | `0.20190509` | [u22.aarch64](/os/u22.aarch64) | pgdg | 18.2 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pgpcre` | `0.20190509` | [u24.x86_64](/os/u24.x86_64) | pgdg | 18.0 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pgpcre` | `0.20190509` | [u24.aarch64](/os/u24.aarch64) | pgdg | 17.9 KiB | [postgresql-13-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgpcre/postgresql-13-pgpcre_0.20190509-9.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petere/pgpcre" title="Repository" icon="github" subtitle="github.com/petere/pgpcre" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgpcre-0.20190509.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgpcre;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgpcre;		# install via package name, for the active PG version

pig install pgpcre -v 18;   # install for PG 18
pig install pgpcre -v 17;   # install for PG 17
pig install pgpcre -v 16;   # install for PG 16
pig install pgpcre -v 15;   # install for PG 15
pig install pgpcre -v 14;   # install for PG 14
pig install pgpcre -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgpcre;
```
