---
title: "pg_readme"
linkTitle: "pg_readme"
description: "Generate a Markdown README from PostgreSQL COMMENT objects"
weight: 4300
categories: ["UTIL"]
width: full
---

[**pg_readme**](https://github.com/bigsmoke/pg_readme) : Generate a Markdown README from PostgreSQL COMMENT objects


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4300** | {{< badge content="pg_readme" link="https://github.com/bigsmoke/pg_readme" >}} | {{< ext "pg_readme" >}} | `0.7.1` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "hstore" >}} |
|   **See Also**    | {{< ext "ddlx" >}} {{< ext "pg_render" >}} {{< ext "schedoc" >}} {{< ext "pgdd" >}} {{< ext "meta" >}} {{< ext "pgpdf" >}} {{< ext "pg_get_functiondef" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_catcheck" >}} {{< ext "pg_query_rewrite" >}} |
|    **Siblings**   | {{< ext "pg_readme_test_extension" >}} |

> [!Note] Catalog release is 0.7.1; PGDG remains the RPM maintainer at 0.7.0, so the PIGSTY 0.7.1 RPM must not be published; PIGSTY maintains the 0.7.1 DEB package.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.7.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_readme` | `hstore` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.7.0` | {{< bg "18" "pg_readme_18" "green" >}} {{< bg "17" "pg_readme_17" "green" >}} {{< bg "16" "pg_readme_16" "green" >}} {{< bg "15" "pg_readme_15" "green" >}} {{< bg "14" "pg_readme_14" "green" >}} | `pg_readme_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.7.1` | {{< bg "18" "postgresql-18-pg-readme" "green" >}} {{< bg "17" "postgresql-17-pg-readme" "green" >}} {{< bg "16" "postgresql-16-pg-readme" "green" >}} {{< bg "15" "postgresql-15-pg-readme" "green" >}} {{< bg "14" "postgresql-14-pg-readme" "green" >}} | `postgresql-$v-pg-readme` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.7.0" "pg_readme_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.7.0" "pg_readme_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-18-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-17-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-16-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-15-pg-readme : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.1" "postgresql-14-pg-readme : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readme_18` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.5 KiB | [pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.5 KiB | [pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_readme_18-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_18-0.7.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_readme_18-0.7.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.7 KiB | [pg_readme_18-0.7.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_readme_18-0.7.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.7 KiB | [pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_readme_18-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.0 KiB | [pg_readme_18-0.7.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_readme_18-0.7.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [pg_readme_18-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_readme_18-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.0 KiB | [pg_readme_18-0.7.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_readme_18-0.7.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_readme_18` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [pg_readme_18-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_readme_18-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-18-pg-readme` | `0.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.5 KiB | [postgresql-18-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-readme` | `0.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.5 KiB | [postgresql-18-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-readme` | `0.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.5 KiB | [postgresql-18-pg-readme_0.7.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.1-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-readme` | `0.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.5 KiB | [postgresql-18-pg-readme_0.7.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.1-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-readme` | `0.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.0 KiB | [postgresql-18-pg-readme_0.7.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.1-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-readme` | `0.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.0 KiB | [postgresql-18-pg-readme_0.7.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.1-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-readme` | `0.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.0 KiB | [postgresql-18-pg-readme_0.7.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.1-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-readme` | `0.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.0 KiB | [postgresql-18-pg-readme_0.7.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.1-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-readme` | `0.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.1 KiB | [postgresql-18-pg-readme_0.7.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.1-1PIGSTY~resolute_all.deb) |
| `postgresql-18-pg-readme` | `0.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.1 KiB | [postgresql-18-pg-readme_0.7.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-readme/postgresql-18-pg-readme_0.7.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readme_17` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.5 KiB | [pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.5 KiB | [pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_readme_17-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_17-0.7.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_readme_17-0.7.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.7 KiB | [pg_readme_17-0.7.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_readme_17-0.7.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.8 KiB | [pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_readme_17-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.0 KiB | [pg_readme_17-0.7.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_readme_17-0.7.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [pg_readme_17-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_readme_17-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.0 KiB | [pg_readme_17-0.7.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_readme_17-0.7.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_readme_17` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [pg_readme_17-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_readme_17-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pg-readme` | `0.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.5 KiB | [postgresql-17-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-readme` | `0.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.5 KiB | [postgresql-17-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-readme` | `0.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.5 KiB | [postgresql-17-pg-readme_0.7.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.1-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-readme` | `0.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.5 KiB | [postgresql-17-pg-readme_0.7.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.1-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-readme` | `0.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.0 KiB | [postgresql-17-pg-readme_0.7.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.1-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-readme` | `0.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.0 KiB | [postgresql-17-pg-readme_0.7.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.1-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-readme` | `0.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.0 KiB | [postgresql-17-pg-readme_0.7.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.1-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-readme` | `0.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.0 KiB | [postgresql-17-pg-readme_0.7.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.1-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-readme` | `0.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.1 KiB | [postgresql-17-pg-readme_0.7.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.1-1PIGSTY~resolute_all.deb) |
| `postgresql-17-pg-readme` | `0.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.1 KiB | [postgresql-17-pg-readme_0.7.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-readme/postgresql-17-pg-readme_0.7.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readme_16` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.5 KiB | [pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.5 KiB | [pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_readme_16-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_16-0.7.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_readme_16-0.7.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.8 KiB | [pg_readme_16-0.7.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_readme_16-0.7.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.8 KiB | [pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_readme_16-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.0 KiB | [pg_readme_16-0.7.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_readme_16-0.7.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [pg_readme_16-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_readme_16-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.0 KiB | [pg_readme_16-0.7.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_readme_16-0.7.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_readme_16` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [pg_readme_16-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_readme_16-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pg-readme` | `0.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.5 KiB | [postgresql-16-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-readme` | `0.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.5 KiB | [postgresql-16-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-readme` | `0.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.5 KiB | [postgresql-16-pg-readme_0.7.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.1-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-readme` | `0.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.5 KiB | [postgresql-16-pg-readme_0.7.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.1-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-readme` | `0.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.0 KiB | [postgresql-16-pg-readme_0.7.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.1-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-readme` | `0.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.0 KiB | [postgresql-16-pg-readme_0.7.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.1-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-readme` | `0.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.0 KiB | [postgresql-16-pg-readme_0.7.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.1-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-readme` | `0.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.0 KiB | [postgresql-16-pg-readme_0.7.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.1-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-readme` | `0.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.1 KiB | [postgresql-16-pg-readme_0.7.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.1-1PIGSTY~resolute_all.deb) |
| `postgresql-16-pg-readme` | `0.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.1 KiB | [postgresql-16-pg-readme_0.7.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-readme/postgresql-16-pg-readme_0.7.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readme_15` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.5 KiB | [pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.5 KiB | [pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_readme_15-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_15-0.7.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_readme_15-0.7.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.7 KiB | [pg_readme_15-0.7.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_readme_15-0.7.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.8 KiB | [pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_readme_15-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.0 KiB | [pg_readme_15-0.7.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_readme_15-0.7.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [pg_readme_15-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_readme_15-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.0 KiB | [pg_readme_15-0.7.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_readme_15-0.7.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_readme_15` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [pg_readme_15-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_readme_15-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pg-readme` | `0.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.5 KiB | [postgresql-15-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-readme` | `0.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.5 KiB | [postgresql-15-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-readme` | `0.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.5 KiB | [postgresql-15-pg-readme_0.7.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.1-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-readme` | `0.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.5 KiB | [postgresql-15-pg-readme_0.7.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.1-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-readme` | `0.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.0 KiB | [postgresql-15-pg-readme_0.7.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.1-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-readme` | `0.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.0 KiB | [postgresql-15-pg-readme_0.7.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.1-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-readme` | `0.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.0 KiB | [postgresql-15-pg-readme_0.7.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.1-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-readme` | `0.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.0 KiB | [postgresql-15-pg-readme_0.7.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.1-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-readme` | `0.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.1 KiB | [postgresql-15-pg-readme_0.7.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.1-1PIGSTY~resolute_all.deb) |
| `postgresql-15-pg-readme` | `0.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.1 KiB | [postgresql-15-pg-readme_0.7.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-readme/postgresql-15-pg-readme_0.7.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_readme_14` | `0.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.5 KiB | [pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.5 KiB | [pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_readme_14-0.7.0-1PGDG.rhel8.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_14-0.7.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_readme_14-0.7.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.8 KiB | [pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.7 KiB | [pg_readme_14-0.7.0-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_readme_14-0.7.0-1PGDG.rhel9.8.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.8 KiB | [pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_readme_14-0.7.0-1PGDG.rhel9.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.0 KiB | [pg_readme_14-0.7.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_readme_14-0.7.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [pg_readme_14-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_readme_14-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.0 KiB | [pg_readme_14-0.7.0-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_readme_14-0.7.0-1PGDG.rhel10.2.noarch.rpm) |
| `pg_readme_14` | `0.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [pg_readme_14-0.7.0-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_readme_14-0.7.0-1PGDG.rhel10.noarch.rpm) |
| `postgresql-14-pg-readme` | `0.7.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.5 KiB | [postgresql-14-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pg-readme` | `0.7.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.5 KiB | [postgresql-14-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pg-readme` | `0.7.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.5 KiB | [postgresql-14-pg-readme_0.7.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.1-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pg-readme` | `0.7.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.5 KiB | [postgresql-14-pg-readme_0.7.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.1-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pg-readme` | `0.7.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.0 KiB | [postgresql-14-pg-readme_0.7.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.1-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pg-readme` | `0.7.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.0 KiB | [postgresql-14-pg-readme_0.7.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.1-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pg-readme` | `0.7.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 20.0 KiB | [postgresql-14-pg-readme_0.7.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.1-1PIGSTY~noble_all.deb) |
| `postgresql-14-pg-readme` | `0.7.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.0 KiB | [postgresql-14-pg-readme_0.7.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.1-1PIGSTY~noble_all.deb) |
| `postgresql-14-pg-readme` | `0.7.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 20.1 KiB | [postgresql-14-pg-readme_0.7.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.1-1PIGSTY~resolute_all.deb) |
| `postgresql-14-pg-readme` | `0.7.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 20.1 KiB | [postgresql-14-pg-readme_0.7.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-readme/postgresql-14-pg-readme_0.7.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_readme" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_readme" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_readme-0.7.1.tar.gz" >}}
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

Sources:

- [pg_readme 0.7.1 README](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/README.md)
- [pg_readme 0.7.1 control file](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/pg_readme.control)
- [pg_readme 0.7.1 upgrade SQL](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/sql/pg_readme--0.7.0--0.7.1.sql)
- [Pigsty package matrix](https://pgext.cloud/ext/pg_readme)

`pg_readme` generates Markdown documentation for a PostgreSQL extension or schema from `COMMENT` objects and live catalog metadata. Use it to keep an extension's README close to its SQL definitions and verify the generated output in source control.

### Install and Generate Markdown

```sql
CREATE EXTENSION pg_readme CASCADE;

SELECT pg_extension_readme('my_extension'::name);
SELECT pg_schema_readme('my_schema'::regnamespace);
```

The control file requires `hstore`, is relocatable, and permits non-superuser installation when the caller can install its dependencies and create the objects.

### Add Processing Instructions

Put Markdown and processing instructions in the extension or schema comment:

```sql
COMMENT ON EXTENSION my_extension IS $markdown$
### `my_extension`

What the extension does.

### Reference

<?pg-readme-reference?>

### Colophon

<?pg-readme-colophon?>
$markdown$;
```

`<?pg-readme-reference?>` expands to a catalog-derived object reference. `<?pg-readme-colophon?>` adds generation metadata. Optional instruction attributes can adjust the heading depth when embedding generated sections.

### Settings

- `pg_readme.include_view_definitions`: include view definitions; default `true`.
- `pg_readme.include_routine_definitions_like`: array of routine-name patterns whose definitions are included; default `'{test__%}'`.
- `pg_readme.include_this_routine_definition`: routine-local override for including the current definition.
- `pg_readme.readme_url`: upstream README link used by generated material.

Use `SET` options on a wrapper function or transaction when a project needs reproducible generation settings.

### Version 0.7.1 and Caveats

- Version 0.7.1 fixes PostgreSQL 18 reference generation that could duplicate array/composite table types and `NOT NULL` markers.
- Upstream and the current Pigsty DEB package are 0.7.1, while the current Pigsty RPM package remains 0.7.0. Check `pg_available_extension_versions` before relying on the PostgreSQL 18 fix.
- Generated output reflects the current database catalog, installed extension versions, comments, and generation time. Review diffs instead of assuming two environments produce identical text.
- Catalog introspection does not replace hand-written operational guidance. Keep prerequisites, preload/restart behavior, upgrade notes, and unsafe operations in curated prose.
- The singular setting `pg_readme.include_routine_definition_like` appears in an old README wrapper example, but the documented current GUC is the plural `pg_readme.include_routine_definitions_like`.
