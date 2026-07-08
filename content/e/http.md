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
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_14 : AVAIL 7" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_14 : AVAIL 7" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_14 : AVAIL 4" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7.1" "pgsql_http_14 : AVAIL 3" "green" >}} |
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
| `pgsql_http_18` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.1 KiB | [pgsql_http_18-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsql_http_18-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_18` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.6 KiB | [pgsql_http_18-1.7.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_http_18-1.7.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_http_18` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pgsql_http_18-1.7.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsql_http_18-1.7.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pgsql_http_18` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.6 KiB | [pgsql_http_18-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsql_http_18-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_18` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [pgsql_http_18-1.7.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_http_18-1.7.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_http_18` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.7 KiB | [pgsql_http_18-1.7.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsql_http_18-1.7.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pgsql_http_18` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.8 KiB | [pgsql_http_18-1.7.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgsql_http_18-1.7.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsql_http_18` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.8 KiB | [pgsql_http_18-1.7.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_http_18-1.7.1-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_http_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [pgsql_http_18-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsql_http_18-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `pgsql_http_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.3 KiB | [pgsql_http_18-1.7.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgsql_http_18-1.7.0-1PGDG.rhel10.aarch64.rpm) |
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
| `pgsql_http_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.2 KiB | [pgsql_http_17-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_http_17-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_17` | `1.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.2 KiB | [pgsql_http_17-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_http_17-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_17` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.3 KiB | [pgsql_http_17-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_http_17-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_17` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.8 KiB | [pgsql_http_17-1.6.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsql_http_17-1.6.0-2PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_17` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.7 KiB | [pgsql_http_17-1.7.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_http_17-1.7.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_http_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pgsql_http_17-1.7.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_http_17-1.7.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pgsql_http_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.7 KiB | [pgsql_http_17-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_http_17-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_17` | `1.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.9 KiB | [pgsql_http_17-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_http_17-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_17` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.0 KiB | [pgsql_http_17-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_http_17-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_17` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.4 KiB | [pgsql_http_17-1.6.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsql_http_17-1.6.0-2PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_17` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [pgsql_http_17-1.7.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_http_17-1.7.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_http_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.8 KiB | [pgsql_http_17-1.7.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsql_http_17-1.7.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pgsql_http_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.0 KiB | [pgsql_http_17-1.7.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsql_http_17-1.7.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsql_http_17` | `1.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.1 KiB | [pgsql_http_17-1.6.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgsql_http_17-1.6.3-2PGDG.rhel10.x86_64.rpm) |
| `pgsql_http_17` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.9 KiB | [pgsql_http_17-1.7.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_http_17-1.7.1-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_http_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgsql_http_17-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_http_17-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `pgsql_http_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.5 KiB | [pgsql_http_17-1.7.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_http_17-1.7.0-1PGDG.rhel10.aarch64.rpm) |
| `pgsql_http_17` | `1.6.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.6 KiB | [pgsql_http_17-1.6.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgsql_http_17-1.6.3-2PGDG.rhel10.aarch64.rpm) |
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
| `pgsql_http_16` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.2 KiB | [pgsql_http_16-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_http_16-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_16` | `1.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.2 KiB | [pgsql_http_16-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_http_16-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_16` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.3 KiB | [pgsql_http_16-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_http_16-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_16` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.7 KiB | [pgsql_http_16-1.6.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_http_16-1.6.0-2PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_16` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.5 KiB | [pgsql_http_16-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsql_http_16-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_16` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.5 KiB | [pgsql_http_16-1.7.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_http_16-1.7.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_http_16` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [pgsql_http_16-1.7.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.7.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pgsql_http_16` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.6 KiB | [pgsql_http_16-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_16` | `1.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.9 KiB | [pgsql_http_16-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_16` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.0 KiB | [pgsql_http_16-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_16` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.4 KiB | [pgsql_http_16-1.6.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.6.0-2PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_16` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.2 KiB | [pgsql_http_16-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsql_http_16-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_16` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.0 KiB | [pgsql_http_16-1.7.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_http_16-1.7.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_http_16` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.9 KiB | [pgsql_http_16-1.7.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsql_http_16-1.7.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pgsql_http_16` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.9 KiB | [pgsql_http_16-1.7.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsql_http_16-1.7.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsql_http_16` | `1.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.0 KiB | [pgsql_http_16-1.6.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgsql_http_16-1.6.3-2PGDG.rhel10.x86_64.rpm) |
| `pgsql_http_16` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.8 KiB | [pgsql_http_16-1.7.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_http_16-1.7.1-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_http_16` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgsql_http_16-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_http_16-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `pgsql_http_16` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgsql_http_16-1.7.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_http_16-1.7.0-1PGDG.rhel10.aarch64.rpm) |
| `pgsql_http_16` | `1.6.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.6 KiB | [pgsql_http_16-1.6.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgsql_http_16-1.6.3-2PGDG.rhel10.aarch64.rpm) |
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
| `pgsql_http_15` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.5 KiB | [pgsql_http_15-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_http_15-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_15` | `1.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.4 KiB | [pgsql_http_15-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_http_15-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_15` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.7 KiB | [pgsql_http_15-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_http_15-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_15` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.1 KiB | [pgsql_http_15-1.6.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_http_15-1.6.0-2PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_15` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.9 KiB | [pgsql_http_15-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsql_http_15-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_15` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pgsql_http_15-1.7.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_http_15-1.7.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_http_15` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pgsql_http_15-1.7.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.7.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pgsql_http_15` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [pgsql_http_15-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_15` | `1.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.1 KiB | [pgsql_http_15-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_15` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.7 KiB | [pgsql_http_15-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_15` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.1 KiB | [pgsql_http_15-1.6.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.6.0-2PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_15` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.8 KiB | [pgsql_http_15-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsql_http_15-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_15` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.3 KiB | [pgsql_http_15-1.7.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_http_15-1.7.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_http_15` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.2 KiB | [pgsql_http_15-1.7.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsql_http_15-1.7.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pgsql_http_15` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.1 KiB | [pgsql_http_15-1.7.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsql_http_15-1.7.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsql_http_15` | `1.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.3 KiB | [pgsql_http_15-1.6.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgsql_http_15-1.6.3-2PGDG.rhel10.x86_64.rpm) |
| `pgsql_http_15` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.1 KiB | [pgsql_http_15-1.7.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_http_15-1.7.1-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_http_15` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.7 KiB | [pgsql_http_15-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_http_15-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `pgsql_http_15` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.7 KiB | [pgsql_http_15-1.7.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_http_15-1.7.0-1PGDG.rhel10.aarch64.rpm) |
| `pgsql_http_15` | `1.6.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.8 KiB | [pgsql_http_15-1.6.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgsql_http_15-1.6.3-2PGDG.rhel10.aarch64.rpm) |
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
| `pgsql_http_14` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.5 KiB | [pgsql_http_14-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_http_14-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_14` | `1.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.4 KiB | [pgsql_http_14-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_http_14-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_14` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.7 KiB | [pgsql_http_14-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_http_14-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_14` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.1 KiB | [pgsql_http_14-1.6.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_http_14-1.6.0-2PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_14` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.9 KiB | [pgsql_http_14-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsql_http_14-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsql_http_14` | `1.7.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pgsql_http_14-1.7.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsql_http_14-1.7.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsql_http_14` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.1 KiB | [pgsql_http_14-1.7.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.7.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pgsql_http_14` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [pgsql_http_14-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_14` | `1.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.1 KiB | [pgsql_http_14-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_14` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.7 KiB | [pgsql_http_14-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_14` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.1 KiB | [pgsql_http_14-1.6.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.6.0-2PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_14` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.8 KiB | [pgsql_http_14-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsql_http_14-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsql_http_14` | `1.7.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.3 KiB | [pgsql_http_14-1.7.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsql_http_14-1.7.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsql_http_14` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.0 KiB | [pgsql_http_14-1.7.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsql_http_14-1.7.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pgsql_http_14` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.2 KiB | [pgsql_http_14-1.7.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsql_http_14-1.7.0-1PGDG.rhel10.x86_64.rpm) |
| `pgsql_http_14` | `1.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.3 KiB | [pgsql_http_14-1.6.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgsql_http_14-1.6.3-2PGDG.rhel10.x86_64.rpm) |
| `pgsql_http_14` | `1.7.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.1 KiB | [pgsql_http_14-1.7.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsql_http_14-1.7.1-1PIGSTY.el10.aarch64.rpm) |
| `pgsql_http_14` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.7 KiB | [pgsql_http_14-1.7.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsql_http_14-1.7.0-1PGDG.rhel10.aarch64.rpm) |
| `pgsql_http_14` | `1.6.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [pgsql_http_14-1.6.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgsql_http_14-1.6.3-2PGDG.rhel10.aarch64.rpm) |
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

Sources: [README](https://github.com/pramsey/pgsql-http/blob/v1.7.1/README.md), [v1.7.1 release](https://github.com/pramsey/pgsql-http/releases/tag/v1.7.1)

`http` lets SQL code make HTTP requests through libcurl. Use it for controlled integration points such as triggers that notify an external service, SQL jobs that fetch a small remote payload, or database-side webhook calls.

```sql
CREATE EXTENSION http;
```

### Request And Response Types

Every request uses `http_request` and returns `http_response`:

```text
http_request(method http_method, uri varchar, headers http_header[], content_type varchar, content varchar)
http_response(status integer, content_type varchar, headers http_header[], content varchar)
```

Convenience wrappers call the same underlying `http(http_request)` function:

- `http_get(uri varchar)`
- `http_get(uri varchar, data jsonb)`
- `http_post(uri varchar, content varchar, content_type varchar)`
- `http_post(uri varchar, data jsonb)`
- `http_put(uri varchar, content varchar, content_type varchar)`
- `http_patch(uri varchar, content varchar, content_type varchar)`
- `http_delete(uri varchar)`
- `http_head(uri varchar)`

### Examples

```sql
SELECT status, content_type, content
FROM http_get('https://httpbun.com/ip');

SELECT content::json->'headers'->>'Authorization'
FROM http((
  'GET',
  'https://httpbun.com/headers',
  http_headers('Authorization', 'Bearer token'),
  NULL,
  NULL
)::http_request);

SELECT status, content::json->'form' AS form
FROM http_post(
  'https://httpbun.com/post',
  jsonb_build_object('myvar', 'myval', 'foo', 'bar')
);

SELECT status, content_type, content::json->>'data' AS data
FROM http_put('https://httpbun.com/put', 'some text', 'text/plain');
```

Inspect response headers by unnesting the `headers` array:

```sql
SELECT (unnest(headers)).*
FROM http_get('https://httpbun.com/');
```

### Binary Content

The README warns that `varchar::bytea` is not safe for binary response bodies because it stops at zero-valued bytes. Use `text_to_bytea(content)` for response content and `bytea_to_text(bytea)` when sending binary request bodies.

```sql
WITH http AS (
  SELECT * FROM http_get('https://httpbingo.org/image/png')
)
SELECT content_type, length(text_to_bytea(content)) AS bytes
FROM http;
```

### Timeout And Version Notes

`pg_http` 1.7.1 is a compatibility and documentation release: it adds timeout examples, adds PostgreSQL 17 wait-event hooks, and includes PostgreSQL 19 support fixes. The user-facing SQL API remains the README surface above.
