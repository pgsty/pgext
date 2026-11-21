---
title: "ddlx"
linkTitle: "ddlx"
description: "DDL eXtractor functions"
weight: 5080
categories: ["ADMIN"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_ddlx**](https://github.com/lacanoid/pgddl) : DDL eXtractor functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5080** | {{< badge content="ddlx" link="https://github.com/lacanoid/pgddl" >}} | {{< ext "ddlx" "pg_ddlx" >}} | `0.32` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_readme" >}} {{< ext "pgdd" >}} {{< ext "meta" >}} {{< ext "pg_catcheck" >}} {{< ext "pg_get_functiondef" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_render" >}} {{< ext "schedoc" >}} {{< ext "pg_query_rewrite" >}} {{< ext "pgpdf" >}} |

> [!Note] Upstream 0.31 tag is marked prerelease, as were the previously packaged 0.30 and 0.29 tags.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.32` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_ddlx` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.32` | {{< bg "19" "ddlx_19" "green" >}} {{< bg "18" "ddlx_18" "green" >}} {{< bg "17" "ddlx_17" "green" >}} {{< bg "16" "ddlx_16" "green" >}} {{< bg "15" "ddlx_15" "green" >}} | `ddlx_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.32` | {{< bg "19" "postgresql-19-ddlx" "green" >}} {{< bg "18" "postgresql-18-ddlx" "green" >}} {{< bg "17" "postgresql-17-ddlx" "green" >}} {{< bg "16" "postgresql-16-ddlx" "green" >}} {{< bg "15" "postgresql-15-ddlx" "green" >}} | `postgresql-$v-ddlx` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "ddlx_19 : N/A 0" "gray" >}} | {{< bg "PGDG 0.32" "ddlx_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_15 : AVAIL 11" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "ddlx_19 : N/A 0" "gray" >}} | {{< bg "PGDG 0.32" "ddlx_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_15 : AVAIL 11" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.32" "ddlx_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_15 : AVAIL 4" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.32" "ddlx_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_15 : AVAIL 4" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.32" "ddlx_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_15 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.32" "ddlx_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 0.32" "ddlx_15 : AVAIL 8" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "MISS" "postgresql-19-ddlx : MISS 0" "red" >}} | {{< bg "PIGSTY 0.31" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-15-ddlx : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "MISS" "postgresql-19-ddlx : MISS 0" "red" >}} | {{< bg "PIGSTY 0.31" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-15-ddlx : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "MISS" "postgresql-19-ddlx : MISS 0" "red" >}} | {{< bg "PIGSTY 0.31" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-15-ddlx : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "MISS" "postgresql-19-ddlx : MISS 0" "red" >}} | {{< bg "PIGSTY 0.31" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-15-ddlx : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "MISS" "postgresql-19-ddlx : MISS 0" "red" >}} | {{< bg "PIGSTY 0.31" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-15-ddlx : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "MISS" "postgresql-19-ddlx : MISS 0" "red" >}} | {{< bg "PIGSTY 0.31" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-15-ddlx : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "MISS" "postgresql-19-ddlx : MISS 0" "red" >}} | {{< bg "PIGSTY 0.31" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-15-ddlx : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "MISS" "postgresql-19-ddlx : MISS 0" "red" >}} | {{< bg "PIGSTY 0.31" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-15-ddlx : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "MISS" "postgresql-19-ddlx : MISS 0" "red" >}} | {{< bg "PIGSTY 0.31" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-15-ddlx : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "MISS" "postgresql-19-ddlx : MISS 0" "red" >}} | {{< bg "PIGSTY 0.31" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.31" "postgresql-15-ddlx : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddlx_19` | `0.32` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [ddlx_19-0.32-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/ddlx_19-0.32-1PGDG.rhel9.9.noarch.rpm) |
| `ddlx_19` | `0.32` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.3 KiB | [ddlx_19-0.32-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/ddlx_19-0.32-1PGDG.rhel9.9.noarch.rpm) |
| `ddlx_19` | `0.32` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.5 KiB | [ddlx_19-0.32-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/ddlx_19-0.32-1PGDG.rhel10.3.noarch.rpm) |
| `ddlx_19` | `0.32` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.5 KiB | [ddlx_19-0.32-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/ddlx_19-0.32-1PGDG.rhel10.3.noarch.rpm) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddlx_18` | `0.32` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.0 KiB | [ddlx_18-0.32-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/ddlx_18-0.32-1PGDG.rhel8.10.noarch.rpm) |
| `ddlx_18` | `0.31` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.4 KiB | [ddlx_18-0.31-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_18-0.31-1PGSTY.el8.noarch.rpm) |
| `ddlx_18` | `0.30` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [ddlx_18-0.30-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/ddlx_18-0.30-5PGDG.rhel8.10.noarch.rpm) |
| `ddlx_18` | `0.30` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.9 KiB | [ddlx_18-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/ddlx_18-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_18` | `0.32` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.0 KiB | [ddlx_18-0.32-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/ddlx_18-0.32-1PGDG.rhel8.10.noarch.rpm) |
| `ddlx_18` | `0.31` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.4 KiB | [ddlx_18-0.31-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_18-0.31-1PGSTY.el8.noarch.rpm) |
| `ddlx_18` | `0.30` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.6 KiB | [ddlx_18-0.30-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/ddlx_18-0.30-5PGDG.rhel8.10.noarch.rpm) |
| `ddlx_18` | `0.30` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.8 KiB | [ddlx_18-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/ddlx_18-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_18` | `0.32` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [ddlx_18-0.32-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ddlx_18-0.32-1PGDG.rhel9.9.noarch.rpm) |
| `ddlx_18` | `0.31` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.2 KiB | [ddlx_18-0.31-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_18-0.31-1PGSTY.el9.noarch.rpm) |
| `ddlx_18` | `0.30` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.0 KiB | [ddlx_18-0.30-5PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ddlx_18-0.30-5PGDG.rhel9.9.noarch.rpm) |
| `ddlx_18` | `0.32` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.3 KiB | [ddlx_18-0.32-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ddlx_18-0.32-1PGDG.rhel9.9.noarch.rpm) |
| `ddlx_18` | `0.31` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.1 KiB | [ddlx_18-0.31-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_18-0.31-1PGSTY.el9.noarch.rpm) |
| `ddlx_18` | `0.30` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.0 KiB | [ddlx_18-0.30-5PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ddlx_18-0.30-5PGDG.rhel9.9.noarch.rpm) |
| `ddlx_18` | `0.32` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.5 KiB | [ddlx_18-0.32-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ddlx_18-0.32-1PGDG.rhel10.3.noarch.rpm) |
| `ddlx_18` | `0.31` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.3 KiB | [ddlx_18-0.31-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddlx_18-0.31-1PGSTY.el10.noarch.rpm) |
| `ddlx_18` | `0.30` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.2 KiB | [ddlx_18-0.30-5PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ddlx_18-0.30-5PGDG.rhel10.3.noarch.rpm) |
| `ddlx_18` | `0.32` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.5 KiB | [ddlx_18-0.32-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ddlx_18-0.32-1PGDG.rhel10.3.noarch.rpm) |
| `ddlx_18` | `0.32` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.5 KiB | [ddlx_18-0.32-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ddlx_18-0.32-1PGDG.rhel10.2.noarch.rpm) |
| `ddlx_18` | `0.31` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.3 KiB | [ddlx_18-0.31-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddlx_18-0.31-1PGSTY.el10.noarch.rpm) |
| `ddlx_18` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [ddlx_18-0.30-5PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ddlx_18-0.30-5PGDG.rhel10.3.noarch.rpm) |
| `ddlx_18` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [ddlx_18-0.30-5PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ddlx_18-0.30-5PGDG.rhel10.2.noarch.rpm) |
| `ddlx_18` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.7 KiB | [ddlx_18-0.30-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ddlx_18-0.30-2PGDG.rhel10.2.noarch.rpm) |
| `ddlx_18` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [ddlx_18-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ddlx_18-0.30-1PGDG.rhel10.noarch.rpm) |
| `postgresql-18-ddlx` | `0.31` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.9 KiB | [postgresql-18-ddlx_0.31-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-18-ddlx_0.31-1PGSTY~bookworm_all.deb) |
| `postgresql-18-ddlx` | `0.31` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.9 KiB | [postgresql-18-ddlx_0.31-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-18-ddlx_0.31-1PGSTY~bookworm_all.deb) |
| `postgresql-18-ddlx` | `0.31` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.8 KiB | [postgresql-18-ddlx_0.31-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-18-ddlx_0.31-1PGSTY~trixie_all.deb) |
| `postgresql-18-ddlx` | `0.31` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.8 KiB | [postgresql-18-ddlx_0.31-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-18-ddlx_0.31-1PGSTY~trixie_all.deb) |
| `postgresql-18-ddlx` | `0.31` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.1 KiB | [postgresql-18-ddlx_0.31-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-18-ddlx_0.31-1PGSTY~jammy_all.deb) |
| `postgresql-18-ddlx` | `0.31` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.1 KiB | [postgresql-18-ddlx_0.31-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-18-ddlx_0.31-1PGSTY~jammy_all.deb) |
| `postgresql-18-ddlx` | `0.31` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.9 KiB | [postgresql-18-ddlx_0.31-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-18-ddlx_0.31-1PGSTY~noble_all.deb) |
| `postgresql-18-ddlx` | `0.31` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.9 KiB | [postgresql-18-ddlx_0.31-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-18-ddlx_0.31-1PGSTY~noble_all.deb) |
| `postgresql-18-ddlx` | `0.31` | [u26.x86_64](/os/u26.x86_64) | pigsty | 25.9 KiB | [postgresql-18-ddlx_0.31-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddlx/postgresql-18-ddlx_0.31-1PGSTY~resolute_all.deb) |
| `postgresql-18-ddlx` | `0.31` | [u26.aarch64](/os/u26.aarch64) | pigsty | 25.9 KiB | [postgresql-18-ddlx_0.31-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddlx/postgresql-18-ddlx_0.31-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddlx_17` | `0.32` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.0 KiB | [ddlx_17-0.32-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ddlx_17-0.32-1PGDG.rhel8.10.noarch.rpm) |
| `ddlx_17` | `0.31` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.4 KiB | [ddlx_17-0.31-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_17-0.31-1PGSTY.el8.noarch.rpm) |
| `ddlx_17` | `0.30` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [ddlx_17-0.30-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ddlx_17-0.30-5PGDG.rhel8.10.noarch.rpm) |
| `ddlx_17` | `0.30` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.4 KiB | [ddlx_17-0.30-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_17-0.30-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_17` | `0.30` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.9 KiB | [ddlx_17-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ddlx_17-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_17` | `0.29` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.8 KiB | [ddlx_17-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ddlx_17-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_17` | `0.32` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.0 KiB | [ddlx_17-0.32-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ddlx_17-0.32-1PGDG.rhel8.10.noarch.rpm) |
| `ddlx_17` | `0.31` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.4 KiB | [ddlx_17-0.31-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_17-0.31-1PGSTY.el8.noarch.rpm) |
| `ddlx_17` | `0.30` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.6 KiB | [ddlx_17-0.30-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ddlx_17-0.30-5PGDG.rhel8.10.noarch.rpm) |
| `ddlx_17` | `0.30` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.3 KiB | [ddlx_17-0.30-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_17-0.30-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_17` | `0.30` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.8 KiB | [ddlx_17-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ddlx_17-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_17` | `0.29` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.8 KiB | [ddlx_17-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ddlx_17-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_17` | `0.32` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [ddlx_17-0.32-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ddlx_17-0.32-1PGDG.rhel9.9.noarch.rpm) |
| `ddlx_17` | `0.31` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.2 KiB | [ddlx_17-0.31-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_17-0.31-1PGSTY.el9.noarch.rpm) |
| `ddlx_17` | `0.30` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.0 KiB | [ddlx_17-0.30-5PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ddlx_17-0.30-5PGDG.rhel9.9.noarch.rpm) |
| `ddlx_17` | `0.30` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.2 KiB | [ddlx_17-0.30-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_17-0.30-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_17` | `0.32` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.3 KiB | [ddlx_17-0.32-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ddlx_17-0.32-1PGDG.rhel9.9.noarch.rpm) |
| `ddlx_17` | `0.31` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.1 KiB | [ddlx_17-0.31-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_17-0.31-1PGSTY.el9.noarch.rpm) |
| `ddlx_17` | `0.30` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.0 KiB | [ddlx_17-0.30-5PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ddlx_17-0.30-5PGDG.rhel9.9.noarch.rpm) |
| `ddlx_17` | `0.30` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.1 KiB | [ddlx_17-0.30-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_17-0.30-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_17` | `0.32` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.5 KiB | [ddlx_17-0.32-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ddlx_17-0.32-1PGDG.rhel10.3.noarch.rpm) |
| `ddlx_17` | `0.31` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.3 KiB | [ddlx_17-0.31-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddlx_17-0.31-1PGSTY.el10.noarch.rpm) |
| `ddlx_17` | `0.30` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.2 KiB | [ddlx_17-0.30-5PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ddlx_17-0.30-5PGDG.rhel10.3.noarch.rpm) |
| `ddlx_17` | `0.32` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.5 KiB | [ddlx_17-0.32-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ddlx_17-0.32-1PGDG.rhel10.3.noarch.rpm) |
| `ddlx_17` | `0.32` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.5 KiB | [ddlx_17-0.32-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ddlx_17-0.32-1PGDG.rhel10.2.noarch.rpm) |
| `ddlx_17` | `0.31` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.3 KiB | [ddlx_17-0.31-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddlx_17-0.31-1PGSTY.el10.noarch.rpm) |
| `ddlx_17` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [ddlx_17-0.30-5PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ddlx_17-0.30-5PGDG.rhel10.3.noarch.rpm) |
| `ddlx_17` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [ddlx_17-0.30-5PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ddlx_17-0.30-5PGDG.rhel10.2.noarch.rpm) |
| `ddlx_17` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.7 KiB | [ddlx_17-0.30-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ddlx_17-0.30-2PGDG.rhel10.2.noarch.rpm) |
| `ddlx_17` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [ddlx_17-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ddlx_17-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_17` | `0.29` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [ddlx_17-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ddlx_17-0.29-1PGDG.rhel10.noarch.rpm) |
| `postgresql-17-ddlx` | `0.31` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.9 KiB | [postgresql-17-ddlx_0.31-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-17-ddlx_0.31-1PGSTY~bookworm_all.deb) |
| `postgresql-17-ddlx` | `0.31` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.9 KiB | [postgresql-17-ddlx_0.31-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-17-ddlx_0.31-1PGSTY~bookworm_all.deb) |
| `postgresql-17-ddlx` | `0.31` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.8 KiB | [postgresql-17-ddlx_0.31-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-17-ddlx_0.31-1PGSTY~trixie_all.deb) |
| `postgresql-17-ddlx` | `0.31` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.8 KiB | [postgresql-17-ddlx_0.31-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-17-ddlx_0.31-1PGSTY~trixie_all.deb) |
| `postgresql-17-ddlx` | `0.31` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.1 KiB | [postgresql-17-ddlx_0.31-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-17-ddlx_0.31-1PGSTY~jammy_all.deb) |
| `postgresql-17-ddlx` | `0.31` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.1 KiB | [postgresql-17-ddlx_0.31-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-17-ddlx_0.31-1PGSTY~jammy_all.deb) |
| `postgresql-17-ddlx` | `0.31` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.9 KiB | [postgresql-17-ddlx_0.31-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-17-ddlx_0.31-1PGSTY~noble_all.deb) |
| `postgresql-17-ddlx` | `0.31` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.9 KiB | [postgresql-17-ddlx_0.31-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-17-ddlx_0.31-1PGSTY~noble_all.deb) |
| `postgresql-17-ddlx` | `0.31` | [u26.x86_64](/os/u26.x86_64) | pigsty | 26.0 KiB | [postgresql-17-ddlx_0.31-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddlx/postgresql-17-ddlx_0.31-1PGSTY~resolute_all.deb) |
| `postgresql-17-ddlx` | `0.31` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.0 KiB | [postgresql-17-ddlx_0.31-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddlx/postgresql-17-ddlx_0.31-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddlx_16` | `0.32` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.0 KiB | [ddlx_16-0.32-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ddlx_16-0.32-1PGDG.rhel8.10.noarch.rpm) |
| `ddlx_16` | `0.31` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.4 KiB | [ddlx_16-0.31-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_16-0.31-1PGSTY.el8.noarch.rpm) |
| `ddlx_16` | `0.30` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [ddlx_16-0.30-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ddlx_16-0.30-5PGDG.rhel8.10.noarch.rpm) |
| `ddlx_16` | `0.30` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.3 KiB | [ddlx_16-0.30-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_16-0.30-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_16` | `0.30` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.9 KiB | [ddlx_16-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ddlx_16-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | `0.29` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.8 KiB | [ddlx_16-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ddlx_16-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | `0.27` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [ddlx_16-0.27-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ddlx_16-0.27-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | `0.26` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.9 KiB | [ddlx_16-0.26-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ddlx_16-0.26-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | `0.24` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.4 KiB | [ddlx_16-0.24-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ddlx_16-0.24-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | `0.32` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.9 KiB | [ddlx_16-0.32-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ddlx_16-0.32-1PGDG.rhel8.10.noarch.rpm) |
| `ddlx_16` | `0.31` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.4 KiB | [ddlx_16-0.31-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_16-0.31-1PGSTY.el8.noarch.rpm) |
| `ddlx_16` | `0.30` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.6 KiB | [ddlx_16-0.30-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ddlx_16-0.30-5PGDG.rhel8.10.noarch.rpm) |
| `ddlx_16` | `0.30` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.3 KiB | [ddlx_16-0.30-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_16-0.30-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_16` | `0.30` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.8 KiB | [ddlx_16-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ddlx_16-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | `0.29` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.8 KiB | [ddlx_16-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ddlx_16-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | `0.27` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.9 KiB | [ddlx_16-0.27-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ddlx_16-0.27-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | `0.26` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.9 KiB | [ddlx_16-0.26-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ddlx_16-0.26-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | `0.24` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.4 KiB | [ddlx_16-0.24-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ddlx_16-0.24-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | `0.32` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [ddlx_16-0.32-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ddlx_16-0.32-1PGDG.rhel9.9.noarch.rpm) |
| `ddlx_16` | `0.31` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.1 KiB | [ddlx_16-0.31-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_16-0.31-1PGSTY.el9.noarch.rpm) |
| `ddlx_16` | `0.30` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.0 KiB | [ddlx_16-0.30-5PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ddlx_16-0.30-5PGDG.rhel9.9.noarch.rpm) |
| `ddlx_16` | `0.30` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.2 KiB | [ddlx_16-0.30-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_16-0.30-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_16` | `0.32` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.3 KiB | [ddlx_16-0.32-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ddlx_16-0.32-1PGDG.rhel9.9.noarch.rpm) |
| `ddlx_16` | `0.31` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.1 KiB | [ddlx_16-0.31-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_16-0.31-1PGSTY.el9.noarch.rpm) |
| `ddlx_16` | `0.30` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.0 KiB | [ddlx_16-0.30-5PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ddlx_16-0.30-5PGDG.rhel9.9.noarch.rpm) |
| `ddlx_16` | `0.30` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.2 KiB | [ddlx_16-0.30-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_16-0.30-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_16` | `0.32` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.5 KiB | [ddlx_16-0.32-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ddlx_16-0.32-1PGDG.rhel10.3.noarch.rpm) |
| `ddlx_16` | `0.31` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.3 KiB | [ddlx_16-0.31-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddlx_16-0.31-1PGSTY.el10.noarch.rpm) |
| `ddlx_16` | `0.30` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.2 KiB | [ddlx_16-0.30-5PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ddlx_16-0.30-5PGDG.rhel10.3.noarch.rpm) |
| `ddlx_16` | `0.32` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.4 KiB | [ddlx_16-0.32-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ddlx_16-0.32-1PGDG.rhel10.3.noarch.rpm) |
| `ddlx_16` | `0.32` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.4 KiB | [ddlx_16-0.32-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ddlx_16-0.32-1PGDG.rhel10.2.noarch.rpm) |
| `ddlx_16` | `0.31` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.2 KiB | [ddlx_16-0.31-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddlx_16-0.31-1PGSTY.el10.noarch.rpm) |
| `ddlx_16` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [ddlx_16-0.30-5PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ddlx_16-0.30-5PGDG.rhel10.3.noarch.rpm) |
| `ddlx_16` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [ddlx_16-0.30-5PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ddlx_16-0.30-5PGDG.rhel10.2.noarch.rpm) |
| `ddlx_16` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.7 KiB | [ddlx_16-0.30-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ddlx_16-0.30-2PGDG.rhel10.2.noarch.rpm) |
| `ddlx_16` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [ddlx_16-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ddlx_16-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_16` | `0.29` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [ddlx_16-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ddlx_16-0.29-1PGDG.rhel10.noarch.rpm) |
| `postgresql-16-ddlx` | `0.31` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.8 KiB | [postgresql-16-ddlx_0.31-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-16-ddlx_0.31-1PGSTY~bookworm_all.deb) |
| `postgresql-16-ddlx` | `0.31` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.8 KiB | [postgresql-16-ddlx_0.31-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-16-ddlx_0.31-1PGSTY~bookworm_all.deb) |
| `postgresql-16-ddlx` | `0.31` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.8 KiB | [postgresql-16-ddlx_0.31-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-16-ddlx_0.31-1PGSTY~trixie_all.deb) |
| `postgresql-16-ddlx` | `0.31` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.8 KiB | [postgresql-16-ddlx_0.31-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-16-ddlx_0.31-1PGSTY~trixie_all.deb) |
| `postgresql-16-ddlx` | `0.31` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.0 KiB | [postgresql-16-ddlx_0.31-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-16-ddlx_0.31-1PGSTY~jammy_all.deb) |
| `postgresql-16-ddlx` | `0.31` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.0 KiB | [postgresql-16-ddlx_0.31-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-16-ddlx_0.31-1PGSTY~jammy_all.deb) |
| `postgresql-16-ddlx` | `0.31` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.9 KiB | [postgresql-16-ddlx_0.31-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-16-ddlx_0.31-1PGSTY~noble_all.deb) |
| `postgresql-16-ddlx` | `0.31` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.9 KiB | [postgresql-16-ddlx_0.31-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-16-ddlx_0.31-1PGSTY~noble_all.deb) |
| `postgresql-16-ddlx` | `0.31` | [u26.x86_64](/os/u26.x86_64) | pigsty | 25.9 KiB | [postgresql-16-ddlx_0.31-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddlx/postgresql-16-ddlx_0.31-1PGSTY~resolute_all.deb) |
| `postgresql-16-ddlx` | `0.31` | [u26.aarch64](/os/u26.aarch64) | pigsty | 25.9 KiB | [postgresql-16-ddlx_0.31-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddlx/postgresql-16-ddlx_0.31-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddlx_15` | `0.32` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.9 KiB | [ddlx_15-0.32-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.32-1PGDG.rhel8.10.noarch.rpm) |
| `ddlx_15` | `0.31` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.3 KiB | [ddlx_15-0.31-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_15-0.31-1PGSTY.el8.noarch.rpm) |
| `ddlx_15` | `0.30` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.6 KiB | [ddlx_15-0.30-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.30-5PGDG.rhel8.10.noarch.rpm) |
| `ddlx_15` | `0.30` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.3 KiB | [ddlx_15-0.30-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_15-0.30-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_15` | `0.30` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.8 KiB | [ddlx_15-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | `0.29` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.8 KiB | [ddlx_15-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | `0.27` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [ddlx_15-0.27-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.27-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | `0.26` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.9 KiB | [ddlx_15-0.26-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.26-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | `0.24` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.4 KiB | [ddlx_15-0.24-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.24-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | `0.23` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.2 KiB | [ddlx_15-0.23-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.23-1.rhel8.noarch.rpm) |
| `ddlx_15` | `0.22` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.8 KiB | [ddlx_15-0.22-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.22-1.rhel8.noarch.rpm) |
| `ddlx_15` | `0.32` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.9 KiB | [ddlx_15-0.32-1PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.32-1PGDG.rhel8.10.noarch.rpm) |
| `ddlx_15` | `0.31` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.3 KiB | [ddlx_15-0.31-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_15-0.31-1PGSTY.el8.noarch.rpm) |
| `ddlx_15` | `0.30` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.5 KiB | [ddlx_15-0.30-5PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.30-5PGDG.rhel8.10.noarch.rpm) |
| `ddlx_15` | `0.30` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.2 KiB | [ddlx_15-0.30-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_15-0.30-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_15` | `0.30` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.8 KiB | [ddlx_15-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | `0.29` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.7 KiB | [ddlx_15-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | `0.27` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.9 KiB | [ddlx_15-0.27-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.27-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | `0.26` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.9 KiB | [ddlx_15-0.26-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.26-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | `0.24` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.4 KiB | [ddlx_15-0.24-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.24-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | `0.23` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.2 KiB | [ddlx_15-0.23-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.23-1.rhel8.noarch.rpm) |
| `ddlx_15` | `0.22` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.8 KiB | [ddlx_15-0.22-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.22-1.rhel8.noarch.rpm) |
| `ddlx_15` | `0.32` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.3 KiB | [ddlx_15-0.32-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ddlx_15-0.32-1PGDG.rhel9.9.noarch.rpm) |
| `ddlx_15` | `0.31` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.1 KiB | [ddlx_15-0.31-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_15-0.31-1PGSTY.el9.noarch.rpm) |
| `ddlx_15` | `0.30` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.0 KiB | [ddlx_15-0.30-5PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ddlx_15-0.30-5PGDG.rhel9.9.noarch.rpm) |
| `ddlx_15` | `0.30` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.1 KiB | [ddlx_15-0.30-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_15-0.30-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_15` | `0.32` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.2 KiB | [ddlx_15-0.32-1PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ddlx_15-0.32-1PGDG.rhel9.9.noarch.rpm) |
| `ddlx_15` | `0.31` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.0 KiB | [ddlx_15-0.31-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_15-0.31-1PGSTY.el9.noarch.rpm) |
| `ddlx_15` | `0.30` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.9 KiB | [ddlx_15-0.30-5PGDG.rhel9.9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ddlx_15-0.30-5PGDG.rhel9.9.noarch.rpm) |
| `ddlx_15` | `0.30` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.1 KiB | [ddlx_15-0.30-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_15-0.30-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_15` | `0.32` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.4 KiB | [ddlx_15-0.32-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ddlx_15-0.32-1PGDG.rhel10.3.noarch.rpm) |
| `ddlx_15` | `0.31` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.2 KiB | [ddlx_15-0.31-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddlx_15-0.31-1PGSTY.el10.noarch.rpm) |
| `ddlx_15` | `0.30` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.1 KiB | [ddlx_15-0.30-5PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ddlx_15-0.30-5PGDG.rhel10.3.noarch.rpm) |
| `ddlx_15` | `0.32` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [ddlx_15-0.32-1PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ddlx_15-0.32-1PGDG.rhel10.3.noarch.rpm) |
| `ddlx_15` | `0.32` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.4 KiB | [ddlx_15-0.32-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ddlx_15-0.32-1PGDG.rhel10.2.noarch.rpm) |
| `ddlx_15` | `0.31` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.2 KiB | [ddlx_15-0.31-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddlx_15-0.31-1PGSTY.el10.noarch.rpm) |
| `ddlx_15` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.0 KiB | [ddlx_15-0.30-5PGDG.rhel10.3.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ddlx_15-0.30-5PGDG.rhel10.3.noarch.rpm) |
| `ddlx_15` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [ddlx_15-0.30-5PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ddlx_15-0.30-5PGDG.rhel10.2.noarch.rpm) |
| `ddlx_15` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.6 KiB | [ddlx_15-0.30-2PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ddlx_15-0.30-2PGDG.rhel10.2.noarch.rpm) |
| `ddlx_15` | `0.30` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.1 KiB | [ddlx_15-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ddlx_15-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_15` | `0.29` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [ddlx_15-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ddlx_15-0.29-1PGDG.rhel10.noarch.rpm) |
| `postgresql-15-ddlx` | `0.31` | [d12.x86_64](/os/d12.x86_64) | pigsty | 28.8 KiB | [postgresql-15-ddlx_0.31-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-15-ddlx_0.31-1PGSTY~bookworm_all.deb) |
| `postgresql-15-ddlx` | `0.31` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.8 KiB | [postgresql-15-ddlx_0.31-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-15-ddlx_0.31-1PGSTY~bookworm_all.deb) |
| `postgresql-15-ddlx` | `0.31` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.8 KiB | [postgresql-15-ddlx_0.31-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-15-ddlx_0.31-1PGSTY~trixie_all.deb) |
| `postgresql-15-ddlx` | `0.31` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.8 KiB | [postgresql-15-ddlx_0.31-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-15-ddlx_0.31-1PGSTY~trixie_all.deb) |
| `postgresql-15-ddlx` | `0.31` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.0 KiB | [postgresql-15-ddlx_0.31-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-15-ddlx_0.31-1PGSTY~jammy_all.deb) |
| `postgresql-15-ddlx` | `0.31` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.0 KiB | [postgresql-15-ddlx_0.31-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-15-ddlx_0.31-1PGSTY~jammy_all.deb) |
| `postgresql-15-ddlx` | `0.31` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.9 KiB | [postgresql-15-ddlx_0.31-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-15-ddlx_0.31-1PGSTY~noble_all.deb) |
| `postgresql-15-ddlx` | `0.31` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.9 KiB | [postgresql-15-ddlx_0.31-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-15-ddlx_0.31-1PGSTY~noble_all.deb) |
| `postgresql-15-ddlx` | `0.31` | [u26.x86_64](/os/u26.x86_64) | pigsty | 25.8 KiB | [postgresql-15-ddlx_0.31-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddlx/postgresql-15-ddlx_0.31-1PGSTY~resolute_all.deb) |
| `postgresql-15-ddlx` | `0.31` | [u26.aarch64](/os/u26.aarch64) | pigsty | 25.8 KiB | [postgresql-15-ddlx_0.31-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddlx/postgresql-15-ddlx_0.31-1PGSTY~resolute_all.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/lacanoid/pgddl" title="Repository" icon="github" subtitle="github.com/lacanoid/pgddl" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgddl-0.32.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_ddlx;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_ddlx;		# install via package name, for the active PG version
pig install ddlx;		# install by extension name, for the current active PG version

pig install ddlx -v 19;   # install for PG 19
pig install ddlx -v 18;   # install for PG 18
pig install ddlx -v 17;   # install for PG 17
pig install ddlx -v 16;   # install for PG 16
pig install ddlx -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION ddlx;
```




## Usage

> [ddlx: DDL eXtractor functions](https://github.com/lacanoid/pgddl)

ddlx is a SQL-only extension that generates DDL scripts from PostgreSQL system catalogs. It provides three main functions that accept an OID and work with all `reg*` object identifier types.

### Core Functions

```sql
-- Generate CREATE statement for an object
SELECT ddlx_create('my_table'::regclass);
SELECT ddlx_create('my_type'::regtype);
SELECT ddlx_create('my_function'::regproc);
SELECT ddlx_create(current_role::regrole);

-- Generate DROP statement
SELECT ddlx_drop('my_table'::regclass);

-- Generate full DDL script with dependency tree
SELECT ddlx_script('my_table'::regclass);
SELECT ddlx_script('my_enum');
SELECT ddlx_script(current_role::regrole);
```

### Options

Options are passed as a text array (e.g., `'{ine,nodcl}'`):

```sql
SELECT ddlx_create('my_table'::regclass, '{ine}');        -- add IF NOT EXISTS
SELECT ddlx_create('my_type'::regtype, '{noowner}');       -- omit ALTER SET OWNER
SELECT ddlx_script('my_table'::regclass, '{drop}');        -- include DROP statements
```

Available options: `drop`, `nodrop`, `owner`, `noowner`, `nogrants`, `nodcl`, `noalter`, `ine` (IF NOT EXISTS), `ie` (IF EXISTS), `ext`, `lite`, `nowrap`, `nopartitions`, `comments`, `nocomments`, `nostorage`, `noconstraints`, `noindexes`, `nosettings`, `notriggers`, `grantor`, `data`.

### For Objects Without reg* Types

```sql
SELECT ddlx_create(oid) FROM pg_foreign_data_wrapper WHERE fdwname = 'postgres_fdw';
SELECT ddlx_create(oid) FROM pg_database WHERE datname = current_database();
```

### Additional Functions

```sql
-- Identify any object by OID
SELECT * FROM ddlx_identify(oid);

-- Describe columns of a class
SELECT * FROM ddlx_describe('my_table'::regclass);

-- Get individual definition parts
SELECT * FROM ddlx_definitions(oid);

-- Generate pre-data creation statements only
SELECT ddlx_createonly('my_table'::regclass);

-- Generate post-data alteration statements
SELECT ddlx_alter('my_table'::regclass);

-- Search function/view bodies by regex
SELECT ddlx_create(objid) FROM ddlx_apropos('users');

-- Get GRANT statements
SELECT ddlx_grants('my_table'::regclass);
```
