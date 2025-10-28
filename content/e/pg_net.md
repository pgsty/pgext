---
title: "pg_net"
linkTitle: "pg_net"
description: "Async HTTP Requests"
weight: 4080
categories: ["UTIL"]
width: full
---

Async HTTP Requests


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4080** | {{< badge content="pg_net" link="https://github.com/supabase/pg_net" >}} | {{< ext "pg_net" >}} | `0.9.2` | {{< category "UTIL" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "http" >}} {{< ext "pg_curl" >}} {{< ext "pgjwt" >}} {{< ext "pg_smtp_client" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "pgjq" >}} |

> [!Note] break on lower libcurl version


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_net" >}} | `0.9.2` | {{< bg "18" "pg_net_18*" "red" >}} {{< bg "17" "pg_net_17*" "green" >}} {{< bg "16" "pg_net_16*" "green" >}} {{< bg "15" "pg_net_15*" "green" >}} {{< bg "14" "pg_net_14*" "green" >}} {{< bg "13" "pg_net_13*" "green" >}} | `pg_net_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_net" >}} | `0.9.2` | {{< bg "18" "postgresql-18-pg-net" "red" >}} {{< bg "17" "postgresql-17-pg-net" "green" >}} {{< bg "16" "postgresql-16-pg-net" "green" >}} {{< bg "15" "postgresql-15-pg-net" "green" >}} {{< bg "14" "postgresql-14-pg-net" "green" >}} {{< bg "13" "postgresql-13-pg-net" "green" >}} | `postgresql-$v-pg-net` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_net_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.9.2" "pg_net_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_13 : AVAIL 3" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_net_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.9.2" "pg_net_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_13 : AVAIL 3" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_net_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.9.2" "pg_net_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_13 : AVAIL 3" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_net_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.9.2" "pg_net_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.9.2" "pg_net_13 : AVAIL 3" "green" >}} |
|    `el10.x86_64`    | {{< bg "PGDG 0.20.0" "pg_net_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.20.0" "pg_net_17 : AVAIL 10" "blue" >}} | {{< bg "PGDG 0.20.0" "pg_net_16 : AVAIL 10" "blue" >}} | {{< bg "PGDG 0.20.0" "pg_net_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 0.20.0" "pg_net_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 0.20.0" "pg_net_13 : AVAIL 10" "blue" >}} |
|    `el10.aarch64`    | {{< bg "PGDG 0.20.0" "pg_net_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.20.0" "pg_net_17 : AVAIL 10" "blue" >}} | {{< bg "PGDG 0.20.0" "pg_net_16 : AVAIL 10" "blue" >}} | {{< bg "PGDG 0.20.0" "pg_net_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 0.20.0" "pg_net_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 0.20.0" "pg_net_13 : AVAIL 10" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-net : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-15-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-14-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-13-pg-net : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-net : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-15-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-14-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-13-pg-net : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-net : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-net : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-net : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-net : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-net : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-net : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-net : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-net : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-net : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-net : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-net : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-net : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-net : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.9.2" "postgresql-17-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-16-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-15-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-14-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-13-pg-net : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-net : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.9.2" "postgresql-17-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-16-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-15-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-14-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.2" "postgresql-13-pg-net : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-net : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-15-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-14-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-13-pg-net : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-net : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-15-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-14-pg-net : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-13-pg-net : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_net_18` | 0.20.0 | `el10.x86_64` | pgdg | 33.9 KiB | [pg_net_18-0.20.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_net_18-0.20.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_18` | 0.19.7 | `el10.x86_64` | pgdg | 33.0 KiB | [pg_net_18-0.19.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_net_18-0.19.7-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_18` | 0.20.0 | `el10.aarch64` | pgdg | 33.4 KiB | [pg_net_18-0.20.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_net_18-0.20.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_18` | 0.19.7 | `el10.aarch64` | pgdg | 32.5 KiB | [pg_net_18-0.19.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_net_18-0.19.7-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_net_17` | 0.9.2 | `el8.x86_64` | pigsty | 26.6 KiB | [pg_net_17-0.9.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_net_17-0.9.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_net_17` | 0.9.2 | `el8.aarch64` | pigsty | 26.3 KiB | [pg_net_17-0.9.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_net_17-0.9.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_net_17` | 0.9.2 | `el9.x86_64` | pigsty | 26.7 KiB | [pg_net_17-0.9.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_net_17-0.9.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_net_17` | 0.9.2 | `el9.aarch64` | pigsty | 26.2 KiB | [pg_net_17-0.9.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_net_17-0.9.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_net_17` | 0.20.0 | `el10.x86_64` | pgdg | 33.9 KiB | [pg_net_17-0.20.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_net_17-0.20.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_17` | 0.19.7 | `el10.x86_64` | pgdg | 33.0 KiB | [pg_net_17-0.19.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_net_17-0.19.7-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_17` | 0.19.6 | `el10.x86_64` | pgdg | 32.7 KiB | [pg_net_17-0.19.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_net_17-0.19.6-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_17` | 0.19.5 | `el10.x86_64` | pgdg | 32.4 KiB | [pg_net_17-0.19.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_net_17-0.19.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_17` | 0.19.4 | `el10.x86_64` | pgdg | 31.9 KiB | [pg_net_17-0.19.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_net_17-0.19.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_17` | 0.19.3 | `el10.x86_64` | pgdg | 31.5 KiB | [pg_net_17-0.19.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_net_17-0.19.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_17` | 0.19.1 | `el10.x86_64` | pgdg | 31.2 KiB | [pg_net_17-0.19.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_net_17-0.19.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_17` | 0.19.0 | `el10.x86_64` | pgdg | 30.9 KiB | [pg_net_17-0.19.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_net_17-0.19.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_17` | 0.16.0 | `el10.x86_64` | pgdg | 28.3 KiB | [pg_net_17-0.16.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_net_17-0.16.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_17` | 0.15.1 | `el10.x86_64` | pgdg | 28.0 KiB | [pg_net_17-0.15.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_net_17-0.15.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_17` | 0.20.0 | `el10.aarch64` | pgdg | 33.4 KiB | [pg_net_17-0.20.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_net_17-0.20.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_17` | 0.19.7 | `el10.aarch64` | pgdg | 32.5 KiB | [pg_net_17-0.19.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_net_17-0.19.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_17` | 0.19.6 | `el10.aarch64` | pgdg | 32.2 KiB | [pg_net_17-0.19.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_net_17-0.19.6-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_17` | 0.19.5 | `el10.aarch64` | pgdg | 31.8 KiB | [pg_net_17-0.19.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_net_17-0.19.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_17` | 0.19.4 | `el10.aarch64` | pgdg | 31.4 KiB | [pg_net_17-0.19.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_net_17-0.19.4-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_17` | 0.19.3 | `el10.aarch64` | pgdg | 31.0 KiB | [pg_net_17-0.19.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_net_17-0.19.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_17` | 0.19.1 | `el10.aarch64` | pgdg | 30.7 KiB | [pg_net_17-0.19.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_net_17-0.19.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_17` | 0.19.0 | `el10.aarch64` | pgdg | 30.3 KiB | [pg_net_17-0.19.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_net_17-0.19.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_17` | 0.16.0 | `el10.aarch64` | pgdg | 27.6 KiB | [pg_net_17-0.16.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_net_17-0.16.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_17` | 0.15.1 | `el10.aarch64` | pgdg | 27.4 KiB | [pg_net_17-0.15.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_net_17-0.15.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-net` | 0.14.0 | `d12.x86_64` | pigsty | 52.3 KiB | [postgresql-17-pg-net_0.14.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-net/postgresql-17-pg-net_0.14.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-net` | 0.14.0 | `d12.aarch64` | pigsty | 51.4 KiB | [postgresql-17-pg-net_0.14.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-net/postgresql-17-pg-net_0.14.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-net` | 0.9.2 | `u22.x86_64` | pigsty | 44.0 KiB | [postgresql-17-pg-net_0.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-net/postgresql-17-pg-net_0.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-net` | 0.9.2 | `u22.aarch64` | pigsty | 43.4 KiB | [postgresql-17-pg-net_0.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-net/postgresql-17-pg-net_0.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-net` | 0.14.0 | `u24.x86_64` | pigsty | 54.2 KiB | [postgresql-17-pg-net_0.14.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-net/postgresql-17-pg-net_0.14.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-net` | 0.14.0 | `u24.aarch64` | pigsty | 53.4 KiB | [postgresql-17-pg-net_0.14.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-net/postgresql-17-pg-net_0.14.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_net_16` | 0.9.2 | `el8.x86_64` | pigsty | 26.7 KiB | [pg_net_16-0.9.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_net_16-0.9.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_net_16` | 0.9.2 | `el8.x86_64` | pgdg | 21.9 KiB | [pg_net_16-0.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_net_16-0.9.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_net_16` | 0.9.1 | `el8.x86_64` | pgdg | 21.0 KiB | [pg_net_16-0.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_net_16-0.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_net_16` | 0.9.2 | `el8.aarch64` | pigsty | 26.2 KiB | [pg_net_16-0.9.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_net_16-0.9.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_net_16` | 0.9.2 | `el8.aarch64` | pgdg | 21.4 KiB | [pg_net_16-0.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_net_16-0.9.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_net_16` | 0.9.1 | `el8.aarch64` | pgdg | 20.6 KiB | [pg_net_16-0.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_net_16-0.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_net_16` | 0.9.2 | `el9.x86_64` | pigsty | 26.7 KiB | [pg_net_16-0.9.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_net_16-0.9.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_net_16` | 0.9.2 | `el9.x86_64` | pgdg | 21.9 KiB | [pg_net_16-0.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_net_16-0.9.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_net_16` | 0.9.1 | `el9.x86_64` | pgdg | 20.9 KiB | [pg_net_16-0.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_net_16-0.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_net_16` | 0.9.2 | `el9.aarch64` | pigsty | 26.2 KiB | [pg_net_16-0.9.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_net_16-0.9.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_net_16` | 0.9.2 | `el9.aarch64` | pgdg | 21.4 KiB | [pg_net_16-0.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_net_16-0.9.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_net_16` | 0.9.1 | `el9.aarch64` | pgdg | 20.4 KiB | [pg_net_16-0.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_net_16-0.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_net_16` | 0.20.0 | `el10.x86_64` | pgdg | 33.9 KiB | [pg_net_16-0.20.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_net_16-0.20.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_16` | 0.19.7 | `el10.x86_64` | pgdg | 33.0 KiB | [pg_net_16-0.19.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_net_16-0.19.7-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_16` | 0.19.6 | `el10.x86_64` | pgdg | 32.7 KiB | [pg_net_16-0.19.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_net_16-0.19.6-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_16` | 0.19.5 | `el10.x86_64` | pgdg | 32.3 KiB | [pg_net_16-0.19.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_net_16-0.19.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_16` | 0.19.4 | `el10.x86_64` | pgdg | 31.9 KiB | [pg_net_16-0.19.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_net_16-0.19.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_16` | 0.19.3 | `el10.x86_64` | pgdg | 31.5 KiB | [pg_net_16-0.19.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_net_16-0.19.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_16` | 0.19.1 | `el10.x86_64` | pgdg | 31.2 KiB | [pg_net_16-0.19.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_net_16-0.19.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_16` | 0.19.0 | `el10.x86_64` | pgdg | 30.8 KiB | [pg_net_16-0.19.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_net_16-0.19.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_16` | 0.16.0 | `el10.x86_64` | pgdg | 28.3 KiB | [pg_net_16-0.16.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_net_16-0.16.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_16` | 0.15.1 | `el10.x86_64` | pgdg | 28.1 KiB | [pg_net_16-0.15.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_net_16-0.15.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_16` | 0.20.0 | `el10.aarch64` | pgdg | 33.4 KiB | [pg_net_16-0.20.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_net_16-0.20.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_16` | 0.19.7 | `el10.aarch64` | pgdg | 32.5 KiB | [pg_net_16-0.19.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_net_16-0.19.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_16` | 0.19.6 | `el10.aarch64` | pgdg | 32.2 KiB | [pg_net_16-0.19.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_net_16-0.19.6-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_16` | 0.19.5 | `el10.aarch64` | pgdg | 31.8 KiB | [pg_net_16-0.19.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_net_16-0.19.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_16` | 0.19.4 | `el10.aarch64` | pgdg | 31.3 KiB | [pg_net_16-0.19.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_net_16-0.19.4-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_16` | 0.19.3 | `el10.aarch64` | pgdg | 31.0 KiB | [pg_net_16-0.19.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_net_16-0.19.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_16` | 0.19.1 | `el10.aarch64` | pgdg | 30.7 KiB | [pg_net_16-0.19.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_net_16-0.19.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_16` | 0.19.0 | `el10.aarch64` | pgdg | 30.3 KiB | [pg_net_16-0.19.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_net_16-0.19.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_16` | 0.16.0 | `el10.aarch64` | pgdg | 27.6 KiB | [pg_net_16-0.16.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_net_16-0.16.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_16` | 0.15.1 | `el10.aarch64` | pgdg | 27.3 KiB | [pg_net_16-0.15.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_net_16-0.15.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-net` | 0.14.0 | `d12.x86_64` | pigsty | 52.2 KiB | [postgresql-16-pg-net_0.14.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-net/postgresql-16-pg-net_0.14.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-net` | 0.14.0 | `d12.aarch64` | pigsty | 51.3 KiB | [postgresql-16-pg-net_0.14.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-net/postgresql-16-pg-net_0.14.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-net` | 0.9.2 | `u22.x86_64` | pigsty | 44.1 KiB | [postgresql-16-pg-net_0.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-net/postgresql-16-pg-net_0.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-net` | 0.9.2 | `u22.aarch64` | pigsty | 43.4 KiB | [postgresql-16-pg-net_0.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-net/postgresql-16-pg-net_0.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-net` | 0.14.0 | `u24.x86_64` | pigsty | 54.1 KiB | [postgresql-16-pg-net_0.14.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-net/postgresql-16-pg-net_0.14.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-net` | 0.14.0 | `u24.aarch64` | pigsty | 53.4 KiB | [postgresql-16-pg-net_0.14.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-net/postgresql-16-pg-net_0.14.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_net_15` | 0.9.2 | `el8.x86_64` | pigsty | 26.7 KiB | [pg_net_15-0.9.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_net_15-0.9.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_net_15` | 0.9.2 | `el8.x86_64` | pgdg | 21.9 KiB | [pg_net_15-0.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_net_15-0.9.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_net_15` | 0.9.1 | `el8.x86_64` | pgdg | 21.0 KiB | [pg_net_15-0.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_net_15-0.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_net_15` | 0.9.2 | `el8.aarch64` | pigsty | 26.3 KiB | [pg_net_15-0.9.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_net_15-0.9.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_net_15` | 0.9.2 | `el8.aarch64` | pgdg | 21.5 KiB | [pg_net_15-0.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_net_15-0.9.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_net_15` | 0.9.1 | `el8.aarch64` | pgdg | 20.7 KiB | [pg_net_15-0.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_net_15-0.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_net_15` | 0.9.2 | `el9.x86_64` | pigsty | 27.0 KiB | [pg_net_15-0.9.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_net_15-0.9.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_net_15` | 0.9.2 | `el9.x86_64` | pgdg | 22.1 KiB | [pg_net_15-0.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_net_15-0.9.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_net_15` | 0.9.1 | `el9.x86_64` | pgdg | 21.1 KiB | [pg_net_15-0.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_net_15-0.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_net_15` | 0.9.2 | `el9.aarch64` | pigsty | 26.4 KiB | [pg_net_15-0.9.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_net_15-0.9.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_net_15` | 0.9.2 | `el9.aarch64` | pgdg | 21.6 KiB | [pg_net_15-0.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_net_15-0.9.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_net_15` | 0.9.1 | `el9.aarch64` | pgdg | 20.6 KiB | [pg_net_15-0.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_net_15-0.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_net_15` | 0.20.0 | `el10.x86_64` | pgdg | 35.0 KiB | [pg_net_15-0.20.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_net_15-0.20.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_15` | 0.19.7 | `el10.x86_64` | pgdg | 33.9 KiB | [pg_net_15-0.19.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_net_15-0.19.7-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_15` | 0.19.6 | `el10.x86_64` | pgdg | 33.6 KiB | [pg_net_15-0.19.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_net_15-0.19.6-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_15` | 0.19.5 | `el10.x86_64` | pgdg | 33.2 KiB | [pg_net_15-0.19.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_net_15-0.19.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_15` | 0.19.4 | `el10.x86_64` | pgdg | 32.9 KiB | [pg_net_15-0.19.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_net_15-0.19.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_15` | 0.19.3 | `el10.x86_64` | pgdg | 32.5 KiB | [pg_net_15-0.19.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_net_15-0.19.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_15` | 0.19.1 | `el10.x86_64` | pgdg | 32.2 KiB | [pg_net_15-0.19.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_net_15-0.19.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_15` | 0.19.0 | `el10.x86_64` | pgdg | 31.9 KiB | [pg_net_15-0.19.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_net_15-0.19.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_15` | 0.16.0 | `el10.x86_64` | pgdg | 29.2 KiB | [pg_net_15-0.16.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_net_15-0.16.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_15` | 0.15.1 | `el10.x86_64` | pgdg | 29.1 KiB | [pg_net_15-0.15.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_net_15-0.15.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_15` | 0.20.0 | `el10.aarch64` | pgdg | 34.3 KiB | [pg_net_15-0.20.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_net_15-0.20.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_15` | 0.19.7 | `el10.aarch64` | pgdg | 33.4 KiB | [pg_net_15-0.19.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_net_15-0.19.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_15` | 0.19.6 | `el10.aarch64` | pgdg | 33.0 KiB | [pg_net_15-0.19.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_net_15-0.19.6-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_15` | 0.19.5 | `el10.aarch64` | pgdg | 32.6 KiB | [pg_net_15-0.19.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_net_15-0.19.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_15` | 0.19.4 | `el10.aarch64` | pgdg | 32.2 KiB | [pg_net_15-0.19.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_net_15-0.19.4-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_15` | 0.19.3 | `el10.aarch64` | pgdg | 31.7 KiB | [pg_net_15-0.19.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_net_15-0.19.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_15` | 0.19.1 | `el10.aarch64` | pgdg | 31.4 KiB | [pg_net_15-0.19.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_net_15-0.19.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_15` | 0.19.0 | `el10.aarch64` | pgdg | 31.1 KiB | [pg_net_15-0.19.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_net_15-0.19.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_15` | 0.16.0 | `el10.aarch64` | pgdg | 28.4 KiB | [pg_net_15-0.16.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_net_15-0.16.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_15` | 0.15.1 | `el10.aarch64` | pgdg | 28.1 KiB | [pg_net_15-0.15.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_net_15-0.15.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-net` | 0.14.0 | `d12.x86_64` | pigsty | 52.6 KiB | [postgresql-15-pg-net_0.14.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-net/postgresql-15-pg-net_0.14.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-net` | 0.14.0 | `d12.aarch64` | pigsty | 51.6 KiB | [postgresql-15-pg-net_0.14.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-net/postgresql-15-pg-net_0.14.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-net` | 0.9.2 | `u22.x86_64` | pigsty | 44.3 KiB | [postgresql-15-pg-net_0.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-net/postgresql-15-pg-net_0.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-net` | 0.9.2 | `u22.aarch64` | pigsty | 43.5 KiB | [postgresql-15-pg-net_0.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-net/postgresql-15-pg-net_0.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-net` | 0.14.0 | `u24.x86_64` | pigsty | 54.8 KiB | [postgresql-15-pg-net_0.14.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-net/postgresql-15-pg-net_0.14.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-net` | 0.14.0 | `u24.aarch64` | pigsty | 54.1 KiB | [postgresql-15-pg-net_0.14.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-net/postgresql-15-pg-net_0.14.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_net_14` | 0.9.2 | `el8.x86_64` | pigsty | 26.7 KiB | [pg_net_14-0.9.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_net_14-0.9.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_net_14` | 0.9.2 | `el8.x86_64` | pgdg | 21.9 KiB | [pg_net_14-0.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_net_14-0.9.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_net_14` | 0.9.1 | `el8.x86_64` | pgdg | 21.0 KiB | [pg_net_14-0.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_net_14-0.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_net_14` | 0.9.2 | `el8.aarch64` | pigsty | 26.3 KiB | [pg_net_14-0.9.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_net_14-0.9.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_net_14` | 0.9.2 | `el8.aarch64` | pgdg | 21.5 KiB | [pg_net_14-0.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_net_14-0.9.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_net_14` | 0.9.1 | `el8.aarch64` | pgdg | 20.7 KiB | [pg_net_14-0.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_net_14-0.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_net_14` | 0.9.2 | `el9.x86_64` | pigsty | 26.9 KiB | [pg_net_14-0.9.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_net_14-0.9.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_net_14` | 0.9.2 | `el9.x86_64` | pgdg | 22.1 KiB | [pg_net_14-0.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_net_14-0.9.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_net_14` | 0.9.1 | `el9.x86_64` | pgdg | 21.1 KiB | [pg_net_14-0.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_net_14-0.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_net_14` | 0.9.2 | `el9.aarch64` | pigsty | 26.4 KiB | [pg_net_14-0.9.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_net_14-0.9.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_net_14` | 0.9.2 | `el9.aarch64` | pgdg | 21.6 KiB | [pg_net_14-0.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_net_14-0.9.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_net_14` | 0.9.1 | `el9.aarch64` | pgdg | 20.6 KiB | [pg_net_14-0.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_net_14-0.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_net_14` | 0.20.0 | `el10.x86_64` | pgdg | 34.8 KiB | [pg_net_14-0.20.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_net_14-0.20.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_14` | 0.19.7 | `el10.x86_64` | pgdg | 33.9 KiB | [pg_net_14-0.19.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_net_14-0.19.7-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_14` | 0.19.6 | `el10.x86_64` | pgdg | 33.6 KiB | [pg_net_14-0.19.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_net_14-0.19.6-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_14` | 0.19.5 | `el10.x86_64` | pgdg | 33.2 KiB | [pg_net_14-0.19.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_net_14-0.19.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_14` | 0.19.4 | `el10.x86_64` | pgdg | 32.9 KiB | [pg_net_14-0.19.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_net_14-0.19.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_14` | 0.19.3 | `el10.x86_64` | pgdg | 32.4 KiB | [pg_net_14-0.19.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_net_14-0.19.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_14` | 0.19.1 | `el10.x86_64` | pgdg | 32.2 KiB | [pg_net_14-0.19.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_net_14-0.19.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_14` | 0.19.0 | `el10.x86_64` | pgdg | 31.8 KiB | [pg_net_14-0.19.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_net_14-0.19.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_14` | 0.16.0 | `el10.x86_64` | pgdg | 29.2 KiB | [pg_net_14-0.16.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_net_14-0.16.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_14` | 0.15.1 | `el10.x86_64` | pgdg | 29.1 KiB | [pg_net_14-0.15.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_net_14-0.15.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_14` | 0.20.0 | `el10.aarch64` | pgdg | 34.2 KiB | [pg_net_14-0.20.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_net_14-0.20.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_14` | 0.19.7 | `el10.aarch64` | pgdg | 33.3 KiB | [pg_net_14-0.19.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_net_14-0.19.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_14` | 0.19.6 | `el10.aarch64` | pgdg | 33.0 KiB | [pg_net_14-0.19.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_net_14-0.19.6-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_14` | 0.19.5 | `el10.aarch64` | pgdg | 32.5 KiB | [pg_net_14-0.19.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_net_14-0.19.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_14` | 0.19.4 | `el10.aarch64` | pgdg | 32.1 KiB | [pg_net_14-0.19.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_net_14-0.19.4-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_14` | 0.19.3 | `el10.aarch64` | pgdg | 31.7 KiB | [pg_net_14-0.19.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_net_14-0.19.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_14` | 0.19.1 | `el10.aarch64` | pgdg | 31.4 KiB | [pg_net_14-0.19.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_net_14-0.19.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_14` | 0.19.0 | `el10.aarch64` | pgdg | 31.0 KiB | [pg_net_14-0.19.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_net_14-0.19.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_14` | 0.16.0 | `el10.aarch64` | pgdg | 28.4 KiB | [pg_net_14-0.16.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_net_14-0.16.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_14` | 0.15.1 | `el10.aarch64` | pgdg | 28.1 KiB | [pg_net_14-0.15.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_net_14-0.15.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-net` | 0.14.0 | `d12.x86_64` | pigsty | 52.5 KiB | [postgresql-14-pg-net_0.14.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-net/postgresql-14-pg-net_0.14.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-net` | 0.14.0 | `d12.aarch64` | pigsty | 51.5 KiB | [postgresql-14-pg-net_0.14.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-net/postgresql-14-pg-net_0.14.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-net` | 0.9.2 | `u22.x86_64` | pigsty | 44.2 KiB | [postgresql-14-pg-net_0.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-net/postgresql-14-pg-net_0.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-net` | 0.9.2 | `u22.aarch64` | pigsty | 43.5 KiB | [postgresql-14-pg-net_0.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-net/postgresql-14-pg-net_0.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-net` | 0.14.0 | `u24.x86_64` | pigsty | 54.7 KiB | [postgresql-14-pg-net_0.14.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-net/postgresql-14-pg-net_0.14.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-net` | 0.14.0 | `u24.aarch64` | pigsty | 54.1 KiB | [postgresql-14-pg-net_0.14.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-net/postgresql-14-pg-net_0.14.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_net_13` | 0.9.2 | `el8.x86_64` | pigsty | 26.6 KiB | [pg_net_13-0.9.2-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_net_13-0.9.2-2PIGSTY.el8.x86_64.rpm) |
| `pg_net_13` | 0.9.2 | `el8.x86_64` | pgdg | 21.8 KiB | [pg_net_13-0.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_net_13-0.9.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_net_13` | 0.9.1 | `el8.x86_64` | pgdg | 20.9 KiB | [pg_net_13-0.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_net_13-0.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_net_13` | 0.9.2 | `el8.aarch64` | pigsty | 26.3 KiB | [pg_net_13-0.9.2-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_net_13-0.9.2-2PIGSTY.el8.aarch64.rpm) |
| `pg_net_13` | 0.9.2 | `el8.aarch64` | pgdg | 21.5 KiB | [pg_net_13-0.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_net_13-0.9.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_net_13` | 0.9.1 | `el8.aarch64` | pgdg | 20.7 KiB | [pg_net_13-0.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_net_13-0.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_net_13` | 0.9.2 | `el9.x86_64` | pigsty | 26.9 KiB | [pg_net_13-0.9.2-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_net_13-0.9.2-2PIGSTY.el9.x86_64.rpm) |
| `pg_net_13` | 0.9.2 | `el9.x86_64` | pgdg | 22.1 KiB | [pg_net_13-0.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_net_13-0.9.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_net_13` | 0.9.1 | `el9.x86_64` | pgdg | 21.1 KiB | [pg_net_13-0.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_net_13-0.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_net_13` | 0.9.2 | `el9.aarch64` | pigsty | 26.4 KiB | [pg_net_13-0.9.2-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_net_13-0.9.2-2PIGSTY.el9.aarch64.rpm) |
| `pg_net_13` | 0.9.2 | `el9.aarch64` | pgdg | 21.6 KiB | [pg_net_13-0.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_net_13-0.9.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_net_13` | 0.9.1 | `el9.aarch64` | pgdg | 20.6 KiB | [pg_net_13-0.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_net_13-0.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_net_13` | 0.20.0 | `el10.x86_64` | pgdg | 34.9 KiB | [pg_net_13-0.20.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_net_13-0.20.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_13` | 0.19.7 | `el10.x86_64` | pgdg | 33.9 KiB | [pg_net_13-0.19.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_net_13-0.19.7-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_13` | 0.19.6 | `el10.x86_64` | pgdg | 33.5 KiB | [pg_net_13-0.19.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_net_13-0.19.6-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_13` | 0.19.5 | `el10.x86_64` | pgdg | 33.2 KiB | [pg_net_13-0.19.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_net_13-0.19.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_13` | 0.19.4 | `el10.x86_64` | pgdg | 32.8 KiB | [pg_net_13-0.19.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_net_13-0.19.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_13` | 0.19.3 | `el10.x86_64` | pgdg | 32.4 KiB | [pg_net_13-0.19.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_net_13-0.19.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_13` | 0.19.1 | `el10.x86_64` | pgdg | 32.1 KiB | [pg_net_13-0.19.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_net_13-0.19.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_13` | 0.19.0 | `el10.x86_64` | pgdg | 31.8 KiB | [pg_net_13-0.19.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_net_13-0.19.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_13` | 0.16.0 | `el10.x86_64` | pgdg | 29.2 KiB | [pg_net_13-0.16.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_net_13-0.16.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_13` | 0.15.1 | `el10.x86_64` | pgdg | 29.1 KiB | [pg_net_13-0.15.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_net_13-0.15.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_net_13` | 0.20.0 | `el10.aarch64` | pgdg | 34.3 KiB | [pg_net_13-0.20.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_net_13-0.20.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_13` | 0.19.7 | `el10.aarch64` | pgdg | 33.3 KiB | [pg_net_13-0.19.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_net_13-0.19.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_13` | 0.19.6 | `el10.aarch64` | pgdg | 33.0 KiB | [pg_net_13-0.19.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_net_13-0.19.6-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_13` | 0.19.5 | `el10.aarch64` | pgdg | 32.5 KiB | [pg_net_13-0.19.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_net_13-0.19.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_13` | 0.19.4 | `el10.aarch64` | pgdg | 32.1 KiB | [pg_net_13-0.19.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_net_13-0.19.4-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_13` | 0.19.3 | `el10.aarch64` | pgdg | 31.7 KiB | [pg_net_13-0.19.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_net_13-0.19.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_13` | 0.19.1 | `el10.aarch64` | pgdg | 31.5 KiB | [pg_net_13-0.19.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_net_13-0.19.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_13` | 0.19.0 | `el10.aarch64` | pgdg | 31.0 KiB | [pg_net_13-0.19.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_net_13-0.19.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_13` | 0.16.0 | `el10.aarch64` | pgdg | 28.3 KiB | [pg_net_13-0.16.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_net_13-0.16.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_net_13` | 0.15.1 | `el10.aarch64` | pgdg | 28.1 KiB | [pg_net_13-0.15.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_net_13-0.15.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-net` | 0.14.0 | `d12.x86_64` | pigsty | 52.0 KiB | [postgresql-13-pg-net_0.14.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-net/postgresql-13-pg-net_0.14.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-net` | 0.14.0 | `d12.aarch64` | pigsty | 51.1 KiB | [postgresql-13-pg-net_0.14.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-net/postgresql-13-pg-net_0.14.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-net` | 0.9.2 | `u22.x86_64` | pigsty | 43.7 KiB | [postgresql-13-pg-net_0.9.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-net/postgresql-13-pg-net_0.9.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-net` | 0.9.2 | `u22.aarch64` | pigsty | 43.1 KiB | [postgresql-13-pg-net_0.9.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-net/postgresql-13-pg-net_0.9.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-net` | 0.14.0 | `u24.x86_64` | pigsty | 54.3 KiB | [postgresql-13-pg-net_0.14.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-net/postgresql-13-pg-net_0.14.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-net` | 0.14.0 | `u24.aarch64` | pigsty | 53.6 KiB | [postgresql-13-pg-net_0.14.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-net/postgresql-13-pg-net_0.14.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/pg_net" title="Repository" icon="github" subtitle="github.com/supabase/pg_net" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_net-0.9.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_net; # get pg_net source code
pig build dep pg_net; # install build dependencies
pig build pkg pg_net; # build extension rpm or deb
pig build ext pg_net; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_net; # install by extension name, for the current active PG version
pig ext install pg_net; # install via package alias, for the active PG version
pig ext install pg_net -v 17;   # install for PG 17
pig ext install pg_net -v 16;   # install for PG 16
pig ext install pg_net -v 15;   # install for PG 15
pig ext install pg_net -v 14;   # install for PG 14
pig ext install pg_net -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_net;
```

