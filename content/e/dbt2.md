---
title: "dbt2"
linkTitle: "dbt2"
description: "OSDL-DBT-2 test kit"
weight: 3220
categories: ["LANG"]
width: full
---

[**dbt2**](https://github.com/osdldbt/dbt2) : OSDL-DBT-2 test kit


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3220** | {{< badge content="dbt2" link="https://github.com/osdldbt/dbt2" >}} | {{< ext "dbt2" >}} | `0.61.7` | {{< category "LANG" >}} | {{< license "Artistic-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgtap" >}} {{< ext "faker" >}} {{< ext "plpgsql" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pllua" >}} {{< ext "hstore_pllua" >}} |

> [!Note] Package/source version 0.61.7; SQL extension version 0.45.0. This package contains the PostgreSQL stored-function extension, not the full DBT-2 benchmark toolchain.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.61.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `dbt2` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.61.7` | {{< bg "18" "dbt2-pg18-extensions" "green" >}} {{< bg "17" "dbt2-pg17-extensions" "green" >}} {{< bg "16" "dbt2-pg16-extensions" "green" >}} {{< bg "15" "dbt2-pg15-extensions" "green" >}} {{< bg "14" "dbt2-pg14-extensions" "green" >}} | `dbt2-pg$v-extensions` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.61.7` | {{< bg "18" "postgresql-18-dbt2" "green" >}} {{< bg "17" "postgresql-17-dbt2" "green" >}} {{< bg "16" "postgresql-16-dbt2" "green" >}} {{< bg "15" "postgresql-15-dbt2" "green" >}} {{< bg "14" "postgresql-14-dbt2" "green" >}} | `postgresql-$v-dbt2` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.61.7" "dbt2-pg18-extensions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "dbt2-pg17-extensions : AVAIL 1" "green" >}} | {{< bg "PGDG 0.53.7" "dbt2-pg16-extensions : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.53.7" "dbt2-pg15-extensions : AVAIL 5" "blue" >}} | {{< bg "PGDG 0.53.7" "dbt2-pg14-extensions : AVAIL 6" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.61.7" "dbt2-pg18-extensions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "dbt2-pg17-extensions : AVAIL 1" "green" >}} | {{< bg "PGDG 0.53.7" "dbt2-pg16-extensions : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.53.7" "dbt2-pg15-extensions : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.53.7" "dbt2-pg14-extensions : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg18-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg17-extensions : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg16-extensions : AVAIL 5" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg15-extensions : AVAIL 9" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg14-extensions : AVAIL 8" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg18-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg17-extensions : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg16-extensions : AVAIL 5" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg15-extensions : AVAIL 9" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg14-extensions : AVAIL 9" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg18-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg17-extensions : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg16-extensions : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg15-extensions : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg14-extensions : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg18-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg17-extensions : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg16-extensions : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg15-extensions : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg14-extensions : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-18-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-17-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-16-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-15-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-14-dbt2 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-18-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-17-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-16-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-15-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-14-dbt2 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-18-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-17-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-16-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-15-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-14-dbt2 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-18-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-17-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-16-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-15-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-14-dbt2 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-18-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-17-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-16-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-15-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-14-dbt2 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-18-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-17-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-16-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-15-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-14-dbt2 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-18-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-17-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-16-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-15-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-14-dbt2 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-18-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-17-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-16-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-15-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-14-dbt2 : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-18-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-17-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-16-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-15-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-14-dbt2 : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-18-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-17-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-16-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-15-dbt2 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.61.7" "postgresql-14-dbt2 : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg18-extensions` | `0.61.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.1 KiB | [dbt2-pg18-extensions-0.61.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/dbt2-pg18-extensions-0.61.7-1PIGSTY.el8.x86_64.rpm) |
| `dbt2-pg18-extensions` | `0.61.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.9 KiB | [dbt2-pg18-extensions-0.61.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/dbt2-pg18-extensions-0.61.7-1PIGSTY.el8.aarch64.rpm) |
| `dbt2-pg18-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.1 KiB | [dbt2-pg18-extensions-0.61.7-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/dbt2-pg18-extensions-0.61.7-5PGDG.rhel9.8.x86_64.rpm) |
| `dbt2-pg18-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [dbt2-pg18-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/dbt2-pg18-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg18-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [dbt2-pg18-extensions-0.61.7-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/dbt2-pg18-extensions-0.61.7-5PGDG.rhel9.8.aarch64.rpm) |
| `dbt2-pg18-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg18-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/dbt2-pg18-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg18-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.5 KiB | [dbt2-pg18-extensions-0.61.7-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/dbt2-pg18-extensions-0.61.7-5PGDG.rhel10.2.x86_64.rpm) |
| `dbt2-pg18-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg18-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/dbt2-pg18-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg18-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.0 KiB | [dbt2-pg18-extensions-0.61.7-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/dbt2-pg18-extensions-0.61.7-5PGDG.rhel10.2.aarch64.rpm) |
| `dbt2-pg18-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg18-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/dbt2-pg18-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-dbt2` | `0.61.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 178.0 KiB | [postgresql-18-dbt2_0.61.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/dbt2-extensions/postgresql-18-dbt2_0.61.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-dbt2` | `0.61.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 176.5 KiB | [postgresql-18-dbt2_0.61.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/dbt2-extensions/postgresql-18-dbt2_0.61.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-dbt2` | `0.61.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 177.8 KiB | [postgresql-18-dbt2_0.61.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/dbt2-extensions/postgresql-18-dbt2_0.61.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-dbt2` | `0.61.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 176.3 KiB | [postgresql-18-dbt2_0.61.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/dbt2-extensions/postgresql-18-dbt2_0.61.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-dbt2` | `0.61.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 201.4 KiB | [postgresql-18-dbt2_0.61.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/dbt2-extensions/postgresql-18-dbt2_0.61.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-dbt2` | `0.61.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 200.4 KiB | [postgresql-18-dbt2_0.61.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/dbt2-extensions/postgresql-18-dbt2_0.61.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-dbt2` | `0.61.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 194.4 KiB | [postgresql-18-dbt2_0.61.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/dbt2-extensions/postgresql-18-dbt2_0.61.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-dbt2` | `0.61.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 193.9 KiB | [postgresql-18-dbt2_0.61.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/dbt2-extensions/postgresql-18-dbt2_0.61.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-dbt2` | `0.61.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 196.7 KiB | [postgresql-18-dbt2_0.61.7-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/dbt2-extensions/postgresql-18-dbt2_0.61.7-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-dbt2` | `0.61.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 196.5 KiB | [postgresql-18-dbt2_0.61.7-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/dbt2-extensions/postgresql-18-dbt2_0.61.7-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg17-extensions` | `0.61.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.1 KiB | [dbt2-pg17-extensions-0.61.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/dbt2-pg17-extensions-0.61.7-1PIGSTY.el8.x86_64.rpm) |
| `dbt2-pg17-extensions` | `0.61.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.9 KiB | [dbt2-pg17-extensions-0.61.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/dbt2-pg17-extensions-0.61.7-1PIGSTY.el8.aarch64.rpm) |
| `dbt2-pg17-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.1 KiB | [dbt2-pg17-extensions-0.61.7-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/dbt2-pg17-extensions-0.61.7-5PGDG.rhel9.8.x86_64.rpm) |
| `dbt2-pg17-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [dbt2-pg17-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/dbt2-pg17-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg17-extensions` | `0.61.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg17-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/dbt2-pg17-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg17-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [dbt2-pg17-extensions-0.61.7-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/dbt2-pg17-extensions-0.61.7-5PGDG.rhel9.8.aarch64.rpm) |
| `dbt2-pg17-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg17-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/dbt2-pg17-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg17-extensions` | `0.61.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg17-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/dbt2-pg17-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg17-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.5 KiB | [dbt2-pg17-extensions-0.61.7-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/dbt2-pg17-extensions-0.61.7-5PGDG.rhel10.2.x86_64.rpm) |
| `dbt2-pg17-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg17-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/dbt2-pg17-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg17-extensions` | `0.61.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg17-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/dbt2-pg17-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg17-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.0 KiB | [dbt2-pg17-extensions-0.61.7-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/dbt2-pg17-extensions-0.61.7-5PGDG.rhel10.2.aarch64.rpm) |
| `dbt2-pg17-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg17-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/dbt2-pg17-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm) |
| `dbt2-pg17-extensions` | `0.61.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.3 KiB | [dbt2-pg17-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/dbt2-pg17-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-dbt2` | `0.61.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 177.7 KiB | [postgresql-17-dbt2_0.61.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/dbt2-extensions/postgresql-17-dbt2_0.61.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-dbt2` | `0.61.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 176.4 KiB | [postgresql-17-dbt2_0.61.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/dbt2-extensions/postgresql-17-dbt2_0.61.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-dbt2` | `0.61.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 177.4 KiB | [postgresql-17-dbt2_0.61.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/dbt2-extensions/postgresql-17-dbt2_0.61.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-dbt2` | `0.61.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 176.4 KiB | [postgresql-17-dbt2_0.61.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/dbt2-extensions/postgresql-17-dbt2_0.61.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-dbt2` | `0.61.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 215.0 KiB | [postgresql-17-dbt2_0.61.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/dbt2-extensions/postgresql-17-dbt2_0.61.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-dbt2` | `0.61.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 214.0 KiB | [postgresql-17-dbt2_0.61.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/dbt2-extensions/postgresql-17-dbt2_0.61.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-dbt2` | `0.61.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 193.4 KiB | [postgresql-17-dbt2_0.61.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/dbt2-extensions/postgresql-17-dbt2_0.61.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-dbt2` | `0.61.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 192.8 KiB | [postgresql-17-dbt2_0.61.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/dbt2-extensions/postgresql-17-dbt2_0.61.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-dbt2` | `0.61.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 193.6 KiB | [postgresql-17-dbt2_0.61.7-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/dbt2-extensions/postgresql-17-dbt2_0.61.7-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-dbt2` | `0.61.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 193.1 KiB | [postgresql-17-dbt2_0.61.7-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/dbt2-extensions/postgresql-17-dbt2_0.61.7-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg16-extensions` | `0.53.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.9 KiB | [dbt2-pg16-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/dbt2-pg16-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.53.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [dbt2-pg16-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/dbt2-pg16-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.1 KiB | [dbt2-pg16-extensions-0.61.7-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.61.7-5PGDG.rhel9.8.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [dbt2-pg16-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.61.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg16-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.53.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [dbt2-pg16-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.53.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.4 KiB | [dbt2-pg16-extensions-0.53.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.53.6-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [dbt2-pg16-extensions-0.61.7-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.61.7-5PGDG.rhel9.8.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg16-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.61.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg16-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.53.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.0 KiB | [dbt2-pg16-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.53.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [dbt2-pg16-extensions-0.53.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.53.6-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.5 KiB | [dbt2-pg16-extensions-0.61.7-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/dbt2-pg16-extensions-0.61.7-5PGDG.rhel10.2.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg16-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/dbt2-pg16-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.61.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg16-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/dbt2-pg16-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.0 KiB | [dbt2-pg16-extensions-0.61.7-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/dbt2-pg16-extensions-0.61.7-5PGDG.rhel10.2.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg16-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/dbt2-pg16-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.61.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.3 KiB | [dbt2-pg16-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/dbt2-pg16-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-dbt2` | `0.61.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 164.0 KiB | [postgresql-16-dbt2_0.61.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/dbt2-extensions/postgresql-16-dbt2_0.61.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-dbt2` | `0.61.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 162.8 KiB | [postgresql-16-dbt2_0.61.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/dbt2-extensions/postgresql-16-dbt2_0.61.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-dbt2` | `0.61.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 164.1 KiB | [postgresql-16-dbt2_0.61.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/dbt2-extensions/postgresql-16-dbt2_0.61.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-dbt2` | `0.61.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 163.2 KiB | [postgresql-16-dbt2_0.61.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/dbt2-extensions/postgresql-16-dbt2_0.61.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-dbt2` | `0.61.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 201.3 KiB | [postgresql-16-dbt2_0.61.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/dbt2-extensions/postgresql-16-dbt2_0.61.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-dbt2` | `0.61.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 200.3 KiB | [postgresql-16-dbt2_0.61.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/dbt2-extensions/postgresql-16-dbt2_0.61.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-dbt2` | `0.61.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 178.6 KiB | [postgresql-16-dbt2_0.61.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/dbt2-extensions/postgresql-16-dbt2_0.61.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-dbt2` | `0.61.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 177.9 KiB | [postgresql-16-dbt2_0.61.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/dbt2-extensions/postgresql-16-dbt2_0.61.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-dbt2` | `0.61.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 180.1 KiB | [postgresql-16-dbt2_0.61.7-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/dbt2-extensions/postgresql-16-dbt2_0.61.7-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-dbt2` | `0.61.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 179.4 KiB | [postgresql-16-dbt2_0.61.7-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/dbt2-extensions/postgresql-16-dbt2_0.61.7-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg15-extensions` | `0.53.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.9 KiB | [dbt2-pg15-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.53.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.8 KiB | [dbt2-pg15-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.50.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.50.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.50.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.49.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.49.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.49.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.48.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.48.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.48.7-1.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.53.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [dbt2-pg15-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/dbt2-pg15-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.53.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/dbt2-pg15-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.50.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [dbt2-pg15-extensions-0.50.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/dbt2-pg15-extensions-0.50.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.49.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [dbt2-pg15-extensions-0.49.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/dbt2-pg15-extensions-0.49.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.1 KiB | [dbt2-pg15-extensions-0.61.7-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.61.7-5PGDG.rhel9.8.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.61.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.53.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [dbt2-pg15-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.53.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [dbt2-pg15-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.50.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.1 KiB | [dbt2-pg15-extensions-0.50.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.50.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.49.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.49.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.49.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.48.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.48.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.48.7-1.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.48.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [dbt2-pg15-extensions-0.48.3-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.48.3-2.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [dbt2-pg15-extensions-0.61.7-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.61.7-5PGDG.rhel9.8.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.6 KiB | [dbt2-pg15-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.61.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg15-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.53.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [dbt2-pg15-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.53.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [dbt2-pg15-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.50.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.50.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.50.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.49.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.49.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.49.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.48.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.4 KiB | [dbt2-pg15-extensions-0.48.7-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.48.7-1.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.48.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.3 KiB | [dbt2-pg15-extensions-0.48.3-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.48.3-2.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.4 KiB | [dbt2-pg15-extensions-0.61.7-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/dbt2-pg15-extensions-0.61.7-5PGDG.rhel10.2.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg15-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/dbt2-pg15-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.61.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg15-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/dbt2-pg15-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.61.7-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/dbt2-pg15-extensions-0.61.7-5PGDG.rhel10.2.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg15-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/dbt2-pg15-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.61.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg15-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/dbt2-pg15-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-dbt2` | `0.61.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 158.0 KiB | [postgresql-15-dbt2_0.61.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/dbt2-extensions/postgresql-15-dbt2_0.61.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-dbt2` | `0.61.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 156.7 KiB | [postgresql-15-dbt2_0.61.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/dbt2-extensions/postgresql-15-dbt2_0.61.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-dbt2` | `0.61.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 157.8 KiB | [postgresql-15-dbt2_0.61.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/dbt2-extensions/postgresql-15-dbt2_0.61.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-dbt2` | `0.61.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 156.9 KiB | [postgresql-15-dbt2_0.61.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/dbt2-extensions/postgresql-15-dbt2_0.61.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-dbt2` | `0.61.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 195.2 KiB | [postgresql-15-dbt2_0.61.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/dbt2-extensions/postgresql-15-dbt2_0.61.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-dbt2` | `0.61.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 194.2 KiB | [postgresql-15-dbt2_0.61.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/dbt2-extensions/postgresql-15-dbt2_0.61.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-dbt2` | `0.61.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 172.2 KiB | [postgresql-15-dbt2_0.61.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/dbt2-extensions/postgresql-15-dbt2_0.61.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-dbt2` | `0.61.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 171.3 KiB | [postgresql-15-dbt2_0.61.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/dbt2-extensions/postgresql-15-dbt2_0.61.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-dbt2` | `0.61.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 175.3 KiB | [postgresql-15-dbt2_0.61.7-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/dbt2-extensions/postgresql-15-dbt2_0.61.7-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-dbt2` | `0.61.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 174.5 KiB | [postgresql-15-dbt2_0.61.7-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/dbt2-extensions/postgresql-15-dbt2_0.61.7-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg14-extensions` | `0.53.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.9 KiB | [dbt2-pg14-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.53.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.7 KiB | [dbt2-pg14-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.50.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.5 KiB | [dbt2-pg14-extensions-0.50.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.50.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.49.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.5 KiB | [dbt2-pg14-extensions-0.49.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.49.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.48.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.4 KiB | [dbt2-pg14-extensions-0.48.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.48.7-1.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.48.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.3 KiB | [dbt2-pg14-extensions-0.48.3-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.48.3-2.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.53.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [dbt2-pg14-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/dbt2-pg14-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.53.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.4 KiB | [dbt2-pg14-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/dbt2-pg14-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.50.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [dbt2-pg14-extensions-0.50.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/dbt2-pg14-extensions-0.50.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.49.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [dbt2-pg14-extensions-0.49.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/dbt2-pg14-extensions-0.49.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.1 KiB | [dbt2-pg14-extensions-0.61.7-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.61.7-5PGDG.rhel9.8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [dbt2-pg14-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.61.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg14-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.53.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [dbt2-pg14-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.53.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [dbt2-pg14-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.50.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.1 KiB | [dbt2-pg14-extensions-0.50.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.50.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.49.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg14-extensions-0.49.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.49.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.48.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg14-extensions-0.48.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.48.7-1.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [dbt2-pg14-extensions-0.61.7-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.61.7-5PGDG.rhel9.8.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.6 KiB | [dbt2-pg14-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.61.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg14-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.53.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [dbt2-pg14-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.53.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [dbt2-pg14-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.50.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.5 KiB | [dbt2-pg14-extensions-0.50.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.50.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.49.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.5 KiB | [dbt2-pg14-extensions-0.49.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.49.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.48.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.4 KiB | [dbt2-pg14-extensions-0.48.7-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.48.7-1.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.48.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.3 KiB | [dbt2-pg14-extensions-0.48.3-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.48.3-2.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.4 KiB | [dbt2-pg14-extensions-0.61.7-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/dbt2-pg14-extensions-0.61.7-5PGDG.rhel10.2.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.6 KiB | [dbt2-pg14-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/dbt2-pg14-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.61.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg14-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/dbt2-pg14-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.0 KiB | [dbt2-pg14-extensions-0.61.7-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/dbt2-pg14-extensions-0.61.7-5PGDG.rhel10.2.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg14-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/dbt2-pg14-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.61.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg14-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/dbt2-pg14-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-dbt2` | `0.61.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 139.1 KiB | [postgresql-14-dbt2_0.61.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/dbt2-extensions/postgresql-14-dbt2_0.61.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-dbt2` | `0.61.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 137.6 KiB | [postgresql-14-dbt2_0.61.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/dbt2-extensions/postgresql-14-dbt2_0.61.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-dbt2` | `0.61.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 138.6 KiB | [postgresql-14-dbt2_0.61.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/dbt2-extensions/postgresql-14-dbt2_0.61.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-dbt2` | `0.61.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 137.8 KiB | [postgresql-14-dbt2_0.61.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/dbt2-extensions/postgresql-14-dbt2_0.61.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-dbt2` | `0.61.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 166.0 KiB | [postgresql-14-dbt2_0.61.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/dbt2-extensions/postgresql-14-dbt2_0.61.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-dbt2` | `0.61.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 164.9 KiB | [postgresql-14-dbt2_0.61.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/dbt2-extensions/postgresql-14-dbt2_0.61.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-dbt2` | `0.61.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 151.5 KiB | [postgresql-14-dbt2_0.61.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/dbt2-extensions/postgresql-14-dbt2_0.61.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-dbt2` | `0.61.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 150.7 KiB | [postgresql-14-dbt2_0.61.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/dbt2-extensions/postgresql-14-dbt2_0.61.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-dbt2` | `0.61.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 155.4 KiB | [postgresql-14-dbt2_0.61.7-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/dbt2-extensions/postgresql-14-dbt2_0.61.7-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-dbt2` | `0.61.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 154.8 KiB | [postgresql-14-dbt2_0.61.7-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/dbt2-extensions/postgresql-14-dbt2_0.61.7-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/osdldbt/dbt2" title="Repository" icon="github" subtitle="github.com/osdldbt/dbt2" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="dbt2-0.61.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg dbt2;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install dbt2;		# install via package name, for the active PG version

pig install dbt2 -v 18;   # install for PG 18
pig install dbt2 -v 17;   # install for PG 17
pig install dbt2 -v 16;   # install for PG 16
pig install dbt2 -v 15;   # install for PG 15
pig install dbt2 -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION dbt2;
```




## Usage

> [dbt2: OSDL-DBT-2 test kit](https://github.com/osdldbt/dbt2)

`dbt2` is a TPC-C benchmark implementation for PostgreSQL. The extension provides stored procedures that implement the five standard TPC-C transaction types.

```sql
CREATE EXTENSION dbt2;
```

### TPC-C Transaction Types

The extension provides stored procedures for the five standard TPC-C transactions:

- **New Order**: Creates a new order with multiple line items, updating stock levels
- **Payment**: Processes a customer payment, updating warehouse and district balances
- **Order Status**: Retrieves the status of a customer's most recent order
- **Delivery**: Processes pending orders for delivery across all districts
- **Stock Level**: Checks the count of recently sold items with low stock

### Benchmark Workflow

The `dbt2` system consists of:

1. **Database extension** (`dbt2`): Stored procedures for TPC-C transactions
2. **Data loader**: Populates the benchmark tables with TPC-C data
3. **Driver**: Generates transaction workloads simulating terminal users
4. **Client**: Manages connections between the driver and database

### Running Benchmarks

The benchmark is typically run using the `dbt2` command-line tools (separate from the extension):

```bash
# Build the benchmark database
dbt2 build --dbms pgsql --warehouses 10

# Run the benchmark
dbt2 run --dbms pgsql --warehouses 10 --duration 300 --connections 10

# Generate report
dbt2 report --dbms pgsql
```

### TPC-C Schema

The benchmark uses these standard tables: `warehouse`, `district`, `customer`, `history`, `new_order`, `orders`, `order_line`, `item`, and `stock`.

Refer to the `doc/` directory in the repository for detailed configuration and tuning options.
