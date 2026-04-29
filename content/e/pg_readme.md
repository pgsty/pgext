---
title: "pg_readme"
linkTitle: "pg_readme"
description: "Generate a README.md document for a database extension or schema"
weight: 4300
categories: ["UTIL"]
width: full
---

[**pg_readme**](https://github.com/bigsmoke/pg_readme) : Generate a README.md document for a database extension or schema


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4300** | {{< badge content="pg_readme" link="https://github.com/bigsmoke/pg_readme" >}} | {{< ext "pg_readme" >}} | `0.7.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "hstore" >}} |
|   **See Also**    | {{< ext "ddl_historization" >}} {{< ext "schedoc" >}} {{< ext "pg_render" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |
|    **Siblings**   | {{< ext "pg_readme_test_extension" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.7.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_readme` | `hstore` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.7.0` | {{< bg "18" "pg_readme_18" "green" >}} {{< bg "17" "pg_readme_17" "green" >}} {{< bg "16" "pg_readme_16" "green" >}} {{< bg "15" "pg_readme_15" "green" >}} {{< bg "14" "pg_readme_14" "green" >}} | `pg_readme_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.7.0` | {{< bg "18" "postgresql-18-pg-readme" "green" >}} {{< bg "17" "postgresql-17-pg-readme" "green" >}} {{< bg "16" "postgresql-16-pg-readme" "green" >}} {{< bg "15" "postgresql-15-pg-readme" "green" >}} {{< bg "14" "postgresql-14-pg-readme" "green" >}} | `postgresql-$v-pg-readme` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.0" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-readme : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-readme : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-readme : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-readme : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-readme : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-readme : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-readme : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-readme : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-readme : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-readme : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readme_18` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.5 KiB | [pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.5 KiB | [pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.7 KiB | [pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [pg_readme_18-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_readme_18-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [pg_readme_18-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_readme_18-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pg-readme` | `0.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.3 KiB | [postgresql-18-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-readme` | `0.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.3 KiB | [postgresql-18-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-readme` | `0.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.3 KiB | [postgresql-18-pg-readme_0.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-readme` | `0.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.3 KiB | [postgresql-18-pg-readme_0.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-readme` | `0.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.7 KiB | [postgresql-18-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-readme` | `0.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.7 KiB | [postgresql-18-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-readme` | `0.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.8 KiB | [postgresql-18-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-readme` | `0.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.8 KiB | [postgresql-18-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readme_17` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.5 KiB | [pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.5 KiB | [pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.8 KiB | [pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [pg_readme_17-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_readme_17-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [pg_readme_17-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_readme_17-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pg-readme` | `0.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.3 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-readme` | `0.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.3 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-readme` | `0.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.3 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-readme` | `0.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.3 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-readme` | `0.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.7 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-readme` | `0.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.7 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-readme` | `0.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.8 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-readme` | `0.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.8 KiB | [postgresql-17-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readme_16` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.5 KiB | [pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.5 KiB | [pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.8 KiB | [pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [pg_readme_16-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_readme_16-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [pg_readme_16-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_readme_16-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pg-readme` | `0.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.3 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-readme` | `0.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.3 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-readme` | `0.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.3 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-readme` | `0.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.3 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-readme` | `0.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.7 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-readme` | `0.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.7 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-readme` | `0.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.8 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-readme` | `0.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.8 KiB | [postgresql-16-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readme_15` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.5 KiB | [pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.5 KiB | [pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.8 KiB | [pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [pg_readme_15-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_readme_15-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [pg_readme_15-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_readme_15-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pg-readme` | `0.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.3 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-readme` | `0.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.3 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-readme` | `0.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.3 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-readme` | `0.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.3 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-readme` | `0.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.8 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-readme` | `0.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.8 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-readme` | `0.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.8 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-readme` | `0.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.8 KiB | [postgresql-15-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readme_14` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.5 KiB | [pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.5 KiB | [pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.8 KiB | [pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [pg_readme_14-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_readme_14-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [pg_readme_14-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_readme_14-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-14-pg-readme` | `0.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 18.3 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-readme` | `0.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.3 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-readme` | `0.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 18.3 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-readme` | `0.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.3 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-readme` | `0.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.7 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-readme` | `0.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.7 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-readme` | `0.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.8 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-readme` | `0.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.8 KiB | [postgresql-14-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_readme" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_readme" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_readme-0.7.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_readme;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_readme;		# install via package name, for the active PG version

pig install pg_readme -v 18;   # install for PG 18
pig install pg_readme -v 17;   # install for PG 17
pig install pg_readme -v 16;   # install for PG 16
pig install pg_readme -v 15;   # install for PG 15
pig install pg_readme -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_readme CASCADE; -- requires hstore
```




## Usage

> [pg_readme: Auto-generate README documentation from PostgreSQL COMMENT objects](https://github.com/bigsmoke/pg_readme)

Generates `README.md` documents for extensions or schemas based on `COMMENT` objects in the `pg_description` system catalog.

### Generate Extension README

```sql
SELECT pg_extension_readme('my_extension'::name);
```

### Generate Schema README

```sql
SELECT pg_schema_readme('my_schema'::regnamespace);
```

### Processing Instructions

Include these XML processing instructions in your `COMMENT ON EXTENSION` or `COMMENT ON SCHEMA`:

- `<?pg-readme-reference?>` -- replaced with a full object reference
- `<?pg-readme-colophon?>` -- adds a colophon with pg_readme info

### Typical Workflow

Create a readme-generating function in your extension:

```sql
CREATE FUNCTION my_ext_readme() RETURNS text
    VOLATILE SET search_path FROM current
    SET pg_readme.include_view_definitions TO 'true'
    LANGUAGE plpgsql AS $$
DECLARE _readme text;
BEGIN
    CREATE EXTENSION IF NOT EXISTS pg_readme WITH VERSION '0.1.0';
    _readme := pg_extension_readme('my_extension'::name);
    RAISE transaction_rollback;
EXCEPTION WHEN transaction_rollback THEN RETURN _readme;
END; $$;
```

Then generate with: `make README.md`

### Settings

| Setting | Default |
|---------|---------|
| `pg_readme.include_view_definitions` | `true` |
| `pg_readme.include_routine_definitions_like` | `'{test__%}'` |
| `pg_readme.include_this_routine_definition` | `null` |
