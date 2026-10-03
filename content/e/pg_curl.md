---
title: "pg_curl"
linkTitle: "pg_curl"
description: "Run curl actions for data transfer in URL syntax"
weight: 4090
categories: ["UTIL"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_curl**](https://github.com/RekGRpth/pg_curl) : Run curl actions for data transfer in URL syntax


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4090** | {{< badge content="pg_curl" link="https://github.com/RekGRpth/pg_curl" >}} | {{< ext "pg_curl" >}} | `2.4.6` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "omni_httpc" >}} {{< ext "pg_graphql" >}} {{< ext "documentdb" >}} |

> [!Note] Package 2.4.6; SQL control version 2.4.1. Retains the 2.4 to 2.4.1 SQL compatibility edge.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_curl` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.6` | {{< bg "18" "pg_curl_18" "green" >}} {{< bg "17" "pg_curl_17" "green" >}} {{< bg "16" "pg_curl_16" "green" >}} {{< bg "15" "pg_curl_15" "green" >}} {{< bg "14" "pg_curl_14" "green" >}} | `pg_curl_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.6` | {{< bg "18" "postgresql-18-pg-curl" "green" >}} {{< bg "17" "postgresql-17-pg-curl" "green" >}} {{< bg "16" "postgresql-16-pg-curl" "green" >}} {{< bg "15" "postgresql-15-pg-curl" "green" >}} {{< bg "14" "postgresql-14-pg-curl" "green" >}} | `postgresql-$v-pg-curl` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_14 : AVAIL 4" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_14 : AVAIL 4" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_14 : AVAIL 5" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_14 : AVAIL 5" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_14 : AVAIL 5" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.6" "pg_curl_14 : AVAIL 5" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-18-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-17-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-16-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-15-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-14-pg-curl : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-18-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-17-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-16-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-15-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-14-pg-curl : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-18-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-17-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-16-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-15-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-14-pg-curl : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-18-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-17-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-16-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-15-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-14-pg-curl : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-18-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-17-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-16-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-15-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-14-pg-curl : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-18-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-17-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-16-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-15-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-14-pg-curl : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-18-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-17-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-16-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-15-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-14-pg-curl : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-18-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-17-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-16-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-15-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-14-pg-curl : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-18-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-17-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-16-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-15-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-14-pg-curl : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-18-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-17-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-16-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-15-pg-curl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.6" "postgresql-14-pg-curl : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_curl_18` | `2.4.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 124.5 KiB | [pg_curl_18-2.4.6-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_18-2.4.6-1PGSTY.el8.x86_64.rpm) |
| `pg_curl_18` | `2.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_curl_18-2.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_curl_18-2.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_curl_18` | `2.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.8 KiB | [pg_curl_18-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_curl_18-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_18` | `2.4.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 122.6 KiB | [pg_curl_18-2.4.6-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_18-2.4.6-1PGSTY.el8.aarch64.rpm) |
| `pg_curl_18` | `2.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_curl_18-2.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_curl_18-2.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_curl_18` | `2.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_curl_18-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_curl_18-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_18` | `2.4.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 124.9 KiB | [pg_curl_18-2.4.6-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_18-2.4.6-1PGSTY.el9.x86_64.rpm) |
| `pg_curl_18` | `2.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.5 KiB | [pg_curl_18-2.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_curl_18-2.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_18` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.8 KiB | [pg_curl_18-2.4.4-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_curl_18-2.4.4-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_18` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.5 KiB | [pg_curl_18-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_curl_18-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_18` | `2.4.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 124.0 KiB | [pg_curl_18-2.4.6-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_18-2.4.6-1PGSTY.el9.aarch64.rpm) |
| `pg_curl_18` | `2.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.9 KiB | [pg_curl_18-2.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_curl_18-2.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_18` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.3 KiB | [pg_curl_18-2.4.4-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_curl_18-2.4.4-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_18` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [pg_curl_18-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_curl_18-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_18` | `2.4.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 126.1 KiB | [pg_curl_18-2.4.6-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_curl_18-2.4.6-1PGSTY.el10.x86_64.rpm) |
| `pg_curl_18` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.1 KiB | [pg_curl_18-2.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_curl_18-2.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_18` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_18-2.4.4-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_curl_18-2.4.4-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_18` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.5 KiB | [pg_curl_18-2.4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_curl_18-2.4.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_curl_18` | `2.4.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 125.2 KiB | [pg_curl_18-2.4.6-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_curl_18-2.4.6-1PGSTY.el10.aarch64.rpm) |
| `pg_curl_18` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.6 KiB | [pg_curl_18-2.4.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_curl_18-2.4.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_curl_18` | `2.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.3 KiB | [pg_curl_18-2.4.4-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_curl_18-2.4.4-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_curl_18` | `2.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.5 KiB | [pg_curl_18-2.4.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_curl_18-2.4.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-curl` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 107.0 KiB | [postgresql-18-pg-curl_2.4.6-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-18-pg-curl_2.4.6-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-curl` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.4 KiB | [postgresql-18-pg-curl_2.4.6-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-18-pg-curl_2.4.6-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-curl` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 107.4 KiB | [postgresql-18-pg-curl_2.4.6-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-curl/postgresql-18-pg-curl_2.4.6-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-curl` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 106.1 KiB | [postgresql-18-pg-curl_2.4.6-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-curl/postgresql-18-pg-curl_2.4.6-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-curl` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 122.0 KiB | [postgresql-18-pg-curl_2.4.6-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-18-pg-curl_2.4.6-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-curl` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 120.3 KiB | [postgresql-18-pg-curl_2.4.6-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-18-pg-curl_2.4.6-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-curl` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 115.4 KiB | [postgresql-18-pg-curl_2.4.6-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-18-pg-curl_2.4.6-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-curl` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 114.7 KiB | [postgresql-18-pg-curl_2.4.6-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-18-pg-curl_2.4.6-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-curl` | `2.4.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 121.5 KiB | [postgresql-18-pg-curl_2.4.6-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-curl/postgresql-18-pg-curl_2.4.6-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-curl` | `2.4.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 120.9 KiB | [postgresql-18-pg-curl_2.4.6-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-curl/postgresql-18-pg-curl_2.4.6-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_curl_17` | `2.4.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 124.5 KiB | [pg_curl_17-2.4.6-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_17-2.4.6-1PGSTY.el8.x86_64.rpm) |
| `pg_curl_17` | `2.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_curl_17-2.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_curl_17-2.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_curl_17` | `2.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.8 KiB | [pg_curl_17-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_curl_17-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_17` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.7 KiB | [pg_curl_17-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_curl_17-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_17` | `2.4.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 122.6 KiB | [pg_curl_17-2.4.6-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_17-2.4.6-1PGSTY.el8.aarch64.rpm) |
| `pg_curl_17` | `2.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_curl_17-2.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_curl_17-2.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_curl_17` | `2.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_curl_17-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_curl_17-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_17` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.9 KiB | [pg_curl_17-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_curl_17-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_17` | `2.4.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 124.9 KiB | [pg_curl_17-2.4.6-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_17-2.4.6-1PGSTY.el9.x86_64.rpm) |
| `pg_curl_17` | `2.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.6 KiB | [pg_curl_17-2.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_curl_17-2.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_17` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.9 KiB | [pg_curl_17-2.4.4-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_curl_17-2.4.4-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_17` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.6 KiB | [pg_curl_17-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_curl_17-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_17` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.6 KiB | [pg_curl_17-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_curl_17-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_17` | `2.4.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 124.0 KiB | [pg_curl_17-2.4.6-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_17-2.4.6-1PGSTY.el9.aarch64.rpm) |
| `pg_curl_17` | `2.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.8 KiB | [pg_curl_17-2.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_curl_17-2.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_17` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.1 KiB | [pg_curl_17-2.4.4-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_curl_17-2.4.4-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_17` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.8 KiB | [pg_curl_17-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_curl_17-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_17` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [pg_curl_17-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_curl_17-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_17` | `2.4.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 125.9 KiB | [pg_curl_17-2.4.6-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_curl_17-2.4.6-1PGSTY.el10.x86_64.rpm) |
| `pg_curl_17` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.1 KiB | [pg_curl_17-2.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_curl_17-2.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_17` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_17-2.4.4-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_curl_17-2.4.4-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_17` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.6 KiB | [pg_curl_17-2.4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_curl_17-2.4.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_curl_17` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_17-2.4.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_curl_17-2.4.3-2PGDG.rhel10.x86_64.rpm) |
| `pg_curl_17` | `2.4.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 125.2 KiB | [pg_curl_17-2.4.6-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_curl_17-2.4.6-1PGSTY.el10.aarch64.rpm) |
| `pg_curl_17` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.0 KiB | [pg_curl_17-2.4.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_curl_17-2.4.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_curl_17` | `2.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.4 KiB | [pg_curl_17-2.4.4-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_curl_17-2.4.4-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_curl_17` | `2.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.2 KiB | [pg_curl_17-2.4.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_curl_17-2.4.4-1PGDG.rhel10.aarch64.rpm) |
| `pg_curl_17` | `2.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.0 KiB | [pg_curl_17-2.4.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_curl_17-2.4.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-curl` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 107.6 KiB | [postgresql-17-pg-curl_2.4.6-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.6-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-curl` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.5 KiB | [postgresql-17-pg-curl_2.4.6-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.6-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-curl` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 109.0 KiB | [postgresql-17-pg-curl_2.4.6-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.6-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-curl` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 106.2 KiB | [postgresql-17-pg-curl_2.4.6-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.6-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-curl` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.5 KiB | [postgresql-17-pg-curl_2.4.6-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.6-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-curl` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.8 KiB | [postgresql-17-pg-curl_2.4.6-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.6-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-curl` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 115.4 KiB | [postgresql-17-pg-curl_2.4.6-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.6-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-curl` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 114.8 KiB | [postgresql-17-pg-curl_2.4.6-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.6-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-curl` | `2.4.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 121.4 KiB | [postgresql-17-pg-curl_2.4.6-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.6-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-curl` | `2.4.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 120.7 KiB | [postgresql-17-pg-curl_2.4.6-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-curl/postgresql-17-pg-curl_2.4.6-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_curl_16` | `2.4.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 124.5 KiB | [pg_curl_16-2.4.6-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_16-2.4.6-1PGSTY.el8.x86_64.rpm) |
| `pg_curl_16` | `2.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_curl_16-2.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_curl_16-2.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_curl_16` | `2.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.8 KiB | [pg_curl_16-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_curl_16-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_16` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.8 KiB | [pg_curl_16-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_curl_16-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_16` | `2.4.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 122.6 KiB | [pg_curl_16-2.4.6-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_16-2.4.6-1PGSTY.el8.aarch64.rpm) |
| `pg_curl_16` | `2.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_curl_16-2.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_curl_16-2.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_curl_16` | `2.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_curl_16-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_curl_16-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_16` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.9 KiB | [pg_curl_16-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_curl_16-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_16` | `2.4.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 124.9 KiB | [pg_curl_16-2.4.6-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_16-2.4.6-1PGSTY.el9.x86_64.rpm) |
| `pg_curl_16` | `2.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.6 KiB | [pg_curl_16-2.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_curl_16-2.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_16` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.9 KiB | [pg_curl_16-2.4.4-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_curl_16-2.4.4-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_16` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.7 KiB | [pg_curl_16-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_curl_16-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_16` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.5 KiB | [pg_curl_16-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_curl_16-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_16` | `2.4.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 124.0 KiB | [pg_curl_16-2.4.6-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_16-2.4.6-1PGSTY.el9.aarch64.rpm) |
| `pg_curl_16` | `2.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.8 KiB | [pg_curl_16-2.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_curl_16-2.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_16` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.3 KiB | [pg_curl_16-2.4.4-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_curl_16-2.4.4-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_16` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [pg_curl_16-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_curl_16-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_16` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.1 KiB | [pg_curl_16-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_curl_16-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_16` | `2.4.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 126.0 KiB | [pg_curl_16-2.4.6-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_curl_16-2.4.6-1PGSTY.el10.x86_64.rpm) |
| `pg_curl_16` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.1 KiB | [pg_curl_16-2.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_curl_16-2.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_16` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_16-2.4.4-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_curl_16-2.4.4-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_16` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.6 KiB | [pg_curl_16-2.4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_curl_16-2.4.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_curl_16` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_16-2.4.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_curl_16-2.4.3-2PGDG.rhel10.x86_64.rpm) |
| `pg_curl_16` | `2.4.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 125.2 KiB | [pg_curl_16-2.4.6-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_curl_16-2.4.6-1PGSTY.el10.aarch64.rpm) |
| `pg_curl_16` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.0 KiB | [pg_curl_16-2.4.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_curl_16-2.4.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_curl_16` | `2.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.5 KiB | [pg_curl_16-2.4.4-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_curl_16-2.4.4-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_curl_16` | `2.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.2 KiB | [pg_curl_16-2.4.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_curl_16-2.4.4-1PGDG.rhel10.aarch64.rpm) |
| `pg_curl_16` | `2.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.0 KiB | [pg_curl_16-2.4.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_curl_16-2.4.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-curl` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 108.0 KiB | [postgresql-16-pg-curl_2.4.6-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.6-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-curl` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.5 KiB | [postgresql-16-pg-curl_2.4.6-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.6-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-curl` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 107.9 KiB | [postgresql-16-pg-curl_2.4.6-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.6-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-curl` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 106.2 KiB | [postgresql-16-pg-curl_2.4.6-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.6-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-curl` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.5 KiB | [postgresql-16-pg-curl_2.4.6-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.6-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-curl` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.8 KiB | [postgresql-16-pg-curl_2.4.6-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.6-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-curl` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 115.4 KiB | [postgresql-16-pg-curl_2.4.6-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.6-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-curl` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 114.8 KiB | [postgresql-16-pg-curl_2.4.6-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.6-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-curl` | `2.4.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 121.5 KiB | [postgresql-16-pg-curl_2.4.6-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.6-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-curl` | `2.4.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 120.8 KiB | [postgresql-16-pg-curl_2.4.6-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-curl/postgresql-16-pg-curl_2.4.6-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_curl_15` | `2.4.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 124.4 KiB | [pg_curl_15-2.4.6-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_15-2.4.6-1PGSTY.el8.x86_64.rpm) |
| `pg_curl_15` | `2.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_curl_15-2.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_curl_15-2.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_curl_15` | `2.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.8 KiB | [pg_curl_15-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_curl_15-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_15` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.7 KiB | [pg_curl_15-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_curl_15-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_15` | `2.4.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 122.6 KiB | [pg_curl_15-2.4.6-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_15-2.4.6-1PGSTY.el8.aarch64.rpm) |
| `pg_curl_15` | `2.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_curl_15-2.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_curl_15-2.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_curl_15` | `2.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_curl_15-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_curl_15-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_15` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.9 KiB | [pg_curl_15-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_curl_15-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_15` | `2.4.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 125.6 KiB | [pg_curl_15-2.4.6-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_15-2.4.6-1PGSTY.el9.x86_64.rpm) |
| `pg_curl_15` | `2.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.4 KiB | [pg_curl_15-2.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_curl_15-2.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_15` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.8 KiB | [pg_curl_15-2.4.4-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_curl_15-2.4.4-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_15` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.6 KiB | [pg_curl_15-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_curl_15-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_15` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.6 KiB | [pg_curl_15-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_curl_15-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_15` | `2.4.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 124.0 KiB | [pg_curl_15-2.4.6-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_15-2.4.6-1PGSTY.el9.aarch64.rpm) |
| `pg_curl_15` | `2.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.5 KiB | [pg_curl_15-2.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_curl_15-2.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_15` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.3 KiB | [pg_curl_15-2.4.4-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_curl_15-2.4.4-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_15` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [pg_curl_15-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_curl_15-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_15` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [pg_curl_15-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_curl_15-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_15` | `2.4.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 126.4 KiB | [pg_curl_15-2.4.6-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_curl_15-2.4.6-1PGSTY.el10.x86_64.rpm) |
| `pg_curl_15` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.1 KiB | [pg_curl_15-2.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_curl_15-2.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_15` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_15-2.4.4-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_curl_15-2.4.4-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_15` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.5 KiB | [pg_curl_15-2.4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_curl_15-2.4.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_curl_15` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_15-2.4.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_curl_15-2.4.3-2PGDG.rhel10.x86_64.rpm) |
| `pg_curl_15` | `2.4.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 125.6 KiB | [pg_curl_15-2.4.6-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_curl_15-2.4.6-1PGSTY.el10.aarch64.rpm) |
| `pg_curl_15` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.9 KiB | [pg_curl_15-2.4.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_curl_15-2.4.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_curl_15` | `2.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.4 KiB | [pg_curl_15-2.4.4-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_curl_15-2.4.4-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_curl_15` | `2.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.5 KiB | [pg_curl_15-2.4.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_curl_15-2.4.4-1PGDG.rhel10.aarch64.rpm) |
| `pg_curl_15` | `2.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.4 KiB | [pg_curl_15-2.4.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_curl_15-2.4.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-curl` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 107.0 KiB | [postgresql-15-pg-curl_2.4.6-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.6-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-curl` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.9 KiB | [postgresql-15-pg-curl_2.4.6-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.6-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-curl` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 107.1 KiB | [postgresql-15-pg-curl_2.4.6-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.6-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-curl` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 105.9 KiB | [postgresql-15-pg-curl_2.4.6-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.6-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-curl` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.6 KiB | [postgresql-15-pg-curl_2.4.6-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.6-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-curl` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.8 KiB | [postgresql-15-pg-curl_2.4.6-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.6-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-curl` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 115.5 KiB | [postgresql-15-pg-curl_2.4.6-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.6-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-curl` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 115.1 KiB | [postgresql-15-pg-curl_2.4.6-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.6-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-curl` | `2.4.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 121.3 KiB | [postgresql-15-pg-curl_2.4.6-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.6-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-curl` | `2.4.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 120.7 KiB | [postgresql-15-pg-curl_2.4.6-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-curl/postgresql-15-pg-curl_2.4.6-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_curl_14` | `2.4.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 124.3 KiB | [pg_curl_14-2.4.6-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_14-2.4.6-1PGSTY.el8.x86_64.rpm) |
| `pg_curl_14` | `2.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_curl_14-2.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_curl_14-2.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_curl_14` | `2.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.8 KiB | [pg_curl_14-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_curl_14-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_14` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.7 KiB | [pg_curl_14-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_curl_14-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_14` | `2.4.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 122.5 KiB | [pg_curl_14-2.4.6-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_14-2.4.6-1PGSTY.el8.aarch64.rpm) |
| `pg_curl_14` | `2.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_curl_14-2.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_curl_14-2.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_curl_14` | `2.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_curl_14-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_curl_14-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_14` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.9 KiB | [pg_curl_14-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_curl_14-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_14` | `2.4.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 125.6 KiB | [pg_curl_14-2.4.6-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_14-2.4.6-1PGSTY.el9.x86_64.rpm) |
| `pg_curl_14` | `2.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.4 KiB | [pg_curl_14-2.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_curl_14-2.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_14` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.8 KiB | [pg_curl_14-2.4.4-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_curl_14-2.4.4-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_14` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.5 KiB | [pg_curl_14-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_curl_14-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_14` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.5 KiB | [pg_curl_14-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_curl_14-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_14` | `2.4.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 124.0 KiB | [pg_curl_14-2.4.6-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_14-2.4.6-1PGSTY.el9.aarch64.rpm) |
| `pg_curl_14` | `2.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.9 KiB | [pg_curl_14-2.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_curl_14-2.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_14` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.3 KiB | [pg_curl_14-2.4.4-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_curl_14-2.4.4-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_14` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.8 KiB | [pg_curl_14-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_curl_14-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_14` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [pg_curl_14-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_curl_14-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_14` | `2.4.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 126.6 KiB | [pg_curl_14-2.4.6-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_curl_14-2.4.6-1PGSTY.el10.x86_64.rpm) |
| `pg_curl_14` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.2 KiB | [pg_curl_14-2.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_curl_14-2.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_14` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_14-2.4.4-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_curl_14-2.4.4-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_14` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.5 KiB | [pg_curl_14-2.4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_curl_14-2.4.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_curl_14` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_14-2.4.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_curl_14-2.4.3-2PGDG.rhel10.x86_64.rpm) |
| `pg_curl_14` | `2.4.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 125.5 KiB | [pg_curl_14-2.4.6-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_curl_14-2.4.6-1PGSTY.el10.aarch64.rpm) |
| `pg_curl_14` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.0 KiB | [pg_curl_14-2.4.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_curl_14-2.4.5-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_curl_14` | `2.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.3 KiB | [pg_curl_14-2.4.4-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_curl_14-2.4.4-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_curl_14` | `2.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.5 KiB | [pg_curl_14-2.4.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_curl_14-2.4.4-1PGDG.rhel10.aarch64.rpm) |
| `pg_curl_14` | `2.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.4 KiB | [pg_curl_14-2.4.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_curl_14-2.4.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-curl` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 107.3 KiB | [postgresql-14-pg-curl_2.4.6-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.6-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-curl` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.1 KiB | [postgresql-14-pg-curl_2.4.6-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.6-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-curl` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 107.7 KiB | [postgresql-14-pg-curl_2.4.6-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.6-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-curl` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 105.6 KiB | [postgresql-14-pg-curl_2.4.6-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.6-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-curl` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.8 KiB | [postgresql-14-pg-curl_2.4.6-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.6-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-curl` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.9 KiB | [postgresql-14-pg-curl_2.4.6-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.6-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-curl` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 115.6 KiB | [postgresql-14-pg-curl_2.4.6-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.6-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-curl` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 115.0 KiB | [postgresql-14-pg-curl_2.4.6-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.6-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-curl` | `2.4.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 121.4 KiB | [postgresql-14-pg-curl_2.4.6-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.6-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-curl` | `2.4.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 121.1 KiB | [postgresql-14-pg-curl_2.4.6-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-curl/postgresql-14-pg-curl_2.4.6-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/RekGRpth/pg_curl" title="Repository" icon="github" subtitle="github.com/RekGRpth/pg_curl" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_curl-2.4.6.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_curl;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_curl;		# install via package name, for the active PG version

pig install pg_curl -v 18;   # install for PG 18
pig install pg_curl -v 17;   # install for PG 17
pig install pg_curl -v 16;   # install for PG 16
pig install pg_curl -v 15;   # install for PG 15
pig install pg_curl -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_curl;
```




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
