---
title: "http"
linkTitle: "http"
description: "HTTP client for PostgreSQL, allows web page retrieval inside the database."
weight: 4070
categories: ["UTIL"]
width: full
---

[**pg_http**](https://github.com/pramsey/pgsql-http) : HTTP client for PostgreSQL, allows web page retrieval inside the database.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4070** | {{< badge content="http" link="https://github.com/pramsey/pgsql-http" >}} | {{< ext "http" "pg_http" >}} | `1.7.1` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pgjwt" >}} {{< ext "pg_smtp_client" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "pgjq" >}} {{< ext "pgmb" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.7.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_http` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.1` | {{< bg "18" "pgsql_http_18" "green" >}} {{< bg "17" "pgsql_http_17" "green" >}} {{< bg "16" "pgsql_http_16" "green" >}} {{< bg "15" "pgsql_http_15" "green" >}} {{< bg "14" "pgsql_http_14" "green" >}} | `pgsql_http_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.7.1` | {{< bg "18" "postgresql-18-http" "green" >}} {{< bg "17" "postgresql-17-http" "green" >}} {{< bg "16" "postgresql-16-http" "green" >}} {{< bg "15" "postgresql-15-http" "green" >}} {{< bg "14" "postgresql-14-http" "green" >}} | `postgresql-$v-http` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_15 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_14 : AVAIL 6" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_15 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_14 : AVAIL 6" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_14 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_14 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_14 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.7.1" "postgresql-18-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-16-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-14-http : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.7.1" "postgresql-18-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-16-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-14-http : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.7.1" "postgresql-18-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-16-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-14-http : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.7.1" "postgresql-18-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-16-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-14-http : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.7.1" "postgresql-18-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-16-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-14-http : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.7.1" "postgresql-18-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-16-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-14-http : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.7.1" "postgresql-18-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-16-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-14-http : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.7.1" "postgresql-18-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-16-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-14-http : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.7.1" "postgresql-18-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-16-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-14-http : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.7.1" "postgresql-18-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-17-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-16-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-http : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-14-http : AVAIL 3" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_http_18` | `1.7.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.6 KiB | [pgsql_http_18-1.7.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_http_18-1.7.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsql_http_18` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.1 KiB | [pgsql_http_18-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsql_http_18-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_18` | `1.7.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.8 KiB | [pgsql_http_18-1.7.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_http_18-1.7.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsql_http_18` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.2 KiB | [pgsql_http_18-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsql_http_18-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_18` | `1.7.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.5 KiB | [pgsql_http_18-1.7.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_http_18-1.7.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsql_http_18` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.2 KiB | [pgsql_http_18-1.7.0-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsql_http_18-1.7.0-3PGDG.rhel9.8.x86_64.rpm) |
| `pgsql_http_18` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.6 KiB | [pgsql_http_18-1.7.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_http_18-1.7.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_http_18` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pgsql_http_18-1.7.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsql_http_18-1.7.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pgsql_http_18` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [pgsql_http_18-1.7.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_http_18-1.7.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_http_18` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.7 KiB | [pgsql_http_18-1.7.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsql_http_18-1.7.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pgsql_http_18` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.8 KiB | [pgsql_http_18-1.7.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_http_18-1.7.1-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_http_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [pgsql_http_18-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsql_http_18-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-http` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.0 KiB | [postgresql-18-http_1.7.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-http` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 44.3 KiB | [postgresql-18-http_1.7.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-http` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.5 KiB | [postgresql-18-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-18-http` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.5 KiB | [postgresql-18-http_1.7.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-http` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.8 KiB | [postgresql-18-http_1.7.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-http` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.1 KiB | [postgresql-18-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-18-http` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.2 KiB | [postgresql-18-http_1.7.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-http` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 44.4 KiB | [postgresql-18-http_1.7.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-http` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.6 KiB | [postgresql-18-http_1.7.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg13+1_amd64.deb) |
| `postgresql-18-http` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.7 KiB | [postgresql-18-http_1.7.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-http` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 43.1 KiB | [postgresql-18-http_1.7.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-http` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.2 KiB | [postgresql-18-http_1.7.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg13+1_arm64.deb) |
| `postgresql-18-http` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 45.2 KiB | [postgresql-18-http_1.7.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-http` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 47.2 KiB | [postgresql-18-http_1.7.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-http` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 44.7 KiB | [postgresql-18-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-http` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 43.5 KiB | [postgresql-18-http_1.7.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-http` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 45.8 KiB | [postgresql-18-http_1.7.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-http` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 43.1 KiB | [postgresql-18-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-http` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.0 KiB | [postgresql-18-http_1.7.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-http` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 46.3 KiB | [postgresql-18-http_1.7.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-http` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.5 KiB | [postgresql-18-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-http` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.6 KiB | [postgresql-18-http_1.7.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-http` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 45.0 KiB | [postgresql-18-http_1.7.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-http` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.0 KiB | [postgresql-18-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg24.04+1_arm64.deb) |
| `postgresql-18-http` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 51.6 KiB | [postgresql-18-http_1.7.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-http` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 54.0 KiB | [postgresql-18-http_1.7.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-http` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 51.2 KiB | [postgresql-18-http_1.7.0-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg26.04+1_amd64.deb) |
| `postgresql-18-http` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 50.2 KiB | [postgresql-18-http_1.7.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-http` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 52.9 KiB | [postgresql-18-http_1.7.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-http/postgresql-18-http_1.7.1-1PIGSTY~resolute_arm64.deb) |
| `postgresql-18-http` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 49.9 KiB | [postgresql-18-http_1.7.0-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-18-http_1.7.0-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_http_17` | `1.7.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.5 KiB | [pgsql_http_17-1.7.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_http_17-1.7.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsql_http_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.1 KiB | [pgsql_http_17-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_http_17-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_17` | `1.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.2 KiB | [pgsql_http_17-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_http_17-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_17` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.2 KiB | [pgsql_http_17-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_http_17-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_17` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.8 KiB | [pgsql_http_17-1.6.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsql_http_17-1.6.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_17` | `1.7.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.8 KiB | [pgsql_http_17-1.7.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_http_17-1.7.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsql_http_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.2 KiB | [pgsql_http_17-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_http_17-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_17` | `1.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.3 KiB | [pgsql_http_17-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_http_17-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_17` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.5 KiB | [pgsql_http_17-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_http_17-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_17` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.1 KiB | [pgsql_http_17-1.6.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsql_http_17-1.6.0-2PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_17` | `1.7.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.7 KiB | [pgsql_http_17-1.7.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_http_17-1.7.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsql_http_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.4 KiB | [pgsql_http_17-1.7.0-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_http_17-1.7.0-3PGDG.rhel9.8.x86_64.rpm) |
| `pgsql_http_17` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.7 KiB | [pgsql_http_17-1.7.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_http_17-1.7.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_http_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pgsql_http_17-1.7.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_http_17-1.7.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pgsql_http_17` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [pgsql_http_17-1.7.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_http_17-1.7.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_http_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.8 KiB | [pgsql_http_17-1.7.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsql_http_17-1.7.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pgsql_http_17` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.9 KiB | [pgsql_http_17-1.7.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_http_17-1.7.1-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_http_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgsql_http_17-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_http_17-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-http` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.8 KiB | [postgresql-17-http_1.7.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-http` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 44.2 KiB | [postgresql-17-http_1.7.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-http` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.6 KiB | [postgresql-17-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-17-http` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.5 KiB | [postgresql-17-http_1.7.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-http` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.8 KiB | [postgresql-17-http_1.7.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-http` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.1 KiB | [postgresql-17-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-17-http` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.0 KiB | [postgresql-17-http_1.7.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-http` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 44.3 KiB | [postgresql-17-http_1.7.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-http` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.6 KiB | [postgresql-17-http_1.7.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg13+1_amd64.deb) |
| `postgresql-17-http` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.6 KiB | [postgresql-17-http_1.7.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-http` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 42.9 KiB | [postgresql-17-http_1.7.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-http` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.3 KiB | [postgresql-17-http_1.7.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg13+1_arm64.deb) |
| `postgresql-17-http` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 49.5 KiB | [postgresql-17-http_1.7.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-http` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 51.5 KiB | [postgresql-17-http_1.7.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-http` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 48.9 KiB | [postgresql-17-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-http` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 47.8 KiB | [postgresql-17-http_1.7.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-http` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 50.1 KiB | [postgresql-17-http_1.7.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-http` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 47.4 KiB | [postgresql-17-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-http` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.0 KiB | [postgresql-17-http_1.7.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-http` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 46.3 KiB | [postgresql-17-http_1.7.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-http` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.6 KiB | [postgresql-17-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-http` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.6 KiB | [postgresql-17-http_1.7.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-http` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 45.1 KiB | [postgresql-17-http_1.7.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-http` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.2 KiB | [postgresql-17-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-http` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 51.7 KiB | [postgresql-17-http_1.7.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-http` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 54.1 KiB | [postgresql-17-http_1.7.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-http` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 51.5 KiB | [postgresql-17-http_1.7.0-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg26.04+1_amd64.deb) |
| `postgresql-17-http` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 50.3 KiB | [postgresql-17-http_1.7.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-http` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.0 KiB | [postgresql-17-http_1.7.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-http/postgresql-17-http_1.7.1-1PIGSTY~resolute_arm64.deb) |
| `postgresql-17-http` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 50.0 KiB | [postgresql-17-http_1.7.0-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-17-http_1.7.0-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_http_16` | `1.7.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.4 KiB | [pgsql_http_16-1.7.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_http_16-1.7.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsql_http_16` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.1 KiB | [pgsql_http_16-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_http_16-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_16` | `1.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.2 KiB | [pgsql_http_16-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_http_16-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_16` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.2 KiB | [pgsql_http_16-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_http_16-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_16` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.8 KiB | [pgsql_http_16-1.6.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_http_16-1.6.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_16` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.7 KiB | [pgsql_http_16-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsql_http_16-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_16` | `1.7.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.7 KiB | [pgsql_http_16-1.7.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_http_16-1.7.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsql_http_16` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.2 KiB | [pgsql_http_16-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_http_16-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_16` | `1.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.3 KiB | [pgsql_http_16-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_http_16-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_16` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.4 KiB | [pgsql_http_16-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_http_16-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_16` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.1 KiB | [pgsql_http_16-1.6.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_http_16-1.6.0-2PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_16` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.9 KiB | [pgsql_http_16-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsql_http_16-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_16` | `1.7.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.6 KiB | [pgsql_http_16-1.7.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_http_16-1.7.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsql_http_16` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.4 KiB | [pgsql_http_16-1.7.0-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_http_16-1.7.0-3PGDG.rhel9.8.x86_64.rpm) |
| `pgsql_http_16` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.5 KiB | [pgsql_http_16-1.7.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_http_16-1.7.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_http_16` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pgsql_http_16-1.7.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.7.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pgsql_http_16` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [pgsql_http_16-1.7.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_http_16-1.7.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_http_16` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.9 KiB | [pgsql_http_16-1.7.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsql_http_16-1.7.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pgsql_http_16` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.8 KiB | [pgsql_http_16-1.7.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_http_16-1.7.1-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_http_16` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgsql_http_16-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_http_16-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-http` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.7 KiB | [postgresql-16-http_1.7.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-http` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 43.9 KiB | [postgresql-16-http_1.7.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-http` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.5 KiB | [postgresql-16-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-16-http` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.2 KiB | [postgresql-16-http_1.7.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-http` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 42.5 KiB | [postgresql-16-http_1.7.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-http` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.1 KiB | [postgresql-16-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-16-http` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.7 KiB | [postgresql-16-http_1.7.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1.pgdg13+1_amd64.deb) |
| `postgresql-16-http` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 44.1 KiB | [postgresql-16-http_1.7.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-http` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.6 KiB | [postgresql-16-http_1.7.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg13+1_amd64.deb) |
| `postgresql-16-http` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.4 KiB | [postgresql-16-http_1.7.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1.pgdg13+1_arm64.deb) |
| `postgresql-16-http` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 42.7 KiB | [postgresql-16-http_1.7.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-http` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.3 KiB | [postgresql-16-http_1.7.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg13+1_arm64.deb) |
| `postgresql-16-http` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 49.0 KiB | [postgresql-16-http_1.7.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-http` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 51.1 KiB | [postgresql-16-http_1.7.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-http` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 48.9 KiB | [postgresql-16-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-http` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 47.3 KiB | [postgresql-16-http_1.7.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-http` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 49.7 KiB | [postgresql-16-http_1.7.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-http` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 47.3 KiB | [postgresql-16-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-http` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.7 KiB | [postgresql-16-http_1.7.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-http` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 46.1 KiB | [postgresql-16-http_1.7.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-http` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.6 KiB | [postgresql-16-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-http` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.2 KiB | [postgresql-16-http_1.7.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-http` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 44.7 KiB | [postgresql-16-http_1.7.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-http` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.2 KiB | [postgresql-16-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg24.04+1_arm64.deb) |
| `postgresql-16-http` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 51.5 KiB | [postgresql-16-http_1.7.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-http` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 54.0 KiB | [postgresql-16-http_1.7.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-http` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 51.5 KiB | [postgresql-16-http_1.7.0-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg26.04+1_amd64.deb) |
| `postgresql-16-http` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 50.0 KiB | [postgresql-16-http_1.7.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-http` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 52.7 KiB | [postgresql-16-http_1.7.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-http/postgresql-16-http_1.7.1-1PIGSTY~resolute_arm64.deb) |
| `postgresql-16-http` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 50.1 KiB | [postgresql-16-http_1.7.0-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-16-http_1.7.0-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_http_15` | `1.7.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.8 KiB | [pgsql_http_15-1.7.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_http_15-1.7.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsql_http_15` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [pgsql_http_15-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_http_15-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_15` | `1.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.3 KiB | [pgsql_http_15-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_http_15-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_15` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.5 KiB | [pgsql_http_15-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_http_15-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_15` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.1 KiB | [pgsql_http_15-1.6.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_http_15-1.6.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_15` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.9 KiB | [pgsql_http_15-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsql_http_15-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_15` | `1.7.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.0 KiB | [pgsql_http_15-1.7.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_http_15-1.7.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsql_http_15` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.5 KiB | [pgsql_http_15-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_http_15-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_15` | `1.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.4 KiB | [pgsql_http_15-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_http_15-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_15` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.7 KiB | [pgsql_http_15-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_http_15-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_15` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.3 KiB | [pgsql_http_15-1.6.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_http_15-1.6.0-2PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_15` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.2 KiB | [pgsql_http_15-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsql_http_15-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_15` | `1.7.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [pgsql_http_15-1.7.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_http_15-1.7.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsql_http_15` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.7 KiB | [pgsql_http_15-1.7.0-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_http_15-1.7.0-3PGDG.rhel9.8.x86_64.rpm) |
| `pgsql_http_15` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pgsql_http_15-1.7.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_http_15-1.7.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_http_15` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pgsql_http_15-1.7.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.7.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pgsql_http_15` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.3 KiB | [pgsql_http_15-1.7.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_http_15-1.7.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_http_15` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.2 KiB | [pgsql_http_15-1.7.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsql_http_15-1.7.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pgsql_http_15` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.1 KiB | [pgsql_http_15-1.7.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_http_15-1.7.1-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_http_15` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.7 KiB | [pgsql_http_15-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_http_15-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-http` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.4 KiB | [postgresql-15-http_1.7.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-http` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 44.9 KiB | [postgresql-15-http_1.7.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-http` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.4 KiB | [postgresql-15-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-15-http` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.2 KiB | [postgresql-15-http_1.7.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-http` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 43.5 KiB | [postgresql-15-http_1.7.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-http` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.1 KiB | [postgresql-15-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-15-http` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.8 KiB | [postgresql-15-http_1.7.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1.pgdg13+1_amd64.deb) |
| `postgresql-15-http` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.1 KiB | [postgresql-15-http_1.7.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-http` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.7 KiB | [postgresql-15-http_1.7.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg13+1_amd64.deb) |
| `postgresql-15-http` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.3 KiB | [postgresql-15-http_1.7.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1.pgdg13+1_arm64.deb) |
| `postgresql-15-http` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 43.6 KiB | [postgresql-15-http_1.7.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-http` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.1 KiB | [postgresql-15-http_1.7.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg13+1_arm64.deb) |
| `postgresql-15-http` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.3 KiB | [postgresql-15-http_1.7.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-http` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 52.4 KiB | [postgresql-15-http_1.7.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-http` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.0 KiB | [postgresql-15-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-http` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 48.3 KiB | [postgresql-15-http_1.7.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-http` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 50.7 KiB | [postgresql-15-http_1.7.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-http` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 48.1 KiB | [postgresql-15-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-http` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.7 KiB | [postgresql-15-http_1.7.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-http` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 47.0 KiB | [postgresql-15-http_1.7.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-http` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.5 KiB | [postgresql-15-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-http` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.1 KiB | [postgresql-15-http_1.7.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-http` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 45.6 KiB | [postgresql-15-http_1.7.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-http` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.1 KiB | [postgresql-15-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg24.04+1_arm64.deb) |
| `postgresql-15-http` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 52.6 KiB | [postgresql-15-http_1.7.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-http` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 55.0 KiB | [postgresql-15-http_1.7.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-http` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 52.6 KiB | [postgresql-15-http_1.7.0-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg26.04+1_amd64.deb) |
| `postgresql-15-http` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 51.0 KiB | [postgresql-15-http_1.7.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-http` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.7 KiB | [postgresql-15-http_1.7.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-http/postgresql-15-http_1.7.1-1PIGSTY~resolute_arm64.deb) |
| `postgresql-15-http` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 50.8 KiB | [postgresql-15-http_1.7.0-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-15-http_1.7.0-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsql_http_14` | `1.7.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.8 KiB | [pgsql_http_14-1.7.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsql_http_14-1.7.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsql_http_14` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [pgsql_http_14-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_http_14-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_14` | `1.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.3 KiB | [pgsql_http_14-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_http_14-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_14` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.5 KiB | [pgsql_http_14-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_http_14-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_14` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.0 KiB | [pgsql_http_14-1.6.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_http_14-1.6.0-2PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_14` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.9 KiB | [pgsql_http_14-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsql_http_14-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsql_http_14` | `1.7.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.0 KiB | [pgsql_http_14-1.7.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsql_http_14-1.7.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsql_http_14` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.5 KiB | [pgsql_http_14-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_http_14-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_14` | `1.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.4 KiB | [pgsql_http_14-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_http_14-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_14` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.7 KiB | [pgsql_http_14-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_http_14-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_14` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.3 KiB | [pgsql_http_14-1.6.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_http_14-1.6.0-2PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_14` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.2 KiB | [pgsql_http_14-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsql_http_14-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsql_http_14` | `1.7.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [pgsql_http_14-1.7.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsql_http_14-1.7.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsql_http_14` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.7 KiB | [pgsql_http_14-1.7.0-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_http_14-1.7.0-3PGDG.rhel9.8.x86_64.rpm) |
| `pgsql_http_14` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pgsql_http_14-1.7.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_http_14-1.7.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_http_14` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pgsql_http_14-1.7.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.7.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pgsql_http_14` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.3 KiB | [pgsql_http_14-1.7.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_http_14-1.7.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_http_14` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.0 KiB | [pgsql_http_14-1.7.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsql_http_14-1.7.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pgsql_http_14` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.1 KiB | [pgsql_http_14-1.7.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_http_14-1.7.1-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_http_14` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.7 KiB | [pgsql_http_14-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsql_http_14-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-http` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.5 KiB | [postgresql-14-http_1.7.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-http` | `1.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 44.8 KiB | [postgresql-14-http_1.7.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-http` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.3 KiB | [postgresql-14-http_1.7.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg12+1_amd64.deb) |
| `postgresql-14-http` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.2 KiB | [postgresql-14-http_1.7.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-http` | `1.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 43.5 KiB | [postgresql-14-http_1.7.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-http` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.0 KiB | [postgresql-14-http_1.7.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg12+1_arm64.deb) |
| `postgresql-14-http` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.8 KiB | [postgresql-14-http_1.7.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1.pgdg13+1_amd64.deb) |
| `postgresql-14-http` | `1.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.0 KiB | [postgresql-14-http_1.7.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-http` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.6 KiB | [postgresql-14-http_1.7.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg13+1_amd64.deb) |
| `postgresql-14-http` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.3 KiB | [postgresql-14-http_1.7.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1.pgdg13+1_arm64.deb) |
| `postgresql-14-http` | `1.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 43.6 KiB | [postgresql-14-http_1.7.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-http` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.1 KiB | [postgresql-14-http_1.7.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg13+1_arm64.deb) |
| `postgresql-14-http` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.3 KiB | [postgresql-14-http_1.7.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-http` | `1.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 52.4 KiB | [postgresql-14-http_1.7.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-http` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.0 KiB | [postgresql-14-http_1.7.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-http` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 48.3 KiB | [postgresql-14-http_1.7.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-http` | `1.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 50.7 KiB | [postgresql-14-http_1.7.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-http` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 48.1 KiB | [postgresql-14-http_1.7.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-http` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.7 KiB | [postgresql-14-http_1.7.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-http` | `1.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 47.0 KiB | [postgresql-14-http_1.7.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-http` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.6 KiB | [postgresql-14-http_1.7.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-http` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.2 KiB | [postgresql-14-http_1.7.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-http` | `1.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 45.6 KiB | [postgresql-14-http_1.7.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-http` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.1 KiB | [postgresql-14-http_1.7.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg24.04+1_arm64.deb) |
| `postgresql-14-http` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 52.5 KiB | [postgresql-14-http_1.7.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-http` | `1.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 54.9 KiB | [postgresql-14-http_1.7.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-http` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 52.4 KiB | [postgresql-14-http_1.7.0-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg26.04+1_amd64.deb) |
| `postgresql-14-http` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 51.0 KiB | [postgresql-14-http_1.7.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-http` | `1.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.7 KiB | [postgresql-14-http_1.7.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgsql-http/postgresql-14-http_1.7.1-1PIGSTY~resolute_arm64.deb) |
| `postgresql-14-http` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 50.8 KiB | [postgresql-14-http_1.7.0-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-http/postgresql-14-http_1.7.0-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pramsey/pgsql-http" title="Repository" icon="github" subtitle="github.com/pramsey/pgsql-http" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsql-http-1.7.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_http;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_http;		# install via package name, for the active PG version
pig install http;		# install by extension name, for the current active PG version

pig install http -v 18;   # install for PG 18
pig install http -v 17;   # install for PG 17
pig install http -v 16;   # install for PG 16
pig install http -v 15;   # install for PG 15
pig install http -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION http;
```


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

