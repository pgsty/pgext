---
title: "ddlx"
linkTitle: "ddlx"
description: "DDL eXtractor functions"
weight: 5080
categories: ["ADMIN"]
width: full
---

DDL eXtractor functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5080** | {{< badge content="ddlx" link="https://github.com/lacanoid/pgddl" >}} | {{< ext "ddlx" "pg_ddlx" >}} | `0.30` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgdd" >}} {{< ext "pg_checksums" >}} {{< ext "pg_permissions" >}} {{< ext "pgextwlist" >}} {{< ext "pg_catcheck" >}} {{< ext "adminpack" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/ddlx" >}} | `0.30` | {{< bg "18" "ddlx_18" "green" >}} {{< bg "17" "ddlx_17" "green" >}} {{< bg "16" "ddlx_16" "green" >}} {{< bg "15" "ddlx_15" "green" >}} {{< bg "14" "ddlx_14" "green" >}} {{< bg "13" "ddlx_13" "green" >}} | `ddlx_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/ddlx" >}} | `0.30` | {{< bg "18" "postgresql-18-ddlx" "green" >}} {{< bg "17" "postgresql-17-ddlx" "green" >}} {{< bg "16" "postgresql-16-ddlx" "green" >}} {{< bg "15" "postgresql-15-ddlx" "green" >}} {{< bg "14" "postgresql-14-ddlx" "green" >}} {{< bg "13" "postgresql-13-ddlx" "green" >}} | `postgresql-$v-ddlx` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.30" "ddlx_18 : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 0.30" "ddlx_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_15 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_14 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_13 : AVAIL 13" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.30" "ddlx_18 : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 0.30" "ddlx_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_15 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_14 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_13 : AVAIL 9" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.30" "ddlx_18 : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 0.30" "ddlx_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_15 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_14 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_13 : AVAIL 8" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.30" "ddlx_18 : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 0.30" "ddlx_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_15 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_14 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 0.30" "ddlx_13 : AVAIL 9" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.30" "ddlx_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.30" "ddlx_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.30" "ddlx_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.30" "ddlx_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.30" "ddlx_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.30" "ddlx_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.30" "ddlx_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.30" "ddlx_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.30" "ddlx_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.30" "ddlx_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.30" "ddlx_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.30" "ddlx_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.30" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-15-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-14-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-13-ddlx : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.30" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-15-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-14-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-13-ddlx : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.30" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-15-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-14-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-13-ddlx : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.30" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-15-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-14-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-13-ddlx : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.30" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-15-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-14-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-13-ddlx : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.30" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-15-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-14-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-13-ddlx : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.30" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-15-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-14-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-13-ddlx : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.30" "postgresql-18-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-17-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-16-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-15-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-14-ddlx : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.30" "postgresql-13-ddlx : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddlx_18` | 0.30 | `el8.x86_64` | pgdg | 33.9 KiB | [ddlx_18-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/ddlx_18-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_18` | 0.30 | `el8.aarch64` | pgdg | 33.8 KiB | [ddlx_18-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/ddlx_18-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_18` | 0.30 | `el9.x86_64` | pgdg | 31.7 KiB | [ddlx_18-0.30-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ddlx_18-0.30-1PGDG.rhel9.noarch.rpm) |
| `ddlx_18` | 0.30 | `el9.aarch64` | pgdg | 31.7 KiB | [ddlx_18-0.30-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ddlx_18-0.30-1PGDG.rhel9.noarch.rpm) |
| `ddlx_18` | 0.30 | `el10.x86_64` | pgdg | 32.2 KiB | [ddlx_18-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ddlx_18-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_18` | 0.30 | `el10.aarch64` | pgdg | 32.2 KiB | [ddlx_18-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ddlx_18-0.30-1PGDG.rhel10.noarch.rpm) |
| `postgresql-18-ddlx` | 0.30 | `d12.x86_64` | pigsty | 28.8 KiB | [postgresql-18-ddlx_0.30-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-18-ddlx_0.30-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-ddlx` | 0.30 | `d12.aarch64` | pigsty | 28.8 KiB | [postgresql-18-ddlx_0.30-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-18-ddlx_0.30-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-ddlx` | 0.30 | `d13.x86_64` | pigsty | 28.7 KiB | [postgresql-18-ddlx_0.30-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-18-ddlx_0.30-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-ddlx` | 0.30 | `d13.aarch64` | pigsty | 28.7 KiB | [postgresql-18-ddlx_0.30-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-18-ddlx_0.30-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-ddlx` | 0.30 | `u22.x86_64` | pigsty | 26.0 KiB | [postgresql-18-ddlx_0.30-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-18-ddlx_0.30-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-ddlx` | 0.30 | `u22.aarch64` | pigsty | 26.0 KiB | [postgresql-18-ddlx_0.30-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-18-ddlx_0.30-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-ddlx` | 0.30 | `u24.x86_64` | pigsty | 25.9 KiB | [postgresql-18-ddlx_0.30-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-18-ddlx_0.30-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-ddlx` | 0.30 | `u24.aarch64` | pigsty | 25.9 KiB | [postgresql-18-ddlx_0.30-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-18-ddlx_0.30-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddlx_17` | 0.30 | `el8.x86_64` | pigsty | 32.4 KiB | [ddlx_17-0.30-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_17-0.30-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_17` | 0.30 | `el8.x86_64` | pgdg | 33.9 KiB | [ddlx_17-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ddlx_17-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_17` | 0.29 | `el8.x86_64` | pigsty | 31.4 KiB | [ddlx_17-0.29-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_17-0.29-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_17` | 0.29 | `el8.x86_64` | pgdg | 32.8 KiB | [ddlx_17-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ddlx_17-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_17` | 0.30 | `el8.aarch64` | pigsty | 32.3 KiB | [ddlx_17-0.30-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_17-0.30-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_17` | 0.30 | `el8.aarch64` | pgdg | 33.8 KiB | [ddlx_17-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ddlx_17-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_17` | 0.29 | `el8.aarch64` | pigsty | 31.4 KiB | [ddlx_17-0.29-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_17-0.29-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_17` | 0.29 | `el8.aarch64` | pgdg | 32.8 KiB | [ddlx_17-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ddlx_17-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_17` | 0.30 | `el9.x86_64` | pigsty | 31.2 KiB | [ddlx_17-0.30-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_17-0.30-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_17` | 0.30 | `el9.x86_64` | pgdg | 31.7 KiB | [ddlx_17-0.30-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ddlx_17-0.30-1PGDG.rhel9.noarch.rpm) |
| `ddlx_17` | 0.29 | `el9.x86_64` | pigsty | 30.3 KiB | [ddlx_17-0.29-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_17-0.29-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_17` | 0.29 | `el9.x86_64` | pgdg | 30.9 KiB | [ddlx_17-0.29-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ddlx_17-0.29-1PGDG.rhel9.noarch.rpm) |
| `ddlx_17` | 0.30 | `el9.aarch64` | pigsty | 31.1 KiB | [ddlx_17-0.30-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_17-0.30-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_17` | 0.30 | `el9.aarch64` | pgdg | 31.7 KiB | [ddlx_17-0.30-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ddlx_17-0.30-1PGDG.rhel9.noarch.rpm) |
| `ddlx_17` | 0.29 | `el9.aarch64` | pigsty | 30.3 KiB | [ddlx_17-0.29-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_17-0.29-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_17` | 0.29 | `el9.aarch64` | pgdg | 30.8 KiB | [ddlx_17-0.29-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ddlx_17-0.29-1PGDG.rhel9.noarch.rpm) |
| `ddlx_17` | 0.30 | `el10.x86_64` | pgdg | 32.2 KiB | [ddlx_17-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ddlx_17-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_17` | 0.29 | `el10.x86_64` | pgdg | 31.4 KiB | [ddlx_17-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ddlx_17-0.29-1PGDG.rhel10.noarch.rpm) |
| `ddlx_17` | 0.30 | `el10.aarch64` | pgdg | 32.2 KiB | [ddlx_17-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ddlx_17-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_17` | 0.29 | `el10.aarch64` | pgdg | 31.3 KiB | [ddlx_17-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ddlx_17-0.29-1PGDG.rhel10.noarch.rpm) |
| `postgresql-17-ddlx` | 0.30 | `d12.x86_64` | pigsty | 28.7 KiB | [postgresql-17-ddlx_0.30-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-17-ddlx_0.30-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-ddlx` | 0.30 | `d12.aarch64` | pigsty | 28.7 KiB | [postgresql-17-ddlx_0.30-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-17-ddlx_0.30-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-ddlx` | 0.30 | `d13.x86_64` | pigsty | 28.7 KiB | [postgresql-17-ddlx_0.30-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-17-ddlx_0.30-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-ddlx` | 0.30 | `d13.aarch64` | pigsty | 28.7 KiB | [postgresql-17-ddlx_0.30-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-17-ddlx_0.30-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-ddlx` | 0.30 | `u22.x86_64` | pigsty | 26.0 KiB | [postgresql-17-ddlx_0.30-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-17-ddlx_0.30-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-ddlx` | 0.30 | `u22.aarch64` | pigsty | 26.0 KiB | [postgresql-17-ddlx_0.30-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-17-ddlx_0.30-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-ddlx` | 0.30 | `u24.x86_64` | pigsty | 25.9 KiB | [postgresql-17-ddlx_0.30-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-17-ddlx_0.30-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-ddlx` | 0.30 | `u24.aarch64` | pigsty | 25.9 KiB | [postgresql-17-ddlx_0.30-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-17-ddlx_0.30-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddlx_16` | 0.30 | `el8.x86_64` | pigsty | 32.3 KiB | [ddlx_16-0.30-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_16-0.30-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_16` | 0.30 | `el8.x86_64` | pgdg | 33.9 KiB | [ddlx_16-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ddlx_16-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | 0.29 | `el8.x86_64` | pigsty | 31.4 KiB | [ddlx_16-0.29-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_16-0.29-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_16` | 0.29 | `el8.x86_64` | pgdg | 32.8 KiB | [ddlx_16-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ddlx_16-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | 0.27 | `el8.x86_64` | pgdg | 32.0 KiB | [ddlx_16-0.27-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ddlx_16-0.27-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | 0.26 | `el8.x86_64` | pgdg | 30.9 KiB | [ddlx_16-0.26-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ddlx_16-0.26-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | 0.24 | `el8.x86_64` | pgdg | 30.4 KiB | [ddlx_16-0.24-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ddlx_16-0.24-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | 0.30 | `el8.aarch64` | pigsty | 32.3 KiB | [ddlx_16-0.30-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_16-0.30-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_16` | 0.30 | `el8.aarch64` | pgdg | 33.8 KiB | [ddlx_16-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ddlx_16-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | 0.29 | `el8.aarch64` | pigsty | 31.4 KiB | [ddlx_16-0.29-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_16-0.29-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_16` | 0.29 | `el8.aarch64` | pgdg | 32.8 KiB | [ddlx_16-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ddlx_16-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | 0.27 | `el8.aarch64` | pgdg | 31.9 KiB | [ddlx_16-0.27-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ddlx_16-0.27-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | 0.26 | `el8.aarch64` | pgdg | 30.9 KiB | [ddlx_16-0.26-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ddlx_16-0.26-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | 0.24 | `el8.aarch64` | pgdg | 30.4 KiB | [ddlx_16-0.24-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ddlx_16-0.24-1PGDG.rhel8.noarch.rpm) |
| `ddlx_16` | 0.30 | `el9.x86_64` | pigsty | 31.2 KiB | [ddlx_16-0.30-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_16-0.30-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_16` | 0.30 | `el9.x86_64` | pgdg | 31.7 KiB | [ddlx_16-0.30-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ddlx_16-0.30-1PGDG.rhel9.noarch.rpm) |
| `ddlx_16` | 0.29 | `el9.x86_64` | pigsty | 30.3 KiB | [ddlx_16-0.29-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_16-0.29-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_16` | 0.29 | `el9.x86_64` | pgdg | 30.9 KiB | [ddlx_16-0.29-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ddlx_16-0.29-1PGDG.rhel9.noarch.rpm) |
| `ddlx_16` | 0.27 | `el9.x86_64` | pgdg | 30.2 KiB | [ddlx_16-0.27-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ddlx_16-0.27-1PGDG.rhel9.noarch.rpm) |
| `ddlx_16` | 0.26 | `el9.x86_64` | pgdg | 29.2 KiB | [ddlx_16-0.26-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ddlx_16-0.26-1PGDG.rhel9.noarch.rpm) |
| `ddlx_16` | 0.24 | `el9.x86_64` | pgdg | 28.8 KiB | [ddlx_16-0.24-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ddlx_16-0.24-1PGDG.rhel9.noarch.rpm) |
| `ddlx_16` | 0.30 | `el9.aarch64` | pigsty | 31.2 KiB | [ddlx_16-0.30-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_16-0.30-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_16` | 0.30 | `el9.aarch64` | pgdg | 31.6 KiB | [ddlx_16-0.30-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ddlx_16-0.30-1PGDG.rhel9.noarch.rpm) |
| `ddlx_16` | 0.29 | `el9.aarch64` | pigsty | 30.3 KiB | [ddlx_16-0.29-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_16-0.29-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_16` | 0.29 | `el9.aarch64` | pgdg | 30.8 KiB | [ddlx_16-0.29-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ddlx_16-0.29-1PGDG.rhel9.noarch.rpm) |
| `ddlx_16` | 0.27 | `el9.aarch64` | pgdg | 30.0 KiB | [ddlx_16-0.27-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ddlx_16-0.27-1PGDG.rhel9.noarch.rpm) |
| `ddlx_16` | 0.26 | `el9.aarch64` | pgdg | 29.0 KiB | [ddlx_16-0.26-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ddlx_16-0.26-1PGDG.rhel9.noarch.rpm) |
| `ddlx_16` | 0.24 | `el9.aarch64` | pgdg | 28.6 KiB | [ddlx_16-0.24-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ddlx_16-0.24-1PGDG.rhel9.noarch.rpm) |
| `ddlx_16` | 0.30 | `el10.x86_64` | pgdg | 32.2 KiB | [ddlx_16-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ddlx_16-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_16` | 0.29 | `el10.x86_64` | pgdg | 31.4 KiB | [ddlx_16-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ddlx_16-0.29-1PGDG.rhel10.noarch.rpm) |
| `ddlx_16` | 0.30 | `el10.aarch64` | pgdg | 32.1 KiB | [ddlx_16-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ddlx_16-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_16` | 0.29 | `el10.aarch64` | pgdg | 31.3 KiB | [ddlx_16-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ddlx_16-0.29-1PGDG.rhel10.noarch.rpm) |
| `postgresql-16-ddlx` | 0.30 | `d12.x86_64` | pigsty | 28.7 KiB | [postgresql-16-ddlx_0.30-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-16-ddlx_0.30-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-ddlx` | 0.30 | `d12.aarch64` | pigsty | 28.7 KiB | [postgresql-16-ddlx_0.30-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-16-ddlx_0.30-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-ddlx` | 0.30 | `d13.x86_64` | pigsty | 28.7 KiB | [postgresql-16-ddlx_0.30-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-16-ddlx_0.30-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-ddlx` | 0.30 | `d13.aarch64` | pigsty | 28.7 KiB | [postgresql-16-ddlx_0.30-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-16-ddlx_0.30-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-ddlx` | 0.30 | `u22.x86_64` | pigsty | 26.0 KiB | [postgresql-16-ddlx_0.30-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-16-ddlx_0.30-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-ddlx` | 0.30 | `u22.aarch64` | pigsty | 26.0 KiB | [postgresql-16-ddlx_0.30-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-16-ddlx_0.30-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-ddlx` | 0.30 | `u24.x86_64` | pigsty | 25.8 KiB | [postgresql-16-ddlx_0.30-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-16-ddlx_0.30-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-ddlx` | 0.30 | `u24.aarch64` | pigsty | 25.8 KiB | [postgresql-16-ddlx_0.30-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-16-ddlx_0.30-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddlx_15` | 0.30 | `el8.x86_64` | pigsty | 32.3 KiB | [ddlx_15-0.30-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_15-0.30-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_15` | 0.30 | `el8.x86_64` | pgdg | 33.8 KiB | [ddlx_15-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | 0.29 | `el8.x86_64` | pigsty | 31.3 KiB | [ddlx_15-0.29-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_15-0.29-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_15` | 0.29 | `el8.x86_64` | pgdg | 32.8 KiB | [ddlx_15-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | 0.27 | `el8.x86_64` | pgdg | 32.0 KiB | [ddlx_15-0.27-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.27-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | 0.26 | `el8.x86_64` | pgdg | 30.9 KiB | [ddlx_15-0.26-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.26-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | 0.24 | `el8.x86_64` | pgdg | 30.4 KiB | [ddlx_15-0.24-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.24-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | 0.23 | `el8.x86_64` | pgdg | 30.2 KiB | [ddlx_15-0.23-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.23-1.rhel8.noarch.rpm) |
| `ddlx_15` | 0.22 | `el8.x86_64` | pgdg | 29.8 KiB | [ddlx_15-0.22-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ddlx_15-0.22-1.rhel8.noarch.rpm) |
| `ddlx_15` | 0.30 | `el8.aarch64` | pigsty | 32.2 KiB | [ddlx_15-0.30-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_15-0.30-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_15` | 0.30 | `el8.aarch64` | pgdg | 33.8 KiB | [ddlx_15-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | 0.29 | `el8.aarch64` | pigsty | 31.3 KiB | [ddlx_15-0.29-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_15-0.29-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_15` | 0.29 | `el8.aarch64` | pgdg | 32.7 KiB | [ddlx_15-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | 0.27 | `el8.aarch64` | pgdg | 31.9 KiB | [ddlx_15-0.27-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.27-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | 0.26 | `el8.aarch64` | pgdg | 30.9 KiB | [ddlx_15-0.26-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.26-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | 0.24 | `el8.aarch64` | pgdg | 30.4 KiB | [ddlx_15-0.24-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.24-1PGDG.rhel8.noarch.rpm) |
| `ddlx_15` | 0.23 | `el8.aarch64` | pgdg | 30.2 KiB | [ddlx_15-0.23-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.23-1.rhel8.noarch.rpm) |
| `ddlx_15` | 0.22 | `el8.aarch64` | pgdg | 29.8 KiB | [ddlx_15-0.22-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ddlx_15-0.22-1.rhel8.noarch.rpm) |
| `ddlx_15` | 0.30 | `el9.x86_64` | pigsty | 31.1 KiB | [ddlx_15-0.30-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_15-0.30-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_15` | 0.30 | `el9.x86_64` | pgdg | 31.6 KiB | [ddlx_15-0.30-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ddlx_15-0.30-1PGDG.rhel9.noarch.rpm) |
| `ddlx_15` | 0.29 | `el9.x86_64` | pigsty | 30.2 KiB | [ddlx_15-0.29-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_15-0.29-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_15` | 0.29 | `el9.x86_64` | pgdg | 30.8 KiB | [ddlx_15-0.29-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ddlx_15-0.29-1PGDG.rhel9.noarch.rpm) |
| `ddlx_15` | 0.27 | `el9.x86_64` | pgdg | 30.2 KiB | [ddlx_15-0.27-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ddlx_15-0.27-1PGDG.rhel9.noarch.rpm) |
| `ddlx_15` | 0.26 | `el9.x86_64` | pgdg | 29.2 KiB | [ddlx_15-0.26-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ddlx_15-0.26-1PGDG.rhel9.noarch.rpm) |
| `ddlx_15` | 0.24 | `el9.x86_64` | pgdg | 28.8 KiB | [ddlx_15-0.24-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ddlx_15-0.24-1PGDG.rhel9.noarch.rpm) |
| `ddlx_15` | 0.23 | `el9.x86_64` | pgdg | 28.8 KiB | [ddlx_15-0.23-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ddlx_15-0.23-1.rhel9.noarch.rpm) |
| `ddlx_15` | 0.22 | `el9.x86_64` | pgdg | 29.0 KiB | [ddlx_15-0.22-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ddlx_15-0.22-1.rhel9.noarch.rpm) |
| `ddlx_15` | 0.30 | `el9.aarch64` | pigsty | 31.1 KiB | [ddlx_15-0.30-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_15-0.30-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_15` | 0.30 | `el9.aarch64` | pgdg | 31.6 KiB | [ddlx_15-0.30-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ddlx_15-0.30-1PGDG.rhel9.noarch.rpm) |
| `ddlx_15` | 0.29 | `el9.aarch64` | pigsty | 30.2 KiB | [ddlx_15-0.29-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_15-0.29-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_15` | 0.29 | `el9.aarch64` | pgdg | 30.8 KiB | [ddlx_15-0.29-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ddlx_15-0.29-1PGDG.rhel9.noarch.rpm) |
| `ddlx_15` | 0.27 | `el9.aarch64` | pgdg | 30.0 KiB | [ddlx_15-0.27-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ddlx_15-0.27-1PGDG.rhel9.noarch.rpm) |
| `ddlx_15` | 0.26 | `el9.aarch64` | pgdg | 29.0 KiB | [ddlx_15-0.26-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ddlx_15-0.26-1PGDG.rhel9.noarch.rpm) |
| `ddlx_15` | 0.24 | `el9.aarch64` | pgdg | 28.6 KiB | [ddlx_15-0.24-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ddlx_15-0.24-1PGDG.rhel9.noarch.rpm) |
| `ddlx_15` | 0.23 | `el9.aarch64` | pgdg | 28.6 KiB | [ddlx_15-0.23-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ddlx_15-0.23-1.rhel9.noarch.rpm) |
| `ddlx_15` | 0.22 | `el9.aarch64` | pgdg | 28.8 KiB | [ddlx_15-0.22-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ddlx_15-0.22-1.rhel9.noarch.rpm) |
| `ddlx_15` | 0.30 | `el10.x86_64` | pgdg | 32.1 KiB | [ddlx_15-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ddlx_15-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_15` | 0.29 | `el10.x86_64` | pgdg | 31.3 KiB | [ddlx_15-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ddlx_15-0.29-1PGDG.rhel10.noarch.rpm) |
| `ddlx_15` | 0.30 | `el10.aarch64` | pgdg | 32.1 KiB | [ddlx_15-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ddlx_15-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_15` | 0.29 | `el10.aarch64` | pgdg | 31.3 KiB | [ddlx_15-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ddlx_15-0.29-1PGDG.rhel10.noarch.rpm) |
| `postgresql-15-ddlx` | 0.30 | `d12.x86_64` | pigsty | 28.7 KiB | [postgresql-15-ddlx_0.30-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-15-ddlx_0.30-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-ddlx` | 0.30 | `d12.aarch64` | pigsty | 28.7 KiB | [postgresql-15-ddlx_0.30-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-15-ddlx_0.30-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-ddlx` | 0.30 | `d13.x86_64` | pigsty | 28.7 KiB | [postgresql-15-ddlx_0.30-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-15-ddlx_0.30-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-ddlx` | 0.30 | `d13.aarch64` | pigsty | 28.7 KiB | [postgresql-15-ddlx_0.30-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-15-ddlx_0.30-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-ddlx` | 0.30 | `u22.x86_64` | pigsty | 25.9 KiB | [postgresql-15-ddlx_0.30-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-15-ddlx_0.30-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-ddlx` | 0.30 | `u22.aarch64` | pigsty | 25.9 KiB | [postgresql-15-ddlx_0.30-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-15-ddlx_0.30-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-ddlx` | 0.30 | `u24.x86_64` | pigsty | 25.8 KiB | [postgresql-15-ddlx_0.30-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-15-ddlx_0.30-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-ddlx` | 0.30 | `u24.aarch64` | pigsty | 25.8 KiB | [postgresql-15-ddlx_0.30-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-15-ddlx_0.30-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddlx_14` | 0.30 | `el8.x86_64` | pigsty | 32.1 KiB | [ddlx_14-0.30-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_14-0.30-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_14` | 0.30 | `el8.x86_64` | pgdg | 33.7 KiB | [ddlx_14-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ddlx_14-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_14` | 0.29 | `el8.x86_64` | pigsty | 31.2 KiB | [ddlx_14-0.29-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_14-0.29-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_14` | 0.29 | `el8.x86_64` | pgdg | 32.6 KiB | [ddlx_14-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ddlx_14-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_14` | 0.27 | `el8.x86_64` | pgdg | 31.9 KiB | [ddlx_14-0.27-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ddlx_14-0.27-1PGDG.rhel8.noarch.rpm) |
| `ddlx_14` | 0.26 | `el8.x86_64` | pgdg | 30.9 KiB | [ddlx_14-0.26-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ddlx_14-0.26-1PGDG.rhel8.noarch.rpm) |
| `ddlx_14` | 0.24 | `el8.x86_64` | pgdg | 30.4 KiB | [ddlx_14-0.24-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ddlx_14-0.24-1PGDG.rhel8.noarch.rpm) |
| `ddlx_14` | 0.23 | `el8.x86_64` | pgdg | 30.2 KiB | [ddlx_14-0.23-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ddlx_14-0.23-1.rhel8.noarch.rpm) |
| `ddlx_14` | 0.22 | `el8.x86_64` | pgdg | 29.9 KiB | [ddlx_14-0.22-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ddlx_14-0.22-1.rhel8.noarch.rpm) |
| `ddlx_14` | 0.30 | `el8.aarch64` | pigsty | 32.1 KiB | [ddlx_14-0.30-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_14-0.30-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_14` | 0.30 | `el8.aarch64` | pgdg | 33.6 KiB | [ddlx_14-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ddlx_14-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_14` | 0.29 | `el8.aarch64` | pigsty | 31.2 KiB | [ddlx_14-0.29-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_14-0.29-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_14` | 0.29 | `el8.aarch64` | pgdg | 32.6 KiB | [ddlx_14-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ddlx_14-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_14` | 0.27 | `el8.aarch64` | pgdg | 31.8 KiB | [ddlx_14-0.27-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ddlx_14-0.27-1PGDG.rhel8.noarch.rpm) |
| `ddlx_14` | 0.26 | `el8.aarch64` | pgdg | 30.9 KiB | [ddlx_14-0.26-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ddlx_14-0.26-1PGDG.rhel8.noarch.rpm) |
| `ddlx_14` | 0.24 | `el8.aarch64` | pgdg | 30.4 KiB | [ddlx_14-0.24-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ddlx_14-0.24-1PGDG.rhel8.noarch.rpm) |
| `ddlx_14` | 0.23 | `el8.aarch64` | pgdg | 30.2 KiB | [ddlx_14-0.23-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ddlx_14-0.23-1.rhel8.noarch.rpm) |
| `ddlx_14` | 0.22 | `el8.aarch64` | pgdg | 29.8 KiB | [ddlx_14-0.22-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ddlx_14-0.22-1.rhel8.noarch.rpm) |
| `ddlx_14` | 0.30 | `el9.x86_64` | pigsty | 31.0 KiB | [ddlx_14-0.30-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_14-0.30-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_14` | 0.30 | `el9.x86_64` | pgdg | 31.5 KiB | [ddlx_14-0.30-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ddlx_14-0.30-1PGDG.rhel9.noarch.rpm) |
| `ddlx_14` | 0.29 | `el9.x86_64` | pigsty | 30.1 KiB | [ddlx_14-0.29-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_14-0.29-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_14` | 0.29 | `el9.x86_64` | pgdg | 30.7 KiB | [ddlx_14-0.29-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ddlx_14-0.29-1PGDG.rhel9.noarch.rpm) |
| `ddlx_14` | 0.27 | `el9.x86_64` | pgdg | 30.1 KiB | [ddlx_14-0.27-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ddlx_14-0.27-1PGDG.rhel9.noarch.rpm) |
| `ddlx_14` | 0.26 | `el9.x86_64` | pgdg | 29.2 KiB | [ddlx_14-0.26-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ddlx_14-0.26-1PGDG.rhel9.noarch.rpm) |
| `ddlx_14` | 0.24 | `el9.x86_64` | pgdg | 28.8 KiB | [ddlx_14-0.24-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ddlx_14-0.24-1PGDG.rhel9.noarch.rpm) |
| `ddlx_14` | 0.23 | `el9.x86_64` | pgdg | 28.8 KiB | [ddlx_14-0.23-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ddlx_14-0.23-1.rhel9.noarch.rpm) |
| `ddlx_14` | 0.22 | `el9.x86_64` | pgdg | 29.0 KiB | [ddlx_14-0.22-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ddlx_14-0.22-1.rhel9.noarch.rpm) |
| `ddlx_14` | 0.30 | `el9.aarch64` | pigsty | 30.9 KiB | [ddlx_14-0.30-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_14-0.30-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_14` | 0.30 | `el9.aarch64` | pgdg | 31.4 KiB | [ddlx_14-0.30-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ddlx_14-0.30-1PGDG.rhel9.noarch.rpm) |
| `ddlx_14` | 0.29 | `el9.aarch64` | pigsty | 30.1 KiB | [ddlx_14-0.29-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_14-0.29-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_14` | 0.29 | `el9.aarch64` | pgdg | 30.6 KiB | [ddlx_14-0.29-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ddlx_14-0.29-1PGDG.rhel9.noarch.rpm) |
| `ddlx_14` | 0.27 | `el9.aarch64` | pgdg | 29.9 KiB | [ddlx_14-0.27-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ddlx_14-0.27-1PGDG.rhel9.noarch.rpm) |
| `ddlx_14` | 0.26 | `el9.aarch64` | pgdg | 29.0 KiB | [ddlx_14-0.26-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ddlx_14-0.26-1PGDG.rhel9.noarch.rpm) |
| `ddlx_14` | 0.24 | `el9.aarch64` | pgdg | 28.6 KiB | [ddlx_14-0.24-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ddlx_14-0.24-1PGDG.rhel9.noarch.rpm) |
| `ddlx_14` | 0.23 | `el9.aarch64` | pgdg | 28.6 KiB | [ddlx_14-0.23-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ddlx_14-0.23-1.rhel9.noarch.rpm) |
| `ddlx_14` | 0.22 | `el9.aarch64` | pgdg | 28.8 KiB | [ddlx_14-0.22-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ddlx_14-0.22-1.rhel9.noarch.rpm) |
| `ddlx_14` | 0.30 | `el10.x86_64` | pgdg | 32.0 KiB | [ddlx_14-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ddlx_14-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_14` | 0.29 | `el10.x86_64` | pgdg | 31.2 KiB | [ddlx_14-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ddlx_14-0.29-1PGDG.rhel10.noarch.rpm) |
| `ddlx_14` | 0.30 | `el10.aarch64` | pgdg | 31.9 KiB | [ddlx_14-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/ddlx_14-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_14` | 0.29 | `el10.aarch64` | pgdg | 31.1 KiB | [ddlx_14-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/ddlx_14-0.29-1PGDG.rhel10.noarch.rpm) |
| `postgresql-14-ddlx` | 0.30 | `d12.x86_64` | pigsty | 28.6 KiB | [postgresql-14-ddlx_0.30-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-14-ddlx_0.30-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-ddlx` | 0.30 | `d12.aarch64` | pigsty | 28.6 KiB | [postgresql-14-ddlx_0.30-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-14-ddlx_0.30-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-ddlx` | 0.30 | `d13.x86_64` | pigsty | 28.5 KiB | [postgresql-14-ddlx_0.30-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-14-ddlx_0.30-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-ddlx` | 0.30 | `d13.aarch64` | pigsty | 28.5 KiB | [postgresql-14-ddlx_0.30-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-14-ddlx_0.30-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-ddlx` | 0.30 | `u22.x86_64` | pigsty | 25.8 KiB | [postgresql-14-ddlx_0.30-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-14-ddlx_0.30-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-ddlx` | 0.30 | `u22.aarch64` | pigsty | 25.8 KiB | [postgresql-14-ddlx_0.30-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-14-ddlx_0.30-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-ddlx` | 0.30 | `u24.x86_64` | pigsty | 25.6 KiB | [postgresql-14-ddlx_0.30-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-14-ddlx_0.30-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-ddlx` | 0.30 | `u24.aarch64` | pigsty | 25.6 KiB | [postgresql-14-ddlx_0.30-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-14-ddlx_0.30-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddlx_13` | 0.30 | `el8.x86_64` | pigsty | 31.9 KiB | [ddlx_13-0.30-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_13-0.30-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_13` | 0.30 | `el8.x86_64` | pgdg | 33.4 KiB | [ddlx_13-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ddlx_13-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_13` | 0.29 | `el8.x86_64` | pigsty | 31.0 KiB | [ddlx_13-0.29-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddlx_13-0.29-1PIGSTY.el8.x86_64.rpm) |
| `ddlx_13` | 0.29 | `el8.x86_64` | pgdg | 32.4 KiB | [ddlx_13-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ddlx_13-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_13` | 0.27 | `el8.x86_64` | pgdg | 31.8 KiB | [ddlx_13-0.27-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ddlx_13-0.27-1PGDG.rhel8.noarch.rpm) |
| `ddlx_13` | 0.26 | `el8.x86_64` | pgdg | 30.9 KiB | [ddlx_13-0.26-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ddlx_13-0.26-1PGDG.rhel8.noarch.rpm) |
| `ddlx_13` | 0.24 | `el8.x86_64` | pgdg | 30.4 KiB | [ddlx_13-0.24-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ddlx_13-0.24-1PGDG.rhel8.noarch.rpm) |
| `ddlx_13` | 0.23 | `el8.x86_64` | pgdg | 30.2 KiB | [ddlx_13-0.23-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ddlx_13-0.23-1.rhel8.noarch.rpm) |
| `ddlx_13` | 0.22 | `el8.x86_64` | pgdg | 29.8 KiB | [ddlx_13-0.22-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ddlx_13-0.22-1.rhel8.noarch.rpm) |
| `ddlx_13` | 0.21 | `el8.x86_64` | pgdg | 29.2 KiB | [ddlx_13-0.21-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ddlx_13-0.21-1.rhel8.noarch.rpm) |
| `ddlx_13` | 0.19 | `el8.x86_64` | pgdg | 28.4 KiB | [ddlx_13-0.19-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ddlx_13-0.19-1.rhel8.noarch.rpm) |
| `ddlx_13` | 0.17 | `el8.x86_64` | pgdg | 28.1 KiB | [ddlx_13-0.17-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ddlx_13-0.17-1.rhel8.noarch.rpm) |
| `ddlx_13` | 0.16 | `el8.x86_64` | pgdg | 28.0 KiB | [ddlx_13-0.16-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ddlx_13-0.16-1.rhel8.noarch.rpm) |
| `ddlx_13` | 0.30 | `el8.aarch64` | pigsty | 31.8 KiB | [ddlx_13-0.30-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_13-0.30-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_13` | 0.30 | `el8.aarch64` | pgdg | 33.4 KiB | [ddlx_13-0.30-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ddlx_13-0.30-1PGDG.rhel8.noarch.rpm) |
| `ddlx_13` | 0.29 | `el8.aarch64` | pigsty | 31.0 KiB | [ddlx_13-0.29-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddlx_13-0.29-1PIGSTY.el8.aarch64.rpm) |
| `ddlx_13` | 0.29 | `el8.aarch64` | pgdg | 32.4 KiB | [ddlx_13-0.29-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ddlx_13-0.29-1PGDG.rhel8.noarch.rpm) |
| `ddlx_13` | 0.27 | `el8.aarch64` | pgdg | 31.8 KiB | [ddlx_13-0.27-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ddlx_13-0.27-1PGDG.rhel8.noarch.rpm) |
| `ddlx_13` | 0.26 | `el8.aarch64` | pgdg | 30.8 KiB | [ddlx_13-0.26-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ddlx_13-0.26-1PGDG.rhel8.noarch.rpm) |
| `ddlx_13` | 0.24 | `el8.aarch64` | pgdg | 30.3 KiB | [ddlx_13-0.24-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ddlx_13-0.24-1PGDG.rhel8.noarch.rpm) |
| `ddlx_13` | 0.23 | `el8.aarch64` | pgdg | 30.1 KiB | [ddlx_13-0.23-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ddlx_13-0.23-1.rhel8.noarch.rpm) |
| `ddlx_13` | 0.22 | `el8.aarch64` | pgdg | 29.8 KiB | [ddlx_13-0.22-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ddlx_13-0.22-1.rhel8.noarch.rpm) |
| `ddlx_13` | 0.30 | `el9.x86_64` | pigsty | 30.7 KiB | [ddlx_13-0.30-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_13-0.30-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_13` | 0.30 | `el9.x86_64` | pgdg | 31.2 KiB | [ddlx_13-0.30-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ddlx_13-0.30-1PGDG.rhel9.noarch.rpm) |
| `ddlx_13` | 0.29 | `el9.x86_64` | pigsty | 29.9 KiB | [ddlx_13-0.29-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddlx_13-0.29-1PIGSTY.el9.x86_64.rpm) |
| `ddlx_13` | 0.29 | `el9.x86_64` | pgdg | 30.5 KiB | [ddlx_13-0.29-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ddlx_13-0.29-1PGDG.rhel9.noarch.rpm) |
| `ddlx_13` | 0.27 | `el9.x86_64` | pgdg | 30.0 KiB | [ddlx_13-0.27-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ddlx_13-0.27-1PGDG.rhel9.noarch.rpm) |
| `ddlx_13` | 0.26 | `el9.x86_64` | pgdg | 29.2 KiB | [ddlx_13-0.26-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ddlx_13-0.26-1PGDG.rhel9.noarch.rpm) |
| `ddlx_13` | 0.24 | `el9.x86_64` | pgdg | 28.7 KiB | [ddlx_13-0.24-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ddlx_13-0.24-1PGDG.rhel9.noarch.rpm) |
| `ddlx_13` | 0.23 | `el9.x86_64` | pgdg | 28.7 KiB | [ddlx_13-0.23-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ddlx_13-0.23-1.rhel9.noarch.rpm) |
| `ddlx_13` | 0.30 | `el9.aarch64` | pigsty | 30.7 KiB | [ddlx_13-0.30-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_13-0.30-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_13` | 0.30 | `el9.aarch64` | pgdg | 31.2 KiB | [ddlx_13-0.30-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ddlx_13-0.30-1PGDG.rhel9.noarch.rpm) |
| `ddlx_13` | 0.29 | `el9.aarch64` | pigsty | 29.8 KiB | [ddlx_13-0.29-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddlx_13-0.29-1PIGSTY.el9.aarch64.rpm) |
| `ddlx_13` | 0.29 | `el9.aarch64` | pgdg | 30.4 KiB | [ddlx_13-0.29-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ddlx_13-0.29-1PGDG.rhel9.noarch.rpm) |
| `ddlx_13` | 0.27 | `el9.aarch64` | pgdg | 29.8 KiB | [ddlx_13-0.27-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ddlx_13-0.27-1PGDG.rhel9.noarch.rpm) |
| `ddlx_13` | 0.26 | `el9.aarch64` | pgdg | 29.0 KiB | [ddlx_13-0.26-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ddlx_13-0.26-1PGDG.rhel9.noarch.rpm) |
| `ddlx_13` | 0.24 | `el9.aarch64` | pgdg | 28.6 KiB | [ddlx_13-0.24-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ddlx_13-0.24-1PGDG.rhel9.noarch.rpm) |
| `ddlx_13` | 0.23 | `el9.aarch64` | pgdg | 28.5 KiB | [ddlx_13-0.23-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ddlx_13-0.23-1.rhel9.noarch.rpm) |
| `ddlx_13` | 0.22 | `el9.aarch64` | pgdg | 28.8 KiB | [ddlx_13-0.22-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ddlx_13-0.22-1.rhel9.noarch.rpm) |
| `ddlx_13` | 0.30 | `el10.x86_64` | pgdg | 31.7 KiB | [ddlx_13-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/ddlx_13-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_13` | 0.29 | `el10.x86_64` | pgdg | 31.0 KiB | [ddlx_13-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/ddlx_13-0.29-1PGDG.rhel10.noarch.rpm) |
| `ddlx_13` | 0.30 | `el10.aarch64` | pgdg | 31.7 KiB | [ddlx_13-0.30-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/ddlx_13-0.30-1PGDG.rhel10.noarch.rpm) |
| `ddlx_13` | 0.29 | `el10.aarch64` | pgdg | 30.9 KiB | [ddlx_13-0.29-1PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/ddlx_13-0.29-1PGDG.rhel10.noarch.rpm) |
| `postgresql-13-ddlx` | 0.30 | `d12.x86_64` | pigsty | 28.3 KiB | [postgresql-13-ddlx_0.30-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-13-ddlx_0.30-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-ddlx` | 0.30 | `d12.aarch64` | pigsty | 28.3 KiB | [postgresql-13-ddlx_0.30-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddlx/postgresql-13-ddlx_0.30-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-ddlx` | 0.30 | `d13.x86_64` | pigsty | 28.3 KiB | [postgresql-13-ddlx_0.30-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-13-ddlx_0.30-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-ddlx` | 0.30 | `d13.aarch64` | pigsty | 28.3 KiB | [postgresql-13-ddlx_0.30-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddlx/postgresql-13-ddlx_0.30-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-ddlx` | 0.30 | `u22.x86_64` | pigsty | 25.5 KiB | [postgresql-13-ddlx_0.30-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-13-ddlx_0.30-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-ddlx` | 0.30 | `u22.aarch64` | pigsty | 25.5 KiB | [postgresql-13-ddlx_0.30-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddlx/postgresql-13-ddlx_0.30-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-ddlx` | 0.30 | `u24.x86_64` | pigsty | 25.4 KiB | [postgresql-13-ddlx_0.30-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-13-ddlx_0.30-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-ddlx` | 0.30 | `u24.aarch64` | pigsty | 25.4 KiB | [postgresql-13-ddlx_0.30-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddlx/postgresql-13-ddlx_0.30-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/lacanoid/pgddl" title="Repository" icon="github" subtitle="github.com/lacanoid/pgddl" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgddl-0.30.tar.gz" >}}
{{< /cards >}}


```bash
pig build get ddlx; # get ddlx source code
pig build dep ddlx; # install build dependencies
pig build pkg ddlx; # build extension rpm or deb
pig build ext ddlx; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install ddlx; # install by extension name, for the current active PG version
pig ext install pg_ddlx; # install via package alias, for the active PG version
pig ext install ddlx -v 18;   # install for PG 18
pig ext install ddlx -v 17;   # install for PG 17
pig ext install ddlx -v 16;   # install for PG 16
pig ext install ddlx -v 15;   # install for PG 15
pig ext install ddlx -v 14;   # install for PG 14
pig ext install ddlx -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION ddlx;
```

