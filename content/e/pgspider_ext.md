---
title: "pgspider_ext"
linkTitle: "pgspider_ext"
description: "foreign-data wrapper for remote PGSpider servers"
weight: 8540
categories: ["FDW"]
width: full
---

[**pgspider_ext**](https://github.com/pgspider/pgspider_ext) : foreign-data wrapper for remote PGSpider servers


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8540** | {{< badge content="pgspider_ext" link="https://github.com/pgspider/pgspider_ext" >}} | {{< ext "pgspider_ext" >}} | `1.3.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plproxy" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "postgres_fdw" >}} {{< ext "citus" >}} {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "mongo_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pgspider_ext` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.0` | {{< bg "18" "pgspider_ext_18" "red" >}} {{< bg "17" "pgspider_ext_17" "green" >}} {{< bg "16" "pgspider_ext_16" "green" >}} {{< bg "15" "pgspider_ext_15" "green" >}} {{< bg "14" "pgspider_ext_14" "red" >}} | `pgspider_ext_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.0` | {{< bg "18" "postgresql-18-pgspider-ext" "red" >}} {{< bg "17" "postgresql-17-pgspider-ext" "green" >}} {{< bg "16" "postgresql-16-pgspider-ext" "green" >}} {{< bg "15" "postgresql-15-pgspider-ext" "green" >}} {{< bg "14" "postgresql-14-pgspider-ext" "red" >}} | `postgresql-$v-pgspider-ext` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pgspider_ext_18 : MISS 0" "red" >}}      | {{< bg "PGDG 1.3.0" "pgspider_ext_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "pgspider_ext_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "pgspider_ext_15 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pgspider_ext_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pgspider_ext_18 : MISS 0" "red" >}}      | {{< bg "PGDG 1.3.0" "pgspider_ext_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "pgspider_ext_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "pgspider_ext_15 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pgspider_ext_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pgspider_ext_18 : MISS 0" "red" >}}      | {{< bg "PGDG 1.3.0" "pgspider_ext_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "pgspider_ext_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "pgspider_ext_15 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pgspider_ext_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pgspider_ext_18 : MISS 0" "red" >}}      | {{< bg "PGDG 1.3.0" "pgspider_ext_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "pgspider_ext_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "pgspider_ext_15 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pgspider_ext_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pgspider_ext_18 : MISS 0" "red" >}}      | {{< bg "PGDG 1.3.0" "pgspider_ext_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "pgspider_ext_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "pgspider_ext_15 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pgspider_ext_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pgspider_ext_18 : MISS 0" "red" >}}      | {{< bg "PGDG 1.3.0" "pgspider_ext_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "pgspider_ext_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "pgspider_ext_15 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pgspider_ext_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgspider-ext : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.0" "postgresql-17-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pgspider-ext : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pgspider-ext : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgspider-ext : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.0" "postgresql-17-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pgspider-ext : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pgspider-ext : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgspider-ext : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.0" "postgresql-17-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pgspider-ext : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pgspider-ext : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgspider-ext : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.0" "postgresql-17-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pgspider-ext : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pgspider-ext : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgspider-ext : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.0" "postgresql-17-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pgspider-ext : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pgspider-ext : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgspider-ext : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.0" "postgresql-17-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pgspider-ext : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pgspider-ext : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgspider-ext : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.0" "postgresql-17-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pgspider-ext : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pgspider-ext : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgspider-ext : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.3.0" "postgresql-17-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-16-pgspider-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.0" "postgresql-15-pgspider-ext : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pgspider-ext : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgspider-ext : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgspider-ext : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgspider-ext : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgspider-ext : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgspider-ext : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgspider-ext : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgspider-ext : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgspider-ext : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgspider-ext : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgspider-ext : MISS 0" "red" >}}      |


{{< tabs items="PG17,PG16,PG15" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgspider_ext_17` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.6 KiB | [pgspider_ext_17-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgspider_ext_17-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgspider_ext_17` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.7 KiB | [pgspider_ext_17-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgspider_ext_17-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgspider_ext_17` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.1 KiB | [pgspider_ext_17-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgspider_ext_17-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgspider_ext_17` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.4 KiB | [pgspider_ext_17-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgspider_ext_17-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgspider_ext_17` | `1.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.1 KiB | [pgspider_ext_17-1.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgspider_ext_17-1.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pgspider_ext_17` | `1.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.1 KiB | [pgspider_ext_17-1.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgspider_ext_17-1.3.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgspider-ext` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 48.8 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgspider-ext` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 47.1 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgspider-ext` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 48.8 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgspider-ext` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 47.3 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgspider-ext` | `1.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 61.7 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgspider-ext` | `1.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 60.7 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgspider-ext` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 50.9 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgspider-ext` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.6 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgspider_ext_16` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.6 KiB | [pgspider_ext_16-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgspider_ext_16-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgspider_ext_16` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.8 KiB | [pgspider_ext_16-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgspider_ext_16-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgspider_ext_16` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.2 KiB | [pgspider_ext_16-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgspider_ext_16-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgspider_ext_16` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.5 KiB | [pgspider_ext_16-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgspider_ext_16-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgspider_ext_16` | `1.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.2 KiB | [pgspider_ext_16-1.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgspider_ext_16-1.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pgspider_ext_16` | `1.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.2 KiB | [pgspider_ext_16-1.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgspider_ext_16-1.3.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgspider-ext` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 48.7 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgspider-ext` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 47.0 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgspider-ext` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 48.8 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgspider-ext` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 47.2 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgspider-ext` | `1.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 61.3 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgspider-ext` | `1.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 60.3 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgspider-ext` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 50.8 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgspider-ext` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.6 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgspider_ext_15` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.0 KiB | [pgspider_ext_15-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgspider_ext_15-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgspider_ext_15` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 28.0 KiB | [pgspider_ext_15-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgspider_ext_15-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgspider_ext_15` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.6 KiB | [pgspider_ext_15-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgspider_ext_15-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgspider_ext_15` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 28.8 KiB | [pgspider_ext_15-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgspider_ext_15-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgspider_ext_15` | `1.3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.4 KiB | [pgspider_ext_15-1.3.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgspider_ext_15-1.3.0-1PGDG.rhel10.x86_64.rpm) |
| `pgspider_ext_15` | `1.3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 29.6 KiB | [pgspider_ext_15-1.3.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgspider_ext_15-1.3.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgspider-ext` | `1.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 48.9 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgspider-ext` | `1.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 47.0 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgspider-ext` | `1.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 49.0 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgspider-ext` | `1.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 47.3 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgspider-ext` | `1.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 61.3 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgspider-ext` | `1.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 60.3 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgspider-ext` | `1.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 51.0 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgspider-ext` | `1.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.7 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgspider/pgspider_ext" title="Repository" icon="github" subtitle="github.com/pgspider/pgspider_ext" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgspider_ext-1.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgspider_ext;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgspider_ext;		# install via package name, for the active PG version

pig install pgspider_ext -v 17;   # install for PG 17
pig install pgspider_ext -v 16;   # install for PG 16
pig install pgspider_ext -v 15;   # install for PG 15

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgspider_ext;
```




## Usage

> [pgspider_ext: Foreign data wrapper for remote PGSpider servers](https://github.com/pgspider/pgspider_ext)

PGSpider Extension turns PostgreSQL into a distributed query engine by creating virtual partitioned tables that transparently query data across multiple remote nodes in parallel.

### Setup Child Servers

First, create servers for each data source using their respective FDWs:

```sql
CREATE EXTENSION pgspider_ext;
CREATE EXTENSION postgres_fdw;

CREATE SERVER pgsrv1 FOREIGN DATA WRAPPER postgres_fdw
  OPTIONS (host '127.0.0.1', port '5433', dbname 'postgres');
CREATE SERVER pgsrv2 FOREIGN DATA WRAPPER postgres_fdw
  OPTIONS (host '127.0.0.1', port '5434', dbname 'postgres');

CREATE USER MAPPING FOR CURRENT_USER SERVER pgsrv1
  OPTIONS (user 'user', password 'pass');
CREATE USER MAPPING FOR CURRENT_USER SERVER pgsrv2
  OPTIONS (user 'user', password 'pass');
```

### Create Child Foreign Tables

```sql
CREATE FOREIGN TABLE t1_pg1_child (i int, t text)
  SERVER pgsrv1 OPTIONS (table_name 't1');
CREATE FOREIGN TABLE t1_pg2_child (i int, t text)
  SERVER pgsrv2 OPTIONS (table_name 't1');
```

### Create PGSpider Partitioned Table

Create a PGSpider server and a partitioned parent table with an extra partition key column:

```sql
CREATE SERVER spdsrv FOREIGN DATA WRAPPER pgspider_ext;
CREATE USER MAPPING FOR CURRENT_USER SERVER spdsrv;

CREATE TABLE t1 (i int, t text, node text)
  PARTITION BY LIST (node);

CREATE FOREIGN TABLE t1_pg1 PARTITION OF t1
  FOR VALUES IN ('node1') SERVER spdsrv;
CREATE FOREIGN TABLE t1_pg2 PARTITION OF t1
  FOR VALUES IN ('node2') SERVER spdsrv
  OPTIONS (child_name 't1_pg2_child');
```

By default, PGSpider finds child tables using the pattern `[parent_table_name]_child`. Use `child_name` to specify a different name.

### Query Across All Nodes

```sql
SELECT * FROM t1;
 i  | t | node
----+---+-------
 10 | a | node1
 11 | b | node1
 20 | c | node2
 21 | d | node2
```

Queries automatically fan out to all child nodes in parallel. WHERE clauses and aggregate functions are pushed down to child nodes:

```sql
SET enable_partitionwise_aggregate TO on;
SELECT count(*), node FROM t1 GROUP BY node;
```

**Note:** Only SELECT operations are supported; INSERT, UPDATE, and DELETE are not supported.
