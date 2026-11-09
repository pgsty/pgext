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
| **4090** | {{< badge content="pg_curl" link="https://github.com/RekGRpth/pg_curl" >}} | {{< ext "pg_curl" >}} | `2.5.0` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "omni_httpc" >}} {{< ext "pg_graphql" >}} {{< ext "documentdb" >}} |

> [!Note] SQL 2.5; no upstream 2.4.1 upgrade edge.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_curl` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.0` | {{< bg "18" "pg_curl_18" "green" >}} {{< bg "17" "pg_curl_17" "green" >}} {{< bg "16" "pg_curl_16" "green" >}} {{< bg "15" "pg_curl_15" "green" >}} {{< bg "14" "pg_curl_14" "green" >}} | `pg_curl_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.0` | {{< bg "18" "postgresql-18-pg-curl" "green" >}} {{< bg "17" "postgresql-17-pg-curl" "green" >}} {{< bg "16" "postgresql-16-pg-curl" "green" >}} {{< bg "15" "postgresql-15-pg-curl" "green" >}} {{< bg "14" "postgresql-14-pg-curl" "green" >}} | `postgresql-$v-pg-curl` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_14 : AVAIL 4" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_14 : AVAIL 4" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_14 : AVAIL 5" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_14 : AVAIL 5" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_14 : AVAIL 5" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_16 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "pg_curl_14 : AVAIL 5" "green" >}} |
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
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_curl_18` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 136.0 KiB | [pg_curl_18-2.5.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_18-2.5.0-1PGSTY.el8.x86_64.rpm) |
| `pg_curl_18` | `2.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_curl_18-2.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_curl_18-2.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_curl_18` | `2.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.8 KiB | [pg_curl_18-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_curl_18-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_18` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 133.6 KiB | [pg_curl_18-2.5.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_18-2.5.0-1PGSTY.el8.aarch64.rpm) |
| `pg_curl_18` | `2.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_curl_18-2.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_curl_18-2.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_curl_18` | `2.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_curl_18-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_curl_18-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_18` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 139.2 KiB | [pg_curl_18-2.5.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_18-2.5.0-1PGSTY.el9.x86_64.rpm) |
| `pg_curl_18` | `2.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.5 KiB | [pg_curl_18-2.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_curl_18-2.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_18` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.8 KiB | [pg_curl_18-2.4.4-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_curl_18-2.4.4-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_18` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.5 KiB | [pg_curl_18-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_curl_18-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_18` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 137.4 KiB | [pg_curl_18-2.5.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_18-2.5.0-1PGSTY.el9.aarch64.rpm) |
| `pg_curl_18` | `2.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.9 KiB | [pg_curl_18-2.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_curl_18-2.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_18` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.3 KiB | [pg_curl_18-2.4.4-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_curl_18-2.4.4-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_18` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [pg_curl_18-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_curl_18-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_18` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 141.9 KiB | [pg_curl_18-2.5.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_curl_18-2.5.0-1PGSTY.el10.x86_64.rpm) |
| `pg_curl_18` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.1 KiB | [pg_curl_18-2.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_curl_18-2.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_18` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_18-2.4.4-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_curl_18-2.4.4-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_18` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.5 KiB | [pg_curl_18-2.4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_curl_18-2.4.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_curl_18` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 140.4 KiB | [pg_curl_18-2.5.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_curl_18-2.5.0-1PGSTY.el10.aarch64.rpm) |
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
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_curl_17` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 136.0 KiB | [pg_curl_17-2.5.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_17-2.5.0-1PGSTY.el8.x86_64.rpm) |
| `pg_curl_17` | `2.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_curl_17-2.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_curl_17-2.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_curl_17` | `2.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.8 KiB | [pg_curl_17-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_curl_17-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_17` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.7 KiB | [pg_curl_17-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_curl_17-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_17` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 133.5 KiB | [pg_curl_17-2.5.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_17-2.5.0-1PGSTY.el8.aarch64.rpm) |
| `pg_curl_17` | `2.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_curl_17-2.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_curl_17-2.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_curl_17` | `2.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_curl_17-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_curl_17-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_17` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.9 KiB | [pg_curl_17-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_curl_17-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_17` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 139.4 KiB | [pg_curl_17-2.5.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_17-2.5.0-1PGSTY.el9.x86_64.rpm) |
| `pg_curl_17` | `2.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.6 KiB | [pg_curl_17-2.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_curl_17-2.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_17` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.9 KiB | [pg_curl_17-2.4.4-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_curl_17-2.4.4-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_17` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.6 KiB | [pg_curl_17-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_curl_17-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_17` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.6 KiB | [pg_curl_17-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_curl_17-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_17` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 137.5 KiB | [pg_curl_17-2.5.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_17-2.5.0-1PGSTY.el9.aarch64.rpm) |
| `pg_curl_17` | `2.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.8 KiB | [pg_curl_17-2.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_curl_17-2.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_17` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.1 KiB | [pg_curl_17-2.4.4-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_curl_17-2.4.4-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_17` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.8 KiB | [pg_curl_17-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_curl_17-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_17` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [pg_curl_17-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_curl_17-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_17` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 141.8 KiB | [pg_curl_17-2.5.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_curl_17-2.5.0-1PGSTY.el10.x86_64.rpm) |
| `pg_curl_17` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.1 KiB | [pg_curl_17-2.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_curl_17-2.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_17` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_17-2.4.4-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_curl_17-2.4.4-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_17` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.6 KiB | [pg_curl_17-2.4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_curl_17-2.4.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_curl_17` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_17-2.4.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_curl_17-2.4.3-2PGDG.rhel10.x86_64.rpm) |
| `pg_curl_17` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 140.3 KiB | [pg_curl_17-2.5.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_curl_17-2.5.0-1PGSTY.el10.aarch64.rpm) |
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
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_curl_16` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 136.0 KiB | [pg_curl_16-2.5.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_16-2.5.0-1PGSTY.el8.x86_64.rpm) |
| `pg_curl_16` | `2.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_curl_16-2.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_curl_16-2.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_curl_16` | `2.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.8 KiB | [pg_curl_16-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_curl_16-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_16` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.8 KiB | [pg_curl_16-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_curl_16-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_16` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 133.5 KiB | [pg_curl_16-2.5.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_16-2.5.0-1PGSTY.el8.aarch64.rpm) |
| `pg_curl_16` | `2.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_curl_16-2.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_curl_16-2.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_curl_16` | `2.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_curl_16-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_curl_16-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_16` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.9 KiB | [pg_curl_16-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_curl_16-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_16` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 139.4 KiB | [pg_curl_16-2.5.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_16-2.5.0-1PGSTY.el9.x86_64.rpm) |
| `pg_curl_16` | `2.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.6 KiB | [pg_curl_16-2.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_curl_16-2.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_16` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.9 KiB | [pg_curl_16-2.4.4-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_curl_16-2.4.4-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_16` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.7 KiB | [pg_curl_16-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_curl_16-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_16` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.5 KiB | [pg_curl_16-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_curl_16-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_16` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 137.5 KiB | [pg_curl_16-2.5.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_16-2.5.0-1PGSTY.el9.aarch64.rpm) |
| `pg_curl_16` | `2.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.8 KiB | [pg_curl_16-2.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_curl_16-2.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_16` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.3 KiB | [pg_curl_16-2.4.4-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_curl_16-2.4.4-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_16` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [pg_curl_16-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_curl_16-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_16` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.1 KiB | [pg_curl_16-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_curl_16-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_16` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 141.8 KiB | [pg_curl_16-2.5.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_curl_16-2.5.0-1PGSTY.el10.x86_64.rpm) |
| `pg_curl_16` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.1 KiB | [pg_curl_16-2.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_curl_16-2.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_16` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_16-2.4.4-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_curl_16-2.4.4-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_16` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.6 KiB | [pg_curl_16-2.4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_curl_16-2.4.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_curl_16` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_16-2.4.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_curl_16-2.4.3-2PGDG.rhel10.x86_64.rpm) |
| `pg_curl_16` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 140.3 KiB | [pg_curl_16-2.5.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_curl_16-2.5.0-1PGSTY.el10.aarch64.rpm) |
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
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_curl_15` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 135.7 KiB | [pg_curl_15-2.5.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_15-2.5.0-1PGSTY.el8.x86_64.rpm) |
| `pg_curl_15` | `2.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_curl_15-2.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_curl_15-2.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_curl_15` | `2.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.8 KiB | [pg_curl_15-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_curl_15-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_15` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.7 KiB | [pg_curl_15-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_curl_15-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_15` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 133.7 KiB | [pg_curl_15-2.5.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_15-2.5.0-1PGSTY.el8.aarch64.rpm) |
| `pg_curl_15` | `2.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_curl_15-2.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_curl_15-2.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_curl_15` | `2.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_curl_15-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_curl_15-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_15` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.9 KiB | [pg_curl_15-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_curl_15-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_15` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 139.7 KiB | [pg_curl_15-2.5.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_15-2.5.0-1PGSTY.el9.x86_64.rpm) |
| `pg_curl_15` | `2.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.4 KiB | [pg_curl_15-2.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_curl_15-2.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_15` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.8 KiB | [pg_curl_15-2.4.4-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_curl_15-2.4.4-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_15` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.6 KiB | [pg_curl_15-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_curl_15-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_15` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.6 KiB | [pg_curl_15-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_curl_15-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_15` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 138.0 KiB | [pg_curl_15-2.5.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_15-2.5.0-1PGSTY.el9.aarch64.rpm) |
| `pg_curl_15` | `2.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.5 KiB | [pg_curl_15-2.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_curl_15-2.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_15` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.3 KiB | [pg_curl_15-2.4.4-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_curl_15-2.4.4-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_15` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [pg_curl_15-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_curl_15-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_15` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [pg_curl_15-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_curl_15-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_15` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 142.0 KiB | [pg_curl_15-2.5.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_curl_15-2.5.0-1PGSTY.el10.x86_64.rpm) |
| `pg_curl_15` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.1 KiB | [pg_curl_15-2.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_curl_15-2.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_15` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_15-2.4.4-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_curl_15-2.4.4-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_15` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.5 KiB | [pg_curl_15-2.4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_curl_15-2.4.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_curl_15` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_15-2.4.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_curl_15-2.4.3-2PGDG.rhel10.x86_64.rpm) |
| `pg_curl_15` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 140.8 KiB | [pg_curl_15-2.5.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_curl_15-2.5.0-1PGSTY.el10.aarch64.rpm) |
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
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_curl_14` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 135.6 KiB | [pg_curl_14-2.5.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_curl_14-2.5.0-1PGSTY.el8.x86_64.rpm) |
| `pg_curl_14` | `2.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_curl_14-2.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_curl_14-2.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_curl_14` | `2.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.8 KiB | [pg_curl_14-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_curl_14-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_14` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.7 KiB | [pg_curl_14-2.4.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_curl_14-2.4.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_curl_14` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 133.7 KiB | [pg_curl_14-2.5.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_curl_14-2.5.0-1PGSTY.el8.aarch64.rpm) |
| `pg_curl_14` | `2.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_curl_14-2.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_curl_14-2.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_curl_14` | `2.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_curl_14-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_curl_14-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_14` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.9 KiB | [pg_curl_14-2.4.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_curl_14-2.4.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_curl_14` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 139.8 KiB | [pg_curl_14-2.5.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_curl_14-2.5.0-1PGSTY.el9.x86_64.rpm) |
| `pg_curl_14` | `2.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.4 KiB | [pg_curl_14-2.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_curl_14-2.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_14` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.8 KiB | [pg_curl_14-2.4.4-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_curl_14-2.4.4-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_curl_14` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.5 KiB | [pg_curl_14-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_curl_14-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_14` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.5 KiB | [pg_curl_14-2.4.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_curl_14-2.4.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_curl_14` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 138.0 KiB | [pg_curl_14-2.5.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_curl_14-2.5.0-1PGSTY.el9.aarch64.rpm) |
| `pg_curl_14` | `2.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.9 KiB | [pg_curl_14-2.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_curl_14-2.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_14` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.3 KiB | [pg_curl_14-2.4.4-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_curl_14-2.4.4-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_curl_14` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.8 KiB | [pg_curl_14-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_curl_14-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_14` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.0 KiB | [pg_curl_14-2.4.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_curl_14-2.4.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_curl_14` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 142.0 KiB | [pg_curl_14-2.5.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_curl_14-2.5.0-1PGSTY.el10.x86_64.rpm) |
| `pg_curl_14` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 47.2 KiB | [pg_curl_14-2.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_curl_14-2.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_14` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_14-2.4.4-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_curl_14-2.4.4-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_curl_14` | `2.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.5 KiB | [pg_curl_14-2.4.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_curl_14-2.4.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_curl_14` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.4 KiB | [pg_curl_14-2.4.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_curl_14-2.4.3-2PGDG.rhel10.x86_64.rpm) |
| `pg_curl_14` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 140.8 KiB | [pg_curl_14-2.5.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_curl_14-2.5.0-1PGSTY.el10.aarch64.rpm) |
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
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/RekGRpth/pg_curl" title="Repository" icon="github" subtitle="github.com/RekGRpth/pg_curl" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_curl-2.5.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_curl;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_curl;		# install via package name, for the active PG version

pig install pg_curl -v 18;   # install for PG 18
pig install pg_curl -v 17;   # install for PG 17
pig install pg_curl -v 16;   # install for PG 16
pig install pg_curl -v 15;   # install for PG 15
pig install pg_curl -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_curl;
```

## Usage

Sources:

- [Official source](https://api.pgxn.org/src/pg_curl/pg_curl-2.5.0/README.md)
- [Official source](https://api.pgxn.org/src/pg_curl/pg_curl-2.5.0/pg_curl.control)
- [Official source](https://api.pgxn.org/src/pg_curl/pg_curl-2.5.0/pg_curl--2.5.sql)
- [Official source](https://api.pgxn.org/src/pg_curl/pg_curl-2.5.0/pg_curl--2.4--2.5.sql)
- [Official source](https://api.pgxn.org/dist/pg_curl/2.5.0/META.json)

`pg_curl` exposes libcurl transfers as SQL functions. Upstream release 2.5.0 uses SQL extension version 2.5. Requests run in the database backend with the server's network and filesystem access; installing the extension requires an administrator. Ordinary calls need no preload.

### A Request in One Transaction

```sql
CREATE EXTENSION pg_curl;
BEGIN;
SELECT curl_easy_reset();
SELECT curl_easy_setopt_url('https://www.postgresql.org/');
SELECT curl_easy_perform();
SELECT convert_from(curl_easy_getinfo_data_in(), 'UTF8');
COMMIT;
```

`curl_easy_reset` clears the handle, `curl_easy_setopt_url` chooses the destination, and `curl_easy_perform` performs the transfer. `curl_easy_getinfo_data_in` returns response bytes and `curl_easy_getinfo_header_in` returns headers. `pg_curl.transaction` defaults to true: handles are released at transaction end, so keep a multi-statement request in one transaction. Network/file side effects are not rolled back with SQL.

### Access and Whitelists

A superuser grants allowed prefixes to an existing application role:

```sql
ALTER ROLE app_user SET pg_curl.whitelist = 'https://api.example.com/';
```

Reconnect to apply a role default. With an unset `pg_curl.whitelist`, superusers are unrestricted and other roles are denied. A configured whitelist also restricts superusers. Only matching HTTP, HTTPS and file URLs are allowed while it applies; FTP, SMTP and other protocols are denied. HTTP paths match at segment boundaries, not arbitrary string prefixes.

Local-file options are checked too. File paths are resolved and must already exist, including an output cookie-jar file. Redirect targets are checked; restricted redirect following requires libcurl 7.80.0 or later. Explicit proxy and DNS-over-HTTPS options are denied while restricted, although the server's proxy environment can still apply. Limit allowed hosts and paths deliberately; stored credentials and server-local files remain sensitive.

### Other Operations and Upgrade

`curl_postfield_append`, `curl_easy_setopt_postfields` and `curl_header_append` construct form or JSON requests. `curl_mime_data` and `curl_mime_file` construct MIME parts; 2.5 adds a `head text[]` overload for multiple headers. Header array elements cannot be NULL or contain carriage returns/newlines. The MIME data overload accepts text or binary data.

Check `pg_extension.extversion` before upgrading. This distribution includes the 2.4-to-2.5 update script; it does not supply a 2.4.1-to-2.5 edge, so do not assume an existing installation reporting 2.4.1 can use an automatic update. For an installation actually at SQL 2.4, after installing matching files:

```sql
ALTER EXTENSION pg_curl UPDATE TO '2.5';
```

Use the new whitelist rules before granting execution to application roles. Review libcurl's supported protocols and use the matching library on all database hosts.
